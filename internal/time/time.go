// Package time specifies the standard [time.Time] for the Beget time format.
package time

import (
	"encoding/json"
	"strings"
	"time"
)

// BegetTimeFormat is a time format supported by Beget.
const BegetTimeFormat = "2006-01-02 15:04:05"

// BegetTime is an extension of [time.Time] with Beget time format.
type BegetTime struct {
	time.Time
}

// UnmarshalJSON implements unmarshalling of BegetTime from JSON string.
func (bt *BegetTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	tm, err := time.Parse(BegetTimeFormat, s)
	if err != nil {
		return err
	}

	*bt = BegetTime{Time: tm}

	return nil
}

// MarshalJSON implements marshaling of BegetTime to JSON string.
func (bt *BegetTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(bt.Time)
}

// String returns the time formatted by Beget time Format.
func (bt *BegetTime) String() string {
	return bt.Format(BegetTimeFormat)
}

// Parse returns the BegetTime parsed from string.
func Parse(value string) (BegetTime, error) {
	tm, err := time.Parse(BegetTimeFormat, value)
	bt := BegetTime{Time: tm}

	return bt, err
}

// MustParse returns the BegetTime parsed from string.
// If an error appears during the parsing, the function will panic it's an error.
func MustParse(value string) BegetTime {
	bt, err := Parse(value)
	if err != nil {
		panic(err)
	}

	return bt
}
