<template>
  <h4>主机添加</h4>

  <div class="d-flex justify-content-center align-items-center container ">
    <el-form :model="form" label-width="120px" width="50%">
      <el-form-item label="主机名">
        <el-input v-model=host.HostName />
      </el-form-item>
      <el-form-item label="主机地址">
        <el-input v-model=host.HostIP />
      </el-form-item>
      <el-form-item label="主机端口">
        <el-input v-model=host.HostPort />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">创建</el-button>
        <el-button type="Danger" :icon="Delete" @click="onClear">清空</el-button>
        <el-button type="text" @click="$router.back()">返回上一步</el-button>
      </el-form-item>
    </el-form>
  </div>

</template>

<script>
import hostService from "@/services/host.service";

export default {
  name: "HostAdd",
  data: function() {
    return {
      host: {},
    }
  },

  methods: {
    onSubmit: function() {
      hostService.addHost(this.host).then((response) => {
        console.log(response)
      })

    },

    onClear: function() {
      this.host.HostName = ""
      this.host.HostIP = ""
      this.host.HostPort = ""
    }
  }


}
</script>

<style scoped>

</style>