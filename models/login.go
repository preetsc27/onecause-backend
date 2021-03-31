package models

import (
	"errors"
	"strconv"
)

// LoginReqModel is a struct to unmarshal and validate login POST request
type LoginReqModel struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	OTT      string `json:"ott" validate:"required,len=4,numeric"`
	Hour     uint8
	Minute   uint8
}

// IsOTTValid checks other validation
func (l *LoginReqModel) IsOTTValid() error {
	// 1. Checking if it is even a number
	_, toIntErr := strconv.Atoi(l.OTT)
	if toIntErr != nil {
		return errors.New("Cannot convert OTT to number")
	}
	// 2. Cheking hour range
	hourStr := l.OTT[0:2]
	hourNum, _ := strconv.Atoi(hourStr)
	if hourNum < 0 || hourNum > 23 {
		return errors.New("OTT hour not in range")
	}
	// 3. Cheking minute range
	minuteStr := l.OTT[2:4]
	minuteNum, _ := strconv.Atoi(minuteStr)
	if hourNum < 0 || hourNum > 59 {
		return errors.New("OTT minute not in range")
	}
	l.Hour = uint8(hourNum)
	l.Minute = uint8(minuteNum)
	return nil
}
