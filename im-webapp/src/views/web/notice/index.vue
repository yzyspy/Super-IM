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
  <el-menu>
    <el-sub-menu index="1">
      <template #title>
        <span>我的好友申请</span>
      </template>
      <el-menu-item index="1-3">
        <el-dropdown @click="handleCommand">
    <span class="el-dropdown-link">
      同意<el-icon class="el-icon--right"><arrow-down /></el-icon>
    </span>

          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>拒绝</el-dropdown-item>
              <el-dropdown-item>忽略</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-menu-item>
      <el-menu-item index="1-3">
        <el-dropdown @click="handleCommand">
    <span class="el-dropdown-link">
      同意<el-icon class="el-icon--right"><arrow-down /></el-icon>
    </span>

          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>拒绝</el-dropdown-item>
              <el-dropdown-item>忽略</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-menu-item>
    </el-sub-menu>
  </el-menu>

</template>

<style scoped>
.el-dropdown-link {
  cursor: pointer;
  color: var(--el-color-primary);
  display: flex;
  align-items: center;
}
</style>