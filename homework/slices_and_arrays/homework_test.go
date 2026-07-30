package main

import (
	"log/slog"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go test -v ./homework/slices_and_arrays

func TestEmptyWhenEmpty(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
}

func TestFullWhenEmpty(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.False(t, queue.Full(), "expect false, fact:%t", queue.Full())
}

func TestFullWhenFull(t *testing.T) {
	const queueSize = 2
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Full())
}

func TestFrontEmpty(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.Equal(t, -1, queue.Front())
}

func TestBackEmpty(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.Equal(t, -1, queue.Back())
}

func TestPopEmpty(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.False(t, queue.Pop())
}

func TestPushFull(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))
}

func TestEmptyWhenNotEmpty(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Push(1))
	assert.False(t, queue.Empty())
}

func TestPushThreeElem(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))
}

func TestEasyCaseFront(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	assert.Equal(t, 1, queue.Front())
}

func TestEasyCaseBack(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	assert.Equal(t, 3, queue.Back())
}

func TestPopEasyCase(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	assert.True(t, queue.Pop())
}

func TestPopWhenHeadAfterTail(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	// 1,2,3
	queue.Pop()
	// _,2,3
	queue.Pop()
	// _,_,3 : [_,_,head]tail
	queue.Push(4)
	// 4,_,3
	assert.True(t, queue.Pop())
}

func TestEmptyAfterPop(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	queue.Pop()
	assert.False(t, queue.Empty())
}

func TestFullAfterPop(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	queue.Pop()
	assert.False(t, queue.Full())
}

func TestPushAfterPopToFirstElem(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	queue.Pop()
	assert.True(t, queue.Push(4))
	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

}

func TestFrontAfterPushFirstElem(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	//1, 2, 3

	queue.Pop()
	//_, 2, 3

	queue.Push(4)
	//4, 2, 3

	assert.Equal(t, 2, queue.Front())
}

func TestBackAfterPushFirstElem(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	//1, 2, 3

	queue.Pop()
	//_, 2, 3

	queue.Push(4)
	//4, 2, 3

	assert.Equal(t, 4, queue.Back())
}

func TestPopToEmpty(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	//1, 2, 3

	queue.Pop()
	//_, 2, 3

	queue.Push(4)
	//4, 2, 3

	assert.True(t, queue.Pop()) //4, _, 3
	assert.True(t, queue.Pop()) //4, _, _
	assert.True(t, queue.Pop()) //_, _, _
	assert.False(t, queue.Pop())
}

func TestEmptyAfterFourPop(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	//1, 2, 3

	queue.Pop()
	//_, 2, 3

	queue.Push(4)
	//4, 2, 3

	queue.Pop() //4, _, 3
	queue.Pop() //4, _, _
	queue.Pop() //_, _, _
	queue.Pop()

	assert.True(t, queue.Empty())
}

func TestFullAfterFourPop(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	//1, 2, 3

	queue.Pop()
	//_, 2, 3

	queue.Push(4)
	//4, 2, 3

	queue.Pop() //4, _, 3
	queue.Pop() //4, _, _
	queue.Pop() //_, _, _
	queue.Pop()

	assert.False(t, queue.Full())
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))

	assert.True(t, queue.Pop())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}

func TestPopWrappedNotAtEnd(t *testing.T) {
	queue := NewCircularQueue(4) //_,_,_,_
	queue.Push(1)                //1,_,_,_
	queue.Push(2)                //1,2,_,_
	queue.Push(3)                //1,2,3,_
	queue.Push(4)                //1,2,3,4
	queue.Pop()                  //_,2,3,4
	queue.Pop()                  //_,_,3,4
	queue.Push(5)                //5,_,3,4

	queue.Pop()                       //5,_,_,4
	assert.Equal(t, 4, queue.Front()) // сейчас упадёт
}

func TestPushIntoGapWhenWrapped(t *testing.T) {
	const queueSize = 4
	queue := NewCircularQueue(queueSize) //_,_,_,_
	queue.Push(1)                        //1,_,_,_
	queue.Push(2)                        //1,2,_,_
	queue.Push(3)                        //1,2,3,_
	queue.Push(4)                        //1,2,3,4
	queue.Pop()                          //_,2,3,4
	queue.Pop()                          //_,_,3,4
	queue.Push(5)                        //5,_,3,4
	queue.Push(6)                        //5,6,3,4

	assert.Equal(t, 3, queue.Front())
	assert.Equal(t, 6, queue.Back())
	assert.True(t, reflect.DeepEqual([]int{5, 6, 3, 4}, queue.values))
}

func TestPushAfterEmptyQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize) //_,_,_

	queue.Push(1) //1,_,_
	queue.Push(2) //1,2,_
	queue.Push(3) //1,2,3
	queue.Pop()   //_,2,3
	queue.Push(4) //4,2,3
	queue.Pop()   //4,_,3
	queue.Pop()   //4,_,_
	queue.Pop()   //_,_,_ empty, но head=0 tail=1

	assert.True(t, queue.Empty())
	assert.True(t, queue.Push(10)) //10,_,_
	assert.Equal(t, 10, queue.Front())
	assert.Equal(t, 10, queue.Back())
}

func init() {

	logHandler := slog.NewTextHandler(
		os.Stderr,
		&slog.HandlerOptions{Level: slog.LevelDebug})

	logger := slog.New(logHandler)

	slog.SetDefault(logger)
}
