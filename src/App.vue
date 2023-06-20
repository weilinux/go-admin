<template>
  <el-container>
    <el-header height="40px" style="padding:0;margin:0">

      <div id="nav">
        <a href="/" class="logo">
          <span class="logo-lg"><b>Go</b>Admin</span>
        </a>
        <nav id="navbar" class="navbar navbar-expand-sm navbar-dark bg-dark">
          <div class="navbar-nav mr-auto">
            <li class="nav-item">
              <router-link to="/main" class="nav-link">
                <font-awesome-icon icon="home" /> 主页
              </router-link>
            </li>
            <li v-if="showAdminBoard" class="nav-item">
              <router-link to="/admin" class="nav-link">Admin Board</router-link>
            </li>
            <li v-if="showModeratorBoard" class="nav-item">
              <router-link to="/mod" class="nav-link">Moderator Board</router-link>
            </li>
            <li class="nav-item">
              <router-link v-if="currentUser" to="/user" class="nav-link">User</router-link>
            </li>
          </div>

          <div v-if="!currentUser" class="navbar-nav ml-auto">
            <li class="nav-item">
              <router-link to="/register" class="nav-link">
                <font-awesome-icon icon="user-plus" /> 注册
              </router-link>
            </li>
            <li class="nav-item">
              <router-link to="/login" class="nav-link">
                <font-awesome-icon icon="sign-in-alt" /> 登录
              </router-link>
            </li>
          </div>

          <div v-if="currentUser" class="navbar-nav ml-auto">
            <li class="nav-item">
              <router-link to="/profile" class="nav-link">
                <font-awesome-icon icon="user" />
                {{ currentUser.username }}
              </router-link>
            </li>
            <li class="nav-item">
              <a class="nav-link" @click.prevent="logOut">
                <font-awesome-icon icon="sign-out-alt" /> 退出
              </a>
            </li>
          </div>
        </nav>
      </div>
      <!--          <div id="header">页面主题</div>-->
    </el-header>

  </el-container>
    <router-view>

    </router-view>
</template>

<script>
export default {
  computed: {
    currentUser() {
      return this.$store.state.auth.user;
    },
    showAdminBoard() {
      if (this.currentUser && this.currentUser['roles']) {
        return this.currentUser['roles'].includes('ROLE_ADMIN');
      }

      return false;
    },
    showModeratorBoard() {
      if (this.currentUser && this.currentUser['roles']) {
        return this.currentUser['roles'].includes('ROLE_MODERATOR');
      }

      return false;
    }
  },
  methods: {
    logOut() {
      this.$store.dispatch('auth/logout');
      this.$router.push('/login');
    }
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 0px;
  height: 46px;
}
#navbar {
  padding: 0px;
  margin-left: 200px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}


.header {
  font-size: 30px;
  line-height: 50px;
  background-color: #f1f1f1;
}

.aside {
  background-color: wheat;
  height: 600px;
}

.subHeader {
  background-color:cornflowerblue;
}

.desc {
  font-size: 25px;
  line-height: 80px;
  color: white;
  width: 800px;
}

.content {
  height: 490px;
}

.footer {
  background-color:dimgrey;
  color: white;
  font-size: 17px;
  line-height: 30px;
}

body, html {
  height: 100%;
}

.logo {
  display: block;
  float: left;
  width: 200px;
  height: 40px;
  padding-top: 10px;
  padding-bottom: 10px;
  border-right-style: solid;
  border-right-width: 1px;
  background-color: #001529;

}

</style>
