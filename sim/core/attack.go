package core

import (
	"slices"
	"time"

	"github.com/wowsims/sod/sim/core/proto"
	"github.com/wowsims/sod/sim/core/stats"
)

// ReplaceMHSwing is called right before a main hand auto attack fires.
// It must never return nil, but either a replacement spell or the passed in regular mhSwingSpell.
// This allows for abilities that convert a white attack into a yellow attack.
type ReplaceMHSwing func(sim *Simulation, mhSwingSpell *Spell) *Spell

// Represents a generic weapon. Pets / unarmed / various other cases don't use
// actual weapon items so this is an abstraction of a Weapon.
type Weapon struct {
	BaseDamageMin        float64
	BaseDamageMax        float64
	AttackPowerPerDPS    float64
	SwingSpeed           float64
	NormalizedSwingSpeed float64
	SpellSchool          SpellSchool
	BonusCoefficient     float64
	MinRange             float64
	MaxRange             float64
}

func (weapon *Weapon) DPS() float64 {
	if weapon.SwingSpeed == 0 {
		return 0
	}
	return (weapon.BaseDamageMin + weapon.BaseDamageMax) / 2.0 / weapon.SwingSpeed
}

func newWeaponFromUnarmed() Weapon {
	// These numbers are probably wrong but nobody cares.
	return Weapon{
		BaseDamageMin:        0,
		BaseDamageMax:        0,
		SwingSpeed:           1,
		NormalizedSwingSpeed: 1,
		AttackPowerPerDPS:    DefaultAttackPowerPerDPS,
		MinRange:             0,
		MaxRange:             MaxMeleeAttackRange,
	}
}

func getWeaponMinRange(item *Item) float64 {
	switch item.RangedWeaponType {
	case proto.RangedWeaponType_RangedWeaponTypeThrown:
		return 5
	case proto.RangedWeaponType_RangedWeaponTypeUnknown, proto.RangedWeaponType_RangedWeaponTypeWand:
		return 0
	default:
		return MinRangedAttackRange
	}
}

func getWeaponMaxRange(item *Item) float64 {
	switch item.RangedWeaponType {
	case proto.RangedWeaponType_RangedWeaponTypeUnknown:
		return MaxMeleeAttackRange
	default:
		return MaxRangedAttackRange
	}
}

func newWeaponFromItem(item *Item, bonusDps float64) Weapon {
	normalizedWeaponSpeed := 2.4
	if item.WeaponType == proto.WeaponType_WeaponTypeDagger {
		normalizedWeaponSpeed = 1.7
	} else if item.HandType == proto.HandType_HandTypeTwoHand {
		normalizedWeaponSpeed = 3.3
	} else if item.RangedWeaponType != proto.RangedWeaponType_RangedWeaponTypeUnknown {
		normalizedWeaponSpeed = 2.8
	}

	spellSchool := SpellSchoolFromProto(item.SpellSchool)
	return Weapon{
		BaseDamageMin:        item.WeaponDamageMin + bonusDps*item.SwingSpeed,
		BaseDamageMax:        item.WeaponDamageMax + bonusDps*item.SwingSpeed,
		SwingSpeed:           item.SwingSpeed,
		NormalizedSwingSpeed: normalizedWeaponSpeed,
		AttackPowerPerDPS:    DefaultAttackPowerPerDPS,
		SpellSchool:          spellSchool,
		BonusCoefficient:     Ternary(spellSchool == SpellSchoolPhysical, 0, 1.0),
		MinRange:             getWeaponMinRange(item),
		MaxRange:             getWeaponMaxRange(item),
	}
}

// Returns weapon stats using the main hand equipped weapon.
func (character *Character) WeaponFromMainHand() Weapon {
	if weapon := character.GetMHWeapon(); weapon != nil {
		return newWeaponFromItem(weapon, character.PseudoStats.BonusMHDps)
	} else {
		return newWeaponFromUnarmed()
	}
}

// Returns weapon stats using the off-hand equipped weapon.
func (character *Character) WeaponFromOffHand() Weapon {
	if weapon := character.GetOHWeapon(); weapon != nil {
		return newWeaponFromItem(weapon, character.PseudoStats.BonusOHDps)
	} else {
		return Weapon{}
	}
}

// Returns weapon stats using the ranged equipped weapon.
func (character *Character) WeaponFromRanged() Weapon {
	if weapon := character.GetRangedWeapon(); weapon != nil {
		return newWeaponFromItem(weapon, character.PseudoStats.BonusRangedDps)
	} else {
		return Weapon{}
	}
}

func (weapon *Weapon) GetSpellSchool() SpellSchool {
	if weapon.SpellSchool == SpellSchoolNone {
		return SpellSchoolPhysical
	} else {
		return weapon.SpellSchool
	}
}

func (weapon *Weapon) GetBonusCoefficient() float64 {
	if weapon.BonusCoefficient > 0 {
		return weapon.BonusCoefficient
	} else {
		return Ternary(weapon.GetSpellSchool() == SpellSchoolPhysical, 1.0, 0.0)
	}
}

func (weapon *Weapon) EnemyWeaponDamage(sim *Simulation, attackPower float64, damageSpread float64) float64 {
	// Maximum damage range is 133% of minimum damage; AP contribution is % of minimum damage roll.
	// Patchwerk follows special damage range rules.
	// TODO: Scrape more logs to determine these values more accurately. AP defined in constants.go

	rand := 1 + damageSpread*sim.RandomFloat("Enemy Weapon Damage")

	return weapon.BaseDamageMin * (rand + attackPower*EnemyAutoAttackAPCoefficient)
}

func (weapon *Weapon) BaseDamage(sim *Simulation) float64 {
	return weapon.BaseDamageMin + (weapon.BaseDamageMax-weapon.BaseDamageMin)*sim.RandomFloat("Weapon Base Damage")
}

func (weapon *Weapon) AverageDamage() float64 {
	return (weapon.BaseDamageMin + weapon.BaseDamageMax) / 2
}

func (weapon *Weapon) CalculateWeaponDamage(sim *Simulation, attackPower float64) float64 {
	return weapon.BaseDamage(sim) + (weapon.SwingSpeed*attackPower)/weapon.AttackPowerPerDPS
}

func (weapon *Weapon) CalculateAverageWeaponDamage(attackPower float64) float64 {
	return weapon.AverageDamage() + (weapon.SwingSpeed*attackPower)/weapon.AttackPowerPerDPS
}

