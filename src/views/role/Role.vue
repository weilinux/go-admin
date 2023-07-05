<template>
  <el-descriptions title="用户详细" :column="1" border>
    <el-descriptions-item
        label="用户ID"
        label-align="right"
        align="center"
        label-class-name="my-label"
        class="my-content"
        width="150px">{{ role.id }}</el-descriptions-item>
    <el-descriptions-item label="用户名" label-align="right" align="center">{{ role.RoleName }}</el-descriptions-item>
    <el-descriptions-item label="用户密码" label-align="right" align="center">{{ role.Password }}</el-descriptions-item>
  </el-descriptions>
  <el-button type="text" @click="$router.back()">返回上一步</el-button>
</template>

<script>
import RoleService from '@/services/role.service';

export default {
  data: function() {
    return {
      role: {}
    }
  },
  mounted: function() {
    RoleService.showRole(this.$route.params.id).then((response) => {
      this.role = response.data.data
    })
  },
}
</script>

<style scoped>
el-descriptions {
  margin-top: 20px;
}

.my-label {
  background: var(--el-color-success-light-9);
}
.my-content {
  background: var(--el-color-danger-light-9);
}

</style>