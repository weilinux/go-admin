<template>
  <el-tabs
      v-model="currentTabName"
      class="demo-tabs"
      @tab-click="handleTab">
    <el-tab-pane label="MySQL数据库" name="MySQL"></el-tab-pane>
    <el-tab-pane label="MongoDB数据库" name="MongoDB"></el-tab-pane>
    <el-tab-pane label="Redis数据库" name="Redis"></el-tab-pane>
    <div style="height:480px;">
      <component v-bind:is="currentTabName" @close="handleSubClose" :ID="userID"></component>
    </div>
  </el-tabs>
</template>


<style>
.demo-tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}
</style>

<script>
import MySQL from './MySQL.vue';
import MongoDB from './MongoDB.vue';
import Redis from './Redis.vue'
// import {reactive} from "vue";


export default {
  name: "Middleware",
  components: {
    MySQL,
    MongoDB,
    Redis,
  },
  data: function (){
    return {
      currentTabName: 'MySQL'
    }
  },
  methods: {
    handleTab: function(tab) {
      this.currentTabName = tab.name
    }
  }

}
</script>

