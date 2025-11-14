<script setup lang="ts">
import { useUserInfoStore } from "@/stores/index";
import {type AuthLogoutRequest, doLogout} from "@/api/auth_api"
import {onMounted, reactive} from "vue";
import { toRefs } from 'vue'
import { useRouter } from "vue-router";
import ImageUpload from '@/components/ImageUpload.vue';


const userStore = useUserInfoStore();
const router = useRouter();

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

async function logout() {
  let logoutRet = await doLogout({})

  userStore.clearUserInfo()
  console.log("退出登录", logoutRet);
  router.push({ name: 'login' })

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
  <ImageUpload />

</template>

<style scoped>

</style>