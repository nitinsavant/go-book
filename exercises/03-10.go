package exercises

import "bytes"

func Comma(number string) string {
	length := len(number)
	if length <= 3 {
		return number
	}

	var buf bytes.Buffer
	var i int
	var commaTracker int

	for i = len(number) - 1; i >= 0; i-- {
		buf.WriteByte(number[i])
		commaTracker++
		if commaTracker > 0 && i > 0 && commaTracker%3 == 0 {
			buf.WriteByte(',')
		}
	}

	return reverse(buf.String())
}

func reverse(s string) string {
	var buf bytes.Buffer
	var i int

	for i = len(s) - 1; i >= 0; i-- {
		buf.WriteByte(s[i])
	}

	return buf.String()
}
