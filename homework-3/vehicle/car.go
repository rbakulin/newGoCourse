package vehicle

import "fmt"

type PassengerCar struct {
	Vehicle
	//Let's imagine that only cars have a trunk
	TrunkVolume float64
	TrunkLoad float64
}

func (c *PassengerCar) Fullness() float64 {
	return (c.TrunkLoad) / (c.TrunkVolume) * 100
}

func (c *PassengerCar) Print() {
	fmt.Printf(
		"Model: %s\nYear: %d\nEngine is on: %t\nWindows are open: %t\nFullness: %g\n-------------\n",
		c.Brand, c.Year, c.IsEngineOn, c.AreWindowsOpen, c.Fullness())
}
