package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	Shop_Seat_Capacity   int
	Hair_Cut_Duration    time.Duration
	Number_Of_Barbers    int
	Clients_Channel      chan string
	Barbers_Done_Channel chan bool // we send true when barber done
	Is_Open              bool
}

func (bs *BarberShop) AddBarber(barber string) {
	// wake this barber up
	isSleeping := false

	// add the number of barbers by one
	bs.Number_Of_Barbers++

	// the routine of the barber
	for {
		// if there is no clients the barber can take a nap
		if len(bs.Clients_Channel) == 0 {
			color.Yellow("there is no clients, so the %s will go to take a nap .. ", barber)
			isSleeping = true
		}

		// while the barber is sleeping, he is waiting for any client to enter, if any client enter the client must wake the barber up
		client, ok := <-bs.Clients_Channel
		if ok {
			// if barber is sleeping
			if isSleeping {

				color.Blue("customer %s wakes the %s barber up ", client, barber)
				// the barber wake up
				isSleeping = false
			}
			// starts cut hair
			bs.CutHair(barber, client)
		} else {

			bs.BarberGoesHome(barber)
			// end the go routine
			return
		}
	}
}

func (bs *BarberShop) CutHair(barber, client string) {
	color.Cyan("%s barber is cutting client %s 's hair", barber, client)
	time.Sleep(bs.Hair_Cut_Duration)
	color.Green("%s barber finished cutting the client %s 's hair", barber, client)
}

func (bs *BarberShop) BarberGoesHome(barber string) {
	color.Green("%s barber left to home", barber)
	// the barber should go home so we set the barber done channel to true
	bs.Barbers_Done_Channel <- true
}
