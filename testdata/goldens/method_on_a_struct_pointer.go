package testdata

import "testing"

func TestBarFoo7(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		i int
		// Expected results.
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		got, err := b.Foo7(tt.i)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.Foo7() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Bar.Foo7() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
