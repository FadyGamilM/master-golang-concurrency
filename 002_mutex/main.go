package main

import (
	"fmt"
	"sync"
)

// Here is the problem
type User struct {
	id       int
	username string
}

func UpdateUserData(user *User, username string, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()

	mx.Lock()
	user.username = username
	mx.Unlock()
}

// with mutex
func main() {
	user := User{
		id:       1,
		username: "fady",
	}

	// tell the main routine to wait for 3 concurrent go routines to finish and then continue
	wg := sync.WaitGroup{}
	wg.Add(3)

	mx := sync.Mutex{}

	// spin the go routines to edit the user concurrently
	go UpdateUserData(&user, "updated fady #1 !! ", &wg, &mx)
	go UpdateUserData(&user, "updated fady #2 !! ", &wg, &mx)
	go UpdateUserData(&user, "updated fady #3 !! ", &wg, &mx)

	// block the main routine
	wg.Wait()

	fmt.Println("user after update is : ", user.username)

}

// without  mutex
// func main() {
// 	user := User{
// 		id:       1,
// 		username: "fady",
// 	}

// 	// tell the main routine to wait for 3 concurrent go routines to finish and then continue
// 	wg := sync.WaitGroup{}
// 	wg.Add(3)

// 	// spin the go routines to edit the user concurrently
// 	go UpdateUserData(&user, "updated fady #1 !! ", &wg)
// 	go UpdateUserData(&user, "updated fady #2 !! ", &wg)
// 	go UpdateUserData(&user, "updated fady #3 !! ", &wg)

// 	// block the main routine
// 	wg.Wait()

// 	fmt.Println("user after update is : ", user.username)

// }
