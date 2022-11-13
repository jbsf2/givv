package givv

import "fmt"

func givvPanic(messageFormat string, args... any) {
	message := fmt.Sprintf(messageFormat, args)
	panic(message)
}