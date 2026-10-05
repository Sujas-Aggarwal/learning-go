package BasicServer

import (
	"fmt"
	"net/http"
)

func StartServer() {
	fmt.Println("Starting Server at 8000")
	err := http.ListenAndServe(":8000", nil)
	fmt.Println(err)
}
