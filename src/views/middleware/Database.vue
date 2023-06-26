<template>
  <el-tabs
      v-model="activeName"
      class="demo-tabs"
      @tab-click="handleClick">
    <el-tab-pane label="MySQL数据库" name="first">MySQL数据库</el-tab-pane>
    <el-tab-pane label="MongoDB数据库" name="second">MongoDB数据库</el-tab-pane>
    <el-tab-pane label="Redis数据库" name="third">Redis数据库</el-tab-pane>
  </el-tabs>
</template>


<!--<script lang="ts" setup>-->
<!--import { ref } from 'vue'-->
<!--import type { TabsPaneContext } from 'element-plus'-->

<!--const activeName = ref('first')-->

<!--const handleClick = (tab: TabsPaneContext, event: Event) => {-->
<!--  console.log(tab, event)-->
<!--}-->
<!--</script>-->

<style>
.demo-tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}
</style>

<script>
export default {
  name: "Database"
}
</script>

