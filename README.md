# floatcheck

[![Go Report Card](https://goreportcard.com/badge/github.com/yudppp/floatcheck)](https://goreportcard.com/report/github.com/yudppp/floatcheck)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
This is a custom linter that checks for potential issues related to floating-point number precision in Go code. Specifically, it currently detects:

- Usage of `fmt.Sprintf` with floating-point format specifiers (`%.nf`, `%.f`, `%.ng`, `%e`).
- Division operations where at least one of the operands is a floating-point number.
- Equality comparisons (`==`, `!=`) between floating-point numbers.
- Relational comparisons (`<`, `<=`, `>`, `>=`) between floating-point numbers.

**Warning:** This linter provides warnings for potentially problematic patterns. It's up to the user to evaluate if the reported instances are actual issues in their specific context.

## Installation and Usage

To use floatcheck, you need to build it and then run it via `go vet` with the `-vettool` flag.

1.  **Clone this repository:**
    ```bash
    git clone https://github.com/yudppp/floatcheck.git
    cd floatcheck
    ```

2.  **Build the linter:**
    ```bash
    go build -o floatcheck ./cmd/floatcheck
    ```

3.  **Run the linter on your Go project using `go vet`:**
    ```bash
    go vet -vettool=./floatcheck your_project_path/...
    ```

    Replace `your_project_path/...` with the path to the Go packages you want to analyze.

## Integration with `go vet`

This linter is designed to be used with `go vet` via the `-vettool` flag. You need to build floatcheck as an executable and then point `go vet` to it.

## Example Usage

Consider the following Go code:

```go
package main

import "fmt"

func main() {
  var pi float64 = 3.14159
  var result float32 = 1.0 / 3.0

  formatted := fmt.Sprintf("%.2f", pi) // This will trigger a warning
  if result == 0.333 { // This will trigger a warning
    println("Result is approximately one-third")
  }
}
````

To analyze this code, navigate to the root of your project and run:

```bash
go vet -vettool=./floatcheck ./example
```

This will report the potential issues in the `Sprintf` format and the equality comparison.

## Contributing

Contributions to this linter are welcome\! If you have ideas for new checks or improvements, feel free to open an issue or submit a pull request.

Please ensure that your contributions adhere to the Go style guidelines and include appropriate tests.

## License

[The MIT License (MIT)](https://github.com/yudppp/floatcheck/blob/main/LICENSE)
