package chunk

func Placeholders(i int) string {
	s := ""
	for ; i > 0; i-- {
		if i > 1 {
			s += "?, "
			continue
		}
		s += "?"
	}
	return s
}
