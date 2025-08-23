package handler

import (
	"auction"
	"auction/database"
	"context"
	"fmt"
	"log"
	
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	
)

type Handler struct{
	DB *database.Database
}

func NewHandler( db *database.Database) *Handler{
	return &Handler{
		DB: db,
	}
}

func (h *Handler)AuctionAdd(c *gin.Context){
	var auction auction.Auction

	err:=c.ShouldBindJSON(&auction)
	if err!=nil{
		log.Fatalln("error binding json",err.Error())
	}

	ctx , cancel := context.WithTimeout(context.Background(),5*time.Second)
	
	defer cancel()

	result , err := h.DB.DB.Collection("auctions").InsertOne(ctx,auction)
	if err != nil {
		log.Println("error inserting auction:", err.Error())
		c.JSON(500, gin.H{"error": "failed to insert auction"})
		return
	}

	c.JSON(200, gin.H{
		"message": "auction created",
		"id":      result.InsertedID,
	})
}

func (h *Handler) GetBlockchainStatus(c *gin.Context){

	if h.DB.BlockchainClient == nil{
		c.JSON(200,gin.H{
			"blockchain_connected": false,
            "message": "Blockchain client not initialized",
		})
		return
	}

	blockNumber , err := h.DB.BlockchainClient.Client.BlockNumber(context.Background())
	if err != nil {
        c.JSON(500, gin.H{
            "blockchain_connected": false,
            "error": err.Error(),
        })
        return
    }


	

    c.JSON(200, gin.H{
        "blockchain_connected": true,
        "chain_id":           h.DB.BlockchainClient.ChainID.String(),
        "latest_block":       blockNumber,
		"contract_address":		h.DB.BlockchainClient.ContractAddress.Hex(),
		"contract_connected":	true,
	
        "message":           "Blockchain connection healthy",
    })
}

func (h *Handler) GetUserBalance(c *gin.Context){
	address :=c.Param("address")

	    if h.DB.BlockchainClient == nil || h.DB.BlockchainClient.AuctionVault == nil {
        c.JSON(500, gin.H{
            "error": "Contract not connected",
        })
        return
    }

	userAddress := common.HexToAddress(address)

	balance , err :=h.DB.BlockchainClient.AuctionVault.Balances(nil,userAddress)
	if err != nil {
			c.JSON(500, gin.H{
				"error": fmt.Sprintf("Failed to get balance: %v", err),
			})
			return
		}

	 c.JSON(200, gin.H{
        "address": address,
        "balance": balance.String(), // Convert big.Int to string
        "balance_eth": fmt.Sprintf("%.6f", float64(balance.Int64())/1e18), // Convert wei to ETH
    })

}