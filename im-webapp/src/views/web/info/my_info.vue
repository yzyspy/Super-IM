<script setup lang="ts">
import { useUserInfoStore } from "@/stores/index";
import {type AuthLogoutRequest, doLogout} from "@/api/auth_api"
import {onMounted, reactive, ref} from "vue";
import { toRefs } from 'vue'
import { useRouter } from "vue-router";
import ImageUpload from '@/components/ImageUpload.vue';
import SvgIcon from "@/components/SvgIcon.vue";
import {ElMessage} from "element-plus";
import {updateUserInfo, type UpdateUserInfoRequest} from "@/api/user_api";


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


const showEditName = ref(false)
function editNickName() {
  if (showEditName.value == true) {
    showEditName.value = false
  } else {
    showEditName.value = true
  }
}

const showEditAbstract = ref(false)
function editAbstract() {
  if (showEditAbstract.value == true) {
    showEditAbstract.value = false
  } else {
    showEditAbstract.value = true
  }
}

async function onUploadSuccess(url : string) {

  userInfo.avatar = url

  var userId = userStore.userInfo.user_id;

  let request : UpdateUserInfoRequest = {
    uid: userId,
    avatar: url,
  }
  let ret = await updateUserInfo(request)
  //调用接口更新用户头像 和 useStore
  ElMessage.info("上传成功" + ret.data)
}

async function onEditNameBlur() {
  console.log("onEditNameBlur" )
  ElMessage.info("onEditNameBlur" )

  var userId = userStore.userInfo.user_id;

  let request : UpdateUserInfoRequest = {
    uid: userId,
    nickname: url,
  }
  let ret = await updateUserInfo(request)
  //调用接口更新用户昵称 和 useStore
}

async function onEditAbstractBlur() {
  console.log("onEditAbstractBlur" )
  ElMessage.info("onEditAbstractBlur" )

  var userId = userStore.userInfo.user_id;

  let request : UpdateUserInfoRequest = {
    uid: userId,
    abstract: url,
  }
  let ret = await updateUserInfo(request)
  //调用接口更新用户昵称 和 useStore
}

</script>

<template>
  <div class="my-info">
    <el-form-item label="头像">
      <el-avatar :src="userInfo.avatar"></el-avatar>
      <ImageUpload :image_type="'avatar'" @onUploadSuccess="onUploadSuccess" />
    </el-form-item>
    <el-form-item label="用户号">
      <span>{{userInfo.uid}}</span>
    </el-form-item>
    <el-form-item label="昵称">
      <div class="edit">
        <span v-if="!showEditName">{{userInfo.name}}</span>
        <el-input  v-else v-model="userInfo.name" placeholder="请输入昵称" @blur="onEditNameBlur"></el-input>
        <svg-icon iconName="icon-bianji" @click="editNickName()"></svg-icon>
      </div>
    </el-form-item>
    <el-form-item label="简介">
      <div class="edit">
        <span v-if="!showEditAbstract">{{userInfo.abstract}}</span>
        <el-input  v-else v-model="userInfo.abstract" placeholder="请输入简介" @blur="onEditAbstractBlur"></el-input>
        <svg-icon iconName="icon-bianji" @click="editAbstract()"></svg-icon>
      </div>
    </el-form-item>
    <a @click="logout">退出登录</a>
  </div>
</template>

<style scoped>
.my-info {
  padding: 10px;
}
.edit {
  display: flex;
}
</style>