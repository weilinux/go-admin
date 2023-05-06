<template>
<!--  <div id="nav">-->
<!--    <router-link to="/">Home</router-link> |-->
<!--    <router-link to="/about">About2</router-link>-->
<!--    <Nav></Nav>-->
<!--  </div>-->
  <router-view>

  </router-view>

</template>

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
