package main

import (
	"reflect"
	"testing"
)

func Test_findOverlappingEvents(t *testing.T) {
	type args struct {
		events []Event
	}
	tests := []struct {
		name string
		args args
		want [][]Event
	}{
		{
			name: "no overlapping events",
			args: args{
				events: []Event{
					{Start: 1, End: 2},
					{Start: 3, End: 4},
				},
			},
			want: nil,
		},
		{
			name: "overlapping events",
			args: args{
				events: []Event{
					{Start: 1, End: 3},
					{Start: 2, End: 4},
				},
			},
			want: [][]Event{
				{
					{Start: 1, End: 3},
					{Start: 2, End: 4},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOverlappingEvents(tt.args.events); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOverlappingEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}
