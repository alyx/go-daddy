// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

// LegalAgreement represents a singular agreement document
// (such as Terms of Service, etc.)
type LegalAgreement struct {
	AgreementKey string
	Content      string
	Title        string
	URL          string
}
