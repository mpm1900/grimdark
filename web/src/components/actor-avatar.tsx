import type { Actor } from '#/lib/game/actor'

function ActorAvatar({ actor }: { actor: Actor }) {
  return (
    <div className="relative h-48 w-52">
      <div className="absolute -bottom-13 left-2 bg-[url(/gothic/AvatarCircleFrame.png)] bg-contain bg-center size-60">
        <div className="bg-neutral-950 border-neutral-800 rounded-xs border-2 ring ring-black absolute top-0 inset-x-6 h-9 font-cinzel-dec text-foreground/80 font-semibold leading-9 px-2">
          {actor.name}
        </div>
      </div>
    </div>
  )
}

export { ActorAvatar }
