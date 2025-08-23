package main

import (
	"auction/database"
	"auction/handler"
	"log"
	"github.com/joho/godotenv"


	"github.com/gin-gonic/gin"
)

func main(){


	 err := godotenv.Load("../.env")
    if err != nil {
        log.Println("Warning: No .env file found or error loading it")
	}
	r :=gin.Default()

	
	db, err :=database.ConnectMongoDB()

	 if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Disconnect()

	h :=handler.NewHandler(db)

	r.POST("/auction",h.AuctionAdd)
	r.GET("/blockchain/status", h.GetBlockchainStatus)
	r.GET("/balance/:address", h.GetUserBalance)
    
    log.Println("Auction application started!")
	

	r.Run(":8080")


}
