package aftermarket

// ListingAction contains the resulting action ID from performing a change
// to the Aftermarket listings
type ListingAction struct {
	ListingActionID int
}

// ListingExpiryCreate handles the input structure for creating a new listing
// in the Aftermarket registry.
type ListingExpiryCreate struct {
	Domain            string
	ExpiresAt         string
	LosingRegistrarID int
	PageViewsMonthly  int
	RevenueMonthly    int
}
