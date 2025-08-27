package auction

import (
	"encoding/json"
	"fmt"
	"kotosAuctionAgent/agent/group"
)

// Auction management messages from admin

func Create(commandList []string) string {
	switch len(commandList) {
	case 1:
		return "Error: missing "

	case 2:
		if commandList[1] == "-help" {
			return "usage: create_auction auction_serialized_json_data"
		}

		// Send message to community
		err := group.SendMessage([]byte(commandList[1]))
		if err != nil {
			return fmt.Sprintf("Error: create_auction msg was not sent: %v", err)
		}

		// Update auction store
		err = CreateAuctionObj([]byte(commandList[1]))
		if err != nil {
			return err.Error()
		}

		// Success
		return "Auction created and sent to community"

	default:
		return "Error: parameter mismatch"
	}
}

func Get(commandList []string) string {
	switch len(commandList) {
	case 1:
		return "Error: missing "

	case 2:
		if commandList[1] == "-help" {
			return "usage: get_auction auctionId"
		}

		obj, err := GetAuctionObj(commandList[1])
		if err != nil {
			return err.Error()
		}

		jsonStr, err := json.Marshal(obj)
		if err != nil {
			return fmt.Sprintf("Error: get_auction json marshal: %v", err)
		}
		return fmt.Sprintf("Auction: %s", jsonStr)

	default:
		return "Error: parameter mismatch"
	}
}

func List(commandList []string) string {
	switch len(commandList) {
	case 1:
		return "Error: missing "

	case 2:
		if commandList[1] == "-help" {
			return "usage: get_auction_list"
		}

		objList, err := GetAuctionList()
		if err != nil {
			return err.Error()
		}

		jsonStr, err := json.Marshal(objList)
		if err != nil {
			return fmt.Sprintf("Error: get_auction_list json marshal: %v", err)
		}
		return fmt.Sprintf("Auction List: %s", jsonStr)

	default:
		return "Error: parameter mismatch"
	}
}

func Winner(commandList []string) string {
	switch len(commandList) {
	case 1:
		return "Error: missing "

	case 2:
		if commandList[1] == "-help" {
			return "usage: get_auction_winner auctionId"
		}

		obj, err := GetAuctionWinner(commandList[1])
		if err != nil {
			return err.Error()
		}

		jsonStr, err := json.Marshal(obj)
		if err != nil {
			return fmt.Sprintf("Error: get_auction_winner json marshal: %v", err)
		}
		return fmt.Sprintf("Auction Winner: %s", jsonStr)

	default:
		return "Error: parameter mismatch"
	}
}

func Start(commandList []string) string {
	switch len(commandList) {
	case 1:
		return "Error: missing "

	case 2:
		if commandList[1] == "-help" {
			return "usage: start_auction auctionId"
		}

		// Send message to community
		err := group.SendMessage([]byte(commandList[1]))
		if err != nil {
			return fmt.Sprintf("Error: start_auction msg was not sent: %v", err)
		}

		// Update object store
		err = StartAuction(commandList[1])
		if err != nil {
			return err.Error()
		}

		return fmt.Sprintf("Auction %s was started", commandList[1])

	default:
		return "Error: parameter mismatch"
	}
}

func Stop(commandList []string) string {
	switch len(commandList) {
	case 1:
		return "Error: missing "

	case 2:
		if commandList[1] == "-help" {
			return "usage: stop_auction auctionId"
		}

		// Send message to community
		err := group.SendMessage([]byte(commandList[1]))
		if err != nil {
			return fmt.Sprintf("Error: stop_auction msg was not sent: %v", err)
		}

		err = StopAuction(commandList[1])
		if err != nil {
			return err.Error()
		}

		return fmt.Sprintf("Auction %s was stopped", commandList[1])

	default:
		return "Error: parameter mismatch"
	}
}

func GetBid(commandList []string) string {
	switch len(commandList) {
	case 1:
		return "Error: missing "

	case 2:
		if commandList[1] == "-help" {
			return "usage: get_bid bidId"
		}

		obj, err := GetBidObj(commandList[1])
		if err != nil {
			return err.Error()
		}

		jsonStr, err := json.Marshal(obj)
		if err != nil {
			return fmt.Sprintf("Error: get_bid json marshal: %v", err)
		}
		return fmt.Sprintf("Bid: %s", jsonStr)

	default:
		return "Error: parameter mismatch"
	}
}

