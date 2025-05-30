package vanilla

import (
	"fmt"
	"time"

	"github.com/wowsims/sod/sim/common/itemhelpers"
	"github.com/wowsims/sod/sim/core"
	"github.com/wowsims/sod/sim/core/proto"
	"github.com/wowsims/sod/sim/core/stats"
)

// Ordered by ID
const (
	Destiny                         = 647
	ShortswordOfVengeance           = 754
	FieryWarAxe                     = 870
	Bloodrazor                      = 809
	HammerOfTheNorthernWind         = 810
	FlurryAxe                       = 871
	SkullflameShield                = 1168
	Nightblade                      = 1982
	Shadowblade                     = 2163
	GutRipper                       = 2164
	HandOfEdwardTheOdd              = 2243
	BowOfSearingArrows              = 2825
	Gutwrencher                     = 5616
	Ravager                         = 7717
	HanzoSword                      = 8190
	TheJackhammer                   = 9423
	PendulumOfDoom                  = 9425
	BloodletterScalpel              = 9511
	TheHandOfAntusul                = 9639
	GryphonRidersStormhammer        = 9651
	Firebreather                    = 10797
	VilerendSlicer                  = 11603
	HookfangShanker                 = 11635
	Naglering                       = 11669
	LinkensSwordOfMastery           = 11902
	SearingNeedle                   = 12531
	PipsSkinner                     = 12709
	ArcaniteChampion                = 12790
	MasterworkStormhammer           = 12794
	Frostguard                      = 12797
	SerpentSlicer                   = 13035
	TheNeedler                      = 13060
	SealOfTheDawn                   = 13209
	CloudkeeperLegplates            = 14554
	JoonhosMercy                    = 17054
	Deathbringer                    = 17068
	ViskagTheBloodletter            = 17075
	ThrashBlade                     = 17705
	SatyrsLash                      = 17752
	MarkOfTheChosen                 = 17774
	ForceReactiveDisk               = 18168
	RazorGauntlets                  = 18326
	Nightfall                       = 19169
	EbonHand                        = 19170
	RuneOfTheDawn                   = 19812
	ZandalariHeroBadge              = 19948
	ZandalariHeroMedallion          = 19949
	ZandalariHeroCharm              = 19950
	BlisteringRagehammer            = 220569 // 10626
	SulfurasHandOfRagnaros          = 227683 // 17182
	SulfuronHammer                  = 227684 // 17193
	TemperedBlackAmnesty            = 227832 // 19166
	EbonFist                        = 227842
	HardenedFrostguard              = 227887
	FlameWrath                      = 227934 // 11809
	LordGeneralsSword               = 227940 // 11817
	WraithScythe                    = 227941
	SecondWind                      = 227967 // 11819
	BurstOfKnowledge                = 227972
	Ironfoe                         = 227991 // 11684
	EbonHiltOfMarduk                = 227993 // 14576
	FrightskullShaft                = 227994 // 14531
	BarovianFamilySword             = 227997 // 14541
	Frightalon                      = 228015 // 14024
	HeadmastersCharge               = 228022 // 13937
	GravestoneWarAxe                = 228029 // 13983
	FiendishMachete                 = 228056 // 18310
	RefinedArcaniteChampion         = 228125
	TalismanOfEphemeralPower        = 228255 // 18820
	DrillborerDisk                  = 228266 // 17066
	GutgoreRipper                   = 228267 // 17071
	Shadowstrike                    = 228272 // 17074
	Thunderstrike                   = 228273 // 17223
	BonereaversEdge                 = 228288 // 17076
	BonereaversEdgeMolten           = 228461
	EssenceOfThePureFlame           = 228293 // 18815
	PerditionsBlade                 = 228296 // 18816
	Typhoon                         = 228347 // 18542
	EskhandarsLeftClaw              = 228349 // 18202
	EskhandarsRightClaw             = 228350 // 18203
	BlazefuryMedallion              = 228354 // 17111
	EmpyreanDemolisher              = 228397 // 17112
	DreadbladeOfTheDestructor       = 228410
	PerditionsBladeMolten           = 228511
	SkullforgeReaver                = 228542 // 13361
	RunebladeOfBaronRivendare       = 228543 // 13505
	HeartOfWyrmthalak               = 228599 // 22321
	Venomspitter                    = 228573 // 13183
	SmolderwebsEye                  = 228576 // 13213
	Chillpike                       = 228586 // 13148
	FangOfTheCrystalSpider          = 228592 // 13218
	BlackhandDoomsaw                = 228603 // 12583
	BlackbladeOfShahram             = 228606 // 12592
	SeepingWillow                   = 228666 // 12969
	DraconicInfusedEmblem           = 228678 // 22268
	QuelSerrar                      = 228679 // 18348
	DrillborerDiskMolten            = 228702
	HandOfJustice                   = 228722 // 11815
	Felstriker                      = 228757 // 12590
	GutgoreRipperMolten             = 229372
	EskhandarsRightClawMolten       = 229379
	Thunderfury                     = 230224 // 19019
	TheUntamedBlade                 = 230242 // 19334
	DrakeTalonCleaver               = 230271 // 19353
	JekliksCrusher                  = 230911 // 19918
	ZulianSlicer                    = 230930 // 19901
	HalberdOfSmiting                = 230991 // 19874
	NatPaglesBrokenReel             = 231271 // 19947
	TigulesHarpoon                  = 231272 // 19946
	GrileksCarver                   = 231273 // 19962
	GrileksGrinder                  = 231274 // 19961
	PitchforkOfMadness              = 231277 // 19963
	GrileksCarverBloodied           = 231846
	GrileksGrinderBloodied          = 231847
	TigulesHarpoonBloodied          = 231849
	JekliksCrusherBloodied          = 231861
	PitchforkOfMadnessBloodied      = 231864
	HalberdOfSmitingBloodied        = 231870
	ZulianSlicerBloodied            = 231876
	DrakeTalonCleaverShadowflame    = 232562
	TheUntamedBladeShadowflame      = 232566
	ScarabBrooch                    = 233601 // 21625
	KalimdorsRevenge                = 233621
	JomGabbar                       = 233627 // 23570
	TheBurrowersShield              = 233628 // 23558
	NeretzekBloodDrinker            = 233647
	ManslayerOfTheQiraji            = 234067
	EyeOfMoam                       = 234080 // 21473
	FetishOfChitinousSpikes         = 234092 // 21488
	DarkmoonCardHeroism             = 234176 // 19287
	DarkmoonCardBlueDragon          = 234177 // 19288
	DarkmoonCardMaelstrom           = 234178 // 19289
	Earthstrike                     = 234462
	WrathOfCenarius                 = 234463 // 21190
	KalimdorsRevengeVoidTouched     = 234981
	NeretzekBloodDrinkerVoidTouched = 234987
	ManslayerOfTheQirajiVoidTouched = 234990
	MisplacedServoArm               = 236221 // 23221
	KissOfTheSpider                 = 236268 // 22954
	WarmthOfForgiveness             = 236320 // 23027
	TheRestrainedEssenceOfSapphiron = 236331 // 23046
	SlayersCrest                    = 236334 // 23041
	GlyphOfDeflection               = 236337 // 23040
	MarkOfTheChampionSpell          = 236351 // 23207
	MarkOfTheChampionPhys           = 236352 // 23206
	TalismanOfAscendance            = 237283 // 22678
	CorruptedAshbringer             = 241081 // 22691
	CorruptedAshbringerLego         = 239301
)

