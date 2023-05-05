package main

import (
	"fmt"
)

type Car struct {
	Type   string
	FuelIn float32
}

func (c *Car) Estimasi() float32 {

	konsumsi := 0.6666666666 //jarak yang ditempuh untuk setiapl 1 L, dihitung dengan patokan 1.5 L/mill.
	jarak := c.FuelIn * float32(konsumsi)
	return jarak
}

func main() {

	myCar := Car{Type: "Pick Up", FuelIn: 3}
	estimasiJarak := myCar.Estimasi()
	fmt.Printf("Dengan bahan bakar terisi %.1f Liter, mobil %s dapat menempuh jarak sejauh %.1f Mill", myCar.FuelIn, myCar.Type, estimasiJarak)

}
