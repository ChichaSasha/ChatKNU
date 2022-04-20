package types_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ChatKNU/pkg/types"
	"github.com/ChatKNU/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestTypeProfileValidate(t *testing.T) {
	locationKyiv, _ := time.LoadLocation("Europe/Kiev")
	tests := []struct {
		name string
		mdl  types.UserProfile
		want error
	}{
		{
			name: "succeeded",
			mdl:  types.UserProfile{DoB: utils.TimeToMillis(time.Date(2001, 4, 16, 0, 0, 0, 0, locationKyiv))},
			want: nil,
		},
		{name: "exactly 13",
			mdl:  types.UserProfile{DoB: utils.TimeToMillis(time.Date(2009, 2, 15, 0, 0, 0, 0, locationKyiv))},
			want: nil,
		},
		{
			name: "younger than 13",
			mdl:  types.UserProfile{DoB: utils.TimeToMillis(time.Date(2015, 2, 15, 0, 0, 0, 0, locationKyiv))},
			want: types.ErrInvalidAge,
		},
		{
			name: "empty age",
			mdl:  types.UserProfile{},
			want: fmt.Errorf("invalid timestamp: negative or zero value"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.mdl.Validate()
			assert.Equal(t, err, tt.want)
		})
	}
}
