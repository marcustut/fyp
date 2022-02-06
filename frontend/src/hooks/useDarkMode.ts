import { useEffect, useMemo } from 'react'

import { useLocalStorage, useMedia } from '@/hooks'

// This hook handles all the stateful logic required to add a ☾ dark mode toggle to your website
export const useDarkMode = () => {
  // Use our useLocalStorage hook to persist state through a page refresh.
  // Read the recipe for this hook to learn more: usehooks.com/useLocalStorage
  const [enabledState, setEnabledState] = useLocalStorage<boolean>(
    'dark-mode-enabled',
    false
  )
  // See if user has set a browser or OS preference for dark mode.
  // The usePrefersDarkMode hook composes a useMedia hook (see code below).
  const prefersDarkMode = usePrefersDarkMode()
  // If enabledState is defined use it, otherwise fallback to prefersDarkMode.
  // This allows user to override OS level setting on our website.
  const enabled = useMemo(
    () => enabledState ?? prefersDarkMode,
    [enabledState, prefersDarkMode]
  )
  // Fire off effect that add/removes dark mode class
  useEffect(
    () => {
      const className = 'dark'
      const element = window.document.body
      if (enabled) {
        element.classList.add(className)
      } else {
        element.classList.remove(className)
      }
    },
    [enabled] // Only re-call effect when value changes
  )
  // Return enabled state and setter
  return { darkMode: enabled, setDarkMode: setEnabledState }
}

// Compose our useMedia hook to detect dark mode preference.
// The API for useMedia looks a bit weird, but that's because ...
// ... it was designed to support multiple media queries and return values.
// Thanks to hook composition we can hide away that extra complexity!
// Read the recipe for useMedia to learn more: usehooks.com/useMedia
const usePrefersDarkMode = () =>
  useMedia<boolean>(['(prefers-color-scheme: dark)'], [true], false)
