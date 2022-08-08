package exercises

func RemoveDupes(strings []string) []string {
	prev := ""
	i := 0
	for _, val := range strings {
		if val != prev {
			strings[i] = val
			i++
		}
		prev = val
	}
	return strings[:i]
}
