function getSocketUrl(instanceID?: string): string {
  const envUrl = import.meta.env.VITE_BACKEND_URL
  const end = instanceID ? `${instanceID}/connect` : 'connect'
  if (envUrl) {
    const protocol = envUrl.startsWith('https') ? 'wss' : 'ws'
    const host = envUrl.replace(/^https?:\/\//, '')
    return `${protocol}://${host}/socket/${end}`
  }

  // Same host as the site (e.g. Traefik routes /socket on 443). Avoid hard-coded :3005 on HTTPS.
  if (typeof window !== 'undefined' && window.location.protocol === 'https:') {
    return `wss://${window.location.host}/socket/${end}`
  }

  const protocol = window.location.protocol === 'https:' ? 'wss' : 'ws'
  const hostname = window.location.hostname
  const port = '3005'

  return `${protocol}://${hostname}:${port}/socket/${end}`
}

export { getSocketUrl }
