package attributes

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type VgaType string

const (
	VgaTypeNone     VgaType = "none"
	VgaTypeCirrus   VgaType = "cirrus" // not recommended
	VgaTypeQxl      VgaType = "qxl"
	VgaTypeQxl2     VgaType = "qxl2"
	VgaTypeQxl3     VgaType = "qxl3"
	VgaTypeQxl4     VgaType = "qxl4"
	VgaTypeSerial0  VgaType = "serial0"
	VgaTypeSerial1  VgaType = "serial1"
	VgaTypeSerial2  VgaType = "serial2"
	VgaTypeSerial3  VgaType = "serial3"
	VgaTypeStd      VgaType = "std" // default
	VgaTypeVirtio   VgaType = "virtio"
	VgaTypeVirtioGl VgaType = "virtio-gl"
	VgaTypeVMWare   VgaType = "vmware"
)

type Vga struct {
	Type   VgaType `json:"type,omitempty"`
	Memory uint64  `json:"memory,omitempty"`
}

func NewVga(t VgaType) *Vga {
	return &Vga{
		Type: t,
	}
}

func (v *Vga) ToApi() string {
	if v == nil {
		return ""
	}
	switch v.Type {
	case VgaTypeSerial0, VgaTypeSerial1, VgaTypeSerial2, VgaTypeSerial3:
		return string(v.Type)
	case VgaTypeQxl, VgaTypeQxl2, VgaTypeQxl3, VgaTypeQxl4:
		v.Memory = 64
		return fmt.Sprintf("%s,memory=%d", v.Type, v.Memory)
	default:
		return string(v.Type)
	}
}

func (v *Vga) ToDomain(s string) *Vga {
	if s == "" {
		return nil
	}
	v.Type = VgaType(strings.Split(s, ",")[0])

	if strings.Contains(s, "memory=") {
		mem := regexp.MustCompile(`memory=(\w+)`).FindStringSubmatch(s)[1]
		val, err := strconv.Atoi(mem)
		if err != nil {
			v.Memory = 0
		} else {
			v.Memory = uint64(val)
		}
	}

	return v
}
