package mage

import (
	"math"
	"time"

	"github.com/wowsims/sod/sim/core"
	"github.com/wowsims/sod/sim/core/proto"
)

// TODO: Classic verify numbers / snapshot / travel time
// https://www.wowhead.com/classic/news/patch-1-15-build-52124-ptr-datamining-season-of-discovery-runes-336044#news-post-336044
// https://www.wowhead.com/classic/spell=401558/living-flame
func (mage *Mage) registerLivingFlameSpell() {
	if !mage.HasRune(proto.MageRune_RuneLegsLivingFlame) {
		return
	}

	baseDamage := mage.baseRuneAbilityDamage() * 1
	spellCoeff := .143
	manaCost := .11
	cooldown := time.Second * 30

	ticks := int32(10)
	tickLength := time.Second * 1

	hasArcaneBlastRune := mage.HasRune(proto.MageRune_RuneHandsArcaneBlast)

	mage.LivingFlame = mage.RegisterSpell(core.SpellConfig{
		ActionID:       core.ActionID{SpellID: int32(proto.MageRune_RuneLegsLivingFlame)},
		ClassSpellMask: ClassSpellMask_MageLivingFlame,
		SpellSchool:    core.SpellSchoolArcane | core.SpellSchoolFire,
		ProcMask:       core.ProcMaskSpellDamage,
		Flags:          core.SpellFlagAPL | core.SpellFlagPureDot,
		MissileSpeed:   6.02,

		ManaCost: core.ManaCostOptions{
			BaseCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    mage.NewTimer(),
				Duration: cooldown,
			},
		},

		// Not affected by hit
		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: "Living Flame",
			},

			NumberOfTicks:    ticks,
			TickLength:       tickLength,
			BonusCoefficient: spellCoeff,

			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
				dot.Snapshot(target, baseDamage, isRollover)

				// We have to deactivate AB here because otherwise the stacks are removed before the snapshot is calculated
				if hasArcaneBlastRune && mage.ArcaneBlastAura.IsActive() {
					mage.ArcaneBlastAura.Deactivate(sim)
				}
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, aoeTarget, dot.OutcomeTick)
				}
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			dot := spell.Dot(target)
			// Ticks lost to travel time
			dot.NumberOfTicks = ticks - int32(math.Floor(spell.TravelTime().Seconds()))
			spell.WaitTravelTime(sim, func(sim *core.Simulation) {
				dot.Apply(sim)
			})
		},
	})

	mage.AddMajorCooldown(core.MajorCooldown{
		Spell: mage.LivingFlame,
		Type:  core.CooldownTypeDPS,
	})
}
