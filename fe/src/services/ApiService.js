export const appUrl = () => {
  return window.location.port !== 6970 ? `http://${window.location.hostname}:6970` : ''
}

export const isLocal = () => {
  return window.location.hostname === '127.0.0.1' || window.location.hostname === 'localhost'
}
