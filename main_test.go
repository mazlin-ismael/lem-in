package main

import (
	// "fmt"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)


var example01 string = `L1-t L2-h L3-0
L1-E L2-A L3-o L4-t L5-h L6-0
L1-a L2-c L3-n L4-E L5-A L6-o L7-t L8-h L9-0
L1-m L2-k L3-e L4-a L5-c L6-n L7-E L8-A L9-o L10-t
L1-end L2-end L3-end L4-m L5-k L6-e L7-a L8-c L9-n L10-E
L4-end L5-end L6-end L7-m L8-k L9-e L10-a
L7-end L8-end L9-end L10-m
L10-end
`


func TestMain(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "FILES/example01.txt", "--test")
	_, errOutput := cmd.CombinedOutput()
	if errOutput != nil {
		t.Error("errOutput", errOutput)
	}

	exampleLines := strings.Fields(example01)
	// for _, v := range exampleLines {
	// 	fmt.Println(v)
	// }
	fmt.Println(exampleLines)
}