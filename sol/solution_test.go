package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	s := "aa"
	p := "c*d*a*"
	for idx := 0; idx < b.N; idx++ {
		isMatch(s, p)
	}
}
func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "s = \"aa\", p = \"a\"",
			args: args{s: "aa", p: "a"},
			want: false,
		},
		{
			name: "s = \"aa\", p = \"a*\"",
			args: args{s: "aa", p: "a*"},
			want: true,
		},
		{
			name: "s = \"ab\", p = \".*\"",
			args: args{s: "ab", p: ".*"},
			want: true,
		},
		{
			name: "s = \"a\", p = \"a*c\"",
			args: args{s: "a", p: "a*c"},
			want: false,
		},
		{
			name: "s = \"d\", p = \"a*b*\"",
			args: args{s: "d", p: "a*b*"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
