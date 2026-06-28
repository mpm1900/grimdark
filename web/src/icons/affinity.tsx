import type { Affinity } from '#/lib/game/core'
import { BsFire } from 'react-icons/bs'
import { FaHandSparkles } from 'react-icons/fa6'
import {
  GiBrain,
  GiAllForOne,
  GiPowerLightning,
  GiSkullWithSyringe,
  GiSnowflake1,
} from 'react-icons/gi'
import type { IconType } from 'react-icons/lib'

export const AFFINITY_ICONS: Record<Affinity, IconType> = {
  arcane: FaHandSparkles,
  cryo: GiSnowflake1,
  fire: BsFire,
  kinetic: GiAllForOne,
  lightning: GiPowerLightning,
  poison: GiSkullWithSyringe,
  psychic: GiBrain,
}
