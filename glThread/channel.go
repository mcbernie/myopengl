package glThread

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

//Add a new function to scope
func Add(fn func()) {
	c.scope <- fn
}

//Runs all Function in scope
func Runs() {
	select {
	case fn := <-c.scope:
		//log.Println("gets a function and run it...")
		fn()
	default:
	}
}
