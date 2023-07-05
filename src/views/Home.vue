<template>
<!--  <div class="home">-->
<!--    <img alt="Vue logo" src="../assets/logo.png" />-->
<!--    <HelloWorld msg="Welcome to Your Vue.js App" />-->
<!--  </div>-->
  <el-container id="middle-area">
    <el-container width="200px">
      <el-aside width="200px" style="alignment:center; height: 100%; padding-top: 0px;">
        <el-menu id="el-menu"
            mode="vertical"
            class="aside"
            router
            @select="selectItem"
            :default-active="$route.path"
            :unique-opened="true">
          <el-menu-item index="/main">概览</el-menu-item>

          <el-menu-item index="/app">软件仓库</el-menu-item>
          <el-sub-menu index="3">
            <template #title>主机管理</template>
            <el-menu-item index="/hosts">列表</el-menu-item>
            <el-menu-item index="/host/auth">授权</el-menu-item>
            <el-menu-item index="/host/monitor">监控</el-menu-item>
            <el-menu-item index="/host/warn">告警</el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="4">
            <template #title>中间件</template>
            <el-menu-item index="/middle">数据库</el-menu-item>
<!--            <el-menu-item index="4-2">Redis数据库</el-menu-item>-->
            <el-menu-item index="4-3">RabbitMQ</el-menu-item>
<!--            <el-menu-item index="4-4">Mongo数据库</el-menu-item>-->
          </el-sub-menu>
          <el-menu-item index="/container">容器管理</el-menu-item>
          <el-menu-item index="/schedular">计划任务</el-menu-item>

          <el-menu-item index="/log">日志审计</el-menu-item>
          <el-sub-menu index="8">
            <template #title>权限管理</template>
            <el-menu-item index="/users">用户列表</el-menu-item>
            <el-menu-item index="/roles">角色管理</el-menu-item>
            <el-menu-item index="/register">菜单管理</el-menu-item>
          </el-sub-menu>
        </el-menu>
      </el-aside>
    </el-container>

      <el-container>
<!--        <el-header height="40px" style="padding:0;margin:0">-->
<!--&lt;!&ndash;          <div id="header">页面主题</div>&ndash;&gt;-->
<!--        </el-header>-->
        <el-main id="el-main" style="padding:0; min-height: 450px;">
          <div class="content">
            <router-view>
            </router-view>
          </div>
        </el-main>
      </el-container>
  </el-container>


  <el-container>
    <el-footer height="30px" class="footer" style="padding:0;background-color: #36332D">
      <div>@horse设计</div>
    </el-footer>

  </el-container>
</template>

<script>

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


#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}


.header {
  font-size: 30px;
  line-height: 30px;
  background-color: #f1f1f1;
}

.aside {
  background-color: #001529;
  height: 100%;
}

#el-menu {
  background-color: #001529;
}

#middle-area {
  height: 100%;
  height: -webkit-calc(100% - 80px);
  height: -moz-calc(100% - 80px);
  height: calc(100% - 80px);
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
}

.footer {
  background-color:dimgrey;
  color: white;
  font-size: 17px;
  line-height: 30px;
  margin-left: 200px;
  border-top-style: solid;
  border-top-width: 1px;
}
</style>
