<template>
<!--  <div style="margin-bottom: 80px;" class="d-flex justify-content-center align-items-center container ">-->
  <div style="margin-bottom: 80px;" >
    <h3 align="center">主机授权</h3>
  </div>

  <div style="display: inline-block">
     <p style="margin-left: 10px; text-align: left;">用户</p>
    <el-select v-model="user" class="m-2" placeholder="请选择" @change="changeValue" size="large">
     <el-option
         v-for="item in userList"
         :key="item.value"
         :label="item.UserName"
         :value="item.id"
     />
     </el-select>
 </div>

  <div style="display: inline-block; margin-left: 20px">
    <p style="margin-left: 10px; text-align: left;">主机</p>
    <el-select v-model="hosts" class="m-2" multiple placeholder="请选择" size="large">
      <el-option
          v-for="item in hostList"
          :key="item.value"
          :label="item.HostName"
          :value="item.ID"
      />
    </el-select>
    </div>

  <div style="margin-top: 80px;" >
    <el-button type="primary" @click="onSubmit">授权</el-button>
    <el-button type="text" @click="$router.back()" style="align: left">返回上一步</el-button>

  </div>


  <div style="margin-bottom: 80px;" >
    <h3 align="center">主机解绑</h3>
  </div>

</template>

<script>
import axios from "axios";
import authHeader from "@/services/auth-header";

export default {
  name: "HostEdit",
  data: function() {
    return {
      hostList: [],
      userList: [],
      user: '',
      hosts: '',
    }
  },

  methods: {
    onSubmit: function() {
      let header = authHeader()
      header["Content-Type"]  = "application/json"
      axios.post('/api/hosts/assign', {
        "user_id": this.user,
        "host_id": this.hosts
      }).then((response) => {
        console.log(response)
      }, (response => {
        console.error(response)
      }))
    },

    changeValue: function() {
      axios.get('/api/users/' + this.user + '/hosts').then((response) => {
        console.log(response)
        this.hostList = response.data.rows
      }, (response) => {
        console.error(response)
      })
    }
  },

  mounted: function() {
    axios.get('/api/users').then((response) => {
          //成功后的callback
          console.log(response)
          //把远程返回的结果(JSON) 赋予到本地, JS支持JSON
          this.userList = response.data.rows
        }, (response) => {
          //失败后的callback
          console.error(response)
        }
    )

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