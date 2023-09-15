package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mdwhatcott/slogging"
	"github.com/mdwhatcott/waiig/repl"
)

var Version = "dev"

func main() {
	slogging.SetScriptingLogger()

	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", filepath.Base(os.Args[0]), Version), flag.ExitOnError)
	_ = flags.Parse(os.Args[1:])

	fmt.Println("Hello! This is the Monkey programming language!")
	fmt.Println("Feel free to type in commands")

	repl.Start(os.Stdin, os.Stdout)
}
