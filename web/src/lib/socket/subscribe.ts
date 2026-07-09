import type { SocketMessageSubscriber } from './request'

const messageSubscribers = new Set<SocketMessageSubscriber>()

function subscribe(subscriber: SocketMessageSubscriber) {
  messageSubscribers.add(subscriber)
  return () => {
    messageSubscribers.delete(subscriber)
  }
}

export { subscribe, messageSubscribers }
