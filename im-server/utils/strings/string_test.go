package strings

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	suffix := AddShortRandomSuffix("a.png")
	fmt.Printf("%s", suffix)
}
