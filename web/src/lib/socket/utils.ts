import { socketStore } from '../stores/socket'
import { INSTANCE_ID_KEY } from './config'

function clearSocketEventHandlers(socket: WebSocket) {
  socket.onopen = null
  socket.onclose = null
  socket.onerror = null
  socket.onmessage = null
}

function isCurrentSocket(socket: WebSocket): boolean {
  return socketStore.get().socket === socket
}

function readSavedInstanceID(): string | null {
  if (typeof window === 'undefined') {
    return null
  }
  return localStorage.getItem(INSTANCE_ID_KEY)
}

export { clearSocketEventHandlers, isCurrentSocket, readSavedInstanceID }
