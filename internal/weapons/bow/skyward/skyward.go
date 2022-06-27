package skyward

import (
	"fmt"

	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/core/player/weapon"
)

func init() {
	core.RegisterWeaponFunc(keys.SkywardHarp, NewWeapon)
}

type Weapon struct {
	Index int
}

func (w *Weapon) SetIndex(idx int) { w.Index = idx }
func (w *Weapon) Init() error      { return nil }

func NewWeapon(c *core.Core, char *character.CharWrapper, p weapon.WeaponProfile) (weapon.Weapon, error) {
	//Increases CRIT DMG by 20%. Hits have a 60% chance to inflict a small AoE attack, dealing 125% Physical
	//ATK DMG. Can only occur once every 4s.
	w := &Weapon{}
	r := p.Refine

	//free crit damage
	m := make([]float64, attributes.EndStatType)
	m[attributes.CD] = 0.15 + float64(r)*0.05
	char.AddStatMod("skyward harp", -1, attributes.NoStat, func() ([]float64, bool) {
		return m, true
	})

	//procs
	cd := 270 - 30*r
	prob := 0.5 + 0.1*float64(r)
	icd := 0

	c.Events.Subscribe(event.OnDamage, func(args ...interface{}) bool {
		atk := args[1].(*combat.AttackEvent)

		if atk.Info.ActorIndex != char.Index {
			return false
		}
		if c.Player.Active() != char.Index {
			return false
		}

		if icd > c.F {
			return false
		}
		if c.Rand.Float64() > prob {
			return false
		}

		ai := combat.AttackInfo{
			ActorIndex: char.Index,
			Abil:       "Skyward Harp Proc",
			AttackTag:  combat.AttackTagWeaponSkill,
			ICDTag:     combat.ICDTagNone,
			ICDGroup:   combat.ICDGroupDefault,
			Element:    attributes.Physical,
			Durability: 100,
			Mult:       1.25,
		}
		c.QueueAttack(ai, combat.NewDefCircHit(2, true, combat.TargettableEnemy), 0, 1)

		icd = c.F + cd

		return false
	}, fmt.Sprintf("skyward-harp-%v", char.Base.Key.String()))

	return w, nil
}
