<script setup lang="ts">
import { useUserInfoStore } from "@/stores/index";
import {type AuthLogoutRequest, doLogout} from "@/api/auth_api"
import {onMounted, reactive, ref} from "vue";
import { toRefs } from 'vue'
import { useRouter } from "vue-router";
import ImageUpload from '@/components/ImageUpload.vue';
import SvgIcon from "@/components/SvgIcon.vue";


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


const showEdit = ref(false)
function editNickName() {
  if (showEdit.value == true) {
    showEdit.value = false
  } else {
    showEdit.value = true
  }
}
</script>

<template>
  <div class="my-info">
    <el-form-item label="头像">
      <el-avatar :src="userInfo.avatar"></el-avatar>
      <ImageUpload />
    </el-form-item>
    <el-form-item label="用户号">
      <span>{{userInfo.uid}}</span>
    </el-form-item>
    <el-form-item label="昵称">
      <span v-if="!showEdit">{{userInfo.name}}</span>

        <el-input  v-else v-model="userInfo.name"></el-input>
        <svg-icon iconName="icon-bianji" @click="editNickName()"></svg-icon>


    </el-form-item>
    <el-form-item label="简介">
      <span>{{userInfo.abstract}}</span>
    </el-form-item>
    <a @click="logout">退出登录</a>
  </div>


</template>

<style scoped>
.my-info {
  padding: 10px;
}
</style>