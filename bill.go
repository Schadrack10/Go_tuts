package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func createNewBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format bill function
func (b bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0

	//list items
	for k, val := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", val)
		total += val
	}

	// add tip 
	fs += fmt.Sprintf("%-25v ...$%v \n", "tip:", b.tip)


	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total + b.tip)

	return fs
}

//update
func (b *bill) updateTip (tip float64){
	//(*b).tip = tip // here we de-reference the struct, but we don not need to do these go is samrt enought to know that we pass ina pointer
	b.tip = tip
}  

//add an item to the bill
func (b *bill) addItems(name string , price float64){
	b.items[name] =  price
}


// save bill
func (b *bill) saveBill(){
   data := []byte(b.format())

  err := os.WriteFile("bills/"+b.name+".txt",data,0644)

  if err !=  nil {
	panic(err)
  }

  fmt.Println("Bill was saved to file")
}
