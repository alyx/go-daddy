// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

type ProductGroup struct {
	ProductGroupKey   string
	SubscriptionCount int
}

type Subscription struct {
	Addons           []SubscriptionAddon
	Billing          []SubscriptionBilling
	Cancelable       bool
	CreatedAt        string
	ExpiresAt        string
	Label            string
	LaunchURL        string
	PaymentProfileID int
	PriceLocked      bool
	Product          *SubscriptionProduct
	Relations        *SubscriptionRelations
	RenewAuto        bool
	Renewable        bool
	Status           string
	SubscriptionID   string
	Upgradable       bool
}

type SubscriptionAddon struct {
	Commitment string
	PfID       int
	Quantity   int
}

type SubscriptionBilling struct {
	Commitment   string
	PastDueTypes []string
	RenewAt      string
	Status       string
}

type SubscriptionList struct {
	Pagination    *Pagination
	Subscriptions []Subscription
}

type SubscriptionProduct struct {
	Label             string
	Namespace         string
	PfID              int
	ProductGroupKey   string
	RenewalPeriod     int
	RenewalPeriodUnit string
	RenewalPfID       int
	SupportBillOn     bool
}

type SubscriptionRelations struct {
	Children []string
	Parent   string
}

type SubscriptionUpdate struct {
	PaymentProfileID int
	RenewAuto        bool
}
