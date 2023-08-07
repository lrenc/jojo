// main.go

package main

import (
	"fmt"
	"jojo/repl"
	"os"
)

func main() {
	fmt.Println("This is the JOJO (0.0.1) programming language!")
	fmt.Println("Type \".help\" for more information.")
	repl.Start(os.Stdin, os.Stdout)
}
