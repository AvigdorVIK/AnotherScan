package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)

	scaner.Scan()
	par := scaner.Text()

	fmt.Println(par)

}
