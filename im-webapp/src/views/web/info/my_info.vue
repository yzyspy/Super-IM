<script setup lang="ts">
import { useUserInfoStore } from "@/stores/index";
import {type AuthLogoutRequest, doLogout} from "@/api/auth_api"


const userStore = useUserInfoStore();

function logout() {

  const req : AuthLogoutRequest = {

  }
  let logoutRet = doLogout(req)
  userStore.clearUserInfo()
  console.log("退出登录", logoutRet);
}

</script>

<template>
  <div>
    <div>头像</div>
    <div>用户号</div>
    <div>昵称</div>
    <div>简介</div>
  </div>
  <a @click="logout">退出登录</a>
</template>

<style scoped>

</style>