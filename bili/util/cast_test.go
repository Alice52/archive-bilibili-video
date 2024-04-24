package util

import (
	"fmt"
	"testing"
)

func TestToPrt(t *testing.T) {
	fmt.Println(ToPrt(1))
	fmt.Println(ToPrt("1"))
	fmt.Println(ToPrt(true))
	fmt.Println(ToPrt([]int{1, 2, 3}))
	fmt.Println(ToPrt(map[string]int{"a": 1, "b": 2}))
	fmt.Println(ToPrt(func() {}))
	fmt.Println(ToPrt(struct{}{}))
}
