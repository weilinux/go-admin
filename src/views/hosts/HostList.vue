<template>
  <div>
    {{ title }}
  </div>

  <el-table :data="hosts" style="width: 100%" height="250">
    <el-table-column type="index" label="序号" width="120" />
    <el-table-column prop="ID" label="编号" width="120" />
    <el-table-column prop="HostName" label="主机名" width="120" />
    <el-table-column prop="HostIP" label="主机地址" width="380" />
    <el-table-column prop="HostPort" label="主机端口" width="250" />
    <el-table-column fixed="right" label="操作" width="250">
      <template #default="scope">
        <el-button
            link
            type="primary"
            size="small"
            @click="ssh_host(scope.row)">SSH</el-button>
        <el-button
            link
            type="primary"
            size="small"
            @click="show_host(scope.$index, scope.row)">FTP</el-button>
        <el-button
            link
            type="primary"
            size="small"
            @click="show_host(scope.$index, scope.row)">详细</el-button>
        <el-button
            link
            type="primary"
            size="small"
            @click="edit_host(scope.$index, scope.row)">编辑</el-button>
        <el-button
            link
            type="primary"
            size="small"
            @click="delete_host(scope.$index, scope.row)">删除</el-button>

      </template>
    </el-table-column>
  </el-table>
  <el-button class="mt-4" style="width: 10%;align: right"  @click="add_host">添加</el-button>
</template>

<script>
import axios from 'axios'

export default {
  name: "Hosts",
  data: function() {
    return {
      title: "主机列表",
      hosts: []
    }
  },

  methods: {
    ssh_host: function(row) {
      axios.get('/api/hosts/' + row.ID + '/ssh', {
        headers: {
          'Authorization': 'Bearer ' + localStorage.getItem('Token'),
          'Access-Control-Allow-Origin': 'http://localhost:8084'

          }
      })
    },
    show_host: function(index, row) {
      this.$router.push({name: 'Host', params: {id: row.ID}})
    },
    edit_host: function(index, row) {
      this.$router.push({name: 'HostEdit', params: {id: row.ID}})
    },
    add_host: function() {
      this.$router.push({name: 'HostNew'})
    },
    delete_host: function(index, row) {
      axios.delete("/api/hosts/" + row.ID, {
        headers: {
          'Authorization': 'Bearer ' + localStorage.getItem('Token')
        }
      }).then(response => {
        this.hosts.splice(index, 1)
        console.log(response)
      }, (response) => {
        console.error(response)
      })
    }
  },

  mounted: function() {
    axios.get('/api/users/hosts', {
      headers: {
        'Authorization': 'Bearer ' + localStorage.getItem('Token')
      }
    }).then((response) => {
      console.log(response)
      this.hosts = response.data.rows
    }, (response) => {
      console.error(response)
    })
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