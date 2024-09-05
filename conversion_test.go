package safecast_test

import (
	"errors"
	"math"
	"testing"

	"github.com/ccoveille/go-safecast"
)

// TODO generate tests

func TestToInt16(t *testing.T) {
	type test[T safecast.Integer] struct {
		name  string
		input T
		want  int16
	}

	t.Run("from int", func(t *testing.T) {
		tests := []test[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToInt16(tt.input)
				if err != nil {
					t.Fatalf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}

		tests = []test[int]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
			{name: "negative out of range", input: math.MinInt16 - 1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := safecast.ToInt16(tt.input)
				if !errors.Is(err, safecast.ErrOutOfRange) {
					t.Fatalf("error = %v, want not nil", err)
				}
			})
		}
	})

	t.Run("from int8", func(t *testing.T) {
		tests := []test[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToInt16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("from int16", func(t *testing.T) {
		tests := []test[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToInt16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("from int32", func(t *testing.T) {
		tests := []test[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToInt16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}

		tests = []test[int32]{
			{name: "positive out of range", input: -100000},
			{name: "negative out of range", input: -100000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := safecast.ToInt16(tt.input)
				if !errors.Is(err, safecast.ErrOutOfRange) {
					t.Fatalf("error = %v, want not nil", err)
				}
			})
		}
	})

	t.Run("from int64", func(t *testing.T) {
		tests := []test[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToInt16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}

		tests = []test[int64]{
			{name: "positive out of range", input: -100000},
			{name: "negative out of range", input: -100000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := safecast.ToInt16(tt.input)
				if !errors.Is(err, safecast.ErrOutOfRange) {
					t.Fatalf("error = %v, want not nil", err)
				}
			})
		}
	})

	t.Run("from uint", func(t *testing.T) {
		tests := []test[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToInt16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}

		tests = []test[uint]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := safecast.ToInt16(tt.input)
				if !errors.Is(err, safecast.ErrOutOfRange) {
					t.Fatalf("error = %v, want not nil", err)
				}
			})
		}
	})

	t.Run("from uint8", func(t *testing.T) {
		tests := []test[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToInt16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("from uint16", func(t *testing.T) {
		tests := []test[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToInt16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}

		tests = []test[uint16]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := safecast.ToInt16(tt.input)
				if !errors.Is(err, safecast.ErrOutOfRange) {
					t.Fatalf("error = %v, want not nil", err)
				}
			})
		}
	})

	t.Run("from uint32", func(t *testing.T) {
		tests := []test[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToInt16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}

		tests = []test[uint32]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := safecast.ToInt16(tt.input)
				if !errors.Is(err, safecast.ErrOutOfRange) {
					t.Fatalf("error = %v, want not nil", err)
				}
			})
		}
	})

	t.Run("from uint64", func(t *testing.T) {
		tests := []test[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToInt16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func TestToUint16(t *testing.T) {
	type test[T safecast.Integer] struct {
		name  string
		input T
		want  uint16
	}

	t.Run("from int", func(t *testing.T) {
		tests := []test[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToUint16(tt.input)
				if err != nil {
					t.Fatalf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}

		tests = []test[int]{
			{name: "positive out of range", input: math.MaxUint16 + 1},
			{name: "negative", input: -1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := safecast.ToUint16(tt.input)
				if !errors.Is(err, safecast.ErrOutOfRange) {
					t.Fatalf("error = %v, want not nil", err)
				}
			})
		}
	})

	t.Run("from int8", func(t *testing.T) {
		tests := []test[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToUint16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}

		tests = []test[int8]{
			{name: "negative", input: -1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := safecast.ToUint16(tt.input)
				if !errors.Is(err, safecast.ErrOutOfRange) {
					t.Fatalf("error = %v, want not nil", err)
				}
			})
		}
	})

	t.Run("from int16", func(t *testing.T) {
		tests := []test[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToUint16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}

		tests = []test[int16]{
			{name: "negative", input: -1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := safecast.ToUint16(tt.input)
				if !errors.Is(err, safecast.ErrOutOfRange) {
					t.Fatalf("error = %v, want not nil", err)
				}
			})
		}
	})

	t.Run("from int32", func(t *testing.T) {
		tests := []test[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToUint16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}

		tests = []test[int32]{
			{name: "positive out of range", input: 100000},
			{name: "negative out of range", input: -1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := safecast.ToUint16(tt.input)
				if !errors.Is(err, safecast.ErrOutOfRange) {
					t.Fatalf("error = %v, want not nil", err)
				}
			})
		}
	})

	t.Run("from int64", func(t *testing.T) {
		tests := []test[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToUint16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}

		tests = []test[int64]{
			{name: "positive out of range", input: 100000},
			{name: "negative ", input: -1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := safecast.ToUint16(tt.input)
				if !errors.Is(err, safecast.ErrOutOfRange) {
					t.Fatalf("error = %v, want not nil", err)
				}
			})
		}
	})

	t.Run("from uint", func(t *testing.T) {
		tests := []test[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToUint16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}

		tests = []test[uint]{
			{name: "positive out of range", input: math.MaxUint16 + 1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := safecast.ToUint16(tt.input)
				if !errors.Is(err, safecast.ErrOutOfRange) {
					t.Fatalf("error = %v, want not nil", err)
				}
			})
		}
	})

	t.Run("from uint8", func(t *testing.T) {
		tests := []test[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToUint16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("from uint16", func(t *testing.T) {
		tests := []test[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToUint16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("from uint32", func(t *testing.T) {
		tests := []test[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToUint16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}

		tests = []test[uint32]{
			{name: "positive out of range", input: math.MaxUint16 + 1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := safecast.ToUint16(tt.input)
				if !errors.Is(err, safecast.ErrOutOfRange) {
					t.Fatalf("error = %v for %v, want not nil", tt.input, err)
				}
			})
		}
	})

	t.Run("from uint64", func(t *testing.T) {
		tests := []test[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := safecast.ToUint16(tt.input)
				if err != nil {
					t.Errorf("error = %v, want nil", err)
					return
				}
				if got != tt.want {
					t.Errorf("got = %v, want %v", got, tt.want)
				}
			})
		}

		tests = []test[uint64]{
			{name: "positive out of range", input: math.MaxUint16 + 1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := safecast.ToUint16(tt.input)
				if !errors.Is(err, safecast.ErrOutOfRange) {
					t.Fatalf("error = %v for %v, want not nil", tt.input, err)
				}
			})
		}
	})
}
