package greeting

import "fmt"

func Greet(name string) string {
	message:=fmt.Sprintf("hi，%v. Welcome!",name)
	return message
}
