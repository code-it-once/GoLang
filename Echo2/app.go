package Echo2

import (
	"fmt"
	"os"
)

func Echo() {

	var message, seperator string

	for i := 1; i < len(os.Args); i++ {
		message += seperator + os.Args[i]
		seperator = " "
	}

	fmt.Println(message)
}
