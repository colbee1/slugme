package slugme

func New(opt Options) *slugme {
	s := &slugme{
		allowed:   opt.Allowed,
		lowerCase: !opt.KeepCase,
		toAscii:   !opt.KeepNonAscii,
		shrink:    !opt.NoShrink,
		trim:      !opt.NoTrim,
	}

	if r := []rune(opt.Replace); len(r) > 0 {
		s.replace = r[0]
	}

	return s
}
