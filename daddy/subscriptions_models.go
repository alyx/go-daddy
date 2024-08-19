// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

// ProductGroup represents a group of a select type of products
type ProductGroup struct {
	ProductGroupKey   string
	SubscriptionCount int
}

// Subscription represents a product subscription
type Subscription struct {
	Addons           []SubscriptionAddon
	Billing          SubscriptionBilling
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

// SubscriptionAddon represents additional components of a primary subscription
type SubscriptionAddon struct {
	Commitment string
	PfID       int
	Quantity   int
}

// SubscriptionBilling represents the billing aspects of a subscription
type SubscriptionBilling struct {
	Commitment   string
	PastDueTypes []string
	RenewAt      string
	Status       string
}

// SubscriptionList is a paginated list of subscriptions
type SubscriptionList struct {
	Pagination    *Pagination
	Subscriptions []Subscription
}

// SubscriptionProduct is a specific product within a subscription
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

// SubscriptionRelations represents the relationships between multiple
// subscriptions
type SubscriptionRelations struct {
	Children []string
	Parent   string
}

// SubscriptionUpdate represents confirmation of an updated subscription
type SubscriptionUpdate struct {
	PaymentProfileID int
	RenewAuto        bool
}
