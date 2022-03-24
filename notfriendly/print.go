package notfriendly

import (
	"fmt"

	"github.com/roidelapluie/testgowork/helloprint"
)

func Print() {
	helloprint.Print()
	fmt.Println("That's not true^")
}
