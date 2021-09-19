package main

import "fmt"

func getCar(carType string) (iCar, error) {
	if carType == "crv" {
		return newCRV(), nil
	}
	if carType == "bolt" {
		return newBolt(), nil
	}
	return nil, fmt.Errorf("wrong car type passed")
}
