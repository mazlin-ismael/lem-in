package errFile

import (
	"fmt"
	"os"
)

// Check if there is an error in the fonction in argument
func CheckFunc(f func() error) {
	errorFunc := f()
	if errorFunc != nil {
		fmt.Println(errorFunc)
		os.Exit(0)
	}
}
