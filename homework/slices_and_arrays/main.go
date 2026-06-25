package main

type Queue interface {
	Push(value int) bool // добавить значение в конец очереди (false, если очередь заполнена)
	Pop() bool           // удалить значение из начала очереди (false, если очередь пустая)
	Front() int          // удалить значение из начала очереди (false, если очередь пустая)
	Back() int           // получить значение из конца очереди (-1, если очередь пустая)
	Empty() bool         /// проверить пустая ли очередь
	Full() bool          // проверить заполнена ли очередь
}

type CircularQueue struct {
	values []int
	// todo
}

func NewCircularQueue(size int) CircularQueue {
	// todo
	return CircularQueue{}
}

func (q *CircularQueue) Push(value int) bool {
	//todo
	return false
}

func (q *CircularQueue) Pop() bool {
	// todo
	return false
}

func (q *CircularQueue) Front() int {
	// todo
	return 1
}

func (q *CircularQueue) Back() int {
	// todo
	return 1
}

func (q *CircularQueue) Empty() bool {
	// todo
	return false
}

func (q *CircularQueue) Full() bool {
	// todo
	return false
}
