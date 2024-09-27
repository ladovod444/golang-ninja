package vehicle

import "fmt"

type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding up")
}
func (c Car) Brake() {
	fmt.Println("Stopping")
}
func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding up")
}
func (t Truck) Brake() {
	fmt.Println("Stopping")
}
func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}
func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

func TryVehicle(v Vehicle) {
	v.Accelerate()
	v.Brake()
	v.Steer("Right")

	truck, ok := v.(Truck)
	if ok {
		truck.LoadCargo("Cargo!!!")
	} else {
		fmt.Println("That is not a Truck")
	}

}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(direction string)
}
