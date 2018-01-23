package main

import (
	_ "fmt"
	_ "container/list"
	"fmt"
	"core"
)

func main() {
	task1 := core.Task{}
	task2 := core.Task{}
	task2.Name = "喝牛奶"
	task1.FatherTasks.PushBack(task2)
	//task2.ChildTasks.PushBack(task1)
	for e := task1.FatherTasks.Front(); e != nil; e.Next() {
		fmt.Println(e.Value)
	}
	//fmt.Print(task1.FatherTasks)

}
