package raiden

import "github.com/genshinsim/gsim/pkg/core"

func (c *char) ActionFrames(a core.ActionType, p map[string]int) int {
	switch a {
	case core.ActionAttack:
		f := 0
		if c.Core.Status.Duration("raidenburst") == 0 {
			switch c.NormalCounter {
			//TODO: need to add atkspd mod
			case 0:
				if c.Core.LastAction.Target == "raiden" && c.Core.LastAction.Typ == core.ActionAttack {
					f += 21
				}
				f = 13
			case 1:
				f = 19
			case 2:
				f = 22
			case 3:
				f = 44
			case 4:
				f = 41
			}
		} else {
			switch c.NormalCounter {
			//TODO: need to add atkspd mod
			case 0:
				//add frames if last action is also attack
				if c.Core.LastAction.Target == "raiden" && c.Core.LastAction.Typ == core.ActionAttack {
					f += 32
				}
				f = 12
			case 1:
				f = 32 - 12
			case 2:
				f = 54 - 32
			case 3:
				f = 95 - 54
			case 4:
				f = 139
			}
		}
		f = int(float64(f) / (1 + c.Stats[core.AtkSpd]))
		return f
	case core.ActionCharge:
		if c.Core.Status.Duration("raidenburst") == 0 {
			return 61 //30 if cancelled
		}
		return 79 //37 <- if cancelled
	case core.ActionSkill:
		return 35 // going by first swapable
	case core.ActionBurst:
		return 108
	default:
		c.Core.Log.Warnw("unknown action", "event", core.LogActionEvent, "frame", c.Core.F, "action", a)
		return 0
	}
}