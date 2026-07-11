import { connect } from '#/lib/socket/connect'
import { socketStore } from '#/lib/stores/socket'
import { useSelector } from '@tanstack/react-store'
import { useEffect } from 'react'

function useReconnect(gameID: string) {
  const socket_status = useSelector(socketStore, (s) => s.status)
  useEffect(() => {
    const should_connect =
      socket_status !== 'open' &&
      socket_status !== 'connecting' &&
      socket_status !== 'reconnecting'

    if (should_connect) {
      connect(gameID)
    }
  }, [gameID, socket_status])
}

export { useReconnect }
