package db

import (
	"reflect"
	"testing"
)

func TestPurgeNeant(t *testing.T) {

	tests := []struct {
		name string
		in   []string
		want []string
	}{
		{
			name: "empty",
			in:   []string{""},
			want: []string{},
		},
		{
			name: "teempty2",
			in:   []string{"", "Néant"},
			want: []string{},
		},
		{
			name: "ab",
			in:   []string{"", "a", "Néant", "b"},
			want: []string{"a", "b"},
		},
		{
			name: "c",
			in:   []string{"", "c", "Néant"},
			want: []string{"c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PurgeNeant(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PurgeNeant() = %v, want %v", got, tt.want)
			}
		})
	}
}
