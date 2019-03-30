package main

import "fmt"

func main() {
	s := []string{"H", "e", "l","l", "o"}
    fmt.Printf("%q\n", nonempty(s))
}
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i]=s
			i++
		}
	}
	return strings[:i]
}