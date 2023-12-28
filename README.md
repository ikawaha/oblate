Oblate
===

Oblate is a wrapper for errors that conceals error details that need not be shown to the user, without losing any error information.

## Usage

```go
func Example() {
	e1 := errors.New("e1")
	e2 := errors.New("e2")
	err := oblate.New("user-facing error message", e1, e2)

	fmt.Println("error:")
	fmt.Println(err.Error())

	fmt.Println("cause:")
	fmt.Println(err.(*oblate.Error).Cause())

	// Output:
	// error:
	// user-facing error message
	// cause:
	// e1
	// e2
}
```