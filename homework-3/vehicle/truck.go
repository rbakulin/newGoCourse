package vehicle

import "fmt"

type Truck struct {
	Vehicle
	//Let's imagine that only trucks have a body
	BodyVolume float64
	BodyLoad float64
}

func (t Truck) Fullness() float64 {
	return t.BodyLoad / t.BodyLoad * 100
}

func (t Truck) Print() {
	fmt.Printf(
		"Model: %s\nYear: %d\nEngine is on: %t\nWindows are open: %t\nFullness: %g\n-------------\n",
		t.Brand, t.Year, t.IsEngineOn, t.AreWindowsOpen, t.Fullness())
}
