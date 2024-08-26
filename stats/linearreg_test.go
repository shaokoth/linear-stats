package linear_stats

import "testing"

func TestLinearReg(t *testing.T) {
	type args struct {
		x []float64
		y []float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		// TODO: Add test cases.
		{
			name: "TestCase",
			args: args{x: []float64{0,1,2,3,4,5,6,7,8,9}, y: []float64{1,2,3,4,5,6,7,8,9,10}},
			want: 1.000000,
			want1: 1.000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := LinearReg(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("LinearReg() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LinearReg() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