func (weapon *Weapon) CalculateNormalizedWeaponDamage(sim *Simulation, attackPower float64) float64 {
	return weapon.BaseDamage(sim) + (weapon.NormalizedSwingSpeed*attackPower)/weapon.AttackPowerPerDPS
}

func (unit *Unit) MHWeaponDamage(sim *Simulation, attackPower float64) float64 {
	return unit.AutoAttacks.mh.CalculateWeaponDamage(sim, attackPower)
}
func (unit *Unit) MHNormalizedWeaponDamage(sim *Simulation, attackPower float64) float64 {
	return unit.AutoAttacks.mh.CalculateNormalizedWeaponDamage(sim, attackPower)
}

func (unit *Unit) OHWeaponDamage(sim *Simulation, attackPower float64) float64 {
	return 0.5 * unit.AutoAttacks.oh.CalculateWeaponDamage(sim, attackPower)
}
func (unit *Unit) OHNormalizedWeaponDamage(sim *Simulation, attackPower float64) float64 {
	return 0.5 * unit.AutoAttacks.oh.CalculateNormalizedWeaponDamage(sim, attackPower)
}

func (unit *Unit) RangedWeaponDamage(sim *Simulation, attackPower float64) float64 {
	return unit.AutoAttacks.ranged.CalculateWeaponDamage(sim, attackPower)
}

// Returns whether this hit effect is associated with the main-hand weapon.
func (spell *Spell) IsMH() bool {
	return spell.ProcMask.Matches(ProcMaskMeleeMH)
}

// Returns whether this hit effect is associated with the off-hand weapon.
func (spell *Spell) IsOH() bool {
	return spell.ProcMask.Matches(ProcMaskMeleeOH)
}

// Returns whether this hit effect is associated with either melee weapon.
func (spell *Spell) IsMelee() bool {
	return spell.ProcMask.Matches(ProcMaskMelee)
}

func (aa *AutoAttacks) MH() *Weapon {
	return aa.mh.getWeapon()
}

func (aa *AutoAttacks) SetMH(sim *Simulation, weapon Weapon) {
	extraAttacksAura := aa.mh.extraAttacksAura
	aa.mh.setWeapon(sim, weapon)

	if extraAttacksAura != nil && aa.mh.extraAttacksAura == nil {
		aa.mh.extraAttacksAura = extraAttacksAura
	}
}

func (aa *AutoAttacks) OH() *Weapon {
	return aa.oh.getWeapon()
}

func (aa *AutoAttacks) SetOH(sim *Simulation, weapon Weapon) {
	aa.oh.setWeapon(sim, weapon)
}

func (aa *AutoAttacks) Ranged() *Weapon {
	return aa.ranged.getWeapon()
}

func (aa *AutoAttacks) SetRanged(sim *Simulation, weapon Weapon) {
	aa.ranged.setWeapon(sim, weapon)
}

func (aa *AutoAttacks) MHAuto() *Spell {
	return aa.mh.spell
}

func (aa *AutoAttacks) OHAuto() *Spell {
	return aa.oh.spell
}

func (aa *AutoAttacks) RangedAuto() *Spell {
	return aa.ranged.spell
}

func (aa *AutoAttacks) MainhandSwingAt() time.Duration {
	return aa.mh.swingAt
}

func (aa *AutoAttacks) OffhandSwingAt() time.Duration {
	return aa.oh.swingAt
}

func (aa *AutoAttacks) SetOffhandSwingAt(offhandSwingAt time.Duration) {
	aa.oh.swingAt = offhandSwingAt
}

func (aa *AutoAttacks) SetReplaceMHSwing(replaceSwing ReplaceMHSwing) {
	aa.mh.replaceSwing = replaceSwing
}

func (aa *AutoAttacks) MHConfig() *SpellConfig {
	return &aa.mh.config
}

func (aa *AutoAttacks) OHConfig() *SpellConfig {
	return &aa.oh.config
}

func (aa *AutoAttacks) RangedConfig() *SpellConfig {
	return &aa.ranged.config
}

const (
	tagMainhand    = 1
	tagOffhand     = 2
	tagExtraAttack = 3
)

type WeaponAttack struct {
	Weapon

	agent Agent
	unit  *Unit

	config SpellConfig
	spell  *Spell

	replaceSwing ReplaceMHSwing

	swingAt             time.Duration
	lastSwingAt         time.Duration
	extraAttacks        int32 // extra attacks that happen right away
	extraAttacksStored  int32 // extra attack that happen on next auto (e.g reckoning)
	extraAttacksPending int32 // extraAttacks prior to previous ones resolving for spell metrics
	extraAttacksAura    *Aura

	curSwingSpeed    float64
	curSwingDuration time.Duration
	enabled          bool
	allowed          bool
}

func (wa *WeaponAttack) getWeapon() *Weapon {
	return &wa.Weapon
}

func (wa *WeaponAttack) setWeapon(sim *Simulation, weapon Weapon) {
	wa.Weapon = weapon

	if wa.extraAttacksAura != nil {
		wa.extraAttacksAura.SetStacks(sim, 0)
		wa.extraAttacks = 0
		wa.extraAttacksPending = 0
		wa.extraAttacksStored = 0
	}

	bonusCoeff := weapon.GetBonusCoefficient()
	school := weapon.GetSpellSchool()

	wa.config.BonusCoefficient = bonusCoeff
	wa.config.SpellSchool = school

	// I know these shouldn't really be touched but this is a very niche case
	// where if we're swapping to/from a weapon with another spell school than physical.
	// These are also basically the only "stats" on a weapon that gets handled by the spell, not the WeaponAttack.Weapon.
	wa.spell.BonusCoefficient = bonusCoeff
	wa.spell.SpellSchool = school
	wa.spell.SchoolIndex = school.GetSchoolIndex()
	wa.spell.SchoolBaseIndices = school.GetBaseIndices()

	wa.updateSwingDuration(wa.curSwingSpeed)
}

// inlineable stub for swing
func (wa *WeaponAttack) trySwing(sim *Simulation) time.Duration {
	if sim.CurrentTime < wa.swingAt {
		return wa.swingAt
	} else if !wa.allowed {
		wa.swingAt = sim.CurrentTime + wa.curSwingDuration // Auto Attack Fails due to being disabled by an item or effect (e.g. Ravagane)
        sim.rescheduleWeaponAttack(wa.swingAt)
		return wa.swingAt
	}
	return wa.swing(sim)
}

