<script>
import Navbar from "../components/Navbar.vue";
</script>

<template>
  <el-header style="height: 15px"> <Navbar /></el-header>
  <div v-if="!isLogin" >
    <p style="left: 520px;top: 60px;position: relative;">请先登录！</p>
    <div class="bottom-save">
        <el-button type="warning" @click="Gotologin">去登录</el-button>
      
    </div>
  </div>
  <div class="common-layout" v-if="isLogin">
    <el-container>
      
      <el-container>
        <el-aside width="450px">
          <div>
            <el-carousel :interval="4000" type="card" height="500px">
              <el-carousel-item
                v-for="item in data"
                :key="item.id"
                style="top: 10%"
              >
                <img :src="item.value" alt="error" class="img-sty" />
              </el-carousel-item>
            </el-carousel>
          </div>
        </el-aside>

        <el-main>
          <div class="common-layout">
            <el-container>
              <el-header
                class="header2-sty"
                style="
                  display: flex;
                  justify-content: center;
                  align-items: center;
                  height: 2px;
                "
                >——活动推荐——</el-header
              >
              <el-main>
                <div class="card-container">
                  <el-card
                    v-for="card in cards"
                    :key="card.id"
                    style="max-width: 270px; height: 250px"
                    class="card"
                  >
                    <!-- <template #header> -->
                    <div class="card-header">
                      <span>{{ card.name }}</span>
                    </div>
                    <!-- </template> -->

                    <img
                      :src="card.imageSrc"
                      alt="error"
                      style="max-width: 270px; height: 160px"
                      class="img-sty"
                      @click="goToDetail"
                      @mouseover="addHoverClass"
                      @mouseout="removeHoverClass"
                      :class="{ 'hover-pointer': isHover }"
                    />
                    <!-- <template #footer class="foot-sty"> -->
                    <div
                      style="
                        margin: 10px;
                        display: flex;
                        justify-content: center;
                      "
                    >
                      <el-button
                        class="custom-button"
                        type="info"
                        :icon="Message"
                        circle
                        title="报名"
                      />
                      <span style="margin: 0 10px"></span>
                      <el-button
                        type="warning"
                        :icon="Star"
                        circle
                        title="收藏"
                      />
                      <!-- </template> -->
                    </div>
                  </el-card>
                </div>
              </el-main>
            </el-container>
          </div>
          <div
            style="
              width: 100%;
              height: 100px;
              display: flex;
              align-items: center;
            "
          >
            <div
              class="example-pagination-block"
              style="margin-top: 20px; margin: 0 auto"
            >
              <div class="example-demonstration">When you have few pages</div>
              <el-pagination layout="prev, pager, next" :total="50" />
            </div>
          </div>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import huodong1 from "../assets/huodong1.jpg";
import huodong2 from "../assets/loginBackground.jpg";
import huodong3 from "../assets/ct4.jpg";
import huodong4 from "../assets/vut.jpg";
import { Message, Star } from "@element-plus/icons-vue";
import { useRouter } from "vue-router";
import { ref, onBeforeMount} from "vue";
import {getUserInfo} from "@/api/home";
const isLogin = ref(false);
let data = [
  {
    id: 1,
    value: huodong1,
  },
  {
    id: 2,
    value: huodong2,
  },
  {
    id: 3,
    value: huodong3,
  },
];
let cards = [
  { id: 1, name: "辩论赛", imageSrc: huodong1 },
  { id: 2, name: "摄影", imageSrc: huodong2 },
  { id: 1, name: "CT4", imageSrc: huodong3 },
  { id: 2, name: "志愿", imageSrc: huodong4 },
  // Add more cards as needed
];

const router = useRouter();
const isHover = ref(false);
const goToDetail = () => {
  router.push({ name: "detail" });
};
const addHoverClass = () => {
  isHover.value = true;
};

const removeHoverClass = () => {
  isHover.value = false;
};

const  Gotologin= () => {
      router.push({ name: 'login' });
    };

// onBeforeMount(() => {
//   getUserInfo().then((res: any) => {
//     // console.log(res);
//     if (res.code == 0) {
//       isLogin.value = true;
//       userInfo.value = res.data;
//     }
//   });
// });
const loadUserInfo = async () => {
  const res = await getUserInfo();
    console.log(res)
  if (res.code === 200) {
    isLogin.value = true;
  } else {
    return;
  }
};

onBeforeMount(() => {
  loadUserInfo();
  console.log("load...")
});
</script>

<style scoped>
 .bottom-save {
    display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: relative;
top: 50px;
  }
.hover-pointer {
  cursor: pointer;
}
.custom-button {
  transition: background-color 0.3s ease; /* 添加过渡效果，使颜色变化更平滑 */
}

.custom-button:hover {
  background-color: lightgray; /* 鼠标悬停时改变按钮的背景色 */
}

.card-header {
  padding: 0px;
  text-align: center;
}

.card-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around; /* 可以根据需要更改 */
}

.card {
  margin: 2px; /* 可以根据需要调整 */
  flex: 0 0 auto; /* 可以根据需要调整 */
  height: 150px;
  width: 300px;
}

.img-sty {
  width: 100%;
  height: 100%;
}
.el-carousel__item h3 {
  display: flex;
  color: #475669;
  opacity: 0.75;
  line-height: 300px;
  margin: 0;
  font-size: 16px;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>
