package main

import "bufio"
import "fmt"
import "golang.org/x/crypto/ssh/terminal"
import "os"

func main() {
	if terminal.IsTerminal(0) {
		bup, err := terminal.MakeRaw(0)
		if err != nil {
			panic(err)
		}

		in, _, err := bufio.NewReader(os.Stdin).ReadLine()
		if err != nil {
			panic(err)
		}

		err = terminal.Restore(0, bup)
		if err != nil {
			panic(err)
		}
		fmt.Println()
		fmt.Println("GOT:", in)
	} else {
		fmt.Println("GOT A PIPE")
		pw, err := terminal.ReadPassword(0)
		if err != nil {
			panic(err)
		}
		fmt.Println("GOT:", pw)
	}
}