func (wa *WeaponAttack) castExtraAttacksStored(sim *Simulation) {
	if wa.extraAttacksAura != nil {
		wa.extraAttacksStored = wa.extraAttacksAura.GetStacks()
		wa.extraAttacksAura.SetStacks(sim, 0)
	}

	wa.castExtraAttacks(sim, wa.extraAttacksStored, 0)
	wa.extraAttacksStored = 0

}

func (wa *WeaponAttack) castExtraAttacksTriggered(sim *Simulation, moreAttacks bool) {
	if moreAttacks {
		wa.castExtraAttacks(sim, wa.extraAttacks, 0)
	} else {
		wa.castExtraAttacks(sim, wa.extraAttacks, 1)
	}
	wa.extraAttacks = 0
}

func (wa *WeaponAttack) castExtraAttacks(sim *Simulation, numExtraAttacks int32, startIndex int32) bool {
	if numExtraAttacks > 0 {
		// if startIndex==1, Ignore the first extra attack, that was used to speed up next attack
		wa.spell.SetMetricsSplit(1)
		wa.spell.CastTimeMultiplier -= 1

		for i := int32(startIndex); i < numExtraAttacks; i++ {
			// use original attacks for subsequent extra Attacks
			wa.spell.Cast(sim, wa.unit.CurrentTarget)
		}

		wa.spell.CastTimeMultiplier += 1

		return true
	}
	return false

}

func (wa *WeaponAttack) swing(sim *Simulation) time.Duration {
	isExtraAttack := wa.extraAttacksPending > 0

	if isExtraAttack {
		// Needs to happen before any attacks gets cast
		wa.extraAttacks = wa.extraAttacksPending
		wa.extraAttacksPending = 0
		// Any further procs will be added to extraAttacksPending to be processed next batch
	}

	wa.castExtraAttacksStored(sim)

	attackSpell := wa.spell

	if wa.replaceSwing != nil {
		// Need to check APL here to allow last-moment HS queue casts.
		wa.unit.Rotation.DoNextAction(sim)

		// Allow MH swing to be overridden for abilities like Heroic Strike.
		attackSpell = wa.replaceSwing(sim, attackSpell)
	}

	if attackSpell.CanCast(sim, wa.unit.CurrentTarget) {
		// Update swing timer BEFORE the cast, so that APL checks for TimeToNextAuto behave correctly
		// if the attack causes APL evaluations (e.g. from rage gain).

		wa.swingAt = sim.CurrentTime + wa.curSwingDuration
		wa.lastSwingAt = sim.CurrentTime

		// don't update isExtraAttack here

		var originalCastTime time.Duration
		if isExtraAttack {
			originalCastTime = wa.spell.DefaultCast.CastTime
			wa.spell.DefaultCast.CastTime = 0
			wa.spell.SetMetricsSplit(1)
		} else {
			wa.spell.SetMetricsSplit(0)
		}

		attackSpell.Cast(sim, wa.unit.CurrentTarget)

		moreAttacks := !isExtraAttack && wa.extraAttacksPending > 0 // True if above cast is a normal Auto attack that triggered an Extra Attack
		wa.castExtraAttacksTriggered(sim, moreAttacks)              // more attacks means we don't count the above cast

		if isExtraAttack {
			// For ranged extra attacks, we have to wait for the spell to hit before resettings the cast time and metrics split
			if originalCastTime > 0 {
				wa.spell.WaitTravelTime(sim, func(sim *Simulation) {
					wa.spell.DefaultCast.CastTime = originalCastTime
					wa.spell.SetMetricsSplit(0)
				})
			} else {
				wa.spell.SetMetricsSplit(0)
			}
		}

		if wa.extraAttacksPending > 0 {
			wa.spell.SetMetricsSplit(1)
			wa.swingAt = sim.CurrentTime + SpellBatchWindow
			wa.lastSwingAt = sim.CurrentTime
			sim.rescheduleWeaponAttack(wa.swingAt) // Required to fix extra attack procs triggered during swing
		}

		if !sim.Options.Interactive && wa.unit.Rotation != nil {
			wa.unit.Rotation.DoNextAction(sim)
		}
	} else {
		// Delay till cast finishes if casting or 100 ms if not
		wa.swingAt = max(wa.unit.Hardcast.Expires, sim.CurrentTime+time.Millisecond*100)
	}

	return wa.swingAt
}

func (wa *WeaponAttack) updateSwingDuration(curSwingSpeed float64) {
	wa.curSwingSpeed = curSwingSpeed
	wa.curSwingDuration = DurationFromSeconds(wa.SwingSpeed / wa.curSwingSpeed)
}

func (wa *WeaponAttack) addWeaponAttack(sim *Simulation, swingSpeed float64) {
	if !wa.enabled {
		return
	}

	wa.updateSwingDuration(swingSpeed)
	sim.addWeaponAttack(wa)
	sim.rescheduleWeaponAttack(wa.swingAt)
}

type AutoAttacks struct {
	AutoSwingMelee  bool
	AutoSwingRanged bool

	AutoSwingDisabled bool

	IsDualWielding bool

	character *Character

	mh     WeaponAttack
	oh     WeaponAttack
	ranged WeaponAttack
}

// Options for initializing auto attacks.
type AutoAttackOptions struct {
	MainHand        Weapon
	OffHand         Weapon
	Ranged          Weapon
	AutoSwingMelee  bool // If true, core engine will handle calling SwingMelee() for you.
	AutoSwingRanged bool // If true, core engine will handle calling SwingRanged() for you.
	ReplaceMHSwing  ReplaceMHSwing
}

