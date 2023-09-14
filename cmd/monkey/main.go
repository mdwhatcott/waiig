package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mdwhatcott/exec"
	"github.com/mdwhatcott/must/must"
	"github.com/mdwhatcott/slogging"
	"github.com/mdwhatcott/tui/v2"
)

var Version = "dev"

func main() {
	slogging.SetScriptingLogger()

	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", filepath.Base(os.Args[0]), Version), flag.ExitOnError)
	_ = flags.Parse(os.Args[1:])

	ui := tui.New()
	ui.Println(must.Value(exec.Run(fmt.Sprintf("echo 'Hello, %s'", flags.Name()))))
}
