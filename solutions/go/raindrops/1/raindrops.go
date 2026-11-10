package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	// panic("Please implement the Convert function")

	var isDivisibleByThree = (number % 3) == 0
	var isDivisibleByFive = (number % 5) == 0
	var isDivisibleBySeven = (number % 7) == 0

	if isDivisibleByThree && isDivisibleByFive && !isDivisibleBySeven {
		return ("PlingPlang")
	}
	if isDivisibleByThree && isDivisibleBySeven && !isDivisibleByFive {
		return ("PlingPlong")
	}
	if isDivisibleByFive && isDivisibleBySeven && !isDivisibleByThree {
		return ("PlangPlong")
	}
	if isDivisibleByThree && isDivisibleByFive && isDivisibleBySeven {
		return "PlingPlangPlong"
	}
	if isDivisibleByThree && !isDivisibleByFive && !isDivisibleBySeven {
		return "Pling"
	}
	if isDivisibleByFive && !isDivisibleByThree && !isDivisibleBySeven {
		return "Plang"
	}
	if isDivisibleBySeven && !isDivisibleByThree && !isDivisibleByFive {
		return "Plong"
	} else {
		return strconv.Itoa(number)
	}

}