func (unit *Unit) EnableAutoAttacks(agent Agent, options AutoAttackOptions) {
	if options.MainHand.AttackPowerPerDPS == 0 {
		options.MainHand.AttackPowerPerDPS = DefaultAttackPowerPerDPS
	}
	if options.OffHand.AttackPowerPerDPS == 0 {
		options.OffHand.AttackPowerPerDPS = DefaultAttackPowerPerDPS
	}

	unit.AutoAttacks = AutoAttacks{
		AutoSwingMelee:  options.AutoSwingMelee,
		AutoSwingRanged: options.AutoSwingRanged,

		IsDualWielding: options.OffHand.SwingSpeed != 0,

		character: agent.GetCharacter(),

		mh: WeaponAttack{
			agent:        agent,
			unit:         unit,
			Weapon:       options.MainHand,
			replaceSwing: options.ReplaceMHSwing,
		},
		oh: WeaponAttack{
			agent:  agent,
			unit:   unit,
			Weapon: options.OffHand,
		},
		ranged: WeaponAttack{
			agent:  agent,
			unit:   unit,
			Weapon: options.Ranged,
		},
	}

	unit.AutoAttacks.mh.config = SpellConfig{
		ActionID:    ActionID{OtherID: proto.OtherAction_OtherActionAttack, Tag: tagMainhand},
		SpellSchool: options.MainHand.GetSpellSchool(),
		DefenseType: DefenseTypeMelee,
		ProcMask:    ProcMaskMeleeMHAuto,
		Flags:       SpellFlagMeleeMetrics | SpellFlagNoOnCastComplete,
		CastType:    proto.CastType_CastTypeMainHand,

		MaxRange:     MaxMeleeAttackRange,
		MetricSplits: 2,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: options.MainHand.GetBonusCoefficient(),

		ApplyEffects: func(sim *Simulation, target *Unit, spell *Spell) {
			autoInProgress := *spell
			baseDamage := autoInProgress.Unit.MHWeaponDamage(sim, autoInProgress.MeleeAttackPower())
			result := autoInProgress.CalcDamage(sim, target, baseDamage, autoInProgress.OutcomeMeleeWhite)

			splitIdx := spell.GetMetricsSplitIdx()

			StartDelayedAction(sim, DelayedActionOptions{
				DoAt: sim.CurrentTime + SpellBatchWindow,
				OnAction: func(sim *Simulation) {
					newSplitIdx := autoInProgress.GetMetricsSplitIdx()

					// We have to dynamically update the split metrics to ensure extra attacks' damage are categorized correctly
					autoInProgress.SetMetricsSplit(splitIdx)
					autoInProgress.DealDamage(sim, result)
					autoInProgress.SetMetricsSplit(newSplitIdx)
				},
			})
		},
	}

	unit.AutoAttacks.oh.config = SpellConfig{
		ActionID:    ActionID{OtherID: proto.OtherAction_OtherActionAttack, Tag: tagOffhand},
		SpellSchool: options.OffHand.GetSpellSchool(),
		DefenseType: DefenseTypeMelee,
		ProcMask:    ProcMaskMeleeOHAuto,
		Flags:       SpellFlagMeleeMetrics | SpellFlagNoOnCastComplete,
		CastType:    proto.CastType_CastTypeOffHand,

		MaxRange:     MaxMeleeAttackRange,
		MetricSplits: 2,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: options.OffHand.GetBonusCoefficient(),

		ApplyEffects: func(sim *Simulation, target *Unit, spell *Spell) {
			autoInProgress := *spell
			baseDamage := autoInProgress.Unit.OHWeaponDamage(sim, autoInProgress.MeleeAttackPower())
			result := autoInProgress.CalcDamage(sim, target, baseDamage, autoInProgress.OutcomeMeleeWhite)

			StartDelayedAction(sim, DelayedActionOptions{
				DoAt: sim.CurrentTime + SpellBatchWindow,
				OnAction: func(sim *Simulation) {
					autoInProgress.DealDamage(sim, result)
				},
			})
		},
	}

	unit.AutoAttacks.ranged.config = SpellConfig{
		ActionID:     ActionID{OtherID: proto.OtherAction_OtherActionShoot},
		SpellSchool:  options.Ranged.GetSpellSchool(),
		DefenseType:  DefenseTypeRanged,
		ProcMask:     ProcMaskRangedAuto,
		Flags:        SpellFlagMeleeMetrics,
		CastType:     proto.CastType_CastTypeRanged,
		MissileSpeed: 24,

		MinRange:     unit.AutoAttacks.ranged.MinRange,
		MaxRange:     unit.AutoAttacks.ranged.MaxRange,
		MetricSplits: 2,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: options.Ranged.GetBonusCoefficient(),

		ApplyEffects: func(sim *Simulation, target *Unit, spell *Spell) {
			baseDamage := spell.Unit.RangedWeaponDamage(sim, spell.RangedAttackPower(target, false))
			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeRangedHitAndCrit)

			splitIdx := spell.GetMetricsSplitIdx()

			spell.WaitTravelTime(sim, func(sim *Simulation) {
				newSplitIdx := spell.GetMetricsSplitIdx()

				// We have to dynamically update the split metrics to ensure extra attacks' damage are categorized correctly
				spell.SetMetricsSplit(splitIdx)
				spell.DealDamage(sim, result)
				spell.SetMetricsSplit(newSplitIdx)
			})
		},
	}

	if unit.Type == EnemyUnit {
		unit.AutoAttacks.mh.config.ApplyEffects = func(sim *Simulation, target *Unit, spell *Spell) {
			ap := max(0, spell.Unit.stats[stats.AttackPower])
			baseDamage := spell.Unit.AutoAttacks.mh.EnemyWeaponDamage(sim, ap, spell.Unit.PseudoStats.DamageSpread)

			spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeEnemyMeleeWhite)
		}
		unit.AutoAttacks.oh.config.ApplyEffects = func(sim *Simulation, target *Unit, spell *Spell) {
			ap := max(0, spell.Unit.stats[stats.AttackPower])
			baseDamage := spell.Unit.AutoAttacks.mh.EnemyWeaponDamage(sim, ap, spell.Unit.PseudoStats.DamageSpread) * 0.5

			spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeEnemyMeleeWhite)
		}
	} else {
		unit.AutoAttacks.mh.extraAttacksAura = MakePermanent(unit.RegisterAura(Aura{
			Label:     "Extra Attacks (Main Hand)", // Tracks Stored Extra Attacks from all sources
			ActionID:  ActionID{SpellID: 21919},    // Thrash ID
			Duration:  NeverExpires,
			MaxStacks: 4, // Max is 4 extra attacks stored - more can proc after
		}))

		unit.AutoAttacks.ranged.extraAttacksAura = MakePermanent(unit.RegisterAura(Aura{
			Label:     "Extra Attacks (Ranged)", // Tracks Stored Extra Attacks from all sources
			ActionID:  ActionID{SpellID: 21919}, // Thrash ID
			Duration:  NeverExpires,
			MaxStacks: 4, // Max is 4 extra attacks stored - more can proc after
		}))
	}
}

func (aa *AutoAttacks) finalize() {
	if aa.AutoSwingMelee {
		aa.mh.spell = aa.mh.unit.GetOrRegisterSpell(aa.mh.config)
		aa.mh.spell.TagSplitMetric(1, tagExtraAttack)
		// Will keep the OH spell registered for Item swapping
		aa.oh.spell = aa.oh.unit.GetOrRegisterSpell(aa.oh.config)
		aa.oh.spell.TagSplitMetric(1, tagExtraAttack)
	}
	if aa.AutoSwingRanged {
		aa.ranged.spell = aa.ranged.unit.GetOrRegisterSpell(aa.ranged.config)
		aa.ranged.spell.TagSplitMetric(1, tagExtraAttack)
	}
}

