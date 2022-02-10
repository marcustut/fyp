const STORAGE_PREFIX = 'slides'

export const getItem = <T = unknown>(key: string) => {
  const item = window.localStorage.getItem(`${STORAGE_PREFIX}-${key}`)
  if (!item)
    return {
      data: null,
      error: new Error(
        `unable to get item with key \`${
          STORAGE_PREFIX + '-' + key
        }\` from localStorage`
      ),
    }
  return { data: JSON.parse(item) as T, error: null }
}

export const setItem = <T = unknown>(key: string, item: T) => {
  window.localStorage.setItem(`${STORAGE_PREFIX}-${key}`, JSON.stringify(item))
}

export const clearItem = (key: string) => {
  window.localStorage.removeItem(`${STORAGE_PREFIX}-${key}`)
}
