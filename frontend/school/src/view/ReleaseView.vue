<script>
import Navbar from "../components/Navbar.vue";
</script>
<template>
  <div>
    <el-header style="height: 15px; margin-bottom: 15px"> <Navbar /></el-header>
    <div class="crumbs">
            <el-breadcrumb separator="/">
                <el-breadcrumb-item>
                    <i class="el-icon-lx-calendar"></i> 
                </el-breadcrumb-item>
                <el-breadcrumb-item>活动发布</el-breadcrumb-item>
            </el-breadcrumb>
        </div>
    <div class="container">
      <div class="form-box">
        <el-form ref="formRef" :rules="rules" :model="form" label-width="80px">
          <el-form-item label="活动名称" prop="name">
            <el-input v-model="form.name"></el-input>
          </el-form-item>
          <el-form-item label="活动类型" prop="region">
            <el-select v-model="form.region" placeholder="请选择">
              <el-option key="bbk" label="讲座类" value="bbk"></el-option>
              <el-option key="xtc" label="志愿类" value="xtc"></el-option>
              <el-option key="imoo" label="竞赛类" value="imoo"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="截止时间">
            <el-col :span="11">
              <el-form-item prop="date1">
                <el-date-picker
                  type="date"
                  placeholder="选择日期"
                  v-model="form.date1"
                  style="width: 100%"
                ></el-date-picker>
              </el-form-item>
            </el-col>
            <el-col class="line" :span="2"> ————</el-col>
            <el-col :span="11">
              <el-form-item prop="date2">
                <el-time-picker
                  placeholder="选择时间"
                  v-model="form.date2"
                  style="width: 100%"
                >
                </el-time-picker>
              </el-form-item>
            </el-col>
          </el-form-item>
          <el-form-item label="活动详情" prop="desc">
            <el-input type="textarea" rows="15" v-model="form.desc"></el-input>
          </el-form-item>
          <el-form-item label="上传图片" prop="img">
            <el-upload
        action="/src/assets"
        :on-success="handleSuccess"
        :before-upload="beforeUpload"
        :limit="1"
        list-type="picture-card">
        <i class="el-icon-plus"></i>
      </el-upload>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">提交</el-button>
            <el-button @click="onReset">重置</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from "vue";
import { ElMessage } from "element-plus";


const rules = {
  name: [{ required: true, message: "请输入活动名称", trigger: "blur" }],
  region:[{ required: true, message: "请选择活动类型", trigger: "blur" }],
  date1:[{ required: true, message: "请选择日期", trigger: "blur" }],
  date2:[{ required: true, message: "请选择时间", trigger: "blur" }],
  desc:[{ required: true, message: "请输入活动详情", trigger: "blur" }],
  img:[{ required: true, message: "请上传相关图片", trigger: "blur" }],
};
const formRef = ref(null);
const form = reactive({
  name: "",
  region: "",    //活动类型
  date1: "",
  date1: "",
  desc: "",
});
// 提交
const onSubmit = () => {
  // 表单校验
  formRef.value.validate((valid) => {
    if (valid) {
      console.log(form);
      ElMessage.success("提交成功！");
    } else {
      return false;
    }
  });
};
// 重置
const onReset = () => {
  formRef.value.resetFields();
};
const handleSuccess = (response, file)=> {
      // 图片上传成功后的处理逻辑
      // 可以在这里获取到服务器返回的图片地址等信息，并做相应处理
      console.log(response);
    };
const beforeUpload = (file)=> {
      // 文件上传前的钩子，可以在这里对文件进行一些验证操作，比如文件类型、大小等
      const isJPG = file.type === 'image/jpeg';
      const isPNG = file.type === 'image/png';
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG && !isPNG) {
        message.error('上传图片只能是 JPG/PNG 格式!');
      }
      if (!isLt2M) {
        message.error('上传图片大小不能超过 3MB!');
      }
      return (isJPG || isPNG) && isLt2M;
    }
</script>