func (aa *AutoAttacks) anyEnabled() bool {
	return aa.mh.enabled || aa.oh.enabled || aa.ranged.enabled
}

func (aa *AutoAttacks) reset(sim *Simulation) {
	if !aa.AutoSwingMelee && !aa.AutoSwingRanged {
		return
	}

	aa.mh.enabled = false
	aa.oh.enabled = false
	aa.ranged.enabled = false

	aa.mh.allowed = true
	aa.oh.allowed = true
	aa.ranged.allowed = true

	aa.mh.swingAt = NeverExpires
	aa.oh.swingAt = NeverExpires

	// Make sure extra attacks are reset
	aa.mh.extraAttacks = 0
	aa.mh.extraAttacksPending = 0
	aa.mh.spell.SetMetricsSplit(0)

	if aa.ranged.spell != nil {
		aa.ranged.extraAttacks = 0
		aa.ranged.extraAttacksPending = 0
		aa.ranged.spell.SetMetricsSplit(0)
	}

	if aa.AutoSwingMelee {
		aa.mh.updateSwingDuration(aa.mh.unit.SwingSpeed())
		aa.mh.swingAt = 0

		if aa.IsDualWielding {
			aa.oh.extraAttacks = 0
			aa.oh.extraAttacksPending = 0
			aa.oh.spell.SetMetricsSplit(0)
			aa.oh.updateSwingDuration(aa.mh.curSwingSpeed)
			aa.oh.swingAt = 0

			// Apply random delay of 0 - 50% swing time, to one of the weapons if dual wielding
			if aa.oh.unit.Type == EnemyUnit {
				aa.oh.swingAt = DurationFromSeconds(aa.mh.SwingSpeed / 2)
			} else {
				if sim.RandomFloat("SwingResetWeapon") < 0.5 {
					aa.mh.swingAt = DurationFromSeconds(sim.RandomFloat("SwingResetDelay") * aa.mh.SwingSpeed / 2)
				} else {
					aa.oh.swingAt = DurationFromSeconds(sim.RandomFloat("SwingResetDelay") * aa.mh.SwingSpeed / 2)
				}
			}
		}

	}

	aa.ranged.swingAt = NeverExpires

	if aa.AutoSwingRanged {
		aa.ranged.updateSwingDuration(aa.ranged.unit.RangedSwingSpeed())
		aa.ranged.swingAt = 0
	}
}

func (aa *AutoAttacks) startPull(sim *Simulation) {
	if !aa.AutoSwingMelee && !aa.AutoSwingRanged {
		return
	}

	if aa.mh.unit.CurrentTarget == nil {
		return
	}

	if aa.anyEnabled() {
		return
	}

	if aa.AutoSwingMelee {
		if aa.mh.swingAt == NeverExpires {
			aa.mh.swingAt = 0
		}

		if aa.IsDualWielding {
			if aa.oh.swingAt == NeverExpires {
				aa.oh.swingAt = 0

				// Apply random delay of 0 - 50% swing time, to one of the weapons if dual wielding
				if aa.oh.unit.Type == EnemyUnit {
					aa.oh.swingAt = DurationFromSeconds(aa.mh.SwingSpeed / 2)
				} else {
					if sim.RandomFloat("SwingResetWeapon") < 0.5 {
						aa.mh.swingAt = DurationFromSeconds(sim.RandomFloat("SwingResetDelay") * aa.mh.SwingSpeed / 2)
					} else {
						aa.oh.swingAt = DurationFromSeconds(sim.RandomFloat("SwingResetDelay") * aa.mh.SwingSpeed / 2)
					}
				}
			}
			if aa.oh.IsInRange() {
				aa.oh.enabled = true
				aa.oh.addWeaponAttack(sim, aa.mh.curSwingSpeed)
			}

		}

		if aa.mh.IsInRange() {
			aa.mh.enabled = true
			aa.mh.addWeaponAttack(sim, aa.mh.unit.SwingSpeed())
		}
	}

	if aa.AutoSwingRanged {
		if aa.ranged.swingAt == NeverExpires {
			aa.ranged.swingAt = 0
		}
		if aa.ranged.IsInRange() {
			aa.ranged.enabled = true
			aa.ranged.addWeaponAttack(sim, aa.ranged.unit.RangedSwingSpeed())
		}

	}
}

func (wa *WeaponAttack) IsInRange() bool {
	return (wa.MinRange == 0. || wa.MinRange <= wa.unit.DistanceFromTarget) && (wa.MaxRange == 0. || wa.MaxRange >= wa.unit.DistanceFromTarget)
}

// Stops the auto swing action for the rest of the iteration. Used for pets
// after being disabled.
func (aa *AutoAttacks) CancelAutoSwing(sim *Simulation) {
	aa.CancelMeleeSwing(sim)
	aa.CancelRangedSwing(sim)
}

// Re-enables the auto swing action for the iteration
func (aa *AutoAttacks) EnableAutoSwing(sim *Simulation) {
	aa.EnableMeleeSwing(sim)
	aa.EnableRangedSwing(sim)
}

// Allow the auto swing action for the iteration
func (aa *AutoAttacks) AllowAutoSwing(sim *Simulation, isAllowed bool) {
	aa.mh.allowed = isAllowed
	aa.oh.allowed = isAllowed
	aa.ranged.allowed = isAllowed
}

func (aa *AutoAttacks) EnableMeleeSwing(sim *Simulation) {
	if !aa.AutoSwingMelee {
		return
	}

	if aa.mh.unit.CurrentTarget == nil {
		return
	}

	aa.mh.swingAt = max(aa.mh.swingAt, sim.CurrentTime, 0)
	if aa.mh.IsInRange() && !aa.mh.enabled {
		aa.mh.enabled = true
		aa.mh.addWeaponAttack(sim, aa.mh.unit.SwingSpeed())
	}

	if aa.IsDualWielding && !aa.oh.enabled {
		aa.oh.swingAt = max(aa.oh.swingAt, sim.CurrentTime, 0)
		if aa.oh.IsInRange() {
			aa.oh.enabled = true
			aa.oh.addWeaponAttack(sim, aa.mh.unit.SwingSpeed())
		}
	}

	if (!aa.IsDualWielding && aa.oh.enabled) {
		sim.removeWeaponAttack(&aa.oh)
		aa.oh.enabled = false
	}
}

