package main

import (
	"fmt"
	"unsafe"
)

// В домашнем задании нужно реализовать функцию по конвертации числа
// из прямого порядка следования байт (*Big Endian*)
//в обратный порядок следования байт (*Little Endian*).**
//
// Например, число `0x01020304` при вызове функции должно быть
//сконвертировано в число `0x04030201`,
//а число `0x0000FFFF` в `0xFFFF0000`.
//
// Когда число больше одного байта (например, uint32 — 4 байта), его нужно разложить по памяти в несколько ячеек подряд.
// Endianness — это правило: какой байт числа кладётся по младшему адресу (в «начало»).
// Рассмотрим на примере числа в 32 бита 0x12345678

//Little endian (LE)Когда число больше одного байта (например, uint32 — 4 байта), его нужно разложить по памяти в несколько ячеек подряд. Endianness — это правило: какой байт числа кладётся по младшему адресу (в «начало»).Когда число больше одного байта (например, uint32 — 4 байта), его нужно разложить по памяти в несколько ячеек подряд. Endianness — это правило: какой байт числа кладётся по младшему адресу (в «начало»).
//Младший байт — по младшему адресу.
//Память (адреса растут вправо):
//Адрес:     0x00   0x01   0x02   0x03
//┌──────┬──────┬──────┬──────┐
//│ 0x78 │ 0x56 │ 0x34 │ 0x12 │
//└──────┴──────┴──────┴──────┘
//LSB              MSB
//Читается «с конца»: в памяти байты выглядят как 78 56 34 12, хотя число логически — 0x12345678.
//Так работают x86, x86-64, ARM в типичном Linux/Android, почти все Windows на ПaК.
//
//Big endian (BE)
//Старший байт — по младшему адресу.
//Адрес:     0x00   0x01   0x02   0x03
//┌──────┬──────┬──────┬──────┐
//│ 0x12 │ 0x34 │ 0x56 │ 0x78 │
//└──────┴──────┴──────┴──────┘
//MSB              LSB
// Поподробнее про порядки следования байт можно прочитать [здесь](https://betterexplained.com/articles/understanding-big-and-little-endian-byte-order/).

func ToLittleEndian(numberInBigEndian uint32) uint32 {
	numberInLittleEndian := new(uint32)
	ptrLittleEnd := unsafe.Pointer(numberInLittleEndian)

	ptrBigEnd := unsafe.Pointer(&numberInBigEndian)

	for i := 0; i < 4; i++ {
		curByteBigEndian := *(*uint8)(unsafe.Add(ptrBigEnd, i))

		curPtrLittleEndian := (*uint8)(unsafe.Add(ptrLittleEnd, 3-i))
		*curPtrLittleEndian = curByteBigEndian
	}

	return *numberInLittleEndian
}

func main() {
	var number uint32 = 0x12345678
	littleEndian := ToLittleEndian(number)
	fmt.Printf("%x \n", littleEndian)
}
