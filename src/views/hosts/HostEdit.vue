<template>
  <h4>主机编辑</h4>

  <div class="d-flex justify-content-center align-items-center container ">
    <el-form :model="form" label-width="120px">
      <el-form-item label="主机名">
        <el-input v-model=host.HostName />
      </el-form-item>
      <el-form-item label="主机IP">
        <el-input v-model=host.HostIP />
      </el-form-item>
      <el-form-item label="主机地址" >
        <el-input v-model=host.HostPort />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">更新</el-button>
        <el-button>取消</el-button>
        <el-button type="text" @click="$router.back()">返回上一步</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "HostEdit",
  data: function() {
    return {
      host: {}
    }
  },

  methods: {
    onSubmit: function() {
      axios.put('/api/hosts/' + this.$route.params.id,{
        HostName: this.host.HostName,
        HostIP: this.host.HostIP,
        HostPort: this.host.HostPort
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


    }


  },

  mounted: function() {
    // http://localhost:9550/host/edit/2
    // 同样的,this.$route.query.id 如果这样访问的话hosts?id=2
    axios.get('/api/hosts/' + this.$route.params.id).then((response => {
      console.info(response.data)
      this.host = response.data.host
    }), (response) => {
      console.error(response)
    })

  }

}
</script>

<style scoped>

</style>