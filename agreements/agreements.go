package agreements

import "github.com/alyx/godaddy"

// Get retrieves Legal Agreements for provided agreements keys
func Get(c *godaddy.Client, privateLabelID int, marketID int, keys []string) {
	//XXX: This function requires significant investment in refactoring c.Get()
	//     for additional headers. TBD.
	return
}
