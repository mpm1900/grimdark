import type { Affinity } from '#/lib/game/core'
import { FaBiohazard } from 'react-icons/fa6'
import {
  GiPowerLightning,
  GiSmallFire,
  GiHadesSymbol,
  GiBeveledStar,
  GiBleedingEye,
  GiPerpendicularRings,
} from 'react-icons/gi'
import type { IconType } from 'react-icons/lib'

export const AFFINITY_ICONS: Record<Affinity, IconType> = {
  arcane: GiHadesSymbol,
  cryo: GiBeveledStar,
  fire: GiSmallFire,
  kinetic: GiPerpendicularRings,
  lightning: GiPowerLightning,
  poison: FaBiohazard,
  psychic: GiBleedingEye,
}
