package ebcdic

import (
	"testing"
	"unicode/utf8"
)

var encTable = map[byte]byte{
	'\x00': '\x00', // -> NULL
	'\x01': '\x01', // -> START OF HEADING
	'\x02': '\x02', // -> START OF TEXT
	'\x03': '\x03', // -> END OF TEXT
	'\x04': '\x9c', // -> CONTROL
	'\x05': '\t',   // -> HORIZONTAL TABULATION
	'\x06': '\x86', // -> CONTROL
	'\x07': '\x7f', // -> DELETE
	'\x08': '\x97', // -> CONTROL
	'\x09': '\x8d', // -> CONTROL
	'\x0A': '\x8e', // -> CONTROL
	'\x0B': '\x0b', // -> VERTICAL TABULATION
	'\x0C': '\x0c', // -> FORM FEED
	'\x0D': '\r',   // -> CARRIAGE RETURN
	'\x0E': '\x0e', // -> SHIFT OUT
	'\x0F': '\x0f', // -> SHIFT IN
	'\x10': '\x10', // -> DATA LINK ESCAPE
	'\x11': '\x11', // -> DEVICE CONTROL ONE
	'\x12': '\x12', // -> DEVICE CONTROL TWO
	'\x13': '\x13', // -> DEVICE CONTROL THREE
	'\x14': '\x9d', // -> CONTROL
	'\x15': '\x85', // -> CONTROL
	'\x16': '\x08', // -> BACKSPACE
	'\x17': '\x87', // -> CONTROL
	'\x18': '\x18', // -> CANCEL
	'\x19': '\x19', // -> END OF MEDIUM
	'\x1A': '\x92', // -> CONTROL
	'\x1B': '\x8f', // -> CONTROL
	'\x1C': '\x1c', // -> FILE SEPARATOR
	'\x1D': '\x1d', // -> GROUP SEPARATOR
	'\x1E': '\x1e', // -> RECORD SEPARATOR
	'\x1F': '\x1f', // -> UNIT SEPARATOR
	'\x20': '\x80', // -> CONTROL
	'\x21': '\x81', // -> CONTROL
	'\x22': '\x82', // -> CONTROL
	'\x23': '\x83', // -> CONTROL
	'\x24': '\x84', // -> CONTROL
	'\x25': '\n',   // -> LINE FEED
	'\x26': '\x17', // -> END OF TRANSMISSION BLOCK
	'\x27': '\x1b', // -> ESCAPE
	'\x28': '\x88', // -> CONTROL
	'\x29': '\x89', // -> CONTROL
	'\x2A': '\x8a', // -> CONTROL
	'\x2B': '\x8b', // -> CONTROL
	'\x2C': '\x8c', // -> CONTROL
	'\x2D': '\x05', // -> ENQUIRY
	'\x2E': '\x06', // -> ACKNOWLEDGE
	'\x2F': '\x07', // -> BELL
	'\x30': '\x90', // -> CONTROL
	'\x31': '\x91', // -> CONTROL
	'\x32': '\x16', // -> SYNCHRONOUS IDLE
	'\x33': '\x93', // -> CONTROL
	'\x34': '\x94', // -> CONTROL
	'\x35': '\x95', // -> CONTROL
	'\x36': '\x96', // -> CONTROL
	'\x37': '\x04', // -> END OF TRANSMISSION
	'\x38': '\x98', // -> CONTROL
	'\x39': '\x99', // -> CONTROL
	'\x3A': '\x9a', // -> CONTROL
	'\x3B': '\x9b', // -> CONTROL
	'\x3C': '\x14', // -> DEVICE CONTROL FOUR
	'\x3D': '\x15', // -> NEGATIVE ACKNOWLEDGE
	'\x3E': '\x9e', // -> CONTROL
	'\x3F': '\x1a', // -> SUBSTITUTE
	'\x40': ' ',    // -> SPACE
	'\x41': '\xa0', // -> NO-BREAK SPACE
	'\x42': '\xe2', // -> LATIN SMALL LETTER A WITH CIRCUMFLEX
	'\x43': '\xe4', // -> LATIN SMALL LETTER A WITH DIAERESIS
	'\x44': '\xe0', // -> LATIN SMALL LETTER A WITH GRAVE
	'\x45': '\xe1', // -> LATIN SMALL LETTER A WITH ACUTE
	'\x46': '\xe3', // -> LATIN SMALL LETTER A WITH TILDE
	'\x47': '\xe5', // -> LATIN SMALL LETTER A WITH RING ABOVE
	'\x48': '\xe7', // -> LATIN SMALL LETTER C WITH CEDILLA
	'\x49': '\xf1', // -> LATIN SMALL LETTER N WITH TILDE
	'\x4A': '[',    // -> LEFT SQUARE BRACKET
	'\x4B': '.',    // -> FULL STOP
	'\x4C': '<',    // -> LESS-THAN SIGN
	'\x4D': '(',    // -> LEFT PARENTHESIS
	'\x4E': '+',    // -> PLUS SIGN
	'\x4F': '!',    // -> EXCLAMATION MARK
	'\x50': '&',    // -> AMPERSAND
	'\x51': '\xe9', // -> LATIN SMALL LETTER E WITH ACUTE
	'\x52': '\xea', // -> LATIN SMALL LETTER E WITH CIRCUMFLEX
	'\x53': '\xeb', // -> LATIN SMALL LETTER E WITH DIAERESIS
	'\x54': '\xe8', // -> LATIN SMALL LETTER E WITH GRAVE
	'\x55': '\xed', // -> LATIN SMALL LETTER I WITH ACUTE
	'\x56': '\xee', // -> LATIN SMALL LETTER I WITH CIRCUMFLEX
	'\x57': '\xef', // -> LATIN SMALL LETTER I WITH DIAERESIS
	'\x58': '\xec', // -> LATIN SMALL LETTER I WITH GRAVE
	'\x59': '\xdf', // -> LATIN SMALL LETTER SHARP S (GERMAN)
	'\x5A': ']',    // -> RIGHT SQUARE BRACKET
	'\x5B': '$',    // -> DOLLAR SIGN
	'\x5C': '*',    // -> ASTERISK
	'\x5D': ')',    // -> RIGHT PARENTHESIS
	'\x5E': ';',    // -> SEMICOLON
	'\x5F': '^',    // -> CIRCUMFLEX ACCENT
	'\x60': '-',    // -> HYPHEN-MINUS
	'\x61': '/',    // -> SOLIDUS
	'\x62': '\xc2', // -> LATIN CAPITAL LETTER A WITH CIRCUMFLEX
	'\x63': '\xc4', // -> LATIN CAPITAL LETTER A WITH DIAERESIS
	'\x64': '\xc0', // -> LATIN CAPITAL LETTER A WITH GRAVE
	'\x65': '\xc1', // -> LATIN CAPITAL LETTER A WITH ACUTE
	'\x66': '\xc3', // -> LATIN CAPITAL LETTER A WITH TILDE
	'\x67': '\xc5', // -> LATIN CAPITAL LETTER A WITH RING ABOVE
	'\x68': '\xc7', // -> LATIN CAPITAL LETTER C WITH CEDILLA
	'\x69': '\xd1', // -> LATIN CAPITAL LETTER N WITH TILDE
	'\x6A': '\xa6', // -> BROKEN BAR
	'\x6B': ',',    // -> COMMA
	'\x6C': '%',    // -> PERCENT SIGN
	'\x6D': '_',    // -> LOW LINE
	'\x6E': '>',    // -> GREATER-THAN SIGN
	'\x6F': '?',    // -> QUESTION MARK
	'\x70': '\xf8', // -> LATIN SMALL LETTER O WITH STROKE
	'\x71': '\xc9', // -> LATIN CAPITAL LETTER E WITH ACUTE
	'\x72': '\xca', // -> LATIN CAPITAL LETTER E WITH CIRCUMFLEX
	'\x73': '\xcb', // -> LATIN CAPITAL LETTER E WITH DIAERESIS
	'\x74': '\xc8', // -> LATIN CAPITAL LETTER E WITH GRAVE
	'\x75': '\xcd', // -> LATIN CAPITAL LETTER I WITH ACUTE
	'\x76': '\xce', // -> LATIN CAPITAL LETTER I WITH CIRCUMFLEX
	'\x77': '\xcf', // -> LATIN CAPITAL LETTER I WITH DIAERESIS
	'\x78': '\xcc', // -> LATIN CAPITAL LETTER I WITH GRAVE
	'\x79': '`',    // -> GRAVE ACCENT
	'\x7A': ':',    // -> COLON
	'\x7B': '#',    // -> NUMBER SIGN
	'\x7C': '@',    // -> COMMERCIAL AT
	'\x7D': '\'',   // -> APOSTROPHE
	'\x7E': '=',    // -> EQUALS SIGN
	'\x7F': '"',    // -> QUOTATION MARK
	'\x80': '\xd8', // -> LATIN CAPITAL LETTER O WITH STROKE
	'\x81': 'a',    // -> LATIN SMALL LETTER A
	'\x82': 'b',    // -> LATIN SMALL LETTER B
	'\x83': 'c',    // -> LATIN SMALL LETTER C
	'\x84': 'd',    // -> LATIN SMALL LETTER D
	'\x85': 'e',    // -> LATIN SMALL LETTER E
	'\x86': 'f',    // -> LATIN SMALL LETTER F
	'\x87': 'g',    // -> LATIN SMALL LETTER G
	'\x88': 'h',    // -> LATIN SMALL LETTER H
	'\x89': 'i',    // -> LATIN SMALL LETTER I
	'\x8A': '\xab', // -> LEFT-POINTING DOUBLE ANGLE QUOTATION MARK
	'\x8B': '\xbb', // -> RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
	'\x8C': '\xf0', // -> LATIN SMALL LETTER ETH (ICELANDIC)
	'\x8D': '\xfd', // -> LATIN SMALL LETTER Y WITH ACUTE
	'\x8E': '\xfe', // -> LATIN SMALL LETTER THORN (ICELANDIC)
	'\x8F': '\xb1', // -> PLUS-MINUS SIGN
	'\x90': '\xb0', // -> DEGREE SIGN
	'\x91': 'j',    // -> LATIN SMALL LETTER J
	'\x92': 'k',    // -> LATIN SMALL LETTER K
	'\x93': 'l',    // -> LATIN SMALL LETTER L
	'\x94': 'm',    // -> LATIN SMALL LETTER M
	'\x95': 'n',    // -> LATIN SMALL LETTER N
	'\x96': 'o',    // -> LATIN SMALL LETTER O
	'\x97': 'p',    // -> LATIN SMALL LETTER P
	'\x98': 'q',    // -> LATIN SMALL LETTER Q
	'\x99': 'r',    // -> LATIN SMALL LETTER R
	'\x9A': '\xaa', // -> FEMININE ORDINAL INDICATOR
	'\x9B': '\xba', // -> MASCULINE ORDINAL INDICATOR
	'\x9C': '\xe6', // -> LATIN SMALL LIGATURE AE
	'\x9D': '\xb8', // -> CEDILLA
	'\x9E': '\xc6', // -> LATIN CAPITAL LIGATURE AE
	'\x9F': '\xa4', // -> CURRENCY SIGN
	'\xA0': '\xb5', // -> MICRO SIGN
	'\xA1': '~',    // -> TILDE
	'\xA2': 's',    // -> LATIN SMALL LETTER S
	'\xA3': 't',    // -> LATIN SMALL LETTER T
	'\xA4': 'u',    // -> LATIN SMALL LETTER U
	'\xA5': 'v',    // -> LATIN SMALL LETTER V
	'\xA6': 'w',    // -> LATIN SMALL LETTER W
	'\xA7': 'x',    // -> LATIN SMALL LETTER X
	'\xA8': 'y',    // -> LATIN SMALL LETTER Y
	'\xA9': 'z',    // -> LATIN SMALL LETTER Z
	'\xAA': '\xa1', // -> INVERTED EXCLAMATION MARK
	'\xAB': '\xbf', // -> INVERTED QUESTION MARK
	'\xAC': '\xd0', // -> LATIN CAPITAL LETTER ETH (ICELANDIC)
	'\xAD': '\xdd', // -> LATIN CAPITAL LETTER Y WITH ACUTE
	'\xAE': '\xde', // -> LATIN CAPITAL LETTER THORN (ICELANDIC)
	'\xAF': '\xae', // -> REGISTERED SIGN
	'\xB0': '\xa2', // -> CENT SIGN
	'\xB1': '\xa3', // -> POUND SIGN
	'\xB2': '\xa5', // -> YEN SIGN
	'\xB3': '\xb7', // -> MIDDLE DOT
	'\xB4': '\xa9', // -> COPYRIGHT SIGN
	'\xB5': '\xa7', // -> SECTION SIGN
	'\xB6': '\xb6', // -> PILCROW SIGN
	'\xB7': '\xbc', // -> VULGAR FRACTION ONE QUARTER
	'\xB8': '\xbd', // -> VULGAR FRACTION ONE HALF
	'\xB9': '\xbe', // -> VULGAR FRACTION THREE QUARTERS
	'\xBA': '\xac', // -> NOT SIGN
	'\xBB': '|',    // -> VERTICAL LINE
	'\xBC': '\xaf', // -> MACRON
	'\xBD': '\xa8', // -> DIAERESIS
	'\xBE': '\xb4', // -> ACUTE ACCENT
	'\xBF': '\xd7', // -> MULTIPLICATION SIGN
	'\xC0': '{',    // -> LEFT CURLY BRACKET
	'\xC1': 'A',    // -> LATIN CAPITAL LETTER A
	'\xC2': 'B',    // -> LATIN CAPITAL LETTER B
	'\xC3': 'C',    // -> LATIN CAPITAL LETTER C
	'\xC4': 'D',    // -> LATIN CAPITAL LETTER D
	'\xC5': 'E',    // -> LATIN CAPITAL LETTER E
	'\xC6': 'F',    // -> LATIN CAPITAL LETTER F
	'\xC7': 'G',    // -> LATIN CAPITAL LETTER G
	'\xC8': 'H',    // -> LATIN CAPITAL LETTER H
	'\xC9': 'I',    // -> LATIN CAPITAL LETTER I
	'\xCA': '\xad', // -> SOFT HYPHEN
	'\xCB': '\xf4', // -> LATIN SMALL LETTER O WITH CIRCUMFLEX
	'\xCC': '\xf6', // -> LATIN SMALL LETTER O WITH DIAERESIS
	'\xCD': '\xf2', // -> LATIN SMALL LETTER O WITH GRAVE
	'\xCE': '\xf3', // -> LATIN SMALL LETTER O WITH ACUTE
	'\xCF': '\xf5', // -> LATIN SMALL LETTER O WITH TILDE
	'\xD0': '}',    // -> RIGHT CURLY BRACKET
	'\xD1': 'J',    // -> LATIN CAPITAL LETTER J
	'\xD2': 'K',    // -> LATIN CAPITAL LETTER K
	'\xD3': 'L',    // -> LATIN CAPITAL LETTER L
	'\xD4': 'M',    // -> LATIN CAPITAL LETTER M
	'\xD5': 'N',    // -> LATIN CAPITAL LETTER N
	'\xD6': 'O',    // -> LATIN CAPITAL LETTER O
	'\xD7': 'P',    // -> LATIN CAPITAL LETTER P
	'\xD8': 'Q',    // -> LATIN CAPITAL LETTER Q
	'\xD9': 'R',    // -> LATIN CAPITAL LETTER R
	'\xDA': '\xb9', // -> SUPERSCRIPT ONE
	'\xDB': '\xfb', // -> LATIN SMALL LETTER U WITH CIRCUMFLEX
	'\xDC': '\xfc', // -> LATIN SMALL LETTER U WITH DIAERESIS
	'\xDD': '\xf9', // -> LATIN SMALL LETTER U WITH GRAVE
	'\xDE': '\xfa', // -> LATIN SMALL LETTER U WITH ACUTE
	'\xDF': '\xff', // -> LATIN SMALL LETTER Y WITH DIAERESIS
	'\xE0': '\\',   // -> REVERSE SOLIDUS
	'\xE1': '\xf7', // -> DIVISION SIGN
	'\xE2': 'S',    // -> LATIN CAPITAL LETTER S
	'\xE3': 'T',    // -> LATIN CAPITAL LETTER T
	'\xE4': 'U',    // -> LATIN CAPITAL LETTER U
	'\xE5': 'V',    // -> LATIN CAPITAL LETTER V
	'\xE6': 'W',    // -> LATIN CAPITAL LETTER W
	'\xE7': 'X',    // -> LATIN CAPITAL LETTER X
	'\xE8': 'Y',    // -> LATIN CAPITAL LETTER Y
	'\xE9': 'Z',    // -> LATIN CAPITAL LETTER Z
	'\xEA': '\xb2', // -> SUPERSCRIPT TWO
	'\xEB': '\xd4', // -> LATIN CAPITAL LETTER O WITH CIRCUMFLEX
	'\xEC': '\xd6', // -> LATIN CAPITAL LETTER O WITH DIAERESIS
	'\xED': '\xd2', // -> LATIN CAPITAL LETTER O WITH GRAVE
	'\xEE': '\xd3', // -> LATIN CAPITAL LETTER O WITH ACUTE
	'\xEF': '\xd5', // -> LATIN CAPITAL LETTER O WITH TILDE
	'\xF0': '0',    // -> DIGIT ZERO
	'\xF1': '1',    // -> DIGIT ONE
	'\xF2': '2',    // -> DIGIT TWO
	'\xF3': '3',    // -> DIGIT THREE
	'\xF4': '4',    // -> DIGIT FOUR
	'\xF5': '5',    // -> DIGIT FIVE
	'\xF6': '6',    // -> DIGIT SIX
	'\xF7': '7',    // -> DIGIT SEVEN
	'\xF8': '8',    // -> DIGIT EIGHT
	'\xF9': '9',    // -> DIGIT NINE
	'\xFA': '\xb3', // -> SUPERSCRIPT THREE
	'\xFB': '\xdb', // -> LATIN CAPITAL LETTER U WITH CIRCUMFLEX
	'\xFC': '\xdc', // -> LATIN CAPITAL LETTER U WITH DIAERESIS
	'\xFD': '\xd9', // -> LATIN CAPITAL LETTER U WITH GRAVE
	'\xFE': '\xda', // -> LATIN CAPITAL LETTER U WITH ACUTE
	'\xFF': '\x9f', // -> CONTROL
}

func TestDecodeRune(t *testing.T) {

	type testScheme struct {
		p     []byte
		want  rune
		width int
	}

	var tests []testScheme

	for ebc, asc := range encTable {
		tests = append(tests, testScheme{[]byte{ebc}, rune(asc), 1})
	}

	for _, test := range tests {
		if r, w := DecodeRune(test.p); r != test.want || w != test.width {
			t.Errorf("decode(%v), want: %c, got: %c ", test.p, test.want, r)
		}
	}
}

func TestRuneError(t *testing.T) {
	if RuneError != utf8.RuneError {
		t.Errorf("RuneError mismatched. %v", RuneError)
	}
}
