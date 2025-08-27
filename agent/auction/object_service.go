package auction

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"kotosAuctionAgent/agent/utilities"
	"log"
	"net/http"
	"time"
)

func CreateAuction(byteData []byte) error {
	// Define the URL to the auction manager
	url := fmt.Sprintf("%s/createAuction", utilities.AUCTION_MGR_URI)

	// Create HTTP request with byteData as body
	req, err := http.NewRequest("POST", url, bytes.NewReader(byteData))
	if err != nil {
		errMsg := fmt.Sprintf("Error: create_auction creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error: create_auction HTTP POST %s: %v", url, err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("Error: create_auction HTTP status code %d", resp.StatusCode)
		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	id, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg := fmt.Sprintf("Error: reading response body: %v", err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}
	log.Printf("AuctionId = %s", string(id))

	return nil
}

func GetAuctionObj(id string) (AuctionObject, error) {
	var obj AuctionObject

	// Define the URL to the auction manager
	url := fmt.Sprintf("%s/getAuction/%s", utilities.AUCTION_MGR_URI, id)

	// Create HTTP request with byteData as body
	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("Error: get_auction creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return obj, errors.New(errMsg)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error: get_auction HTTP Get %s: %v", url, err)
		log.Println(errMsg)
		return obj, errors.New(errMsg)
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("Error: get_auction HTTP status code %d", resp.StatusCode)
		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return obj, errors.New(errMsg)
	}

	if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		errMsg := fmt.Sprintf("Error: get_auction decoding assignments JSON: %v", err)
		log.Println(errMsg)
		return obj, err
	}

	return obj, nil
}

func GetAuctionList() ([]AuctionObject, error) {
	var _list []AuctionObject

	// Define the URL to the meta data manager
	url := fmt.Sprintf("%s/getAuctionList", utilities.AUCTION_MGR_URI)

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("Error: get_auction_list creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return _list, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error: get_auction_list HTTP POST %s: %v", url, err)
		log.Println(errMsg)
		return _list, err
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("Error: get_auction_list HTTP status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return _list, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&_list); err != nil {
		errMsg := fmt.Sprintf("Error: get_auction_list decoding assignments JSON: %v", err)
		log.Println(errMsg)
		return _list, err
	}

	return _list, nil
}

func GetAuctionWinner(id string) (AuctionObject, error) {
	var obj AuctionObject

	// Define the URL to the auction manager
	url := fmt.Sprintf("%s/getAuctionWinner/%s", utilities.AUCTION_MGR_URI, id)

	// Create HTTP request with byteData as body
	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("Error: get_auction_winner creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return obj, errors.New(errMsg)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error: get_auction_winner HTTP Get %s: %v", url, err)
		log.Println(errMsg)
		return obj, errors.New(errMsg)
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("Error: get_auction_winner HTTP status code %d", resp.StatusCode)
		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return obj, errors.New(errMsg)
	}

	if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		errMsg := fmt.Sprintf("Error: get_auction_winner decoding assignments JSON: %v", err)
		log.Println(errMsg)
		return obj, err
	}

	return obj, nil
}

func StartAuction(id string) error {
	// Define the URL to the meta data manager
	url := fmt.Sprintf("%s/startAuction/%s", utilities.AUCTION_MGR_URI, id)

	// Create HTTP request
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("Error: start_auction creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return err
	}

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error: start_auction HTTP PUT %s: %v", url, err)
		log.Println(errMsg)
		return err
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("Error: start_auction HTTP status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return err
	}

	return nil
}

func StopAuction(id string) error {
	// Define the URL to the meta data manager
	url := fmt.Sprintf("%s/stopAuction/%s", utilities.AUCTION_MGR_URI, id)

	// Create HTTP request
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("Error: stop_auction creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return err
	}

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error: stop_auction HTTP PUT %s: %v", url, err)
		log.Println(errMsg)
		return err
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("Error: stop_auction HTTP status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return err
	}

	return nil
}

func GetBidObj(id string) (AuctionObject, error) {
	var obj AuctionObject

	// Define the URL to the auction manager
	url := fmt.Sprintf("%s/getBid/%s", utilities.AUCTION_MGR_URI, id)

	// Create HTTP request with byteData as body
	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("Error: get_bid creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return obj, errors.New(errMsg)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error: get_bid HTTP Get %s: %v", url, err)
		log.Println(errMsg)
		return obj, errors.New(errMsg)
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("Error: get_bid HTTP status code %d", resp.StatusCode)
		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return obj, errors.New(errMsg)
	}

	if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		errMsg := fmt.Sprintf("Error: get_bid decoding assignments JSON: %v", err)
		log.Println(errMsg)
		return obj, err
	}

	return obj, nil
}

