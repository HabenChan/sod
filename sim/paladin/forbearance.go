package paladin

import (
	"time"

	"github.com/wowsims/sod/sim/core"
)

func (paladin *Paladin) registerForbearance() {

	actionID := core.ActionID{SpellID: 25771}

	forbearanceAura := paladin.RegisterAura(core.Aura{
		Label:    "Forbearance",
		ActionID: actionID,
		Duration: time.Minute * 1,
	})

	paladin.OnSpellRegistered(func(spell *core.Spell) {

		if spell.Flags.Matches(SpellFlag_Forbearance) {
			if !paladin.bypassAvengingWrathForbearance || spell.ClassSpellMask != ClassSpellMask_PaladinAvengingWrath {
				oldEffect := spell.ApplyEffects

				spell.ApplyEffects = func(sim *core.Simulation, unit *core.Unit, spell *core.Spell) {
					oldEffect(sim, unit, spell)
					forbearanceAura.Activate(sim)
				}
			}

			if spell.ExtraCastCondition != nil {
				oldCondition := spell.ExtraCastCondition
				spell.ExtraCastCondition = func(sim *core.Simulation, target *core.Unit) bool {
					return (!forbearanceAura.IsActive()) && oldCondition(sim, target)
				}
			} else {
				spell.ExtraCastCondition = func(sim *core.Simulation, target *core.Unit) bool {
					return !forbearanceAura.IsActive()
				}
			}
		}
	})
}
