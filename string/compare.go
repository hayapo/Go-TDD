package string

//compare two strings lexicographically and return integer
func Compare(a, b string) (compared int) {
	switch {
	case a < b:
		compared = -1
	case a > b:
		compared = +1
	case a == b:
		compared = 0
	}
	return
}