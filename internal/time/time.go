package time

import (
	"encoding/json"
	"strings"
	"time"
)

const begetTimeFormat = "2006-01-02 15:04:05"

type BegetTime struct {
	time.Time
}

func (bt *BegetTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	tm, err := time.Parse(begetTimeFormat, s)
	if err != nil {
		return err
	}

	*bt = BegetTime{Time: tm}

	return nil
}

func (bt *BegetTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(bt.Time)
}

func (bt *BegetTime) String() string {
	return bt.Format(begetTimeFormat)
}

func Parse(value string) (BegetTime, error) {
	tm, err := time.Parse(begetTimeFormat, value)
	bt := BegetTime{Time: tm}

	return bt, err
}

func MustParse(value string) BegetTime {
	bt, err := Parse(value)
	if err != nil {
		panic(err)
	}

	return bt
}
