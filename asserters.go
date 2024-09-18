package safecast

func checkUpperBoundary[T Type, T2 Type](value T, boundary T2) error {
	if value <= 0 {
		return nil
	}

	var overflow bool
	switch f := any(value).(type) {
	case float64:
		overflow = isFloatOverflow(f, boundary)

	case float32:
		overflow = isFloatOverflow(f, boundary)

	default:
		// for all other integer types, it fits in an uint64 without overflow as we know value is positive.
		overflow = uint64(value) > uint64(boundary)
	}

	if overflow {
		return Error{
			value:    value,
			boundary: boundary,
			err:      ErrExceedMaximumValue,
		}
	}

	return nil
}

func checkLowerBoundary[T Type, T2 Type](value T, boundary T2) error {
	if value >= 0 {
		return nil
	}

	var underflow bool
	switch f := any(value).(type) {
	case float64:
		underflow = isFloatUnderOverflow(f, boundary)
	case float32:
		underflow = isFloatUnderOverflow(f, boundary)
	default:
		// for all other integer types, it fits in an int64 without overflow as we know value is negative.
		underflow = int64(value) < int64(boundary)
	}

	if underflow {
		return Error{
			value:    value,
			boundary: boundary,
			err:      ErrExceedMinimumValue,
		}
	}

	return nil
}

func isFloatOverflow[T Type, T2 Type](value T, boundary T2) bool {
	// boundary is positive when checking for an overflow

	// everything fits in float64 without overflow.
	v := float64(value)
	b := float64(boundary)

	if v > b*1.01 {
		// way greater than the maximum value
		return true
	}

	if v < b*0.99 {
		// we are way below the maximum value
		return false
	}
	// we are close to the maximum value

	// let's try to create the overflow
	// by converting back and forth with type juggling
	conv := float64(T(T2(v)))

	// the number was between 0.99 and 1.01 of the maximum value
	// once converted back and forth, we need to check if the value is in the same range
	// if not, so it's an overflow
	return conv <= b*0.99
}

func isFloatUnderOverflow[T Type, T2 Type](value T, boundary T2) bool {
	// everything fits in float64 without overflow.
	v := float64(value)
	b := float64(boundary)

	if b == 0 {
		// boundary is 0
		// we can check easily
		return value < 0
	}

	if v < b*1.01 { // please note value and boundary are negative here
		// way below than the minimum value, it would underflow
		return true
	}

	if v > b*0.99 { // please note value and boundary are negative here
		// way greater than the minimum value
		return false
	}

	// we are just above to the minimum value
	// let's try to create the underflow
	conv := float64(T(T2(v)))

	// the number was between 0.99 and 1.01 of the minimum value
	// once converted back and forth, we need to check if the value is in the same range
	// if not, so it's an underflow
	return conv >= b*0.99
}