func (aa *AutoAttacks) EnableRangedSwing(sim *Simulation) {
	if !aa.AutoSwingRanged || aa.ranged.enabled {
		return
	}

	if aa.ranged.unit.CurrentTarget == nil {
		return
	}

	aa.ranged.swingAt = max(aa.ranged.swingAt, sim.CurrentTime, 0)
	if aa.ranged.IsInRange() {
		aa.ranged.enabled = true
		aa.ranged.addWeaponAttack(sim, aa.ranged.unit.RangedSwingSpeed())
	}
}

func (aa *AutoAttacks) CancelMeleeSwing(sim *Simulation) {
	if !aa.AutoSwingMelee {
		return
	}

	if aa.mh.enabled {
		sim.removeWeaponAttack(&aa.mh)
		aa.mh.enabled = false
	}

	if aa.IsDualWielding && aa.oh.enabled {
		aa.oh.enabled = false
		sim.removeWeaponAttack(&aa.oh)
	}
}

func (aa *AutoAttacks) CancelRangedSwing(sim *Simulation) {
	if !aa.AutoSwingRanged || !aa.ranged.enabled {
		return
	}

	aa.ranged.enabled = false
	sim.removeWeaponAttack(&aa.ranged)
}

// The amount of time between two MH swings.
func (aa *AutoAttacks) MainhandSwingSpeed() time.Duration {
	return aa.mh.curSwingDuration
}

// The amount of time between two OH swings.
func (aa *AutoAttacks) OffhandSwingSpeed() time.Duration {
	return aa.oh.curSwingDuration
}

// The amount of time between two Ranged swings.
func (aa *AutoAttacks) RangedSwingSpeed() time.Duration {
	return aa.ranged.curSwingDuration - aa.RangedAuto().CastTime()
}

func (aa *AutoAttacks) UpdateSwingTimers(sim *Simulation) {
	if !aa.anyEnabled() {
		return
	}

	if aa.AutoSwingRanged && aa.ranged.enabled {
		aa.ranged.updateSwingDuration(aa.ranged.unit.RangedSwingSpeed())
		// ranged attack speed changes aren't applied mid-"swing"
	}

	if aa.AutoSwingMelee && aa.mh.enabled {
		oldSwingSpeed := aa.mh.curSwingSpeed
		aa.mh.updateSwingDuration(aa.mh.unit.SwingSpeed())
		f := oldSwingSpeed / aa.mh.curSwingSpeed

		if remainingSwingTime := aa.mh.swingAt - sim.CurrentTime; remainingSwingTime > 0 {
			aa.mh.swingAt = sim.CurrentTime + time.Duration(float64(remainingSwingTime)*f)
		}

		sim.rescheduleWeaponAttack(aa.mh.swingAt)

		if aa.IsDualWielding && aa.oh.enabled {
			aa.oh.updateSwingDuration(aa.mh.curSwingSpeed)

			if remainingSwingTime := aa.oh.swingAt - sim.CurrentTime; remainingSwingTime > 0 {
				aa.oh.swingAt = sim.CurrentTime + time.Duration(float64(remainingSwingTime)*f)
			}

			sim.rescheduleWeaponAttack(aa.oh.swingAt)
		}
	}
}

func (aa *AutoAttacks) ExtraMHAttackProc(sim *Simulation, attacks int32, actionID ActionID, spell *Spell) {
	if spell.Flags.Matches(SpellFlagBatchStopAttackMacro) {
		aa.StoreExtraMHAttack(sim, attacks, actionID, spell.ActionID)
	} else {
		aa.ExtraMHAttack(sim, attacks, actionID, spell.ActionID)
	}
}

// ExtraMHAttack should be used for all "extra attack" procs in Classic Era versions, including Wild Strikes and Hand of Justice. In vanilla, these procs don't actually grant a full extra attack, but instead just advance the MH swing timer.
func (aa *AutoAttacks) ExtraMHAttack(sim *Simulation, attacks int32, actionID ActionID, triggerAction ActionID) {
	if attacks == 0 {
		return
	}
	if sim.Log != nil {
		attacksText := Ternary(attacks == 1, "attack", "attacks")
		aa.mh.unit.Log(sim, "Gained %d extra main-hand %s from %s triggered by %s", attacks, attacksText, actionID, triggerAction)
	}

	aa.mh.swingAt = sim.CurrentTime
	aa.mh.spell.SetMetricsSplit(1)
	sim.rescheduleWeaponAttack(aa.mh.swingAt)
	aa.mh.extraAttacksPending += attacks
}

func (aa *AutoAttacks) StoreExtraMHAttack(sim *Simulation, attacks int32, actionID ActionID, triggerAction ActionID) {
	if attacks == 0 {
		return
	}

	if !aa.mh.extraAttacksAura.IsActive() {
		aa.mh.extraAttacksAura.Activate(sim)
	}

	aa.mh.extraAttacksAura.AddStacks(sim, attacks)

	if sim.Log != nil {
		attacksText := Ternary(attacks == 1, "attack", "attacks")
		aa.mh.unit.Log(sim, "Stored %d extra main-hand %s from %s triggered by %s, %d total attacks stored", attacks, attacksText, actionID, triggerAction, aa.mh.extraAttacksAura.GetStacks())
	}
}

// ExtraMHAttack should be used for all "extra attack" procs in Classic Era versions, including Wild Strikes and Hand of Justice. In vanilla, these procs don't actually grant a full extra attack, but instead just advance the MH swing timer.
func (aa *AutoAttacks) ExtraOHAttack(sim *Simulation, attacks int32, actionID ActionID, triggerAction ActionID) {
	if attacks == 0 {
		return
	}
	if sim.Log != nil {
		attacksText := Ternary(attacks == 1, "attack", "attacks")
		aa.oh.unit.Log(sim, "Gained %d extra off-hand %s from %s triggered by %s", attacks, attacksText, actionID, triggerAction)
	}
	aa.oh.swingAt = sim.CurrentTime + SpellBatchWindow
	aa.oh.spell.SetMetricsSplit(1)
	sim.rescheduleWeaponAttack(aa.oh.swingAt)
	aa.oh.extraAttacksPending += attacks
}

