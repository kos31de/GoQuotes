package cmd

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "gorilla",
			want: "hello gorilla",
		},
		{
			in:   "cat",
			want: "hello cat",
		},
	}

	for _, tt := range tests {
		got := Hello(tt.in)
		if tt.want != got {
			t.Errorf("unexpected result. want=%s, got=%s", tt.want, got)
		}
	}
}
