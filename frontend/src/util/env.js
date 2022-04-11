export function isDevMode() {
  return import.meta.env.DEV
}

export function isProdMode() {
  return import.meta.env.PROD
}

export function apiPrefix() {
  const api = import.meta.env.VITE_BASE_API
  if (api) {
    return api + ''
  }
  return ''
}
