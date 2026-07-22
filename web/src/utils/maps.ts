export function keys<T extends string>(map: Record<T, any>) {
  return Object.keys(map) as Array<T>
}
export function entries<K extends string, V>(
  map: Record<K, V> | Partial<Record<K, V>>
) {
  return Object.entries(map) as Array<[K, V]>
}