func GetBidList(id string) ([]BidObject, error) {
	var _list []BidObject

	// Define the URL to the meta data manager
	url := fmt.Sprintf("%s/getBidList/%s", utilities.AUCTION_MGR_URI, id)

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("GetBidList: creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return _list, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("GetBidList: HTTP POST %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return _list, err
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("GetBidList: status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return _list, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&_list); err != nil {
		errMsg := fmt.Sprintf("GetStartAuctionList: decoding assignments JSON: %v", err)
		log.Println(errMsg)
		return _list, err
	}

	return _list, nil
}

/*
func GetStopAuctionList(duration int) ([]AuctionObject, error) {
	var _list []AuctionObject

	// Define the URL to the meta data manager
	url := fmt.Sprintf("%s/getStopList/%d", utilities.AUCTION_DATA_MGR_URI, duration)

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("GetStopAuctionList: creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return _list, err
	}

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("GetStopAuctionList: HTTP POST %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return _list, err
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("GetStopAuctionList: status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http status", errMsg)
		return _list, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&_list); err != nil {
		errMsg := fmt.Sprintf("GetStopAuctionList: decoding assignments JSON: %v", err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http status", errMsg)
		return _list, err
	}

	return _list, nil
}

func GetAwaitingWinnerList() ([]AuctionObject, error) {
	var _list []AuctionObject

	// Define the URL to the meta data manager
	url := fmt.Sprintf("%s/getAwaitingWinnerList", utilities.AUCTION_DATA_MGR_URI)

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("GetAwaitingWinnerList: creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return _list, err
	}

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("GetAwaitingWinnerList: HTTP POST %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return _list, err
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("GetAwaitingWinnerList: status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http status", errMsg)
		return _list, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&_list); err != nil {
		errMsg := fmt.Sprintf("GetAwaitingWinnerList: decoding assignments JSON: %v", err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http status", errMsg)
		return _list, err
	}

	return _list, nil
}

func SetAuctionStarted(id string, startTime time.Time) error {
	// Define the URL to the meta data manager
	url := fmt.Sprintf("%s/startAuction/%s/%s", utilities.AUCTION_DATA_MGR_URI, id, startTime.Format(time.RFC3339))
	//log.Printf("URL = %s", url)

	// Create HTTP request
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("SetAuctionStarted: creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return err
	}

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("SetAuctionStarted: HTTP PUT %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return err
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("SetAuctionStarted: status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http status", errMsg)
		return err
	}

	return nil
}

func SetAuctionStopped(id string, stopTime time.Time) error {
	// Define the URL to the meta data manager
	url := fmt.Sprintf("%s/stopAuction/%s/%s", utilities.AUCTION_DATA_MGR_URI, id, stopTime.Format(time.RFC3339))
	//log.Printf("URL = %s", url)

	// Create HTTP request
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("SetAuctionStopped: creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return err
	}

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("SetAuctionStopped: HTTP PUT %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return err
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("SetAuctionStopped: status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http status", errMsg)
		return err
	}

	return nil
}

func SetAuctionWinner(id string, bidId string) error {
	// Define the URL to the meta data manager
	url := fmt.Sprintf("%s/setAuctionWinner/%s/%s", utilities.AUCTION_DATA_MGR_URI, id, bidId)

	// Create HTTP request
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("SetAuctionWinner: creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return err
	}

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("SetAuctionWinner: HTTP PUT %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return err
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("SetAuctionStopped: status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http status", errMsg)
		return err
	}

	return nil
}

func AddBid(byteData []byte) error {
	// Define the URL to the bid data manager
	url := fmt.Sprintf("%s/addBid", utilities.AUCTION_DATA_MGR_URI)

	// Create HTTP request with byteData as body
	req, err := http.NewRequest("POST", url, bytes.NewReader(byteData))
	if err != nil {
		errMsg := fmt.Sprintf("Bid_Response_Event_Received: creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return errors.New(errMsg)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Bid_Response_Event_Received: HTTP POST %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return errors.New(errMsg)
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("Bid_Response_Event_Received: status code %d", resp.StatusCode)
		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http status", errMsg)
		return errors.New(errMsg)
	}

	return nil
}

func GetBidList(id string) ([]BidObject, error) {
	var _list []BidObject

	// Define the URL to the meta data manager
	url := fmt.Sprintf("%s/getBidList/%s", utilities.AUCTION_DATA_MGR_URI, id)

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("GetBidList: creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return _list, err
	}
	//req.Header.Set("Content-Type", "application/json")

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("GetBidList: HTTP POST %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return _list, err
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("GetBidList: status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http status", errMsg)
		return _list, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&_list); err != nil {
		errMsg := fmt.Sprintf("GetStartAuctionList: decoding assignments JSON: %v", err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http status", errMsg)
		return _list, err
	}

	return _list, nil
}
*/
