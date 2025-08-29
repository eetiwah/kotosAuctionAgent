package auction

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func Messages(data string, conversationID int, onion string) {
	cmdList := strings.Split(data, " ")

	/*
		var groupMsg group.GroupMessage

		// convert string to []byte
		jsonData := []byte(data)

		// Unmarshal string to JSON cmd structure
		if err := json.Unmarshal(jsonData, &groupMsg); err != nil {
			msg := fmt.Sprintf("Error: unmarshalling failure = %v\n", err)
			log.Println(msg)
			return
		}
			// Marshal Data back to JSON and unmarshal into the correct struct
			dataBytes, err := json.Marshal(groupMsg.Data)
			if err != nil {
				msg := fmt.Sprintf("Error: failed to marshal data: %v\n", err)
				log.Println(msg)
				return
			}
	*/

	// Determine auction message type
	//switch strings.ToLower(groupMsg.Type) {
	switch strings.ToLower(cmdList[0]) {

	case "ping_auction":
		log.Println("ping_auction received")

	case "bid_offer":
		//log.Printf("dataBytes = %s", string(dataBytes))
		log.Println("bid_offer received")

		log.Printf("Data = %s\n", cmdList[1])

		var bidObject BidObject
		if err := json.Unmarshal([]byte(cmdList[1]), &bidObject); err != nil {
			msg := fmt.Sprintf("Error: unmarshalling failure = %v\n", err)
			log.Println(msg)
			return
		}

		// Add onion address so that we can track who did what
		bidObject.Onion = onion

		// Add bid to bid store
		err := AddBid(bidObject)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

	default:
		//log.Printf("Auction MessageType error: %v from %d", groupMsg.Type, conversationID)
		log.Printf("Auction MessageType error: %v from %d", cmdList[0], conversationID)
	}
}
