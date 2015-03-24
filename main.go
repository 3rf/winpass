package main

import "bufio"
import "fmt"
import "github.com/howeyc/gopass"
import "golang.org/x/crypto/ssh/terminal"
import "os"
import "syscall"

func main() {
	if !terminal.IsTerminal(syscall.STD_INPUT_HANDLE) {
		fmt.Println("GOT A PIPE")
		in, _, err := bufio.NewReader(os.Stdin).ReadLine()
		if err != nil {
			panic(err)
		}
		fmt.Println()
		fmt.Println("GOT:", string(in))
	} else {
		fmt.Printf("Password: ")
		pass := gopass.GetPasswd()
		fmt.Println("GOT:", string(pass))
	}
}
