package queue

type loadablePrintable interface {
	Fullness() float64
	Print()
}

type vehicleQueue struct {
	Store []loadablePrintable
}

func (q *vehicleQueue) Enqueue(vehicle loadablePrintable) {
	_ = append(q.Store, vehicle)
}


func (q *vehicleQueue) Dequeue() loadablePrintable {
	if len(q.Store) == 0 {
		return nil
	}

	itemToDequeue := q.Store[0]
	q.Store = q.Store[1:]
	return itemToDequeue
}
