package waldorf_test

import (
	"fmt"
	"github.com/CIP-NL/waldorf"
)

func Example() {
	w := waldorf.Observe()

	for i := 0; i < 3; i++ {
		w.ShouldBeTrue(i == 0, "%d is one too many alread!", i)
	}
	complaints := w.Ridicule()
	fmt.Println(complaints.String())
	// Output:
	// [1] is one too many alread!
	// [2] is one too many alread!
}
