package common

import (
	"testing"
)

func TestX(t *testing.T) {

	s := NewState("one", 1, "two", 2)
	s.Set("something", "yes")

}
