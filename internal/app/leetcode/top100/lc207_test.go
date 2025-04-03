package top100

import "testing"

func Test_Lc207(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}

	tests := []args{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := canFinish(tt.numCourses, tt.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
