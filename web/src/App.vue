<script setup lang="ts">
import { onMounted } from 'vue'
import { useSettingsStore } from '@/store/settings'

onMounted(() => {
  useSettingsStore().initFromStorage()
})
</script>

<template>
  <router-view />
</template>
