package attributes

import (
	"fmt"
	"regexp"
	"strings"
)

type Agent struct {
	Enabled           bool   `json:"enabled"`
	FreezeFsOnBackup  bool   `json:"freeze-fs-on-backup"`
	FstrimClonedDisks bool   `json:"fstrim_cloned_disks"`
	Type              string `json:"type"`
}

func NewAgent() *Agent {
	return &Agent{
		Enabled:           true,
		FreezeFsOnBackup:  true,
		FstrimClonedDisks: true,
		Type:              "virtio",
	}
}

func (a *Agent) ToApi() string {
	str := fmt.Sprintf("enabled=%t,freeze-fs-on-backup=%t,fstrim_cloned_disks=%t,type=%s", a.Enabled, a.FreezeFsOnBackup, a.FstrimClonedDisks, a.Type)
	str = strings.ReplaceAll(str, "true", "1")
	str = strings.ReplaceAll(str, "false", "0")
	return str
}

func (a *Agent) ToDomain(s string) *Agent {
	a.Enabled = strings.Contains(s, "enabled=1") || strings.HasPrefix(s, "1")
	a.FreezeFsOnBackup = strings.Contains(s, "freeze-fs-on-backup=1")
	a.FstrimClonedDisks = strings.Contains(s, "fstrim_cloned_disks=1")

	virtio := regexp.MustCompile("(virtio|isa)").FindString(s)
	if virtio != "" {
		a.Type = virtio
	}

	return a
}
