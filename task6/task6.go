package main

import (
	"context"
	"math/rand"
)

func just_end(c chan int) {
	for i := 0; i < 100; i++ {
		c <- rand.Intn(1000)
	}
}

func forever(c chan int) {
	for {
		c <- rand.Intn(1000)
	}
}

func on_event(c chan int, s chan struct{}) {
	for {
		select {
		case _ = <-s:
			break
		default:
			c <- rand.Intn(1000)
		}
	}
}

func on_context(cont context.Context, c chan int) {
	for {
		select {
		case _ = <-cont.Done():
			break
		default:
			c <- rand.Intn(1000)
		}
	}
}
