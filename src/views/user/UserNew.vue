<template>
  <h4>创建用户</h4>
  <el-form :model="form" label-width="120px" width="50%">
    <el-form-item label="用户ID">
      <el-input v-model=user.ID />
    </el-form-item>
    <el-form-item label="用户名">
      <el-input v-model=user.UserName />
    </el-form-item>
    <el-form-item label="用户密码" >
      <el-input v-model=user.Password  type="password"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">创建</el-button>
      <el-button>取消</el-button>
      <el-button type="text" @click="$router.back()">返回上一步</el-button>
    </el-form-item>
  </el-form>

</template>

<script>
import axios from "axios";

export default {
  name: "UserEdit",
  data: function() {
    return {
      user: {},
    }
  },

  methods: {
    onSubmit: function() {
      axios.post('/api/users',{
        ID: this.user.ID,
        UserName: this.user.UserName,
        Password: this.user.Password,
      }, {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        }
      } ).then((response) => {
        console.log(response)
        alert('添加成功')
        this.$router.go(-1)
      }, (response) => {
        console.error(response)
        alert('添加失败')
      })
    },
  }

}
</script>

<style scoped>
</style>