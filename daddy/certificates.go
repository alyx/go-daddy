// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

import (
	"encoding/json"
)

// CertificatesService handles communication with the certificates related
// methods of the GoDaddy API.
//
// GoDaddy API docs: https://developer.godaddy.com/doc/endpoint/certificates
type CertificatesService service

// Create will create a pending order for a certificate
func (s *CertificatesService) Create(body *CertificateCreate) (*CertificateIdentifier, error) {
	var res = new(CertificateIdentifier)

	enc, err := json.Marshal(body)
	if err != nil {
		return res, err
	}

	data, err := s.client.Post("/v1/certificates", enc)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// Validate will validate a pending order for a certificate
func (s *CertificatesService) Validate(body *CertificateCreate) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = s.client.Post("/v1/certificates/validate", enc)
	if err != nil {
		return err
	}

	return err
}

// Get retrieves certificate details for a specified certificate
func (s *CertificatesService) Get(id string) (*Certificate, error) {
	res := new(Certificate)

	data, err := s.client.Get("/v1/certificates/" + id)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// GetActions retrieves all certificate actions
func (s *CertificatesService) GetActions(id string) ([]CertificateAction, error) {
	var res []CertificateAction

	data, err := s.client.Get("/v1/certificates/" + id + "/actions")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// ResendEmail resends a validation email
func (s *CertificatesService) ResendEmail(cert string, email string) error {
	_, err := s.client.Post("/v1/certificates/"+cert+"/email/"+email+"/resend", nil)

	return err
}

// AddEmail adds an alternate email address
func (s *CertificatesService) AddEmail(cert string, email string) (*EmailHistory, error) {
	res := new(EmailHistory)

	data, err := s.client.Post("/v1/certificates/"+cert+"/email/resend/"+email, nil)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// ResendEmailByID resents a specific email to an email address
func (s *CertificatesService) ResendEmailByID(cert string, emailID string, email string) error {
	_, err := s.client.Post("/v1/certificates/"+cert+"/email/"+emailID+"/resend/"+email, nil)

	return err
}

// History retrieves email history
func (s *CertificatesService) History(id string) (*EmailHistory, error) {
	res := new(EmailHistory)

	data, err := s.client.Get("/v1/certificates/" + id + "/email/history")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// DeleteCallback unregisters a system callback
func (s *CertificatesService) DeleteCallback(id string) error {
	_, err := s.client.Delete("/v1/certificates/"+id+"/callback", nil)

	return err
}

// GetCallback retrieves a system stateful action callback URL
func (s *CertificatesService) GetCallback(id string) (*CertificateCallback, error) {
	res := new(CertificateCallback)

	data, err := s.client.Get("/v1/certificates/" + id + "/callback")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// AddCallback registers a certificate action callback
func (s *CertificatesService) AddCallback(id string, cb string) error {
	uri, err := BuildQuery("/v1/certificates/"+id+"/callback",
		map[string]interface{}{"callbackUrl": cb})
	if err != nil {
		return err
	}
	_, err = s.client.Put(uri, nil)

	return err
}

// Cancel a pending certificate
func (s *CertificatesService) Cancel(id string) error {
	_, err := s.client.Post("/v1/certificates/"+id+"/cancel", nil)

	return err
}

// Download certificate
func (s *CertificatesService) Download(id string) (*CertificateBundle, error) {
	res := new(CertificateBundle)

	data, err := s.client.Get("/v1/certificates/" + id + "/download")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// Reissue the active certificate
func (s *CertificatesService) Reissue(id string, body *CertificateReissue) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = s.client.Post("/v1/certificates/"+id+"/reissue", enc)

	return err
}

// Renew the active certificate
func (s *CertificatesService) Renew(id string, body *CertificateRenew) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = s.client.Post("/v1/certificates/"+id+"/renew", enc)

	return err
}

// Revoke the active certificate
func (s *CertificatesService) Revoke(id string, body *CertificateRevoke) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = s.client.Post("/v1/certificates/"+id+"/revoke", enc)

	return err
}

// SiteSeal generates and retrieves the site seal
func (s *CertificatesService) SiteSeal(id string, theme string, locale string) (*CertificateSiteSeal, error) {
	res := new(CertificateSiteSeal)

	data, err := s.client.Get("/v1/certificates/" + id + "/siteSeal")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

// VerifyDomainControl confirms domain control
func (s *CertificatesService) VerifyDomainControl(id string) error {
	_, err := s.client.Post("/v1/certificates/"+id+"/verifyDomainControl", nil)

	return err
}
