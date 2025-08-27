package auction

func Messages(data string, conversationID int, onion string) {
	/*	var groupMsg group.GroupMessage

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

		// Determine auction message type
		switch strings.ToLower(groupMsg.Type) {

		case "bid_response":
			Bid_Response_Event_Received(conversationID, onion, dataBytes)

		default:
			log.Printf("Auction MessageType error: %v from %d", groupMsg.Type, conversationID)
			// sendErrorMessage(envelope.ConversationID, "MessageType error", m.Type)
		}
	*/
}
