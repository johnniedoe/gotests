package testdata

import "testing"

func TestBookOpen(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Expected results.
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Book{}
		if err := b.Open(); (err != nil) != tt.wantErr {
			t.Errorf("%q. Book.Open() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestDoorOpen(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Expected results.
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		d := &door{}
		if err := d.Open(); (err != nil) != tt.wantErr {
			t.Errorf("%q. door.Open() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestXmlOpen(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Expected results.
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		x := &xml{}
		if err := x.Open(); (err != nil) != tt.wantErr {
			t.Errorf("%q. xml.Open() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
