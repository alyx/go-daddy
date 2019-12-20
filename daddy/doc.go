// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

/*
Package daddy provides a client for using the GoDaddy API.

Usage:

		import "github.com/alyx/go-daddy/daddy"

Construct a new GoDaddy client, then use the various services on the client to
access different parts of the GoDaddy API. For example:

		client := daddy.NewClient("API Key", "Secret Key", true)

		// list all domains owned by the owner of the API key
		domains, err := client.Domains.List(nil, nil, 0, "", nil, "")

The services of a client divide the API into chunks corresponding to the
structure of the GoDaddy API documentation at https://developer.godaddy.com/doc

For more sample code snippets, visit the
https://github.com/alyx/go-daddy/tree/master/example directory.

Authentication

Upon initialization, the go-daddy library will require three parameters: an
API key, a secret key and a boolean on whether to connect to the GoDaddy OTE
(development) environment, or the production environment.

*/
package daddy
