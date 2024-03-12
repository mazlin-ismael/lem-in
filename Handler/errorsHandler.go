package Handler

import (
	"fmt"
	"os"
)

func CheckFunc(f func() error) {
	errorFunc := f()
	if errorFunc != nil {
		fmt.Println(errorFunc)
		os.Exit(0)
	}
}
