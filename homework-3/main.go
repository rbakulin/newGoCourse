package main

import (
	"github.com/rbakulin/newGoCourse/homework-3/queue"
	"github.com/rbakulin/newGoCourse/homework-3/vehicle"
)

func main() {
	toyota := vehicle.PassengerCar{
		Vehicle: vehicle.Vehicle{
			Brand:"Toyota",
			Year:2010,
			IsEngineOn:false,
			AreWindowsOpen:false,
		},
		TrunkVolume: 10,
		TrunkLoad: 5,
	}
	bmw := vehicle.PassengerCar{
		Vehicle: vehicle.Vehicle{
			Brand:"BMW",
			Year:2020,
			IsEngineOn:true,
			AreWindowsOpen:true,
		},
		TrunkVolume: 10,
		TrunkLoad: 0,
	}
	mercedes := vehicle.Truck{
		Vehicle: vehicle.Vehicle{
			Brand:"Mercedes",
			Year:2018,
			IsEngineOn:true,
			AreWindowsOpen:false,
		},
		BodyVolume: 50,
		BodyLoad: 45,
	}

	queue.Enqueue(toyota)
	queue.Enqueue(bmw)
	queue.Enqueue(mercedes)
	dequeuedVehicle := queue.Dequeue()
	for dequeuedVehicle != nil {
		dequeuedVehicle.Print()
		dequeuedVehicle = queue.Dequeue()
	}
}
