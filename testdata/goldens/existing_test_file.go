package testdata

import (
	"reflect"
	"testing"
)

func TestBarBar100(t *testing.T) {
	tests := []struct {
		name    string
		b       *Bar
		i       interface{}
		wantErr bool
	}{
		{
			name:    "Basic test",
			b:       &Bar{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		if err := tt.b.Bar100(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar100() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestBaz100(t *testing.T) {
	tests := []struct {
		name string
		f    *float64
		want float64
	}{
		{
			name: "Basic test",
			f:    func() *float64 { var x float64 = 64; return &x }(),
			want: 64,
		},
	}
	// TestBaz100 contains a comment.
	for _, tt := range tests {
		if got := baz100(tt.f); got != tt.want {
			t.Errorf("%q. baz100() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFoo100(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		strs []string
		// Expected results.
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo100(tt.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo100() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo100() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
