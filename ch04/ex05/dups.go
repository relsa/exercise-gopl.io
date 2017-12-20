package main

func rmAdjacentDups(strings []string) []string {

	if len(strings) == 0 {
		return strings
	}

	prev := strings[0]
	ptr := 1

	for _, s := range strings {
		if s != prev {
			prev = s
			strings[ptr] = s
			ptr++
		}
	}

	return strings[:ptr]
}
