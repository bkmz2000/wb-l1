package task11

import (
	"fmt"
	"math/rand"
)

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

type Set[T comparable] struct {
	data map[T]bool
	size int
}

func NewSet[T comparable]() Set[T] {
	ret := Set[T]{}
	ret.data = make(map[T]bool)
	return ret
}

// amortized O(1)
func (s *Set[T]) Add(v T) {
	if !s.Has(v) {
		s.data[v] = true
		s.size++
	}
}

// amortized O(1)
func (s *Set[T]) Has(v T) bool {
	_, ok := s.data[v]
	return ok
}

// amortized O(len(s.data))
func (s *Set[T]) Equals(other Set[T]) bool {
	if s.size != other.size {
		return false
	}

	for k, _ := range s.data {
		if !other.Has(k) {
			return false
		}
	}

	return true
}

// amortized O(len(s.data))
func (s *Set[T]) Intersect(other Set[T]) Set[T] {
	ret := NewSet[T]()

	for k, _ := range s.data {
		if other.Has(k) {
			ret.Add(k)
		}
	}

	return ret
}

func (s *Set[T]) Iterator() <-chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)
		for key, _ := range s.data {
			ch <- key
		}
	}()

	return ch
}

func main11() {
	LA := rand.Intn(1000)
	sa := NewSet[string]()

	LB := rand.Intn(1000)
	sb := NewSet[string]()

	for i := 0; i < LA; i++ {
		sa.Add(generateRandomString(10))
	}

	for i := 0; i < LB; i++ {
		sb.Add(generateRandomString(10))
	}

	LC := rand.Intn(1000)
	sc := NewSet[string]()

	for i := 0; i < LC; i++ {
		str := generateRandomString(10)
		sa.Add(str)
		sb.Add(str)
		sc.Add(str)
	}

	if sc.Equals(sa.Intersect(sb)) {
		fmt.Println("Ok")
	} else {
		fmt.Println("No ok")
	}
}
