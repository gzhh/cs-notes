# Vue
- https://github.com/vuejs/core
- https://vuejs.org

Docs
- https://vuejs.org/guide
- https://play.vuejs.org
- https://cn.vuejs.org
- https://cn.vuejs.org/guide/introduction
- https://cn.vuejs.org/guide/quick-start
- https://cn.vuejs.org/guide/essentials/application.html
- https://cn.vuejs.org/tutorial
- https://cn.vuejs.org/examples
- https://cn.vuejs.org/api
- https://cn.vuejs.org/api/application.html

Vue2
- https://github.com/vuejs/vue
- https://v2.vuejs.org
- https://v2.cn.vuejs.org/v2/guide


## Best Practice
- https://github.com/vuejs/awesome-vue
- 手摸手，带你用vue撸后台 系列一（基础篇）https://juejin.cn/post/6844903476661583880
- 揭秘从新手到Vue高级工程师的养成之路 http://www.imooc.com/article/270401


## Vue 生态圈
- https://github.com/vuejs/router
- https://github.com/vuejs/vuex
- https://github.com/vuejs/vue-loader
- https://github.com/vuejs/vue-test-utils
- https://github.com/vuejs/devtools-v6
- https://github.com/vuejs/vue-cli
- https://github.com/vuejs/vetur
- https://github.com/vuejs-templates/webpack

Ref
- https://panjiachen.github.io/vue-element-admin-site/zh/guide/#vue-生态圈



## Tookit
Vite
- https://github.com/vitejs/vite
- https://vite.dev

Vue tooling for VS Code.
- https://github.com/vuejs/vetur


## Library
### UI Library
- Element - Vue2  https://github.com/ElemeFE/element
  - Mobile UI elements https://github.com/ElemeFE/mint-ui
- Element Plus - Vue3https://github.com/element-plus/element-plus


## Vue admin interfaces
Vue2
- https://github.com/PanJiaChen/vue-element-admin
  - 开发建议：把 vue-element-admin当做工具箱或者集成方案仓库，在 vue-admin-template 的基础上进行二次开发，想要什么功能或者组件就去 vue-element-admin 那里复制过来。
  - https://panjiachen.github.io/vue-element-admin-site/
  - https://panjiachen.github.io/vue-element-admin-site/zh/guide/
  - https://github.com/PanJiaChen/vue-admin-template
  - https://github.com/Armour/vue-typescript-admin-template

Vue3
- https://github.com/zxwk1998/vue-admin-better


## Quick Start - Demo
1.安装 Vue cli 命令
- `npm install -g @vue/cli`

2.项目初始化
- `vue create admin-dashboard`

3.配置 Vue Router
- 后台管理系统通常有多个页面，需要使用 Vue Router 来管理路由。编辑 src/router/index.js 文件，定义基本的路由结构：
```
import { createRouter, createWebHistory } from 'vue-router';
import Dashboard from '@/views/Dashboard.vue';
import Users from '@/views/Users.vue';
import Settings from '@/views/Settings.vue';

const routes = [
  { path: '/', name: 'Dashboard', component: Dashboard },
  { path: '/users', name: 'Users', component: Users },
  { path: '/settings', name: 'Settings', component: Settings },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
```

4.创建基础页面组件
- 在 src/views 目录下，创建 Dashboard.vue、Users.vue 和 Settings.vue 页面组件：类似的，你可以创建 Users.vue 和 Settings.vue。
```
<!-- Dashboard.vue -->
<template>
  <div>
    <h1>Dashboard</h1>
    <p>这里显示仪表盘内容。</p>
  </div>
</template>

<script>
export default {
  name: 'UserDashboard',
};
</script>
```

5.添加导航栏
- 后台管理系统通常有一个侧边栏或顶部导航栏。可以在 App.vue 中添加一个简单的导航栏：
```
<template>
  <div id="app">
    <nav>
      <ul>
        <li><router-link to="/">仪表盘</router-link></li>
        <li><router-link to="/users">用户管理</router-link></li>
        <li><router-link to="/settings">系统设置</router-link></li>
      </ul>
    </nav>
    <router-view></router-view> <!-- 路由内容渲染位置 -->
  </div>
</template>

<script>
export default {
  name: 'App',
};
</script>

<style>
nav {
  background-color: #333;
  padding: 10px;
}

nav ul {
  list-style-type: none;
  padding: 0;
}

nav ul li {
  display: inline;
  margin-right: 15px;
}

nav ul li a {
  color: white;
  text-decoration: none;
}
</style>
```
- 改成 element ui
```
<template>
  <el-menu :router="true" mode="horizontal" background-color="#333" text-color="#fff" active-text-color="#ffd04b">
    <el-menu-item index="/">仪表盘</el-menu-item>
    <el-menu-item index="/users">用户管理</el-menu-item>
    <el-menu-item index="/settings">系统设置</el-menu-item>
  </el-menu>
  <router-view></router-view>
</template>

<script>
export default {
  name: 'App',
};
</script>
```

6.样式美化
- 在 main.js 中引入 Element UI： 
```
import { createApp } from 'vue'
import App from './App.vue'
import router from './router';
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

const app = createApp(App)

app.use(router);
app.use(ElementPlus)
app.mount('#app')
```

7.部署与打包
- `npm run serve` 本地实时测试
- `npm run build` 这将生成一个 dist 文件夹，里面包含打包后的项目，接着你可以将其部署到任何支持静态文件的服务器上。
