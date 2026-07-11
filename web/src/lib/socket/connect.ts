import { getSocketUrl } from '#/utils/get-socket-url'
import {
  closeSocket,
  closeSocketEmpty,
  deleteSocket,
  openSocket,
  resetReconnect,
  setSocket,
  setSocketError,
  socketStore,
  startReconnect,
  type SocketResponse,
} from '../stores/socket'
import {
  INITIAL_RECONNECT_DELAY,
  INSTANCE_ID_KEY,
  MAX_RECONNECT_DELAY,
} from './config'
import { socket_reducer } from './reducer'
import { messageSubscribers } from './subscribe'
import {
  clearSocketEventHandlers,
  isCurrentSocket,
  readSavedInstanceID,
} from './utils'

let connectionAbortController: AbortController | null = null
let reconnectTimer: number | null = null

function connect(instanceID?: string | null): Promise<SocketResponse> {
  return new Promise<SocketResponse>((resolve, reject) => {
    let settled = false
    const finishResolve = (value: SocketResponse) => {
      if (settled) return
      settled = true
      resolve(value)
    }
    const finishReject = (reason: unknown) => {
      if (settled) return
      settled = true
      reject(reason instanceof Error ? reason : new Error(String(reason)))
    }

    if (connectionAbortController) {
      connectionAbortController.abort()
    }

    connectionAbortController = new AbortController()
    const signal = connectionAbortController.signal
    signal.addEventListener(
      'abort',
      () => {
        finishReject(new Error('Socket connect aborted'))
      },
      { once: true }
    )

    if (reconnectTimer) {
      clearTimeout(reconnectTimer)
      reconnectTimer = null
    }

    const previous = socketStore.state.socket
    if (previous) {
      clearSocketEventHandlers(previous)
      if (
        previous.readyState === WebSocket.CONNECTING ||
        previous.readyState === WebSocket.OPEN
      ) {
        previous.close(1000, 'Switching instance')
      }
    }

    const url = getSocketUrl(instanceID)
    const socket = new WebSocket(url)
    setSocket(socket)

    socket.onopen = () => {
      if (signal.aborted || !isCurrentSocket(socket)) {
        socket.close(1000, 'Aborted')
        return
      }

      openSocket()
      resetReconnect(socket, signal)
    }

    socket.onmessage = (event) => {
      if (signal.aborted || !isCurrentSocket(socket)) return
      let message: SocketResponse | null = null

      if (typeof event.data === 'string') {
        try {
          message = JSON.parse(event.data) as SocketResponse
        } catch {
          console.error('non-JSON payload')
        }
      }

      if (message?.type === 'on-connect') {
        finishResolve(message)
      }
      socket_reducer(message)
      for (const subscriber of messageSubscribers) {
        try {
          subscriber(event, message)
        } catch (error) {
          console.error('socket message subscriber error', error)
        }
      }
    }

    socket.onerror = (error) => {
      finishReject(new Error('WebSocket error during connect'))
      if (signal.aborted || !isCurrentSocket(socket)) return
      console.error('WebSocket error:', error)
      setSocketError()
    }

    socket.onclose = (event) => {
      console.log(
        `WebSocket connection closed: code=${event.code}, reason=${event.reason}, wasClean=${event.wasClean}`
      )
      if (signal.aborted || !isCurrentSocket(socket)) return
      const should_reconnect = socketStore.state.status !== 'closing'
      if (should_reconnect && !settled) {
        finishReject(
          new Error(
            `WebSocket closed before on-connect (code=${event.code}, reason=${event.reason || 'none'})`
          )
        )
      }

      deleteSocket()
      if (should_reconnect) {
        reconnect()
      }
    }
  })
}

function reconnect() {
  const { reconnect_count } = socketStore.state

  const delay = Math.min(
    INITIAL_RECONNECT_DELAY * Math.pow(2, reconnect_count),
    MAX_RECONNECT_DELAY
  )

  console.log(
    `Attempting to reconnect in ${delay}ms... (attempt ${reconnect_count + 1})`
  )

  startReconnect()
  reconnectTimer = window.setTimeout(() => {
    reconnectTimer = null
    const instance_ID = readSavedInstanceID()
    void connect(instance_ID).catch(() => {
      // Reconnect loop is tracked in store state; ignore per-attempt promise rejection.
    })
  }, delay)
}

function disconnect(code = 1000, reason = 'Manual disconnect') {
  if (reconnectTimer) {
    clearTimeout(reconnectTimer)
    reconnectTimer = null
  }

  if (!socketStore.state.socket) {
    closeSocketEmpty()
    if (typeof window !== 'undefined') {
      localStorage.removeItem(INSTANCE_ID_KEY)
    }
    return
  }

  closeSocket(code, reason)
  localStorage.removeItem(INSTANCE_ID_KEY)
}

export { connect, disconnect }
