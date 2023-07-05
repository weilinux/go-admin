<template>
<div>
    {{ title }}
</div>
  <!-- 功能操作区一 -->
  <div class="flex">
    <el-form :model="searchForm" inline>
      <el-form-item label="主机名称">
        <el-input v-model="roleName" placeholder="请输入角色名称" style="width: 150px;alignment: left" class="filter-item"/>
      </el-form-item>
      <el-form-item>
        <el-button @click="handleFilter">查询</el-button>
      </el-form-item>
      <el-form-item>
        <el-button @click="add_role">+添加角色</el-button>
      </el-form-item>
      <el-form-item>
        <el-button @click="delete_role">批量删除</el-button>
      </el-form-item>
    </el-form>
  </div>

  <!-- TABLE渲染区: 列要对齐从右边开始对齐,还有功能操作大小 -->
  <el-table :data="roles" style="width: 100%" height="250">
    <el-table-column prop="id" label="ID" width="220" />
    <el-table-column prop="name" label="角色名称" width="220" />
    <el-table-column prop="status" label="状态" width="220" />
    <el-table-column prop="sort" label="排序" width="100" />
    <el-table-column prop="created_time" label="添加时间" width="210" />
    <el-table-column prop="updated_time" label="添加时间" width="210" />
    <!-- 操作功能区二 -->
    <el-table-column fixed="right" label="功能操作" width="450">
      <template #default="scope">
        <el-button
            link
            type="primary"
            size="small"
            @click="edit_role(scope.$index, scope.row)">编辑</el-button>
        <el-button
            link
            type="primary"
            size="small"
            @click="delete_role(scope.$index, scope.row)">删除</el-button>
        <el-button
            link
            type="primary"
            size="small"
            @click="show_role(scope.$index, scope.row)">角色权限</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import RoleService from '@/services/role.service';
import { ElMessageBox } from 'element-plus'

export default {
  data: function() {
    return {
      title: "角色列表",
      roles: [],
      roleName: '',
    }
  },
  methods: {
    edit_role: function(index, row) {
      this.$router.push({name: 'RoleEdit', params: {id: row.id}})
    },
    add_role: function() {
      this.$router.push({name: 'RoleNew'})
    },
    delete_role: function(index, row) {
      ElMessageBox.confirm("确实要删除这个角色吗?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
      }).then(() => {
        RoleService.delRole(row.id).then(
            () => {
              this.roles.splice(index, 1)
            })
        alert("删除成功!");
      })
    },
    get_roles: function() {
      RoleService.getRoleList().then(response => {
        this.roles = response.data.data.rows;
      }, (error) => {
        console.log(error)
      })
    }
  },
  created: function() {
    this.get_roles()
  }
}
</script>

<style scoped>
td {
  border: 3px solid grey;
}
.flex {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>