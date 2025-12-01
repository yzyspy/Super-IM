
<template>
  <div class="contact_view">
    <div class="contact_slide">
        <div class="head">
           <div>
             <el-input
               placeholder="搜索"
               prefix-icon="location"
               v-model="searchTxt"
               @blur="handleSearch"
               />
           </div>
           <span> <svg-icon iconName="icon-tianjiaqunliao"></svg-icon></span>
        </div>
         <div class="content_menu">
           <el-scrollbar height="480px">
                <el-menu>
                     <el-sub-menu index="1">
                       <template #title>
                         <span>我创建的群聊</span>
                       </template>
                       <el-menu-item index="1-3">item three</el-menu-item>
                       <el-menu-item index="1-3">item three</el-menu-item>
                     </el-sub-menu>

                     <el-sub-menu index="2">
                       <template #title>
                         <span>我参加的群聊</span>
                       </template>
                       <el-menu-item index="1-3">item three</el-menu-item>
                       <el-menu-item index="1-3">item three</el-menu-item>
                     </el-sub-menu>

                     <el-sub-menu index="3">
                       <template #title>
                         <span>我的好友</span>
                       </template>
                       <el-menu-item index="1-1" v-for="(item, index) in friendList">
                         <div class="user_item" @click="onUserClick(item)">
                           <div>
                             <el-avatar :src="item.avatar" width="30px" height="30px" />
                           </div>
                           <div class="nickname">
                             {{item.nickname}}
                           </div>
                         </div>
                       </el-menu-item>
                       <el-sub-menu index="1-4">
                         <template #title>二级菜单</template>
                         <el-menu-item index="1-4-1">item one</el-menu-item>
                       </el-sub-menu>
                     </el-sub-menu>
               </el-menu>
           </el-scrollbar>
         </div>
    </div>
    <div class="line"></div>
    <div class="contact_main">
         <router-view></router-view>
    </div>
  </div>
</template>

<style scoped>
.head {
  display: flex;
}
.contact_view {
  display: flex;
}
.contact_slide {
  height: 500px;
  width: 140px;
}
.user_item {
  display: flex;
  .nickname {
     margin-left: 10px;
  }
}
.contact_main {

}
.line {
  width: 1px;
  background-color: #e6e6e6;
}
</style>
<script lang="ts" setup>
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
} from '@element-plus/icons-vue'
import {onMounted, reactive, ref} from "vue";
import {queryMyFriendList, searchUser, type SearchUserRequest, type UserInfo} from "@/api/user_api";
import {ElMessage} from "element-plus";
import {useRouter} from "vue-router";
import SvgIcon from "@/components/SvgIcon.vue";

const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}

const router = useRouter();

const searchTxt = ref('')
const friendList = ref<UserInfo[]>([])

onMounted(() => {
    queryMyFriendList().then((res) => {
       friendList.value = res.list
     })
  })

async function handleSearch() {
  const req: SearchUserRequest = {
    uid: searchTxt.value,
    nickname: searchTxt.value,
  }
  let ret = await searchUser(req)
  console.log(ret.list)
}

function onUserClick(item: UserInfo) {
 // router.push("/user_chat")
  router.push({ name: 'user_chat' , params: { id: item.uid} })
}
</script>