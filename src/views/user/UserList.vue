<!--suppress ALL -->
<template>
<div>
    {{ title }}
  </div>
<!--  <el-table :data="users" style="width: 100%" height="250" :table-layout="tableLayout">-->
  <el-table :data="users" style="width: 100%" height="250">
    <el-table-column type="index" label="序号" width="120" />
    <el-table-column prop="ID" label="用户编号" width="120" />
    <el-table-column prop="UserName" label="名称" width="120" />
    <el-table-column prop="Password" label="密钥" width="520" />
    <el-table-column prop="CreatedAt" label="创建时间" width="250" />
    <el-table-column fixed="right" label="操作" width="150">
      <template #default="scope">
        <el-button
            link
            type="primary"
            size="small"
            @click="show_user(scope.$index, scope.row)">详细</el-button>
        <el-button
            link
            type="primary"
            size="small"
            @click="edit_user(scope.$index, scope.row)">编辑</el-button>
        <el-button
            link
            type="primary"
            size="small"
            @click="delete_user(scope.$index, scope.row)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
  <el-button class="mt-4" style="width: 10%;align: right"  @click="add_user">添加</el-button>

</template>



<script>
import axios from 'axios'

export default {
  // data() 钩子方法: 声明页面会出现的变量，并且赋予初始值
  data: function() {
    return {
      title: "用户列表",
      users: [
      ]
    }
  },
  // 比较核心的方法，Vue页面中用到的事件都要写在这里
  methods: {
    // change_title: function() {
    //   this.title = "用户列表(修改过2)"
    // },
    show_user: function(index, row) {
      this.$router.push({name: 'User', params: {id: row.ID}})
    },
    edit_user: function(index, row) {
      this.$router.push({name: 'UserEdit', params: {id: row.ID}})
    },
    add_user: function() {
      this.$router.push({name: 'UserNew'})
    },
    delete_user: function(index, row) {
      axios.delete("/api/users/" + row.ID).then(response => {
        this.users.splice(index, 1)
        console.log(response)
      }, (response) => {
        console.error(response)
      })
    }
  },
  // 表示页面被Vue渲染好之后的钩子方法，会立刻执行, 钩子方法. http request写这里，还有created，相似
  mounted: function() {
    // 注意这是一个模拟数据的方法
    // import Mock from '../../mock/Mock'
    // this.tradeList = Mock.getTradeList()
    axios.get('/api/users').then((response) => {
      //成功后的callback
        console.log(response)
        //把远程返回的结果(JSON) 赋予到本地, JS支持JSON
        this.users = response.data.rows
        }, (response) => {
      //失败后的callback
        console.error(response)
        }
    )
  },
  components: {
  }
}
</script>

<style scoped>
td {
  border: 3px solid grey;
}
</style>