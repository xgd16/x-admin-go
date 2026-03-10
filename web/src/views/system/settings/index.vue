<script setup lang="ts">
import { useSettingsStore } from '@/store/settings'
import { COLOR_PRESETS, GLOBAL_COLOR_PRESETS } from '@/store/settings'

const settings = useSettingsStore()
</script>

<template>
  <div class="settings-page">
    <h2 class="page-title mb-3 text-gray-800 dark:text-gray-100">后台设置</h2>

    <el-card class="page-card rounded-card" shadow="always">
      <template #header>
        <span>主题与配色</span>
      </template>

      <div class="space-y-6">
        <div>
          <div class="text-sm font-medium text-gray-700 dark:text-gray-300 mb-3">主题模式</div>
          <div class="flex items-center gap-2">
            <el-switch
              :model-value="settings.darkMode"
              @update:model-value="settings.setDarkMode"
              inline-prompt
              active-text="暗色"
              inactive-text="亮色"
            />
          </div>
        </div>

        <div>
          <div class="text-sm font-medium text-gray-700 dark:text-gray-300 mb-3">主色预设</div>
          <div class="flex flex-wrap gap-3">
            <div
              v-for="preset in COLOR_PRESETS"
              :key="preset.id"
              class="preset-item"
              :class="{ active: settings.colorPresetId === preset.id }"
              @click="settings.setColorPreset(preset.id)"
            >
              <div class="preset-color" :style="{ background: preset.primary }" />
              <span class="preset-name">{{ preset.name }}</span>
            </div>
          </div>
        </div>

        <div>
          <div class="text-sm font-medium text-gray-700 dark:text-gray-300 mb-3">全局色预设</div>
          <p class="text-xs text-gray-500 dark:text-gray-400 mb-3">含白天与深夜两套配色</p>
          <div class="flex flex-wrap gap-3">
            <div
              v-for="preset in GLOBAL_COLOR_PRESETS"
              :key="preset.id"
              class="global-preset-item"
              :class="{ active: settings.globalColorPresetId === preset.id }"
              @click="settings.setGlobalColorPreset(preset.id)"
            >
              <div class="global-preset-preview">
                <div class="preview-light" :style="{ background: preset.light.bgLayer0, color: preset.light.textPrimary }">亮</div>
                <div class="preview-dark" :style="{ background: preset.dark.bgLayer0, color: preset.dark.textPrimary }">暗</div>
              </div>
              <span class="preset-name">{{ preset.name }}</span>
            </div>
          </div>
        </div>
      </div>
    </el-card>
  </div>
</template>

<style scoped>
.preset-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 14px;
  border-radius: 8px;
  border: 2px solid transparent;
  cursor: pointer;
  transition: all 0.2s;
}

.preset-item:hover {
  background: var(--el-fill-color-light);
}

.preset-item.active {
  border-color: var(--el-color-primary);
  background: var(--el-color-primary-light-9);
}

.preset-color {
  width: 20px;
  height: 20px;
  border-radius: 4px;
  flex-shrink: 0;
}

.preset-name {
  font-size: 13px;
  color: var(--text-primary);
}

.global-preset-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  padding: 10px 12px;
  border-radius: 8px;
  border: 2px solid transparent;
  cursor: pointer;
  transition: all 0.2s;
}

.global-preset-item:hover {
  background: var(--el-fill-color-light);
}

.global-preset-item.active {
  border-color: var(--el-color-primary);
  background: var(--el-color-primary-light-9);
}

.global-preset-preview {
  display: flex;
  gap: 2px;
}

.global-preset-preview .preview-light,
.global-preset-preview .preview-dark {
  width: 28px;
  height: 24px;
  border-radius: 4px;
  font-size: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.global-preset-item .preset-name {
  font-size: 12px;
}
</style>
