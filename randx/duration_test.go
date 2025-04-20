package tools

import (
	"testing"
	"time"
)

func TestRandDuration_AroundDuration(t *testing.T) {

	type args struct {
		deviation float64
		base      time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				base:      1 * time.Second,
				deviation: 0.5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRandDuration(tt.args.deviation)
			got := r.AroundDuration(tt.args.base)
			t.Logf("AroundDuration: %v", got)
		})
	}
}

func TestRandDuration_JitterDuration(t *testing.T) {

	type args struct {
		deviation float64
		base      time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				base:      1 * time.Second,
				deviation: 0.5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRandDuration(tt.args.deviation)
			got := r.JitterDuration(tt.args.base)
			t.Logf("JitterDuration: %v", got)
		})
	}
}
