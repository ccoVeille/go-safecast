package safecast

import "fmt"

func assertNotNegative[T Type](i T) error {
	if i < 0 {
		return fmt.Errorf("%w: %v is negative", ErrOutOfRange, i)
	}
	return nil
}
