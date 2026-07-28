import type { Affinity } from '#/lib/game/core'
import {
  GiPowerLightning,
  GiSmallFire,
  GiHadesSymbol,
  GiBeveledStar,
  GiBleedingEye,
  GiPerpendicularRings,
  GiDeathJuice,
  GiPoison,
  GiLightningTree,
  GiJerusalemCross,
  GiHeptagram,
} from 'react-icons/gi'
import type { IconType } from 'react-icons/lib'

export const AFFINITY_ICONS: Record<Affinity, IconType> = {
  arcane: GiHeptagram,
  blood: GiDeathJuice,
  holy: GiJerusalemCross,
  fire: GiSmallFire,
  physical: GiPerpendicularRings,
  lightning: GiLightningTree,
  poison: GiPoison,
  psychic: GiBleedingEye,
}
