# floatcheck

[![Go Report Card](https://goreportcard.com/badge/github.com/yudppp/floatcheck)](https://goreportcard.com/report/github.com/yudppp/floatcheck)

floatcheck is a custom linter for Go that detects potential floating-point precision issues in your code.

## What it detects
- Usage of floating-point format specifiers (e.g., `%.nf`, `%.f`) in `fmt.Sprintf`
- Division operations involving floating-point numbers
- Equality comparisons (`==`, `!=`) between floating-point numbers
- Relational comparisons (`<`, `<=`, `>`, `>=`) between floating-point numbers

**Warning:** This linter warns about potential patterns. It is up to the user to determine if the warning is an actual issue in their context.

## Installation & Usage

### 1. To check for all patterns at once
```bash
go install github.com/yudppp/floatcheck/cmd/floatcheck
```
```bash
go vet -vettool=$(which floatcheck) your_project_path/...
```

### 2. To use each check individually
- Format only: `floatcheck_format`
- Division only: `floatcheck_division`
- Comparison only: `floatcheck_comparison`

```bash
go install github.com/yudppp/floatcheck/cmd/floatcheck_format
# or
go install github.com/yudppp/floatcheck/cmd/floatcheck_division
# or
go install github.com/yudppp/floatcheck/cmd/floatcheck_comparison
```

## Testing

```bash
go test ./...
```

Or you can manually check the samples under the `testdata` directory.

## Example

```go
package main
import "fmt"
func main() {
  var pi float64 = 3.14159
  var result float32 = 1.0 / 3.0
  formatted := fmt.Sprintf("%.2f", pi) // warning
  if result == 0.333 { // warning
    println("Result is approximately one-third")
  }
}
```

## Contributing
Pull requests and issues are welcome! Please follow the Go style guide and add tests for new features or bug fixes.

## License
[MIT License](https://github.com/yudppp/floatcheck/blob/main/LICENSE)
