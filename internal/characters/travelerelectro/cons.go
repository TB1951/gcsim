package travelerelectro

import (
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
)

// C2 - Violet Vehemence
// When Falling Thunder created by Bellowing Thunder hits an opponent, it will decrease their Electro RES by 15% for 8s.
func (c *char) c2(t combat.Target) combat.AttackCBFunc {
	return func(a combat.AttackCB) {
		if c.Base.Cons < 2 {
			return
		}
		e, ok := a.Target.(core.Enemy)
		if !ok {
			return
		}
		e.AddResistMod("travelerelectro-c2", 480, attributes.Electro, -0.15)
	}
}

// c4 - When a character obtains Abundance Amulets generated by Lightning Blade, if this character's Energy
//   is less than 35%, the Energy restored by the Abundance Amulets is increased by 100%.
func (c *char) c4(buffEnergy float64) float64 {
	if c.Base.Cons >= 4 {
		collector := c.Core.Player.ActiveChar()
		currentEnergyP := collector.Energy / collector.EnergyMax
		if currentEnergyP < 0.35 {
			buffEnergy = buffEnergy * 2.0
		}
	}
	return buffEnergy
}

// World-Shaker
//  Every 2 Falling Thunder attacks triggered by Bellowing Thunder will greatly increase the DMG
//  dealt by the next Falling Thunder, which will deal 200% of its original DMG [..]
// * Electro traveller's C6 is a multiplicative buff
func (c *char) c6Damage(ai *combat.AttackEvent) {
	if c.Base.Cons >= 6 {
		c.burstC6Hits++
		if c.burstC6Hits >= 3 {
			// TODO will this properly multiply? Should we use a mod instead?
			ai.Info.Mult *= 2
			c.burstC6Hits = 0
			c.burstC6WillGiveEnergy = true
		}
	}
}

// World-Shaker
//  [..] and will restore an additional 1 Energy to the current character.
func (c *char) c6Energy() combat.AttackCBFunc {
	return func(a combat.AttackCB) {
		if c.burstC6WillGiveEnergy {
			active := c.Core.Player.ActiveChar()
			active.AddEnergy("travelerelectro-c6", 1)
			c.burstC6WillGiveEnergy = false
		}
	}
}
