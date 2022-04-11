
export function get(url: string, options?: any): Promise<Record<string, any>> {
  if (!options) {
    options = {
      headers: {
        'Content-Type': 'application/json'
      }
    }
  }
  return new Promise<Record<string, any>>(resolve => {
    fetch(url, options).then(res => {
      if (res.status === 204) {
        return {}
      }
      return res.json()
    }).then(res => resolve(res)).catch(e => {
      console.error(e)
    })
  })
}