func init() {
	core.AddEffectsToTest = false

	// ! Please keep items ordered alphabetically within a given category !

	///////////////////////////////////////////////////////////////////////////
	//                                 Weapons
	///////////////////////////////////////////////////////////////////////////

	// https://www.wowhead.com/classic/spell=16916/strength-of-the-champion
	// Chance on hit: Heal self for 270 to 450 and Increases Strength by 120 for 30 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcAura(ArcaniteChampion, "Arcanite Champion", 1.0, StrengthOfTheChampionAura)

	// https://www.wowhead.com/classic/item=227997/barovian-family-sword
	// Chance on hit: Deals 30 Shadow damage every 3 sec for 15 sec. All damage done is then transferred to the caster.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(BarovianFamilySword, "Barovian Family Sword", 0.5, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 18652}

		// Keep track of damage taken by each enemy
		enemyDamageTaken := map[int32]float64{}
		for _, target := range character.Env.Encounter.TargetUnits {
			enemyDamageTaken[target.UnitIndex] = 0
		}

		healthMetrics := character.NewHealthMetrics(actionID)

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot,

			Dot: core.DotConfig{
				NumberOfTicks: 5,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Siphon Health (Barovian Family Sword)",
				},
				OnSnapshot: func(_ *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					enemyDamageTaken[target.UnitIndex] = 0
					dot.Snapshot(target, 30, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					result := dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
					enemyDamageTaken[target.UnitIndex] += result.Damage
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				enemyDamageTaken[target.UnitIndex] = 0
				spell.Dot(target).Apply(sim)
			},
		})

		// The healing is applied at the end of the DoT and can crit according to old comments
		for _, dot := range spell.Dots() {
			if dot != nil {
				unit := dot.Unit
				dot.ApplyOnExpire(func(aura *core.Aura, sim *core.Simulation) {
					// TODO: This may not be quite correct but it's close enough
					result := spell.CalcDamage(sim, unit, enemyDamageTaken[unit.UnitIndex], spell.OutcomeHealingCrit)
					character.GainHealth(sim, result.Damage, healthMetrics)
				})
			}
		}

		return spell
	})

	// https://www.wowhead.com/classic/item=228606/blackblade-of-shahram
	// Chance on hit: Summons the infernal spirit of Shahram.
	// Summons an NPC "Shahram" who has an equal chance to cast one of 6 spells:
	// Curse of Shahram: -50% movement speed and -25% attack speed on all enemies within 10 yards of Shahram for 10 seconds.
	// Might of Shahram: 5-second stun on all enemies within 10 yards of Shahram.
	// Fist of Shahram: +30% Melee Attack Speed for all party members within 30 yards of Shahram for 8 seconds.
	// Blessing of Shahram: Restores 50 health and mana every 5 seconds for all party members within 30 yards of Shahram for 20 seconds. The Healing portion of this effect scales at 100% of self-healing buffs such as Amplify Magic.
	// Will of Shahram: +50 all stats for yourself for 20 seconds.
	// Flames of Shahram: Deals 100-150 Fire damage to all enemies within 10 yards of Shahram. Damage scales at 100% with +spelldmg debuffs placed on enemies such as Flame Buffet.
	//
	// Implementing this without the guardian as it seems to just cast a spell and depart and guardians are expensive
	core.NewItemEffect(BlackbladeOfShahram, func(agent core.Agent) {
		character := agent.GetCharacter()

		curseOfShahramAuras := character.NewEnemyAuraArray(func(target *core.Unit, _ int32) *core.Aura {
			aura := target.RegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16597},
				Label:    "Curse of Shahram",
				Duration: time.Second * 10,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.MultiplyAttackSpeed(sim, 1/1.25)
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.MultiplyAttackSpeed(sim, 1.25)
				},
			})
			core.AtkSpeedReductionEffect(aura, 1.25)
			return aura
		})
		curseOfShahram := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16597},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				curseOfShahramAuras.Get(target).Activate(sim)
			},
		})

		mightOfShahramAuras := character.NewEnemyAuraArray(func(target *core.Unit, _ int32) *core.Aura {
			return target.RegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16600},
				Label:    "Might of Shahram",
				Duration: time.Second * 5,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.PseudoStats.Stunned = true
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.PseudoStats.Stunned = false
				},
			})
		})

		mightOfShahram := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16600},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					mightOfShahramAuras.Get(aoeTarget).Activate(sim)
				}
			},
		})

		// This isn't explicit in-game but using a safe value that will likely never be hit
		numFistOfShahramAuras := 8
		fistOfShahramAuras := []*core.Aura{}
		for i := 0; i < numFistOfShahramAuras; i++ {
			fistOfShahramAuras = append(fistOfShahramAuras, character.RegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16601},
				Label:    fmt.Sprintf("Fist of Shahram (%d)", i),
				Duration: time.Second * 8,
			}).AttachMultiplyAttackSpeed(&character.Unit, 1.3))
		}

		fistOfShahram := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16601},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for i := 0; i < numFistOfShahramAuras; i++ {
					if aura := fistOfShahramAuras[i]; !aura.IsActive() {
						aura.Activate(sim)
						break
					}
				}
			},
		})

		blessingOfShahramManaMetrics := character.NewPartyManaMetrics(core.ActionID{SpellID: 16599})
		blessingOfShahram := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16599},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagIgnoreAttackerModifiers | core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,
			Hot: core.DotConfig{
				Aura: core.Aura{
					Label: "Blessing of Shahram",
				},
				NumberOfTicks: 4,
				TickLength:    time.Second * 5,
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, _ bool) {
					dot.SnapshotBaseDamage = 50
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotHealing(sim, target, dot.OutcomeTick)
					if target.HasManaBar() {
						target.AddMana(sim, 50, blessingOfShahramManaMetrics[target.UnitIndex])
					}
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, agent := range character.Party.PlayersAndPets {
					spell.Hot(&agent.GetCharacter().Unit).Apply(sim)
				}
			},
		})

		// This isn't explicit in-game but using a safe value that will likely never be hit
		numWillOfShahramAuras := 8
		willOfShahramAuras := []*core.Aura{}
		willOfShahramStats := stats.Stats{
			stats.Agility:   50,
			stats.Intellect: 50,
			stats.Stamina:   50,
			stats.Spirit:    50,
			stats.Strength:  50,
		}

		for i := 0; i < numWillOfShahramAuras; i++ {
			willOfShahramAuras = append(willOfShahramAuras, character.RegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16598},
				Label:    fmt.Sprintf("Will of Shahram (%d)", i),
				Duration: time.Second * 20,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					character.AddStatsDynamic(sim, willOfShahramStats)
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					character.AddStatsDynamic(sim, willOfShahramStats.Invert())
				},
			}))
		}

		willOfShahram := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16598},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for i := 0; i < numWillOfShahramAuras; i++ {
					if aura := willOfShahramAuras[i]; !aura.IsActive() {
						aura.Activate(sim)
						break
					}
				}
			},
		})

		flamesOfShahram := character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 16596},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagIgnoreAttackerModifiers | core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, 90, spell.OutcomeMagicCrit)
				}
			},
		})

		castableSpells := []*core.Spell{curseOfShahram, mightOfShahram, fistOfShahram, blessingOfShahram, willOfShahram, flamesOfShahram}
		aura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Summon Shahram",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			PPM:               1,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				spellIdx := int32(sim.Roll(0, 6))
				castableSpells[spellIdx].Cast(sim, result.Target)
			},
		})

		character.ItemSwap.RegisterProc(BlackbladeOfShahram, aura)
	})

	// https://www.wowhead.com/classic/item=228603/blackhand-doomsaw
	// Chance on hit: Wounds the target for 324 to 540 damage.
	// TODO: Proc rate based on the original item
	itemhelpers.CreateWeaponCoHProcDamage(BlackhandDoomsaw, "Blackhand Doomsaw", 0.4, 16549, core.SpellSchoolPhysical, 324, 216, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=220569/blistering-ragehammer
	// Chance on hit: Increases damage done by 20 and attack speed by 5% for 15 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcAura(BlisteringRagehammer, "Blistering Ragehammer", 1.0, EnrageAura446327)

	itemhelpers.CreateWeaponCoHProcDamage(BloodletterScalpel, "Bloodletter Scalpel", 1.0, 18081, core.SpellSchoolPhysical, 60, 10, 0, core.DefenseTypeMelee)

	itemhelpers.CreateWeaponProcSpell(Bloodrazor, "Bloodrazor", 1.0, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 17504},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Rend (Bloodrazor)",
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, 12, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=228288/bonereavers-edge
	// https://www.wowhead.com/classic/item=228461/bonereavers-edge
	// Chance on hit: Your attacks ignore 700 of your enemies' armor for 10 sec. This effect stacks up to 3 times.
	itemhelpers.CreateWeaponProcSpell(BonereaversEdge, "Bonereaver's Edge", 2.0, bonereaversEdgeEffect)
	itemhelpers.CreateWeaponProcSpell(BonereaversEdgeMolten, "Bonereaver's Edge (Molten)", 2.0, bonereaversEdgeEffect)

	itemhelpers.CreateWeaponProcSpell(BowOfSearingArrows, "Bow of Searing Arrows", 3.35, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 29638},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeRanged,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				dmg := sim.Roll(18, 26)
				spell.CalcAndDealDamage(sim, target, dmg, spell.OutcomeRangedCritOnly)
			},
		})
	})

	// https://www.wowhead.com/classic/item=228586/chillpike
	// Chance on hit: Blasts a target for 160 to 250 Frost damage.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponCoHProcDamage(Chillpike, "Chillpike", 1.0, 19260, core.SpellSchoolFrost, 160, 90, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=241081/corrupted-ashbringer
	// Chance on hit: Steals 475 to 525 life from target enemy.
	// PPM confirmed 2.0
	itemhelpers.CreateWeaponProcSpell(CorruptedAshbringer, "Corrupted Ashbringer", 2.0, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 1220711}
		healthMetrics := character.NewHealthMetrics(actionID)

		return character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellProc | core.ProcMaskSpellDamageProc,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, sim.Roll(475, 525), spell.OutcomeMagicHit)
				character.GainHealth(sim, result.Damage, healthMetrics)
			},
		})
	})

	// https://www.wowhead.com/classic/item=239301/corrupted-ashbringer
	// Chance on hit: Steals 475 to 525 life from up to 5 enemies, granting 30 Strength or Agility per enemy siphoned, stacking up to 5 times.
	// PPM confirmed 5.0
	core.NewItemEffect(CorruptedAshbringerLego, func(agent core.Agent) {
		character := agent.GetCharacter()

		targetCount := int(min(character.Env.GetNumTargets(), 5))
		chainMultiplier := 1.1 // EffectChainAmplitude: 1.1
		wrathSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 1231377},
			SpellSchool: core.SpellSchoolHoly,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellProc | core.ProcMaskSpellDamageProc,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BonusCoefficient: 0.1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				origMult := spell.GetDamageMultiplier()

				results := make([]*core.SpellResult, targetCount)
				for i := 0; i < targetCount; i++ {
					results[i] = spell.CalcDamage(sim, target, sim.Roll(1080, 1320), spell.OutcomeMagicCrit)
					target = sim.Environment.NextTargetUnit(target)
					spell.ApplyMultiplicativeDamageBonus(chainMultiplier)
				}

				spell.SetMultiplicativeDamageBonus(origMult)

				for _, result := range results {
					spell.DealDamage(sim, result)
				}
			},
		})

		wrathOfTheAshbringerAura := character.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 1239620},
			Label:    "Wrath of the Ashbringer",
			Duration: time.Millisecond * 100,
		})

		activateAshbringerAura := func(sim *core.Simulation, corruptedAura *core.Aura, ashbringerAura *core.Aura) {
			corruptedAura.Deactivate(sim)
			ashbringerAura.Activate(sim)

			// Wrath of the Ashbringer is always triggered upon activation of the Ashbringer aura
			wrathOfTheAshbringerAura.Activate(sim)
			core.StartDelayedAction(sim, core.DelayedActionOptions{
				DoAt: sim.CurrentTime + core.SpellBatchWindow,
				OnAction: func(sim *core.Simulation) {
					wrathSpell.Cast(sim, character.CurrentTarget)
				},
			})
		}

		var crusaderAuras []*core.Aura

		checkAshbringerAuraActivation := func(sim *core.Simulation, newStacks int32, corruptedAura *core.Aura, ashbringerAura *core.Aura) {
			if newStacks != 5 || len(crusaderAuras) == 0 {
				return
			}

			for _, aura := range crusaderAuras {
				if aura.IsActive() {
					// Delay by one batch to allow for damage to be dealt
					core.StartDelayedAction(sim, core.DelayedActionOptions{
						DoAt: sim.CurrentTime + core.SpellBatchWindow,
						OnAction: func(sim *core.Simulation) {
							activateAshbringerAura(sim, corruptedAura, ashbringerAura)
						},
					})
					return
				}
			}
		}

		ashbrinigerAgilityAura := character.NewTemporaryStatsAura("Ashbringer (Agility)", core.ActionID{SpellID: 1231409}, stats.Stats{stats.Agility: 200}, time.Second*30)
		corruptedAgilityAura := character.RegisterAura(core.Aura{
			ActionID:  core.ActionID{SpellID: 1231406},
			Label:     "Corrupted Agility",
			MaxStacks: 5,
			Duration:  time.Second * 20,
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks, newStacks int32) {
				character.AddStatDynamic(sim, stats.Agility, 30*float64(newStacks-oldStacks))
				checkAshbringerAuraActivation(sim, newStacks, aura, ashbrinigerAgilityAura)
			},
		})

		ashbrinigerStrengthAura := character.NewTemporaryStatsAura("Ashbringer (Strength)", core.ActionID{SpellID: 1230596}, stats.Stats{stats.Strength: 200}, time.Second*30)
		corruptedStrengthAura := character.RegisterAura(core.Aura{
			ActionID:  core.ActionID{SpellID: 1231339},
			Label:     "Corrupted Strength",
			MaxStacks: 5,
			Duration:  time.Second * 20,
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks, newStacks int32) {
				character.AddStatDynamic(sim, stats.Strength, 30*float64(newStacks-oldStacks))
				checkAshbringerAuraActivation(sim, newStacks, aura, ashbrinigerStrengthAura)
			},
		})

		character.OnAuraRegistered(func(aura *core.Aura) {
			if aura.Tag != "Crusader" {
				return
			}

			crusaderAuras = append(crusaderAuras, aura)
			aura.ApplyOnGain(func(aura *core.Aura, sim *core.Simulation) {
				if corruptedAgilityAura.IsActive() && corruptedAgilityAura.GetStacks() == 5 && !ashbrinigerAgilityAura.IsActive() {
					activateAshbringerAura(sim, corruptedAgilityAura, ashbrinigerAgilityAura)
				} else if corruptedStrengthAura.IsActive() && corruptedStrengthAura.GetStacks() == 5 && !ashbrinigerStrengthAura.IsActive() {
					activateAshbringerAura(sim, corruptedStrengthAura, ashbrinigerStrengthAura)
				}
			})
		})

		consumptionActionID := core.ActionID{SpellID: 1231330}
		healthMetrics := character.NewHealthMetrics(consumptionActionID)
		consumptionSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    consumptionActionID,
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellProc | core.ProcMaskSpellDamageProc,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				results := make([]*core.SpellResult, targetCount)
				for i := 0; i < targetCount; i++ {
					results[i] = spell.CalcDamage(sim, target, sim.Roll(475, 525), spell.OutcomeAlwaysHit)
					target = sim.Environment.NextTargetUnit(target)
				}

				// Consumption has the flag "Missile Speed is Delay (in sec)", with a MissileSpeed of 1
				core.StartDelayedAction(sim, core.DelayedActionOptions{
					DoAt: sim.CurrentTime + time.Second,
					OnAction: func(sim *core.Simulation) {
						for _, result := range results {
							spell.DealDamage(sim, result)
							character.GainHealth(sim, result.Damage, healthMetrics)
						}
					},
				})

				// Confirmed by Zirene to pick the highest of your Agility or Strength
				if character.GetStat(stats.Agility) >= character.GetStat(stats.Strength) {
					corruptedAgilityAura.Activate(sim)
					corruptedAgilityAura.AddStacks(sim, int32(targetCount))
				} else {
					corruptedStrengthAura.Activate(sim)
					corruptedStrengthAura.AddStacks(sim, int32(targetCount))
				}
			},
		})

		consumptionTriggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Consumption Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			ICD:               time.Millisecond * 100,
			DPM:               character.AutoAttacks.NewDynamicProcManagerForWeaponEffect(CorruptedAshbringerLego, 5.0, 0),
			DPMProcCheck:      core.DPMProc,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				consumptionSpell.Cast(sim, result.Target)
			},
		})

		wrathTriggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Wrath of the Ashbringer Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			Duration:          time.Second * 30,
			DPM:               character.AutoAttacks.NewDynamicProcManagerForWeaponEffect(CorruptedAshbringerLego, 5.0, 0),
			DPMProcCheck:      core.DPMProc,
			ExtraCondition: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) bool {
				return !wrathOfTheAshbringerAura.IsActive()
			},
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				wrathOfTheAshbringerAura.Activate(sim)
				wrathSpell.Cast(sim, result.Target)
			},
		})

		ashbrinigerAgilityAura.ApplyOnGain(func(aura *core.Aura, sim *core.Simulation) {
			wrathTriggerAura.Activate(sim)
			consumptionTriggerAura.Deactivate(sim)
		}).ApplyOnExpire(func(aura *core.Aura, sim *core.Simulation) {
			consumptionTriggerAura.Activate(sim)
			wrathTriggerAura.Deactivate(sim)
		})

		ashbrinigerStrengthAura.ApplyOnGain(func(aura *core.Aura, sim *core.Simulation) {
			wrathTriggerAura.Activate(sim)
			consumptionTriggerAura.Deactivate(sim)
		}).ApplyOnExpire(func(aura *core.Aura, sim *core.Simulation) {
			consumptionTriggerAura.Activate(sim)
			wrathTriggerAura.Deactivate(sim)
		})

		character.ItemSwap.RegisterProc(CorruptedAshbringerLego, consumptionTriggerAura)
		character.ItemSwap.RegisterProc(CorruptedAshbringerLego, wrathTriggerAura)
	})

	// https://www.wowhead.com/classic/item=17068/deathbringer
	// Chance on hit: Sends a shadowy bolt at the enemy causing 110 to 140 Shadow damage.
	itemhelpers.CreateWeaponCoHProcDamage(Deathbringer, "Deathbringer", 1.0, 18138, core.SpellSchoolShadow, 110, 30, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=647/destiny
	itemhelpers.CreateWeaponProcAura(Destiny, "Destiny", 1.0, func(character *core.Character) *core.Aura {
		return character.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 17152},
			Label:    "Destiny",
			Duration: time.Second * 10,
		}).AttachStatBuff(stats.Strength, 200)
	})

	// https://www.wowhead.com/classic/item=230271/drake-talon-cleaver
	// Chance on hit: Delivers a fatal wound for 300 damage.
	itemhelpers.CreateWeaponCoHProcDamage(DrakeTalonCleaver, "Drake Talon Cleaver", 1.0, 467167, core.SpellSchoolPhysical, 300, 0, 0.0, core.DefenseTypeMelee)
	// https://www.wowhead.com/classic/item=232562/drake-talon-cleaver
	itemhelpers.CreateWeaponCoHProcDamage(DrakeTalonCleaverShadowflame, "Drake Talon Cleaver", 1.0, 467167, core.SpellSchoolPhysical, 300, 0, 0.0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=228410/dreadblade-of-the-destructor
	// https://www.wowhead.com/classic/item=228498/dreadblade-of-the-destructor
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(DreadbladeOfTheDestructor, "Dreadblade of the Destructor", 1.0, dreadbladeOfTheDestructorEffect)

	// https://www.wowhead.com/classic/item=227842/ebon-fist
	// Chance on hit: Sends a shadowy bolt at the enemy causing 125 to 275 Shadow damage.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponCoHProcDamage(EbonFist, "Ebon Fist", 1.0, 18211, core.SpellSchoolShadow, 125, 150, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=19170/ebon-hand
	// Chance on hit: Sends a shadowy bolt at the enemy causing 125 to 275 Shadow damage.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponCoHProcDamage(EbonHand, "Ebon Hand", 1.0, 18211, core.SpellSchoolShadow, 125, 150, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=227993/ebon-hilt-of-marduk
	// Chance on hit: Corrupts the target, causing 210 damage over 3 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(EbonHiltOfMarduk, "Ebon Hilt of Marduk", 1.0, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18656},
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Corruption (Ebon Hilt of Marduk)",
				},
				TickLength:    time.Second,
				NumberOfTicks: 3,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 70, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=228397/empyrean-demolisher
	// Chance on hit: Increases your attack speed by 20% for 10 sec.
	// Original proc rate 1.0 lowered to 0.6 in SoD phase 5
	itemhelpers.CreateWeaponProcAura(EmpyreanDemolisher, "Empyrean Demolisher", 0.6, func(character *core.Character) *core.Aura {
		return character.RegisterAura(core.Aura{
			Label:    "Empyrean Demolisher Haste Aura",
			ActionID: core.ActionID{SpellID: 21165},
			Duration: time.Second * 10,
		}).AttachMultiplyAttackSpeed(&character.Unit, 1.2)
	})

	// https://www.wowhead.com/classic/item=228349/eskhandars-left-claw
	// Chance on hit: Slows enemy's movement by 60% and causes them to bleed for 150 damage over 30 sec.
	// TODO: Proc rate untested
	itemhelpers.CreateWeaponProcSpell(EskhandarsLeftClaw, "Eskhandar's Left Claw", 1.0, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 22639},
			SpellSchool: core.SpellSchoolPhysical,
			DefenseType: core.DefenseTypeMelee,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot,
			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Eskhandar's Rake",
				},
				TickLength:    time.Second * 3,
				NumberOfTicks: 10,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 15, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=228350/eskhandars-right-claw
	// Chance on hit: Increases your attack speed by 25% for 5 sec.
	// Original proc rate 1.0 lowered to 0.6 in SoD phase 5
	itemhelpers.CreateWeaponProcAura(EskhandarsRightClaw, "Eskhandar's Right Claw", 0.6, eskhandarsRightClawAura)
	itemhelpers.CreateWeaponProcAura(EskhandarsRightClawMolten, "Eskhandar's Right Claw (Molten)", 0.6, eskhandarsRightClawAura)

	// https://www.wowhead.com/classic/item=13218/fang-of-the-crystal-spider
	// Chance on hit: Slows target enemy's casting speed and increases the time between melee and ranged attacks by 10% for 10 sec.
	// TODO: Proc rate assumed and needs testing
	core.NewItemEffect(FangOfTheCrystalSpider, func(agent core.Agent) {
		character := agent.GetCharacter()

		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit, _ int32) *core.Aura {
			aura := unit.RegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 17331},
				Label:    "Fang of the Crystal Spider",
				Duration: time.Second * 10,
			})
			core.AtkSpeedReductionEffect(aura, 1.10)
			return aura
		})

		dpm := character.AutoAttacks.NewDynamicProcManagerForEnchant(FangOfTheCrystalSpider, 1, 0)

		aura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Fang of the Crystal Spider Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			DPM:               dpm,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				debuffAuras.Get(result.Target).Activate(sim)
			},
		})

		character.ItemSwap.RegisterProc(FangOfTheCrystalSpider, aura)
	})

	// https://www.wowhead.com/classic/item=12590/felstriker
	// Chance on hit: All attacks are guaranteed to land and will be critical strikes for the next 3 sec.
	// Original proc rate 1.0 lowered to 0.6 in SoD phase 5
	core.NewItemEffect(Felstriker, func(agent core.Agent) {
		character := agent.GetCharacter()

		effectAura := character.NewTemporaryStatsAura("Felstriker", core.ActionID{SpellID: 16551}, stats.Stats{stats.MeleeCrit: 100 * core.CritRatingPerCritChance, stats.MeleeHit: 100 * core.MeleeHitRatingPerHitChance}, time.Second*3)

		dpm := character.AutoAttacks.NewDynamicProcManagerForEnchant(Felstriker, 0.6, 0)

		aura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Felstriker Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			DPM:               dpm,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				effectAura.Activate(sim)
			},
		})

		character.ItemSwap.RegisterProc(Felstriker, aura)
	})

	core.NewMobTypeAttackPowerEffect(FiendishMachete, []proto.MobType{proto.MobType_MobTypeElemental}, 36)

	itemhelpers.CreateWeaponProcSpell(FieryWarAxe, "Fiery War Axe", 1.0, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18796},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Fiery War Axe Fireball",
				},
				TickLength:    time.Second * 2,
				NumberOfTicks: 3,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				dmg := sim.Roll(155, 197)
				result := spell.CalcAndDealDamage(sim, target, dmg, spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})
	})

	itemhelpers.CreateWeaponProcSpell(Firebreather, "Firebreather", 1.0, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 16413},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 70, spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 3,
				TickLength:    time.Second * 2,
				Aura: core.Aura{
					Label: "Fireball",
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, 3, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=227934/flame-wrath
	// Chance on hit: Envelops the caster with a Fire shield for 15 sec and shoots a ring of fire dealing 130 to 170 damage to all nearby enemies.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcSpell(FlameWrath, "Flame Wrath", 1.0, func(character *core.Character) *core.Spell {
		shieldActionID := core.ActionID{SpellID: 461152}
		shieldSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:         shieldActionID,
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			BonusCoefficient: 1, // Only the shield portion has scaling
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 10, spell.OutcomeAlwaysHit)
			},
		})
		shieldAura := character.RegisterAura(core.Aura{
			ActionID: shieldActionID,
			Label:    "Flame Wrath",
			Duration: time.Second * 15,
			OnSpellHitTaken: func(_ *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() {
					shieldSpell.Cast(sim, spell.Unit)
				}
			},
		}).AttachStatBuff(stats.FireResistance, 30)

		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 461151},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				shieldAura.Activate(sim)

				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, sim.Roll(130, 170), spell.OutcomeMagicHit)
				}
			},
		})
	})

	itemhelpers.CreateWeaponProcSpell(FlurryAxe, "Flurry Axe", 1.9, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 18797},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.AutoAttacks.ExtraMHAttackProc(sim, 1, core.ActionID{SpellID: 18797}, spell)
			},
		})
	})

	// https://www.wowhead.com/classic/item=14024/frightalon
	// Chance on hit: Lowers all attributes of target by 10 for 1 min.
	// TODO: Proc rate assumed and needs testing
	core.NewItemEffect(Frightalon, func(agent core.Agent) {
		character := agent.GetCharacter()

		debuffAuraArray := character.NewEnemyAuraArray(func(target *core.Unit, _ int32) *core.Aura {
			return target.RegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 19755},
				Label:    "Frightalon",
				Duration: time.Minute * 1,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddStatsDynamic(sim, stats.Stats{
						stats.Agility:   -10,
						stats.Intellect: -10,
						stats.Stamina:   -10,
						stats.Spirit:    -10,
						stats.Strength:  -10,
					})
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddStatsDynamic(sim, stats.Stats{
						stats.Agility:   10,
						stats.Intellect: 10,
						stats.Stamina:   10,
						stats.Spirit:    10,
						stats.Strength:  10,
					})
				},
			})
		})

		dpm := character.AutoAttacks.NewDynamicProcManagerForEnchant(Frightalon, 1, 0)
		aura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Frightalon Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			DPM:               dpm,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				debuffAuraArray.Get(result.Target).Activate(sim)
			},
		})

		character.ItemSwap.RegisterProc(Frightalon, aura)
	})

	// https://www.wowhead.com/classic/item=227994/frightskull-shaft
	// Chance on hit: Deals 8 Shadow damage every 2 sec for 30 sec and lowers their Strength for the duration of the disease.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(FrightskullShaft, "Frightskull Shaft", 0.5, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18633},
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot | core.SpellFlagDisease,

			Dot: core.DotConfig{
				NumberOfTicks: 15,
				TickLength:    time.Second * 2,
				Aura: core.Aura{
					Label: "Weakening Disease",
					OnGain: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatDynamic(sim, stats.Strength, -50)
					},
					OnExpire: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatDynamic(sim, stats.Strength, 50)
					},
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=12797/frostguard#comments
	// Chance on hit: Target's movement slowed by 30% and increasing the time between attacks by 25% for 5 sec.
	// TODO: Proc rate assumed and needs testing
	core.NewItemEffect(Frostguard, func(agent core.Agent) {
		character := agent.GetCharacter()
		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit, _ int32) *core.Aura {
			aura := unit.RegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16927},
				Label:    "Chilled (Frostguard)",
				Duration: time.Second * 5,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddMoveSpeedModifier(&aura.ActionID, 0.30)
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.RemoveMoveSpeedModifier(&aura.ActionID)
				},
			})
			core.AtkSpeedReductionEffect(aura, 1.25)
			return aura
		})

		triggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Frostguard",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMeleeMH,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			PPM:               1, // Estimated based on data from WoW Armaments Discord
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				debuffAuras.Get(result.Target).Activate(sim)
			},
		})

		character.ItemSwap.RegisterProc(Frostguard, triggerAura)
	})

	// https://www.wowhead.com/classic/item=228029/gravestone-war-axe
	// Chance on hit: Diseases target enemy for 55 Nature damage every 3 sec for 15 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(GravestoneWarAxe, "Gravestone War Axe", 0.5, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18289},
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagDisease | core.SpellFlagPureDot,

			Dot: core.DotConfig{
				NumberOfTicks: 15,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Creeping Mold",
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 55, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=231273/grileks-carver
	// +141 Attack Power when fighting Dragonkin.
	core.NewMobTypeAttackPowerEffect(GrileksCarver, []proto.MobType{proto.MobType_MobTypeDragonkin}, 141)
	core.NewMobTypeAttackPowerEffect(GrileksCarverBloodied, []proto.MobType{proto.MobType_MobTypeDragonkin}, 141)

	// https://www.wowhead.com/classic/item=231274/grileks-grinder
	// +60 Attack Power when fighting Dragonkin.
	core.NewMobTypeAttackPowerEffect(GrileksGrinder, []proto.MobType{proto.MobType_MobTypeDragonkin}, 60)
	core.NewMobTypeAttackPowerEffect(GrileksGrinderBloodied, []proto.MobType{proto.MobType_MobTypeDragonkin}, 60)

	itemhelpers.CreateWeaponCoHProcDamage(GryphonRidersStormhammer, "Gryphon Rider's Stormhammer", 1.0, 18081, core.SpellSchoolNature, 91, 34, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=228267/gutgore-ripper
	// Chance on hit: Sends a shadowy bolt at the enemy causing 150 Shadow damage and lowering all stats by 25 for 30 sec.
	itemhelpers.CreateWeaponProcSpell(GutgoreRipper, "Gutgore Ripper", 1.0, gutgoreRipperEffect)
	itemhelpers.CreateWeaponProcSpell(GutgoreRipperMolten, "Gutgore Ripper (Molten)", 1.0, gutgoreRipperEffect)

	itemhelpers.CreateWeaponProcSpell(Gutwrencher, "Gutwrencher", 1.0, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 16406},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Rend (Gutwrencher)",
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
		})
	})

	itemhelpers.CreateWeaponCoHProcDamage(GutRipper, "Gut Ripper", 1.0, 18107, core.SpellSchoolPhysical, 95, 26, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=230991/halberd-of-smiting
	// Equip: Chance to decapitate the target on a melee swing, causing 452 to 676 damage.
	itemhelpers.CreateWeaponEquipProcDamage(HalberdOfSmiting, "Halberd of Smiting", 2.1, 467819, core.SpellSchoolPhysical, 452, 224, 0.0, core.DefenseTypeMelee)         // Works as phantom strike
	itemhelpers.CreateWeaponEquipProcDamage(HalberdOfSmitingBloodied, "Halberd of Smiting", 2.1, 467819, core.SpellSchoolPhysical, 452, 224, 0.0, core.DefenseTypeMelee) // Works as phantom strike

	itemhelpers.CreateWeaponCoHProcDamage(HammerOfTheNorthernWind, "Hammer of the Northern Wind", 3.5, 13439, core.SpellSchoolFrost, 20, 10, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=2243/hand-of-edward-the-odd
	// Chance on hit: Next spell cast within 4 sec will cast instantly.
	itemhelpers.CreateWeaponProcAura(HandOfEdwardTheOdd, "Hand of Edward the Odd", 1.0, func(character *core.Character) *core.Aura {
		return character.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 18803},
			Label:    "Focus (Hand of Edward the Odd)",
			Duration: time.Second * 4,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyCastSpeed(100000)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyCastSpeed(1 / 100000.0)
			},
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
				aura.Deactivate(sim)
			},
		})
	})

	itemhelpers.CreateWeaponCoHProcDamage(HanzoSword, "Hanzo Sword", 1.0, 16405, core.SpellSchoolPhysical, 75, 0, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=228022/headmasters-charge#comments
	// Use: Gives 20 additional intellect to party members within 30 yards. (10 Min Cooldown)
	// Originally did not stack with Arcane Intellect, but is reported to stack in SoD
	core.NewItemEffect(HeadmastersCharge, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 18264}

		buffAura := character.NewTemporaryStatsAura("Headmaster's Charge", actionID, stats.Stats{stats.Intellect: 25}, time.Minute*15)

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 10,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		character.AddMajorEquipmentCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=227887/hardened-frostguard
	// Chance on hit: Target's movement slowed by 30% and increasing the time between attacks by 25% for 5 sec.
	// Chance on hit: Inflicts Frost damage to nearby enemies, immobilizing them for up to 8 sec.
	// TODO: Proc rate assumed and needs testing
	core.NewItemEffect(HardenedFrostguard, func(agent core.Agent) {
		character := agent.GetCharacter()
		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit, _ int32) *core.Aura {
			aura := unit.RegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16927},
				Label:    "Chilled (Hardened Frostguard)",
				Duration: time.Second * 5,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddMoveSpeedModifier(&aura.ActionID, 0.30)
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.RemoveMoveSpeedModifier(&aura.ActionID)
				},
			})
			core.AtkSpeedReductionEffect(aura, 1.25)
			return aura
		})

		novaSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 463448},
			SpellSchool:      core.SpellSchoolFrost,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			BonusCoefficient: 1,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, 28, spell.OutcomeMagicHitAndCrit)
				}
			},
		})

		debuffTriggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Hardened Frostguard",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMeleeMH,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			PPM:               1.5, // Estimated based on data from WoW Armaments Discord
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				debuffAuras.Get(result.Target).Activate(sim)
			},
		})

		spellTriggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Hardened Frostguard",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMeleeMH,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			PPM:               1.5, // Estimated based on data from WoW Armaments Discord
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				novaSpell.Cast(sim, result.Target)
			},
		})

		character.ItemSwap.RegisterProc(Frostguard, debuffTriggerAura)
		character.ItemSwap.RegisterProc(Frostguard, spellTriggerAura)
	})

	itemhelpers.CreateWeaponProcSpell(HookfangShanker, "Hookfang Shanker", 1.0, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 13526},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPoison | core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Corrosive Poison",
					OnGain: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.Armor: -50})
					},
					OnExpire: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.Armor: 50})
					},
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, 7, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=227991/ironfoe
	// Chance on hit: Grants 2 extra attacks on your next swing.
	// TODO: Need updated proc rate lowered in SoD phase 5
	// Original proc rate 0.8 lowered to approximately 0.53 in SoD phase 5
	itemhelpers.CreateWeaponProcSpell(Ironfoe, "Ironfoe", 0.53, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 15494},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.AutoAttacks.ExtraMHAttackProc(sim, 2, core.ActionID{SpellID: 15494}, spell)
			},
		})
	})

	// https://www.wowhead.com/classic/item=230911/jekliks-crusher
	// Chance on hit: Wounds the target for 200 to 220 damage.
	// Original proc rate 4.0 lowered to 1.5 in SoD phase 5
	itemhelpers.CreateWeaponCoHProcDamage(JekliksCrusher, "Jeklik's Crusher", 1.5, 467642, core.SpellSchoolPhysical, 200, 20, 0.0, core.DefenseTypeMelee)
	itemhelpers.CreateWeaponCoHProcDamage(JekliksCrusherBloodied, "Jeklik's Crusher", 1.5, 467642, core.SpellSchoolPhysical, 200, 20, 0.0, core.DefenseTypeMelee)

	itemhelpers.CreateWeaponCoHProcDamage(JoonhosMercy, "Joonho's Mercy", 1.0, 20883, core.SpellSchoolArcane, 70, 0, 0, core.DefenseTypeMagic)

	itemhelpers.CreateWeaponProcSpell(KalimdorsRevenge, "Kalimdor's Revenge", 12, KalimdorsRevengeDamageProc)
	itemhelpers.CreateWeaponProcSpell(KalimdorsRevengeVoidTouched, "Kalimdor's Revenge", 12, KalimdorsRevengeDamageProc)

	itemhelpers.CreateWeaponCoHProcDamage(LinkensSwordOfMastery, "Linken's Sword of Mastery", 1.0, 18089, core.SpellSchoolNature, 45, 30, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=227940/lord-generals-sword
	// Chance on hit: Increases attack power by 50 for 30 sec.
	// // TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcAura(LordGeneralsSword, "Lord General's Sword", 1.0, func(character *core.Character) *core.Aura {
		return character.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 15602},
			Label:    "Lord General's Sword",
			Duration: time.Second * 30,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.AttackPower:       50,
					stats.RangedAttackPower: 50,
				})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.AttackPower:       -50,
					stats.RangedAttackPower: -50,
				})
			},
		})
	})

	core.NewItemEffect(ManslayerOfTheQiraji, func(agent core.Agent) {
		character := agent.GetCharacter()

		if !character.AutoAttacks.AutoSwingMelee {
			return
		}

		aura := manslayerOfTheQirajiAura(character)
		character.ItemSwap.RegisterProc(ManslayerOfTheQiraji, aura)
	})

	core.NewItemEffect(ManslayerOfTheQirajiVoidTouched, func(agent core.Agent) {
		character := agent.GetCharacter()

		if !character.AutoAttacks.AutoSwingMelee {
			return
		}

		aura := manslayerOfTheQirajiAura(character)
		character.ItemSwap.RegisterProc(ManslayerOfTheQirajiVoidTouched, aura)
	})

	// https://www.wowhead.com/classic/item=12794/masterwork-stormhammer
	// Chance on hit: Blasts up to 3 targets for 105 to 145 Nature damage.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcSpell(MasterworkStormhammer, "Masterwork Stormhammer", 0.5, func(character *core.Character) *core.Spell {
		maxHits := int(min(3, character.Env.GetNumTargets()))
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 463946},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for numHits := 0; numHits < maxHits; numHits++ {
					spell.CalcAndDealDamage(sim, target, sim.Roll(105, 145), spell.OutcomeMagicHitAndCrit)
					target = character.Env.NextTargetUnit(target)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=236221/misplaced-servo-arm
	// Equip: Chance to discharge electricity causing 250 to 300 Nature damage to your target.
	// If dual-wielding, your other weapon can proc the Misplaced Servo Arm when it strikes as well.
	// Chance-on-hit for the other weapon is determined by it's base weapon speed, set to 2PPM.
	// Same interaction when dual-wielding two Misplaced Servo Arms, one melee from one Arm has a chance to proc both Arms.
	// Assuming same PPM as in Classic for now
	core.NewItemEffect(MisplacedServoArm, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 1220535}
		label := "Electric Discharge Trigger"

		// Register 2 PPM for MH and 4 PPM for MH+OH
		// This is needed due to Item Swap
		dpm := character.AutoAttacks.NewDynamicProcManagerForWeaponEffect(MisplacedServoArm, 1.0, 0)
		dpmDualWield := character.AutoAttacks.NewDynamicProcManagerForWeaponEffect(MisplacedServoArm, 2.0, 0)

		procSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellProc | core.ProcMaskSpellDamageProc,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, sim.Roll(250, 300), spell.OutcomeMagicHitAndCrit)
			},
		})

		aura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              label,
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				var actualProcManager *core.DynamicProcManager
				if spell.ProcMask.Matches(core.ProcMaskMelee) {
					actualProcManager = dpmDualWield
				} else {
					actualProcManager = dpm
				}
				if actualProcManager.Proc(sim, spell.ProcMask, label) {
					procSpell.Cast(sim, result.Target)
				}
			},
		})

		character.ItemSwap.RegisterProc(MisplacedServoArm, aura)
	})

	// https://www.wowhead.com/classic/item=234987/neretzek-the-blood-drinker
	// Chance on hit: Steals 171 to 193 life from target enemy.
	itemhelpers.CreateWeaponProcSpell(NeretzekBloodDrinker, "Neretzek, The Blood Drinker", 0.8, neretzekBloodDrinkerEffect)
	itemhelpers.CreateWeaponProcSpell(NeretzekBloodDrinkerVoidTouched, "Neretzek, The Blood Drinker", 0.8, neretzekBloodDrinkerEffect)

	itemhelpers.CreateWeaponCoHProcDamage(Nightblade, "Nightblade", 1.0, 18211, core.SpellSchoolShadow, 125, 150, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=19169/nightfall
	// Removed from SoD
	// core.NewItemEffect(Nightfall, func(agent core.Agent) {
	// 	makeNightfallProc(agent.GetCharacter(), "Nightfall")
	// })

	itemhelpers.CreateWeaponCoHProcDamage(PendulumOfDoom, "Pendulum of Doom", 0.5, 10373, core.SpellSchoolPhysical, 250, 100, 0, core.DefenseTypeMelee)

	core.NewItemEffect(PipsSkinner, func(agent core.Agent) {
		character := agent.GetCharacter()

		targetWasBeast := false
		procAura := core.MakePermanent(character.RegisterAura(core.Aura{
			Label:    "Beast Slaying 45",
			ActionID: core.ActionID{SpellID: 18067},
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				targetIsBeast := character.CurrentTarget.MobType == proto.MobType_MobTypeBeast
				if targetIsBeast {
					targetWasBeast = targetIsBeast
					character.PseudoStats.MobTypeAttackPower += 45
				}
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				if targetWasBeast {
					character.PseudoStats.MobTypeAttackPower -= 45
				}
			},
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				targetWasBeast = false
			},
		}))

		character.ItemSwap.RegisterProc(Frostguard, procAura)
	})

	// https://www.wowhead.com/classic/item=228296/perditions-blade
	// Chance on hit: Blasts a target for 98 to 122 Fire damage.
	itemhelpers.CreateWeaponCoHProcDamage(PerditionsBlade, "Perdition's Blade", 2.8, 461695, core.SpellSchoolFire, 98, 24, 0, core.DefenseTypeMagic)
	itemhelpers.CreateWeaponCoHProcDamage(PerditionsBladeMolten, "Perdition's Blade", 2.8, 461695, core.SpellSchoolFire, 98, 24, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=231277/pitchfork-of-madness
	// +141 Attack Power when fighting Demons.
	core.NewMobTypeAttackPowerEffect(PitchforkOfMadness, []proto.MobType{proto.MobType_MobTypeDemon}, 141)
	core.NewMobTypeAttackPowerEffect(PitchforkOfMadnessBloodied, []proto.MobType{proto.MobType_MobTypeDemon}, 141)

	// https://www.wowhead.com/classic/item=228679/quelserrar
	// Chance on hit: When active, grants the wielder 25 defense and 300 armor for 10 sec.
	// Proc rate estimated based on data from WoW Armaments Discord for the original item
	itemhelpers.CreateWeaponProcAura(QuelSerrar, "Quel'Serrar", 2.0, func(character *core.Character) *core.Aura {
		return character.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 463105},
			Label:    "Sanctuary",
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.Defense:    25,
					stats.BonusArmor: 300,
				})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.Defense:    -25,
					stats.BonusArmor: -300,
				})
			},
		})
	})

	itemhelpers.CreateWeaponProcAura(Ravager, "Ravager", 1.0, func(character *core.Character) *core.Aura {
		tickActionID := core.ActionID{SpellID: 9633}
		procActionID := core.ActionID{SpellID: 9632}
		auraActionID := core.ActionID{SpellID: 433801}

		ravegerBladestormTickSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    tickActionID,
			SpellSchool: core.SpellSchoolPhysical,
			DefenseType: core.DefenseTypeMelee,
			ProcMask:    core.ProcMaskMeleeMHSpecial,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			BonusCoefficient: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				damage := 5.0 + spell.Unit.MHNormalizedWeaponDamage(sim, spell.MeleeAttackPower())
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, damage, spell.OutcomeMeleeSpecialHitAndCrit)
				}
			},
		})

		character.RegisterSpell(core.SpellConfig{
			SpellSchool: core.SpellSchoolPhysical,
			ActionID:    procActionID,
			ProcMask:    core.ProcMaskMeleeMHSpecial,
			Flags:       core.SpellFlagChanneled,
			Dot: core.DotConfig{
				IsAOE: true,
				Aura: core.Aura{
					Label: "Ravager Whirlwind",
				},
				NumberOfTicks:       3,
				TickLength:          time.Second * 3,
				AffectedByCastSpeed: false,
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					ravegerBladestormTickSpell.Cast(sim, target)
				},
			},
		})

		return character.RegisterAura(core.Aura{
			Label:    "Ravager Bladestorm",
			ActionID: auraActionID,
			Duration: time.Second * 9,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AutoAttacks.CancelAutoSwing(sim)
				dotSpell := character.GetSpell(procActionID)
				dotSpell.AOEDot().Apply(sim)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AutoAttacks.EnableAutoSwing(sim)
				dotSpell := character.GetSpell(procActionID)
				dotSpell.AOEDot().Cancel(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=228125/refined-arcanite-champion
	// Chance on hit: Heal self for 270 to 450 and Increases Strength by 120 for 30 sec.
	// Chance on hit: Increases damage done by 20 and attack speed by 5% for 15 sec.
	// TODO: Proc rate assumed and needs testing
	core.NewItemEffect(RefinedArcaniteChampion, func(agent core.Agent) {
		character := agent.GetCharacter()

		strengthAura := StrengthOfTheChampionAura(character)
		enrageAura := EnrageAura446327(character)

		ppm := 1.0 // Estimated based on data from WoW Armaments Discord
		strDpm := character.AutoAttacks.NewDynamicProcManagerForWeaponEffect(RefinedArcaniteChampion, ppm, 0)
		enrageDpm := character.AutoAttacks.NewDynamicProcManagerForWeaponEffect(RefinedArcaniteChampion, ppm, 0)

		strengthProcTrigger := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Refined Arcanite Champion (Strength)",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			DPM:               strDpm,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				strengthAura.Activate(sim)
			},
		})

		enrageProcTrigger := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Refined Arcanite Champion (Enrage)",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			DPM:               enrageDpm,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				enrageAura.Activate(sim)
			},
		})

		character.ItemSwap.RegisterProc(RefinedArcaniteChampion, strengthProcTrigger)
		character.ItemSwap.RegisterProc(RefinedArcaniteChampion, enrageProcTrigger)
	})

	// https://www.wowhead.com/classic/item=228543/runeblade-of-baron-rivendare
	// Equip: Increases movement speed and life regeneration rate.
	// TODO: Movement speed not implemented
	core.NewItemEffect(RunebladeOfBaronRivendare, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 17625}
		healthMetrics := character.NewHealthMetrics(actionID)
		aura := character.RegisterAura(core.Aura{
			ActionID: actionID,
			Label:    "Unholy Aura",
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				core.StartPeriodicAction(sim, core.PeriodicActionOptions{
					Period:   time.Second * 5,
					Priority: core.ActionPriorityAuto,
					OnAction: func(sim *core.Simulation) {
						character.GainHealth(sim, 20, healthMetrics)
					},
				})
			},
		})

		character.ItemSwap.RegisterProc(RunebladeOfBaronRivendare, aura)
	})

	itemhelpers.CreateWeaponCoHProcDamage(SatyrsLash, "Satyr's Lash", 1.0, 18205, core.SpellSchoolShadow, 55, 30, 0, core.DefenseTypeMagic)

	// TODO Searing Needle adds an "Apply Aura: Mod Damage Done (Fire): 10" aura to the /target/, buffing it; not currently modelled
	itemhelpers.CreateWeaponCoHProcDamage(SearingNeedle, "Searing Needle", 1.0, 16454, core.SpellSchoolFire, 60, 0, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=228666/seeping-willow
	// Chance on hit: Lowers all stats by 20 and deals 20 Nature damage every 3 sec to all enemies within an 8 yard radius of the caster for 30 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(SeepingWillow, "Seeping Willow", 0.5, func(character *core.Character) *core.Spell {
		stats := stats.Stats{
			stats.Agility:   20,
			stats.Intellect: 20,
			stats.Stamina:   20,
			stats.Spirit:    20,
			stats.Strength:  20,
		}
		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit, _ int32) *core.Aura {
			return unit.RegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 17196},
				Label:    "Seeping Willow",
				Duration: time.Second * 30,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					unit.AddStatsDynamic(sim, stats.Invert())
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					unit.AddStatsDynamic(sim, stats)
				},
			})
		})

		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 17196},
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPoison | core.SpellFlagPureDot,
			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Seeping Willow Poison",
				},
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 20, isRollover)
					debuffAuras.Get(target).Activate(sim)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					result := spell.CalcAndDealOutcome(sim, aoeTarget, spell.OutcomeMagicHit)
					if result.Landed() {
						spell.Dot(aoeTarget).Apply(sim)
					}
				}
			},
		})
	})

	itemhelpers.CreateWeaponProcSpell(SerpentSlicer, "Serpent Slicer", 1.0, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 17511},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPoison | core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 2,
				Aura: core.Aura{
					Label: "Poison (Serpent Slicer)",
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
		})
	})

	itemhelpers.CreateWeaponCoHProcDamage(Shadowblade, "Shadowblade", 1.0, 18138, core.SpellSchoolShadow, 110, 30, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=228272/shadowstrike
	// Chance on hit: Steals 180 to 220 life from target enemy.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcSpell(Shadowstrike, "Shadowstrike", 2.2, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 461683}
		healthMetrics := character.NewHealthMetrics(actionID)
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         actionID,
			SpellSchool:      core.SpellSchoolShadow,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BonusCoefficient: 1.0,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, sim.Roll(180, 220), spell.OutcomeMagicHit)
				character.GainHealth(sim, result.Damage, healthMetrics)
			},
		})
	})

	itemhelpers.CreateWeaponCoHProcDamage(ShortswordOfVengeance, "Shortsword of Vengeance", 1.0, 13519, core.SpellSchoolHoly, 30, 0, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=228542/skullforge-reaver
	// Equip: Drains target for 2 Shadow damage every 1 sec and transfers it to the caster. Lasts for 30 sec.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcSpell(SkullforgeReaver, "Skullforge Reaver", 1.7, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 17484}
		healthMetrics := character.NewHealthMetrics(actionID)
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot,
			Dot: core.DotConfig{
				NumberOfTicks: 30,
				TickLength:    time.Second,
				Aura: core.Aura{
					Label: "Skullforge Brand",
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 2, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					result := dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
					character.GainHealth(sim, result.Damage, healthMetrics)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=227683/sulfuras-hand-of-ragnaros
	// Chance on hit: Hurls a fiery ball that causes 273 to 333 Fire damage and purges the target's soul, increasing Fire and Holy damage taken by up to 30 and dealing an additional 75 damage over 10 sec.
	// Equip: 20% chance to deal 25 Fire damage to all nearby enemies when you are struck by a melee attack. (Proc chance: 20%)
	core.NewItemEffect(SulfurasHandOfRagnaros, func(agent core.Agent) {
		character := agent.GetCharacter()

		immolationSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 460335},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			BonusCoefficient: .025,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 25, spell.OutcomeAlwaysHit)
			},
		})

		immolationProcTrigger := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Immolation (Hand of Ragnaros)",
			Callback:   core.CallbackOnSpellHitTaken,
			Outcome:    core.OutcomeLanded,
			ProcMask:   core.ProcMaskMelee,
			ProcChance: .20,
			Handler: func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					immolationSpell.Cast(sim, aoeTarget)
				}
			},
		})

		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit, _ int32) *core.Aura {
			return unit.RegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 460338},
				Label:    "Purged by Fire",
				Duration: time.Second * 10,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFire] += 30
					unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] += 30
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFire] -= 30
					unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] -= 30
				},
			})
		})

		purgedByFireSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 460338},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Purged By Fire",
				},
				TickLength:    time.Second * 2,
				NumberOfTicks: 5,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 15, isRollover)
					debuffAuras.Get(target).Activate(sim)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, sim.Roll(273, 333), spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})

		purgedByFireProcTrigger := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Purged by Fire Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			PPM:               1, // Estimated based on data from WoW Armaments Discord
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				purgedByFireSpell.Cast(sim, result.Target)
			},
		})

		character.ItemSwap.RegisterProc(SulfurasHandOfRagnaros, immolationProcTrigger)
		character.ItemSwap.RegisterProc(SulfurasHandOfRagnaros, purgedByFireProcTrigger)
	})

	// https://www.wowhead.com/classic/item=17182/sulfuras-hand-of-ragnaros
	// Chance on hit: Hurls a fiery ball that causes 273 to 333 Fire damage and an additional 75 damage over 10 sec.
	// Equip: Deals 5 Fire damage to anyone who strikes you with a melee attack.
	// core.NewItemEffect(SulfurasHandOfRagnaros, func(agent core.Agent) {
	// 	character := agent.GetCharacter()

	// 	immolationActionID := core.ActionID{SpellID: 21142}

	// 	immolationSpell := character.RegisterSpell(core.SpellConfig{
	// 		ActionID:    immolationActionID,
	// 		SpellSchool: core.SpellSchoolFire,
	// 		ProcMask:    core.ProcMaskEmpty,

	// 		DamageMultiplier: 1,
	// 		ThreatMultiplier: 1,

	// 		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
	// 			spell.CalcAndDealDamage(sim, target, 5, spell.OutcomeMagicHit)
	// 		},
	// 	})

	// 	character.RegisterAura(core.Aura{
	// 		ActionID: immolationActionID,
	// 		Label:    "Immolation (Hand of Ragnaros)",
	// 		Duration: core.NeverExpires,
	// 		OnReset: func(aura *core.Aura, sim *core.Simulation) {
	// 			aura.Activate(sim)
	// 		},
	// 		OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
	// 			if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMelee) {
	// 				immolationSpell.Cast(sim, spell.Unit)
	// 			}
	// 		},
	// 	})

	// 	fireballSpell := character.RegisterSpell(core.SpellConfig{
	// 		ActionID:    core.ActionID{SpellID: 21162},
	// 		SpellSchool: core.SpellSchoolFire,
	// 		DefenseType: core.DefenseTypeMagic,
	// 		ProcMask:    core.ProcMaskEmpty,

	// 		DamageMultiplier: 1,
	// 		ThreatMultiplier: 1,

	// 		Dot: core.DotConfig{
	// 			Aura: core.Aura{
	// 				Label: "Fireball (Hand of Ragnaros)",
	// 			},
	// 			TickLength:    time.Second * 2,
	// 			NumberOfTicks: 5,

	// 			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
	// 				dot.Snapshot(target, 15, isRollover)
	// 			},

	// 			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
	// 				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
	// 			},
	// 		},

	// 		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
	// 			result := spell.CalcAndDealDamage(sim, target, sim.Roll(273, 333), spell.OutcomeMagicHitAndCrit)
	// 			if result.Landed() {
	// 				spell.Dot(target).Apply(sim)
	// 			}
	// 		},
	// 	})

	// 	core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
	// 		Name:     "Hand of Ragnaros Trigger",
	// 		Callback: core.CallbackOnSpellHitDealt,
	// 		Outcome:  core.OutcomeLanded,
	// 		ProcMask: core.ProcMaskMelee,
	// 		PPM:      1, // Estimated based on data from WoW Armaments Discord
	// 		Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
	// 			fireballSpell.Cast(sim, result.Target)
	// 		},
	// 	})
	// })

	// https://www.wowhead.com/classic/item=227684/sulfuron-hammer
	// Chance on hit: Hurls a fiery ball that causes 83 to 101 Fire damage and an additional 16 damage over 8 sec.
	// Equip: Deals 5 Fire damage to anyone who strikes you with a melee attack.
	core.NewItemEffect(SulfuronHammer, func(agent core.Agent) {
		character := agent.GetCharacter()

		immolationSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 21142},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 5, spell.OutcomeAlwaysHit)
			},
		})

		immolationProcTrigger := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Immolation (Hand of Ragnaros)",
			Callback: core.CallbackOnSpellHitTaken,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee,
			Handler: func(sim *core.Simulation, spell *core.Spell, _ *core.SpellResult) {
				immolationSpell.Cast(sim, spell.Unit)
			},
		})

		fireballSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 21159},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Fireball (Sulfuron Hammer)",
				},
				TickLength:    time.Second * 2,
				NumberOfTicks: 4,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 4, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, sim.Roll(83, 101), spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})

		sulfuranHammerProcTrigger := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Sulfuron Hammer Trigger",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee,
			PPM:      1, // TODO: Armaments Discord didn't have any data on Sulfuron Hammer
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				fireballSpell.Cast(sim, result.Target)
			},
		})

		character.ItemSwap.RegisterProc(SulfuronHammer, immolationProcTrigger)
		character.ItemSwap.RegisterProc(SulfuronHammer, sulfuranHammerProcTrigger)
	})

	// https://www.wowhead.com/classic/item=227832/tempered-black-amnesty
	// Chance on hit: Reduce your threat to the current target making them less likely to attack you.
	// TODO: Proc rate untested, no way to reduce threat right now
	// itemhelpers.CreateWeaponProcSpell(TemperedBlackAmnesty, "Tempered Black Amnesty", 1.0, func(character *core.Character) *core.Spell {
	// 	return character.RegisterSpell(core.SpellConfig{
	// 		ActionID:         core.ActionID{SpellID: 23604},
	// 		SpellSchool:      core.SpellSchoolPhysical,
	// 		ProcMask:         core.ProcMaskEmpty,
	// 		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
	// 			character.threat
	// 		},
	// 	})
	// })

	itemhelpers.CreateWeaponProcSpell(TheHandOfAntusul, "The Hand of Antu'sul", 1.0, func(character *core.Character) *core.Spell {
		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit, _ int32) *core.Aura {
			aura := unit.RegisterAura(core.Aura{
				Label:    "ThunderClap-Antu'sul",
				ActionID: core.ActionID{SpellID: 13532},
				Duration: time.Second * 10,
			})
			core.AtkSpeedReductionEffect(aura, 1.11)
			return aura
		})

		results := make([]*core.SpellResult, min(4, character.Env.GetNumTargets()))

		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 13532},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for idx := range results {
					results[idx] = spell.CalcDamage(sim, target, 7, spell.OutcomeMagicHitAndCrit)
					target = character.Env.NextTargetUnit(target)
				}
				for _, result := range results {
					spell.DealDamage(sim, result)
					if result.Landed() {
						debuffAuras.Get(result.Target).Activate(sim)
					}
				}
			},
		})
	})

	itemhelpers.CreateWeaponProcAura(TheJackhammer, "The Jackhammer", 1.0, func(character *core.Character) *core.Aura {
		return character.RegisterAura(core.Aura{
			Label:    "The Jackhammer Haste Aura",
			ActionID: core.ActionID{SpellID: 13533},
			Duration: time.Second * 10,
		}).AttachMultiplyAttackSpeed(&character.Unit, 1.3)
	})

	itemhelpers.CreateWeaponCoHProcDamage(TheNeedler, "The Needler", 3.0, 13060, core.SpellSchoolPhysical, 75, 0, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=230242/the-untamed-blade
	// Chance on hit: Increases Strength by 300 for 8 sec.
	// Estimated based on data from WoW Armaments Discord
	// Original proc rate 1.0 lowered to approximately 0.55 in SoD phase 5
	itemhelpers.CreateWeaponProcAura(TheUntamedBlade, "The Untamed Blade", 0.55, func(character *core.Character) *core.Aura {
		return character.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 23719},
			Label:    "Untamed Fury",
			Duration: time.Second * 8,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.Strength: 300})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.Strength: -300})
			},
		})
	})
	// https://www.wowhead.com/classic/item=232566/the-untamed-blade
	itemhelpers.CreateWeaponProcAura(TheUntamedBladeShadowflame, "The Untamed Blade", 0.55, func(character *core.Character) *core.Aura {
		return character.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 23719},
			Label:    "Untamed Fury",
			Duration: time.Second * 8,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.Strength: 300})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.Strength: -300})
			},
		})
	})

	itemhelpers.CreateWeaponProcSpell(ThrashBlade, "Thrash Blade", 1.0, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 21919},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.AutoAttacks.ExtraMHAttackProc(sim, 1, core.ActionID{SpellID: 21919}, spell)
			},
		})
	})

	// https://www.wowhead.com/classic/item=19019/thunderfury-blessed-blade-of-the-windseeker
	// Chance on hit: Blasts your enemy with lightning, dealing 300 Nature damage and then jumping to additional nearby enemies.
	// Each jump reduces that victim's Nature resistance by 25. Affects 5 targets.
	// Your primary target is also consumed by a cyclone, slowing its attack speed by 20% for 12 sec.
	core.NewItemEffect(Thunderfury, func(agent core.Agent) {
		character := agent.GetCharacter()

		thunderfuryASAuras := character.NewEnemyAuraArray(core.ThunderfuryASAura)
		procActionID := core.ActionID{SpellID: 21992}

		singleTargetSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    procActionID.WithTag(1),
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellProc | core.ProcMaskSpellDamageProc,
			Flags:       core.SpellFlagIgnoreAttackerModifiers,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			FlatThreatBonus:  126,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 300, spell.OutcomeMagicHitAndCrit)
				thunderfuryASAuras.Get(target).Activate(sim)
			},
		})

		debuffAuras := character.NewEnemyAuraArray(func(target *core.Unit, _ int32) *core.Aura {
			return character.NewTemporaryStatsAura("Thunderfury", procActionID, stats.Stats{stats.NatureResistance: 25}, time.Second*12)
		})

		results := make([]*core.SpellResult, min(5, character.Env.GetNumTargets()))

		bounceSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    procActionID.WithTag(2),
			SpellSchool: core.SpellSchoolNature,
			ProcMask:    core.ProcMaskEmpty,

			ThreatMultiplier: 1,
			FlatThreatBonus:  126,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for idx := range results {
					results[idx] = spell.CalcDamage(sim, target, 0, spell.OutcomeMagicHit)
					target = sim.Environment.NextTargetUnit(target)
				}
				for _, result := range results {
					if result.Landed() {
						debuffAuras[result.Target.Index].Activate(sim)
					}
					spell.DealDamage(sim, result)
				}
			},
		})

		triggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:         "Thunderfury Trigger",
			DPM:          character.AutoAttacks.NewDynamicProcManagerForWeaponEffect(Thunderfury, 6.0, 0),
			DPMProcCheck: core.DPMProc,
			Callback:     core.CallbackOnSpellHitDealt,
			Outcome:      core.OutcomeLanded,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				singleTargetSpell.Cast(sim, result.Target)
				bounceSpell.Cast(sim, result.Target)
			},
		})

		character.ItemSwap.RegisterProc(Thunderfury, triggerAura)
	})

	// https://www.wowhead.com/classic/item=228273/thunderstrike
	// Chance on hit: Blasts up to 3 targets for 200 to 300 Nature damage. Each target after the first takes less damage.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(Thunderstrike, "Thunderstrike", 1.5, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 461686},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				initialResult := spell.CalcOutcome(sim, target, spell.OutcomeMagicHit)
				// Only the initial hit can be fully resisted according to a wowhead comment
				if initialResult.Landed() {
					damageMultiplier := 1.0
					for numHits := 0; numHits < 3; numHits++ {
						spell.CalcAndDealDamage(sim, target, sim.Roll(150, 250)*damageMultiplier, spell.OutcomeMagicCrit)
						numHits++
						target = character.Env.NextTargetUnit(target)
						// TODO: Couldn't find information on what the multiplier actually is
						damageMultiplier *= .65
					}
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=231272/tigules-harpoon
	// +99 Attack Power when fighting Beasts.
	core.NewMobTypeAttackPowerEffect(TigulesHarpoon, []proto.MobType{proto.MobType_MobTypeBeast}, 99)
	core.NewMobTypeAttackPowerEffect(TigulesHarpoonBloodied, []proto.MobType{proto.MobType_MobTypeBeast}, 99)

	// https://www.wowhead.com/classic/item=228347/typhoon
	// Chance on hit: Grants an extra attack on your next swing.
	// TODO: Proc rate assumed and needs testing
	core.NewItemEffect(Typhoon, func(agent core.Agent) {
		character := agent.GetCharacter()
		dpm := character.AutoAttacks.NewDynamicProcManagerForWeaponEffect(Typhoon, 1.0, 0)
		triggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Typhoon Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			DPM:               dpm,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				character.AutoAttacks.ExtraMHAttackProc(sim, 1, core.ActionID{SpellID: 461985}, spell)
			},
		})

		character.ItemSwap.RegisterProc(Typhoon, triggerAura)
	})

	// https://www.wowhead.com/classic/item=13183/venomspitter
	// Chance on hit: Poisons target for 7 Nature damage every 2 sec for 30 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(Venomspitter, "Venomspitter", 1.0, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18203},
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPoison | core.SpellFlagPureDot,
			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Poison (Venomspitter)",
				},
				TickLength:    time.Second * 2,
				NumberOfTicks: 15,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 7, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})
	})

	itemhelpers.CreateWeaponCoHProcDamage(VilerendSlicer, "Vilerend Slicer", 1.0, 16405, core.SpellSchoolPhysical, 75, 0, 0, core.DefenseTypeMelee)

	itemhelpers.CreateWeaponCoHProcDamage(ViskagTheBloodletter, "Vis'kag the Bloodletter", 0.6, 21140, core.SpellSchoolPhysical, 240, 0, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=227941/wraith-scythe
	// Chance on hit: Steals 45 life from target enemy.
	itemhelpers.CreateWeaponProcSpell(WraithScythe, "Wraith Scythe", 1.0, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 16414}
		healthMetrics := character.NewHealthMetrics(actionID)

		return character.RegisterSpell(core.SpellConfig{
			ActionID:         actionID,
			SpellSchool:      core.SpellSchoolShadow,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			BonusCoefficient: 0.3,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 45, spell.OutcomeAlwaysHit)
				character.GainHealth(sim, result.Damage, healthMetrics)
			},
		})
	})

	// https://www.wowhead.com/classic/item=230930/zulian-slicer
	// Chance on hit: Slices the enemy for 72 to 96 Nature damage.
	itemhelpers.CreateWeaponCoHProcDamage(ZulianSlicer, "Zulian Slicer", 1.2, 467738, core.SpellSchoolNature, 72, 24, 0.35, core.DefenseTypeMelee)
	itemhelpers.CreateWeaponCoHProcDamage(ZulianSlicerBloodied, "Zulian Slicer", 1.2, 467738, core.SpellSchoolNature, 72, 24, 0.35, core.DefenseTypeMelee)

	///////////////////////////////////////////////////////////////////////////
	//                                 Trinkets
	///////////////////////////////////////////////////////////////////////////

	// https://www.wowhead.com/classic/item=227972/burst-of-knowledge
	// Use: Reduces mana cost of all spells by 100 for 10 sec. (5 Min Cooldown)
	core.NewItemEffect(BurstOfKnowledge, func(agent core.Agent) {
		character := agent.GetCharacter()

		aura := character.RegisterAura(core.Aura{
			ActionID: core.ActionID{ItemID: BurstOfKnowledge},
			Label:    "Burst of Knowledge",
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				for _, spell := range aura.Unit.Spellbook {
					if spell.Cost != nil && spell.Cost.CostType() == core.CostTypeMana {
						spell.Cost.Multiplier -= 100
					}
				}
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				for _, spell := range aura.Unit.Spellbook {
					if spell.Cost != nil && spell.Cost.CostType() == core.CostTypeMana {
						spell.Cost.Multiplier += 100
					}
				}
			},
		})

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: core.ActionID{ItemID: BurstOfKnowledge},
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 5,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				aura.Activate(sim)
			},
		})

		character.AddMajorEquipmentCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeMana,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=228678/draconic-infused-emblem
	// Use: Increases your spell damage by up to 100 and your healing by up to 190 for 15 sec. (1 Min, 30 Sec Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(DraconicInfusedEmblem, stats.Stats{stats.SpellDamage: 128, stats.HealingPower: 236}, time.Second*15, time.Second*90)

	// https://www.wowhead.com/classic/item=19288/darkmoon-card-blue-dragon
	// Equip: 5% chance on successful spellcast to allow 100% of your Mana regeneration to continue while casting for 15 sec.
	core.NewItemEffect(DarkmoonCardBlueDragon, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{SpellID: 1213421}

		procAura := character.RegisterAura(core.Aura{
			Label:    "Aura of the Blue Dragon",
			ActionID: actionID,
			Duration: time.Second * 15,
		}).AttachAdditivePseudoStatBuff(&character.PseudoStats.SpiritRegenRateCasting, 1)

		triggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Aura of the Blue Dragon Trigger",
			Callback:   core.CallbackOnCastComplete,
			ProcMask:   core.ProcMaskSpellDamage | core.ProcMaskSpellHealing,
			ProcChance: .05,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procAura.Activate(sim)
			},
		})

		character.ItemSwap.RegisterProc(DarkmoonCardBlueDragon, triggerAura)
	})

	// https://www.wowhead.com/classic/item=19287/darkmoon-card-heroism
	// Equip: Sometimes heals bearer of 170 to 230 damage when damaging an enemy in melee.
	core.NewItemEffect(DarkmoonCardHeroism, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{SpellID: 1213419}
		healthMetrics := character.NewHealthMetrics(actionID)

		procSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolHoly,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.GainHealth(sim, sim.Roll(170, 230), healthMetrics)
			},
		})

		triggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Heroism Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
			PPM:               2,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procSpell.Cast(sim, spell.Unit)
			},
		})

		character.ItemSwap.RegisterProc(DarkmoonCardHeroism, triggerAura)
	})

	// https://www.wowhead.com/classic/item=19289/darkmoon-card-maelstrom
	// Equip: Chance to strike your melee target with lightning for 317 to 517 Nature damage.
	core.NewItemEffect(DarkmoonCardMaelstrom, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{SpellID: 1213417}

		procSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, sim.Roll(317, 517), spell.OutcomeMagicHitAndCrit)
			},
		})

		triggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Lightning Strike Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
			PPM:               1.0,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procSpell.Cast(sim, result.Target)
			},
		})

		character.ItemSwap.RegisterProc(DarkmoonCardMaelstrom, triggerAura)
	})

	// https://www.wowhead.com/classic/item=234462/earthstrike
	// Use: Increases your melee and ranged attack power by 328.  Effect lasts for 20 sec. (2 Min Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(Earthstrike, stats.Stats{stats.AttackPower: 328, stats.RangedAttackPower: 328}, time.Second*20, time.Second*120)

	// https://www.wowhead.com/classic/item=228293/essence-of-the-pure-flame
	// Equip: When struck in combat inflicts 50 Fire damage to the attacker.
	core.NewItemEffect(EssenceOfThePureFlame, func(agent core.Agent) {
		character := agent.GetCharacter()
		character.PseudoStats.ThornsDamage += 50

		procSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 461694},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagBinary | core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 50, spell.OutcomeAlwaysHit)
			},
		})

		triggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Fiery Aura Proc",
			Callback: core.CallbackOnSpellHitTaken,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procSpell.Cast(sim, spell.Unit)
			},
		})

		character.ItemSwap.RegisterProc(EssenceOfThePureFlame, triggerAura)
	})

	// https://www.wowhead.com/classic/item=234080/eye-of-moam
	// Use: Increases damage done by magical spells and effects by up to 150, and decreases the magical resistances of your spell targets by 100 for 30 sec. (3 Min Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(EyeOfMoam, stats.Stats{stats.SpellDamage: 150, stats.SpellPenetration: 100}, time.Second*30, time.Minute*3)

	// https://www.wowhead.com/classic/item=234092/fetish-of-chitinous-spikes#see-also
	// Use: Spikes sprout from you causing 150 Nature damage to attackers when hit.  Lasts 30 sec. (3 Min Cooldown)
	// "Ouch"
	core.NewItemEffect(FetishOfChitinousSpikes, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{ItemID: FetishOfChitinousSpikes}

		thornSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 1214263},
			SpellSchool: core.SpellSchoolNature,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 150, spell.OutcomeMagicHit)
			},
		})

		thornsAura := character.RegisterAura(core.Aura{
			ActionID: actionID,
			Label:    "Chitinous Spikes",
			Duration: time.Second * 30,
			OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() {
					thornSpell.Cast(sim, spell.Unit)
				}
			},
		})

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolPhysical,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 3,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				thornsAura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=236337/glyph-of-deflection
	// Use: Increases the block value of your shield by 400 for 20 sec. (2 Min Cooldown)
	core.NewSimpleStatDefensiveTrinketEffect(GlyphOfDeflection, stats.Stats{stats.BlockValue: 400}, time.Second*20, time.Minute*2)

	core.NewItemEffect(HandOfJustice, func(agent core.Agent) {
		character := agent.GetCharacter()
		if !character.AutoAttacks.AutoSwingMelee {
			return
		}

		triggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Hand of Justice Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
			ProcChance:        0.02,
			ICD:               time.Second * 2,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				spell.Unit.AutoAttacks.ExtraMHAttackProc(sim, 1, core.ActionID{SpellID: 15600}, spell)
			},
		})

		character.ItemSwap.RegisterProc(HandOfJustice, triggerAura)
	})

	// https://www.wowhead.com/classic/item=228599/heart-of-wyrmthalak
	// Equip: Chance to bathe your melee target in flames for 120 to 180 Fire damage.
	// TODO: Proc rate assumed from a wowhead comment and needs testing
	core.NewItemEffect(HeartOfWyrmthalak, func(agent core.Agent) {
		character := agent.GetCharacter()
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 462385},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, sim.Roll(120, 180), spell.OutcomeMagicHitAndCrit)
			},
		})
		triggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Heart of Wyrmthalak Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
			PPM:               0.4,
			Handler: func(sim *core.Simulation, _ *core.Spell, result *core.SpellResult) {
				spell.Cast(sim, result.Target)
			},
		})

		character.ItemSwap.RegisterProc(HeartOfWyrmthalak, triggerAura)
	})

	// https://www.wowhead.com/classic/item=233627/jom-gabbar
	// Use: Increases attack power by 70 and an additional 70 every 2 sec.  Lasts 20 sec. (2 Min Cooldown)
	core.NewItemEffect(JomGabbar, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 1213366}
		duration := time.Second * 20
		bonusPerStack := stats.Stats{
			stats.AttackPower:       70,
			stats.RangedAttackPower: 70,
		}

		jomGabbarAura := character.RegisterAura(core.Aura{
			Label:     "Jom Gabbar",
			ActionID:  actionID,
			Duration:  duration,
			MaxStacks: 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				core.StartPeriodicAction(sim, core.PeriodicActionOptions{
					Period:          time.Second * 2,
					NumTicks:        10,
					Priority:        core.ActionPriorityAuto,
					TickImmediately: true,
					OnAction: func(sim *core.Simulation) {
						aura.AddStack(sim)
					},
				})
			},
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks, newStacks int32) {
				bonusStats := bonusPerStack.Multiply(float64(newStacks - oldStacks))
				character.AddStatsDynamic(sim, bonusStats)
			},
		})
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 2,
				},
				SharedCD: core.Cooldown{
					Timer:    character.GetOffensiveTrinketCD(),
					Duration: duration,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				jomGabbarAura.Activate(sim)
			},
		})

		character.AddMajorEquipmentCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=236268/kiss-of-the-spider
	// Use: Increases your attack speed by 20% for 15 sec. (1 Min, 30 Sec Cooldown)
	core.NewItemEffect(KissOfTheSpider, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{ItemID: KissOfTheSpider}
		duration := time.Second * 15
		buffAura := character.RegisterAura(core.Aura{
			ActionID: actionID,
			Label:    "Kiss of the Spider",
			Duration: duration,
		}).AttachMultiplyAttackSpeed(&character.Unit, 1.20)

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolPhysical,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Second * 90,
				},
				SharedCD: core.Cooldown{
					Timer:    character.GetOffensiveTrinketCD(),
					Duration: duration,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=236352/mark-of-the-champion
	// Equip: +157 Attack Power when fighting Undead and Demons.
	core.NewMobTypeAttackPowerEffect(MarkOfTheChampionPhys, []proto.MobType{proto.MobType_MobTypeUndead, proto.MobType_MobTypeDemon}, 157)

	// https://www.wowhead.com/classic/item=236351/mark-of-the-champion
	// Equip: Increases damage done to Undead and Demons by magical spells and effects by up to 89.
	core.NewMobTypeSpellPowerEffect(MarkOfTheChampionSpell, []proto.MobType{proto.MobType_MobTypeUndead, proto.MobType_MobTypeDemon}, 89)

	// https://www.wowhead.com/classic/item=23046/the-restrained-essence-of-sapphiron
	// Use: Increases damage and healing done by magical spells and effects by up to 180 for 20 sec. (2 Min Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(TheRestrainedEssenceOfSapphiron, stats.Stats{stats.SpellPower: 180}, time.Second*20, time.Minute*2)

	// https://www.wowhead.com/classic/item=17774/mark-of-the-chosen
	// Equip: Has a 2% chance when struck in combat of increasing all stats by 25 for 1 min. (Proc chance: 2%)
	core.NewItemEffect(MarkOfTheChosen, func(agent core.Agent) {
		character := agent.GetCharacter()
		statIncrease := float64(25)

		procAura := character.RegisterAura(core.Aura{
			Label:    "Mark of the Chosen Effect",
			ActionID: core.ActionID{SpellID: 21970},
			Duration: time.Minute,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.Stamina:   statIncrease,
					stats.Agility:   statIncrease,
					stats.Strength:  statIncrease,
					stats.Intellect: statIncrease,
					stats.Spirit:    statIncrease,
				})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.Stamina:   -statIncrease,
					stats.Agility:   -statIncrease,
					stats.Strength:  -statIncrease,
					stats.Intellect: -statIncrease,
					stats.Spirit:    -statIncrease,
				})
			},
		})

		triggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Mark of the Chosen",
			Callback:   core.CallbackOnSpellHitTaken,
			ProcMask:   core.ProcMaskMelee,
			Outcome:    core.OutcomeLanded,
			ProcChance: 0.02,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procAura.Activate(sim)
			},
		})

		character.ItemSwap.RegisterProc(MarkOfTheChosen, triggerAura)
	})

	// https://www.wowhead.com/classic/item=231271/nat-pagles-broken-reel
	core.NewSimpleStatOffensiveTrinketEffect(NatPaglesBrokenReel, stats.Stats{
		stats.SpellHit: 10 * core.SpellHitRatingPerHitChance,
		stats.MeleeHit: 10 * core.MeleeHitRatingPerHitChance,
	}, time.Second*15, time.Second*90)

	// https://www.wowhead.com/classic/item=19812/rune-of-the-dawn
	// Equip: Increases damage done to Undead by magical spells and effects by up to 48.
	core.NewMobTypeSpellPowerEffect(RuneOfTheDawn, []proto.MobType{proto.MobType_MobTypeUndead}, 48)

	// https://www.wowhead.com/classic/item=233601/scarab-brooch
	// Use: Your magical heals provide the target with a shield that absorbs damage equal to 15% of the amount healed for 30 sec. (3 Min Cooldown)
	core.NewItemEffect(ScarabBrooch, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{ItemID: ScarabBrooch}

		shieldSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 26470},
			SpellSchool: core.SpellSchoolNature,
			ProcMask:    core.ProcMaskSpellHealing,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell | core.SpellFlagHelpful,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Shield: core.ShieldConfig{
				Aura: core.Aura{
					Label:    "Scarab Brooch Shield",
					Duration: time.Second * 30,
				},
			},
		})

		activeAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Persistent Shield",
			Callback: core.CallbackOnHealDealt,
			Duration: time.Second * 30,
			Handler: func(sim *core.Simulation, _ *core.Spell, result *core.SpellResult) {
				shieldSpell.Shield(result.Target).Apply(sim, result.Damage*0.15)
			},
		})

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 3,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				activeAura.Activate(sim)
			},
		})

		character.AddMajorEquipmentCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=227967/second-wind
	// Use: Restores 30 mana every 1 sec for 10 sec. (2 Min Cooldown)
	core.NewItemEffect(SecondWind, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 15604}
		manaMetrics := character.NewManaMetrics(actionID)
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			ProcMask: core.ProcMaskEmpty,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 2,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				core.StartPeriodicAction(sim, core.PeriodicActionOptions{
					Period:   time.Second * 1,
					NumTicks: 10,
					Priority: core.ActionPriorityAuto,
					OnAction: func(sim *core.Simulation) {
						character.AddMana(sim, 30, manaMetrics)
					},
				})
			},
		})
		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=236334/slayers-crest
	// Use: Increases Attack Power by 280 for 20 sec. (2 Min Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(SlayersCrest, stats.Stats{stats.AttackPower: 280, stats.RangedAttackPower: 280}, time.Second*20, time.Minute*2)

	// https://www.wowhead.com/classic/item=228576/smolderwebs-eye#see-also
	// Use: Poisons target for 20 Nature damage every 2 sec for 20 sec. (2 Min Cooldown)
	core.NewItemEffect(SmolderwebsEye, func(agent core.Agent) {
		character := agent.GetCharacter()
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 17330},
			SpellSchool: core.SpellSchoolNature,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPoison | core.SpellFlagPureDot | core.SpellFlagNoOnCastComplete,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 2,
				},
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 2,
				Aura: core.Aura{
					Label: "Poison (Smolderweb's Eye)",
				},
				OnSnapshot: func(_ *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 20, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=13209/seal-of-the-dawn
	// Equip: +81 Attack Power when fighting Undead.
	core.NewMobTypeAttackPowerEffect(SealOfTheDawn, []proto.MobType{proto.MobType_MobTypeUndead}, 81)

	// https://www.wowhead.com/classic/item=237283/talisman-of-ascendance
	// Use: Your next 5 damage or healing spells cast within 20 seconds will grant a bonus of up to 40 damage and up to 75 healing, stacking up to 5 times.
	// Expires after 6 damage or healing spells or 20 seconds, whichever occurs first. (50 Sec Cooldown)
	core.NewItemEffect(TalismanOfAscendance, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{ItemID: TalismanOfAscendance}
		duration := time.Second * 20
		bonusPerStack := stats.Stats{
			stats.SpellDamage:  40,
			stats.HealingPower: 75,
		}

		buffAura := character.GetOrRegisterAura(core.Aura{
			ActionID:  actionID,
			Label:     "Ascendance",
			Duration:  duration,
			MaxStacks: 5,
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks, newStacks int32) {
				bonusStats := bonusPerStack.Multiply(float64(newStacks - oldStacks))
				character.AddStatsDynamic(sim, bonusStats)
			},
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
				if !spell.ProcMask.Matches(core.ProcMaskSpellDamage | core.ProcMaskSpellHealing) {
					return
				}

				if aura.GetStacks() == 5 {
					aura.Deactivate(sim)
				} else {
					aura.AddStack(sim)
				}
			},
		})

		cdSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Second * 50,
				},
				SharedCD: core.Cooldown{
					Timer:    character.GetOffensiveTrinketCD(),
					Duration: duration,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		character.AddMajorEquipmentCooldown(core.MajorCooldown{
			Spell: cdSpell,
			Type:  core.CooldownTypeDPS,
		})
	})

	// https://www.wowhead.com/classic/item=228255/talisman-of-ephemeral-power
	// Use: Increases damage and healing done by magical spells and effects by up to 184 for 15 sec. (1 Min, 30 Sec Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(TalismanOfEphemeralPower, stats.Stats{stats.SpellPower: 184}, time.Second*15, time.Second*90)

	// https://www.wowhead.com/classic/item=233628/the-burrowers-shell#see-also
	// Use: Absorbs 1200 damage.  Lasts 20 sec. (2 Min Cooldown)
	core.NewItemEffect(TheBurrowersShield, func(agent core.Agent) {
		character := agent.GetCharacter()

		shieldSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 1213375},
			SpellSchool: core.SpellSchoolHoly,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell | core.SpellFlagHelpful,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Shield: core.ShieldConfig{
				Aura: core.Aura{
					Label:    "The Burrower's Shield",
					Duration: time.Second * 20,
				},
			},
		})

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{ItemID: TheBurrowersShield},
			SpellSchool: core.SpellSchoolPhysical,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 2,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				shieldSpell.Shield(spell.Unit).Apply(sim, 1200)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeSurvival,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=236320/warmth-of-forgiveness
	// Use: Restores 6000 mana. (3 Min Cooldown)
	core.NewItemEffect(WarmthOfForgiveness, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{ItemID: WarmthOfForgiveness}
		manaMetrics := character.NewManaMetrics(actionID)

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolPhysical,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(), // Doesn't share the trinket timer
					Duration: time.Minute * 3,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.AddMana(sim, 6000, manaMetrics)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=19948/zandalarian-hero-badge
	// Increases your armor by 2000 and defense skill by 30 for 20 sec.
	// Every time you take melee or ranged damage, this bonus is reduced by 200 armor and 3 defense.
	core.NewItemEffect(ZandalariHeroBadge, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{ItemID: ZandalariHeroBadge}
		duration := time.Second * 20
		bonusPerStack := stats.Stats{
			stats.Armor:   200,
			stats.Defense: 3,
		}

		buffAura := character.RegisterAura(core.Aura{
			Label:     "Fragile Armor",
			ActionID:  actionID,
			Duration:  duration,
			MaxStacks: 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.SetStacks(sim, aura.MaxStacks)
			},
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks, newStacks int32) {
				bonusStats := bonusPerStack.Multiply(float64(newStacks - oldStacks))
				character.AddStatsDynamic(sim, bonusStats)
			},
			OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if spell.ProcMask.Matches(core.ProcMaskMeleeOrRanged) && result.Landed() {
					aura.RemoveStack(sim)
				}
			},
		})

		cdSpell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 2,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		character.AddMajorEquipmentCooldown(core.MajorCooldown{
			Spell: cdSpell,
			Type:  core.CooldownTypeSurvival,
		})
	})

	// https://www.wowhead.com/classic/item=19950/zandalarian-hero-charm
	// Increases your spell damage by up to 204 and your healing by up to 408 for 20 sec.
	// Every time you cast a spell, the bonus is reduced by 17 spell damage and 34 healing.
	core.NewItemEffect(ZandalariHeroCharm, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{ItemID: ZandalariHeroCharm}
		duration := time.Second * 20
		bonusPerStack := stats.Stats{
			stats.SpellDamage:  17,
			stats.HealingPower: 34,
		}

		buffAura := character.RegisterAura(core.Aura{
			ActionID:  actionID,
			Label:     "Unstable Power",
			Duration:  duration,
			MaxStacks: 12,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.SetStacks(sim, aura.MaxStacks)
			},
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks, newStacks int32) {
				bonusStats := bonusPerStack.Multiply(float64(newStacks - oldStacks))
				character.AddStatsDynamic(sim, bonusStats)
			},
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
				if !spell.ProcMask.Matches(core.ProcMaskSpellDamage) {
					return
				}
				aura.RemoveStack(sim)
			},
		})

		cdSpell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 2,
				},
				SharedCD: core.Cooldown{
					Timer:    character.GetOffensiveTrinketCD(),
					Duration: duration,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		character.AddMajorEquipmentCooldown(core.MajorCooldown{
			Spell: cdSpell,
			Type:  core.CooldownTypeDPS,
		})
	})

	core.NewItemEffect(ZandalariHeroMedallion, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{ItemID: ZandalariHeroMedallion}
		duration := time.Second * 20

		buffAura := character.RegisterAura(core.Aura{
			ActionID:  actionID,
			Label:     "Restless Strength",
			Duration:  duration,
			MaxStacks: 20,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.SetStacks(sim, aura.MaxStacks)
			},
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks, newStacks int32) {
				character.PseudoStats.BonusPhysicalDamage += 2 * float64(newStacks-oldStacks)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					aura.RemoveStack(sim)
				}
			},
		})

		cdSpell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 2,
				},
				SharedCD: core.Cooldown{
					Timer:    character.GetOffensiveTrinketCD(),
					Duration: duration,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		character.AddMajorEquipmentCooldown(core.MajorCooldown{
			Spell: cdSpell,
			Type:  core.CooldownTypeDPS,
		})
	})

	///////////////////////////////////////////////////////////////////////////
	//                                 Other
	///////////////////////////////////////////////////////////////////////////

	// https://www.wowhead.com/classic/item=228354/blazefury-medallion
	// Equip: Adds 2 fire damage to your melee attacks.
	core.NewItemEffect(BlazefuryMedallion, func(agent core.Agent) {
		character := agent.GetCharacter()
		BlazefuryTriggerAura(character, BlazefuryMedallion, 7711, 7712, core.SpellSchoolFire, 2)
	})

	// https://www.wowhead.com/classic/item=14554/cloudkeeper-legplates
	// Use: Increases Attack Power by 100 for 30 sec. (15 Min Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(CloudkeeperLegplates, stats.Stats{stats.AttackPower: 100, stats.RangedAttackPower: 100}, time.Second*30, time.Minute*15)

	// https://www.wowhead.com/classic/item=228266/drillborer-disk
	// Equip: When struck in combat inflicts 3 Arcane damage to the attacker.
	core.NewItemEffect(DrillborerDisk, func(agent core.Agent) {
		thornsArcaneDamageEffect(agent, DrillborerDisk, "Drillborer Disk", 3)
	})
	core.NewItemEffect(DrillborerDiskMolten, func(agent core.Agent) {
		thornsArcaneDamageEffect(agent, DrillborerDiskMolten, "Drillborer Disk (Molten)", 3)
	})

	// https://www.wowhead.com/classic/item=18168/force-reactive-disk
	// Equip: When the shield blocks it releases an electrical charge that damages all nearby enemies. (1s cooldown)
	core.NewItemEffect(ForceReactiveDisk, func(agent core.Agent) {
		character := agent.GetCharacter()

		procSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{ItemID: ForceReactiveDisk},
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, 25, spell.OutcomeMagicHitAndCrit)
				}
			},
		})

		triggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Force Reactive Disk",
			Callback: core.CallbackOnSpellHitTaken,
			ProcMask: core.ProcMaskMelee,
			Outcome:  core.OutcomeBlock,
			ICD:      time.Second,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procSpell.Cast(sim, spell.Unit)
			},
		})

		character.ItemSwap.RegisterProc(ForceReactiveDisk, triggerAura)
	})

	// https://www.wowhead.com/classic/item=11669/naglering
	// Equip: When struck in combat inflicts 3 Arcane damage to the attacker.
	core.NewItemEffect(Naglering, func(agent core.Agent) {
		thornsArcaneDamageEffect(agent, Naglering, "Naglering", 3)
	})

	// https://www.wowhead.com/classic/item=18326/razor-gauntlets
	// Equip: When struck in combat inflicts 3 Arcane damage to the attacker.
	core.NewItemEffect(RazorGauntlets, func(agent core.Agent) {
		thornsArcaneDamageEffect(agent, RazorGauntlets, "Razor Gauntlets", 3)
	})

	// https://www.wowhead.com/classic/item=1168/skullflame-shield
	// Equip: When struck in combat has a 3% chance of stealing 35 life from target enemy. (Proc chance: 3%)
	// Equip: When struck in combat has a 1% chance of dealing 75 to 125 Fire damage to all targets around you. (Proc chance: 1%)
	core.NewItemEffect(SkullflameShield, func(agent core.Agent) {
		character := agent.GetCharacter()

		drainLifeActionID := core.ActionID{SpellID: 18817}
		healthMetrics := character.NewHealthMetrics(drainLifeActionID)
		drainLifeSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    drainLifeActionID,
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BonusCoefficient: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 35, spell.OutcomeAlwaysHit)
				character.GainHealth(sim, result.Damage, healthMetrics)
			},
		})

		flamestrikeSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18818},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, sim.Roll(75, 125), spell.OutcomeMagicHit)
				}
			},
		})

		drainLifeTriggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Drain Life Trigger",
			Callback:   core.CallbackOnSpellHitTaken,
			Outcome:    core.OutcomeLanded,
			ProcMask:   core.ProcMaskMelee,
			ProcChance: 0.03,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				drainLifeSpell.Cast(sim, spell.Unit)
			},
		})

		flameStrikeTriggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Flamestrike Trigger",
			Callback:   core.CallbackOnSpellHitTaken,
			Outcome:    core.OutcomeLanded,
			ProcMask:   core.ProcMaskMelee,
			ProcChance: 0.01,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				flamestrikeSpell.Cast(sim, spell.Unit)
			},
		})

		character.ItemSwap.RegisterProc(SkullflameShield, drainLifeTriggerAura)
		character.ItemSwap.RegisterProc(SkullflameShield, flameStrikeTriggerAura)
	})

	// https://www.wowhead.com/classic/item=234463/wrath-of-cenarius
	// Gives a chance when your harmful spells land to increase the damage of your spells and effects by 193 for 10 sec.
	// (Proc chance: 5%)
	core.NewItemEffect(WrathOfCenarius, func(agent core.Agent) {
		character := agent.GetCharacter()
		buffAura := character.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 1214279},
			Label:    "Spell Blasting",
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatDynamic(sim, stats.SpellDamage, 193)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatDynamic(sim, stats.SpellDamage, -193)
			},
		})

		triggerAura := core.MakeProcTriggerAura(&agent.GetCharacter().Unit, core.ProcTrigger{
			Name:              "Spell Blasting Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskSpellDamage,
			SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
			CanProcFromProcs:  true,
			ProcChance:        0.05,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				buffAura.Activate(sim)
			},
		})

		character.ItemSwap.RegisterProc(SkullflameShield, triggerAura)
	})

	core.AddEffectsToTest = true
}

