package main

type crv struct {
	car
}

func newCRV() iCar {
	return &crv{
		car: car{
			name:  "crv",
			power: 11,
		},
	}
}
