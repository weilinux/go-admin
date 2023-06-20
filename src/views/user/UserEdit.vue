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
      axios.put('/api/users/' + this.$route.params.id,{
        UserName: this.user.UserName,
        Password: this.user.Password,
      }, {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        }
      } ).then((response) => {
        console.log(response)
        alert('修改成功')
        this.$router.go(-1)
      }, (response) => {
        console.error(response)
        alert('修改失败')
      })
    },
  },

  mounted: function() {
    // http://localhost:9550/user/edit/2
    // 同样的,this.$route.query.id 如果这样访问的话users?id=2
    axios.get('/api/users/' + this.$route.params.id).then((response => {
      console.info(response.data)
      this.user = response.data.user
    }), (response) => {
      console.error(response)
    })

  }

}
</script>

<style scoped>
</style>