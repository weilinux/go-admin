<template>
  <h4>用户编辑</h4>
  <el-form :model="form" label-width="120px">
    <el-form-item label="用户名">
      <el-input v-model=user.UserName />
    </el-form-item>
    <el-form-item label="用户密码" >
      <el-input v-model=user.Password  type="password"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">更新</el-button>
      <el-button>取消</el-button>
      <el-button type="text" @click="$router.back()">返回上一步</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import UserService from "@/services/user.service";

export default {
  name: "UserEdit",
  data: function() {
    return {
      user: {},
    }
  },

  methods: {
    onSubmit: function() {
      UserService.updateUser(this.user.id, this.user.UserName, this.user.Password).then(
          () => {
            alert('更新用户成功')
            this.$router.go(-1)
          })
    },
  },

  created: function() {
    UserService.showUser(this.$route.params.id).then((response) => {
      this.user = response.data.data
    })
  }
}
</script>

<style scoped>
</style>