# x-admin 后台前端

基于 Vue 3 + TypeScript + Vite + Tailwind CSS + Element Plus 的管理后台。

## 技术栈

- **Vue 3** - Composition API + `<script setup>`
- **TypeScript**
- **Vite** - 构建工具
- **Tailwind CSS** - 原子化 CSS
- **Element Plus** - UI 组件库，支持深色模式

## 样式规范

- 主色调：白色为主，不同深浅度的白/灰阶梯区分层级
- 圆角：除 body 外，卡片、侧边栏、内容区等采用圆角
- 深色模式：通过 Element Plus + VueUse 的 useDark 实现切换

## 开发

```bash
npm install
npm run dev
```

访问 http://localhost:5174

## 构建

```bash
npm run build
```

## API 代理

开发环境 `/api` 会代理到 `http://127.0.0.1:8000`（后端服务）。
