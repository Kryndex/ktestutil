package utils

import "time"

// Retry retries f until f return nil error.
// It makes max attempts and adds delay between each attempt.
func Retry(attempts int, delay time.Duration, f func() error) error {
	var err error

	for i := 0; i < attempts; i++ {
		err = f()
		if err == nil {
			break
		}

		if i < attempts-1 {
			time.Sleep(delay)
		}
	}

	return err
}
