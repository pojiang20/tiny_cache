package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	c := New()

	chanA := make(chan string, 5)
	chanB := make(chan string, 5)
	chanC := make(chan string, 5)

	go func() {
		for v := range chanA {
			_, _ = c.GetWithFn(context.Background(), v, func() (interface{}, error) {
				return 1, nil
			})
		}
	}()

	go func() {
		for v := range chanB {
			_, _ = c.GetWithFn(context.Background(), v, func() (interface{}, error) {
				return 1, nil
			})
		}
	}()

	go func() {
		for v := range chanC {
			_, _ = c.GetWithFn(context.Background(), v, func() (interface{}, error) {
				return 1, nil
			})
		}
	}()

	for j := 0; j < 100; j++ {
		for i := 0; i < 5; i++ {
			chanA <- fmt.Sprint(j + i)
			chanB <- fmt.Sprint(j + i)
			chanC <- fmt.Sprint(j + i)
		}
	}
}
