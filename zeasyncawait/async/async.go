package async

type asyncFn func() interface{}

func Call(fn asyncFn) chan interface{} {

	c := make(chan interface{})

	go func() {
		c <- fn()
	}()

	return c

}

// Qs and Todos
// 1. Make func type aware with generics
// 2. Do we have a limit on the number of channels we can create?
