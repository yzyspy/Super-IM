package collection

import (
	"fmt"
	"testing"
)

func TestDeduplicateUids(t *testing.T) {
	uids := []uint64{1, 2, 2, 3, 4, 5, 6, 6, 7, 7, 7}
	deduplicatedUids := deduplicateUids(uids)
	if len(deduplicatedUids) != 20 {
		t.Error("deduplicateUids failed")
	}
	fmt.Printf("deduplicatedUids: %v\n", deduplicatedUids)
}
