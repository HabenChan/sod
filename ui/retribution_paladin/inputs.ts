import * as InputHelpers from '../core/components/input_helpers.js';
import { Player } from '../core/player.js';
import { Spec } from '../core/proto/common.js';
import { PaladinAura, PaladinSeal } from '../core/proto/paladin.js';
import { ActionId } from '../core/proto_utils/action_id.js';
import { TypedEvent } from '../core/typed_event.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const AuraSelection = InputHelpers.makeSpecOptionsEnumIconInput<Spec.SpecRetributionPaladin, PaladinAura>({
	fieldName: 'aura',
	values: [
		{ value: PaladinAura.NoPaladinAura, tooltip: 'No Aura' },
		{
			actionId: () => ActionId.fromSpellId(20218),
			value: PaladinAura.SanctityAura,
		},
	],
});

// The below is used in the custom APL action "Cast Primary Seal".
// Only shows SoC if it's talented, only shows SoM if the relevant rune is equipped.
export const PrimarySealSelection = InputHelpers.makeSpecOptionsEnumIconInput<Spec.SpecRetributionPaladin, PaladinSeal>({
	fieldName: 'primarySeal',
	values: [
		{
			actionId: player =>
				player.getMatchingSpellActionId([
					{ id: 20154, maxLevel: 9 },
					{ id: 20287, minLevel: 10, maxLevel: 17 },
					{ id: 20288, minLevel: 18, maxLevel: 25 },
					{ id: 20289, minLevel: 26, maxLevel: 33 },
					{ id: 20290, minLevel: 34, maxLevel: 41 },
					{ id: 20291, minLevel: 42, maxLevel: 49 },
					{ id: 20292, minLevel: 50, maxLevel: 57 },
					{ id: 20293, minLevel: 58 },
				]),
			value: PaladinSeal.Righteousness,
		},
		{
			actionId: player =>
				player.getMatchingSpellActionId([
					{ id: 20375, maxLevel: 29 },
					{ id: 20915, minLevel: 30, maxLevel: 39 },
					{ id: 20918, minLevel: 40, maxLevel: 49 },
					{ id: 20919, minLevel: 50, maxLevel: 59 },
					{ id: 20920, minLevel: 60 },
				]),
			value: PaladinSeal.Command,
			showWhen: (player: Player<Spec.SpecRetributionPaladin>) => player.getTalents().sealOfCommand,
		},
		{
			actionId: () => ActionId.fromSpellId(407798),
			value: PaladinSeal.Martyrdom,
		},
		{
			actionId: player =>
				player.getMatchingSpellActionId([
					{ id: 21082, minLevel: 6, maxLevel: 11 },
					{ id: 20162, minLevel: 12, maxLevel: 21 },
					{ id: 20305, minLevel: 22, maxLevel: 31 },
					{ id: 20306, minLevel: 32, maxLevel: 41 },
					{ id: 20307, minLevel: 42, maxLevel: 51 },
					{ id: 20308, minLevel: 52 },
				]),
			value: PaladinSeal.Crusader,
		},
	],
	// changeEmitter: (player: Player<Spec.SpecRetributionPaladin>) => player.changeEmitter,
	changeEmitter: (player: Player<Spec.SpecRetributionPaladin>) =>
		TypedEvent.onAny([player.gearChangeEmitter, player.talentsChangeEmitter, player.specOptionsChangeEmitter, player.levelChangeEmitter]),
});

export const CrusaderStrikeStopAttack = InputHelpers.makeSpecOptionsBooleanInput<Spec.SpecRetributionPaladin>({
	fieldName: 'isUsingCrusaderStrikeStopAttack',
	label: 'Using Crusader Strike /stopattack macro',
	labelTooltip: 'Allows saving of extra attacks',
});

export const DivineStormStopAttack = InputHelpers.makeSpecOptionsBooleanInput<Spec.SpecRetributionPaladin>({
	fieldName: 'isUsingDivineStormStopAttack',
	label: 'Using Divine Storm /stopattack macro',
	labelTooltip: 'Allows saving of extra attacks',
});

export const JudgementStopAttack = InputHelpers.makeSpecOptionsBooleanInput<Spec.SpecRetributionPaladin>({
	fieldName: 'isUsingJudgementStopAttack',
	label: 'Using Judgement /stopattack macro',
	labelTooltip: 'Allows saving of extra attacks',
});

export const ExorcismStopAttack = InputHelpers.makeSpecOptionsBooleanInput<Spec.SpecRetributionPaladin>({
	fieldName: 'isUsingExorcismStopAttack',
	label: 'Using Exorcism /stopattack macro',
	labelTooltip: 'Makes for more precise timings when twisting close to or below 3s swing speed',
});

export const ManualStartAttack = InputHelpers.makeSpecOptionsBooleanInput<Spec.SpecRetributionPaladin>({
	fieldName: 'isUsingManualStartAttack',
	label: 'Using manual /startattack macro (for stacking only)',
	labelTooltip: 'Mainly useful for slow seal stacking utilizing /startattack macros',
});
