package pkga

import "fmt"

var Prefix = "Echo"
var Postfix = ""

func Echo(text string) {

	fmt.Println(Prefix + ": " + text)

}
