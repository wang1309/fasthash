package hashtable

import (
	"errors"
)

type FixedMap struct {
	cellCnt uint32
	elemCnt uint32
	cells   []uint64
}

type FixedMapIterator struct {
	table *FixedMap
	idx   uint32
}

func (ht *FixedMap) Init(cellCnt uint32) {
	ht.cellCnt = cellCnt
	ht.cells = make([]uint64, cellCnt)
}

func (ht *FixedMap) Insert(key uint32) uint64 {
	value := ht.cells[key]
	if value == 0 {
		ht.elemCnt++
		value = uint64(ht.elemCnt)
		ht.cells[key] = value
	}
	return value
}

func (ht *FixedMap) Cells() []uint64 {
	return ht.cells
}

func (ht *FixedMap) Cardinality() uint64 {
	return uint64(ht.elemCnt)
}

func (it *FixedMapIterator) Init(ht *FixedMap) {
	it.table = ht
	it.idx = 0
}

func (it *FixedMapIterator) Next() (key uint32, value uint64, err error) {
	for it.idx < it.table.cellCnt && it.table.cells[it.idx] == 0 {
		it.idx++
	}

	if it.idx == it.table.cellCnt {
		err = errors.New("out of range")
		return
	}

	key = it.idx
	value = it.table.cells[key]

	return
}
