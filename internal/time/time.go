package time

import (
	"encoding/json"
	"strings"
	"time"
)

const begetTimeFormat = "2006-01-02 15:04:05"

type BegetTime time.Time

func (bt *BegetTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(begetTimeFormat, s)
	if err != nil {
		return err
	}
	*bt = BegetTime(t)
	return nil
}

func (bt *BegetTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*bt))
}
