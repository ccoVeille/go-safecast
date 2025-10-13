package safecast

import (
	"math"
	"strconv"
	"testing"
)

const (
	intSize64bits = uint64(8) // size of int on 64-bit systems
	intSize32bits = uint64(4) // size of int on 32-bit systems
)

func Test_sizeOf(t *testing.T) {

	var intSize = intSize64bits
	if strconv.IntSize == 32 {
		intSize = intSize32bits
	}

	for _, tc := range []struct {
		name         string
		fn           func() uint64
		expectedSize uint64
	}{
		{name: "float32", fn: sizeOf[float32], expectedSize: 4},
		{name: "float64", fn: sizeOf[float64], expectedSize: 8},
		{name: "int", fn: sizeOf[int], expectedSize: intSize},
		{name: "int8", fn: sizeOf[int8], expectedSize: 1},
		{name: "int16", fn: sizeOf[int16], expectedSize: 2},
		{name: "int32", fn: sizeOf[int32], expectedSize: 4},
		{name: "int64", fn: sizeOf[int64], expectedSize: 8},
		{name: "uint", fn: sizeOf[uint], expectedSize: intSize},
		{name: "uint8", fn: sizeOf[uint8], expectedSize: 1},
		{name: "uint16", fn: sizeOf[uint16], expectedSize: 2},
		{name: "uint32", fn: sizeOf[uint32], expectedSize: 4},
		{name: "uint64", fn: sizeOf[uint64], expectedSize: 8},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.fn(); got != tc.expectedSize {
				t.Errorf("sizeOf() = %v, want %v", got, tc.expectedSize)
			}
		})
	}
}

func Test_minOf(t *testing.T) {
	for _, tc := range []struct {
		name        string
		fn          func() any
		expectedMin any
	}{
		{name: "float32", fn: minOf[float32], expectedMin: float32(-math.MaxFloat32)},
		{name: "float64", fn: minOf[float64], expectedMin: float64(-math.MaxFloat64)},
		{name: "int", fn: minOf[int], expectedMin: int(math.MinInt)},
		{name: "int8", fn: minOf[int8], expectedMin: int8(math.MinInt8)},
		{name: "int16", fn: minOf[int16], expectedMin: int16(math.MinInt16)},
		{name: "int32", fn: minOf[int32], expectedMin: int32(math.MinInt32)},
		{name: "int64", fn: minOf[int64], expectedMin: int64(math.MinInt64)},
		{name: "uint", fn: minOf[uint], expectedMin: uint(0)},
		{name: "uint8", fn: minOf[uint8], expectedMin: uint8(0)},
		{name: "uint16", fn: minOf[uint16], expectedMin: uint16(0)},
		{name: "uint32", fn: minOf[uint32], expectedMin: uint32(0)},
		{name: "uint64", fn: minOf[uint64], expectedMin: uint64(0)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.fn(); got != tc.expectedMin {
				t.Errorf("minOf() = %v (%[1]T), want %v (%[2]T)", got, tc.expectedMin)
			}
		})
	}
}

func Test_maxOf(t *testing.T) {
	for _, tc := range []struct {
		name        string
		fn          func() any
		expectedMax any
	}{
		{name: "float32", fn: maxOf[float32], expectedMax: float32(math.MaxFloat32)},
		{name: "float64", fn: maxOf[float64], expectedMax: float64(math.MaxFloat64)},
		{name: "int", fn: maxOf[int], expectedMax: int(math.MaxInt)},
		{name: "int8", fn: maxOf[int8], expectedMax: int8(math.MaxInt8)},
		{name: "int16", fn: maxOf[int16], expectedMax: int16(math.MaxInt16)},
		{name: "int32", fn: maxOf[int32], expectedMax: int32(math.MaxInt32)},
		{name: "int64", fn: maxOf[int64], expectedMax: int64(math.MaxInt64)},
		{name: "uint", fn: maxOf[uint], expectedMax: uint(math.MaxUint)},
		{name: "uint8", fn: maxOf[uint8], expectedMax: uint8(math.MaxUint8)},
		{name: "uint16", fn: maxOf[uint16], expectedMax: uint16(math.MaxUint16)},
		{name: "uint32", fn: maxOf[uint32], expectedMax: uint32(math.MaxUint32)},
		{name: "uint64", fn: maxOf[uint64], expectedMax: uint64(math.MaxUint64)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.fn(); got != tc.expectedMax {
				t.Errorf("maxOf() = %v (%[1]T), want %v (%[2]T)", got, tc.expectedMax)
			}
		})
	}
}
