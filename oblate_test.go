package oblate_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ikawaha/oblate"
)

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

func TestNew(t *testing.T) {
	t.Run("w/ no error", func(t *testing.T) {
		err := oblate.New("test", nil)
		if err == nil {
			t.Error("expected error")
		}
		if err.Error() != "test" {
			t.Errorf("expected %q, got %q", "test", err.Error())
		}
		if err.(*oblate.Error).Cause() != "" {
			t.Errorf("expected %q, got %q", "", err.(*oblate.Error).Cause())
		}
	})
	t.Run("w/ some error", func(t *testing.T) {
		e1 := errors.New("e1")
		e2 := errors.New("e2")
		err := oblate.New("test", e1, e2)
		if err == nil {
			t.Error("expected error")
		}
		if err.Error() != "test" {
			t.Errorf("expected %q, got %q", "test", err.Error())
		}
		if err.(*oblate.Error).Cause() != "e1\ne2" {
			t.Errorf("expected %q, got %q", "e1\ne2", err.(*oblate.Error).Cause())
		}
	})
}

func TestJoin(t *testing.T) {
	e1 := errors.New("e1")
	e2 := errors.New("e2")
	tests := []struct {
		name string
		errs []error
		want error
	}{
		{name: "empty", errs: nil, want: nil},
		{name: "nil errors", errs: []error{nil, nil, nil}, want: nil},
		{name: "one error", errs: []error{e1}, want: e1},
		{name: "some errors w/ nil", errs: []error{e1, nil, e2}, want: e1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := oblate.Join(tt.errs...)
			if tt.want == nil {
				if got != nil {
					t.Errorf("Join() error = %v, want %v", got, tt.want)
				}
				return
			}
			if got == nil {
				t.Errorf("Join() error = %v, want %v", got, tt.want)
				return
			}
			if got.Error() != tt.want.Error() {
				t.Errorf("Join() error = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Cause(t *testing.T) {
	e1 := errors.New("e1")
	e2 := errors.New("e2")
	e3 := errors.New("e3")
	tests := []struct {
		name string
		errs []error
		want string
	}{
		{name: "one error", errs: []error{e1}, want: ""},
		{name: "some errors", errs: []error{e1, e2, e3}, want: "e2\ne3"},
		{name: "some errors w/ nil", errs: []error{e1, nil, e2, nil, e3}, want: "e2\ne3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := oblate.Join(tt.errs...)
			var got *oblate.Error
			if !errors.As(e, &got) {
				t.Errorf("Join() error = %v, want %v", got, tt.want)
				return
			}
			if v := got.Cause(); v != tt.want {
				t.Errorf("Cause() = %v, want %v", v, tt.want)
			}
		})
	}
}
