import type { Affinity } from '#/lib/game/core'
import {
  GiPowerLightning,
  GiSmallFire,
  GiHadesSymbol,
  GiBeveledStar,
  GiBleedingEye,
  GiPerpendicularRings,
  GiFly,
  GiDeathJuice,
  GiHolySymbol,
} from 'react-icons/gi'
import type { IconType } from 'react-icons/lib'

export const AFFINITY_ICONS: Record<Affinity, IconType> = {
  arcane: GiHadesSymbol,
  blood: GiDeathJuice,
  holy: GiHolySymbol,
  fire: GiSmallFire,
  physical: GiPerpendicularRings,
  lightning: GiPowerLightning,
  poison: GiFly,
  psychic: GiBleedingEye,
}
