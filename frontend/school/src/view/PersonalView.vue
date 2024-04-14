<script>
import Navbar from "../components/Navbar.vue";
</script>
<template>
  <div>
    <el-header style="height: 15px; margin-bottom: 15px"> <Navbar /></el-header>
    <el-row :gutter="20">
      <el-col :span="12">
        <el-card shadow="hover">
          <template #header>
            <div class="clearfix">
              <span>基础信息</span>
            </div>
          </template>
          <div class="info">
            <div class="info-image" @click="showDialog">
              <img :src="avatarImg" />
              <span class="info-edit">
                <i class="el-icon-lx-camerafill"></i>
              </span>
            </div>
            <div class="info-name">{{ userinfo.username }}</div>
            <div class="info-desc">可点击头像即可传头像</div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card shadow="hover">
          <template #header>
            <div class="clearfix">
              <span>账户编辑</span>
            </div>
          </template>
          <el-form label-width="90px">
            <el-form-item label="用户名：">
              {{ userinfo.username }}
            </el-form-item>
            <el-form-item label="新密码：">
              <el-input type="password" v-model="form.old"></el-input>
            </el-form-item>
            <el-form-item label="确认密码：">
              <el-input type="password" v-model="form.new"></el-input>
            </el-form-item>
            <el-form-item label="个人简介：">
              <el-input v-model="form.desc"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="onSubmit">保存</el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
    </el-row>
    <el-dialog title="裁剪图片" v-model="dialogVisible" width="600px">
      <vue-cropper
        ref="cropper"
        :src="imgSrc"
        :ready="cropImage"
        :zoom="cropImage"
        :cropmove="cropImage"
        style="width: 100%; height: 400px"
      ></vue-cropper>

      <template #footer>
        <span class="dialog-footer">
          <el-button class="crop-demo-btn" type="primary"
            >选择图片
            <input
              class="crop-input"
              type="file"
              name="image"
              accept="image/*"
              @change="setImage"
            />
          </el-button>
          <el-button type="primary" @click="saveAvatar">上传并保存</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { onBeforeMount, reactive, ref } from "vue";
import VueCropper from "vue-cropperjs";
import "cropperjs/dist/cropper.css";
import avatar from "../assets/chushi.png";
import { getUserInfo } from "@/api/home";
import { ResetPsd } from "@/api/personal";

const form = ref({
  old: "",
  new: "",
  desc: "不可能！我的代码怎么可能会有bug！",
});
const onSubmit = () => {
  console.log(form.value.new, form.value.old);
  console.log(form.value.old === form.value.new)
  if (form.value.old == "" && form.value.new == "") {
    console.log("填写为空 ")
    ElMessage({
      message: "填写不能为空",
      type: "error",
    });
  } else if (form.value.old !== form.value.new) {
    ElMessage({
      
      message: "两次输入的密码不同",
      type: "error",
    });
  } else {
    console.log("post.....");
    const formData = new FormData();
    formData.append("password", form.value.new);
    formData.append("name", userinfo.username);
    formData.append("desc", form.value.desc);
    ResetPsd(formData) //发送修改请求
      .then((res) => {
        console.log(res.data);
        if (res.data.code == 200) {
          ElMessage({
            message: "修改成功.",
            type: "success",
          });
          changeType();
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
  }
};

const avatarImg = ref(avatar);
const imgSrc = ref("");
const cropImg = ref("");
const dialogVisible = ref(false);
const cropper = ref(null);

const showDialog = () => {
  dialogVisible.value = true;
  imgSrc.value = avatarImg.value;
};

const setImage = (e) => {
  const file = e.target.files[0];
  if (!file.type.includes("image/")) {
    return;
  }
  const reader = new FileReader();
  reader.onload = (event) => {
    dialogVisible.value = true;
    imgSrc.value = event.target.result;
    cropper.value && cropper.value.replace(event.target.result);
  };
  reader.readAsDataURL(file);
};

const cropImage = () => {
  cropImg.value = cropper.value.getCroppedCanvas().toDataURL();
};

const saveAvatar = () => {
  avatarImg.value = cropImg.value;
  dialogVisible.value = false;
};

onBeforeMount(() => {
  loadUserInfo();
  console.log("load...");
});
var userinfo = ref({
  username: "",
});
const loadUserInfo = async () => {
  const res = await getUserInfo();
  console.log(res);
  userinfo.value.username = res.Data.Name;
  if (res.code === 200) {
  } else {
    return;
  }
};
</script>

<style scoped>
.info {
  text-align: center;
  padding: 35px 0;
}
.info-image {
  position: relative;
  margin: auto;
  width: 100px;
  height: 100px;
  background: #f8f8f8;
  border: 1px solid #eee;
  border-radius: 50px;
  overflow: hidden;
}
.info-image img {
  width: 100%;
  height: 100%;
}
.info-edit {
  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  opacity: 0;
  transition: opacity 0.3s ease;
}
.info-edit i {
  color: #eee;
  font-size: 25px;
}
.info-image:hover .info-edit {
  opacity: 1;
}
.info-name {
  margin: 15px 0 10px;
  font-size: 24px;
  font-weight: 500;
  color: #262626;
}
.crop-demo-btn {
  position: relative;
}
.crop-input {
  position: absolute;
  width: 100px;
  height: 40px;
  left: 0;
  top: 0;
  opacity: 0;
  cursor: pointer;
}
</style>
