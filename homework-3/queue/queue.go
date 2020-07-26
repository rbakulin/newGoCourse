package queue

type LoadablePrintable interface {
	Fullness() float64
	Print()
}

type VehicleQueue struct {
	store []LoadablePrintable
}

func (q *VehicleQueue) Enqueue(vehicle LoadablePrintable) {
	q.store = append(q.store, vehicle)
}


func (q *VehicleQueue) Dequeue() LoadablePrintable {
	if len(q.store) == 0 {
		return nil
	}

	itemToDequeue := q.store[0]
	q.store = q.store[1:]
	return itemToDequeue
}

func (q *VehicleQueue) GetQueueLength() int {
	return len(q.store)
}
