[![Go Report Card](https://goreportcard.com/badge/github.com/rashedmyt/trtl-api-go)](https://goreportcard.com/report/github.com/rashedmyt/trtl-api-go)
[![GoDoc](https://godoc.org/github.com/rashedmyt/trtl-api-go?status.svg)](https://godoc.org/github.com/rashedmyt/trtl-api-go)

# TRTL Service Golang API Interface

This wrapper allows you to easily interact with the [TRTL TS](https://trtl.TS) API to quickly develop applications that interact with the [TurtleCoin](https://turtlecoin.lol) Network.


# Table of Contents

1. [Installation](#installation)
2. [Intialization](#intialization)
3. [Documentation](#documentation)
   1. [Methods](#methods)

# Installation

```bash
go get github.com/trtl-TS/ts-api-go
```

# Intialization

```go
import (
    "fmt"

    "github.com/trtl-TS/ts-api-go"
)

TS := TRTLServices.TSwrapper {
    token: "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzdCIsImFwcElkIjo0LCJ1c2VySWQiOjYsInBlcm1pc3Npb25zIjpbImFkZHJlc3M6bmV3Il0sImlhdCI6MTUzNjU4NTM2NywiZXhwIjoxNTM5MTc3MzY3LCJhdWQiOiJ0dXJ0bGV3YWxsZXQuaW8iLCJpc3MiOiJUUlRMIFNlcnZpY2VzIiwianRpIjoiMzMifQ.AEHXmvTo8RfNuZ15Y3IGPRhZPaJxFSmOZvVv2YGN9L4We7bXslIPxhMv_n_5cNW8sIgE2Fr-46OTb5H5AFgpjA",
    timeout: 2000
}
```


# Documentation

API documentation is available at https://trtl.services/docs


## Methods

### createAddress()
Create a new TRTL addresses

```go
response, err := TS.createAddress()
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```

### getAddress(address)
Get address details by address
```go
response, err := TS.getAddress("TRTLuxH78akDMCsXycnU5HjJE6zPCgM4KRNNQSboqh1yiTnvxuhNVUL9tK92j9kurSKdXVHFmjSRkaNBxM6Nb3G8eQGL7aj113A")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```

### deleteAddress(address)
Delete a selected TRTL address

```go
response, err := TS.deleteAddress("TRTLuxH78akDMCsXycnU5HjJE6zPCgM4KRNNQSboqh1yiTnvxuhNVUL9tK92j9kurSKdXVHFmjSRkaNBxM6Nb3G8eQGL7aj113A")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### getAddresses()
View all addresses.

```go
response, err := TS.getAddresses()
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### scanAddress(address, blockIndex)
Scan an address for transactions between a 100 block range starting from the specified blockIndex.

```go
response, err := TS.scanAddress("TRTLuxH78akDMCsXycnU5HjJE6zPCgM4KRNNQSboqh1yiTnvxuhNVUL9tK92j9kurSKdXVHFmjSRkaNBxM6Nb3G8eQGL7aj113A", 899093)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### getAddressKeys(address)
Get the public and secret spend key of an address.


```go
response, err := TS.getAddressKeys("TRTLuxH78akDMCsXycnU5HjJE6zPCgM4KRNNQSboqh1yiTnvxuhNVUL9tK92j9kurSKdXVHFmjSRkaNBxM6Nb3G8eQGL7aj113A")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### integrateAddress(address, paymentId)
Create an integrated address with an address and payment ID.

```go
response, err := TS.getAddressKeys("TRTLuxH78akDMCsXycnU5HjJE6zPCgM4KRNNQSboqh1yiTnvxuhNVUL9tK92j9kurSKdXVHFmjSRkaNBxM6Nb3G8eQGL7aj113A", "7d89a2d16365a1198c46db5bbe1af03d2b503a06404f39496d1d94a0a46f8804")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### getIntegratedAddresses(address)
Get all integrated addresses by address.

```go
response, err := TS.getIntegratedAddresses("TRTLuxH78akDMCsXycnU5HjJE6zPCgM4KRNNQSboqh1yiTnvxuhNVUL9tK92j9kurSKdXVHFmjSRkaNBxM6Nb3G8eQGL7aj113A", "7d89a2d16365a1198c46db5bbe1af03d2b503a06404f39496d1d94a0a46f8804")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### getFee(amount)
Calculate the TRTL Services fee for an amount specified in TRTL with two decimal points.

```go
response, err := TS.getFee(1092.19)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### createTransfer()
Send a TRTL transaction with an address with the amount specified two decimal points.

```go
response, err := TS.createTransfer(
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


### getTransfer()
Get a transaction details specified by transaction hash.

```go
response, err := TS.getTransfer("EohMUzR1DELyeQM9RVVwpmn5Y1DP0lh1b1ZpLQrfXQsgtvGHnDdJSG31nX2yESYZ")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### getWallet()
Get wallet container info and health check.

```go
response, err := TS.getWallet()
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### getStatus()
Get the current status of the TRTL Services infrastructure.

```go
response, err := TS.getStatus()
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
