package async

// NOTE! This is just an example to to grok how concurrency works in go
//       Basing from the assumption that we have existing knowledge of the async/await concepts

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
// 3. Is creating channels cheap? can we do use it like a regular var
