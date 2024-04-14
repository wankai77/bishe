import './assets/main.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'element-plus/dist/index.css'
import { createApp } from 'vue';
import {createPinia} from 'pinia'
import App from './App.vue';
import router from '../router/index';
import ElementPlus from 'element-plus'; // 导入 Element Plus
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
const app = createApp(App);

// 全局注册 Element Plus 的所有图标组件
Object.entries(ElementPlusIconsVue).forEach(([key, component]) => {
    app.component(key, component);
});

// 使用 Element Plus 插件
app.use(ElementPlus);

// 使用 Vue Router 插件
app.use(router);

const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);
app.use(pinia)
// 将应用挂载到 DOM 中
app.mount('#app');
