package string

func Join(slice []string, sep string) string{
	var joined string
	for i, r := range slice {
		if i != len(slice)-1 {
			joined += r + sep
		} else {
			joined += r
		}
	}
	return joined
}
