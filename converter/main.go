package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	rc "github.com/ghostzali/romanconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Choose operation mode:")
		fmt.Println("1: Roman to Arabic")
		fmt.Println("2: Arabic to Roman")
		fmt.Println("exit: Close the program")
		fmt.Print(">> ")
		option, _ := reader.ReadString('\n')
		option = strings.Replace(option, "\n", "", -1)
		if option == "exit" {
			break
		}
		switch option {
		case "1":
			fmt.Println("Input roman numeral:")
			fmt.Print(">> ")
			roman, _ := reader.ReadString('\n')
			roman = strings.Replace(roman, "\n", "", -1)
			valid := rc.Validate(roman)
			if !valid {
				fmt.Println("Invalid input! Please enter valid roman numeral")
				continue
			}
			dec, err := rc.Convert(roman)
			if err != nil {
				fmt.Printf("Error with message: %s\n", err.Error())
				continue
			}
			fmt.Printf("Roman %q equal to Arabic \"%d\"\n", roman, dec)
		case "2":
			fmt.Println("Input decimal number")
			fmt.Print(">> ")
			arabic, _ := reader.ReadString('\n')
			arabic = strings.Replace(arabic, "\n", "", -1)
			dec, err := strconv.ParseInt(arabic, 10, 0)
			if err != nil {
				fmt.Println("Invalid input! Please enter arabic number")
				continue
			}
			roman, err := rc.Parse(int(dec))
			if err != nil {
				fmt.Printf("Error with message: %s\n", err.Error())
				continue
			}
			fmt.Printf("Arabic \"%d\" equal to Roman %q\n", dec, roman)
		default:
			continue
		}
		time.Sleep(30 * time.Second)
	}
}
