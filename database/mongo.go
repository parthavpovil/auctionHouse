package database

import (
	"auction/blockchain"
	"context"
	"fmt"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
	DB *mongo.Database
	AuctionCollection *mongo.Collection
	BidsCollection *mongo.Collection
	BlockchainClient *blockchain.BlockchainClient
}

func ConnectMongoDB() (*Database, error){
	uri :="mongodb://localhost:27017"

	clientOptions :=options.Client().ApplyURI(uri)

	ctx , cancel :=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	client , err := mongo.Connect(ctx,clientOptions)
	if err != nil {
        return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
    }
    
	err =client.Ping(ctx,nil)
	if err != nil {
        return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
    }

	

	 db :=client.Database("auction_db")

	 auctionCollection := db.Collection("auctions")
	 bidsCollection	:= db.Collection("bids")

	  fmt.Println("Connected to MongoDB!",db.Name())
	  fmt.Println("Created collections",auctionCollection.Name())
	  fmt.Println("Created collection:",bidsCollection.Name())

	  blockchainClient , err := blockchain.NewBlockchainClient()

	  if err !=nil{
		fmt.Printf("⚠️  Warning: Could not connect to blockchain: %v\n", err)
        fmt.Println("🔄 Continuing without blockchain functionality...")
        
        return &Database{
            Client:            client,
            DB:                db,
            AuctionCollection: auctionCollection,
            BidsCollection:    bidsCollection,
            BlockchainClient:  nil, // No blockchain client
        }, nil

	  }

	  // Test the blockchain connection
    err = blockchainClient.TestConnection()
    if err != nil {
        fmt.Printf("⚠️  Warning: Blockchain connection test failed: %v\n", err)
    }


	 return &Database{
		Client: client,
		DB: db,
		AuctionCollection: auctionCollection,
		BidsCollection: bidsCollection,
		BlockchainClient: blockchainClient,
	 },nil
}

func (d *Database) Disconnect() error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    return d.Client.Disconnect(ctx)
}