// ExtraRangedAttack should be used for all "extra ranged attack" procs in Classic Era versions, including Hand of Injustice. In vanilla, these procs don't actually grant a full extra attack, but instead just advance the Ranged swing timer.
// Note that Hand of Injustice doesn't seem to reset the swing timer however.
func (aa *AutoAttacks) ExtraRangedAttack(sim *Simulation, attacks int32, actionID ActionID, triggerAction ActionID) {
	if sim.Log != nil {
		attacksText := Ternary(attacks == 1, "attack", "attacks")
		aa.mh.unit.Log(sim, "Gained %d extra ranged %s from %s triggered by %s", attacks, attacksText, actionID, triggerAction)
	}
	aa.ranged.swingAt = sim.CurrentTime + SpellBatchWindow
	aa.ranged.spell.SetMetricsSplit(1)
	sim.rescheduleWeaponAttack(aa.ranged.swingAt)
	aa.ranged.extraAttacks += attacks
}

func (aa *AutoAttacks) StoreExtraRangedAttack(sim *Simulation, attacks int32, actionID ActionID, triggerAction ActionID) {
	if attacks == 0 {
		return
	}

	if !aa.ranged.extraAttacksAura.IsActive() {
		aa.ranged.extraAttacksAura.Activate(sim)
	}

	aa.ranged.extraAttacksAura.AddStacks(sim, attacks)

	if sim.Log != nil {
		attacksText := Ternary(attacks == 1, "attack", "attacks")
		aa.ranged.unit.Log(sim, "Stored %d extra main-hand %s from %s triggered by %s, %d total attacks stored", attacks, attacksText, actionID, triggerAction, aa.mh.extraAttacksAura.GetStacks())
	}
}

// StopMeleeUntil should be used whenever a non-melee spell is cast. It stops melee, then restarts it
// at end of cast, but with a reset swing timer (as if swings had just landed).
func (aa *AutoAttacks) StopMeleeUntil(sim *Simulation, readyAt time.Duration, desyncOH bool) {
	if !aa.AutoSwingMelee { // if not auto swinging, don't auto restart.
		return
	}

	aa.mh.swingAt = readyAt + aa.mh.curSwingDuration
	sim.rescheduleWeaponAttack(aa.mh.swingAt)

	if aa.IsDualWielding {
		aa.oh.swingAt = readyAt + aa.oh.curSwingDuration
		if desyncOH {
			// Used by warrior to desync offhand after unglyphed Shattering Throw.
			aa.oh.swingAt += aa.oh.curSwingDuration / 2
		}
		sim.rescheduleWeaponAttack(aa.oh.swingAt)
	}
}

func (aa *AutoAttacks) StopRangedUntil(sim *Simulation, readyAt time.Duration) {
	if !aa.AutoSwingRanged { // if not auto swinging, don't auto restart.
		return
	}

	aa.ranged.swingAt = readyAt + aa.ranged.curSwingDuration
	sim.rescheduleWeaponAttack(aa.ranged.swingAt)
}

// Delays all swing timers for the specified amount. Only used by Slam.
func (aa *AutoAttacks) DelayMeleeBy(sim *Simulation, delay time.Duration) {
	if delay <= 0 {
		return
	}

	aa.mh.swingAt += delay
	sim.rescheduleWeaponAttack(aa.mh.swingAt)

	if aa.IsDualWielding {
		aa.oh.swingAt += delay
		sim.rescheduleWeaponAttack(aa.oh.swingAt)
	}
}

func (aa *AutoAttacks) DelayRangedUntil(sim *Simulation, readyAt time.Duration) {
	if readyAt <= aa.ranged.swingAt {
		return
	}

	aa.ranged.swingAt = readyAt
	sim.rescheduleWeaponAttack(aa.ranged.swingAt)
}

// Returns the time at which the next melee attack will occur.
func (aa *AutoAttacks) NextAttackAt() time.Duration {
	return min(aa.mh.swingAt, aa.oh.swingAt)
}

// Returns the time at which the next attack will occur.
func (aa *AutoAttacks) NextAnyAttackAt() time.Duration {
	return min(min(aa.mh.swingAt, aa.oh.swingAt), aa.ranged.swingAt)
}

// Returns the time at which the next ranged attack will occur.
func (aa *AutoAttacks) NextRangedAttackAt() time.Duration {
	return aa.ranged.swingAt
}

type DynamicProcManager struct {
	ppm             float64
	fixedProcChance float64
	procMasks       []ProcMask
	procChances     []float64

	// For feral druids, certain PPM effects use their equipped weapon speed
	// instead of their paw attack speed.
	mhSpecialProcChance float64
	ohSpecialProcChance float64
}

// Returns whether the effect procced.
func (dpm *DynamicProcManager) Proc(sim *Simulation, procMask ProcMask, label string) bool {
	for i, m := range dpm.procMasks {
		if m.Matches(procMask) {
			return sim.RandomFloat(label) < dpm.procChances[i]
		}
	}
	return false
}

// Returns whether the effect procced.
// This is different from Proc() in that yellow melee hits use a proc chance based on the equipped
// weapon speed rather than the base attack speed. This distinction matters for feral druids.
func (dpm *DynamicProcManager) ProcWithWeaponSpecials(sim *Simulation, procMask ProcMask, label string) bool {
	if procMask.Matches(ProcMaskMeleeMHSpecial) {
		return sim.RandomFloat(label) < dpm.mhSpecialProcChance
	} else if procMask.Matches(ProcMaskMeleeOHSpecial) {
		return sim.RandomFloat(label) < dpm.ohSpecialProcChance
	} else {
		return dpm.Proc(sim, procMask, label)
	}
}

func (dpm *DynamicProcManager) Chance(procMask ProcMask) float64 {
	for i, m := range dpm.procMasks {
		if m.Matches(procMask) {
			return dpm.procChances[i]
		}
	}
	return 0
}

func (dpm *DynamicProcManager) GetPPM() float64 {
	return dpm.ppm
}

// PPMManager for static ProcMasks
func (aa *AutoAttacks) NewPPMManager(ppm float64, procMask ProcMask) *DynamicProcManager {
	dpm := aa.newDynamicProcManager(ppm, 0, procMask)

	if aa.character != nil {
		aa.character.RegisterItemSwapCallback(AllWeaponSlots(), func(sim *Simulation, slot proto.ItemSlot, _ bool) {
			dpm = aa.character.AutoAttacks.newDynamicProcManager(ppm, 0, procMask)
		})
	}

	return &dpm
}

// PPMManager for static ProcMasks and no item swap callback
func (aa *AutoAttacks) NewStaticPPMManager(ppm float64, procMask ProcMask) *DynamicProcManager {
	dpm := aa.newDynamicProcManager(ppm, 0, procMask)

	return &dpm
}

