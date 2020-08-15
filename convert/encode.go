package convert

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"net/url"
	"strconv"
	"strings"
)

func ToURLEncode(s string) string {
	return url.QueryEscape(s)
}

func ToURLDecode(s string) string {
	values, err := url.ParseQuery(s)
	if err != nil {
		return ""
	}

	if len(values) == 0 {
		return ""
	}

	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	return keys[0]
}

func Unicode2S(s string) string {
	bs, err := hex.DecodeString(strings.Replace(s, `\u`, ``, -1))
	if err != nil {
		return ""
	}
	var to string
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		if err = binary.Read(br, binary.BigEndian, &r); err != nil {
			return ""
		}
		to += string(r)
	}
	return to
}

func S2Unicode(s string) string {
	unicode := strconv.QuoteToASCII(s)
	return unicode[1 : len(unicode)-1]
}
