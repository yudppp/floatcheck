# floatcheck

[![Go Report Card](https://goreportcard.com/badge/github.com/yudppp/floatcheck)](https://goreportcard.com/report/github.com/yudppp/floatcheck)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
This is a custom linter for [golangci-lint](https://golangci-lint.run/) that checks for potential issues related to floating-point number precision in Go code. Specifically, it currently detects:

- Usage of `fmt.Sprintf` with floating-point format specifiers (`%.nf`, `%.f`, `%.ng`, `%e`).
- Division operations where at least one of the operands is a floating-point number.
- Equality comparisons (`==`, `!=`) between floating-point numbers.
- Relational comparisons (`<`, `<=`, `>`, `>=`) between floating-point numbers.

**Warning:** This linter provides warnings for potentially problematic patterns. It's up to the user to evaluate if the reported instances are actual issues in their specific context.

## Installation

1.  **Install `golangci-lint`** if you haven't already:
    ```bash
    go install [github.com/golangci/golangci-lint/cmd/golangci-lint@latest](https://github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
    ```

2.  **Clone this repository:**
    ```bash
    git clone [https://github.com/yudppp/floatcheck.git](https://github.com/yudppp/floatcheck.git)
    cd floatcheck
    ```

3.  **Build the custom linter plugin:**
    ```bash
    go build -buildmode=plugin -o floatcheck.so ./cmd/floatcheck
    ```

## Configuration

To enable this custom linter in your project, you need to configure your `.golangci.yml` file. Add the following to your `linters` and `linters-settings` sections:

```yaml
linters:
  enable:
    - floatcheck # Enable the custom linter

linters-settings:
  custom:
    floatcheck:
      path: /path/to/your/floatcheck/floatcheck.so # Replace with the actual path to floatcheck.so
      description: "Checks for potentially problematic float formatting and operations"
      original-url: "[https://github.com/yudppp/floatcheck](https://github.com/yudppp/floatcheck)" 
````

**Important:** Make sure to replace `/path/to/your/floatcheck/floatcheck.so` with the actual absolute or relative path to the built `floatcheck.so` file in your system. Also, update the `original-url` to your repository's URL.

## Usage

Once configured, you can run `golangci-lint` as you normally would:

```bash
golangci-lint run
```

The linter will then analyze your Go code and report any findings related to floating-point formatting, division, and equality comparisons based on the rules implemented in `floatcheck.go`.

## Examples

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
```

Running `golangci-lint run` with `floatcheck` enabled will report the potential issues in the `Sprintf` format and the equality comparison.

## Contributing

Contributions to this linter are welcome\! If you have ideas for new checks or improvements, feel free to open an issue or submit a pull request.

Please ensure that your contributions adhere to the Go style guidelines and include appropriate tests.

## License

[The MIT License (MIT)](https://github.com/yudppp/pprof-mcp-agent/blob/main/LICENSE)

