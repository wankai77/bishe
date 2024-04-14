<template>
    <div class="not-found">
      <el-result title="404" sub-title="您访问的页面不存在">
        <template #icon>
          <el-image
            src="https://iconfont.alicdn.com/s/f415cb8b-816d-4c6b-a84e-b7d7c591e3e6_origin.svg"
          ></el-image>
        </template>
        <template #extra>
          <el-button type="primary" @click="handleBackHome">返回首页</el-button>
        </template>
      </el-result>
    </div>
  </template>
  
  <script setup>
  import {} from "vue";
  import { useRouter } from "vue-router";
  
  const router = useRouter();
  
  // 返回首页
  const handleBackHome = () => {
    router.push("/home");
  };
  </script>
  
  <style scoped>

</style>