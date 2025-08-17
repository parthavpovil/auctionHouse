package handler

import (
	"auction"
	"auction/database"
	"context"
	"log"
	"time"

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