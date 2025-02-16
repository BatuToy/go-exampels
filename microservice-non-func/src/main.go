package src

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
