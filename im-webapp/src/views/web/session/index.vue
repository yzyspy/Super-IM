<script setup lang="ts">
import {onMounted, ref} from "vue";
import {
  type QueryChatHistoryRequest,
  chatSessionListApi,
  type SessionUserInfo,
  type ChatSessionRequest
} from "@/api/chat_api";
import type {UserInfo} from "@/api/user_api";

const req : ChatSessionRequest = {
  page: 0,
  limit: 0,
  key: "1"
}

const sessionUserList = ref<SessionUserInfo[]>([])
onMounted(async ()=> {

//const list = []


  const chatRes = await chatSessionListApi(req)

  console.log(chatRes)
  console.log(chatRes.list.length)

  if ( chatRes.list.length > 0) {
    for (let item of chatRes.list) {
      sessionUserList.value.push(
          {
            user_id : item.user_id,
            nickname : item.nickname,
            avatar : item.avatar
          }
      )
    }
  }
})


</script>

<template>
  <div class="session_container">
    <!-- 最近参与单聊的列表 和 参与的群聊的列表 -->
    <div class="session_list">
      <div v-for="item in sessionUserList">
        {{item.nickname}}
        <el-avatar :src="item.avatar" width="30px" height="30px" />
      </div>
    </div>
    <div>
      <router-view></router-view>
    </div>
  </div>

</template>

<style scoped>
.session_container {
  display: flex;
}

</style>