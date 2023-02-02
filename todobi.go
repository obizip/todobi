// simple To-Do manager
package main

import (
	"fmt"

	"github.com/obi-3/todobi/pkg/task"
)

func main() {
	fmt.Println("Welcome to todobi!")
	t := task.NewTask("hello", 2, 3)
	fmt.Printf("%v (%v/%v) %v", t.Content(), t.Date().Month(), t.Date().Day(), t.Done())
}
