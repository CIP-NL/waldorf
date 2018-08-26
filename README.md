[![Go Report Card](https://goreportcard.com/badge/github.com/CIP-NL/waldorf)](https://goreportcard.com/report/github.com/CIP-NL/waldorf)
[![CircleCI](https://circleci.com/gh/CIP-NL/waldorf/tree/master.svg?style=svg)](https://circleci.com/gh/CIP-NL/waldorf/tree/master)

# waldorf

An easy to use library to validate data, for example incomming API requests.
Waldorf checks everything that is wrong and returns a slice of complaints.

```go
package waldorf_test

import (
	"fmt"
	"github.com/CIP-NL/waldorf"
	)

func Example() {
	w := waldorf.Observe()

	for i := 0; i < 3; i++ {
		w.ShouldBeTrue(i==0, "%d is one too many alread!", i)
	}
	complaints := w.Ridicule()
	fmt.Println(complaints.String())
	// Output:
	// [1] is one too many alread!
	// [2] is one too many alread!
}
```
