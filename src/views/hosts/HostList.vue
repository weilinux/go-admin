<template>
  <div>
    {{ title }}
    <el-button class="mt-1" style="width: 10%;margin-left: 500px"  @click="add_host">添加</el-button>
  </div>

  <el-table :data="hosts" style="width: 100%;" height="460">
    <el-table-column type="index" label="序号" width="80" />
    <el-table-column prop="ID" label="编号" width="80" />
    <el-table-column prop="HostName" label="主机名" width="180" />
    <el-table-column prop="HostIP" label="主机地址" width="250" />
    <el-table-column prop="HostPort" label="主机端口" width="250" />
    <el-table-column fixed="right" label="操作" width="350">
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


  <hr class="my-2" />
  <div class="demo-pagination-block;" style="padding-left: 580px">
    <el-pagination
        class="pagination"
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        v-model:total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
    />
  </div>
  <!--        :background="background"-->

</template>

<script>
import axios from "@/router/carry";
import hostService from "@/services/host.service";
// import authHeader from "@/services/auth-header";

export default {
  name: "Hosts",
  data: function() {
    return {
      title: "主机列表",
      hosts: [],
      total: 1000,
      pageSize: 10,
      currentPage: 1,
    }
  },

  methods: {
    ssh_host: function(row) {
      axios.get('/api/hosts/' + row.ID + '/ssh', {
        headers: {
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
      hostService.delUserHost(row.ID).then(
          (response) => {
            console.log(response)
            this.hosts.splice(index, 1)
          }, (error) => {
            this.content =
                (error.response &&
                error.response.data &&
                error.response.data.message) ||
                error.message ||
                error.toString();
      })
    },

    get_user_hosts: function () {
      hostService.getUserHosts(this.currentPage, this.pageSize).then(
          (response) => {
            console.log(response)
            this.hosts = response.data.rows
            this.total = response.data.total
          }, (error) => {
            this.content =
                (error.response &&
                    error.response.data &&
                    error.response.data.message) ||
                error.message ||
                error.toString();
          })
    },

    handleSizeChange: function () {
      this.get_user_hosts()
    },
    handleCurrentChange: function () {
      this.get_user_hosts()
    }
  },

  mounted: function() {
    this.get_user_hosts()
  },
}
</script>

<style scoped>
td {
  border: 3px solid grey;
}

.demo-pagination-block + {
  margin-top: 10px;
}
.demo-pagination-block .pagination {
  margin-bottom: 16px;
}

</style>