package main

import "fmt"

func main() {
	var mes int
	var dia int
	fmt.Scan(&dia)
	fmt.Scan(&mes)
	switch {
	case mes == 1:
		if dia < 21 {
			fmt.Println("capricornio")
		} else if dia <= 31 {
			fmt.Println("acuario")
		}
	case mes == 2:
		if dia < 21 {
			fmt.Println("acuario")
		} else if dia <= 29 {
			fmt.Println("piscis")
		}
	case mes == 3:
		if dia < 21 {
			fmt.Println("piscis")
		} else if dia <= 31 {
			fmt.Println("aries")
		}
	case mes == 4:
		if dia < 21 {
			fmt.Println("aries")
		} else if dia <= 31 {
			fmt.Println("tauro")
		}
	case mes == 5:
		if dia < 22 {
			fmt.Println("tauro")
		} else if dia <= 31 {
			fmt.Println("geminis")
		}
	case mes == 6:
		if dia < 22 {
			fmt.Println("geminis")
		} else if dia <= 31 {
			fmt.Println("cancer")
		}
	case mes == 7:
		if dia < 23 {
			fmt.Println("cancer")
		} else if dia <= 31 {
			fmt.Println("leo")
		}
	case mes == 8:
		if dia < 23 {
			fmt.Println("leo")
		} else if dia <= 31 {
			fmt.Println("virgo")
		}
	case mes == 9:
		if dia < 24 {
			fmt.Println("virgo")
		} else if dia <= 31 {
			fmt.Println("libra")
		}
	case mes == 10:
		if dia < 25 {
			fmt.Println("libra")
		} else if dia <= 31 {
			fmt.Println("escorpio")
		}
	case mes == 11:
		if dia < 23 {
			fmt.Println("escorpio")
		} else if dia <= 31 {
			fmt.Println("sagitario")
		}
	case mes == 12:
		if dia < 22 {
			fmt.Println("sagitario")
		} else if dia <= 31 {
			fmt.Println("capricornio")
		}
	default:

	}

}
