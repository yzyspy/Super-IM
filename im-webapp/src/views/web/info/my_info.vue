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
const nickNameInputRef = ref()
function editNickName() {
  if (showEditName.value == true) {
    showEditName.value = false
  } else {
    showEditName.value = true
    //这个地方获得焦点，后续才能触发 失去焦点的blur事件
    //另外这个地方需要注意的是需要延迟10ms，否者不会触发获取焦点
    setTimeout(() => {
      nickNameInputRef.value.focus()
    }, 10);
  }
}

const showEditAbstract = ref(false)
// 声明 ref（变量名要与模板中的一致）
const abstractInputRef = ref()
function editAbstract() {
  if (showEditAbstract.value == true) {
    showEditAbstract.value = false
  } else {
    showEditAbstract.value = true
    //这个地方获得焦点，后续才能触发 失去焦点的blur事件
    //另外这个地方需要注意的是需要延迟10ms，否者不会触发获取焦点
    setTimeout(() => {
      abstractInputRef.value.focus()
    }, 10);

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
        <el-input ref="nickNameInputRef"  v-else v-model="userInfo.name" placeholder="请输入昵称" @blur="onEditNameBlur"></el-input>
        <svg-icon iconName="icon-bianji" @click="editNickName()"></svg-icon>
      </div>
    </el-form-item>
    <el-form-item label="简介">
      <div class="edit">
        <span v-if="!showEditAbstract">{{userInfo.abstract}}</span>
        <el-input ref="abstractInputRef" v-else v-model="userInfo.abstract" placeholder="请输入简介" type="textarea" @blur="onEditAbstractBlur"></el-input>
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