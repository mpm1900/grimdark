package game

import "maps"

func (a *Actor) mapBaseStat(stat Stat, stats map[Stat]float64, offset float64) float64 {
	base := stats[stat]*2 + offset
	ratio := float64(a.Level) / 100
	result := (base*ratio + 5)
	if stat == Health {
		result += float64(a.Level)
	}
	return result
}

func (a *Actor) getStatOffset(stat Stat) float64 {
	offset, ok := a.OffsetStats[stat]
	if !ok {
		offset = 0
	}

	for _, w := range a.Weapons {
		weapon_offset, ok := w.OffsetStats[stat]
		if ok {
			offset += weapon_offset
		}
	}

	return offset
}

func (a *Actor) mapBaseStats() {
	a.UnmodifiedStats = maps.Clone(a.Stats)

	for stat, _ := range a.Stats {
		if _, ok := mappedStats[stat]; !ok {
			continue
		}

		a.Stats[stat] = a.mapBaseStat(stat, a.Stats, a.getStatOffset(stat))
		a.UnmodifiedStats[stat] = a.mapBaseStat(stat, a.UnmodifiedStats, 0)
	}
}

func (a *Actor) mapStagedStats() {
	builder := NewStageBuilder(a.Stages)
	stats := builder.ResolveAll(a.Stats)

	builder.Mod = 3
	accuracy := builder.Resolve(Accuracy, a.Stats[Accuracy])
	evasion := builder.Resolve(Evasion, a.Stats[Evasion])
	crit_damage := builder.Resolve(CriticalDamage, a.Stats[CriticalDamage])
	effect_chance := builder.Resolve(EffectChance, a.Stats[EffectChance])

	a.Stats = stats
	a.Stats[Accuracy] = accuracy
	a.Stats[Evasion] = evasion
	a.Stats[CriticalDamage] = crit_damage
	a.Stats[EffectChance] = effect_chance
}
