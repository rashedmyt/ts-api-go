package tsapi

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

/*
TRTLServices structure contains the
URL and Token info of the TRTL Services
*/
type TRTLServices struct {
	URL   string
	Token string
}

func (service *TRTLServices) check() error {
	if service.URL == "" {
		service.URL = "https://api.trtl.services"
	}

	if service.Token == "" {
		return errors.New("All methods require an API key. See https://trtl.services/documentation")
	}

	return nil
}

// CreateAddress creates a new TRTL address
func (service *TRTLServices) CreateAddress() (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	data := url.Values{}

	response := service.makePostRequest("address", data)

	return response, nil
}

// DeleteAddress deletes a selected TRTL address
func (service *TRTLServices) DeleteAddress(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeDeleteRequest("address/" + address)

	return response, nil
}

// ViewAddress gets address details by address
func (service *TRTLServices) ViewAddress(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/view/" + address)

	return response, nil
}

// ViewAddresses view all addresses belonging to the specified token
func (service *TRTLServices) ViewAddresses() (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/view/all")

	return response, nil
}

// ScanAddress scan for transactions in the next 100
// blocks specified by blockIndex and address
func (service *TRTLServices) ScanAddress(address string, blockIndex int) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/scan/" + address + "/" + strconv.Itoa(blockIndex))

	return response, nil
}

// GetFee calculates the TRTL Services fee for a specified TRTL amount
func (service *TRTLServices) GetFee(amount float64) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("transfer/fee/" + strconv.FormatFloat(amount, 'f', 2, 64))

	return response, nil
}

// CreateTransfer sends a TRTL transaction with a specified amount
func (service *TRTLServices) CreateTransfer(
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

	data := url.Values{}
	data.Set("from", fromAddress)
	data.Set("to", toAddress)
	data.Set("amount", strconv.FormatFloat(amount, 'f', 2, 64))
	data.Set("fee", strconv.FormatFloat(fee, 'f', 2, 64))
	data.Set("paymentId", paymentID)
	data.Set("extra", extra)

	response := service.makePostRequest("transfer", data)

	return response, nil
}

// ViewTransfer lists transaction details with specified hash
func (service *TRTLServices) ViewTransfer(transactionHash string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("transfer/view/" + transactionHash)

	return response, nil
}

// GetStatus gets the current status of the TRTL Services infrastructure
func (service *TRTLServices) GetStatus() (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("status")

	return response, nil
}

func (service *TRTLServices) makeGetRequest(method string) *bytes.Buffer {
	url := service.URL + "/" + method

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("authorization", "Bearer "+service.Token)

	return decodeResponse(req)
}

func (service *TRTLServices) makePostRequest(method string, data url.Values) *bytes.Buffer {
	if method == "" {
		fmt.Println("No method supplied")
		return nil
	}

	url := service.URL + "/" + method

	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("authorization", "Bearer "+service.Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return decodeResponse(req)
}

func (service *TRTLServices) makeDeleteRequest(method string) *bytes.Buffer {
	if method == "" {
		fmt.Println("No method supplied")
		return nil
	}

	url := service.URL + "/" + method

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("authorization", "Bearer "+service.Token)

	return decodeResponse(req)
}

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
