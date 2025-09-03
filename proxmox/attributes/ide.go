package attributes

import (
	"fmt"
	"strings"
)

// Support CDROM and Cloudinit
type Ide struct {
	Volume string `json:"volume"`
	Iso    string `json:"iso"`
}

func (i *Ide) ToApi() string {
	if i == nil {
		return ""
	}

	if i.Iso == "" {
		return fmt.Sprintf("%s:cloudinit", i.Volume)
	}

	return fmt.Sprintf("%s:%s,media=cdrom", i.Volume, i.Iso)
}

func (i *Ide) ToDomain(s string) *Ide {
	if s == "" {
		return nil
	}
	i.Volume = strings.Split(s, ":")[0]

	if !strings.Contains(s, "cloudinit") {
		i.Iso = strings.Split(strings.Split(s, ",")[0], ":")[1]
	}

	return i
}

func NewIdeIso(volume string, iso string) *Ide {
	return &Ide{
		Volume: volume,
		Iso:    iso,
	}
}

func NewIdeCloudinit(volume string) *Ide {
	return &Ide{
		Volume: volume,
	}
}
