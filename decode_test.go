package sip

import (
	"reflect"
	"testing"
)

func TestDecodeMessages(t *testing.T) {
	tests := []struct {
		input string
		want  Message
	}{
		{
			"2300019810908ZZZZ124533AOAcme co.|AAp123.xyz|ACsecret|ADpass|",
			Message{
				Type: MsgReqPatronStatus,
				Fields: map[fieldType]string{
					FieldLanguage:         "000",
					FieldTransactionDate:  "19810908ZZZZ124533",
					FieldInstitutionID:    "Acme co.",
					FieldPatronIdentifier: "p123.xyz",
					FieldTerminalPassword: "secret",
					FieldPatronPassword:   "pass",
				},
				RepeateableFields: map[fieldType][]string{},
			},
		},
		{
			"11YN20131216    13531620131216    135316AO|AAN0012121212|ABix:1156620,1|ACsecret",
			Message{
				Type: MsgReqCheckout,
				Fields: map[fieldType]string{
					FieldRenewalPolicy:    "Y",
					FieldNoBlock:          "N",
					FieldTransactionDate:  "20131216    135316",
					FieldNbDueDate:        "20131216    135316",
					FieldInstitutionID:    "",
					FieldPatronIdentifier: "N0012121212",
					FieldItemIdentifier:   "ix:1156620,1",
					FieldTerminalPassword: "secret",
				},
				RepeateableFields: map[fieldType][]string{},
			},
		},
		{
			"120NNN20131216    140902AO|AAN000000001|ABRECEIPT|AJ|AH|AG DEICHMANSKE BIBLIOTEK|AG FURUSET|AG Tlf: 23 47 90 12|AG furu@deichman.no|AG www.deichman.no|AG|AGKvittering for utlån 16/12/2013|AGLånenummer: N000000001|AG|AG                  : Take this waltz |AG  (1357155.005) forfall : 23/12/2013|AG Dahl, Kjell Ola  : Svart engel : kr|AG  (766342.026) forfall : 13/01/2014|AG Egeland, Tom     : Ulvenatten : thr|AG  (664329.009) forfall : 13/01/2014|AG|AGTakk for besøket.|AG|AG|AG|AG|",
			Message{
				Type: MsgRespCheckout,
				Fields: map[fieldType]string{
					// fixed-length fields:
					FieldOK:              "0",
					FieldRenewalOK:       "N",
					FieldMagneticMedia:   "N",
					FieldDesentisize:     "N",
					FieldTransactionDate: "20131216    140902",
					// variable-length fields:
					FieldInstitutionID:    "",
					FieldPatronIdentifier: "N000000001",
					FieldItemIdentifier:   "RECEIPT",
					FieldTitleIdentifier:  "",
					FieldDueDate:          "",
				},
				RepeateableFields: map[fieldType][]string{
					FieldPrintLine: []string{
						" DEICHMANSKE BIBLIOTEK",
						" FURUSET",
						" Tlf: 23 47 90 12",
						" furu@deichman.no",
						" www.deichman.no",
						"",
						"Kvittering for utlån 16/12/2013",
						"Lånenummer: N000000001",
						"",
						"                  : Take this waltz ",
						"  (1357155.005) forfall : 23/12/2013",
						" Dahl, Kjell Ola  : Svart engel : kr",
						"  (766342.026) forfall : 13/01/2014",
						" Egeland, Tom     : Ulvenatten : thr",
						"  (664329.009) forfall : 13/01/2014",
						"",
						"Takk for besøket.",
						"",
						"",
						"",
						"",
					},
				},
			},
		},
	}

	for i, tt := range tests {
		got, err := Decode([]byte(tt.input))
		if err != nil {
			t.Errorf("%d: decode failed: %v", i, err)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: got:\n%v\nwant:\n%v", i, got, tt.want)
		}
	}
}
func TestDecodeErrors(t *testing.T) {
	tests := []struct {
		input     string
		errString string
	}{
		{"", "message too short"},
		{"1\r", "message too short"},
		{"xx", `unknown message code: "xx"`},
		{"3301000120131216    140859ABix:664329,9|AJEgeland, Tom : Ulvenatten|", `unknown message code: "33"`},
		{"2300019810908ZZZZ12453AO|AA|AC|AD|", "message too short to contain required fields for MsgReqPatronStatus: 34 < 35"},
	}

	for i, tt := range tests {
		_, err := Decode([]byte(tt.input))
		if err == nil {
			t.Errorf("%d, got nil; want %q", i, tt.errString)
			continue
		}
		if err.Error() != tt.errString {
			t.Errorf("%d: got %q; want %q", i, err.Error(), tt.errString)
		}
	}
}
