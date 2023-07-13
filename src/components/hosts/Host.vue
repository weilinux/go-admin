<template>
  <div class="d-flex justify-content-center align-items-center container ">
    <el-descriptions title="主机详细" :column="1" border class="margin-top">
      <el-descriptions-item
          label="主机ID"
          label-align="left"
          align="left"
          label-class-name="my-label"
          class="cell-item"
          width="20px">{{ host.ID }}
      </el-descriptions-item>
      <el-descriptions-item label="主机名" label-align="left" align="left">{{ host.HostName }}</el-descriptions-item>
      <el-descriptions-item label="主机地址" label-align="left" align="left">{{ host.HostIP }}</el-descriptions-item>
      <el-descriptions-item label="主机端口" label-align="left" align="left">{{ host.HostPort }}</el-descriptions-item>
    </el-descriptions>
  </div>

  <el-button type="text" @click="$router.back()" style="align: left">返回上一步</el-button>


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

  mounted: function () {
    hostService.getHost(this.$route.params.id).then((response) => {
      this.host = response.data.data
      console.log(response)
    })
  }
}
</script>

<style scoped>
.el-descriptions {
  margin-top: 20px;
}

.cell-item {
  display: flex;
  align-items: center;
}

.margin-top {
  margin-top: 20px;
}
</style>