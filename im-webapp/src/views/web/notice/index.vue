<script setup lang="ts">

import {onMounted, ref} from "vue";
import {type FriendApplyUserInfo, getFriendApplyList, type UserInfo} from "@/api/user_api";
import { ArrowDown } from '@element-plus/icons-vue';


const applyList = ref<FriendApplyUserInfo[]>([])

onMounted(() => {
  //获取我的好友申请验证列表
  getFriendApplyList().then(res => {
    applyList.value = res.list
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
      <el-menu-item index="1-3" v-for="(item, index) in applyList" :key="index">
        {{item.nickname}}
        <el-dropdown @click="handleCommand">
    <span class="el-dropdown-link">
      同意<el-icon class="el-icon--right"><arrow-down /></el-icon>
    </span>

          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>同意</el-dropdown-item>
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