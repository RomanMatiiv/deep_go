package main

import (
	"fmt"
	"log/slog"
)

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
	Front() int          // получить значение из начала очереди (-1, если очередь пустая)
	Back() int           // получить значение из конца очереди (-1, если очередь пустая)
	Empty() bool         // проверить пустая ли очередь
	Full() bool          // проверить заполнена ли очередь
}

type CircularQueue struct {
	values   []int
	initSize int
	curSize  int
	headIdx  int // всегда указывает на первый(или нулевой) элемент. Инвариант.
	tailIdx  int // всегда указывает на элемент следующий за последним. Инвариант.
}

func NewCircularQueue(size int) CircularQueue {
	slog.Debug(fmt.Sprintf("Init(size=%d)", size))
	return CircularQueue{
		initSize: size,
		curSize:  0,
		headIdx:  0,
		tailIdx:  0,
		values:   make([]int, size, size)}
}

// Empty проверить пустая ли очередь
func (q *CircularQueue) Empty() bool {
	return q.curSize == 0
}

// Full проверить заполнена ли очередь
func (q *CircularQueue) Full() bool {
	return q.curSize == q.initSize
}

// Push добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueue) Push(value int) bool {
	slog.Debug(fmt.Sprintf("Push(%d) start", value))
	defer func() { slog.Debug(fmt.Sprintf("Push(%d) end", value)) }()

	if q.Full() {
		return false
	}

	slog.Debug(fmt.Sprintf("before Push(%d) when head=%d tail=%d", value, q.headIdx, q.tailIdx))
	defer func() { slog.Debug(fmt.Sprintf("after Push(%d) when head=%d tail=%d", value, q.headIdx, q.tailIdx)) }()

	idx := -1
	// ht,_,_,_,_
	if q.tailIdx == 0 && q.headIdx == 0 {
		idx = 1
	}
	// head,tail,_,_,_

	// _,head,_,tail,_
	if q.tailIdx < q.initSize && q.tailIdx > q.headIdx {
		idx = q.tailIdx + 1
		// _,head,_,_,tail

		// [_,_,_,head]tail
	} else if q.tailIdx == q.initSize {
		idx = 1
		// [_,tail,_,head]

		// _,tail,_,head,_
	} else if q.tailIdx < q.headIdx {
		idx = q.tailIdx + 1
	}

	q.values[idx-1] = value
	q.tailIdx = idx

	q.curSize += 1
	return true
}

// Back получить значение из конца очереди (-1, если очередь пустая)
func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.tailIdx-1]
}

// Pop удалить значение из начала очереди (false, если очередь пустая)
func (q *CircularQueue) Pop() bool {
	slog.Debug("Pop()")

	slog.Debug(fmt.Sprintf("before Pop() when head=%d tail=%d", q.headIdx, q.tailIdx))
	defer func() { slog.Debug(fmt.Sprintf("after Pop() when head=%d tail=%d", q.headIdx, q.tailIdx)) }()

	if q.Empty() {
		return false
	}

	// head,_,_,tail,_
	if q.headIdx < q.tailIdx {
		q.headIdx = q.headIdx + 1
	}
	// _,head,_,tail,_

	// _,_,tail,head,_
	// _,_,tail,_,head
	if q.headIdx >= q.tailIdx {
		if q.headIdx+1 < q.initSize {
			q.headIdx += 1
		} else {
			q.headIdx = 0
		}
	}
	// _,_,tail,_,head
	// head,_,tail,_,_

	q.curSize -= 1

	if q.Empty() {
		q.headIdx = 0
		q.tailIdx = 0
	}

	return true
}

// Front вернуть значение из начала очереди (-1, если очередь пустая)
func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.headIdx]
}
