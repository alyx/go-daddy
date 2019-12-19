// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

type Country struct {
	CountrySummary
	States []State
}

type CountrySummary struct {
	CallingCode string
	CountryKey  string
	Label       string
}

type State struct {
	Label    string
	StateKey string
}
