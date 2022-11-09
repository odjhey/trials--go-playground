package pkga

import "fmt"

var Prefix = "Echo"

func Echo(text string) {

	fmt.Println(Prefix + ": " + text)

}
