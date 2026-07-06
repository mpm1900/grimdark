import type { Actor } from '#/lib/game/actor'

function ActorAvatar({ actor }: { actor: Actor }) {
  return (
    <div className="relative h-48 w-52">
      <div className="absolute -bottom-10 left-2 bg-[url(/gothic/AvatarCircleFrame.png)] bg-contain bg-center size-60"></div>
    </div>
  )
}

export { ActorAvatar }
