<script>
import Navbar from "../components/Navbar.vue";
</script>

<template>
    <el-header style="height: 15px;margin-bottom:15px;"> <Navbar /></el-header>
    <div style="background-color: #f4f4f4;">
    <el-row :gutter="20">
      <el-col :span="8">
        <el-card shadow="hover" class="mgb20" style="height: 500px">
          <div>
            <img src="../assets/vut.jpg" alt="error" style="width: 100%; height:
            100%;"
          </div>
        </el-card>
        <el-card shadow="hover" style="height: 252px">
          <template #header>
            <div class="clearfix">
              <span>....</span>
            </div>
          </template>
          .....
        </el-card>
      </el-col>
      <el-col :span="16">
        <el-row :gutter="20" class="mgb20">
          <el-col :span="8">
            <el-card shadow="hover" :body-style="{ padding: '0px' }">
              <div class="grid-content grid-con-1">
                <i class="el-icon-user-solid grid-con-icon"></i>
                <div class="grid-cont-right">
                  <div class="grid-num">1234</div>
                  <div>用户访问量</div>
                </div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="8">
            <el-card shadow="hover" :body-style="{ padding: '0px' }">
              <div class="grid-content grid-con-2">
                <i class="el-icon-message-solid grid-con-icon"></i>
                <div class="grid-cont-right">
                  <div class="grid-num">321</div>
                  <div>报名人数</div>
                </div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="8" style="padding-right: 1px;">
            <el-card shadow="hover" :body-style="{ padding: '0px' }">
              <div class="grid-content grid-con-3">
                <i class="el-icon-s-goods grid-con-icon"></i>
                <div class="grid-cont-right">
                  <div class="grid-num">2024-3-28</div>
                  <div>截止日期</div>
                </div>
              </div>
            </el-card>
          </el-col>
        </el-row>
        <el-card shadow="hover" style="height: 350px">
          <template #header>
            <div class="clearfix">
              <span style="display: block; text-align: center">活动详情</span>
            </div>
          </template>

          <p>
            ⭐“广大的春天”艺术体验周志愿者招募 ⏰时间：
            1⃣️3月26号（周二）下午15:30～18:30 2⃣️3月27号（周三）下午
            17:00～20:00 3⃣️3月28号（周四）下午 17:00～20:00
            📍志愿地点：党建红色文化广场（何世杰体育馆前）
            ‼️工作内容：协助公共艺术中心活动 📌志愿时长:按工作时间结算
            🔅招募人数： 星期二17人 星期三19人 星期四18人
            报名成功请加入链接下方二维码进入志愿者群，谢谢配合！
          </p>
          <el-button type="info" round style="margin: 5px;margin-right: 20px;" @click="Application">报名</el-button>
    <el-button type="warning" round @click="Collection">收藏</el-button>
        </el-card>
        
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="12">
        <el-card shadow="hover">
          <schart
            ref="bar"
            class="schart"
            canvasId="bar"
            :options="options"
          ></schart>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card shadow="hover">
          <schart
            ref="line"
            class="schart"
            canvasId="line"
            :options="options2"
          ></schart>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import * as Schart from "schart.js";
import { reactive } from "vue";
import { ElMessage, ElMessageBox } from 'element-plus'
// name: "dashboard",
// components: { Schart },
const name = localStorage.getItem("ms_username");
const role = name === "admin" ? "超级管理员" : "普通用户";
const Application = () =>{
  ElMessageBox.confirm(
    '确认报名此项活动?',
    'Confirm',
    {
      confirmButtonText: 'Yeah',
      cancelButtonText: 'Cancel',
      type: 'success',
      draggable: true,
    }
  )
    .then(() => {
      ElMessage({
        type: 'success',
        message: '报名成功',
      })
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '已取消',
      })
    })
}
const Collection = () =>{

}

</script>

<style scoped>
.el-row {
  margin-bottom: 20px;
}

.grid-content {
  display: flex;
  align-items: center;
  height: 100px;
}

.grid-cont-right {
  flex: 1;
  text-align: center;
  font-size: 14px;
  color: #999;
}

.grid-num {
  font-size: 30px;
  font-weight: bold;
}

.grid-con-icon {
  font-size: 50px;
  width: 100px;
  height: 100px;
  text-align: center;
  line-height: 100px;
  color: #fff;
}

.grid-con-1 .grid-con-icon {
  background: rgb(45, 140, 240);
}

.grid-con-1 .grid-num {
  color: rgb(45, 140, 240);
}

.grid-con-2 .grid-con-icon {
  background: rgb(100, 213, 114);
}

.grid-con-2 .grid-num {
  color: rgb(45, 140, 240);
}

.grid-con-3 .grid-con-icon {
  background: rgb(242, 94, 67);
}

.grid-con-3 .grid-num {
  color: rgb(242, 94, 67);
}

.user-info {
  display: flex;
  align-items: center;
  padding-bottom: 20px;
  border-bottom: 2px solid #ccc;
  margin-bottom: 20px;
}

.user-avator {
  width: 120px;
  height: 120px;
  border-radius: 50%;
}

.user-info-cont {
  padding-left: 50px;
  flex: 1;
  font-size: 14px;
  color: #999;
}

.user-info-cont div:first-child {
  font-size: 30px;
  color: #222;
}

.user-info-list {
  font-size: 14px;
  color: #999;
  line-height: 25px;
}

.user-info-list span {
  margin-left: 70px;
}

.mgb20 {
  margin-bottom: 20px;
}

.todo-item {
  font-size: 14px;
}

.todo-item-del {
  text-decoration: line-through;
  color: #999;
}

.schart {
  width: 100%;
  height: 300px;
}
</style>
