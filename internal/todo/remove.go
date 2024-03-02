package todo

import (
	"fmt"
	"os"
)

func Remove() {
	if len(os.Args) >= 2 {
		fmt.Println("Remove: Please specify the item IDs.")
		return
	}
}
