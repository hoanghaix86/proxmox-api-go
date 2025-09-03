package attributes

import (
	"fmt"
	"regexp"
	"strings"
)

type TpmState struct {
	Volume  string `json:"volume"`
	Version string `json:"version"`
}

func NewDefaultTpmState(volume string) *TpmState {
	return &TpmState{
		Volume:  volume,
		Version: "v2.0",
	}
}

func (t *TpmState) ToApi() string {
	if t == nil {
		return ""
	}
	return fmt.Sprintf("%s:1,version=%s", t.Volume, t.Version)
}

func (t *TpmState) ToDomain(s string) *TpmState {
	if t == nil {
		return nil
	}
	t.Volume = strings.Split(s, ":")[0]
	t.Version = regexp.MustCompile(`(v1.2|v2.0)`).FindString(s)
	return t
}
