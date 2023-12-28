Oblate
===

Oblate serves as a wrapper for errors, concealing details unnecessary for user presentation while simultaneously preserving comprehensive error information.

## Usage

```go
func Example() {
    e1 := errors.New("e1")
    e2 := errors.New("e2")
    e3 := fmt.Errorf("wrapped error: %w", io.EOF)

    // New returns an error that wraps the given errors.
    err := oblate.New("user-facing error message", e1, e2, e3)

    // The error formats as the strings obtained by calling the Error
    // method of the first error in errs.
    fmt.Println("error:")
    fmt.Println(err.Error())

    // Cause returns the error formats as the concatenation of the strings
    // obtained by calling the Error method of each element of errs
    // except the first error, with a newline between each string.
    fmt.Println("cause:")
    fmt.Println(err.(*oblate.Error).Cause())

    // The error wrapped by oblate can be checked with errors.Is and errors.As.
    fmt.Println("error details:")
    fmt.Printf("errors.Is(err, io.EOF): %v\n", errors.Is(err, io.EOF))

    // Output:
    // error:
    // user-facing error message
    // cause:
    // e1
    // e2
    // EOF
    // error details:
    // errors.Is(err, io.EOF): true
}
```