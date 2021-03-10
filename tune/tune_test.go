package tune

import "testing"

func TestCongestionTune_Level(t *testing.T) {
	type fields struct {
		upperBound int64
	}
	type args struct {
		now int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   level
	}{
		// TODO: Add test cases.
		{
			"case-normal",
			fields{10},
			args{3},
			levelNormal,
		},
		{
			"case-normal-bound",
			fields{10},
			args{5},
			levelWarn,
		},
		{
			"case-warn",
			fields{10},
			args{6},
			levelWarn,
		},
		{
			"case-warn-bound",
			fields{10},
			args{8},
			levelCritical,
		},
		{
			"case-error-lower-bound",
			fields{10},
			args{10},
			levelError,
		},
		{
			"case-error",
			fields{10},
			args{11},
			levelError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ct := &CongestionTune{
				upperBound: tt.fields.upperBound,
			}
			if got := ct.getLevel(tt.args.now); got != tt.want {
				t.Errorf("Level() = %v, want %v", got, tt.want)
			}
		})
	}
}
