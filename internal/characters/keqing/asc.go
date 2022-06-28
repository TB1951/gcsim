package keqing

import (
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

func (c *char) a1() {
	//account for it starting somewhere around hitmark
	dur := 300 + skillRecastHitmark
	c.Core.Status.Add("keqinginfuse", dur)
	c.Core.Player.AddWeaponInfuse(
		c.Index,
		"keqing-a1",
		attributes.Electro,
		dur,
		true,
		combat.AttackTagNormal, combat.AttackTagExtra, combat.AttackTagPlunge,
	)
}

func (c *char) a4() {
	m := make([]float64, attributes.EndStatType)
	m[attributes.CR] = 0.15
	m[attributes.ER] = 0.15

	c.AddStatMod(character.StatMod{Base: modifier.NewBase("keqing-a4", 480), AffectedStat: attributes.NoStat, Amount: func() ([]float64, bool) {
		return m, true
	}})
}
