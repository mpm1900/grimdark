import {
  GiBoots,
  GiNecklace,
  GiShinyApple,
  GiSkullSignet,
} from 'react-icons/gi'
import type { IconType } from 'react-icons/lib'

export const ITEM_ICONS: Record<string, IconType> = {
  // corrupted necklace
  '019fca79-fdc7-757d-9214-5a4952a86358': GiNecklace,
  // cursed boots
  '019fca52-9a89-77f8-a58d-8e5933f5e291': GiBoots,
  // cursed ring
  '019fca3b-2cee-76f8-8f07-b576848c4026': GiSkullSignet,
  // rations
  '019fcab0-c659-779d-a8e3-a440d272561a': GiShinyApple,
}
