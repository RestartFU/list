package list

import (
	"github.com/df-mc/dragonfly/server/cmd"
)

func NewRunnable(w *List, allower func(src cmd.Source) bool) Command {
	return Command{whitelist: w, allower: allower}
}

type Command struct {
	whitelist *List
	allower   func(src cmd.Source) bool
	Status    status
}

func (c Command) Run(src cmd.Source, o *cmd.Output) {
	switch string(c.Status) {
	case "enable", "on":
		if c.whitelist.Enabled {
			o.Print("server is already whitelisted")
			return
		}
		c.whitelist.Enabled = true
		o.Print("server is now whitelisted")
	case "disable", "off":
		if !c.whitelist.Enabled {
			o.Print("server is not whitelisted")
			return
		}
		c.whitelist.Enabled = false
		o.Print("server is no longer whitelisted")
	default:
		o.Printf("'%s' is not a valid argument for this command!", c.Status)
	}
}

func (c Command) Allow(src cmd.Source) bool { return c.allower(src) }

type status string

func (status) Type() string                    { return "status" }
func (status) Options(src cmd.Source) []string { return []string{"enable", "disable"} }
