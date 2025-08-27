package auction

import (
	"log"
	"time"
)

func Watcher() {
	// Set tick interval
	tickInterval := time.Duration(10) * time.Second

	// Create a ticker that ticks every tickInterval
	ticker := time.NewTicker(tickInterval)

	// Define the function to be executed on each tick
	tickerFunc := func() error {

		/*

			err := StartAuctions()
			if err != nil {
				log.Printf("Warning: StartAuctions = %s", err)
			}

			err = StopAuctions()
			if err != nil {
				log.Printf("Warning: StopAuctions = %s", err)
			}

			err = SelectAuctionWinners()
			if err != nil {
				log.Printf("Warning: SelectionAuctionWinners = %s", err)
			}
		*/

		return nil
	}

	// Create start message
	log.Println("auctionWatcher started ...")

	// Create channels for completion and errors
	done := make(chan struct{})
	errChan := make(chan string, 1)

	// Start a goroutine to execute the function on each tick
	go func() {
		defer close(done)   // Ensure done is closed when the goroutine exits
		defer ticker.Stop() // Ensure ticker is stopped when the goroutine exits

		for range ticker.C {
			// Execute the function on each tick
			if err := tickerFunc(); err != nil {
				errChan <- "Watcher Error:  " + err.Error()
				return
			}
		}
	}()

	// Wait for the goroutine to complete or fail
	select {
	case <-done:
		// Check if an error occurred
		select {
		case errMsg := <-errChan:
			log.Printf("Error = %v", errMsg)

		default:
			log.Println("Inventory check completed")
		}
	case errMsg := <-errChan:
		log.Printf("Error = %v", errMsg)
	}
}

/*
func StartAuctions() error {
	// Get a list of auction IDs that need to be started
	startAuctionList, err := GetStartAuctionList()
	if err != nil {
		return err
	}

	// Check to determine if there are any auctions that need to be started
	if len(startAuctionList) == 0 {
		log.Println("No auctions need to be started")
		return nil
	}

	// Start the auctions
	for _, auction := range startAuctionList {
		// Update startDate to now
		startDate := time.Now().UTC()
		auction.StartDate = startDate

		// Get the order associated with the auction from mongoDB

		// placeholder, needs to be re-factored
		var order OrderObject

		// Send order object to the auction community (full order object)
		err = Add_Order_Event_Send(order)
		if err != nil {
			log.Printf("Add_Order_Event_Send failed: %v\n", err)
			return err
		}

		// Send auction start event to the auction community (full auction object)
		err = Auction_Start_Event_Send(auction)
		if err != nil {
			log.Printf("Auction_Start_Event_Send failed: %v\n", err)
			return err
		}

		// Update the auction obj in MongoDB to indicate that the auction has been started
		err = SetAuctionStarted(auction.AuctionId, startDate)
		if err != nil {
			log.Printf("setAuctionStarted failed: %v\n", err)
			return err
		}

		log.Printf("Auction: %s was started @ %v\n", auction.AuctionId, startDate)
	}

	return nil
}
*/

/*
func StopAuctions() error {
	// This is the duration that an auction will run for, in secs
	auction_duration := 60

	// Check to determine if there are any auctions that need to be stoppped
	stopAuctionList, err := GetStopAuctionList(auction_duration) // This returns a list of auctions that need to be stopped
	if err != nil {
		return err
	}

	if len(stopAuctionList) == 0 {
		log.Println("No auctions need to be stopped")
		return nil
	}

	for _, auction := range stopAuctionList {
		endDate := time.Now().UTC()
		err := Auction_Stop_Event_Send(auction.AuctionId, endDate) // stop an auctionId
		if err != nil {
			return err
		}

		// Update the object stored in MongoDB
		err = SetAuctionStopped(auction.AuctionId, endDate)
		if err != nil {
			return err
		}

		log.Printf("Auction: %s was stopped @ %v\n", auction.AuctionId, endDate)
	}

	return nil
}


func SelectAuctionWinners() error {
	// Get a list of all the auctions that have been closed are awaiting a winner selection
	awaitingWinnerList, err := GetAwaitingWinnerList()
	if err != nil {
		return err
	}

	// Check to determine if there are any auctions that need to be started
	if len(awaitingWinnerList) == 0 {
		log.Println("No auctions need to have a winner selected")
		return nil
	}

	log.Printf("Number of awaitingWinnerList entries = %d", len(awaitingWinnerList))

	// Determine auctions for the list
	for _, auction := range awaitingWinnerList {
		// Get the bidList for AuctionId
		bidList, err := GetBidList(auction.AuctionId)
		if err != nil {
			return err
		}

		if len(bidList) == 0 {
			return errors.New("bidList was empty")
		}

		// Determine winner
		bidId, err := GetAuctionWinner(auction, bidList)
		if err != nil {
			return err
		}

		log.Printf("Winning bid: %s for Auction: %s\n", bidId, auction.AuctionId)

		// Update auction with the winning bidId
		err = SetAuctionWinner(auction.AuctionId, bidId)
		if err != nil {
			return err
		}

		// Announce auction winner to community
		err = Auction_Winner_Event_Send(auction.AuctionId, bidId)
		if err != nil {
			return err
		}
	}

	return nil
}
*/

/*
func GetAuctionWinner(auction AuctionObject, bidList []BidObject) (string, error) {
	//winningbidID := ""

		// Get the order
		order, err := GetOrder(auction.OrderId)
		if err != nil {
			return winningbidID, err
		}

	return "123123", nil

		for _, bid := range bidList {
			if order.Price == bid.Price &&
				order.Quantity == bid.Quantity &&
				order.DeliveryDate.Equal(bid.DeliveryDate) {
				return bid.BidId, nil
			}
		}

		return winningbidID, errors.New("no bids matched auction params")
}
*/
