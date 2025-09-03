package attributes

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Scsi struct {
	Volume string `json:"volume"`
	// unit: GiB
	Size     uint64 `json:"size"`
	Iothread bool   `json:"iothread"`
	Ssd      bool   `json:"ssd"`
	Discard  bool   `json:"discard"`
}

func NewScsi(volume string, size uint64) *Scsi {
	return &Scsi{
		Volume:   volume,
		Size:     size,
		Iothread: true,
		Ssd:      true,
		Discard:  true,
	}
}

func (s *Scsi) ToApi() string {
	if s == nil {
		return ""
	}
	str := fmt.Sprintf("%s:%d,iothread=%t,ssd=%t,discard=%t", s.Volume, s.Size, s.Iothread, s.Ssd, s.Discard)
	str = strings.ReplaceAll(str, "discard=true", "discard=on")
	str = strings.ReplaceAll(str, "discard=false", "discard=ignore")
	str = strings.ReplaceAll(str, "true", "1")
	str = strings.ReplaceAll(str, "false", "0")
	return str
}

// ex: local-lvm:vm-300-disk-1,discard=on,iothread=1,size=16G,ssd=1
func (scsi *Scsi) ToDomain(s string) *Scsi {
	if s == "" {
		return nil
	}

	scsi.Volume = strings.Split(s, ":")[0]

	size := regexp.MustCompile(`size=(\d+)`).FindStringSubmatch(s)[1]
	val, err := strconv.Atoi(size)
	if err != nil {
		scsi.Size = 0
	} else {
		scsi.Size = uint64(val)
	}

	scsi.Iothread = strings.Contains(s, "iothread=1")
	scsi.Ssd = strings.Contains(s, "ssd=1")
	scsi.Discard = strings.Contains(s, "discard=on")
	return scsi
}
