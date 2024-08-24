package linear_stats

import "testing"

func TestPearsonCorr(t *testing.T) {
	type args struct {
		x []float64
		y []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "Testcase",
			args: args{x: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, y: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 1.0000000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PearsonCorr(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("PearsonCorr() = %v, want %v", got, tt.want)
			}
		})
	}
}
