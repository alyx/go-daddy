// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

// Country represents a country, containing information about the country
// and a list of states
type Country struct {
	CountrySummary
	States []State
}

// CountrySummary represents a brief, less in-depth summary of a country
type CountrySummary struct {
	CallingCode string
	CountryKey  string
	Label       string
}

// State represents a state, a political/legal segment of a country
type State struct {
	Label    string
	StateKey string
}
