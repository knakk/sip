// package sip is a library for working with the SIP protocol used for
// communication between self-service circulation automats and library vendor systems.
//
// Definitions:
//   ACS = Automated Circulation System (SIP-server)
//   SC = Self Checkout Machine

package sip

// A Message defines a request sent by the SC to the ACS, or the response
// sent by the ACS to the SC. The Message Type defines which kind it is.
type Message struct {
	Type              msgType
	Fields            map[fieldType]string
	RepeateableFields map[fieldType][]string
}

/*

// Valid returns true if Message is valid. That is, if the Type is not MsgUnknown, and
// all the required fields are present. It does not validate that the contents of
// the fields conforms to expected patterns.
func (m Message) Valid() bool                  { return false }
func (m Message) Field(t fieldType) []string   { return nil }
func (m Message) HasField(t fieldType) bool    { return false }
func (m Message) MustField(t fieldType) string { return "" }
func (m Message) Encode(w io.Writer) error     { return nil }

*/

// msgType represents the request or response message type.
type msgType int

// All possible message types
const (
	MsgUnknown msgType = iota
	// Requests:
	MsgReqPatronStatus      // 23
	MsgReqCheckout          // 11
	MsgReqCheckin           // 09
	MsgReqBlockPatron       // 01
	MsgReqStatus            // 99
	MsgReqResend            // 97
	MsgReqLogin             // 93
	MsgReqPatronInformation // 63
	MsgReqEndPatronSession  // 35
	MsgReqFeePaid           // 37
	MsgReqItemInformation   // 17
	MsgReqItemStatusUpdate  // 19
	MsgReqPatronEnable      // 25
	MsgReqHold              // 15
	MsgReqRenew             // 29
	MsgReqRenewAll          // 65
	// Responses:
	MsgRespPatronStatus      // 24
	MsgRespCheckout          // 12
	MsgRespCheckin           // 10
	MsgRespStatus            // 98
	MsgRespLogin             // 94
	MsgRespPatronInformation // 64
	MsgRespEndPatronSession  // 36
	MsgRespFeePaid           // 38
	MsgRespItemInformation   // 18
	MsgRespItemStatusUpdate  // 20
	MsgRespPatronEnable      // 26
	MsgRespHold              // 16
	MsgRespRenew             // 30
	MsgRespRenewAll          // 66
)

// fieldType represents a request/response message field ID.
type fieldType int

