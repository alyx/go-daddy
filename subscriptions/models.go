package subscriptions

import (
	godaddy "github.com/alyx/go-daddy"
)

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
	Pagination    *godaddy.Pagination
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
