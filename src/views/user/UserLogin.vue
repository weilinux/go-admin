<template>
  <el-container>
    <el-header height="50px" style="padding:0">
      <div class="header">服务器管理系统</div>
    </el-header>

    <div id="login">
      <h2 align="center">{{ title }}</h2>
      <div v-if="noLogin"><label>账号：</label><input type="text" v-model="UserName" /></div>
      <div v-if="noLogin"><label>密码：</label><input type="password" v-model="Password" /></div>
      <p align="center"><button @click="click">{{ buttonTitle }}</button></p>
      <p align="center"><button @click="click2">注销</button></p>
      <p align="center">
        <router-link to="enroll">申请注册</router-link>
      </p>
      <el-button>I am ElButton</el-button>
    </div>

  </el-container>

</template>

<script>
import axios from 'axios'

export default {
  data: function() {
    return {
      title:"欢迎您：未登录",
      noLogin:true,
      UserName: '',
      Password: '',
      buttonTitle:"登录"
    }
  },
  methods: {
    click() {
      if (this.noLogin) {
        this.login()
      } else {
        this.logout()
      }
    },
    click2(){
      axios.get('api/logout',{
        headers: {
          'Authorization': 'Bearer ' + localStorage.getItem('Token')
        }
      }).then((response) =>{
        console.log(response)
      }).then((response) =>{
        console.error(response)
      })

    },
    login: function() {
      axios.post('/api/login',
          {
            UserName: this.UserName,
            Password: this.Password,
          },{
            headers: {
              'Content-Type': 'application/x-www-form-urlencoded'
            }
          }
      ).then((response) => {
        console.info("提交成功!, 刚才提交的内容是：" + response.data.UserName  +  response.data.Password + response.data.Token)
        localStorage.setItem("Token", response.data.Token);
  // axios.get('/main')
                this.$router.push({name: 'Main',})
              }, (response) => {
                alert('登录失败')
                console.error(response)
              }
          )
      if (this.UserName.length > 0 && this.Password.length > 0) {
        this.noLogin = false
        this.title = `欢迎您:${this.UserName}`
        this.buttonTitle = "注销"
        this.UserName = ""
        this.Password = ""
      } else {
        alert("请输入账号密码")
      }

    },
    logout: function () {
      this.noLogin = true
      this.title = `欢迎您:未登录`
      this.buttonTitle = "登录"
    }
  }


}
</script>

<style scoped>

</style>