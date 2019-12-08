package domains

import "github.com/alyx/godaddy"

type Consent struct {
	AgreedAt      string
	AgreedBy      string
	AgreementKeys []string
}

type Contact struct {
	AddressMailing godaddy.Address
	Email          string
	Fax            string
	JobTitle       string
	NameFirst      string
	NameLast       string
	NameMiddle     string
	Organization   string
	Phone          string
}

type DNSRecord struct {
	Data     string
	Name     string
	Port     int
	Priority int
	Protocol int
	Service  int
	TTL      int
	Type     string
	Weight   int
}

type DomainAvailableResponse struct {
	Available  bool
	Currency   string
	Definitive bool
	Domain     string
	Period     int
	Price      int
}

type DomainAvailableError struct {
	Code    string
	Domain  string
	Message string
	Path    string
	Status  int
}

type DomainAvailableBulkMixed struct {
	Domains []DomainAvailableResponse
	Errors  []DomainAvailableError
}

type DomainContacts struct {
	Admin      Contact
	Billing    Contact
	Registrant Contact
	Tech       Contact
}

type DomainDetail struct {
	DomainSummary
	SubaccountID  string
	Verifications VerificationsDomain
}

type DomainPurchase struct {
	Consent Consent
	DomainContacts
	Domain      string
	NameServers []string
	Period      int
	Privacy     bool
	RenewAuto   bool
}

type DomainPurchaseResponse struct {
	Currency  string
	ItemCount int
	OrderID   int
	Total     int
}

type DomainRenew struct {
	Period int
}

type DomainSuggestion struct {
	Domain string
}

type DomainSummary struct {
	AuthCode string
	DomainContacts
	CreatedAt               string
	DeletedAt               string
	TransferAwayEligibileAt string
	Domain                  string
	DomainID                int
	ExpirationProtected     bool
	Expires                 string
	HoldRegistrar           bool
	Locked                  bool
	NameServers             []string
	Privacy                 bool
	RenewAuto               bool
	RenewDeadline           string
	Renewable               bool
	Status                  string
	TransferProtected       bool
}

type DomainForwardingMask struct {
	Title       string
	Description string
	Keywords    string
}

type DomainForwardingCreate struct {
	Type string
	URL  string
	Mask DomainForwardingMask
}

type DomainForwarding struct {
	FQDN string
	DomainForwardingCreate
}

type DomainTransferIn struct {
	AuthCode  string
	Consent   Consent
	Period    int
	Privacy   bool
	RenewAuto bool
	DomainContacts
}

type DomainUpdate struct {
	Locked       bool
	NameServers  []string
	RenewAuto    bool
	SubaccountID string
}

type DomainContactsBulk struct {
	DomainContacts
	ContactPresence Contact
	Domains         []string
	EntityType      string
}

type ErrorDomainContactsValidate struct {
	Code    string
	Fields  []ErrorFieldDomainContactsValidate
	Message string
	Stack   []string
}

type ErrorFieldDomainContactsValidate struct {
	godaddy.ErrorField
	Domains []string
}

type LegalAgreement struct {
	AgreementKey string
	Content      string
	Title        string
	URL          string
}

type PrivacyPurchase struct {
	Consent Consent
}

type RealNameValidation struct {
	Status string
}

type TldSummary struct {
	Name string
	Type string
}

type VerificationDomainName struct {
	Status string
}

type VerificationRealName struct {
	Status string
}

type VerificationsDomain struct {
	DomainName VerificationDomainName
	RealName   VerificationRealName
}