func BidList(commandList []string) string {
	switch len(commandList) {
	case 1:
		return "Error: missing "

	case 2:
		if commandList[1] == "-help" {
			return "usage: get_bid_list auctionId"
		}

		objList, err := GetBidList(commandList[1])
		if err != nil {
			return err.Error()
		}

		jsonStr, err := json.Marshal(objList)
		if err != nil {
			return fmt.Sprintf("Error: get_bid_list json marshal: %v", err)
		}
		return fmt.Sprintf("Bid List: %s", jsonStr)

	default:
		return "Error: parameter mismatch"
	}
}

// Auction messages from community
/*
func Auction_Start_Event_Send(auctionObj AuctionObject) error {
	log.Println("Auction_Start_Event_Send")

	// Create a group message of the "auction_start" type
	groupMsg := group.GroupMessage{
		Type:    "auction_start",
		Version: "1.0",
		Data:    auctionObj,
	}

	// Serialize the message
	dataBytes, err := json.Marshal(groupMsg)
	if err != nil {
		errMsg := fmt.Sprintf("Auction_Start_Event_Send: marshalling group message failed: %v", err)
		log.Println(errMsg)
		// sendErrorMessage(conversation.ID, "Thread_Services: error", errMsg)
		return errors.New(errMsg)
	}

	// Get the conversation
	conversation, err := utilities.Cwtchbot.Peer.FetchConversationInfo(utilities.AuctionCommunityOnion)
	if err != nil {
		errMsg := fmt.Sprintf("Auction_Start_Event_Send: failed to find conversation for: %s, err: %v", utilities.AGENT_ADMIN_ID, err)
		log.Println(errMsg)
		//sendErrorMessage(conversation.ID, "Thread_Services: error", errMsg)
		return errors.New(errMsg)
	}

	// Send response to the group/community
	_, err = utilities.Cwtchbot.Peer.SendMessage(conversation.ID, string(utilities.Cwtchbot.PackMessage(model.OverlayChat, string(dataBytes))))
	if err != nil {
		log.Printf("Auction_Start_Event_Send: send message error: %v", err)
	}

	//log.Printf("Auction_Start_Event_Send: auction %s was started @ %v", auctionObj.AuctionId, auctionObj.StartDate)
	return nil
}

func Auction_Stop_Event_Send(id string, endTime time.Time) error {
	log.Println("Auction_Stop_Event_Send")

	auctionEnd := AuctionEnd{
		AuctionID: id,
		EndDate:   endTime,
	}
	// Create a group message of the "auction_stop" type
	groupMsg := group.GroupMessage{
		Type:    "auction_stop",
		Version: "1.0",
		Data:    auctionEnd,
	}

	// Serialize the message
	dataBytes, err := json.Marshal(groupMsg)
	if err != nil {
		errMsg := fmt.Sprintf("Auction_Stop_Event_Send: marshalling group message failed: %v", err)
		log.Println(errMsg)
		// sendErrorMessage(conversation.ID, "Thread_Services: error", errMsg)
		return errors.New(errMsg)
	}

	// Get the conversation
	conversation, err := utilities.Cwtchbot.Peer.FetchConversationInfo(utilities.AuctionCommunityOnion)
	if err != nil {
		errMsg := fmt.Sprintf("Auction_Stop_Event_Send: failed to find conversation for: %s, err: %v", utilities.AGENT_ADMIN_ID, err)
		log.Println(errMsg)
		//sendErrorMessage(conversation.ID, "Thread_Services: error", errMsg)
		return errors.New(errMsg)
	}

	// Send response to the group/community
	_, err = utilities.Cwtchbot.Peer.SendMessage(conversation.ID, string(utilities.Cwtchbot.PackMessage(model.OverlayChat, string(dataBytes))))
	if err != nil {
		log.Printf("Auction_Stop_Event_Send: send message error: %v", err)
	}

	log.Printf("Auction_Stop_Event_Send: auction %s was ended @ %v", id, endTime)
	return nil
}

func Add_Order_Event_Send(orderObj OrderObject) error {
	log.Println("Add_Order_Event_Send")

	// Create a group message of the "auction_start" type
	groupMsg := group.GroupMessage{
		Type:    "add_order",
		Version: "1.0",
		Data:    orderObj,
	}

	// Serialize the message
	dataBytes, err := json.Marshal(groupMsg)
	if err != nil {
		errMsg := fmt.Sprintf("Add_Order_Event_Send: marshalling group message failed: %v", err)
		log.Println(errMsg)
		// sendErrorMessage(conversation.ID, "Thread_Services: error", errMsg)
		return errors.New(errMsg)
	}

	// Get the conversation
	conversation, err := utilities.Cwtchbot.Peer.FetchConversationInfo(utilities.AuctionCommunityOnion)
	if err != nil {
		errMsg := fmt.Sprintf("Add_Order_Event_Send: failed to find conversation for: %s, err: %v", utilities.AGENT_ADMIN_ID, err)
		log.Println(errMsg)
		//sendErrorMessage(conversation.ID, "Thread_Services: error", errMsg)
		return errors.New(errMsg)
	}

	// Send response to the group/community
	_, err = utilities.Cwtchbot.Peer.SendMessage(conversation.ID, string(utilities.Cwtchbot.PackMessage(model.OverlayChat, string(dataBytes))))
	if err != nil {
		log.Printf("Add_Order_Event_Send: send message error: %v", err)
	}

	return nil
}

func Auction_Winner_Event_Send(auctionID string, bidID string) error {
	log.Println("Auction_Winner_Event_Send")

	auctionWinner := AuctionWinner{
		AuctionID: auctionID,
		BidID:     bidID,
	}

	// Create a group message of the "auction_winner" type
	groupMsg := group.GroupMessage{
		Type:    "auction_winner",
		Version: "1.0",
		Data:    auctionWinner,
	}

	// Serialize the message
	dataBytes, err := json.Marshal(groupMsg)
	if err != nil {
		errMsg := fmt.Sprintf("Auction_Winner_Event_Send: marshalling group message failed: %v", err)
		log.Println(errMsg)
		// sendErrorMessage(conversation.ID, "Thread_Services: error", errMsg)
		return errors.New(errMsg)
	}

	// Get the conversation
	conversation, err := utilities.Cwtchbot.Peer.FetchConversationInfo(utilities.AuctionCommunityOnion)
	if err != nil {
		errMsg := fmt.Sprintf("Auction_Winner_Event_Send: failed to find conversation for: %s, err: %v", utilities.AGENT_ADMIN_ID, err)
		log.Println(errMsg)
		//sendErrorMessage(conversation.ID, "Thread_Services: error", errMsg)
		return errors.New(errMsg)
	}

	// Send response to the group/community
	_, err = utilities.Cwtchbot.Peer.SendMessage(conversation.ID, string(utilities.Cwtchbot.PackMessage(model.OverlayChat, string(dataBytes))))
	if err != nil {
		log.Printf("Auction_Winner_Event_Send: send message error: %v", err)
	}

	log.Printf("Auction_Winner_Event_Send: auction %s was won by %s", auctionID, bidID)
	return nil
}

func Bid_Response_Event_Received(conversationID int, onion string, data []byte) {
	log.Println("Bid_Response_Event_Received")

	// Add bid to bidCollection
	err := AddBid(data)
	if err != nil {
		msg := fmt.Sprintf("Bid_Event_Received: AddBid: %v\n", err)
		log.Println(msg)
		//sendErrorMessage(conversationID, "Thread_Services: add_received error", msg)
		return
	}

	// Unmarshall the data -> bidObject
	var bidObj BidObject
	if err := json.Unmarshal(data, &bidObj); err != nil {
		msg := fmt.Sprintf("Bid_Event_Received: failed to unmarshal addObj: %v\n", err)
		log.Println(msg)
		//sendErrorMessage(conversationID, "Thread_Services: add_received error", msg)
		return
	}

	// SetResponseDate and BidId in Auction
	err = SetResponseDate(bidObj.AuctionId, bidObj.BidId, bidObj.Timestamp)
		if err != nil {
			msg := fmt.Sprintf("Bid_Event_Received: SetResponseDate: %v\n", err)
			log.Println(msg)
			//sendErrorMessage(conversationID, "Thread_Services: add_received error", msg)
			return
		}

	log.Printf("Bid_Response_Event_Received: bid %s added", bidObj.AuctionId)
}

*/
