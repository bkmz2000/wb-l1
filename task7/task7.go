package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func map_abuser(m map[int]string, mu *sync.RWMutex) {
	start := time.Now()

	for time.Now().Sub(start).Seconds() < float64(5) {
		r := rand.Intn(2)

		if r == 0 && len(m) > 0 {
			mu.RLock()
			fmt.Println("regular: ", m[rand.Intn(len(m))])
			mu.RUnlock()
		} else {
			mu.Lock()
			m[len(m)+1] = generateRandomString(10)
			mu.Unlock()
		}
	}
}

func concurrent_map_abuser(m *sync.Map, count *atomic.Int32) {
	start := time.Now()

	for time.Now().Sub(start).Seconds() < float64(5) {
		r := rand.Intn(2)

		if r == 0 && count.Load() > 0 {
			v, _ := m.Load(rand.Int31n(count.Load()))
			fmt.Println("concurent: ", v)
		} else {
			count.Add(1)
			m.Store(count.Load()+1, generateRandomString(10))
		}
	}
}

func main_7() {
	m := make(map[int]string)
	cm := sync.Map{}
	mu := sync.RWMutex{}
	wg := sync.WaitGroup{}
	count := atomic.Int32{}

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			map_abuser(m, &mu)
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			concurrent_map_abuser(&cm, &count)
			wg.Done()
		}()
	}

	wg.Wait()
}
