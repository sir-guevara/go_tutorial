// takes name and address as input and prints it in json format

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(" Please enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf(" Please enter your address: ")
	addr, _ := reader.ReadString('\n')
	userData := map[string]string{
		"name":    strings.TrimSpace(name),
		"address": strings.TrimSpace(addr),
	}

	jsonData, _ := json.Marshal(userData)
	fmt.Println(string(jsonData))
}
