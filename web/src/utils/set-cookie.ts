import { setResponseHeader } from '@tanstack/react-start/server'
import type { AxiosResponse } from 'axios'

function setResponseCookie(response: Pick<AxiosResponse, 'headers'>) {
  const getSetCookie =
    (
      response.headers as {
        getSetCookie?: () => string[]
      }
    ).getSetCookie?.() ?? []
  const headerCookie = response.headers['set-cookie']
  const fallbackCookie = Array.isArray(headerCookie)
    ? headerCookie[0]
    : headerCookie
  const setCookie = getSetCookie[0] ?? fallbackCookie

  if (setCookie && setCookie.length > 0) {
    setResponseHeader('set-cookie', setCookie)
  }
}

export { setResponseCookie }
