package main

import (
	"fmt"
	"math"
	"wb-l1/task11"
)

func backet(data []float64) map[int]task11.Set[float64] {
	ret := make(map[int]task11.Set[float64])

	for _, n := range data {
		var ind int = (int(math.Trunc(n)) / 10) * 10

		b, ok := ret[ind]

		if !ok {
			empty := task11.NewSet[float64]()
			empty.Add(n)
			ret[ind] = empty
		} else {
			b.Add(n)
		}
	}

	return ret
}

func main() {
	backets := backet([]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5})
	for key, value := range backets {
		fmt.Print(key, " {")

		for n := range value.Iterator() {
			fmt.Print(n, ",")
		}

		fmt.Println("}")
	}
}
