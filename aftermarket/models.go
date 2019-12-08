package aftermarket

type ListingAction struct {
	ListingActionId int
}

type ListingExpiryCreate struct {
	Domain            string
	ExpiresAt         string
	LosingRegistrarId int
	PageViewsMonthly  int
	RevenueMonthly    int
}

type Expiry struct {
	Id int
}
