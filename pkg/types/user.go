package types

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/ChatKNU/pkg/utils"
)

var ErrInvalidAge = errors.New("age must be over 13")

type UserProfile struct {
	ID        string `firestore:"id"`
	Name      string `firestore:"name"`
	DoB       int64  `firestore:"dob"`
	Location  string `firestore:"location"`
	Bio       string `firestore:"bio"`
	Email     string `firestore:"email"`
	PushToken string `firestore:"push_token"`
}

func (p UserProfile) String() string {
	return fmt.Sprintf("UserProfile(%s, %s, %s)", p.ID, p.Name, p.Location)
}

func (p UserProfile) Validate() error {
	now := utils.TimeToMillis(time.Now())
	ageInDays, err := utils.DaysBetweenUnixMillis(p.DoB, now)
	if err != nil {
		return err
	}

	if math.Abs(ageInDays/365) <= 13 || math.Abs(ageInDays/365) >= 150 {
		return ErrInvalidAge
	}
	return nil
}
