package control

import (
	"kboard/core"
	"fmt"
	"log"
)

type IControl interface {
	Register(string, func()) *Control
	Run(string)
}

type Control struct {
	Config *core.Config
	TplEngine *core.TplEngine
	Control string
	Actions map[string]func()
}


func (c *Control) Register(action string, f func()) *Control {
	if c.Actions == nil {
		c.Actions = map[string]func(){}
	}
	if c.Control == "" {
		core.CheckError(core.NewError("error: control is empty!"), 999)
	}
	c.Actions[action] = f
	return c
}

func (c *Control) Index() {
	fmt.Fprintln(c.TplEngine.W, "hello world, this is default index")
}


func (c *Control) Run(action string) {
	if c.Actions[action] == nil {
		//
		fmt.Fprintln(c.TplEngine.W, "404 page not found!")
		log.Println("404")
	}else{
		// run action
		c.Actions[action]()
	}
}