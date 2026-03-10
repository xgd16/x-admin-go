## Table 组件使用说明

基于 Element Plus 二次封装的表格组件，保留 `columns` 配置与插槽两种列定义方式，内置分页、请求与搜索联动、选择行操作等能力，且支持多种后端返回结构的适配。

### 基础示例

```vue
<script setup lang="ts">
import { ref } from 'vue'

const apiUrl = '/api/demo/list'
const page = ref(1)
const size = ref(10)
const form = ref({ keyword: '' })

const columns = [
  { type: 'selection', width: 48 },
  { type: 'index', label: '#', width: 60 },
  { label: '姓名', prop: 'name', minWidth: 120 },
  { label: '年龄', prop: 'age', width: 100, align: 'center' },
]

const tableRef = ref()

function refresh() {
  tableRef.value?.refresh()
}
</script>

<template>
  <Table
    ref="tableRef"
    :url="apiUrl"
    :search="[form]"
    :columns="columns"
    v-model:currentPage="page"
    v-model:pageSize="size"
  />
</template>
```

### Props

- **url**: 请求地址，必填，字符串。
- **method**: 请求方法，默认 `post`，可选 `get|post|put|delete|patch`。
- **search**: 搜索条件数组，组件会将数组内所有对象浅合并为请求参数。
- **pageSize**: 每页条数，默认 `10`。
- **currentPage**: 当前页，默认 `1`，支持 `v-model:currentPage`。
- **columns**: 列配置数组，等价于 `ElTableColumn` 的属性，支持 `children` 实现多级表头；也可不传，改用插槽 `columns` 自定义。
- **immediate**: 是否在挂载后立即请求，默认 `true`。
- **pagination**: 是否显示分页，默认 `true`。
- 分页参数键：
  - **pageNumKey**: 默认 `pageNum`
  - **pageSizeKey**: 默认 `pageSize`
- 响应数据键：
  - **listKey**: 默认 `list`
  - **totalKey**: 默认 `total`
- **beforeRequest(payload)**: 发送请求前对最终 `payload` 进行处理，需返回处理后的对象。
- **afterResponse(res)**: 请求返回后的回调（无论成功失败都会触发）。
- **extraConfig**: 额外请求配置（如 headers、timeout 等），将透传到内部请求方法。

### Emits

- **update:currentPage**: 当前页变化时触发。
- **update:pageSize**: 每页条数变化时触发。
- **loaded(payload)**: 数据加载成功后触发，`payload = { list, total }`。
- **error(error)**: 请求失败或后端返回非 0 状态时触发。

### 插槽

- **toolbar**: 表格上方工具栏区域。
- **columns**: 完全自定义列时使用（不传 `columns` 时生效）。
- `ElTable` 原生插槽（如 `empty`）可直接覆盖，例如：

```vue
<Table :url="url">
  <template #empty>暂无数据</template>
</Table>
```

使用 `columns` 配置且需要自定义单元格/表头时，可在列配置中设定 `slot` 或 `headerSlot`，并在父组件中提供对应命名插槽：

```vue
<script setup lang="ts">
const columns = [
  { label: '姓名', prop: 'name', slot: 'nameCell' },
  { label: '年龄', prop: 'age', headerSlot: 'ageHeader' },
]
</script>

<template>
  <Table :url="url" :columns="columns">
    <template #nameCell="{ row }">{{ row.name }}（自定义）</template>
    <template #ageHeader>年龄(岁)</template>
  </Table>
</template>
```

### 搜索与防抖

- 将多个搜索对象放入 `search` 数组即可：`[formA, formB]`，组件会浅合并。
- 搜索条件变化会 300ms 防抖后自动重置到第一页并刷新。

### 自定义请求与响应结构

后端若使用不同的分页参数或响应字段：

```vue
<Table
  :url="url"
  :pageNumKey="'page'"
  :pageSizeKey="'size'"
  :listKey="'rows'"
  :totalKey="'totalCount'"
  :beforeRequest="(p) => ({ ...p, tenantId })"
  :afterResponse="(res) => console.log('response', res)"
  :extraConfig="{ headers: { Authorization: token } }"
/>
```

### 分页控制

- 通过 `v-model:currentPage` 与 `v-model:pageSize` 可与父组件联动。
- 设为 `pagination=false` 可隐藏分页区域。

### 暴露方法（通过 ref）

```ts
// 获取选择行
tableRef.value.getSelection()
// 清空选择
tableRef.value.clearSelection()
// 切换某行选择状态
tableRef.value.toggleRowSelection(row, true)
// 刷新（回到第一页）
tableRef.value.refresh()
// 刷新（保留当前页）
tableRef.value.refreshKeepPage()
```

在模板中：

```vue
<Table ref="tableRef" :url="url" />
```

### 使用插槽完全自定义列

```vue
<Table :url="url">
  <template #columns>
    <el-table-column type="index" label="#" width="60" />
    <el-table-column prop="name" label="姓名" />
    <el-table-column label="操作" width="160">
      <template #default="{ row }">
        <el-button link type="primary">查看</el-button>
      </template>
    </el-table-column>
  </template>
</Table>
```

### 注意事项

- 组件内部默认认为后端成功码为 `res.code === 0`，如有差异请在外层请求工具中统一适配，或通过 `afterResponse` 自行处理。
- `search` 为数组，非对象项会被忽略；若存在相同键，后者会覆盖前者（浅合并）。
- 若需要传递更多 `ElTable` 原生属性/事件，直接通过 `v-bind="$attrs"` 透传即可。