// Dynamic Proc Manager for dynamic ProcMasks on weapon enchants
func (aa *AutoAttacks) NewDynamicProcManagerForEnchant(effectID int32, ppm float64, fixedProcChance float64) *DynamicProcManager {
	return aa.newDynamicProcManagerWithDynamicProcMask(ppm, fixedProcChance, func() ProcMask {
		return aa.character.getCurrentProcMaskForWeaponEnchant(effectID)
	})
}

// Dynamic Proc Manager for dynamic ProcMasks on weapon effects
func (aa *AutoAttacks) NewDynamicProcManagerForWeaponEffect(itemID int32, ppm float64, fixedProcChance float64) *DynamicProcManager {
	return aa.newDynamicProcManagerWithDynamicProcMask(ppm, fixedProcChance, func() ProcMask {
		return aa.character.getCurrentProcMaskForWeaponEffect(itemID)
	})
}

func (aa *AutoAttacks) newDynamicProcManagerWithDynamicProcMask(ppm float64, fixedProcChance float64, procMaskFn func() ProcMask) *DynamicProcManager {
	dpm := aa.newDynamicProcManager(ppm, fixedProcChance, procMaskFn())

	if aa.character != nil {
		aa.character.RegisterItemSwapCallback(AllWeaponSlots(), func(sim *Simulation, slot proto.ItemSlot, _ bool) {
			dpm = aa.character.AutoAttacks.newDynamicProcManager(ppm, fixedProcChance, procMaskFn())
		})
	}

	return &dpm

}

func (aa *AutoAttacks) newDynamicProcManager(ppm float64, fixedProcChance float64, procMask ProcMask) DynamicProcManager {
	if (ppm != 0) && (fixedProcChance != 0) {
		panic("Cannot simultaneously specify both a ppm and a fixed proc chance!")
	}

	if !aa.AutoSwingMelee && !aa.AutoSwingRanged {
		return DynamicProcManager{}
	}

	dpm := DynamicProcManager{ppm: ppm, fixedProcChance: fixedProcChance, procMasks: make([]ProcMask, 0, 2), procChances: make([]float64, 0, 2)}

	mergeOrAppend := func(speed float64, mask ProcMask) {
		if speed == 0 || mask == 0 {
			return
		}

		if i := slices.Index(dpm.procChances, speed); i != -1 {
			dpm.procMasks[i] |= mask
			return
		}

		dpm.procMasks = append(dpm.procMasks, mask)
		dpm.procChances = append(dpm.procChances, speed)
	}

	mergeOrAppend(aa.mh.SwingSpeed, procMask&^ProcMaskRanged&^ProcMaskMeleeOH) // "everything else", even if not explicitly flagged MH
	mergeOrAppend(aa.oh.SwingSpeed, procMask&ProcMaskMeleeOH)
	mergeOrAppend(aa.ranged.SwingSpeed, procMask&ProcMaskRanged)

	for i := range dpm.procChances {
		if fixedProcChance != 0 {
			dpm.procChances[i] = fixedProcChance
		} else {
			dpm.procChances[i] *= ppm / 60
		}
	}

	character := aa.mh.agent.GetCharacter()
	if character != nil {
		if procMask.Matches(ProcMaskMeleeMH) {
			if mhWeapon := character.GetMHWeapon(); mhWeapon != nil {
				if fixedProcChance != 0 {
					dpm.mhSpecialProcChance = fixedProcChance
				} else {
					dpm.mhSpecialProcChance = mhWeapon.SwingSpeed * ppm / 60
				}
			}
		}
		if procMask.Matches(ProcMaskMeleeOH) {
			if ohWeapon := character.GetOHWeapon(); ohWeapon != nil {
				if fixedProcChance != 0 {
					dpm.ohSpecialProcChance = fixedProcChance
				} else {
					dpm.ohSpecialProcChance = ohWeapon.SwingSpeed * ppm / 60
				}
			}
		}
	}
	return dpm
}

// Returns whether a PPM-based effect procced.
// Using NewPPMManager() is preferred; this function should only be used when
// the attacker is not known at initialization time.
func (aa *AutoAttacks) PPMProc(sim *Simulation, ppm float64, procMask ProcMask, label string, spell *Spell) bool {
	if !aa.AutoSwingMelee && !aa.AutoSwingRanged {
		return false
	}

	switch {
	case spell.ProcMask.Matches(procMask &^ ProcMaskMeleeOH &^ ProcMaskRanged):
		return sim.RandomFloat(label) < ppm*aa.mh.SwingSpeed/60.0
	case spell.ProcMask.Matches(procMask & ProcMaskMeleeOH):
		return sim.RandomFloat(label) < ppm*aa.oh.SwingSpeed/60.0
	case spell.ProcMask.Matches(procMask & ProcMaskRanged):
		return sim.RandomFloat(label) < ppm*aa.ranged.SwingSpeed/60.0
	}
	return false
}

func (unit *Unit) applyParryHaste() {
	if !unit.PseudoStats.ParryHaste || !unit.AutoAttacks.AutoSwingMelee {
		return
	}

	unit.RegisterAura(Aura{
		Label:    "Parry Haste",
		Duration: NeverExpires,
		OnReset: func(aura *Aura, sim *Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitTaken: func(aura *Aura, sim *Simulation, spell *Spell, result *SpellResult) {
			if !result.DidParry() {
				return
			}

			currentSwingTime := aura.Unit.AutoAttacks.mh.swingAt - aura.Unit.AutoAttacks.mh.lastSwingAt
			swingSpeed := aura.Unit.AutoAttacks.mh.curSwingDuration
			minRemainingTime := time.Duration(float64(swingSpeed) * 0.2) // 20% of Swing Speed
			defaultReduction := minRemainingTime * 2                     // 40% of Swing Speed

			if currentSwingTime <= minRemainingTime {
				return
			}

			newReadyAt := max(aura.Unit.AutoAttacks.mh.swingAt-defaultReduction, aura.Unit.AutoAttacks.mh.lastSwingAt+minRemainingTime)
			parryHasteReduction := newReadyAt - aura.Unit.AutoAttacks.mh.swingAt
			if sim.Log != nil {
				aura.Unit.Log(sim, "MH Swing reduced by %s due to parry haste, will now occur at %s", parryHasteReduction, newReadyAt)
			}

			aura.Unit.AutoAttacks.mh.swingAt = newReadyAt
			sim.rescheduleWeaponAttack(newReadyAt)
		},
	})
}
