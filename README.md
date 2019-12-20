# Unofficial GoDaddy API Client for Go

[![GoDoc](https://godoc.org/github.com/alyx/go-daddy/github?status.svg)](https://godoc.org/github.com/alyx/go-daddy/github)
[![Go Report Card](https://goreportcard.com/badge/github.com/alyx/go-daddy)](https://goreportcard.com/report/github.com/alyx/go-daddy)

`go-daddy` provides a GoDaddy API client for the Go language. This package
provides a `net/http`-based wrapper around the [GoDaddy Developer API](https://developer.godaddy.com/).

## Requirements

This package is currently verified against the most recent version of the Go
compiler. At time of publishing, December 2019, the most recent version is
**1.13**. 

## Installation

`go-daddy` can be installed via `go get`:

```bash
$ go get github.com/alyx/go-daddy/daddy
```

## Configuration

Before use, please visit the [GoDaddy Getting Started](https://developer.godaddy.com/getstarted)
developer page, which walks you through getting access to the API, using the API as a self-serve user vs a reseller, and informs you of the terms of use.

You will need to generate a [GoDaddy API key](https://developer.godaddy.com/keys) to use `go-daddy`.

After importing the core `go-daddy` package, you will need to initialize a
client with `godaddy.NewClient(APIKey, SecretKey, OTE)` where `APIKey` is your
GoDaddy API key, `SecretKey` is the secret key generated with it, and `OTE`
is a value declaring if your API calls should point to the GoDaddy _OTE_
(Test Environment) or the Production environment -- setting `OTE` to `true`
specifies that you do wish to use the test environment, and `false` is the
Production environment.

## Usage

```
import "github.com/alyx/go-daddy/daddy"
```

Construct a new GoDaddy client, then use the various services on the client to
access different parts of the GoDaddy API. For example:

```
client := daddy.NewClient("API Key", "Secret Key", true)

// list all domains owned by the owner of the API key
domains, err := client.Domains.List(nil, nil, 0, "", nil, "")
```

The services of a client divide the API into chunks corresponding to the
structure of the GoDaddy API documentation at https://developer.godaddy.com/doc

The current available client services are:
`.Abuse, .Aftermarket, .Agreements, .Certificates, .Countries, .Domains, .Orders, .Shoppers, .Subscriptions`

### Example

This is a basic example showing how to generate a self-serve client and
retrieve a list of all domains owned by the account.

```go
package main

import (
	"log"

	"github.com/alyx/go-daddy"
)

func main() {
	client, _ := daddy.NewClient("abc", "def", false)

	myDomains, err := client.Domains.List(nil, nil, 0, "", nil, "")
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range myDomains {
		log.Println(value.Domain)
	}
}
```

## License

This package is licensed under the ISC license, see `LICENSE`