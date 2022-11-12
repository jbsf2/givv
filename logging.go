package givv

import "fmt"

func Debug(message string, args ...any) {
	fmt.Printf(message, args)
	fmt.Println("")
}