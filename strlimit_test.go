// Package strlimit provides functions to escape reserved characters in Unix and Windows.
package strlimit

import (
	"fmt"
	"testing"
)

func TestLimitBytes(t *testing.T) {
	type args struct {
		s     string
		limit int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				s:     "Hello, 世界", // "Hello, " is 7 bytes and "世" "界" are 3 bytes.
				limit: 10,
			},
			want: "Hello, 世",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LimitBytes(tt.args.s, tt.args.limit); got != tt.want {
				t.Errorf("LimitBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLimitBytesWithEnd(t *testing.T) {
	type args struct {
		s     string
		limit int
		end   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				s:     "Hello, 世界", // "Hello, " is 7 bytes and "世" "界" are 3 bytes.
				limit: 13,
				end:   "...",
			},
			want: "Hello, 世...",
		},
		{
			name: "rune-safe",
			args: args{
				s:     "Hello, 世界",
				limit: 15,
				end:   "...",
			},
			want: "Hello, 世...",
		},
		{
			name: "surplus bytes",
			args: args{
				s:     "Hello, 世界",
				limit: 20,
				end:   "...",
			},
			want: "Hello, 世界",
		},
		{
			name: "without end",
			args: args{
				s:     "Hello, 世界",
				limit: 10,
				end:   "",
			},
			want: "Hello, 世",
		},
		{
			name: "only end",
			args: args{
				s:     "Hello, 世界",
				limit: 3,
				end:   "...",
			},
			want: "...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LimitBytesWithEnd(tt.args.s, tt.args.limit, tt.args.end); got != tt.want {
				t.Errorf("LimitBytesWithEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleLimitBytesWithEnd() {
	// "Hello, " is 7 bytes and "世" "界" are 3 bytes.
	fmt.Println(LimitBytesWithEnd("Hello, 世界", 13, "..."))
}
