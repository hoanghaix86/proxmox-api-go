package attributes

import (
	"fmt"
	"regexp"
	"strings"
)

type EfIdisk struct {
	Volume          string `json:"volume"`
	EfiType         string `json:"efitype"`
	PreEnrolledKeys bool   `json:"pre-enrolled-keys"`
}

func NewDefaultEfIdisk(volume string) *EfIdisk {
	return &EfIdisk{
		Volume:          volume,
		EfiType:         "4m",
		PreEnrolledKeys: true,
	}
}

func (e *EfIdisk) ToApi() string {
	if e == nil {
		return ""
	}
	str := fmt.Sprintf("%s:0,efitype=%s,pre-enrolled-keys=%t", e.Volume, e.EfiType, e.PreEnrolledKeys)
	str = strings.ReplaceAll(str, "true", "1")
	str = strings.ReplaceAll(str, "false", "0")
	return str
}

func (e *EfIdisk) ToDomain(s string) *EfIdisk {
	if e == nil {
		return nil
	}
	e.Volume = strings.Split(s, ":")[0]
	e.EfiType = regexp.MustCompile(`(4m|2m)`).FindString(s)
	e.PreEnrolledKeys = strings.Contains(s, "pre-enrolled-keys=1")
	return e
}
