package safecast_test

import (
	"math"
	"testing"

	"github.com/ccoveille/go-safecast"
)

func TestConvert(t *testing.T) {
	t.Run("to int8", func(t *testing.T) {
		t.Run("from int", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int
				want  int8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int{
				"positive out of range": math.MaxInt8 + 1,
				"negative out of range": math.MinInt8 - 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int8
				want  int8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from int16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int16
				want  int8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int16{
				"positive out of range": math.MaxInt8 + 1,
				"negative out of range": math.MinInt8 - 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int32
				want  int8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int32{
				"positive out of range": math.MaxInt8 + 1,
				"negative out of range": math.MinInt8 - 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int64
				want  int8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int64{
				"positive out of range": math.MaxInt8 + 1,
				"negative out of range": math.MinInt8 - 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint
				want  int8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint{
				"positive out of range": math.MaxInt8 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint8
				want  int8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint8{
				"positive out of range": math.MaxInt8 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint16
				want  int8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint16{
				"positive out of range": math.MaxInt8 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint32
				want  int8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint16{
				"positive out of range": math.MaxInt8 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint64
				want  int8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint64{
				"positive out of range": math.MaxInt8 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float32
				want  int8
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float32{
				"positive out of range": math.MaxInt8 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float64
				want  int8
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float64{
				"positive out of range": math.MaxInt8 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})
	})

	t.Run("to int16", func(t *testing.T) {
		t.Run("from int", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int
				want  int16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int{
				"positive out of range": math.MaxInt16 + 1,
				"negative out of range": math.MinInt16 - 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int8
				want  int16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from int16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int16
				want  int16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from int32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int32
				want  int16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int32{
				"positive out of range": math.MaxInt16 + 1,
				"negative out of range": math.MinInt16 - 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int64
				want  int16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int64{
				"positive out of range": math.MaxInt16 + 1,
				"negative out of range": math.MinInt16 - 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint
				want  int16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint{
				"positive out of range": math.MaxInt16 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint8
				want  int16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint16
				want  int16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint16{
				"positive out of range": math.MaxInt16 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint32
				want  int16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint32{
				"positive out of range": math.MaxInt16 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint64
				want  int16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint64{
				"positive out of range": math.MaxInt16 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float32
				want  int16
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float32{
				"positive out of range": math.MaxInt16 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float64
				want  int16
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float64{
				"positive out of range": math.MaxInt16 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})
	})

	t.Run("to int32", func(t *testing.T) {
		t.Run("from int", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int
				want  int32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int{
				"positive out of range": math.MaxInt32 + 1,
				"negative out of range": math.MinInt32 - 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int32](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int8
				want  int32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from int16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int16
				want  int32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from int32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int32
				want  int32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from int64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int64
				want  int32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int64{
				"positive out of range": math.MaxInt32 + 1,
				"negative out of range": math.MinInt32 - 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int32](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint
				want  int32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint{
				"positive out of range": math.MaxInt32 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int32](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint8
				want  int32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint16
				want  int32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint32
				want  int32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint32{
				"positive out of range": math.MaxInt32 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int32](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint64
				want  int32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint64{
				"positive out of range": math.MaxInt32 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int32](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float32
				want  int32
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float32{
				"positive out of range": math.MaxInt32 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int32](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float64
				want  int32
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float64{
				"positive out of range": math.MaxInt32 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int32](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})
	})

	t.Run("to int64", func(t *testing.T) {
		t.Run("from int", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int
				want  int64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from int8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int8
				want  int64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from int16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int16
				want  int64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from int32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int32
				want  int64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from int64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int64
				want  int64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint
				want  int64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint8
				want  int64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint16
				want  int64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint32
				want  int64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint64
				want  int64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint64{
				"positive out of range": math.MaxUint64,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int64](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float32
				want  int64
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float32{
				"positive out of range": math.MaxUint64,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int64](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float64
				want  int64
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float64{
				"positive out of range": math.MaxUint64,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int64](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})
	})

	t.Run("to int", func(t *testing.T) {
		t.Run("from int", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int
				want  int
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from int8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int8
				want  int
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from int16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int16
				want  int
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from int32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int32
				want  int
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from int64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int64
				want  int
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
				"negative within range": {input: -100, want: -100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint
				want  int
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint{
				"positive out of range": math.MaxUint,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint8
				want  int
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint16
				want  int
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint32
				want  int
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint64
				want  int
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint64{
				"positive out of range": math.MaxUint,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float32
				want  int
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float32{
				"positive out of range": math.MaxUint,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float64
				want  int
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[int](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float64{
				"positive out of range": math.MaxUint,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[int](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})
	})

	t.Run("to uint8", func(t *testing.T) {
		t.Run("from int", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int
				want  uint8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int{
				"positive out of range": math.MaxUint8 + 1,
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int8
				want  uint8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int{
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int16
				want  uint8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int16{
				"positive out of range": math.MaxUint8 + 1,
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int32
				want  uint8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int32{
				"positive out of range": math.MaxUint8 + 1,
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int64
				want  uint8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int64{
				"positive out of range": math.MaxUint8 + 1,
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint
				want  uint8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint{
				"positive out of range": math.MaxUint8 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint8
				want  uint8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint16
				want  uint8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint16{
				"positive out of range": math.MaxUint8 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint32
				want  uint8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint16{
				"positive out of range": math.MaxUint8 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint64
				want  uint8
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint64{
				"positive out of range": math.MaxUint8 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float32
				want  uint8
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float32{
				"positive out of range": math.MaxUint8 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float64
				want  uint8
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint8](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float64{
				"positive out of range": math.MaxUint8 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint8](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})
	})

	t.Run("to uint16", func(t *testing.T) {
		t.Run("from int", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int
				want  uint16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int{
				"positive out of range": math.MaxUint16 + 1,
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int8
				want  uint16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int8{
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int16
				want  uint16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int16{
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int32
				want  uint16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int32{
				"positive out of range": math.MaxUint16 + 1,
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int64
				want  uint16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int64{
				"positive out of range": math.MaxUint16 + 1,
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint
				want  uint16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint{
				"positive out of range": math.MaxUint16 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint8
				want  uint16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint16
				want  uint16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint32
				want  uint16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint32{
				"positive out of range": math.MaxUint16 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint64
				want  uint16
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint64{
				"positive out of range": math.MaxUint16 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float32
				want  uint16
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float32{
				"positive out of range": math.MaxUint16 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float64
				want  uint16
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint16](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float64{
				"positive out of range": math.MaxUint16 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint16](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})
	})

	t.Run("to uint32", func(t *testing.T) {
		t.Run("from int", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int
				want  uint32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int{
				"positive out of range": math.MaxUint32 + 1,
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint32](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int8
				want  uint32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int8{
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint32](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int16
				want  uint32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int8{
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint32](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int32
				want  uint32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int8{
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint32](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int64
				want  uint32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int64{
				"positive out of range": math.MaxUint32 + 1,
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint32](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint
				want  uint32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint{
				"positive out of range": math.MaxUint32 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint32](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint8
				want  uint32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint16
				want  uint32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint32
				want  uint32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint64
				want  uint32
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]uint64{
				"positive out of range": math.MaxUint32 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint32](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float32
				want  uint32
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float32{
				"positive out of range": math.MaxUint32 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint32](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float64
				want  uint32
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint32](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float64{
				"positive out of range": math.MaxUint32 + 1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint32](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})
	})

	t.Run("to uint64", func(t *testing.T) {
		t.Run("from int", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int
				want  uint64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int{
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint64](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int8
				want  uint64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int8{
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint64](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int16
				want  uint64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int16{
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint64](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int32
				want  uint64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int32{
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint64](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int64
				want  uint64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
			for name, input := range map[string]int64{
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint64](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint
				want  uint64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint8
				want  uint64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint16
				want  uint64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint32
				want  uint64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint64
				want  uint64
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from float32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float32
				want  uint64
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float32{
				"positive out of range": math.MaxUint64,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint64](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float64
				want  uint64
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint64](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float64{
				"positive out of range": math.MaxUint64,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint64](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})
	})

	t.Run("to uint", func(t *testing.T) {
		t.Run("from int", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int
				want  uint
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int{
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int8
				want  uint
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int{
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int16
				want  uint
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int{
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int32
				want  uint
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int{
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from int64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input int64
				want  uint
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]int{
				"negative out of range": -1,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from uint", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint
				want  uint
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint8", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint8
				want  uint
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint16", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint16
				want  uint
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint32
				want  uint
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from uint64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input uint64
				want  uint
			}{
				"zero":                  {input: 0, want: 0},
				"positive within range": {input: 100, want: 100},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}
		})

		t.Run("from float32", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float32
				want  uint
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float32{
				"positive out of range": math.MaxUint,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})

		t.Run("from float64", func(t *testing.T) {
			for name, tt := range map[string]struct {
				input float64
				want  uint
			}{
				"zero":                  {input: 0.0, want: 0},
				"positive within range": {input: 1.1, want: 1},
			} {
				t.Run(name, func(t *testing.T) {
					got, err := safecast.Convert[uint](tt.input)
					assertNoError(t, err)
					assertEqual(t, tt.want, got)
				})
			}

			for name, input := range map[string]float64{
				"positive out of range": math.MaxUint,
			} {
				t.Run(name, func(t *testing.T) {
					_, err := safecast.Convert[uint](input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
				})
			}
		})
	})
}
