// Package likes contains functions for generating messages based on a post's likes
package likes

import "fmt"

// Likes takes a slice of people who like a post and returns a string indicating who likes it
func Likes(list []string) string {
	likeLen := len(list)

	switch likeLen {
	case 0:
		return "no one likes this"
	case 1:
		return fmt.Sprintf("%s likes this", list[0])
	case 2:
		return fmt.Sprintf("%s and %s like this", list[0], list[1])
	case 3:
		return fmt.Sprintf("%s, %s, and %s like this", list[0], list[1], list[2])
	default:
		return fmt.Sprintf("%s, %s, and %d others like this", list[0], list[1], likeLen-2)
	}
}
