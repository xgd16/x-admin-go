<script setup lang="ts">
import { computed, onMounted, ref, watch, defineExpose, defineOptions, defineComponent, h, useSlots } from 'vue'
import type { VNode } from 'vue'
defineOptions({ inheritAttrs: false })
import { ElTable, ElPagination, ElMessage } from 'element-plus'
import { ElTableColumn } from 'element-plus'
import type { TableColumnCtx } from 'element-plus/es/components/table/src/table-column/defaults'
import type { DefaultRow } from 'element-plus/es/components/table/src/table/defaults'
import req from '@/utils/request'
import type { Response } from '@/utils/request'

type HttpMethod = 'get' | 'post' | 'put' | 'delete' | 'patch'

export type TableColumn<T extends DefaultRow = DefaultRow> = {
  type?: 'selection' | 'index' | 'expand'
  label?: string
  prop?: string
  width?: number | string
  minWidth?: number | string
  fixed?: 'left' | 'right' | boolean
  align?: 'left' | 'center' | 'right'
  headerAlign?: 'left' | 'center' | 'right'
  sortable?: boolean | 'custom'
  showOverflowTooltip?: boolean
  formatter?: (row: T, column: TableColumnCtx<T>, cellValue: any, index: number) => any
  slot?: string
  headerSlot?: string
  children?: TableColumn<T>[]
  [key: string]: any
}

const props = withDefaults(defineProps<{
  url: string
  method?: HttpMethod | string
  search?: any[]
  pageSize?: number
  currentPage?: number
  columns?: TableColumn[]
  // 是否组件挂载后立即请求
  immediate?: boolean
  // 是否显示分页
  pagination?: boolean
  // 后端分页参数键
  pageNumKey?: string
  pageSizeKey?: string
  // 响应数据结构可配置
  listKey?: string
  totalKey?: string
  // 请求发送前处理 payload
  beforeRequest?: (payload: Record<string, any>) => Record<string, any>
  // 响应完成后钩子
  afterResponse?: (res: Response<any>) => void
  // 额外请求配置（如 headers）
  extraConfig?: Record<string, any>
}>(), {
  method: 'post',
  search: () => [],
  pageSize: 10,
  currentPage: 1,
  columns: () => [],
  immediate: true,
  pagination: true,
  pageNumKey: 'pageNum',
  pageSizeKey: 'pageSize',
  listKey: 'list',
  totalKey: 'total',
})

const emits = defineEmits<{
  (e: 'update:currentPage', value: number): void
  (e: 'update:pageSize', value: number): void
  (e: 'loaded', payload: { list: any[]; total: number }): void
  (e: 'error', error: unknown): void
}>()

const loading = ref(false)
const tableData = ref<any[]>([])
const total = ref(0)
const currentPage = ref(props.currentPage ?? 1)
const pageSize = ref(props.pageSize)
const tableRef = ref<InstanceType<typeof ElTable>>()

watch(() => props.pageSize, (val) => {
  if (typeof val === 'number' && val > 0) {
    pageSize.value = val
  }
})

watch(() => props.currentPage, (val) => {
  if (typeof val === 'number' && val > 0) {
    currentPage.value = val
  }
})

// 将 any[] 的搜索条件合并为一个对象，非对象项忽略
const normalizedSearch = computed<Record<string, any>>(() => {
  const result: Record<string, any> = {}
  const items = Array.isArray(props.search) ? props.search : []
  for (const item of items) {
    if (item && typeof item === 'object' && !Array.isArray(item)) {
      Object.assign(result, item as Record<string, any>)
    }
  }
  return result
})

const fetchList = async (): Promise<{ list: any[]; total: number } | null> => {
  if (!props.url) return null
  loading.value = true
  try {
    const method = (props.method || 'post').toLowerCase() as HttpMethod
    const payloadBase: Record<string, any> = {
      [props.pageNumKey]: currentPage.value,
      [props.pageSizeKey]: pageSize.value,
      ...normalizedSearch.value,
    }
    const payload = props.beforeRequest ? props.beforeRequest(payloadBase) : payloadBase

    // 兼容 get 与非 get 的参数位置
    const config = method === 'get'
      ? { url: props.url, method, params: payload, ...props.extraConfig }
      : { url: props.url, method, data: payload, ...props.extraConfig }

    const res = await req<Response<any>>(config)
    if (props.afterResponse) props.afterResponse(res)
    if (res.code !== 0) {
      ElMessage.error(res.msg ?? '请求失败')
      emits('error', res)
      return null
    }

    const data = (res.data ?? {}) as unknown as Record<string, unknown>
    const list = data[props.listKey] ?? []
    const totalCount = Number(data[props.totalKey] ?? 0)

    tableData.value = Array.isArray(list) ? list : []
    total.value = Number.isFinite(totalCount) ? totalCount : 0
    const normalized = { list: tableData.value, total: total.value }
    emits('loaded', normalized)
    return normalized
  } catch (err) {
    ElMessage.error('网络错误或服务异常')
    emits('error', err)
    return null
  } finally {
    loading.value = false
  }
}

