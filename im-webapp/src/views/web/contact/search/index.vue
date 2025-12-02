<script setup lang="ts">
import {useRouter} from "vue-router";
import {searchUser, type SearchUserRequest} from "@/api/user_api";

const router = useRouter()
function changeTab(index: number) {
    // TODO: change tab
    if (index == 0) {
        router.push({ name: 'search_user' })
    } else if (index == 1) {
        router.push({ name: 'search_group' })
    }
}
//
// const req: SearchUserRequest = {
//   uid: searchTxt.value,
//   nickname: searchTxt.value,
// }
// let ret = await searchUser(req)
// console.log(ret.list)

</script>

<template>
  <el-button @click="changeTab(0)">找人</el-button>
  <el-button @click="changeTab(1)">找群</el-button>
 <router-view></router-view>

</template>

<style scoped>

</style>