package main

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
	var cntRef *int

	return COWBuffer{
		data: data,
		refs: cntRef,
	}
}

func (b *COWBuffer) Clone() COWBuffer {
	return COWBuffer{} // todo implement
}

func (b *COWBuffer) Close() {
	// todo implement
}

func (b *COWBuffer) Update(index int, value byte) bool {
	return false // todo implement
}

func (b *COWBuffer) String() string {
	return "" // todo implement
}
