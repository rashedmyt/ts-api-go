package TRTLServices

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

/*
TSwrapper structure contains the
url and Token info of the TRTL Services
*/
type TSwrapper struct {
	url string
	token string
	timeout int
}

func (service *TSwrapper) check() error {
	service.url = "https://api.trtl.services/v1"
	
	if service.token == "" {
		return errors.New("All methods require an JWT access token. See https://trtl.services/docs")
	}

	if service.timeout == 0 {
		service.timeout = 2000
	}

	return nil
}

// Create Address
func (service *TSwrapper) createADdress(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	response := service.makePostRequest("address", data)
	return response, nil
}


// Delete Address
func (service *TSwrapper) deleteAddress(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeDeleteRequest("address/" + address)
	return response, nil
}


// Get Adddress
func (service *TSwrapper) getAddress(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/" + address)

	return response, nil
}

// Get Addresses
func (service *TSwrapper) getAddresses() (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/all")
	return response, nil
}


// Scan Address
func (service *TSwrapper) scanAddress(address string, blockIndex int) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/scan/" + address + "/" + strconv.Itoa(blockIndex))
	return response, nil
}


// get Address Keys
func (service *TSwrapper) getAddressKeys(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/keys/" + address)
	return response, nil
}


// Integrate Address
func (service *TSwrapper) integrateAddress(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	data["address"] = address
	response := service.makePostRequest("address/integrate", data)
	return response, nil
}


// Get Integrated Addresses
func (service *TSwrapper) getIntegratedAddresses(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/integrate/" + address)
	return response, nil
}


// GetFee
func (service *TSwrapper) getFee(amount float64) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("transfer/fee/" + strconv.FormatFloat(amount, 'f', 2, 64))
	return response, nil
}


// Create Transfer
func (service *TSwrapper) createTransfer(
	fromAddress string,
	toAddress string,
	amount float64,
	fee float64,
	paymentID string,
	extra string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	data["from"] = fromAddress
	data["to"] = toAddress
	data["amount"] = amount
	data["fee"] = fee
	data["paymentId"] = paymentID
	data["extra"] = extra

	response := service.makePostRequest("transfer", data)
	return response, nil
}


// Get Transfer
func (service *TSwrapper) getTransfer(transactionHash string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("transfer/" + transactionHash)
	return response, nil
}


// Get Wallet
func (service *TSwrapper) getWallet() (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("wallet")
	return response, nil
}


// Get Status
func (service *TSwrapper) getStatus() (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("status")
	return response, nil
}


// Get Method
func (service *TSwrapper) makeGetRequest(method string) *bytes.Buffer {
	url := service.url + "/" + method

	req, err := http.NewRequest("get", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("Authorization", service.token)
	return decodeResponse(req)
}


// Post Method
func (service *TSwrapper) makePostRequest(method string, params map[string]interface{}) *bytes.Buffer {
	if method == "" {
		fmt.Println("No method supplied.")
		return nil
	}

	url := service.url + "/" + method

	jsonPayload, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	body := bytes.NewBuffer(jsonPayload)

	req, err := http.NewRequest("post", url, body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("authorization", service.token)
	return decodeResponse(req)
}

// Delete Method
func (service *TSwrapper) makeDeleteRequest(method string) *bytes.Buffer {
	if method == "" {
		fmt.Println("No method supplied.")
		return nil
	}

	url := service.url + "/" + method

	req, err := http.NewRequest("delete", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("authorization", service.token)
	return decodeResponse(req)
}

// Decode Res 
func decodeResponse(req *http.Request) *bytes.Buffer {
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	responseBody, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println(err1)
		return nil
	}

	return bytes.NewBuffer(responseBody)
}
