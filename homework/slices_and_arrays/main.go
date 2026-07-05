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
	//curSize int
	//curIdx int
	headIdx int // всегда указывает на первый(или нулевой) элемент
	tailIdx int // всегда указывает на элемент следующий за последним
}

func NewCircularQueue(size int) CircularQueue {
	slog.Debug(fmt.Sprintf("Init queue with initSize: %d", size))
	return CircularQueue{
		initSize: size,
		//curSize: 0,
		//curIdx: 0,
		headIdx: 0,
		tailIdx: 0,
		values:  make([]int, size, size)}
}

// Empty проверить пустая ли очередь
func (q *CircularQueue) Empty() bool {
	slog.Debug(fmt.Sprintf("Empty() when head=%d tail=%d", q.headIdx, q.tailIdx))

	if q.headIdx == q.tailIdx {
		return true
	}

	return false
}

// Full проверить заполнена ли очередь
func (q *CircularQueue) Full() bool {
	slog.Debug(fmt.Sprintf("Full() when head=%d tail=%d", q.headIdx, q.tailIdx))

	if q.headIdx == q.tailIdx {
		if q.headIdx == 0 && q.tailIdx == 0 {
			return false
		}
		panic("equal index, not in start")
	}

	// когда headIdx < tailIdx
	// h,_,_,t ; t-h<initSize-1 ; true
	// h,_,t,_ ; t-h<initSize-1 ; false
	if q.headIdx < q.tailIdx {
		return q.tailIdx-q.headIdx == q.initSize
	}

	// todo переделать
	// когда tailIdx < headIdx
	// _,t,_,h min,max = minMax(head, tail); max-min>1 false
	// _,_,t,h min,max = minMax(head, tail); max-min>1 true
	// t,_,h,_ min,max = minMax(head, tail); max-min>1 false
	// _,t,h,_ min,max = minMax(head, tail); max-min>1 true
	minIdx := min(q.headIdx, q.tailIdx)
	maxIdx := max(q.headIdx, q.tailIdx)
	if q.tailIdx < q.headIdx {
		return maxIdx-minIdx > 1
	}

	panic("impossible situation")
}

// Push добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}

	slog.Debug(fmt.Sprintf("Push(%d) when head=%d tail=%d", value, q.headIdx, q.tailIdx))

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

		// _,_,_,head,tail
	} else if q.tailIdx == q.initSize && q.headIdx > 0 {
		idx = 0

		// _,tail,_,head,_
	} else if q.tailIdx < q.headIdx && q.headIdx > 0 {
		idx = min(0, q.headIdx, q.tailIdx)
	}

	if idx == -1 {
		panic("out of range")
	}

	q.values[idx-1] = value
	q.tailIdx = idx

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
	if q.Empty() {
		return false
	}
	return true
}

// Front вернуть значение из начала очереди (-1, если очередь пустая)
func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	//todo
	return 1
}
