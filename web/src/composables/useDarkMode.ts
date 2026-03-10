import { useDark, useToggle } from '@vueuse/core'
import { watch } from 'vue'

const isDark = useDark({ storageKey: 'x-admin-theme' })
const toggleDark = useToggle(isDark)

export function useDarkMode() {
  watch(
    isDark,
    (dark) => {
      if (dark) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }
    },
    { immediate: true }
  )
  return { isDark, toggleDark }
}