// All possible field types:
const (
	FieldUnknown fieldType = iota

	// Fixed length fields identified by position in request/response header:
	FieldAlert                 // 1 char: Y or N
	FieldAvialable             // 1 char: Y or N
	FieldCardRetained          // 1 char: Y or N
	FieldChargedItemsCount     // 4 char: 0000 to 9999
	FieldFineItemsCount        // 4 char: 0000 to 9999
	FieldHoldItemsCount        // 4 char: 0000 to 9999
	FieldOverdueItemsCount     // 4 char: 0000 to 9999
	FieldRecallItemsCount      // 4 char: 0000 to 9999
	FieldUnavailableHoldsCount // 4 char: 0000 to 9999
	FieldCheckinOK             // 1 char: Y or N
	FieldCheckoutOK            // 1 char: Y or N
	FieldCirulationStatus      // 2 char: 00 to 99 (1-13 are defined by the protocol)
	FieldDateTimeSync          // 18 char: YYYYMMDDZZZZHHMMSS
	FieldDesentisize           // 1 char: Y or N
	FieldEndSession            // 1 char: Y or N
	FieldHoldMode              // 1 char: + or - or *
	FieldItemPropertiesOK      // 1 char: '1' = OK, other = not NOK
	FieldLanguage              // 3 char: 000 - 999 (some language codes defined by the protocol)
	FieldMagneticMedia         // 1 char: Y or N or U
	FieldMaxPrintWidth         // 3 char: 000 to 999
	FieldNbDueDate             // 1 char: Y or N
	FieldNoBlock               // 1 char: Y or N
	FieldOffLineOK             // 1 char: Y or N
	FieldOK                    // 1 char: '1' = OK, other = not OK
	FieldOnLineStatus          // 1 char: Y or N
	FieldPatronStatus          // 14 char: each char is Y or N
	FieldPaymentAccepted       // 1 char: Y or N
	FieldPaymentType           // 2 char: 00=cash, 01=VISA, 02=credit card
	FieldProtocolVersion       // 4 char: x.xx
	FieldPWDAlgorithm          // 1 char: '0' = not encrypted
	FieldRenewalOK             // 1 char: Y or N
	FieldRenewedCount          // 4 char: 0000 to 9999
	FieldResentisize           // 1 char: Y or N
	FieldRetriesAllowd         // 3 char: ?
	FieldReturnDate            // 18 char: YYYYMMDDZZZZHHMMSS
	FieldRenewalPolicy         // 1 char: Y or N
	FieldSecurityMarker        // 2 char: 00=other, 01=none, 02,03=3M specific
	FieldStatusCode            // 1 char: 0=OK, 1=printer out of paper, 2=about to shut down
	FieldStatusUpdateOK        // 1 char: Y or N
	FieldSummary               // 10 char:
	FieldThirdPartyAllowd      // 1 char: Y or N
	FieldTimeoutPeriod         // 3 char: 000 - 999
	FieldTransactionDate       // 18 char: YYYYMMDDZZZZHHMMSS
	FieldUIDAlgorithm          // 1 char: 0=not encrypted
	FieldUnrenewedCount        // 4 char: 0000 to 9999

	// Variable length fields identified by a two-character code:
	FieldPatronIdentifier      // AA
	FieldItemIdentifier        // AB
	FieldTerminalPassword      // AC
	FieldPatronPassword        // AD
	FieldPersonalName          // AE
	FieldScreenMessage         // AF
	FieldPrintLine             // AG
	FieldDueDate               // AH
	FieldTitleIdentifier       // AJ
	FieldBlockedCardMsg        // AL
	FieldLibraryName           // AM
	FieldTerminalLocation      // AN
	FieldInstitutionID         // AO
	FieldCurrentLocation       // AP
	FieldPermanentLocation     // AQ
	FieldHoldItems             // AS
	FieldOverdueItems          // AT
	FieldChargedItems          // AU
	FieldFineItems             // AV
	FieldHomeAddress           // BD
	FieldEmailAddress          // BE
	FieldHomePhoneNumber       // BF
	FieldOwner                 // BG
	FieldCurrenyType           // BH
	FieldCancel                // BI
	FieldTransactionID         // BK
	FieldValidPatron           // BL
	FieldRenewedItems          // BM
	FieldUnrenewedItems        // BN
	FieldFeeAcknowledged       // BO
	FieldStartItem             // BP
	FieldEndItem               // BQ
	FieldQueuePosition         // BR
	FieldPickupLocation        // BS
	FieldFeeType               // BT
	FieldRecallItems           // BU
	FieldFeeAmount             // BV
	FieldExpirationDate        // BW
	FieldSupportedMessages     // BX
	FieldHoldType              // BY
	FieldHoldItemsLimit        // BZ
	FieldOverdueItemsLimit     // CA
	FieldChargedItemsLimit     // CB
	FieldFeeLimit              // CC
	FieldUnavailableHoldsItems // CD
	FieldHoldQueueLength       // CF
	FieldFeeIdentifier         // CG
	FieldItemProperties        // CH
	FieldSecurityInhibit       // CI
	FieldRecallDate            // CJ
	FieldMediaType             // CK
	FieldSortBin               // CL
	FieldHoldPickupDate        // CM
	FieldLoginUserID           // CN
	FieldLoginPassword         // CO
	FieldLocationCode          // CP
	FieldValidPatronPassword   // CQ
	FieldSequenceNumber        // AY
	FieldChecksum              // AZ
)
