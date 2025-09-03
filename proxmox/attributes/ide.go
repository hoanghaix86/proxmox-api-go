package attributes

import (
	"fmt"
	"strings"
)

type Ide struct {
	Volume string `json:"volume"`
	Iso    string `json:"iso"`
	Media  string `json:"media"`
}

func (i *Ide) ToApi() string {
	if i == nil {
		return ""
	}

	if i.Iso == "" {
		return fmt.Sprintf("%s:%s,media=disk", i.Volume, i.Media)
	}

	return fmt.Sprintf("%s:%s,media=cdrom", i.Volume, i.Iso)
}

func (i *Ide) ToDomain(s string) *Ide {
	if s == "" {
		return nil
	}
	i.Volume = strings.Split(s, ":")[0]

	if strings.Contains(s, "cdrom") {
		i.Media = "cdrom"
		i.Iso = strings.Split(strings.Split(s, ",")[0], ":")[1]
	} else {
		i.Media = "disk"
	}

	return i
}
