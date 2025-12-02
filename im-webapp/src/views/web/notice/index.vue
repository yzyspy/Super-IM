<script setup lang="ts">

import {onMounted} from "vue";
import {getFriendApplyList} from "@/api/user_api";
import { ArrowDown } from '@element-plus/icons-vue';

onMounted(() => {
  //获取我的好友申请验证列表
  getFriendApplyList().then(res => {
    console.log(res)
  })
})

function handleCommand(command: string) {
  console.log(command);
}

</script>
<template>
  <el-dropdown @click="handleCommand">
    <span class="el-dropdown-link">
      下拉菜单<el-icon class="el-icon--right"><arrow-down /></el-icon>
    </span>

    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item>选项一</el-dropdown-item>
        <el-dropdown-item divided>分割线</el-dropdown-item>
        <el-dropdown-item disabled>禁用选项</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<style scoped>
.el-dropdown-link {
  cursor: pointer;
  color: var(--el-color-primary);
  display: flex;
  align-items: center;
}
</style>