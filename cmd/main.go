package main

import (
	"auction/database"
	"auction/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main(){

	r :=gin.Default()

	
	db, err :=database.ConnectMongoDB()

	 if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Disconnect()

	h :=handler.NewHandler(db)

	r.POST("/auction",h.AuctionAdd)
    
    log.Println("Auction application started!")

	r.Run(":8080")


}
