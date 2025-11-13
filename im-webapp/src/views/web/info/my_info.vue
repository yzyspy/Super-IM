<script setup lang="ts">
import { useUserInfoStore } from "@/stores/index";
import {type AuthLogoutRequest, doLogout} from "@/api/auth_api"


const userStore = useUserInfoStore();

function logout() {

  const req : AuthLogoutRequest = {

  }
  let logoutRet = doLogout(req)
  console.log("退出登录", logoutRet);
}

</script>

<template>
  my_info
  <a @click="logout">退出登录</a>
</template>

<style scoped>

</style>