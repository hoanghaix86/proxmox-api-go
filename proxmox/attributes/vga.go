package attributes

import (
	"fmt"
	"regexp"
	"strings"
)

type VgaType string

const (
	VgaTypeStd VgaType = "std"
	VgaTypeQxl VgaType = "qxl"
)

type Vga struct {
	Type      VgaType `json:"type,omitempty"`
	Clipboard string  `json:"clipboard,omitempty"`
	Memory    string  `json:"memory,omitempty"`
}

func NewDefaultVga(t VgaType) *Vga {
	return &Vga{
		Type:      t,
		Clipboard: "vnc",
		Memory:    "64",
	}
}

func (v *Vga) ToApi() string {
	if v == nil {
		return ""
	}
	if v.Type == "" {
		v.Type = VgaTypeQxl
	}
	if v.Clipboard == "" {
		v.Clipboard = "vnc"
	}
	if v.Memory == "" {
		v.Memory = "64"
	}
	return fmt.Sprintf("%s,clipboard=%s,memory=%s", v.Type, v.Clipboard, v.Memory)
}

func (v *Vga) ToDomain(s string) *Vga {
	if s == "" {
		return nil
	}
	v.Type = VgaType(strings.Split(s, ",")[0])
	v.Clipboard = regexp.MustCompile(`clipboard=(\w+)`).FindString(s)
	v.Memory = regexp.MustCompile(`memory=(\w+)`).FindString(s)
	return v
}
