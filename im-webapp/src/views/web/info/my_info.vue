<script setup lang="ts">

function logout() {
  console.log("退出登录");
}

</script>

<template>
  my_info
  <el-button type="primary" @click="logout">退出登录</el-button>
</template>

<style scoped>

</style>