// https://www.wowhead.com/classic/spell=1213355/shock
// Chance on hit: Instantly lightning shocks the target for 339 to 377 Nature damage.
// SP mod: 0.15
// Used by:
// - https://www.wowhead.com/classic/item=233621/kalimdors-revenge
// - https://www.wowhead.com/classic/item=234981/kalimdors-revenge
func KalimdorsRevengeDamageProc(character *core.Character) *core.Spell {
	procMaskOnAuto := core.ProcMaskDamageProc     // Both spell and melee procs
	procMaskOnSpecial := core.ProcMaskSpellDamage // TODO: check if core.ProcMaskSpellDamage remains on special

	return character.GetOrRegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 1213355},
		SpellSchool: core.SpellSchoolNature,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    procMaskOnAuto,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: 0.15,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			if spell.ProcMask.Matches(core.ProcMaskMeleeSpecial) {
				spell.ProcMask = procMaskOnSpecial
			} else {
				spell.ProcMask = procMaskOnAuto
			}
			spell.CalcAndDealDamage(sim, target, sim.RollWithLabel(339, 377, "Kalimdor's Revenge"), spell.OutcomeMagicHitAndCrit)
		},
	})
}

// https://www.wowhead.com/classic/spell=446327/enrage
// Used by:
// - https://www.wowhead.com/classic/item=220569/blistering-ragehammer and
// - https://www.wowhead.com/classic/item=228125/refined-arcanite-champion
func EnrageAura446327(character *core.Character) *core.Aura {
	return character.RegisterAura(core.Aura{
		ActionID: core.ActionID{SpellID: 446327},
		Label:    "Enrage (446327)",
		Duration: time.Second * 15,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			character.PseudoStats.BonusPhysicalDamage += 20
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			character.PseudoStats.BonusPhysicalDamage -= 20
		},
	}).AttachMultiplyAttackSpeed(&character.Unit, 1.05)
}

