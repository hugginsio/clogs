package main

import (
	"errors"

	"github.com/hugginsio/clogs"
)

func main() {
	clogs.Debugln("this is an example log line at debug level")
	clogs.Println("this is an example log line with info")
	clogs.Infoln("this is another log line with info", "and multiple arguments")
	clogs.Warnln("this is a warning line")
	clogs.Errorln("this is an error line", errors.New("with the error passed as an argument"))
}
