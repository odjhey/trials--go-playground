package main

import "fmt"

type Runner interface {
	run()
}

type Walker interface {
	walk()
}

func main() {
	fmt.Println("main start")

}

func move(subject interface{}) {
	fmt.Println("Move")

	s, ok := subject.(Runner)
	if !ok {
		return
	}

	s.run()  // nice, compiler knows, and since we return/exit at line 23, we super sure that `s` is a Runner
	s.walk() // nice, compiler knows

	subject.run() // <-- should not error because we checked earlier (line 21) that subject is a Runner
}

func move2(subject interface{}) {
	s, ok := subject.(Runner)
	if !ok {
		// return
	}

	s.run() // dagnabit! why this no compile error, compiler should know that we unsure of what `s` is, atleast gimme a warning or sumthing
}
