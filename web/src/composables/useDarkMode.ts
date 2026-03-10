import { useSettingsStore } from '@/store/settings'

export function useDarkMode() {
  const settings = useSettingsStore()
  return {
    isDark: settings.darkMode,
    toggleDark: () => settings.setDarkMode(!settings.darkMode),
  }
}
