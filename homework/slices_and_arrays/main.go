package main

// В домашнем задании нужно реализовать кольцевую очередь (Circular Queue).
//
// **Кольцевая очередь** (*Circular Queue*) — это структура данных,
// которая представляет собой очередь (*FIFO*) фиксированного размера.
// Кольцевая очередь использует буфер фиксированного размера таким образом, как будто бы после последнего элемента
// сразу же снова идет первый (*как представлено ниже*).
//
// NewCircularQueue()
// ┌───┐ ┌───┐ ┌───┐
// │   │ │   │ │   │
// └───┘ └───┘ └───┘
//
// Push(1)
// ┌───┐ ┌───┐ ┌───┐
// │ 1 │ │   │ │   │
// └───┘ └───┘ └───┘
//
// Push(2), Push(3)
// ┌───┐ ┌───┐ ┌───┐
// │ 1 │ │ 2 │ │ 3 │
// └───┘ └───┘ └───┘
//
// Pop()
// ┌───┐ ┌───┐ ┌───┐
// │   │ │ 2 │ │ 3 │   ← «1» удалён, слот освободился
// └───┘ └───┘ └───┘
//
// Push(4)
// ┌───┐ ┌───┐ ┌───┐
// │ 4 │ │ 2 │ │ 3 │   ← «4» занял первый свободный слот (кольцо «замкнулось»)
// └───┘ └───┘ └───┘
// Подробнее можно прочитать [здесь](https://www.programiz.com/dsa/circular-queue).
//
// Такая структура много где используется, например для организации различных очередей сообщений и а также буффер в буфферезированных каналах Go реализован в виде кольцевой очереди.
type Queue interface {
	Push(value int) bool // добавить значение в конец очереди (false, если очередь заполнена)
	Pop() bool           // удалить значение из начала очереди (false, если очередь пустая)
	Front() int          // удалить значение из начала очереди (false, если очередь пустая)
	Back() int           // получить значение из конца очереди (-1, если очередь пустая)
	Empty() bool         /// проверить пустая ли очередь
	Full() bool          // проверить заполнена ли очередь
}

type CircularQueue struct {
	values   []int
	initSize int
}

func NewCircularQueue(size int) CircularQueue {
	values := make([]int, 0, 3)
	return CircularQueue{
		initSize: size,
		values:   values}
}

func (q *CircularQueue) Empty() bool {
	return len(q.values) == 0
}

func (q *CircularQueue) Full() bool {
	return len(q.values) == q.initSize
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