const BlazefuryAuraTag = "Blazefury"

func BlazefuryTriggerAura(character *core.Character, itemID int32, triggerSpellID int32, damageSpellID int32, spellSchool core.SpellSchool, damage float64) {
	if character.HasAuraWithTag(BlazefuryAuraTag) {
		return
	}

	procSpell := character.RegisterSpell(core.SpellConfig{
		ActionID:         core.ActionID{SpellID: damageSpellID},
		SpellSchool:      spellSchool,
		DefenseType:      core.DefenseTypeMagic,
		ProcMask:         core.ProcMaskMeleeDamageProc,
		Flags:            core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell | core.SpellFlagSuppressEquipProcs,
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMagicHitAndCrit)
		},
	})

	procMaskOnAuto := core.ProcMaskDamageProc // Both spell and melee proc combo
	procMaskOnSpecial := core.ProcMaskEmpty

	triggerAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
		Name:              fmt.Sprintf("Blazefury Trigger (%d)", triggerSpellID),
		Tag:               BlazefuryAuraTag,
		Callback:          core.CallbackOnSpellHitDealt,
		Outcome:           core.OutcomeLanded,
		ProcMask:          core.ProcMaskMelee,
		SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
		Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			procSpell.ProcMask = core.Ternary(spell.ProcMask.Matches(core.ProcMaskMeleeSpecial), procMaskOnSpecial, procMaskOnAuto)
			procSpell.Cast(sim, result.Target)
		},
	})

	character.ItemSwap.RegisterProc(itemID, triggerAura)
}

