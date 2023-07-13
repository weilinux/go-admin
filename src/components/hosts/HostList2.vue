<template>
  <div style="margin-top: 40px;">
    {{ title }}
  </div>

  <div class="flex">
    <el-form :model="searchForm" inline>
      <el-form-item label="主机名称">
        <el-input v-model="host" placeholder="输入主机名称查询"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button @click="handleFilter">查询</el-button>
      </el-form-item>
    </el-form>
    <el-button @click="add_host">新增主机</el-button>
  </div>

  <el-table :data="hosts" style="width: 100%;" height="460">
    <el-table-column type="index" label="序号" width="80"/>
    <el-table-column prop="ID" label="编号" width="80"/>
    <el-table-column prop="HostName" label="主机名" width="180"/>
    <el-table-column prop="HostIP" label="主机地址" width="250"/>
    <el-table-column prop="HostPort" label="主机端口" width="250"/>
    <el-table-column fixed="right" label="功能操作" width="350">
      <template #default="scope">
        <el-button
            link
            type="primary"
            size="small"
            @click="ssh_host(scope.row)">SSH
        </el-button>
        <el-button
            link
            type="primary"
            size="small"
            @click="show_host(scope.$index, scope.row)">FTP
        </el-button>
        <el-button
            link
            type="primary"
            size="small"
            @click="show_host(scope.$index, scope.row)">详细
        </el-button>
        <el-button
            link
            type="primary"
            icon="Edit"
            size="small"
            @click="edit_host(scope.$index, scope.row)">编辑
        </el-button>
        <el-button
            link
            type="primary"
            icon="Delete"
            size="small"
            @click="delete_host(scope.$index, scope.row)">删除
        </el-button>
      </template>
    </el-table-column>
  </el-table>

  <hr class="my-2"/>
  <div class="demo-pagination-block;" style="padding-left: 580px">
    <el-pagination
        class="pagination"
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        v-model:total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"/>
  </div>
</template>

<script>
import axios from "@/router/carry";
import hostService from "@/services/host.service";
import {ElMessageBox} from 'element-plus'

export default {
  name: "Hosts",
  data: function () {
    return {
      title: "主机列表",
      host: "",
      hosts: [],
      total: 1000,
      pageSize: 10,
      currentPage: 1,
    }
  },

  methods: {
    ssh_host: function (row) {
      axios.get('/api/hosts/' + row.ID + '/ssh', {
        headers: {
          'Access-Control-Allow-Origin': 'http://www.wllinux.com:8084'
        }
      })
    },
    show_host: function (index, row) {
      this.$router.push({name: 'Host', params: {id: row.ID}})
    },
    edit_host: function (index, row) {
      this.$router.push({name: 'HostEdit', params: {id: row.ID}})
    },
    add_host: function () {
      this.$router.push({name: 'HostNew'})
    },
    delete_host: function (index, row) {
      ElMessageBox.confirm("确实要删除这个主机吗?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
      }).then(() => {
        hostService.delUserHost(row.ID).then(
            () => {
              this.hosts.splice(index, 1)
            })
        alert("删除成功!");
      })
    },

    get_user_hosts: function () {
      hostService.getUserHosts(this.currentPage, this.pageSize).then(
          (response) => {
            this.hosts = response.data.data.rows
            this.total = response.data.data.total
          })
    },

    handleSizeChange: function () {
      this.get_user_hosts()
    },
    handleCurrentChange: function () {
      this.get_user_hosts()
    }
  },

  mounted: function () {
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

.flex {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>