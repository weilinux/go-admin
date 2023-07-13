<template>
  <div>
    <div>
      <div class="subTitle">加入我们，一起创造美好世界</div>
      <h1 class="title">注册账号</h1>
      <div v-for="(item, index) in fields" v-bind:key="item" class="inputContainer">
        <div class="field">{{ item.title }} <span v-if="item.required" style="color: red;">*</span></div>
        <input v-model="item.model" class="input" :type="item.type"/>
        <div class="tip" v-if="index === 2">请确认密码程度需要大于1位</div>
      </div>
      <button @click="createAccount" class="btn">注册</button>
    </div>
  </div>
</template>

<script>
import UserService from '@/services/user.service';

export default {
  name: "RegisterForm",
  data: function () {
    return {
      fields: [
        {
          title: "用户名",
          required: true,
          type: "text",
          model: ""
        }, {
          title: "邮箱地址",
          required: false,
          type: "text",
          model: ""
        }, {
          title: "密码",
          required: true,
          type: "password",
          model: ""
        }
      ],
      receiveMsg: false
    }
  },
  computed: {
    name: {
      get() {
        return this.fields[0].model
      },
      set(value) {
        this.fields[0].model = value
      }
    },
    email: {
      get() {
        return this.fields[1].model
      },
      set(value) {
        this.fields[1].model = value
      }
    },
    password: {
      get() {
        return this.fields[2].model
      },
      set(value) {
        this.fields[2].model = value
      }
    }
  },
  methods: {
    emailCheck() {
      // var verify = /^\w[-\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\.)+[A-Za-z]{2,14}/;
      let verify = /^\w[-\w.+]*@([A-Za-z\d][-A-Za-z\d]+\.)+[A-Za-z]{2,14}/;
      if (!verify.test(this.email)) {
        return false
      } else {
        return true
      }
    },
    createAccount() {
      if (this.name.length === 0) {
        alert("请输入用户名")
        return
      } else if (this.password.length <= 1) {
        alert("密码设置需要大于1位字符")
        return
      } else if (this.email.length > 0 && !this.emailCheck(this.email)) {
        alert("请输入正确的邮箱")
        return
      }

      UserService.signupUser(this.name, this.password).then(() => {
        alert("注册成功")
        this.$router.push("/login")
      }, (error) => {
        console.error(error)
      })
    }
  }
}
</script>

<style scoped>

</style>