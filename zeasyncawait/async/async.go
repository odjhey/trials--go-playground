package async

type asyncFn func() interface{}

func Call(fn asyncFn) chan interface{} {

	c := make(chan interface{})

	go func() {
		c <- fn()
	}()

	return c

}
