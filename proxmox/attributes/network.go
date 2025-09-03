package attributes

import (
	"fmt"
	"regexp"
	"strings"
)

type Network struct {
	Type     string `json:"type"`
	Bridge   string `json:"bridge"`
	Firewall bool   `json:"firewall"`
}

func NewDefaultNetwork(bridge string) *Network {
	return &Network{
		Type:     "virtio",
		Bridge:   bridge,
		Firewall: false,
	}
}

func (n *Network) ToApi() string {
	if n == nil {
		return ""
	}

	if n.Type == "" {
		n.Type = "virtio"
	}

	str := fmt.Sprintf("%s,bridge=%s,firewall=%t", n.Type, n.Bridge, n.Firewall)
	str = strings.ReplaceAll(str, "true", "1")
	str = strings.ReplaceAll(str, "false", "0")
	return str
}

func (n *Network) ToDomain(s string) *Network {
	if n == nil {
		return nil
	}
	n.Type = regexp.MustCompile(`(virtio|isa)`).FindString(s)
	n.Bridge = regexp.MustCompile(`bridge=(\w+)`).FindStringSubmatch(s)[1]
	n.Firewall = strings.Contains(s, "firewall=1")
	return n
}
