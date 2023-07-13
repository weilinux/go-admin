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
      <el-form-item label="主机地址">
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
import hostService from "@/services/host.service";

export default {
  name: "HostEdit",
  data: function () {
    return {
      host: {}
    }
  },

  methods: {
    onSubmit: function () {
      hostService.updateHost(this.$route.params.id, this.host).then((response) => {
        console.log(response)
        this.$router.go(-1)
      })
    }
  },

  //TODO: 这里不应该再次发出请求的了，应该使用客户端存在的数据进行处理
  mounted: function () {
    hostService.getHost(this.$route.params.id).then((response) => {
      this.host = response.data.data
      console.log(response)
    })
  }
}
</script>

<style scoped>

</style>