// TODO: This is treated as a buff, NOT a debuff in-game
// We don't have the ability to remove resistances for only one agent at a time right now
func bonereaversEdgeEffect(character *core.Character) *core.Spell {
	actionID := core.ActionID{SpellID: 21153}
	buffAura := character.RegisterAura(core.Aura{
		ActionID:  actionID,
		Label:     "Bonereaver's Edge",
		Duration:  time.Second * 10,
		MaxStacks: 3,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks, newStacks int32) {
			for _, target := range sim.Encounter.TargetUnits {
				target.AddStatDynamic(sim, stats.Armor, 700*float64(oldStacks))
				target.AddStatDynamic(sim, stats.Armor, -700*float64(newStacks))
			}
		},
	})
	return character.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskEmpty,
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			buffAura.Activate(sim)
			buffAura.AddStack(sim)
		},
	})
}

// Chance on hit: Reduces an enemy's Strength by 125 and its Stamina by 50 for 2 min.
// Equip: When struck in combat has a chance of causing the attacker to flee in terror for 2 seconds. (Proc chance: 2%)
func dreadbladeOfTheDestructorEffect(character *core.Character) *core.Spell {
	actionID := core.ActionID{SpellID: 462178}
	procAuras := character.NewEnemyAuraArray(func(target *core.Unit, _ int32) *core.Aura {
		return target.RegisterAura(core.Aura{
			ActionID: actionID,
			Label:    "Enfeeble (Dreadblade of the Destructor)",
			Duration: time.Minute * 2,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatsDynamic(sim, stats.Stats{
					stats.Stamina:  -50,
					stats.Strength: -125,
				})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatsDynamic(sim, stats.Stats{
					stats.Stamina:  50,
					stats.Strength: 125,
				})
			},
		})
	})

	core.MakePermanent(character.RegisterAura(core.Aura{
		Label:      "Cursed Blade",
		ActionID:   core.ActionID{SpellID: 462228},
		BuildPhase: core.CharacterBuildPhaseBuffs,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			character.AddBuildPhaseStatDynamic(sim, stats.MeleeCrit, 2*core.CritRatingPerCritChance)
			character.PseudoStats.MeleeSpeedMultiplier *= 1.05
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			character.AddBuildPhaseStatDynamic(sim, stats.MeleeCrit, -2*core.CritRatingPerCritChance)
			character.PseudoStats.MeleeSpeedMultiplier /= 1.05
		},
	}))

	return character.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskEmpty,
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			procAuras.Get(target).Activate(sim)
		},
	})
}

