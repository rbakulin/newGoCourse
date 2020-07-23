package queue

type LoadablePrintable interface {
	Fullness() float64
	Print()
}

var queueArray []LoadablePrintable

func Enqueue(vehicle LoadablePrintable) {
	queueArray = append(queueArray, vehicle)
}

func Dequeue() LoadablePrintable {
	if len(queueArray) == 0 {
		return nil
	}

	itemToDequeue := queueArray[0]
	queueArray = queueArray[1:]
	itemToDequeue.Print()
	return itemToDequeue
}
