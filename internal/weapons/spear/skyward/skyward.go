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
	core.RegisterWeaponFunc(keys.SkywardSpine, NewWeapon)
}

type Weapon struct {
	Index int
}

func (w *Weapon) SetIndex(idx int) { w.Index = idx }
func (w *Weapon) Init() error      { return nil }

//Increases CRIT Rate by 8% and increases Normal ATK SPD by 12%. Additionally,
//Normal and Charged Attacks hits on opponents have a 50% chance to trigger a
//vacuum blade that deals 40% of ATK as DMG in a small AoE. This effect can
//occur no more than once every 2s.
func NewWeapon(c *core.Core, char *character.CharWrapper, p weapon.WeaponProfile) (weapon.Weapon, error) {
	w := &Weapon{}
	r := p.Refine

	//perm buff
	m := make([]float64, attributes.EndStatType)
	m[attributes.CR] = 0.06 + float64(r)*0.02
	m[attributes.AtkSpd] = 0.12
	char.AddStatMod("skyward spine", -1, attributes.NoStat, func() ([]float64, bool) {
		return m, true
	})

	icd := 0
	atk := .25 + .15*float64(r)
	c.Events.Subscribe(event.OnDamage, func(args ...interface{}) bool {
		ae := args[1].(*combat.AttackEvent)
		//check if char is correct?
		if ae.Info.ActorIndex != char.Index {
			return false
		}
		if ae.Info.AttackTag != combat.AttackTagNormal && ae.Info.AttackTag != combat.AttackTagExtra {
			return false
		}
		//check if cd is up
		if icd > c.F {
			return false
		}
		if c.Rand.Float64() > .5 {
			return false
		}

		//add a new action that deals % dmg immediately
		ai := combat.AttackInfo{
			ActorIndex: char.Index,
			Abil:       "Skyward Spine Proc",
			AttackTag:  combat.AttackTagWeaponSkill,
			ICDTag:     combat.ICDTagNone,
			ICDGroup:   combat.ICDGroupDefault,
			Element:    attributes.Physical,
			Durability: 100,
			Mult:       atk,
		}
		c.QueueAttack(ai, combat.NewDefCircHit(0.1, false, combat.TargettableEnemy), 0, 1)

		//trigger cd
		icd = c.F + 120
		return false
	}, fmt.Sprintf("skyward-spine-%v", char.Base.Key.String()))
	return w, nil
}
