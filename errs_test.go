package errs

import (
	"sync"
	"testing"
)

func Test_errors_String(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		errors  []error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errors{
				errors: tt.fields.errors,
			}
			if got := e.String(); got != tt.want {
				t.Errorf("errors.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
