package string

func StringReverse(to_reverse_str string)string {
	r := []rune(to_reverse_str)
	for i:=0;i<len(r)/2;i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return string(r)
}