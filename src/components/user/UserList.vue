<template>
  <div>
    {{ title }}
  </div>
  <div class="flex">
    <el-form :model="searchForm" inline>
      <el-form-item label="用户名称">
        <el-input v-model="user" placeholder="输入用户名称查询"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button @click="handleFilter">查询</el-button>
      </el-form-item>
      <el-form-item>
        <el-button @click="add_user">新增用户</el-button>
      </el-form-item>
    </el-form>
  </div>

  <el-table :data="users" style="width: 100%" height="250">
    <el-table-column type="index" label="序号" width="120"/>
    <el-table-column prop="id" label="用户编号" width="120"/>
    <el-table-column prop="UserName" label="名称" width="120"/>
    <el-table-column prop="Password" label="密钥" width="520"/>
    <el-table-column fixed="right" label="功能操作" width="150">
      <template #default="scope">
        <el-button
            link
            type="primary"
            size="small"
            @click="show_user(scope.$index, scope.row)">详细
        </el-button>
        <el-button
            link
            type="primary"
            size="small"
            @click="edit_user(scope.$index, scope.row)">编辑
        </el-button>
        <el-button
            link
            type="primary"
            size="small"
            @click="delete_user(scope.$index, scope.row)">删除
        </el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import UserService from '@/services/user.service';
import {ElMessageBox} from 'element-plus';

export default {
  data: function () {
    return {
      title: "用户列表",
      users: [],
    }
  },
  methods: {
    show_user: function (index, row) {
      this.$router.push({name: 'User', params: {id: row.id}})
    },
    edit_user: function (index, row) {
      this.$router.push({name: 'UserEdit', params: {id: row.id}})
    },
    add_user: function () {
      this.$router.push({name: 'UserNew'})
    },
    delete_user: function (index, row) {
      ElMessageBox.confirm("确实要删除这个用户吗?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
      }).then(() => {
        UserService.delUser(row.id).then(() => {
          this.users.splice(index, 1)
        })
        alert("删除成功!");
      })
          .catch(() => {
            alert("取消删除!");
          })
    },
    get_users: function () {
      UserService.getUserList().then(response => {
        this.users = response.data.data.rows;
      }, (error) => {
        console.log(error)
      })
    }
  },
  created: function () {
    this.get_users()
  }
}
</script>

<style scoped>
td {
  border: 3px solid grey;
}

.flex {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>