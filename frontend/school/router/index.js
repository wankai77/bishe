import { createRouter, createWebHistory } from 'vue-router';
import Login from '@/view/LoginView.vue';
import Home from '@/view/Home.vue';
import Detail from '@/view/DetailView.vue';
import Personal from '@/view/PersonalView.vue';
import Message from '@/view/MessageView.vue';
import Audit from '@/view/AuditView.vue';
import Release from '@/view/ReleaseView.vue'
import { createApp } from 'vue'; // 导入createApp函数

const routes = [
    {
      path: '/',
      name: 'login',
      component: Login
    },
    {
      path: '/home',
      name: 'home',
      component: Home
    },
    {
      path: '/detail',
      name: 'detail',
      component: Detail
    },
    {
      path: '/personal',
      name: 'personal',
      component: Personal
    },
    {
      path: '/message',
      name: 'message',
      component: Message
    },
    {
      path: '/audit',
      name: 'audit',
      component: Audit
    },
    {
      path: '/release',
      name: 'release',
      component: Release
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes
});
export default router;
const app = createApp({}); // 使用createApp函数创建Vue应用实例
app.use(router); // 将router挂载到Vue应用
app.mount('#app'); // 挂载Vue应用到页面上，假设您的根元素id为app
