import { Store } from '@tanstack/store'
import type {
  SocketMessageSubscriber,
  SocketRequest,
  SocketResponse,
} from '../socket/request'

type SocketStatus =
  | 'idle'
  | 'connecting'
  | 'open'
  | 'closing'
  | 'closed'
  | 'error'
  | 'reconnecting'

type SocketState = {
  instance_ID: string | null
  socket: WebSocket | null
  status: SocketStatus
  reconnect_count: number
}

const socketStore = new Store<SocketState>({
  instance_ID: null,
  socket: null,
  status: 'idle',
  reconnect_count: 0,
})

function sendSocketMessage(
  payload: string | ArrayBufferLike | Blob | ArrayBufferView
): boolean {
  const socket = socketStore.state.socket
  if (!socket || socket.readyState !== WebSocket.OPEN) {
    console.warn('Attempted to send message while socket is not open')
    return false
  }

  socket.send(payload as any)
  return true
}

function sendContextMessage(request: SocketRequest) {
  return sendSocketMessage(JSON.stringify(request))
}

function setSocket(socket: WebSocket) {
  socketStore.setState((s) => ({
    ...s,
    socket,
    status: s.reconnect_count > 0 ? 'reconnecting' : 'connecting',
  }))
}

function openSocket() {
  console.log('WebSocket connection opened')
  socketStore.setState((s) => ({
    ...s,
    status: 'open',
  }))
}

function closeSocketEmpty() {
  socketStore.setState((s) => ({
    ...s,
    status: 'closed',
    instance_ID: null,
    reconnectCount: 0,
  }))
}

function closeSocket(code: number, reason: string) {
  socketStore.setState((s) => ({
    ...s,
    status: 'closing',
  }))

  if (
    socketStore.state.socket?.readyState === WebSocket.CONNECTING ||
    socketStore.state.socket?.readyState === WebSocket.OPEN
  ) {
    socketStore.state.socket.close(code, reason)
    return
  }

  socketStore.setState((s) => ({
    ...s,
    socket: null,
    status: 'closed',
    instance_ID: null,
    reconnectCount: 0,
  }))
}

function deleteSocket() {
  socketStore.setState((s) => ({
    ...s,
    socket: null,
    status: s.status ? 'error' : 'closed',
  }))
}

function startReconnect() {
  socketStore.setState((s) => ({
    ...s,
    reconnectCount: s.reconnect_count + 1,
    status: 'reconnecting',
  }))
}

function resetReconnect(socket: WebSocket, signal: AbortSignal) {
  setTimeout(() => {
    socketStore.setState((s) => {
      if (
        !signal.aborted &&
        socketStore.state.socket === socket &&
        socketStore.state.status === 'open'
      ) {
        return {
          ...s,
          reconnect_count: 0,
        }
      }

      return s
    })
  }, 5000)
}

function setSocketError() {
  socketStore.setState((s) => ({
    ...s,
    status: 'error',
  }))
}

export {
  closeSocketEmpty,
  closeSocket,
  deleteSocket,
  openSocket,
  resetReconnect,
  sendContextMessage,
  sendSocketMessage,
  socketStore,
  startReconnect,
  setSocket,
  setSocketError,
}
export type { SocketMessageSubscriber, SocketResponse }
