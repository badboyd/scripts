package trans

import (
	// "errors"
	"fmt"
	"log"

	"git.chotot.org/go-common/trans"
)

// Command ...
type Command struct {
	trans.TransCmd
}

func Init(addr string, port int) {
	// trans.SetTransServer(addr, port)
}

// SendCommand ...
func (c *Command) SendCommand(cmd string) (token string, err error) {
	c.Cmd = cmd
	c.Commit = true
	return c.executeCommand()
}

func printErrorLog(err error) {
	if err != nil {
		log.Println(err)
	}
}
func (c *Command) executeCommand() (token string, err error) {
	err = c.Execute()
	token = c.GetResValue("token")
	defer printErrorLog(err)

	if err != nil {
		return
	}

	if !c.CheckStatusOK() {
		err = fmt.Errorf("error=%s,status=%s,token=%s",
			c.GetResValue("error"),
			c.GetResValue("status"),
			token)
		return
	}

	return
}
