package helper

//Channel For scoping to main Thread
type Channel struct {
	scope chan func()
}

var c *Channel

// InitScoping
// Simple create a channel to call function only inside main thread
func InitScoping() {
	c = &Channel{
		scope: make(chan func()),
	}
}

//AddFunction a new function to scope
//for calling a gl function in main thread
func AddFunction(fn func()) {
	c.scope <- fn
}

//RunFunctions all Function in scope
//for run all function in main thread
func RunFunctions() {
	select {
	case fn := <-c.scope:
		//log.Println("gets a function and run it...")
		fn()
	default:
	}
}
