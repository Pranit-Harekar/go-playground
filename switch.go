package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its weekend")
	default:
		fmt.Println("its weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("its before noon")
	default:
		fmt.Println("its after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println(i, "is boolean")
		case string:
			fmt.Println(i, "is string")
		case int:
			fmt.Println(i, "is integer")
		default:
			fmt.Printf("Don't know the type %T\n", t)
		}
	}

	whatAmI(1)
	whatAmI("hello")
	whatAmI(true)
	whatAmI(1.999999)
}
