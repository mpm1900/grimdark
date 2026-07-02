import type { Actor } from '#/lib/game/actor'
import { NULL_CONTEXT, type Context } from '#/lib/game/context'
import { useEffect, useState } from 'react'

function useContext(
  targets_context: Context,
  target_type?: 'position_IDs' | 'actor_IDs'
) {
  target_type =
    (target_type ?? targets_context.position_IDs.length > 0)
      ? 'position_IDs'
      : 'actor_IDs'
  const [value, set] = useState({
    ...NULL_CONTEXT,
    action_ID: targets_context.action_ID,
    source_ID: targets_context.source_ID,
    parent_ID: targets_context.parent_ID,
    player_ID: targets_context.player_ID,
  })

  function reset() {
    set({
      ...NULL_CONTEXT,
      action_ID: targets_context.action_ID,
      source_ID: targets_context.source_ID,
      parent_ID: targets_context.parent_ID,
      player_ID: targets_context.player_ID,
    })
  }

  useEffect(() => {
    reset()
  }, [targets_context.source_ID])

  return {
    value,
    set,
    addTarget: (target: Actor) => {
      if (target_type === 'position_IDs' && target.position_ID) {
        set((c) => ({
          ...c,
          position_IDs: [...c.position_IDs, target.position_ID as string],
        }))
      }
      if (target_type === 'actor_IDs') {
        set((c) => ({
          ...c,
          actor_IDs: [...c.actor_IDs, target.ID],
        }))
      }
    },
    hasTarget: (target: Actor) => {
      if (target_type === 'position_IDs' && target.position_ID) {
        return value.position_IDs.includes(target.position_ID)
      }
      if (target_type === 'actor_IDs') {
        return value.actor_IDs.includes(target.ID)
      }

      return false
    },
    removeTarget: (target: Actor) => {
      set((c) => ({
        ...c,
        actor_IDs: c.actor_IDs.filter((id) => id !== target.ID),
        position_IDs: c.position_IDs.filter(
          (pid) => pid !== target.position_ID
        ),
      }))
    },
    reset,
  }
}

export { useContext }
