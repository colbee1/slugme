package slugme

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

/*
From https://github.com/Regis24GmbH/go-diacritics/blob/v2.0.2/diacritics.go

MIT License

Copyright (c) 2019 Regis24GmbH

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

var unavailableMapping = map[rune]rune{
	'\u0181': 'B',
	'\u1d81': 'd',
	'\u1d85': 'l',
	'\u1d89': 'r',
	'\u028b': 'v',
	'\u1d8d': 'x',
	'\u1d83': 'g',
	'\u0191': 'F',
	'\u0199': 'k',
	'\u019d': 'N',
	'\u0220': 'N',
	'\u01a5': 'p',
	'\u0224': 'Z',
	'\u0126': 'H',
	'\u01ad': 't',
	'\u01b5': 'Z',
	'\u0234': 'l',
	'\u023c': 'c',
	'\u0240': 'z',
	'\u0142': 'l',
	'\u0244': 'U',
	'\u2c60': 'L',
	'\u0248': 'J',
	'\ua74a': 'O',
	'\u024c': 'R',
	'\ua752': 'P',
	'\ua756': 'Q',
	'\ua75a': 'R',
	'\ua75e': 'V',
	'\u0260': 'g',
	'\u01e5': 'g',
	'\u2c64': 'R',
	'\u0166': 'T',
	'\u0268': 'i',
	'\u2c66': 't',
	'\u026c': 'l',
	'\u1d6e': 'f',
	'\u1d87': 'n',
	'\u1d72': 'r',
	'\u2c74': 'v',
	'\u1d76': 'z',
	'\u2c78': 'e',
	'\u027c': 'r',
	'\u1eff': 'y',
	'\ua741': 'k',
	'\u0182': 'B',
	'\u1d86': 'm',
	'\u0288': 't',
	'\u018a': 'D',
	'\u1d8e': 'z',
	'\u0111': 'd',
	'\u0290': 'z',
	'\u0192': 'f',
	'\u1d96': 'i',
	'\u019a': 'l',
	'\u019e': 'n',
	'\u1d88': 'p',
	'\u02a0': 'q',
	'\u01ae': 'T',
	'\u01b2': 'V',
	'\u01b6': 'z',
	'\u023b': 'C',
	'\u023f': 's',
	'\u0141': 'L',
	'\u0243': 'B',
	'\ua745': 'k',
	'\u0247': 'e',
	'\ua749': 'l',
	'\u024b': 'q',
	'\ua74d': 'o',
	'\u024f': 'y',
	'\ua751': 'p',
	'\u0253': 'b',
	'\ua755': 'p',
	'\u0257': 'd',
	'\ua759': 'q',
	'\u00d8': 'O',
	'\u2c63': 'P',
	'\u2c67': 'H',
	'\u026b': 'l',
	'\u1d6d': 'd',
	'\u1d71': 'p',
	'\u0273': 'n',
	'\u1d75': 't',
	'\u1d91': 'd',
	'\u00f8': 'o',
	'\u2c7e': 'S',
	'\u1d7d': 'p',
	'\u2c7f': 'Z',
	'\u0183': 'b',
	'\u0187': 'C',
	'\u1d80': 'b',
	'\u0289': 'u',
	'\u018b': 'D',
	'\u1d8f': 'a',
	'\u0291': 'z',
	'\u0110': 'D',
	'\u0193': 'G',
	'\u1d82': 'f',
	'\u0197': 'I',
	'\u029d': 'j',
	'\u019f': 'O',
	'\u2c6c': 'z',
	'\u01ab': 't',
	'\u01b3': 'Y',
	'\u0236': 't',
	'\u023a': 'A',
	'\u023e': 'T',
	'\ua740': 'K',
	'\u1d8a': 's',
	'\ua744': 'K',
	'\u0246': 'E',
	'\ua748': 'L',
	'\ua74c': 'O',
	'\u024e': 'Y',
	'\ua750': 'P',
	'\ua754': 'P',
	'\u0256': 'd',
	'\ua758': 'Q',
	'\u2c62': 'L',
	'\u0266': 'h',
	'\u2c73': 'w',
	'\u2c6a': 'k',
	'\u1d6c': 'b',
	'\u2c6e': 'M',
	'\u1d70': 'n',
	'\u0272': 'n',
	'\u1d92': 'e',
	'\u1d74': 's',
	'\u2c7a': 'o',
	'\u2c6b': 'Z',
	'\u027e': 'r',
	'\u0180': 'b',
	'\u0282': 's',
	'\u1d84': 'k',
	'\u0188': 'c',
	'\u018c': 'd',
	'\ua742': 'K',
	'\u1d99': 'u',
	'\u0198': 'K',
	'\u1d8c': 'v',
	'\u0221': 'd',
	'\u2c71': 'v',
	'\u0225': 'z',
	'\u01a4': 'P',
	'\u0127': 'h',
	'\u01ac': 'T',
	'\u0235': 'n',
	'\u01b4': 'y',
	'\u2c72': 'W',
	'\u023d': 'L',
	'\ua743': 'k',
	'\u0249': 'j',
	'\ua74b': 'o',
	'\u024d': 'r',
	'\ua753': 'p',
	'\u0255': 'c',
	'\ua757': 'q',
	'\u2c68': 'h',
	'\ua75b': 'r',
	'\ua75f': 'v',
	'\u2c61': 'l',
	'\u2c65': 'a',
	'\u01e4': 'G',
	'\u0167': 't',
	'\u2c69': 'K',
	'\u026d': 'l',
	'\u1d6f': 'm',
	'\u0271': 'm',
	'\u1d73': 'r',
	'\u027d': 'r',
	'\u1efe': 'Y',
}

// ToASCII removes diacritical characters and replaces them with their ASCII representation
func ToASCII(input string) string {
	var transformChain = transform.Chain(
		runes.Map(mapDecomposeUnavailable),
		norm.NFD,
		runes.Remove(runes.In(unicode.Mn)),
		norm.NFC,
	)

	input = strings.Replace(input, "\u00df", "ss", -1) // ß to ss handling
	result, _, _ := transform.String(transformChain, input)

	return result
}

func mapDecomposeUnavailable(r rune) rune {
	if v, ok := unavailableMapping[r]; ok {
		r = v
	}

	return r
}