<template>
  <div class="login-page">
    <div class="login-register">
      <div class="contain">
        <div class="big-box" :class="{ active: isLogin }">
          <div class="big-contain" key="bigContainLogin" v-if="isLogin">
            <div class="btitle">账户登录</div>
            <div class="bform">
              <input type="username" placeholder="用户名" v-model="form.username" />
              <input
                type="password"
                placeholder="密码"
                v-model="form.userpwd"
              />
            </div>
            <button class="bbutton" @click="login">登录</button>
          </div>
          <div class="big-contain" key="bigContainRegister" v-else>
            <div class="btitle">创建账户</div>
            <div class="bform">
              <input type="text" placeholder="用户名" v-model="form.username" />
              <input type="email" placeholder="邮箱" v-model="form.useremail" />
              <input
                type="password"
                placeholder="密码"
                v-model="form.userpwd"
              />
            </div>
            <button class="bbutton" @click="register">注册</button>
          </div>
        </div>
        <div class="small-box" :class="{ active: isLogin }">
          <div class="small-contain" key="smallContainRegister" v-if="isLogin">
            <div class="stitle">你好，同学!</div>
            <p class="scontent">开始注册，即可加入</p>
            <button class="sbutton" @click="changeType">注册</button>
          </div>
          <div class="small-contain" key="smallContainLogin" v-else>
            <div class="stitle">欢迎回来!</div>
            <p class="scontent">
              欢迎回来，请登录你的账户，你也可向我们反馈问题
            </p>
            <button class="sbutton" @click="changeType">登录</button>

            <button class="sbutton" @click="feedbackType">反馈</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import axios from "axios";
import { useRouter } from 'vue-router';
const router = useRouter();

import { ElMessage } from "element-plus";
import { registerUser,loginUser} from "@/api";

import { useUserStore } from "@/stores/user";

const userStore = useUserStore();
let isLogin = ref(false);
let emailError = ref(false);
let passwordError = ref(false);
let existed = ref(false);
let form = ref({
  username: "",
  useremail: "",
  userpwd: "",
});

function changeType() {
  isLogin.value = !isLogin.value;
  form.value.username = "";
  form.value.useremail = "";
  form.value.userpwd = "";
}
function login() {
  if (form.value.userename != "" && form.value.userpwd != "") {
	const formData = new FormData();
    formData.append("username", form.value.username);
    formData.append("password", form.value.userpwd);
	
    loginUser(formData)  //登录请求
      .then((res) => {
        if (res.data.code == 200){
			ElMessage({
            message: "登录成功.",
            type: "success",
          });
          console.log(res)
          userStore.setToken(res.data.Data.token);
        userStore.setUsername(form.value.username);

        // 必须延缓一段时间，避免持久化存储失败！
        timer = setTimeout(() => {
		  router.push({ name: 'home' });
        }, 1000);
      
		}else{
			ElMessage({
            message: res.data.message,
            type: "error",
          });
		}
      })
      .catch((err) => {
        console.log(err)
      });
  } else {
    alert("填写不能为空！");
  }
}
function register() {
  if (
    form.value.username != "" &&
    form.value.useremail != "" &&
    form.value.userpwd != ""
  ) {
    const formData = new FormData();
    formData.append("username", form.value.username);
    formData.append("email", form.value.useremail);
    formData.append("password", form.value.userpwd);
    registerUser(formData)  //发送注册请求
      .then((res) => {
		console.log(res.data);
        if (res.data.code == 200) {
          ElMessage({
            message: "注册成功.",
            type: "success",
          });
		  changeType()
        } else {
          ElMessage({
            message: res.data.message,
            type: "error",
          });
        }
      })
      .catch((err) => {
        console.log(err);
      });
  } else {
    ElMessage({
            message: "填写不能为空",
            type: "error",
          });
  }
}
</script>

<style scoped="scoped">
.login-page {
  height: 100%; /* 确保容器铺满整个浏览器视口 */
  background-image: url("../assets/gzdu.jpg"); /* 添加背景图片 */
  background-size: cover; /* 背景图片大小覆盖整个容器 */
  background-repeat: no-repeat; /* 防止背景图片重复 */
  font-family: Arial, sans-serif; /* 设置字体，可选 */
}

.login-register {
  width: 100vw;
  height: 100vh;
  box-sizing: border-box;
}
.contain {
  width: 60%;
  height: 70%;
  position: relative;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 20px;
  box-shadow: 0 0 3px #f0f0f0, 0 0 6px #f0f0f0;
}
.big-box {
  width: 70%;
  height: 100%;
  position: absolute;
  top: 0;
  left: 30%;
  transform: translateX(0%);
  transition: all 1s;
  background-color: rgba(255, 255, 255, 0.3);
}
.big-contain {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.btitle {
  font-size: 1.5em;
  font-weight: bold;
  color: rgb(255, 255, 255);
}
.bform {
  width: 100%;
  height: 40%;
  padding: 2em 0;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
}
.bform .errTips {
  display: block;
  width: 50%;
  text-align: left;
  color: red;
  font-size: 0.7em;
  margin-left: 1em;
}
.bform input {
  width: 50%;
  height: 30px;
  border: none;
  outline: none;
  border-radius: 10px;
  padding-left: 2em;
  background-color: rgba(255, 255, 255, 0.489);
}
.bbutton {
  width: 20%;
  height: 40px;
  border-radius: 24px;
  border: none;
  outline: none;
  background-color: rgba(255, 255, 255, 0.532);
  color: #137196;
  font-size: 0.9em;
  cursor: pointer;
}
.small-box {
  width: 30%;
  height: 100%;
  position: absolute;
  top: 0;
  left: 0;
  transform: translateX(0%);
  transition: all 1s;
  border-top-left-radius: inherit;
  border-bottom-left-radius: inherit;
  background-color: rgba(255, 255, 255, 0.3);
}
.small-contain {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.stitle {
  font-size: 1.5em;
  font-weight: bold;
  color: #fff;
}
.scontent {
  font-size: 0.8em;
  color: #fff;
  text-align: center;
  padding: 2em 4em;
  line-height: 1.7em;
}
.sbutton {
  width: 60%;
  height: 40px;
  border-radius: 24px;
  border: 1px solid #fff;
  outline: none;
  background-color: rgba(255, 255, 255, 0.375);
  color: #156f92;
  font-size: 0.9em;
  cursor: pointer;
  margin-bottom: 20px; /* 为按钮添加下边距 */
}

.big-box.active {
  left: 0;
  transition: all 0.5s;
}
.small-box.active {
  left: 100%;
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
  border-top-right-radius: inherit;
  border-bottom-right-radius: inherit;
  transform: translateX(-100%);
  transition: all 1s;
}
</style>
