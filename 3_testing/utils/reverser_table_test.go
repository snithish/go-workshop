package utils

import "testing"

func Test_Reverse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Reverse for empty string",
			args: args{
				input: "",
			},
			want: "",
		},
		{
			name: "Reverse for ASCII",
			args: args{
				input: "abc1234",
			},
			want: "4321cba",
		},
		{
			name: "Reverse for Unicode",
			args: args{
				input: "🙏ĈØՖআஹఘ✌️",
			},
			want: "✌️ఘஹআՖØĈ🙏",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.input); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}