<template>
  <div v-if="!isreg">
    <div class="input-box">
      <p>用户</p>
      <input type="text" placeholder="请输入用户名" v-model="registered.username" />
      <p v-if="error_message.username" class="error">{{error_message.username}}</p>
    </div>

    <div class="input-box">
      <p>密码</p>
      <input type="password" placeholder="输入密码" v-model="registered.password" />
      <p v-if="error_message.password" class="error">{{error_message.password}}</p>
    </div>

    <div class="input-box">
      <p>邮箱</p>
      <input type="text" placeholder="输入邮箱" v-model="registered.email" />
      <p v-if="error_message.email" class="error">{{error_message.email}}</p>
    </div>

    <div class="input-box">
      <p>年龄</p>
      <input type="number" placeholder="输入年龄" v-model="registered.age" />
    </div>

    <div class="input-box">
      <input type="radio" id="boy" value="boy" v-model="registered.sex" />
      <label for="boy">boy</label>
      <input type="radio" id="girl" value="girl" v-model="registered.sex" />
      <label for="girl">girl</label>
    </div>

    <div class="input-box">
      <input type="checkbox" id="checkbox" v-model="registered.terms" />
      <label for="checkbox">我已阅读使用者条款</label>
    </div>

    <a class="btn" @click="submit">注册</a>
  </div>

  <div v-if="isreg">
    <h1>注册成功</h1>
  </div>


</template>

<script>
import UserService from '@/services/user.service';
import { ref, reactive } from "vue";
// import { ref } from "vue";

export default {
  name: "RegisterForm",
  setup() {
    const isreg = ref(false)

    const registered = reactive({
      username: "",
      password: "",
      sex: "",
      email: "",
      age: "",
      terms: false,
    })

    const error_message = reactive({})

    const handleError = (err) => {
      Object.keys(error_message).forEach((item) => {
        error_message[item] = ""
      })
      console.log(error_message)
      console.log(Object.keys(err))

      Object.keys(err).forEach((error) => {
        error_message[error] = err[error]
        console.log(error)
        console.log(error_message)
      })
    }

    const submit = () => {
      UserService.signupUser(registered.username, registered.password)
          .then(() => {
            isreg.value = true
            this.$router.push("/login")
          })
          .catch(err =>{
            // console.log("hehe")
            console.log(err.data.data.msg)
            // console.log("hehe")
            // handleError(err.data.msg)
          })
    }

    return {
      isreg,
      registered,
      error_message,
      handleError,
      submit,
    }
  }
}
</script>

<style scoped>

</style>