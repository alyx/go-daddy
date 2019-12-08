package certificates

// Create will create a pending order for a certificate
func Create() {}

// Validate will validate a pending order for a certificate
func Validate() {}

// Get retrieves certificate details for a specified certificate
func Get() {}

// GetActions retrieves all certificate actions
func GetActions() {}

// ResendEmail resends a validation email
func ResendEmail() {}

// AddEmail adds an alternate email address
func AddEmail() {}

// ResendEmailByID resents a specific email to an email address
func ResendEmailByID() {}

// History retrieves email history
func History() {}

// DeleteCallback unregisters a system callback
func DeleteCallback() {}

// GetCallback retrieves a system stateful action callback URL
func GetCallback() {}

// AddCallback registers a certificate action callback
func AddCallback() {}

// Cancel a pending certificate
func Cancel() {}

// Download certificate
func Download() {}

// Reissue the active certificate
func Reissue() {}

// Renew the active certificate
func Renew() {}

// Revoke the active certificate
func Revoke() {}

// SiteSeal generates and retrieves the site seal
func SiteSeal() {}

// VerifyDomainControl confirms domain control
func VerifyDomainControl() {}
