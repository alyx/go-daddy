package agreements

// LegalAgreement represents a singular agreement document
// (such as Terms of Service, etc.)
type LegalAgreement struct {
	AgreementKey string
	Content      string
	Title        string
	URL          string
}
