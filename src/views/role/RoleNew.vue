<template>
  <h4>创建用户</h4>
  <el-form :model="form" label-width="120px" width="50%">
    <el-form-item label="用户ID">
      <el-input v-model=user.id />
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
import RoleService from '@/services/role.service';

export default {
  name: "UserEdit",
  data: function() {
    return {
      user: {},
    }
    },

  methods: {
    onSubmit: function() {
      RoleService.addRole(this.user.id, this.user.UserName, this.user.Password).then(
          (response) => {
            console.log(response)
            alert('添加用户成功')
          }
      )
    }
  }

}
</script>

<style scoped>
</style>