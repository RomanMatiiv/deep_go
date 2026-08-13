package main

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestCOWSmoke(t *testing.T) {
	data := []byte{'a', 'b', 'c', 'd'}

	buffer := NewCOWBuffer(data)
	_ = buffer
}

func TestCOWBufferEqualInitBuffer(t *testing.T) {
	t.Skipf("not implement")

	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	assert.Equal(t, unsafe.SliceData(data), unsafe.SliceData(buffer.data))
}

func TestCOWBufferEqualClone(t *testing.T) {
	t.Skipf("not implement")

	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	clone1 := buffer.Clone()

	assert.Equal(t, unsafe.SliceData(buffer.data), unsafe.SliceData(clone1.data))
}

func TestCOWEqualsClone(t *testing.T) {
	t.Skipf("not implement")

	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	clone1 := buffer.Clone()
	clone2 := buffer.Clone()

	assert.Equal(t, unsafe.SliceData(clone1.data), unsafe.SliceData(clone2.data))
}

func TestCOWNotCopyBufferWhenString(t *testing.T) {
	t.Skipf("not implement")

	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	assert.True(t, (*byte)(unsafe.SliceData(data)) == unsafe.StringData(buffer.String()))
}

func TestCOWNotCopyBufferCloneWhenString(t *testing.T) {
	t.Skipf("not implement")

	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	clone1 := buffer.Clone()
	assert.True(t, (*byte)(unsafe.StringData(buffer.String())) == unsafe.StringData(clone1.String()))
}

func TestCOWBufferCopyingEqualWhenString(t *testing.T) {
	t.Skipf("not implement")

	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	clone1 := buffer.Clone()
	clone2 := buffer.Clone()

	assert.True(t, (*byte)(unsafe.StringData(clone1.String())) == unsafe.StringData(clone2.String()))

}

func TestCOWChangeByteSuccess(t *testing.T) {
	t.Skipf("not implement")

	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	assert.True(t, buffer.Update(0, 'g'))
}

func TestCOWChangeByteInvalidIndex(t *testing.T) {
	t.Skipf("not implement")

	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	assert.False(t, buffer.Update(-1, 'g'))
}

func TestCOWChangeByteOutOfRange(t *testing.T) {
	t.Skipf("not implement")

	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	assert.False(t, buffer.Update(4, 'g'))
}

func TestCOWEqualBufferAfterUpdate(t *testing.T) {
	t.Skipf("not implement")

	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	buffer.Update(0, 'g')

	assert.True(t, reflect.DeepEqual([]byte{'g', 'b', 'c', 'd'}, buffer.data))
}

func TestCOWCopyNotUpdateWhenSrcBufferUpdate(t *testing.T) {
	t.Skipf("not implement")

	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	copy1 := buffer.Clone()
	copy2 := buffer.Clone()

	buffer.Update(0, 'g')

	// равны сами массивы
	assert.True(t, reflect.DeepEqual([]byte{'a', 'b', 'c', 'd'}, copy1.data))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'b', 'c', 'd'}, copy2.data))

	// массивы указывают на один и тот же участок
	assert.Equal(t, unsafe.SliceData(copy1.data), unsafe.SliceData(copy2.data))
}

func TestCOWCopyEqualAfterSrcUpdate(t *testing.T) {
	t.Skipf("not implement")

	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	copy1 := buffer.Clone()

	buffer.Update(0, 'g')

	assert.NotEqual(t, unsafe.SliceData(buffer.data), unsafe.SliceData(copy1.data))
}

func TestCOWNotCopyIfReferOnlyOneObj(t *testing.T) {
	t.Skipf("not implement")

	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	copy1 := buffer.Clone()
	copy2 := buffer.Clone()

	copy1.Close()

	previous := copy2.data
	copy2.Update(0, 'f')
	current := copy2.data

	assert.Equal(t, unsafe.SliceData(previous), unsafe.SliceData(current))

	copy2.Close()

}

func TestCOWBuffer(t *testing.T) {
	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	copy1 := buffer.Clone()
	copy2 := buffer.Clone()

	assert.Equal(t, unsafe.SliceData(data), unsafe.SliceData(buffer.data))
	assert.Equal(t, unsafe.SliceData(buffer.data), unsafe.SliceData(copy1.data))
	assert.Equal(t, unsafe.SliceData(copy1.data), unsafe.SliceData(copy2.data))

	assert.True(t, (*byte)(unsafe.SliceData(data)) == unsafe.StringData(buffer.String()))
	assert.True(t, (*byte)(unsafe.StringData(buffer.String())) == unsafe.StringData(copy1.String()))
	assert.True(t, (*byte)(unsafe.StringData(copy1.String())) == unsafe.StringData(copy2.String()))

	assert.True(t, buffer.Update(0, 'g'))
	assert.False(t, buffer.Update(-1, 'g'))
	assert.False(t, buffer.Update(4, 'g'))

	assert.True(t, reflect.DeepEqual([]byte{'g', 'b', 'c', 'd'}, buffer.data))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'b', 'c', 'd'}, copy1.data))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'b', 'c', 'd'}, copy2.data))

	assert.NotEqual(t, unsafe.SliceData(buffer.data), unsafe.SliceData(copy1.data))
	assert.Equal(t, unsafe.SliceData(copy1.data), unsafe.SliceData(copy2.data))

	copy1.Close()

	previous := copy2.data
	copy2.Update(0, 'f')
	current := copy2.data

	// 1 reference - don't need to copy buffer during update
	assert.Equal(t, unsafe.SliceData(previous), unsafe.SliceData(current))

	copy2.Close()
}
