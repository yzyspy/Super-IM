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

</template>

<style scoped>

</style>