package certificates

import (
	"encoding/json"

	"github.com/alyx/godaddy"
)

// Create will create a pending order for a certificate
func Create(c *godaddy.Client, body *CertificateCreate) (*Identifier, error) {
	var res = new(Identifier)

	enc, err := json.Marshal(body)
	if err != nil {
		return res, err
	}

	data, err := c.Post("/v1/certificates", enc)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// Validate will validate a pending order for a certificate
func Validate(c *godaddy.Client, body *CertificateCreate) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = c.Post("/v1/certificates/validate", enc)
	if err != nil {
		return err
	}

	return err
}

// Get retrieves certificate details for a specified certificate
func Get(c *godaddy.Client, id string) (*Certificate, error) {
	res := new(Certificate)

	data, err := c.Get("/v1/certificates/" + id)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// GetActions retrieves all certificate actions
func GetActions(c *godaddy.Client, id string) ([]Action, error) {
	var res []Action

	data, err := c.Get("/v1/certificates/" + id + "/actions")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// ResendEmail resends a validation email
func ResendEmail(c *godaddy.Client, cert string, email string) error {
	_, err := c.Post("/v1/certificates/"+cert+"/email/"+email+"/resend", nil)

	return err
}

// AddEmail adds an alternate email address
func AddEmail(c *godaddy.Client, cert string, email string) (*EmailHistory, error) {
	res := new(EmailHistory)

	data, err := c.Post("/v1/certificates/"+cert+"/email/resend/"+email, nil)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// ResendEmailByID resents a specific email to an email address
func ResendEmailByID(c *godaddy.Client, cert string, emailID string, email string) error {
	_, err := c.Post("/v1/certificates/"+cert+"/email/"+emailID+"/resend/"+email, nil)

	return err
}

// History retrieves email history
func History(c *godaddy.Client, id string) (*EmailHistory, error) {
	res := new(EmailHistory)

	data, err := c.Get("/v1/certificates/" + id + "/email/history")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// DeleteCallback unregisters a system callback
func DeleteCallback(c *godaddy.Client, id string) error {
	_, err := c.Delete("/v1/certificates/"+id+"/callback", nil)

	return err
}

// GetCallback retrieves a system stateful action callback URL
func GetCallback(c *godaddy.Client, id string) (*Callback, error) {
	res := new(Callback)

	data, err := c.Get("/v1/certificates/" + id + "/callback")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// AddCallback registers a certificate action callback
func AddCallback(c *godaddy.Client, id string, cb string) error {
	uri, err := godaddy.BuildQuery("/v1/certificates/"+id+"/callback",
		map[string]string{"callbackUrl": cb})
	if err != nil {
		return err
	}
	_, err = c.Put(uri, nil)

	return err
}

// Cancel a pending certificate
func Cancel(c *godaddy.Client, id string) error {
	_, err := c.Post("/v1/certificates/"+id+"/cancel", nil)

	return err
}

// Download certificate
func Download(c *godaddy.Client, id string) (*Bundle, error) {
	res := new(Bundle)

	data, err := c.Get("/v1/certificates/" + id + "/download")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// Reissue the active certificate
func Reissue(c *godaddy.Client, id string, body *CertificateReissue) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = c.Post("/v1/certificates/"+id+"/reissue", enc)

	return err
}

// Renew the active certificate
func Renew(c *godaddy.Client, id string, body *CertificateRenew) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = c.Post("/v1/certificates/"+id+"/renew", enc)

	return err
}

// Revoke the active certificate
func Revoke(c *godaddy.Client, id string, body *CertificateRevoke) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = c.Post("/v1/certificates/"+id+"/revoke", enc)

	return err
}

// SiteSeal generates and retrieves the site seal
func SiteSeal(c *godaddy.Client, id string, theme string, locale string) (*CertificateSiteSeal, error) {
	res := new(CertificateSiteSeal)

	data, err := c.Get("/v1/certificates/" + id + "/siteSeal")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// VerifyDomainControl confirms domain control
func VerifyDomainControl(c *godaddy.Client, id string) error {
	_, err := c.Post("/v1/certificates/"+id+"/verifyDomainControl", nil)

	return err
}
