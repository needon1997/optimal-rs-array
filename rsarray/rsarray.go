package rsarray

import (
	"fmt"
	"math"
)

type RsArray struct {
	b        indexArr
	sbIndex  int
	sbRecord []int
	k        int
	n        int
}

func NewRSArray() *RsArray {
	arr := &RsArray{sbIndex: 0, sbRecord: make([]int, 1), k: 0, n: 0}
	arr.b = indexArr{k: 0, arr: make([]*dataBlock, 1)}
	return arr
}
func (this *RsArray) Write(index int, data int) {
	if index >= this.n {
		panic(fmt.Sprintf("index out of bound: %v", index))
	}
	hi, lo := this.getHiLo(index)
	this.b.arr[hi].data[lo] = data
	return
}
func (this *RsArray) Read(index int) int {
	if index >= this.n {
		panic(fmt.Sprintf("index out of bound: %v", index))
	}
	hi, lo := this.getHiLo(index)
	return this.b.arr[hi].data[lo]
}

func (this *RsArray) getHiLo(index int) (int, int) {
	sbIndex := int(math.Floor(math.Log2(float64(index) + 1)))
	be := (index + 1) - (1 << sbIndex)
	b := be >> int(math.Ceil(float64(sbIndex)/2))
	e := be - (b << int(math.Ceil(float64(sbIndex)/2)))
	var p int
	if sbIndex%2 == 0 {
		p = int(math.Pow(2, float64(sbIndex)/2+1) - 2)
	} else {
		p = int(3*math.Pow(2, float64(sbIndex-1)/2) - 2)
	}
	hi := p + b
	lo := e
	return hi, lo
}
func (this *RsArray) Grow() {
	//last data block is full
	if this.k == 0 || this.b.arr[this.k-1].count == this.b.arr[this.k-1].size {
		if this.sbRecord[this.sbIndex] >= int(math.Pow(2, math.Floor(float64(this.sbIndex)/2))) { //current super block is not full
			if len(this.sbRecord) < this.sbIndex+2 {
				this.sbRecord = append(this.sbRecord, 0)
			}
			this.sbIndex += 1
		}
		this.sbRecord[this.sbIndex] += 1
		this.b.grow()
		this.b.arr[this.k] = newDataBlock(int(math.Pow(2, math.Ceil(float64(this.sbIndex)/2))))
		this.k += 1
	}
	this.b.arr[this.k-1].count += 1
	this.n += 1
}
func (this *RsArray) Shrink() {
	if this.b.arr[this.k-1].count == 1 {
		this.b.shrink()
		this.k -= 1
		this.sbRecord[this.sbIndex] -= 1
		if this.sbRecord[this.sbIndex] == 0 {
			this.sbIndex -= 1
		}
	} else {
		this.b.arr[this.k-1].count -= 1
	}
	this.n -= 1
}

type indexArr struct {
	arr []*dataBlock
	k   int
}

func (this *indexArr) grow() {
	if this.k >= len(this.arr) {
		newArr := make([]*dataBlock, 2*len(this.arr))
		for i := 0; i < this.k; i++ {
			newArr[i] = this.arr[i]
		}
		this.arr = newArr
	}
	this.k += 1
}

func (this *indexArr) shrink() {
	if this.k > 0 {
		this.k -= 1
		this.arr[this.k] = nil
		if float64(this.k) < float64(len(this.arr))/3 {
			newArr := make([]*dataBlock, len(this.arr)/2)
			for i := 0; i < this.k; i++ {
				newArr[i] = this.arr[i]
			}
			this.arr = newArr
		}
	}
}

type dataBlock struct {
	data  []int
	count int
	size  int
}

func newDataBlock(size int) *dataBlock {
	d := &dataBlock{data: make([]int, size), count: 0, size: size}
	return d
}
