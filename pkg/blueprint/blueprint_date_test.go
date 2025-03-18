package blueprint_test

import (
	"testing"
	"time"

	"github.com/osbuild/blueprint-schema/pkg/blueprint"
)

func TestDateDaysFrom1970(t *testing.T) {
	tests := []struct {
		input blueprint.Date
		want  int
	}{
		{
			input: blueprint.Date{
				time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: 0,
		},
		{
			input: blueprint.Date{
				time.Date(1970, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			want: 1,
		},
		{
			input: blueprint.Date{
				time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC),
			},
			want: 15838,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input.Format("2006-01-02"), func(t *testing.T) {
			got := tt.input.DaysFrom1970()
			if got != tt.want {
				t.Errorf("DateDaysFrom1970() = %v, want %v", got, tt.want)
			}
		})
	}
}
