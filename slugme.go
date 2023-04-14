package slugme

import (
	"strings"
	"unicode"
)

var DefaultOptions = Options{
	Allowed: "-_",
	Replace: "-",
}

type Options struct {
	Allowed      string // Allowed symbols in slug in addition to letters (a-Z) and numbers (0-9).
	Replace      string // Replace all disallowed symbols by this one.
	KeepCase     bool   // Do not lower case the slug
	KeepNonAscii bool   // Do not try to convert letters to ASCII7, aka: do not remove diacritics.
	NoShrink     bool   // Do not shrink repetition of Replace char.
	NoTrim       bool   // Do not trim slug start/end for Replace character.
}

type slugme struct {
	allowed   string
	replace   rune
	lowerCase bool
	toAscii   bool
	shrink    bool
	trim      bool
}

func (s *slugme) Slug(text string) string {
	if text == "" {
		return text
	}

	if s.toAscii {
		text = ToASCII(text)
	}

	containsAllowed := false
	slug := make([]rune, 0, len(text))
	for _, c := range text {
		if unicode.In(c, unicode.Letter, unicode.Digit) || strings.ContainsRune(s.allowed, c) {
			slug = append(slug, c)
			containsAllowed = true
		} else if s.replace != rune(0) {
			slug = append(slug, s.replace)
		}
	}
	if !containsAllowed {
		return ""
	}

	// Shrink
	//
	if s.shrink {
		pos, repeat := 0, false
		for i, c := range slug {
			if c == s.replace {
				if !repeat {
					slug[pos] = slug[i]
					repeat = true
					pos++
				}

				continue
			}

			slug[pos] = slug[i]
			repeat = false
			pos++
		}
		slug = slug[0:pos]
	}

	// Trim
	//
	if s.trim {
		start, stop := 0, 0
		for i, c := range slug {
			if c == s.replace {
				continue
			}
			start = i
			break
		}

		for i := len(slug) - 1; i >= 0; i-- {
			if slug[i] == s.replace {
				continue
			}
			stop = i + 1
			break
		}
		slug = slug[start:stop]
	}

	// lower case
	//
	if s.lowerCase {
		return strings.ToLower(string(slug))
	}

	return string(slug)
}
