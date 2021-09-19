package main

type bolt struct {
	car
}

func newBolt() iCar {
	return &bolt{
		car: car{
			name:  "bolt",
			power: 99,
		},
	}
}
