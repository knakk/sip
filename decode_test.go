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
