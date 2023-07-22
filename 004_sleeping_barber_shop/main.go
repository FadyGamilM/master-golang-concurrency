package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// global package level vars
var seat_capacity = 10                             // 10 seats in the shop
var arrival_rate = 100                             // we have a customer every 100
var time_to_get_cut_done = 1000 * time.Millisecond // time for the barber to cut the custoemr hair
var time_for_shop_to_be_open = 10 * time.Second    // the shop is open for 10 seconds each day

func main() {
	// rand used for arrival_rate
	rand.Seed(time.Now().UnixNano())

	// welcoming message
	color.Red("the sleeping barber shop is welcoming you !")
	color.Red("-----------------------------------------")

	// define the channels we will need
	clients_channel := make(chan string, seat_capacity)
	barbers_done_channel := make(chan bool)

	// define the shop type
	shop := BarberShop{
		Shop_Seat_Capacity:   seat_capacity,
		Hair_Cut_Duration:    time_to_get_cut_done,
		Number_Of_Barbers:    0,
		Is_Open:              true,
		Clients_Channel:      clients_channel,
		Barbers_Done_Channel: barbers_done_channel,
	}

	// add barbers
	shop.AddBarber("fady")

	time.Sleep(5 * time.Second)
}
