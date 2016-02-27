// It attempts to mimic bufio/scan.go
package ebcdic

func isSpace(r rune) bool {

	if r <= '\u00ff' {
		switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}

	//high value ones.
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}

	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}

// ScanWords is a split function for a Scanner that returns each space-seperated
// word of text, with surrounding spaces deleted. It will never return an empty
// string. The definition of space is set by ebcidic.IsSpace.
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {

	// skipping all front-spaces
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = DecodeRune(data[start:])
		if !isSpace(r) {
			break
		}
	}

	// scan until newline is met
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = DecodeRune(data[i:])
		if isSpace(r) {
			return i + width, data[start:i], nil
		}
	}

	// if we are at EOF
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}

	return start, nil, nil
}
