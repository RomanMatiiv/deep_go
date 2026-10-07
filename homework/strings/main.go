package main

import (
	"fmt"
	"log/slog"
	"unsafe"
)

// Домашнее задание №3
//
// В домашнем задании нужно реализовать COW (Copy-On-Write) буффер.
//
// Идея подхода copy-on-write заключается в том, что при чтении данных
// используется общая копия данных буфера,
// но в случае изменения данных — создается новая копия данных буфера.
// Для реализации такого подхода можно
// использовать разделяемый счетчик ссылок: если при изменении данных буфера
// кто-то еще ссылается на этот буфер, то нужно будет сначала произвести
// копию данных буфера, изменить счетчик ссылок и только затем произвести
// изменение (если никто не ссылается на буфер, то копировать данные буфера
// не нужно при изменении данных).
//
// Дополнительно еще нужно реализовать метод конвертации данных буфера
// в строку без копирования и дополнительного выделения памяти.

type COW interface {
	Clone() COWBuffer                  // создать новую копию буфера
	Close()                            // перестать использовать копию буффера
	Update(index int, value byte) bool // изменить определенный байт в буффере
	String() string                    // сконвертировать буффер в строку
}

type COWBuffer struct {
	data []byte
	refs *int
	// need to implement
}

func NewCOWBuffer(data []byte) COWBuffer {
	var cntRef int

	return COWBuffer{
		data: data,
		refs: &cntRef,
	}
}

func (b *COWBuffer) Clone() COWBuffer {
	slog.Debug("Clone()")

	*b.refs += 1

	return COWBuffer{
		data: b.data,
		refs: b.refs,
	}
}

func (b *COWBuffer) Close() {
	if *b.refs > 0 {
		*b.refs -= 1
	}

	var refs int
	b.refs = &refs
	b.data = nil
}

func (b *COWBuffer) Update(index int, value byte) bool {
	slog.Debug(fmt.Sprintf("Update(%d, %b)", index, value))

	if !b.isValidUpdateParams(index, value) {
		return false
	}

	if *b.refs == 0 {
		b.data[index] = value
	} else {
		b.updateWithRefs(index, value)
	}

	return true
}

func (b *COWBuffer) isValidUpdateParams(index int, value byte) bool {
	curLen := len(b.data)
	if index < 0 || index >= curLen {
		return false
	}

	if b.refs == nil || *b.refs < 0 {
		panic("invalid ref count")
	}

	return true
}

func (b *COWBuffer) updateWithRefs(index int, value byte) {

	newDataBuffer := make([]byte, len(b.data))
	copy(newDataBuffer, b.data)

	b.data = newDataBuffer
	b.data[index] = value

	*b.refs -= 1

	var refs int
	b.refs = &refs
}

func (b *COWBuffer) String() string {
	if len(b.data) == 0 {
		return ""
	}

	return unsafe.String(unsafe.SliceData(b.data), len(b.data))
}
