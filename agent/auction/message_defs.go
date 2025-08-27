package auction

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuctionObject struct {
	AuctionId  string    `json:"auctionid"`  // GUID for an auction
	OrderId    string    `json:"orderid"`    // GUID for the order
	StartDate  time.Time `json:"startdate"`  // Start date of an auction
	EndDate    time.Time `json:"enddate"`    // End date of an auction
	WinningBid string    `json:"winningbid"` // Winning bid of an auction
	Subscribed bool      `json:"subscribed"` // Indicates whether or not the GUID asset has been subscribed to
	Cloned     bool      `json:"cloned"`     // Indicates whether or not the GUID asset has been cloned locally
}

type AuctionWinner struct {
	AuctionID string `json:"auctionid"` // GUID for the auction
	BidID     string `json:"bidid"`     // GUID for the bid
}

type AuctionEnd struct {
	AuctionID string    `json:"auctionid"` // GUID for the auction
	EndDate   time.Time `json:"enddate"`   // end date
}

type BidObject struct {
	BidId        string    `json:"bidid"`        // GUID for the bid
	AuctionId    string    `json:"auctionid"`    // GUID for an auction
	Price        string    `json:"price"`        // Target price
	Quantity     int       `json:"quantity"`     // Quantity requested
	DeliveryDate time.Time `json:"deliverydate"` // Delivery date
	Onion        string    `json:"onion"`        // Onion address of the bid submitter
	ResponseDate time.Time `json:"responsedate"` // The time the bid was submitted
}

type BidResponse struct {
	AuctionID    string    `json:"auctionid"`    // GUID for the auction
	BidId        string    `json:"bidid"`        // GUID for the bid
	ResponseDate time.Time `json:"responsedate"` // end date
}

// *** These need to be removed *** //

// Received from inventoryAgent
type OrderObject struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"` // id of the meta data object on MongoDB
	OrderId      string             `json:"orderid"`                 // GUID for the order
	Entity       string             `json:"entity"`                  // Onion address of the entity managing the object
	Price        string             `json:"price"`                   // Target price
	Quantity     int                `json:"quantity"`                // Quantity requested
	DeliveryDate time.Time          `json:"deliverydate"`            // Delivery date
	OrderItem    OrderItemObject    `json:"orderitem"`               // Item for auction
}

type OrderItemObject struct {
	GUID          string         `json:"guid"`          // GUID for data object
	HASH          string         `json:"hash"`          // SHA-256 of the data object -> dsig?
	URI           string         `json:"uri"`           // Location where the data object is stored
	Description   string         `json:"description"`   // Brief description of the data object
	Version       string         `json:"version"`       // Version number
	Relationships []Relationship `json:"relationships"` // Relationship list
	Timestamp     time.Time      `json:"timestamp"`     // When the object was created
	OrderedDate   time.Time      `json:"ordereddate"`   // When the object was ordered
}

type Relationship struct {
	Entity    string    `json:"entity"`    // Onion address of the entity managing the data object
	GUID      string    `json:"guid"`      // GUID for managed data object
	Version   string    `json:"version"`   // Version number
	Type      string    `json:"type"`      // Type = [Parent, ]
	Timestamp time.Time `json:"timestamp"` // When the relationship was established
}