func thornsArcaneDamageEffect(agent core.Agent, itemID int32, itemName string, damage float64) {
	character := agent.GetCharacter()
	character.PseudoStats.ThornsDamage += damage

	procSpell := character.GetOrRegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{ItemID: itemID},
		SpellSchool: core.SpellSchoolArcane,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagBinary | core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMagicHit)
		},
	})

	procAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
		Name:     fmt.Sprintf("Thorns (%s)", itemName),
		ProcMask: core.ProcMaskMelee,
		Callback: core.CallbackOnSpellHitTaken,
		Outcome:  core.OutcomeLanded,
		Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			procSpell.Cast(sim, spell.Unit)
		},
	})

	character.ItemSwap.RegisterProc(itemID, procAura)
}

func eskhandarsRightClawAura(character *core.Character) *core.Aura {
	return character.GetOrRegisterAura(core.Aura{
		Label:    "Eskhandar's Rage",
		ActionID: core.ActionID{SpellID: 22640},
		Duration: time.Second * 5,
	}).AttachMultiplyAttackSpeed(&character.Unit, 1.25)
}

func gutgoreRipperEffect(character *core.Character) *core.Spell {
	procAuras := character.NewEnemyAuraArray(func(target *core.Unit, _ int32) *core.Aura {
		return target.GetOrRegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 461682},
			Label:    "Gutgore Ripper",
			Duration: time.Second * 30,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatsDynamic(sim, stats.Stats{
					stats.Agility:   -25,
					stats.Intellect: -25,
					stats.Stamina:   -25,
					stats.Spirit:    -25,
					stats.Strength:  -25,
				})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatsDynamic(sim, stats.Stats{
					stats.Agility:   25,
					stats.Intellect: 25,
					stats.Stamina:   25,
					stats.Spirit:    25,
					stats.Strength:  25,
				})
			},
		})
	})

	return character.GetOrRegisterSpell(core.SpellConfig{
		ActionID:         core.ActionID{SpellID: 461682},
		SpellSchool:      core.SpellSchoolShadow,
		DefenseType:      core.DefenseTypeMagic,
		ProcMask:         core.ProcMaskEmpty,
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcAndDealDamage(sim, target, 150, spell.OutcomeMagicHitAndCrit)
			if result.Landed() {
				procAuras.Get(target).Activate(sim)
			}
		},
	})
}

