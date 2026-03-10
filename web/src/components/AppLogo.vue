<script setup lang="ts">
defineProps<{
  collapsed?: boolean
}>()
</script>

<template>
  <div class="app-logo flex items-center justify-center h-14 px-4">
    <span v-if="!collapsed" class="text-lg font-semibold text-gray-800 dark:text-gray-100">
      x-admin
    </span>
    <span v-else class="text-lg font-bold text-primary">x</span>
  </div>
</template>
