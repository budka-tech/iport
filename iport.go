package iport

import (
	"fmt"
	"github.com/budka-tech/envo"
)

type Port uint16
type Host string

type Ports struct {
	env *envo.Env
	m   map[Host]Port
}

type Params struct {
	Env *envo.Env
	M   map[Host]Port
}

func NewPorts(params *Params) *Ports {
	return &Ports{
		env: params.Env,
		m:   params.M,
	}
}

func (p *Ports) FormatServiceTCP(name Host) string {
	return FormatServiceTCP(p.env, name, p.m)
}

func Format(port int16) string {
	return fmt.Sprintf("%d", port)
}

func FormatTCP(port int16) string {
	return fmt.Sprintf(":%d", port)
}

func FormatServiceTCP(env *envo.Env, name Host, m map[Host]Port) string {
	port := m[name]
	if env.IsLocal() {
		return FormatLocal(port)
	} else {
		return fmt.Sprintf("%s:%d", name, port)
	}
}

func FormatLocal(port Port) string {
	return fmt.Sprintf("localhost:%d", port)
}
