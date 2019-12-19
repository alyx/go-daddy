// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

type Certificate struct {
	CertificateID           string
	CommonName              string
	Contact                 Contact
	CreatedAt               string
	DeniedReason            string
	Organization            CertificateOrganization
	Period                  int
	ProductType             string
	Progress                int
	RevokedAt               string
	RootType                string
	SerialNumber            string
	SerialNumberHex         string
	SlotSize                string
	Status                  string
	SubjectAlternativeNames []SubjectAlternativeNameDetails
	ValidEnd                string
	ValidStart              string
}

type CertificateAction struct {
	CreatedAt string
	Type      string
}

type CertificateBundle struct {
	PEMs         PEMCertificates
	SerialNumber string
}

type CertificateCallback struct {
	CallbackURL string
}

type CertificateContact struct {
	Email      string
	JobTitle   string
	NameFirst  string
	NameLast   string
	NameMiddle string
	Phone      string
	Suffix     string
}

type CertificateCreate struct {
	CallbackURL             string
	CommonName              string
	Contact                 Contact
	CSR                     string
	IntelVPro               bool
	Organization            CertificateOrganizationCreate
	Period                  int
	ProductType             string
	RootType                string
	SlotSize                string
	SubjectAlternativeNames []string
}

type CertificateIdentifier struct {
	CertificateID string
}

type CertificateOrganization struct {
	CertificateOrganizationCreate
	JurisdictionOfIncorporation JurisdictionOfIncorporation
}

type CertificateOrganizationCreate struct {
	Address            Address
	AssumedName        string
	Name               string
	Phone              string
	RegistrationAgent  string
	RegistrationNumber string
}

type CertificateReissue struct {
	CallbackURL             string
	CommonName              string
	CSR                     string
	DelayExistingRevoke     int
	RootType                string
	SubjectAlternativeNames []string
}

type CertificateRenew struct {
	CertificateReissue
}

type CertificateRevoke struct {
	Reason string
}

type CertificateSiteSeal struct {
	HTML string
}

type EmailHistory struct {
	ID           int
	AccountID    int
	TemplateType string
	FromType     string
	Recipients   string
	Body         string
	DateEntered  string
	Subject      string
}

type JurisdictionOfIncorporation struct {
	City    string
	Country string
	County  string
	State   string
}

type PEMCertificates struct {
	Certificate  string
	Cross        string
	Intermediate string
	Root         string
}

type SubjectAlternativeNameDetails struct {
	Status                 string
	SubjectAlternativeName string
}
