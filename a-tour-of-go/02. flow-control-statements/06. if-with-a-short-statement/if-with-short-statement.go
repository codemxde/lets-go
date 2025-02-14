package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main()  {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20), // notice the additional (,) at the end
	)

	// fmt.Println(pow(1,2, 3), pow(2,3, 4)) // here we don't have to specify extra comma (,) at the end
}