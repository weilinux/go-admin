<template>
<!--  <div class="home">-->
<!--    <img alt="Vue logo" src="../assets/logo.png" />-->
<!--    <HelloWorld msg="Welcome to Your Vue.js App" />-->
<!--  </div>-->
  <el-container>
    <el-header height="50px" style="padding:0">
      <div class="header">服务器管理系统</div>
    </el-header>
    <el-container width="200px">
      <el-aside width="200px" style="align:center">
        <el-menu
            mode="vertical"
            class="aside"
            router
            @select="selectItem"
            :default-active="$route.path"
            :unique-opened="true">
          <el-menu-item index="/main">首页</el-menu-item>

          <el-sub-menu index="3">
            <template #title>主机管理</template>
            <el-menu-item index="/hosts">列表</el-menu-item>
            <el-menu-item index="/host/auth">授权</el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="4">
            <template #title>中间件</template>
            <el-menu-item index="4-1">mysql数据库</el-menu-item>
            <el-menu-item index="4-2">Redis数据库</el-menu-item>
            <el-menu-item index="4-3">RabbitMQ</el-menu-item>
            <el-menu-item index="4-4">Mongo数据库</el-menu-item>
          </el-sub-menu>
          <el-menu-item index="4">监控指标</el-menu-item>

          <el-menu-item index="/log">日志</el-menu-item>
          <el-menu-item index="/log">运维工具</el-menu-item>
          <el-menu-item index="7">告警</el-menu-item>
          <el-sub-menu index="8">
            <template #title>用户</template>
            <el-menu-item index="/login">登录</el-menu-item>
            <el-menu-item index="/logout">退出</el-menu-item>
            <el-menu-item index="/signup">注册</el-menu-item>
            <el-menu-item index="/users">管理</el-menu-item>
          </el-sub-menu>
          <el-menu-item index="/about">联系我们</el-menu-item>
        </el-menu>
      </el-aside>

      <el-container>
        <el-header height="50px" style="padding:0;margin:0">
<!--          <div id="header">页面主题</div>-->
        </el-header>
        <el-main style="padding:0;margin:0">
          <div class="content">
            <router-view>
            </router-view>
          </div>
        </el-main>
        <el-footer height="50px" class="footer" style="padding:0">
          <div>@horse设计</div>
        </el-footer>
      </el-container>
    </el-container>
  </el-container>
</template>

<script>
// @ is an alias to /src
// @ is an alias to /src
// @ is an alias to /src

export default {
  name: "Home",
  components: {
  },
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
  padding: 30px;
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
  height: 410px;
}

.footer {
  background-color:dimgrey;
  color: white;
  font-size: 17px;
  line-height: 30px;
}
</style>
