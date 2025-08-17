package auction

import "time"

// Auction document
type Auction struct {
    ID            string    `bson:"_id,omitempty" json:"id"`
    Seller        string    `bson:"seller" json:"seller"`
    Item          string    `bson:"item" json:"item"`
    MinPrice      float64   `bson:"minPrice" json:"minPrice"`
    HighestBid    float64   `bson:"highestBid" json:"highestBid"`
    HighestBidder string    `bson:"highestBidder,omitempty" json:"highestBidder,omitempty"`
    Deadline      time.Time `bson:"deadline" json:"deadline"`
    CreatedAt     time.Time `bson:"createdAt" json:"createdAt"`
    Status        string    `bson:"status" json:"status"`
}


// Bid document
type Bid struct {
    ID        string    `bson:"_id,omitempty"`
    AuctionID string    `bson:"auctionId"`
    Bidder    string    `bson:"bidder"`
    Amount    float64   `bson:"amount"`
    Timestamp time.Time `bson:"timestamp"`
}
