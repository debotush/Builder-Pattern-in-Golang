package main

import "fmt"

type Bike struct{
	name string
	NoOfCylinders int
	CC int
	fuelCapacity float32
}
func (bike Bike) ShowMyBike(){
	fmt.Println("Bike Name : " , bike.name)
       	fmt.Println("Engine Capacity : " , bike.CC)
       	fmt.Println("No of Cylinders : " , bike.NoOfCylinders)
       	fmt.Println("Fuel Capacity : " , bike.fuelCapacity)
}

type BikeBuilder interface {
	buildBike() Bike
}

type RoyelEnfield struct {
}
func (bk RoyelEnfield) buildBike() Bike {
	Enfield := Bike{}
	Enfield.name = "Royel Enfield classic 350"
	Enfield.CC = 346
	Enfield.NoOfCylinders = 1
	Enfield.fuelCapacity = 13.5
	return Enfield;
}
type Production struct {
	builder BikeBuilder
}
func (p Production)BuildEnfield() Bike {
	return p.builder.buildBike()
}

func main() {
	production := Production{RoyelEnfield{}}
	res := production.BuildEnfield()
	res.ShowMyBike()
}  
