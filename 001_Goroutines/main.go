package main

import (
	"fmt"
	"sync"
)

func DoTask(s string, wg *sync.WaitGroup) {
	fmt.Println("Task ", s, " is finished!")
	// decrement the number of waited tasks that the waitgroup is waiting for
	wg.Done()
}
func main() {
	tasks := []string{"task1", "task2", "task3", "task4", "task5", "task6"}

	wg := sync.WaitGroup{}

	wg.Add(len(tasks))

	for _, task := range tasks {
		go DoTask(task, &wg)
	}
	// let the main routine wait for the all registered tasks in the wg wait group
	wg.Wait() // <-- the main routine will wait and get blocked here untill all registered go-routines are finished

	fmt.Println("all tasks are done!")
}
