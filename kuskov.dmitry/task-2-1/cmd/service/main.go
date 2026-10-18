package main

import "fmt"

// Я так понял по комментариям у одногруппников, что нужен вот такой блок констант, но в моём решении как будто немного излишне.
const (
	Tmn int = 15
	Tmx int = 30
)

func optTemp(s string, mn, mx *int) int {
	switch s {
	case "<=" :
		var T int
		_, err := fmt.Scan(&T)
		if err != nil {
			fmt.Println("Invalid temperature")
		}
		if (T < *mx) {
			*mx = T
		}
	case ">=" :
		var T int
		_, err := fmt.Scan(&T)
		if err != nil {
			fmt.Println("Invalid temperature")
		}
		if (T > *mn) {
			*mn = T
		}
	}

	if (*mx - *mn >= 0) {
		return *mn
	} else {
		return -1
	}
}

func main() {
	var N int
	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println("Invalid N")
		return
	}

	for i := 0; i < N; i++ {
		var K int
		_, err := fmt.Scan(&K)
		if err != nil {
			fmt.Println("Invalid K")
			return
		}

		var mn int = Tmn
		var mx int = Tmx

		for j := 0; j < K; j++ {
			var s string
			_, err := fmt.Scan(&s)
			if (err != nil || (s != "<=" && s != ">=")) {
				fmt.Println("Invalid measure")
				return
			}

			var temp int = optTemp(s, &mn, &mx)

			if (temp == -1) {
				fmt.Println("An optimal T cannot be reached")
				break
			} else {
				fmt.Println(temp)
			}
		}
	}
}
