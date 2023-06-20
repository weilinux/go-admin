<template>
  <div class="navbar">
    <nav class="nav-main navbar navbar-expand-lg navbar-dark bg-dark">
      <div class="container">
        <a class="navbar-brand" href="/">服务器管理</a>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#nav-main" aria-controls="nav-main" aria-expanded="false" aria-label="Toggle navigation"><span class="navbar-toggler-icon"></span></button>
        <div class="collapse navbar-collapse" id="nav-main">
          <ul class="navbar-nav me-auto">
            <li class="nav-item dropdown"><a class="nav-link dropdown-toggle" href="/hosts" id="navbarDropdown" role="button" data-bs-toggle="dropdown" aria-expanded="false">主机</a>
              <div class="dropdown-menu" aria-labelledby="navbarDropdown">
                <a class="dropdown-item" href="/hosts">列表</a>
                <div class="dropdown-divider"></div>
                <a class="dropdown-item" href="/hosts">配置</a>
              </div>
            </li>
            <li class="nav-item"><a class="nav-link" href="/monitor/hosts">指标</a></li>
            <li class="nav-item"><a class="nav-link" href="/monitor/hosts">中间件</a></li>
            <li class="nav-item"><a class="nav-link" href="/monitor/hosts">日志</a></li>
            <li class="nav-item"><a class="nav-link" href="/monitor/hosts">告警</a></li>
            <li class="nav-item dropdown"><a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-bs-toggle="dropdown" aria-expanded="false">用户</a>
              <div class="dropdown-menu" aria-labelledby="navbarDropdown">
                <a class="dropdown-item" href="/login">登录</a>
                <a class="dropdown-item" href="/logout">退出</a>
                <a class="dropdown-item" href="register">注册</a>
                <div class="dropdown-divider"></div>
                <a class="dropdown-item" href="/user">管理</a>
              </div>
            </li>
            <li class="nav-item"><a class="nav-link" href="/about">联系我们</a></li>
          </ul>
          <form class="d-flex" action="index.html">
            <div class="input-group">
              <input class="form-control" name="q" type="search" placeholder="Search" aria-label="Search" />
              <button class="input-group-btn btn btn-primary" type="submit">搜索</button>
            </div>
          </form>
        </div>
      </div>
    </nav>

    <h1>{{ title }}</h1>
  </div>


</template>

<script>
export default {
  name: "Nav",
  props: ['title'],
}
</script>

<style scoped>

</style>