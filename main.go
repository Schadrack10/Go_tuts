package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {

	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name:")
	// name,_ := reader.ReadString('\n')
	// name =  strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name:", reader)

	b := createNewBill(name)
	fmt.Println("Created new bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose options(a- add item, s - save the bill , t - add tip ,c - check current bill): ", reader)

	fmt.Println(opt)

	//check the options returned
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("the Price must be a number: ", err)
			promptOptions(b)
		}

		b.addItems(name, p)

		fmt.Println(name, price)
    promptOptions(b)

	case "t":
		tip, _ := getInput("Enter tip amount: ", reader)


    t,err := strconv.ParseFloat(tip,64)

    if err != nil {
      fmt.Println("The Tip must be a number: ",err)
      promptOptions(b)
    }

    b.updateTip(t)

		fmt.Println("Tip added - ",tip)
    promptOptions(b)

	case "s":
    b.saveBill()
		fmt.Println("you saved the file: ",b.name)

	case "c":
		fmt.Println("the current bill", b)
	default:
		fmt.Println("That was not a valid option")
		promptOptions(b)
	}

}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}** Real world fuel consumption averages may vary according to factors such as driving styles, environmental conditions, load, wheel and tyre fitment, accessories and options fitted.
