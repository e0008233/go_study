package main

import (
	"fmt"
)

func main() {
	var things []string
	things = append(things, "1")
	for _, s := range things {
		fmt.Printf(s)
	}
	fmt.Println(things) // true

	//fmt.Println(amazon.CountFamilyLogins([]string{"ab", "bc", "zz", "aa", "bb", "cc", "cc", "cc", "aa"}))
}

// change to existing kyc enum for event tracking
