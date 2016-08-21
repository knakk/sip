package sip

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// Decode decodes a SIP message.
// Decode does validate a SIP message according to the required fields, other
// than an initial check of length to see if the message is long enough to
// contain the required fields for the given message type. To validate a
// Message use the Message.Validate() function.
func Decode(msg []byte) (Message, error) {
	var m Message

	l := len(msg)

	// trim trailing carriage return if present
	if l > 0 && msg[l-1] == '\r' {
		l--
		msg = msg[:l]
	}

	if l < 2 {
		return m, errors.New("message too short")
	}

	m.Type = codeToMsg[string(msg[:2])]
	if m.Type == MsgUnknown {
		return m, fmt.Errorf("unknown message code: %q", string(msg[:2]))
	}

	if l < minMsgLength[m.Type] {
		return m, fmt.Errorf("message too short to include required fields for %v: %d < %d", m.Type, len(msg), minMsgLength[m.Type])
	}

	m.Fields = make(map[fieldType]string)
	//m.RepeateableFields = make(map[fieldType][]string)

	p := 2 // byte position in message
	for _, f := range msgDefinitions[m.Type].RequiredFixed {
		end := p + fixedFieldLengths[f] // end of token
		m.Fields[f] = string(msg[p:end])
		p = end
	}
outer:
	for {
		start := p + 2 // start of current field
		f := codeToField[string(msg[p:start])]
		p = start
		for {
			r, w := utf8.DecodeRune(msg[p:])
			p += w
			if r == '|' {
				m.Fields[f] = string(msg[start : p-1])
				if p == len(msg) {
					break outer
				}
				continue outer
			}

		}
	}

	return m, nil

}
