[![Go Report Card](https://goreportcard.com/badge/github.com/trtl-services/trtl-api-go)](https://goreportcard.com/report/github.com/trtl-services/trtl-api-go)
[![GoDoc](https://godoc.org/github.com/trtl-services/trtl-api-go?status.svg)](https://godoc.org/github.com/trtl-services/trtl-api-go)

# TRTL Service Golang API Interface

This wrapper allows you to easily interact with the [TRTL Services](https://trtl.services) API to quickly develop applications that interact with the [TurtleCoin](https://turtlecoin.lol) Network.


# Table of Contents

1. [Installation](#installation)
2. [Intialization](#intialization)
3. [Documentation](#documentation)
   1. [Methods](#methods)

# Installation

```bash
go get github.com/trtl-services/ts-api-go
```

# Intialization

```go
import (
    "fmt"

    "github.com/trtl-services/ts-api-go"
)

services := tsapi.TRTLServices{
    URL: "https://api.trtl.services",
    Token: "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzdCIsImFwcElkIjo0LCJ1c2VySWQiOjYsInBlcm1pc3Npb25zIjpbImFkZHJlc3M6bmV3Il0sImlhdCI6MTUzNjU4NTM2NywiZXhwIjoxNTM5MTc3MzY3LCJhdWQiOiJ0dXJ0bGV3YWxsZXQuaW8iLCJpc3MiOiJUUlRMIFNlcnZpY2VzIiwianRpIjoiMzMifQ.AEHXmvTo8RfNuZ15Y3IGPRhZPaJxFSmOZvVv2YGN9L4We7bXslIPxhMv_n_5cNW8sIgE2Fr-46OTb5H5AFgpjA",
}
```

# Documentation

API documentation is available at https://trtl.services/documentation


## Methods

### createAddress()
Create a new TRTL addresses

```go
response, err := services.CreateAddress()
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### deleteAddress(address)
Delete a selected TRTL addresses

```go
response, err := services.DeleteAdddress("TRTLuxH78akDMCsXycnU5HjJE6zPCgM4KRNNQSboqh1yiTnvxuhNVUL9tK92j9kurSKdXVHFmjSRkaNBxM6Nb3G8eQGL7aj113A")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### ViewAddress()
Get address details by address
```go
response, err := services.ViewAddress("TRTLuxH78akDMCsXycnU5HjJE6zPCgM4KRNNQSboqh1yiTnvxuhNVUL9tK92j9kurSKdXVHFmjSRkaNBxM6Nb3G8eQGL7aj113A")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### viewAddresses()
View all addresses belonging to the specified token.

```go
response, err := services.ViewAddresses()
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### scanAddress()
Scan for transactions in the next 100 blocks specified by blockIndex and address.

```go
response, err := services.ScanAddress("TRTLuxH78akDMCsXycnU5HjJE6zPCgM4KRNNQSboqh1yiTnvxuhNVUL9tK92j9kurSKdXVHFmjSRkaNBxM6Nb3G8eQGL7aj113A", 899093)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### getFee()
Calculate the TRTL Services fee for a specified TRTL amount.

```go
response, err := services.GetFee(1092.19)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### createTransfer()
Send a TRTL transaction with a specified account.

```go
response, err := services.CreateTransfer(
    "TRTLuxH78akDMCsXycnU5HjJE6zPCgM4KRNNQSboqh1yiTnvxuhNVUL9tK92j9kurSKdXVHFmjSRkaNBxM6Nb3G8eQGL7aj113A", "TRTLuzAzNs1E1RBFhteX56A5353vyHuSJ5AYYQfoN97PNbcMDvwQo4pUWHs7SYpuD9ThvA7AD3r742kwTmWh5o9WFaB9JXH8evP",
    1000.00,
    1.00,
    "",
    "",
)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```

### viewTransfer()
Lists transaction details with specified hash.

```go
response, err := services.ViewTransfer("EohMUzR1DELyeQM9RVVwpmn5Y1DP0lh1b1ZpLQrfXQsgtvGHnDdJSG31nX2yESYZ")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### getStatus()
Get the current status of the TRTL Services infrastructure.

```go
response, err := services.GetStatus()
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


# License

```
Copyright (C) 2018 Rashed Mohammed, The TurtleCoin Developers

Please see the included LICENSE file for more information.
```
