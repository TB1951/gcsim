package travelerelectro

import (
	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/glog"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

var skillFrames []int

const skillHitmark = 55

func init() {
	skillFrames = frames.InitAbilSlice(55)
}

func (c *char) Skill(p map[string]int) action.ActionInfo {
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Lightning Blade",
		AttackTag:  combat.AttackTagElementalArt,
		ICDTag:     combat.ICDTagElementalArt,
		ICDGroup:   combat.ICDGroupDefault,
		StrikeType: combat.StrikeTypeDefault,
		Element:    attributes.Electro,
		Durability: 25,
		Mult:       skill[c.TalentLvlSkill()],
	}
	snap := c.Snapshot(&ai)

	hits, ok := p["hits"]
	if !ok {
		hits = 1
	} else if hits > 3 {
		hits = 3
	}

	maxAmulets := 2
	if c.Base.Cons >= 1 {
		maxAmulets = 3
	}

	// clear existing amulets
	c.abundanceAmulets = 0

	// accept param to limit the amount of amulets generated
	pMaxAmulets, ok := p["max_amulets"]
	if ok && pMaxAmulets < maxAmulets {
		maxAmulets = pMaxAmulets
	}

	// Counting from the frame E is pressed, it takes an average of 1.79 seconds for a character to be able to pick one up
	// https://library.keqingmains.com/evidence/characters/electro/traveler-electro#amulets-delay
	amuletDelay := p["amulet_delay"]
	//make it so that it can't be faster than 1.79s
	if amuletDelay < 107 {
		amuletDelay = 107 // ~1.79s
	}

	//particles appear to be generated if the blades lands but capped at 1
	partCount := 0
	particlesCB := func(atk combat.AttackCB) {
		if partCount > 0 {
			return
		}
		partCount++
		c.Core.QueueParticle(c.Base.Key.String(), 1, attributes.Electro, 100) //this way we're future proof if for whatever reason this misses
	}

	amuletCB := func(atk combat.AttackCB) {
		// generate amulet if generated amulets < limit
		if c.abundanceAmulets >= maxAmulets {
			return
		}

		// 1 amulet per attack
		c.abundanceAmulets++
		c.SetTag("generated", c.abundanceAmulets)

		c.Core.Log.NewEvent("travelerelectro abundance amulet generated", glog.LogCharacterEvent, c.Index, "amulets", c.abundanceAmulets)
	}

	for i := 0; i < hits; i++ {
		c.Core.QueueAttackWithSnap(ai, snap, combat.NewDefCircHit(0.3, false, combat.TargettableEnemy), skillHitmark, particlesCB, amuletCB)
	}

	// try to pick up amulets
	c.Core.Tasks.Add(func() {
		active := c.Core.Player.ActiveChar()
		c.collectAmulets(active)
	}, amuletDelay)

	c.SetCDWithDelay(action.ActionSkill, 810, 21) //13.5s, starts 21 frames in

	return action.ActionInfo{
		Frames:          frames.NewAbilFunc(skillFrames),
		AnimationLength: skillFrames[action.InvalidAction],
		CanQueueAfter:   skillHitmark,
		State:           action.SkillState,
	}
}

func (c *char) collectAmulets(collector *character.CharWrapper) bool {
	// if there are no amulets to collect, return
	if c.abundanceAmulets <= 0 {
		return false
	}

	// Assume all available amulets are collected simultaneously

	mER := make([]float64, attributes.EndStatType)

	mER[attributes.ER] = 0.20

	// handle a4 - Increases the Energy Recharge effect granted by Lightning Blade's Abundance Amulet by 10% of the
	//	 Traveler's Energy Recharge.
	//   This effect only takes into account the Traveler's original Energy Recharge.
	//   Picking up an Amulet to increase the Traveler's ER will not impact the amount of ER shared by
	//   Resounding Roar for other Amulet pickups.
	//   TODO how do we pull unbuffed energy recharge %? Store on init?
	mER[attributes.ER] += c.BaseStats[attributes.ER] * .1

	// apply flat energy
	buffEnergy := skillRegen[c.Talents.Skill] * float64(c.abundanceAmulets)

	// c4 - When a character obtains Abundance Amulets generated by Lightning Blade, if this character's Energy
	//   is less than 35%, the Energy restored by the Abundance Amulets is increased by 100%.
	buffEnergy = c.c4(buffEnergy)

	collector.AddEnergy("abundance-amulet", buffEnergy)

	// handle a1 - When another nearby character in the party obtains an Abundance Amulet created by Lightning Blade,
	//   Lightning Blade's CD is decreased by 1.5s.
	if collector.Index != c.Index {
		c.ReduceActionCooldown(action.ActionSkill, 90*c.abundanceAmulets)
	}

	// apply ER mod
	collector.AddStatMod(character.StatMod{Base: modifier.NewBase("abundance-amulet", 360), AffectedStat: attributes.ER, Amount: func() ([]float64, bool) {
		return mER, true
	}})

	// Reset amulets
	c.abundanceAmulets = 0

	return true
}
