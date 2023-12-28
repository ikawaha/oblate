Oblate
===

Oblate is a wrapper for errors that conceals error details that need not be shown to the user, without losing any error information.

## Usage

```go
func Example() {
e1 := errors.New("e1")
e2 := errors.New("e2")
e3 := io.EOF
err := oblate.New("user-facing error message", e1, e2, e3)

fmt.Println("error:")
fmt.Println(err.Error())

fmt.Println("cause:")
fmt.Println(err.(*oblate.Error).Cause())

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