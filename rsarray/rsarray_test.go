package rsarray_test

import (
	"optimal-rs-array/rsarray"
	"testing"
)

func TestNewRSArray(t *testing.T) {
	arr := rsarray.NewRSArray()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	arr.Grow()
	for i := 0; i < 30; i++ {
		arr.Write(i,i)
	}
	arr.Write(39,39)
	arr.Write(40,39)
}
