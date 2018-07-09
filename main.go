package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/isaacasensio/memo/cmd"
	"github.com/sirupsen/logrus"
)

const banner = `
 __  __                      
|  \/  |                     
| \  / | ___ _ __ ___   ___  
| |\/| |/ _ \ '_ ' _ \ / _ \ 
| |  | |  __/ | | | | | (_) |
|_|  |_|\___|_| |_| |_|\___/ 
`

func main() {
	fmt.Println(banner)
	var pwdr cmd.StdInPasswordReader
	scanner := bufio.NewScanner(os.Stdin)
	result, err := cmd.Run(pwdr, scanner)
	if err != nil {
		logrus.Fatalf("something went wrong: %v", err)
		os.Exit(1)
	}

	fmt.Printf("\nResult: %s\n", result)
}
