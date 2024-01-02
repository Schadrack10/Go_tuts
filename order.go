package main

import "fmt"

type order struct {
	orderId        int
	orderItems     map[string]float64
	orderReference string
}

func createOrder(ref string) order {
	newOrder := order{
		orderId:        123764783,
		orderItems:     map[string]float64{"Brown Bread": 23.60, "Nike shoes": 50.66},
		orderReference: ref,
	}

	return newOrder
}

func (o order) formatOrder() string {

	var text string = ""

	for k, v := range o.orderItems {
		text += fmt.Sprintf("%-25v R%v \n", k, v)
	}

	return text
}
