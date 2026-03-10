import { defineStore } from 'pinia'
import { ref } from 'vue'

/** 预设配色 */
export interface ColorPreset {
  id: string
  name: string
  primary: string
  primaryRgb: string
}

/** 白天/深夜模式的全局色变量 */
export interface ModeColors {
  bgLayer0: string
  bgLayer1: string
  bgLayer2: string
  bgLayer3: string
  bgLayer4: string
  textPrimary: string
  textSecondary: string
  textTertiary: string
}

/** 全局色预设（含白天、深夜两套） */
export interface GlobalColorPreset {
  id: string
  name: string
  light: ModeColors
  dark: ModeColors
}

/** 预制主色选项 */
export const COLOR_PRESETS: ColorPreset[] = [
  { id: 'default', name: '默认蓝', primary: '#409eff', primaryRgb: '64, 158, 255' },
  { id: 'green', name: '翠绿', primary: '#67c23a', primaryRgb: '103, 194, 58' },
  { id: 'teal', name: '青碧', primary: '#20b2aa', primaryRgb: '32, 178, 170' },
  { id: 'purple', name: '罗兰紫', primary: '#8b5cf6', primaryRgb: '139, 92, 246' },
  { id: 'pink', name: '桃粉', primary: '#ec4899', primaryRgb: '236, 72, 153' },
  { id: 'orange', name: '琥珀橙', primary: '#f97316', primaryRgb: '249, 115, 22' },
  { id: 'red', name: '朱砂红', primary: '#ef4444', primaryRgb: '239, 68, 68' },
  { id: 'indigo', name: '靛蓝', primary: '#6366f1', primaryRgb: '99, 102, 241' },
  { id: 'cyan', name: '天青', primary: '#06b6d4', primaryRgb: '6, 182, 212' },
]

/** 预制全局色选项（白天 + 深夜） */
export const GLOBAL_COLOR_PRESETS: GlobalColorPreset[] = [
  {
    id: 'default',
    name: '默认白',
    light: {
      bgLayer0: '#ffffff',
      bgLayer1: '#fafafa',
      bgLayer2: '#f5f5f5',
      bgLayer3: '#f0f0f0',
      bgLayer4: '#ebebeb',
      textPrimary: '#303133',
      textSecondary: '#606266',
      textTertiary: '#909399',
    },
    dark: {
      bgLayer0: '#1d1e1f',
      bgLayer1: '#252526',
      bgLayer2: '#2d2d30',
      bgLayer3: '#333333',
      bgLayer4: '#3c3c3c',
      textPrimary: '#e5e7eb',
      textSecondary: '#9ca3af',
      textTertiary: '#6b7280',
    },
  },
  {
    id: 'eye-care',
    name: '护眼绿',
    light: {
      bgLayer0: '#f5f9f0',
      bgLayer1: '#eef5e4',
      bgLayer2: '#e5efdb',
      bgLayer3: '#dce9d2',
      bgLayer4: '#d3e3c9',
      textPrimary: '#2d3a24',
      textSecondary: '#5a6b4a',
      textTertiary: '#849678',
    },
    dark: {
      bgLayer0: '#1a2116',
      bgLayer1: '#232d1c',
      bgLayer2: '#2c3822',
      bgLayer3: '#354328',
      bgLayer4: '#3e4e2e',
      textPrimary: '#d4e5c8',
      textSecondary: '#9cb88a',
      textTertiary: '#6b8a5a',
    },
  },
  {
    id: 'warm',
    name: '暖杏',
    light: {
      bgLayer0: '#fdfbf7',
      bgLayer1: '#f9f5ed',
      bgLayer2: '#f5efe3',
      bgLayer3: '#f0e9d9',
      bgLayer4: '#ebe3cf',
      textPrimary: '#3d3629',
      textSecondary: '#6b6152',
      textTertiary: '#9a8f7b',
    },
    dark: {
      bgLayer0: '#241f1a',
      bgLayer1: '#2e2822',
      bgLayer2: '#38312a',
      bgLayer3: '#423a32',
      bgLayer4: '#4c433a',
      textPrimary: '#e8dfd0',
      textSecondary: '#b5a896',
      textTertiary: '#82776a',
    },
  },
  {
    id: 'cool',
    name: '冷灰',
    light: {
      bgLayer0: '#f8fafc',
      bgLayer1: '#f1f5f9',
      bgLayer2: '#e2e8f0',
      bgLayer3: '#cbd5e1',
      bgLayer4: '#94a3b8',
      textPrimary: '#1e293b',
      textSecondary: '#475569',
      textTertiary: '#64748b',
    },
    dark: {
      bgLayer0: '#0f172a',
      bgLayer1: '#1e293b',
      bgLayer2: '#334155',
      bgLayer3: '#475569',
      bgLayer4: '#64748b',
      textPrimary: '#f1f5f9',
      textSecondary: '#94a3b8',
      textTertiary: '#64748b',
    },
  },
  {
    id: 'midnight',
    name: '午夜蓝',
    light: {
      bgLayer0: '#f7f9fc',
      bgLayer1: '#eef2f8',
      bgLayer2: '#e4eaf4',
      bgLayer3: '#dae2ef',
      bgLayer4: '#d0daea',
      textPrimary: '#1e2d3d',
      textSecondary: '#4a5d72',
      textTertiary: '#768da6',
    },
    dark: {
      bgLayer0: '#0d1117',
      bgLayer1: '#161b22',
      bgLayer2: '#21262d',
      bgLayer3: '#30363d',
      bgLayer4: '#484f58',
      textPrimary: '#e6edf3',
      textSecondary: '#8b949e',
      textTertiary: '#6e7681',
    },
  },
  {
    id: 'paper',
    name: '素纸',
    light: {
      bgLayer0: '#fefefe',
      bgLayer1: '#f7f6f3',
      bgLayer2: '#efebe6',
      bgLayer3: '#e7e1d9',
      bgLayer4: '#dfd7cc',
      textPrimary: '#2c2825',
      textSecondary: '#5c5650',
      textTertiary: '#8c847b',
    },
    dark: {
      bgLayer0: '#1c1b18',
      bgLayer1: '#262420',
      bgLayer2: '#302d28',
      bgLayer3: '#3a3630',
      bgLayer4: '#443f38',
      textPrimary: '#ebe6de',
      textSecondary: '#afa89e',
      textTertiary: '#7d766e',
    },
  },
]

