package safecast

func checkUpperBoundary[T Type, T2 Type](value T, boundary T2) error {
	if value <= 0 {
		return nil
	}

	var greater bool
	switch f := any(value).(type) {
	case float64:
		// for float64, everything fits in float64 without overflow.
		// We are using a greater or equal because float cannot be compared easily because of precision loss.
		greater = f >= float64(boundary)
	case float32:
		// everything fits in float32, except float64 greater than math.MaxFloat32.
		// So, we must convert to float64 and check.
		// We are using a greater or equal because float cannot be compared easily because of precision loss.
		greater = float64(f) >= float64(boundary)
	default:
		// for all other integer types, it fits in an uint64 without overflow as we know value is positive.
		greater = uint64(value) > uint64(boundary)
	}

	if greater {
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

	var smaller bool
	switch f := any(value).(type) {
	case float64:
		// everything fits in float64 without overflow.
		// We are using a lower or equal because float cannot be compared easily because of precision loss.
		smaller = f <= float64(boundary)
	case float32:
		// everything fits in float32, except float64 smaller than -math.MaxFloat32.
		// So, we must convert to float64 and check.
		// We are using a lower or equal because float cannot be compared easily because of precision loss.
		smaller = float64(f) <= float64(boundary)
	default:
		// for all other integer types, it fits in an int64 without overflow as we know value is negative.
		smaller = int64(value) < int64(boundary)
	}

	if smaller {
		return Error{
			value:    value,
			boundary: boundary,
			err:      ErrExceedMinimumValue,
		}
	}

	return nil
}
