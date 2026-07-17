import { gameStore } from '#/lib/stores/game'
import { useSelector } from '@tanstack/react-store'
import { useLayoutEffect, useMemo, useRef, useState } from 'react'

type Point = {
  x: number
  y: number
}

type Arrow = {
  id: string
  path: string
  head: string
}

type ArrowLayout = {
  width: number
  height: number
  arrows: Arrow[]
}

const EMPTY_LAYOUT: ArrowLayout = {
  width: 0,
  height: 0,
  arrows: [],
}

function getAnchor(rect: DOMRect, parentRect: DOMRect): Point {
  return {
    x: rect.left - parentRect.left + rect.width / 2,
    y: rect.bottom - parentRect.top - rect.height * 0.3,
  }
}

function getArrow(start: Point, end: Point, index: number, count: number) {
  const dx = end.x - start.x
  const dy = end.y - start.y
  const length = Math.hypot(dx, dy)

  if (length === 0) return null

  const startPadding = 32
  const endPadding = 28
  const headLength = 32
  const headWidth = 32
  const from = {
    x: start.x + (dx / length) * startPadding,
    y: start.y + (dy / length) * startPadding,
  }
  const tip = {
    x: end.x - (dx / length) * endPadding,
    y: end.y - (dy / length) * endPadding,
  }
  const normal = {
    x: -dy / length,
    y: dx / length,
  }
  const arc = count > 1 ? (index - (count - 1) / 2) * 220 : 0
  const lift = count > 1 ? -170 : -60
  const control = {
    x: (from.x + tip.x) / 2 + normal.x * arc,
    y: (from.y + tip.y) / 2 + normal.y * arc + lift,
  }
  const tangent = {
    x: tip.x - control.x,
    y: tip.y - control.y,
  }
  const tangentLength = Math.hypot(tangent.x, tangent.y)

  if (tangentLength === 0) return null

  const unit = {
    x: tangent.x / tangentLength,
    y: tangent.y / tangentLength,
  }
  const headNormal = {
    x: -unit.y,
    y: unit.x,
  }
  const base = {
    x: tip.x - unit.x * headLength,
    y: tip.y - unit.y * headLength,
  }
  const left = {
    x: base.x + headNormal.x * (headWidth / 2),
    y: base.y + headNormal.y * (headWidth / 2),
  }
  const right = {
    x: base.x - headNormal.x * (headWidth / 2),
    y: base.y - headNormal.y * (headWidth / 2),
  }

  return {
    path: `M ${from.x} ${from.y} Q ${control.x} ${control.y} ${base.x} ${base.y}`,
    head: `${tip.x},${tip.y} ${left.x},${left.y} ${right.x},${right.y}`,
  }
}

function queryActor(board: HTMLElement, actor_ID: string) {
  return [...board.querySelectorAll<HTMLElement>('[data-actor-id]')].find(
    (element) => element.dataset.actorId === actor_ID
  )
}

function queryPosition(board: HTMLElement, position_ID: string) {
  return [...board.querySelectorAll<HTMLElement>('[data-position-id]')].find(
    (element) => element.dataset.positionId === position_ID
  )
}

function TargetingArrows() {
  const overlayRef = useRef<HTMLDivElement>(null)
  const active_context = useSelector(gameStore, (g) => g.active_context)
  const actors = useSelector(gameStore, (g) => g.actors)
  const positions = useSelector(gameStore, (g) => g.positions)
  const [layout, setLayout] = useState<ArrowLayout>(EMPTY_LAYOUT)

  const target_position_IDs = useMemo(() => {
    if (!active_context?.source_ID) return []

    const ids = new Set(active_context.position_IDs.filter(Boolean))
    active_context.actor_IDs.forEach((actor_ID) => {
      const actor = actors.find((a) => a.ID === actor_ID)
      if (actor?.position_ID) ids.add(actor.position_ID)
    })

    return [...ids]
  }, [active_context, actors])

  useLayoutEffect(() => {
    if (!active_context?.source_ID || target_position_IDs.length === 0) {
      setLayout(EMPTY_LAYOUT)
      return
    }

    const overlay = overlayRef.current
    const board = overlay?.parentElement
    if (!overlay || !board) return

    const measure = () => {
      const source = actors.find((a) => a.ID === active_context.source_ID)
      if (!active_context.source_ID) return
      const sourceElement =
        queryActor(board, active_context.source_ID) ??
        (source?.position_ID ? queryPosition(board, source.position_ID) : null)

      if (!sourceElement) {
        setLayout(EMPTY_LAYOUT)
        return
      }

      const boardRect = board.getBoundingClientRect()
      const start = getAnchor(sourceElement.getBoundingClientRect(), boardRect)
      const targets = target_position_IDs
        .flatMap((position_ID) => {
          const targetElement = queryPosition(board, position_ID)
          if (!targetElement) return []

          return [
            {
              id: position_ID,
              end: getAnchor(targetElement.getBoundingClientRect(), boardRect),
            },
          ]
        })
        .sort((a, b) => a.end.x - b.end.x)

      const arrows = targets.flatMap((target, index) => {
        const arrow = getArrow(start, target.end, index, targets.length)
        if (!arrow) return []

        return [{ id: target.id, ...arrow }]
      })

      setLayout({
        width: boardRect.width,
        height: boardRect.height,
        arrows,
      })
    }

    measure()

    const resizeObserver = new ResizeObserver(measure)
    resizeObserver.observe(board)
    positions.forEach((position) => {
      const element = queryPosition(board, position.ID)
      if (element) resizeObserver.observe(element)
    })

    const frame = requestAnimationFrame(measure)
    const timeout = window.setTimeout(measure, 280)

    return () => {
      resizeObserver.disconnect()
      cancelAnimationFrame(frame)
      window.clearTimeout(timeout)
    }
  }, [active_context, actors, positions, target_position_IDs])

  return (
    <div
      ref={overlayRef}
      className="pointer-events-none absolute inset-0 z-30 opacity-50"
    >
      {layout.arrows.length > 0 && (
        <svg
          aria-hidden="true"
          className="absolute inset-0 bottom-1/2 size-full overflow-visible"
          viewBox={`0 0 ${layout.width} ${layout.height}`}
        >
          {layout.arrows.map((arrow) => (
            <g key={arrow.id}>
              <path
                d={arrow.path}
                fill="none"
                opacity="0.55"
                stroke="rgb(0 0 0)"
                strokeLinecap="round"
                strokeWidth="24"
              />
              <path
                d={arrow.path}
                fill="none"
                opacity="1"
                stroke="oklch(95.4% 0.038 75.164)"
                strokeLinecap="round"
                strokeWidth="20"
              />
              <polygon
                fill="oklch(95.4% 0.038 75.164)"
                opacity="1"
                points={arrow.head}
              />
            </g>
          ))}
        </svg>
      )}
    </div>
  )
}

export { TargetingArrows }
