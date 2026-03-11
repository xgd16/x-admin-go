<script setup lang="ts">
import type { Placement } from 'element-plus'

withDefaults(
  defineProps<{
    content: string
    placement?: Placement
    type?: 'primary' | 'success' | 'warning' | 'danger' | 'info' | 'default'
    size?: 'large' | 'default' | 'small'
    circle?: boolean
  }>(),
  { placement: 'bottom', type: 'default', size: 'default', circle: true }
)

defineEmits<{
  click: []
}>()
</script>

<template>
  <el-tooltip :content="content" :placement="placement">
    <el-button
      v-motion
      :type="type"
      :size="size"
      :circle="circle"
      text
      :tapped="{ scale: 0.92, transition: { duration: 0.15 } }"
      @click="$emit('click')"
    >
      <slot />
    </el-button>
  </el-tooltip>
</template>
