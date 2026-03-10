/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        /* 白色阶梯层级：以白色为主，不同深浅度区分 */
        'layer': {
          0: '#ffffff',   /* 最顶层 */
          1: '#fafafa',   /* 一级 */
          2: '#f5f5f5',   /* 二级 */
          3: '#f0f0f0',   /* 三级 */
          4: '#ebebeb',   /* 四级 */
        },
      },
      borderRadius: {
        'panel': '12px',
        'card': '10px',
      },
    },
  },
  plugins: [],
}
