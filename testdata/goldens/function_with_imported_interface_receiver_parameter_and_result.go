package testdata

import (
	"io"
	"reflect"
	"testing"
)

func TestFoo17(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		r io.Reader
		// Expected results.
		want io.Reader
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo17(tt.r); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo17() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
