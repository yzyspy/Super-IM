<script setup lang="ts">

import {onMounted} from "vue";
import {getFriendApplyList} from "@/api/user_api";

onMounted(() => {
  //获取我的好友申请验证列表
  getFriendApplyList().then(res => {
    console.log(res)
  })
})
</script>

<template>
<div>
  <el-button-group>
    <el-button type="primary">通知</el-button>
    <el-button type="primary">私信</el-button>
    <el-dropdown>
      <el-button type="primary">通知</el-button>
      <el-button type="primary">私信</el-button>
      <el-dropdown-menu>
        <el-dropdown-item>通知</el-dropdown-item>
        <el-dropdown-item>私信</el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>
  </el-button-group>
</div>

</template>

<style scoped>

</style>