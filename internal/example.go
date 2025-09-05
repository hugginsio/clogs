package main

import (
	"errors"

	"github.com/hugginsio/clogs"
)

func main() {
	clogs.Debugln("This log line won't appear. DEBUG is off by default.")
	clogs.SetDebugMode(true)
	clogs.Debugln("You should be able to see this now that DEBUG is enabled.")

	clogs.Println("PrintLn is an alias for INFO.")
	clogs.Infoln("This method prints INFO.")
	clogs.Warnln("This method prints WARN.")
	clogs.Errorln("This method prints ERROR. All methods support multiple arguments,", errors.New("like this!"))
}
