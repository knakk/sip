package sip

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Decode(msg []byte) (Message, error) {
	var m Message

	if len(msg) < 2 {
		return m, errors.New("message too short")
	}

	m.Type = codeToMsg[string(msg[:2])]
	if m.Type == MsgUnknown {
		return m, fmt.Errorf("unknown message code: %q", string(msg[:2]))
	}

	if len(msg) < minMsgLength[m.Type] {
		return m, fmt.Errorf("message too short to include required fields: %d < %d", len(msg), minMsgLength[m.Type])
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
