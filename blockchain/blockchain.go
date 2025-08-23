package blockchain

import (
	"auction/contracts"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)


type BlockchainClient struct {
	Client *ethclient.Client
	ContractAddress common.Address
	ChainID *big.Int
	AuctionVault *contracts.AuctionVault

}

func  NewBlockchainClient() (*BlockchainClient , error){
	rpcURL :=os.Getenv("ETH_RPC_URL")
	if rpcURL == ""{
		rpcURL = "http://localhost:8545"
	}

	contractAddr :=os.Getenv("CONTRACT_ADDRESS")
	
	client , err := ethclient.Dial(rpcURL)
	   if err != nil {
        return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
    }
	chainID , err := client.NetworkID(context.Background())
	 if err != nil {
        return nil, fmt.Errorf("failed to get network ID: %v", err)
    }

	log.Printf("Connected to blockchain - Chain ID: %v, RPC: %s", chainID, rpcURL)

	
	
	

	bc := &BlockchainClient{
        Client:  client,
        ChainID: chainID,
    }

	bc.ContractAddress =common.HexToAddress(contractAddr)

	// After setting the contract address, add this:
	if contractAddr != "" {
    	bc.ContractAddress = common.HexToAddress(contractAddr)
		
		// Create contract instance
		auctionVault, err := contracts.NewAuctionVault(bc.ContractAddress, client)
		if err != nil {
			return nil, fmt.Errorf("failed to create contract instance: %v", err)
		}
		bc.AuctionVault = auctionVault
		
		log.Printf("contract address: %s", contractAddr)
		log.Printf("✅ Contract instance created successfully")
	}
	


	return bc,nil

}

func (bc *BlockchainClient) TestConnection() error {
    blockNumber, err := bc.Client.BlockNumber(context.Background())
    if err != nil {
        return fmt.Errorf("failed to get block number: %v", err)
    }
    
    log.Printf("✅ Blockchain connected! Latest block: %d", blockNumber)
    return nil
}