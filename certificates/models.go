package certificates

import (
	godaddy "github.com/alyx/go-daddy"
)

type Certificate struct {
	CertificateID           string
	CommonName              string
	Contact                 Contact
	CreatedAt               string
	DeniedReason            string
	Organization            Organization
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

type Action struct {
	CreatedAt string
	Type      string
}

type Bundle struct {
	PEMs         PEMCertificates
	SerialNumber string
}

type Callback struct {
	CallbackURL string
}

type Contact struct {
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
	Organization            OrganizationCreate
	Period                  int
	ProductType             string
	RootType                string
	SlotSize                string
	SubjectAlternativeNames []string
}

type Identifier struct {
	CertificateID string
}

type Organization struct {
	OrganizationCreate
	JurisdictionOfIncorporation JurisdictionOfIncorporation
}

type OrganizationCreate struct {
	Address            godaddy.Address
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
