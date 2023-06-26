<template>
  <el-table :data="tableData" style="width: 100%">
    <el-table-column prop="date" label="Date" width="180" />
    <el-table-column prop="name" label="Name" width="180" />
    <el-table-column prop="address" label="Address" />
  </el-table>
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
  name: "Database",
  data: function()  {
    return {
      tableData: [
        {
          date: '2016-05-03',
          name: 'Tom',
          address: 'this is MySQL',
        },
        {
          date: '2016-05-02',
          name: 'Tom',
          address: 'this is MySQL',
        },
        {
          date: '2016-05-04',
          name: 'Tom',
          address: 'this is MySQL',
        },
        {
          date: '2016-05-01',
          name: 'Tom',
          address: 'this is MySQL',
        },
      ]
    }

  }
}
</script>

