package main

import "fmt"

func main() {
	fmt.Println("Hello, world !!")

	a := 5
	b := 35
	c := 3
	sum := a + b
	mul := a * c
	div := b / a
	frac := float64(b) / float64(c)
	rem := b % c

	fmt.Printf("%v + %v = %v \n", a,b,sum)
	fmt.Printf("%v x %v = %v \n", a,c,mul)
	fmt.Printf("%v / %v = %v \n", b,a,div)
	fmt.Printf("%v / %v = %f \n", b,c,frac)
	fmt.Printf("%d %% %v = %v \n", b,c,rem)
	fmt.Println()
	fmt.Println()

	x := 13 % 3
	fmt.Println(x)
	if x == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
	fmt.Println()
	fmt.Println()

	//nested for loop
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}
	fmt.Println()
	fmt.Println()

	// for while loop
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println()
	fmt.Println()


	// for loop with condition 
	for i := 1; i < 5; i++ {
		if i%2 == 1 {
			fmt.Println(i, "is Odd")
		} else {
			fmt.Println(i, "is Even")
		}
	}
	fmt.Println()
	fmt.Println()

	
}