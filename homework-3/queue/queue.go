package queue

type LoadablePrintable interface {
	Fullness() float64
	Print()
}

type VehicleQueue struct {
	Store []LoadablePrintable
}

func (q *VehicleQueue) Enqueue(vehicle LoadablePrintable) {
	_ = append(q.Store, vehicle)
}


func (q *VehicleQueue) Dequeue() LoadablePrintable {
	if len(q.Store) == 0 {
		return nil
	}

	itemToDequeue := q.Store[0]
	q.Store = q.Store[1:]
	return itemToDequeue
}
