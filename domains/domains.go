package domains

import (
	"github.com/alyx/godaddy"
)

func List(c *godaddy.Client) {
	return
}

func GetAgreements() { return }

func GetAvailable() { return }

func ValidateContactSchema()       { return }
func GetPurchaseSchema(tld string) { return }

func ValidatePurchaseSchema() { return }

func GetSuggestions() { return }

func GetTLDs() { return }

// Individual domain settings

func Purchase()       { return }
func Delete()         { return }
func Get()            { return }
func Update()         { return }
func UpdateContacts() { return }

// Privacy
func DeletePrivacy()   { return }
func PurchasePrivacy() { return }

func AddRecords()                  { return }
func ReplaceRecords()              { return }
func ReplaceRecordsByType()        { return }
func ReplaceRecordsByTypeAndName() { return }
func GetRecords()                  { return }

func Renew()               { return }
func Transfer()            { return }
func SendRegistrantEmail() { return }

// v2

func RemoveForward() { return }
func GetForward()    { return }
func UpdateForward() { return }
