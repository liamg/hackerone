# hackerone

HackerOne API Client in Go.

Current supports the entire _Hackers_ API. Support for the _Customers_ API may be added in future if there is any demand.

Please see the [documentation](https://pkg.go.dev/github.com/liamg/hackerone) for further information.

## Usage Example

```go
package main

import (
	"fmt"

	"github.com/liamg/hackerone"
)

func main() {
	h1 := hackerone.New("your-username-here", "your-api-key-here")
	reports, _, _ := h1.Hackers.GetReports(nil)
	for _, report := range reports {
		fmt.Println(report.Id)
	}
}

```
