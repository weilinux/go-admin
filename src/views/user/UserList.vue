<template>
<div>
    {{ title }}
</div>
  <el-table :data="users" style="width: 100%" height="250">
    <el-table-column type="index" label="序号" width="120" />
    <el-table-column prop="id" label="用户编号" width="120" />
    <el-table-column prop="UserName" label="名称" width="120" />
    <el-table-column prop="Password" label="密钥" width="520" />
    <el-table-column fixed="right" label="功能操作" width="150">
      <template #default="scope">
        <el-button
            link
            type="primary"
            size="small"
            @click="show_user(scope.$index, scope.row)">详细</el-button>
        <el-button
            link
            type="primary"
            size="small"
            @click="edit_user(scope.$index, scope.row)">编辑</el-button>
        <el-button
            link
            type="primary"
            size="small"
            @click="delete_user(scope.$index, scope.row)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
  <el-button class="mt-4" style="width: 10%;align: right"  @click="add_user">添加</el-button>
</template>

<script>
import UserService from '@/services/user.service';

export default {
  data: function() {
    return {
      title: "用户列表",
      users: [],
    }
  },
  methods: {
    show_user: function(index, row) {
      this.$router.push({name: 'User', params: {id: row.id}})
    },
    edit_user: function(index, row) {
      this.$router.push({name: 'UserEdit', params: {id: row.id}})
    },
    add_user: function() {
      this.$router.push({name: 'UserNew'})
    },
    delete_user: function(index, row) {
      UserService.delUser(row.id).then(() => {
        this.users.splice(index, 1)
      }, (error) => {
        console.log(error.toString())
      })
    },
    get_users: function() {
      UserService.getUserList().then(response => {
        this.users = response.data.data.rows;
      }, (error) => {
        console.log(error)
      })
    }
  },
  created: function() {
    this.get_users()
  }
}
</script>

<style scoped>
td {
  border: 3px solid grey;
}
</style>