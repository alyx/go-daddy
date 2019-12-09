# Unofficial GoDaddy API Client for Go

`go-daddy` provides a GoDaddy API client for the Go language. This package
provides a `net/http`-based wrapper around the [GoDaddy Developer API](https://developer.godaddy.com/).

## Requirements

This package is currently verified against the most recent version of the Go
compiler. At time of publishing, December 2019, the most recent version is
**1.13**. 

## Installation

`go-daddy` can be installed via `go get`:

```bash
$ go get github.com/alyx/go-daddy
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

If the environmental variables `GODADDY_API_KEY` or `GODADDY_API_SECRET` are
defined, these values will override the `APIKey` and `SecretKey` arguments
in `NewClient()`, respectively. Additionally, if the `GODADDY_API_URL`
environmental variable is defined, this value will override both the production
and OTE API URLs.

## Usage

To properly use `go-daddy`, you will need to import both the base `go-daddy` package,
as well as the packages for any endpoints you wish to communicate with in your program.

The available packages for import are:

* `github.com/alyx/go-daddy`
* `github.com/alyx/go-daddy/abuse`
* `github.com/alyx/go-daddy/aftermarket`
* `github.com/alyx/go-daddy/agreements`
* `github.com/alyx/go-daddy/certificates`
* `github.com/alyx/go-daddy/domains`
* `github.com/alyx/go-daddy/orders`
* `github.com/alyx/go-daddy/shoppers`
* `github.com/alyx/go-daddy/subscriptions`

### Example

This is a basic example showing how to generate a self-serve client and
retrieve a list of all domains owned by the account.

```go
package main

import (
	"log"

	godaddy "github.com/alyx/go-daddy"
	"github.com/alyx/go-daddy/domains"
)

func main() {
	client, _ := godaddy.NewClient("", "", false)

	myDomains, err := domains.List(client, nil, nil, 0, "", nil, "")
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