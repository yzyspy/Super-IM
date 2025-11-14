<script setup lang="ts">
import { useUserInfoStore } from "@/stores/index";
import {type AuthLogoutRequest, doLogout} from "@/api/auth_api"
import {onMounted, reactive} from "vue";
import { toRefs } from 'vue'

const userStore = useUserInfoStore();

const userInfo = reactive({
  avatar: '',
  uid: 0,
  name: '',
  abstract: ''
})

onMounted(() => {
 // const user = toRefs(userStore.userInfo);
  userInfo.name = userStore.userInfo.user_name
  userInfo.avatar = userStore.userInfo.avatar
  userInfo.uid = userStore.userInfo.user_id
  userInfo.abstract = userStore.userInfo.abstract
});

function logout() {
  let logoutRet = doLogout({})
  userStore.clearUserInfo()
  console.log("退出登录", logoutRet);
}
</script>

<template>
  <div>
    <div>头像:{{userInfo.avatar}}</div>
    <div>用户号:{{userInfo.uid}}</div>
    <div>昵称:{{userInfo.name}}</div>
    <div>简介:{{userInfo.abstract}}</div>
  </div>
  <a @click="logout">退出登录</a>
</template>

<style scoped>

</style>