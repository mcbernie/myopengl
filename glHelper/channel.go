package glHelper

//Channel For scoping to main Thread
type Channel struct {
	scope chan func()
}

var c *Channel

//InitScoping setup a new Scope
func InitScoping() {
	c = &Channel{
		scope: make(chan func()),
	}
}

//AddFunction a new function to scope
func AddFunction(fn func()) {
	c.scope <- fn
}

//RunFunctions all Function in scope
func RunFunctions() {
	select {
	case fn := <-c.scope:
		//log.Println("gets a function and run it...")
		fn()
	default:
	}
}
