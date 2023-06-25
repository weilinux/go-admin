<template>
  <el-descriptions title="用户详细" :column="1" border>
    <el-descriptions-item
      label="用户ID"
      label-align="right"
      align="center"
      label-class-name="my-label"
      class="my-content"
      width="150px"
      >{{ user.id }}</el-descriptions-item
    >
    <el-descriptions-item label="用户名" label-align="right" align="center">{{
      user.UserName
    }}</el-descriptions-item>
    <el-descriptions-item label="用户密码" label-align="right" align="center">{{
      user.Password
    }}</el-descriptions-item>
  </el-descriptions>

  <el-button type="text" @click="$router.back()">返回上一步</el-button>
</template>

<script>
import axios from "axios";

export default {
  data: function () {
    return {
      user: {},
    };
  },
  mounted: function () {
    // http://localhost:9550/user/edit/2
    // 同样的,this.$route.query.id 如果这样访问的话users?id=2
    axios.get("/api/users/" + this.$route.params.id).then(
      (response) => {
        console.info(response.data);
        this.user = response.data.user;
      },
      (response) => {
        console.error(response);
      }
    );
  },
  components: {},
};
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
