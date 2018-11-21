package TRTLServices

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
TSwrapper structure contains the
url, token and timeout info for
the TRTL Services
*/
type TSwrapper struct {
	url     string
	token   string
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

// CreateAddress method creates a new TRTL address
func (service *TSwrapper) CreateAddress() (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	data := url.Values{}

	response := service.makePostRequest("address", data)
	return response, nil
}

// DeleteAddress method deletes the specified address
func (service *TSwrapper) DeleteAddress(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeDeleteRequest("address/" + address)
	return response, nil
}

// GetAddress method gets the address details of the specified address
func (service *TSwrapper) GetAddress(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/" + address)

	return response, nil
}

// GetAddresses method views all addresses
// associated with the API token
func (service *TSwrapper) GetAddresses() (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/all")
	return response, nil
}

// ScanAddress method scans an address for transactions between
// a 100 block range starting from the specified blockIndex.
func (service *TSwrapper) ScanAddress(address string, blockIndex int) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/scan/" + address + "/" + strconv.Itoa(blockIndex))
	return response, nil
}

// GetAddressKeys method gets the public and
// secret spend keys of the specified address
func (service *TSwrapper) GetAddressKeys(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/keys/" + address)
	return response, nil
}

// IntegrateAddress method creates an integrated
// address with specified paymentID
func (service *TSwrapper) IntegrateAddress(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	data := url.Values{}
	data.Set("address", address)

	response := service.makePostRequest("transfer", data)
	return response, nil
}

// GetIntegratedAddresses mthod returns all integrated
// address associated with the given normal address
func (service *TSwrapper) GetIntegratedAddresses(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/integrate/" + address)
	return response, nil
}

// GetFee method calculates the TRTL Services fee for
// an amount specified in TRTL with two decimal points.
func (service *TSwrapper) GetFee(amount float64) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("transfer/fee/" + strconv.FormatFloat(amount, 'f', 2, 64))
	return response, nil
}

// CreateTransfer method sends a TRTL transaction with an
// address with the amount specified two decimal points.
func (service *TSwrapper) CreateTransfer(
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

// GetTransfer method gets transaction details
// specified by transaction hash.
func (service *TSwrapper) GetTransfer(transactionHash string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("transfer/" + transactionHash)
	return response, nil
}

// GetWallet method gets wallet container info and health check
func (service *TSwrapper) GetWallet() (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("wallet")
	return response, nil
}

// GetStatus method gets the current status of the TRTL Services infrastructure
func (service *TSwrapper) GetStatus() (*bytes.Buffer, error) {
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
func (service *TSwrapper) makePostRequest(method string, data url.Values) *bytes.Buffer {
	if method == "" {
		fmt.Println("No method supplied.")
		return nil
	}

	url := service.url + "/" + method

	req, err := http.NewRequest("post", url, strings.NewReader(data.Encode()))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("Authorization", "Bearer "+service.token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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

	req.Header.Add("Authorization", service.token)
	return decodeResponse(req)
}

// Decode Response
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
