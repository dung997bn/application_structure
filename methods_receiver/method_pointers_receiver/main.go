package main

import "fmt"

type car struct {
	brand string
	price int
}

func changeCar(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice //*c:error
}

// type distance *int
// func (d *distance) m1() {
// 	fmt.Println("just a message")
// }
//invalid receiver type *distance( distance is a poiter type)

func main() {
	myCar := car{brand: "Audi", price: 64000}
	changeCar(myCar, "BMW", 20000)
	fmt.Println(myCar) //{Audi 64000} -> no changes

	myCar.changeCar1("BMW", 20000)
	fmt.Println(myCar) //{Audi 64000} -> no changes

	(&myCar).changeCar2("BMW", 20000)
	fmt.Println(myCar) //{BMW 20000} -->changes value, () is required

	(&myCar).changeCar2("Seat", 543333)
	fmt.Println(myCar) //{Seat 543333}

	myCar.changeCar2("Seat1", 543333) //->or this is accepted
	fmt.Println(myCar)                //{Seat1 543333}

	var yourCar *car
	yourCar = &myCar
	fmt.Println(*yourCar) //{Seat1 543333}

	//valid ways
	yourCar.changeCar2("MAZDA", 50000)
	fmt.Println(*yourCar) //{MAZDA 50000}

	(*yourCar).changeCar2("MAZDA", 50000)
	fmt.Println(*yourCar) //{MAZDA 50000}

	fmt.Println(myCar) //{MAZDA 50000}

}
