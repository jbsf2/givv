package givv

import "fmt"

type any interface{}

func givvPanic(messageFormat string, args... any) {
	message := fmt.Sprintf(messageFormat, args)
	panic(message)
}