// Chance on hit: Spell damage taken by target increased by 15% for 5 sec.
// func nightfallProc(character *core.Character, itemName string) {
// 	procAuras := character.NewEnemyAuraArray(func(target *core.Unit, _ int32) *core.Aura {
// 		return target.RegisterAura(core.Aura{
// 			Label:    fmt.Sprintf("Spell Vulnerability (%s)", itemName),
// 			ActionID: core.ActionID{SpellID: 23605},
// 			Duration: time.Second * 5,
// 			OnGain: func(aura *core.Aura, sim *core.Simulation) {
// 				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexArcane] *= 1.15
// 				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFire] *= 1.15
// 				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFrost] *= 1.15
// 				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] *= 1.15
// 				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexNature] *= 1.15
// 				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexShadow] *= 1.15
// 			},
// 			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
// 				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexArcane] /= 1.15
// 				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFire] /= 1.15
// 				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFrost] /= 1.15
// 				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] /= 1.15
// 				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexNature] /= 1.15
// 				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexShadow] /= 1.15
// 			},
// 		})
// 	})

// 	core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
// 		Name:              fmt.Sprintf("%s Trigger", itemName),
// 		Callback:          core.CallbackOnSpellHitDealt,
// 		Outcome:           core.OutcomeLanded,
// 		ProcMask:          core.ProcMaskMelee,
// 		SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
// 		PPM:               2,
// 		Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
// 			procAuras.Get(result.Target).Activate(sim)
// 		},
// 	})
// }

