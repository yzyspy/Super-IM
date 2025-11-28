
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
           <span>添加群聊</span>
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
                       <el-menu-item index="1-1">
                         <img src="http://localhost/api/file/uploads/avatar/logo_8hpzr7.png" width="30px" height="30px" />
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
  width: 150px;
}
</style>
<script lang="ts" setup>
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
} from '@element-plus/icons-vue'
import {ref} from "vue";
import {searchUser, type SearchUserRequest} from "@/api/user_api";
import {ElMessage} from "element-plus";

const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}

const searchTxt = ref('')

async function handleSearch() {
  console.log(searchTxt.value)
  const req: SearchUserRequest = {
    uid: searchTxt.value,
    nickname: searchTxt.value,
  }
  let ret = await searchUser(req)
  console.log(ret)
}
</script>