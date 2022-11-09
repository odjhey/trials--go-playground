package main

import (
	"fmt"
	"gotraining/pkg/pkga"
	"log"
	"os"
	"strconv"
)

// please accept marks also from the command line change the if conditions and program accordingly
func greet() {

	// os.Args is a string type so no matter what we pass that would be always be of string type so conversion is imp
	details := os.Args[1:]

	if len(details) < 2 {
		log.Println("please provide your name, age, marks")
		//log.Fatalln("") // stop the program
		//os.Exit(1)
		return // stops the exec of the current func  // note :- it doesn't stop the program
	}
	fmt.Println(details)

	name := details[0]
	ageString := details[1]
	markString := details[2]

	age, err := strconv.Atoi(ageString)
	if err != nil {
		log.Println(err)
		return
	}

	mark, err := strconv.Atoi(markString)
	if err != nil {
		log.Println(err)
		return
	}

	pkga.Echo("Hey")
	pkga.Prefix = "John Cena"
	pkga.Echo("Hello")
	fmt.Println(name, age, mark)

	// default val of err is nil which indicates no error
	//var defaultErr error
	//fmt.Println("defaultErr", defaultErr)

}

func main() {

	greet()

	fmt.Println("main func ends")

}
