package main

import (
	"fmt"
	"strings"
)

func pingy_pongy(ping <-chan string, pong chan<- string) {
	// ping is a channel to read from only
	// pong is a channel to write to it only
	// => this routine is running in the background forever
	for {
		// receive the data from the ping channel sent via another routine
		data := <-ping

		// sned the response to the routine who listens to pong channel
		pong <- strings.ToUpper(data)
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go pingy_pongy(ping, pong)

	fmt.Println("Type your msg and hit enter to send, or hit q to quit")

	for {
		fmt.Print("➜ ")

		// get the user inp
		var userInp string
		_, _ = fmt.Scanln(&userInp)

		if "q" == strings.ToLower(userInp) {
			break
		}

		ping <- strings.ToUpper(userInp)

		// wait to get the response, so the main rotuine will block and won't go the next iteration to receive another input untill the pong is receiving a msg
		response := <-pong
		fmt.Printf("RECEIVED : %v \n", response)
	}

	close(pong)
	close(ping)
}
