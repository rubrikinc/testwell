package fail

import "testing"

func TestUnifiedDiff(t *testing.T) {
	cases := []struct {
		name        string
		left, right string
		want        string
	}{
		{
			name:  "changed middle line",
			left:  "a\nb\nc",
			right: "a\nB\nc",
			want:  "  a\n- b\n+ B\n  c",
		},
		{
			name:  "added line",
			left:  "a\nc",
			right: "a\nb\nc",
			want:  "  a\n+ b\n  c",
		},
		{
			name:  "removed line",
			left:  "a\nb\nc",
			right: "a\nc",
			want:  "  a\n- b\n  c",
		},
		{
			name:  "identical",
			left:  "a\nb",
			right: "a\nb",
			want:  "  a\n  b",
		},
		{
			name:  "all different",
			left:  "x\ny",
			right: "a\nb",
			want:  "- x\n- y\n+ a\n+ b",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := unifiedDiff(tc.left, tc.right); got != tc.want {
				t.Errorf("unifiedDiff:\ngot:\n%s\nwant:\n%s", got, tc.want)
			}
		})
	}
}