func StrengthOfTheChampionAura(character *core.Character) *core.Aura {
	actionID := core.ActionID{SpellID: 16916}
	healthMetrics := character.NewHealthMetrics(actionID)
	return character.GetOrRegisterAura(core.Aura{
		Label:    "Strength of the Champion",
		ActionID: actionID,
		Duration: time.Second * 30,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			character.GainHealth(sim, sim.Roll(270, 450), healthMetrics)
			character.AddStatDynamic(sim, stats.Strength, 120)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			character.AddStatDynamic(sim, stats.Strength, -120)
		},
	})
}

func manslayerOfTheQirajiAura(character *core.Character) *core.Aura {

	return core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
		Name:              "Manslayer Of The Qiraji",
		Duration:          core.NeverExpires,
		Callback:          core.CallbackOnSpellHitDealt,
		Outcome:           core.OutcomeLanded,
		SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
		ProcMask:          core.ProcMaskMelee,
		ProcChance:        0.01,
		ICD:               time.Millisecond * 200,
		Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			character.Unit.AutoAttacks.ExtraMHAttackProc(sim, 1, core.ActionID{SpellID: 1214927}, spell)
		},
	})
}

func neretzekBloodDrinkerEffect(character *core.Character) *core.Spell { // PPM based on old ppm from Armamaments discord
	actionID := core.ActionID{SpellID: 1214208}
	healthMetrics := character.NewHealthMetrics(actionID)
	return character.RegisterSpell(core.SpellConfig{
		ActionID:         actionID,
		SpellSchool:      core.SpellSchoolShadow,
		DefenseType:      core.DefenseTypeMagic,
		ProcMask:         core.ProcMaskEmpty,
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcAndDealDamage(sim, target, sim.Roll(171, 193), spell.OutcomeMagicHit)
			character.GainHealth(sim, result.Damage, healthMetrics)
		},
	})
}