export const useSettingsStore = defineStore(
  'settings',
  () => {
    const colorPresetId = ref('default')
    const globalColorPresetId = ref('default')
    const darkMode = ref(false)

    function setColorPreset(id: string) {
      const preset = COLOR_PRESETS.find((p) => p.id === id)
      if (preset) {
        colorPresetId.value = id
        applyColorPreset(preset)
      }
    }

    function setDarkMode(v: boolean) {
      darkMode.value = v
      if (v) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }
      applyGlobalColorPreset()
    }

    function setGlobalColorPreset(id: string) {
      const preset = GLOBAL_COLOR_PRESETS.find((p) => p.id === id)
      if (preset) {
        globalColorPresetId.value = id
        applyGlobalColorPreset()
      }
    }

    function applyGlobalColorPreset() {
      const preset = GLOBAL_COLOR_PRESETS.find((p) => p.id === globalColorPresetId.value)
      if (!preset) return
      const colors = darkMode.value ? preset.dark : preset.light
      const root = document.documentElement
      root.style.setProperty('--bg-layer-0', colors.bgLayer0)
      root.style.setProperty('--bg-layer-1', colors.bgLayer1)
      root.style.setProperty('--bg-layer-2', colors.bgLayer2)
      root.style.setProperty('--bg-layer-3', colors.bgLayer3)
      root.style.setProperty('--bg-layer-4', colors.bgLayer4)
      root.style.setProperty('--text-primary', colors.textPrimary)
      root.style.setProperty('--text-secondary', colors.textSecondary)
      root.style.setProperty('--text-tertiary', colors.textTertiary)
      // 同步 Element Plus 变量，使 table/input 等兼容主题
      root.style.setProperty('--el-bg-color', colors.bgLayer1)
      root.style.setProperty('--el-bg-color-overlay', colors.bgLayer0)
      root.style.setProperty('--el-fill-color-blank', colors.bgLayer0)
      root.style.setProperty('--el-fill-color-light', colors.bgLayer2)
      root.style.setProperty('--el-fill-color-lighter', colors.bgLayer3)
      root.style.setProperty('--el-border-color', colors.bgLayer3)
      root.style.setProperty('--el-border-color-lighter', colors.bgLayer4)
      root.style.setProperty('--el-text-color-primary', colors.textPrimary)
      root.style.setProperty('--el-text-color-regular', colors.textPrimary)
      root.style.setProperty('--el-text-color-secondary', colors.textSecondary)
      root.style.setProperty('--el-text-color-placeholder', colors.textTertiary)
    }

    function applyColorPreset(preset: ColorPreset) {
      document.documentElement.style.setProperty('--el-color-primary', preset.primary)
      document.documentElement.style.setProperty('--el-color-primary-rgb', preset.primaryRgb)
    }

    function initFromStorage() {
      // 迁移旧版主题配置 (x-admin-theme)
      try {
        const legacy = localStorage.getItem('x-admin-theme')
        if (legacy === 'true' && !darkMode.value) {
          darkMode.value = true
        }
        if (legacy) localStorage.removeItem('x-admin-theme')
      } catch {}

      const preset = COLOR_PRESETS.find((p) => p.id === colorPresetId.value)
      if (preset) applyColorPreset(preset)
      setDarkMode(darkMode.value)
      applyGlobalColorPreset()
    }

    return {
      colorPresetId,
      globalColorPresetId,
      darkMode,
      setColorPreset,
      setGlobalColorPreset,
      setDarkMode,
      initFromStorage,
    }
  },
  {
    persist: {
      key: 'x-admin-settings',
      storage: localStorage,
    },
  }
)