const handleSizeChange = (val: number) => {
  pageSize.value = val
  emits('update:pageSize', val)
  currentPage.value = 1
  emits('update:currentPage', 1)
  fetchList()
}

const handleCurrentChange = (val: number) => {
  currentPage.value = val
  emits('update:currentPage', val)
  fetchList()
}

// 暴露手动刷新方法
const refresh = () => {
  currentPage.value = 1
  emits('update:currentPage', 1)
  fetchList()
}
const refreshKeepPage = () => {
  fetchList()
}

// 选择行相关暴露
const getSelection = () => tableRef.value?.getSelectionRows?.() ?? []
const clearSelection = () => tableRef.value?.clearSelection?.()
const toggleRowSelection = (row: any, selected?: boolean) => tableRef.value?.toggleRowSelection?.(row, selected)

defineExpose({ refresh, refreshKeepPage, fetchList, getSelection, clearSelection, toggleRowSelection, tableRef })

let searchDebounceTimer: any = null
watch(normalizedSearch, () => {
  // 搜索项变化时重置到第一页并刷新（防抖）
  if (searchDebounceTimer) clearTimeout(searchDebounceTimer)
  searchDebounceTimer = setTimeout(() => {
    currentPage.value = 1
    emits('update:currentPage', 1)
    fetchList()
  }, 300)
})

onMounted(() => {
  if (props.immediate) fetchList()
})

// 递归渲染列 - 本地组件
const rootSlots = useSlots()
const ColumnsRender = defineComponent<{ columns: TableColumn[] }>({
  name: 'ColumnsRender',
  props: {
    columns: { type: Array as unknown as () => TableColumn[], default: () => [] },
  },
  setup(localProps) {
    function renderColumns(cols: TableColumn[]): VNode[] {
      return cols.map((col): VNode => {
        const { children, slot, headerSlot, ...colProps } = col
        return h(
          ElTableColumn as any,
          colProps,
          {
            default: children && children.length
              ? () => renderColumns(children)
              : (scope: any) => {
                  const slotFn = slot ? rootSlots[slot] : undefined
                  return slotFn ? slotFn(scope) : undefined
                },
            header: (scope: any) => {
              const headerFn = headerSlot ? rootSlots[headerSlot] : undefined
              return headerFn ? headerFn(scope) : undefined
            },
          },
        )
      })
    }
    return () => renderColumns(localProps.columns)
  },
})
</script>

<template>
  <div class="table-wrapper">
    <!-- 工具栏插槽（可选） -->
    <div class="table-toolbar">
      <slot name="toolbar" />
    </div>

    <ElTable :data="tableData" v-loading="loading" border style="width: 100%" v-bind="$attrs" ref="tableRef">
      <!-- 列：优先使用 columns 配置；未提供则回退到插槽方式 -->
      <ColumnsRender v-if="props.columns && props.columns.length" :columns="props.columns" />
      <slot v-else name="columns" />

      <!-- 空状态插槽（可覆盖） -->
      <template #empty>
        <slot name="empty">
          <div class="empty-wrapper">
            <el-empty description="暂无数据" />
          </div>
        </slot>
      </template>
    </ElTable>

    <div class="table-pagination" v-if="props.pagination">
      <ElPagination background layout="total, sizes, prev, pager, next, jumper" :total="total"
        :page-sizes="[5,10, 20, 30, 50, 100]" :page-size="pageSize" :current-page="currentPage"
        @size-change="handleSizeChange" @current-change="handleCurrentChange" />
    </div>
  </div>
</template>

<style scoped>
.table-wrapper {
  display: flex;
  flex-direction: column;
  gap: 8px;
}
.table-toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.table-pagination {
  display: flex;
  justify-content: flex-end;
}
</style>



