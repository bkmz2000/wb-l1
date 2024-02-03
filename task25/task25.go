package main

import "time"

func wait(dt time.Duration) {
	now := time.Now()

	end := now.Add(dt)

	for now.Before(end) {
	}
}
