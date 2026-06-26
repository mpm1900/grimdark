export function keys<T extends string>(map: Record<T, any>) {
  return Object.keys(map) as Array<T>
}
