package main

import (
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

	assert.False(t, queue.Full())
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

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