const (
	Active                                  = "ACTIVE"
	AwaitingClaimAck                        = "AWAITING_CLAIM_ACK"
	AwaitingDocumentAfterTransfer           = "AWAITING_DOCUMENT_AFTER_TRANSFER"
	AwaitingDocumentAfterUpdateAccount      = "AWAITING_DOCUMENT_AFTER_UPDATE_ACCOUNT"
	AwaitingDocumentUpload                  = "AWAITING_DOCUMENT_UPLOAD"
	AwaitingFailedTransferWhoisPrivacy      = "AWAITING_FAILED_TRANSFER_WHOIS_PRIVACY"
	AwaitingPayment                         = "AWAITING_PAYMENT"
	AwaitingRenewalTransferInComplete       = "AWAITING_RENEWAL_TRANSFER_IN_COMPLETE"
	AwaitingTransferInAck                   = "AWAITING_TRANSFER_IN_ACK"
	AwaitingTransferInAuth                  = "AWAITING_TRANSFER_IN_AUTH"
	AwaitingTransferInAuto                  = "AWAITING_TRANSFER_IN_AUTO"
	AwaitingTransferInWhois                 = "AWAITING_TRANSFER_IN_WHOIS"
	AwaitingTransferInWhoisFix              = "AWAITING_TRANSFER_IN_WHOIS_FIX"
	AwaitingVerificationIcann               = "AWAITING_VERIFICATION_ICANN"
	AwaitingVerificationIcannManual         = "AWAITING_VERIFICATION_ICANN_MANUAL"
	Cancelled                               = "CANCELLED"
	CancelledHeld                           = "CANCELLED_HELD"
	CancelledRedeemable                     = "CANCELLED_REDEEMABLE"
	CancelledTransfer                       = "CANCELLED_TRANSFER"
	Confiscated                             = "CONFISCATED"
	DisabledSpecial                         = "DISABLED_SPECIAL"
	ExcludedInvalidClaimFirehose            = "EXCLUDED_INVALID_CLAIM_FIREHOSE"
	ExpiredReassigned                       = "EXPIRED_REASSIGNED"
	FailedBackorderCapture                  = "FAILED_BACKORDER_CAPTURE"
	FailedDropImmediateThenAdd              = "FAILED_DROP_IMMEDIATE_THEN_ADD"
	FailedPreRegistration                   = "FAILED_PRE_REGISTRATION"
	FailedRedemption                        = "FAILED_REDEMPTION"
	FailedRedemptionReport                  = "FAILED_REDEMPTION_REPORT"
	FailedRegistration                      = "FAILED_REGISTRATION"
	FailedRegistrationFirehose              = "FAILED_REGISTRATION_FIREHOSE"
	FailedRestorationRedemptionMock         = "FAILED_RESTORATION_REDEMPTION_MOCK"
	FailedSetup                             = "FAILED_SETUP"
	FailedTransferIn                        = "FAILED_TRANSFER_IN"
	FailedTransferInBadStatus               = "FAILED_TRANSFER_IN_BAD_STATUS"
	FailedTransferInRegistry                = "FAILED_TRANSFER_IN_REGISTRY"
	HeldCourtOrdered                        = "HELD_COURT_ORDERED"
	HeldDisputed                            = "HELD_DISPUTED"
	HeldExpirationProtection                = "HELD_EXPIRATION_PROTECTION"
	HeldExpiredRedemptionMock               = "HELD_EXPIRED_REDEMPTION_MOCK"
	HeldRegistrarAdd                        = "HELD_REGISTRAR_ADD"
	HeldRegistrarRemove                     = "HELD_REGISTRAR_REMOVE"
	HeldShopper                             = "HELD_SHOPPER"
	HeldTemporary                           = "HELD_TEMPORARY"
	LockedAbuse                             = "LOCKED_ABUSE"
	LockedCopyright                         = "LOCKED_COPYRIGHT"
	LockedRegistry                          = "LOCKED_REGISTRY"
	LockedSuper                             = "LOCKED_SUPER"
	ParkedAndHeld                           = "PARKED_AND_HELD"
	ParkedExpired                           = "PARKED_EXPIRED"
	ParkedVerificationIcann                 = "PARKED_VERIFICATION_ICANN"
	PendingAbortCancelSetup                 = "PENDING_ABORT_CANCEL_SETUP"
	PendingAgreementPreRegistration         = "PENDING_AGREEMENT_PRE_REGISTRATION"
	PendingApplyRenewalCredits              = "PENDING_APPLY_RENEWAL_CREDITS"
	PendingBackorderCapture                 = "PENDING_BACKORDER_CAPTURE"
	PendingBlockedRegistry                  = "PENDING_BLOCKED_REGISTRY"
	PendingCancelRegistrantProfile          = "PENDING_CANCEL_REGISTRANT_PROFILE"
	PendingCompleteRedemptionWithoutReceipt = "PENDING_COMPLETE_REDEMPTION_WITHOUT_RECEIPT"
	PendingCompleteRegistrantProfile        = "PENDING_COMPLETE_REGISTRANT_PROFILE"
	PendingCoo                              = "PENDING_COO"
	PendingCooComplete                      = "PENDING_COO_COMPLETE"
	PendingDNS                              = "PENDING_DNS"
	PendingDNSActive                        = "PENDING_DNS_ACTIVE"
	PendingDNSInactive                      = "PENDING_DNS_INACTIVE"
	PendingDocumentValidation               = "PENDING_DOCUMENT_VALIDATION"
	PendingDocumentVerification             = "PENDING_DOCUMENT_VERIFICATION"
	PendingDropImmediate                    = "PENDING_DROP_IMMEDIATE"
	PendingDropImmediateThenAdd             = "PENDING_DROP_IMMEDIATE_THEN_ADD"
	PendingEppCreate                        = "PENDING_EPP_CREATE"
	PendingEppDelete                        = "PENDING_EPP_DELETE"
	PendingEppUpdate                        = "PENDING_EPP_UPDATE"
	PendingEscalationRegistry               = "PENDING_ESCALATION_REGISTRY"
	PendingExpiration                       = "PENDING_EXPIRATION"
	PendingExpirationResponse               = "PENDING_EXPIRATION_RESPONSE"
	PendingExpirationSync                   = "PENDING_EXPIRATION_SYNC"
	PendingExpiredReassignment              = "PENDING_EXPIRED_REASSIGNMENT"
	PendingExpireAutoAdd                    = "PENDING_EXPIRE_AUTO_ADD"
	PendingExtendRegistrantProfile          = "PENDING_EXTEND_REGISTRANT_PROFILE"
	PendingFailedCoo                        = "PENDING_FAILED_COO"
	PendingFailedEppCreate                  = "PENDING_FAILED_EPP_CREATE"
	PendingFailedHeld                       = "PENDING_FAILED_HELD"
	PendingFailedPurchasePremium            = "PENDING_FAILED_PURCHASE_PREMIUM"
	PendingFailedReconcileFirehose          = "PENDING_FAILED_RECONCILE_FIREHOSE"
	PendingFailedRedemptionWithoutReceipt   = "PENDING_FAILED_REDEMPTION_WITHOUT_RECEIPT"
	PendingFailedReleasePremium             = "PENDING_FAILED_RELEASE_PREMIUM"
	PendingFailedRenewExpirationProtection  = "PENDING_FAILED_RENEW_EXPIRATION_PROTECTION"
	PendingFailedReservePremium             = "PENDING_FAILED_RESERVE_PREMIUM"
	PendingFailedSubmitFirehose             = "PENDING_FAILED_SUBMIT_FIREHOSE"
	PendingFailedTransferAckPremium         = "PENDING_FAILED_TRANSFER_ACK_PREMIUM"
	PendingFailedTransferInAckPremium       = "PENDING_FAILED_TRANSFER_IN_ACK_PREMIUM"
	PendingFailedTransferInPremium          = "PENDING_FAILED_TRANSFER_IN_PREMIUM"
	PendingFailedTransferPremium            = "PENDING_FAILED_TRANSFER_PREMIUM"
	PendingFailedTransferSubmitPremium      = "PENDING_FAILED_TRANSFER_SUBMIT_PREMIUM"
	PendingFailedUnlockPremium              = "PENDING_FAILED_UNLOCK_PREMIUM"
	PendingFailedUpdateAPI                  = "PENDING_FAILED_UPDATE_API"
	PendingFraudVerification                = "PENDING_FRAUD_VERIFICATION"
	PendingFraudVerified                    = "PENDING_FRAUD_VERIFIED"
	PendingGetContacts                      = "PENDING_GET_CONTACTS"
	PendingGetHosts                         = "PENDING_GET_HOSTS"
	PendingGetNameServers                   = "PENDING_GET_NAME_SERVERS"
	PendingGetStatus                        = "PENDING_GET_STATUS"
	PendingHoldEscrow                       = "PENDING_HOLD_ESCROW"
	PendingHoldRedemption                   = "PENDING_HOLD_REDEMPTION"
	PendingLockClientRemove                 = "PENDING_LOCK_CLIENT_REMOVE"
	PendingLockDataQuality                  = "PENDING_LOCK_DATA_QUALITY"
	PendingLockThenHoldRedemption           = "PENDING_LOCK_THEN_HOLD_REDEMPTION"
	PendingParkingDetermination             = "PENDING_PARKING_DETERMINATION"
	PendingParkInvalidWhois                 = "PENDING_PARK_INVALID_WHOIS"
	PendingParkInvalidWhoisRemoval          = "PENDING_PARK_INVALID_WHOIS_REMOVAL"
	PendingPurchasePremium                  = "PENDING_PURCHASE_PREMIUM"
	PendingReconcile                        = "PENDING_RECONCILE"
	PendingReconcileFirehose                = "PENDING_RECONCILE_FIREHOSE"
	PendingRedemption                       = "PENDING_REDEMPTION"
	PendingRedemptionReport                 = "PENDING_REDEMPTION_REPORT"
	PendingRedemptionReportComplete         = "PENDING_REDEMPTION_REPORT_COMPLETE"
	PendingRedemptionReportSubmitted        = "PENDING_REDEMPTION_REPORT_SUBMITTED"
	PendingRedemptionWithoutReceipt         = "PENDING_REDEMPTION_WITHOUT_RECEIPT"
	PendingRedemptionWithoutReceiptMock     = "PENDING_REDEMPTION_WITHOUT_RECEIPT_MOCK"
	PendingReleasePremium                   = "PENDING_RELEASE_PREMIUM"
	PendingRemoval                          = "PENDING_REMOVAL"
	PendingRemovalHeld                      = "PENDING_REMOVAL_HELD"
	PendingRemovalParked                    = "PENDING_REMOVAL_PARKED"
	PendingRemovalUnpark                    = "PENDING_REMOVAL_UNPARK"
	PendingRenewal                          = "PENDING_RENEWAL"
	PendingRenewExpirationProtection        = "PENDING_RENEW_EXPIRATION_PROTECTION"
	PendingRenewInfinite                    = "PENDING_RENEW_INFINITE"
	PendingRenewLocked                      = "PENDING_RENEW_LOCKED"
	PendingRenewWithoutReceipt              = "PENDING_RENEW_WITHOUT_RECEIPT"
	PendingReportRedemptionWithoutReceipt   = "PENDING_REPORT_REDEMPTION_WITHOUT_RECEIPT"
	PendingReservePremium                   = "PENDING_RESERVE_PREMIUM"
	PendingResetVerificationIcann           = "PENDING_RESET_VERIFICATION_ICANN"
	PendingResponseFirehose                 = "PENDING_RESPONSE_FIREHOSE"
	PendingRestoration                      = "PENDING_RESTORATION"
	PendingRestorationInactive              = "PENDING_RESTORATION_INACTIVE"
	PendingRestorationRedemptionMock        = "PENDING_RESTORATION_REDEMPTION_MOCK"
	PendingRetryEppCreate                   = "PENDING_RETRY_EPP_CREATE"
	PendingRetryHeld                        = "PENDING_RETRY_HELD"
	PendingSendAuthCode                     = "PENDING_SEND_AUTH_CODE"
	PendingSetup                            = "PENDING_SETUP"
	PendingSetupAbandon                     = "PENDING_SETUP_ABANDON"
	PendingSetupAgreementLandrush           = "PENDING_SETUP_AGREEMENT_LANDRUSH"
	PendingSetupAgreementSunrise2A          = "PENDING_SETUP_AGREEMENT_SUNRISE2_A"
	PendingSetupAgreementSunrise2B          = "PENDING_SETUP_AGREEMENT_SUNRISE2_B"
	PendingSetupAgreementSunrise2C          = "PENDING_SETUP_AGREEMENT_SUNRISE2_C"
	PendingSetupAuth                        = "PENDING_SETUP_AUTH"
	PendingSetupDNS                         = "PENDING_SETUP_DNS"
	PendingSetupFailed                      = "PENDING_SETUP_FAILED"
	PendingSetupReview                      = "PENDING_SETUP_REVIEW"
	PendingSetupSunrise                     = "PENDING_SETUP_SUNRISE"
	PendingSetupSunrisePre                  = "PENDING_SETUP_SUNRISE_PRE"
	PendingSetupSunriseResponse             = "PENDING_SETUP_SUNRISE_RESPONSE"
	PendingSubmitFailure                    = "PENDING_SUBMIT_FAILURE"
	PendingSubmitFirehose                   = "PENDING_SUBMIT_FIREHOSE"
	PendingSubmitHoldFirehose               = "PENDING_SUBMIT_HOLD_FIREHOSE"
	PendingSubmitHoldLandrush               = "PENDING_SUBMIT_HOLD_LANDRUSH"
	PendingSubmitHoldSunrise                = "PENDING_SUBMIT_HOLD_SUNRISE"
	PendingSubmitLandrush                   = "PENDING_SUBMIT_LANDRUSH"
	PendingSubmitResponseFirehose           = "PENDING_SUBMIT_RESPONSE_FIREHOSE"
	PendingSubmitResponseLandrush           = "PENDING_SUBMIT_RESPONSE_LANDRUSH"
	PendingSubmitResponseSunrise            = "PENDING_SUBMIT_RESPONSE_SUNRISE"
	PendingSubmitSuccessFirehose            = "PENDING_SUBMIT_SUCCESS_FIREHOSE"
	PendingSubmitSuccessLandrush            = "PENDING_SUBMIT_SUCCESS_LANDRUSH"
	PendingSubmitSuccessSunrise             = "PENDING_SUBMIT_SUCCESS_SUNRISE"
	PendingSubmitSunrise                    = "PENDING_SUBMIT_SUNRISE"
	PendingSubmitWaitingLandrush            = "PENDING_SUBMIT_WAITING_LANDRUSH"
	PendingSuccessPreRegistration           = "PENDING_SUCCESS_PRE_REGISTRATION"
	PendingSuspendedDataQuality             = "PENDING_SUSPENDED_DATA_QUALITY"
	PendingTransferAckPremium               = "PENDING_TRANSFER_ACK_PREMIUM"
	PendingTransferIn                       = "PENDING_TRANSFER_IN"
	PendingTransferInAck                    = "PENDING_TRANSFER_IN_ACK"
	PendingTransferInAckPremium             = "PENDING_TRANSFER_IN_ACK_PREMIUM"
	PendingTransferInBadRegistrant          = "PENDING_TRANSFER_IN_BAD_REGISTRANT"
	PendingTransferInCancel                 = "PENDING_TRANSFER_IN_CANCEL"
	PendingTransferInCancelRegistry         = "PENDING_TRANSFER_IN_CANCEL_REGISTRY"
	PendingTransferInCompleteAck            = "PENDING_TRANSFER_IN_COMPLETE_ACK"
	PendingTransferInDelete                 = "PENDING_TRANSFER_IN_DELETE"
	PendingTransferInLock                   = "PENDING_TRANSFER_IN_LOCK"
	PendingTransferInNack                   = "PENDING_TRANSFER_IN_NACK"
	PendingTransferInNotification           = "PENDING_TRANSFER_IN_NOTIFICATION"
	PendingTransferInPremium                = "PENDING_TRANSFER_IN_PREMIUM"
	PendingTransferInRelease                = "PENDING_TRANSFER_IN_RELEASE"
	PendingTransferInResponse               = "PENDING_TRANSFER_IN_RESPONSE"
	PendingTransferInUnderage               = "PENDING_TRANSFER_IN_UNDERAGE"
	PendingTransferOut                      = "PENDING_TRANSFER_OUT"
	PendingTransferOutAck                   = "PENDING_TRANSFER_OUT_ACK"
	PendingTransferOutNack                  = "PENDING_TRANSFER_OUT_NACK"
	PendingTransferOutPremium               = "PENDING_TRANSFER_OUT_PREMIUM"
	PendingTransferOutUnderage              = "PENDING_TRANSFER_OUT_UNDERAGE"
	PendingTransferOutValidation            = "PENDING_TRANSFER_OUT_VALIDATION"
	PendingTransferPremium                  = "PENDING_TRANSFER_PREMIUM"
	PendingTransferPremuim                  = "PENDING_TRANSFER_PREMUIM"
	PendingTransferSubmitPremium            = "PENDING_TRANSFER_SUBMIT_PREMIUM"
	PendingUnlockDataQuality                = "PENDING_UNLOCK_DATA_QUALITY"
	PendingUnlockPremium                    = "PENDING_UNLOCK_PREMIUM"
	PendingUpdate                           = "PENDING_UPDATE"
	PendingUpdatedRegistrantDataQuality     = "PENDING_UPDATED_REGISTRANT_DATA_QUALITY"
	PendingUpdateAccount                    = "PENDING_UPDATE_ACCOUNT"
	PendingUpdateAPI                        = "PENDING_UPDATE_API"
	PendingUpdateAPIResponse                = "PENDING_UPDATE_API_RESPONSE"
	PendingUpdateAuth                       = "PENDING_UPDATE_AUTH"
	PendingUpdateContacts                   = "PENDING_UPDATE_CONTACTS"
	PendingUpdateContactsPrivacy            = "PENDING_UPDATE_CONTACTS_PRIVACY"
	PendingUpdateDNS                        = "PENDING_UPDATE_DNS"
	PendingUpdateDNSSecurity                = "PENDING_UPDATE_DNS_SECURITY"
	PendingUpdateEligibility                = "PENDING_UPDATE_ELIGIBILITY"
	PendingUpdateEppContacts                = "PENDING_UPDATE_EPP_CONTACTS"
	PendingUpdateMembership                 = "PENDING_UPDATE_MEMBERSHIP"
	PendingUpdateOwnership                  = "PENDING_UPDATE_OWNERSHIP"
	PendingUpdateOwnershipAuthAuction       = "PENDING_UPDATE_OWNERSHIP_AUTH_AUCTION"
	PendingUpdateOwnershipHeld              = "PENDING_UPDATE_OWNERSHIP_HELD"
	PendingUpdateRegistrant                 = "PENDING_UPDATE_REGISTRANT"
	PendingUpdateRepo                       = "PENDING_UPDATE_REPO"
	PendingValidationDataQuality            = "PENDING_VALIDATION_DATA_QUALITY"
	PendingVerificationFraud                = "PENDING_VERIFICATION_FRAUD"
	PendingVerificationStatus               = "PENDING_VERIFICATION_STATUS"
	PendingVerifyRegistrantDataQuality      = "PENDING_VERIFY_REGISTRANT_DATA_QUALITY"
	Reserved                                = "RESERVED"
	ReservedPremium                         = "RESERVED_PREMIUM"
	Reverted                                = "REVERTED"
	SuspendedVerificationIcann              = "SUSPENDED_VERIFICATION_ICANN"
	TransferredOut                          = "TRANSFERRED_OUT"
	UnlockedAbuse                           = "UNLOCKED_ABUSE"
	UnlockedSuper                           = "UNLOCKED_SUPER"
	UnparkedAndUnheld                       = "UNPARKED_AND_UNHELD"
	UpdatedOwnership                        = "UPDATED_OWNERSHIP"
	UpdatedOwnershipHeld                    = "UPDATED_OWNERSHIP_HELD"
)
