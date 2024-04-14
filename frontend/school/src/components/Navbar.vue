<template>
  <nav class="navbar navbar-expand-lg bg-body-tertiary" style="width:100%;left: 0%; background-color: white;">
    <div class="container-fluid">
      <a class="navbar-brand" >活动中心</a>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarScroll"
        aria-controls="navbarScroll"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarScroll">
        <ul
          class="navbar-nav me-auto my-2 my-lg-0 navbar-nav-scroll"
          style="--bs-scroll-height: 100px"
        >
          <li class="nav-item">
            <a class="nav-link active" aria-current="page" href="#" @click="GotoHome">首页</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#" @click="GotoPersonal">个人中心</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#" @click="GotoMessage">消息中心</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#" @click="GotoAudit">内容审核</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#" @click="GotoRelease">活动发布</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#" @click="Gotologin" style="margin-left: 300px;">退出登录</a>
          </li>
        </ul>
        <form class="d-flex" role="search">
          <input
            class="form-control me-2"
            type="search"
            placeholder="Search"
            aria-label="Search"
          />
          <button class="btn btn-outline-success" type="submit">Search</button>
        </form>
      </div>
    </div>
  </nav>
</template>

<script setup>

// export default {
//   name: "Navbar",
// };

import { useRouter } from 'vue-router';

    const router = useRouter();

    const GotoPersonal = () => {
      router.push({ name: 'personal' });
    };
    const GotoHome = () => {
      router.push({ name: 'home' });
    };
    const GotoMessage = () => {
      router.push({ name: 'message' });
    };
    const GotoAudit = () => {
      router.push({ name: 'audit' });
    };
    const  Gotologin= () => {
      router.push({ name: 'login' });
    };
    const GotoRelease = () => {
      router.push({ name: 'release' });
    };
    
function GetUserInfoF(){
  axios({
      method: "get",
      url: "http://127.0.0.1:8888/api/user/getInfo",
    })
}

</script>
<style scoped>

.navbar {
  position: fixed; /* 将导航栏固定在页面上方 */
  top: 0; /* 将导航栏放置在页面顶部 */
  width: 100%; /* 确保导航栏占据整个页面宽度 */
  z-index: 1000; /* 设置 z-index 以确保导航栏显示在其他元素之上 */
  left: 0%;
}
.d-flex {
  margin-right: 0%;
}

</style>
