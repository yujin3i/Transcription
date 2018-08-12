package echo

func Echo1(strings []string) string {
	var s, sep string
	for i := 1; i < len(strings); i++ {
		s += sep + strings[i]
		sep = " "
	}
	return s
}

func Echo2(strings []string) string {
	s, sep := "", ""
	for _, str := range strings {
		s += sep + str
		sep = " "
	}
	return s
}
