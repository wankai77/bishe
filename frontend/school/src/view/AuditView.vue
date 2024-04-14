<script>
import Navbar from "../components/Navbar.vue";
import { getActInfo } from "../api/audit";
</script>
<template>
  <el-header style="height: 15px; margin-bottom: 15px"> <Navbar /></el-header>
  <div>
    <div class="crumbs">
      <p style="margin-left: 8%">审核</p>
    </div>
    <div class="container">
      <div class="handle-box">
        <el-select
          v-model="query.address"
          placeholder="类型"
          class="handle-select mr10"
        >
          <el-option key="1" label="志愿类" value="志愿类"></el-option>
          <el-option key="2" label="竞赛类" value="竞赛类"></el-option>
          <el-option key="3" label="讲座" value="讲座"></el-option>
          <el-option key="4" label="其他" value="其他"></el-option>
        </el-select>
        <el-input
          v-model="query.name"
          placeholder="活动名"
          class="handle-input mr10"
          style="width: 150px"
        ></el-input>
        <el-button @click="handleSearch">
          <div
            style="display: flex; flex-direction: column; align-items: center"
          >
            <el-icon><Search /></el-icon>
            <p style="margin: 0">search</p>
          </div>
        </el-button>
      </div>
      <el-table
        :data="tableData"
        border
        class="table"
        ref="multipleTable"
        header-cell-class-name="table-header"
      >
        <el-table-column
          prop="id"
          label="ID"
          width="55"
          align="center"
        ></el-table-column>
        <el-table-column prop="name" label="活动名"></el-table-column>
        <el-table-column label="申请部门">
          <template #default="scope">￥{{  }}</template>
        </el-table-column>
        <el-table-column label="封面(查看大图)" align="center">
          <template #default="scope">
            <el-image
              class="table-td-thumb"
              :src="scope.row.thumb"
              :preview-src-list="[scope.row.thumb]"
            >
            </el-image>
          </template>
        </el-table-column>
        <el-table-column prop="address" label="查看详情"></el-table-column>
        <el-table-column label="状态" align="center">
          <template #default="scope">
            <el-tag
              :type="
                scope.row.state === '通过'
                  ? 'success'
                  : scope.row.state === '已退回'
                  ? 'danger'
                  : scope.row.state === '待审核'
                  ? 'otherStateClass'
                  : ''
              "
              >{{ scope.row.state }}</el-tag
            >
          </template>
        </el-table-column>

        <el-table-column prop="date" label="申请时间"></el-table-column>
        <el-table-column label="操作" width="180" align="center">
          <template #default="scope">
            <el-button
              type="text"
              icon="el-icon-edit"
              @click="handleAudit(scope.$index, scope.row)"
              >通过
            </el-button>
            <el-button
              type="text"
              icon="el-icon-delete"
              class="red"
              @click="handleReturn(scope.$index, scope.row)"
              >退回</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination
          background
          layout="total, prev, pager, next"
          :current-page="query.pageIndex"
          :page-size="query.pageSize"
          :total="pageTotal"
          @current-change="handlePageChange"
        ></el-pagination>
      </div>
    </div>

    <!-- 编辑弹出框 -->
    <el-dialog title="退回" v-model="editVisible" width="30%">
      <el-form label-width="70px">
        <el-form-item label="活动名">
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="原因">
          <el-input v-model="form.address"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="editVisible = false">取 消</el-button>
          <el-button type="primary" @click="saveEdit">确 定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import { getActInfo } from "../api/audit";

const query = reactive({
  address: "",
  name: "",
  pageIndex: 1,
  pageSize: 6,
});
const tableData = ref([]);
const pageTotal = ref(0);
// 获取表格数据
// const getData = () => {
//   getActInfo().then((res) => {
//     console.log(res)
//     tableData.value = res;
//     pageTotal.value = res.pageTotal || 50;
//   });
// };

// 查询操作
const handleSearch = () => {
  query.pageIndex = 1;
  getData();
};
// 分页导航
const handlePageChange = (val) => {
  query.pageIndex = val;
  getData();
};

// 审核操作
const handleAudit = (index) => {
  // 二次确认
  ElMessageBox.confirm("确认通过", "提示", {
    type: "warning",
  })
    .then(() => {
      ElMessage.success("通过");
    //   console.log(index);
      tableData.value[index].state="通过";
    })
    .catch(() => {});
};

// 退回
const editVisible = ref(false);
let form = reactive({
  name: "",
  address: "",
});
let idx = -1;
const handleReturn = (index, row) => {
  idx = index;
  Object.keys(form).forEach((item) => {
    form[item] = row[item];
  });
  editVisible.value = true;
};
const saveEdit = () => {
  editVisible.value = false;
  console.log(form);
  ElMessage.success(`${form.name}已退回`);
  tableData.value[idx].state="已退回";
};
</script>

<style scoped>
.handle-box {
  margin-bottom: 20px;
}

.handle-select {
  width: 120px;
}

.handle-input {
  width: 300px;
  display: inline-block;
}
.table {
  width: 100%;
  font-size: 14px;
}
.red {
  color: #ff0000;
}
.mr10 {
  margin-right: 10px;
}
.table-td-thumb {
  display: block;
  margin: auto;
  width: 40px;
  height: 40px;
}
</style>
