package main

import (
	"fmt"

	"github.com/oarkflow/tail"
)

func main() {
	done := make(chan bool, 1)
	tail.Dir("/Users/sujit/Sites/frame/examples/logs", func(file, line string) error {
		fmt.Println(file)
		fmt.Println(line)
		return nil
	})
	<-done
}
