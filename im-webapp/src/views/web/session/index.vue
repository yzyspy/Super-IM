<script lang="ts">
import {onMounted} from "vue";
import {type QueryChatHistoryRequest, queryChatSession} from "@/api/chat_api";

 onMounted(()=> {
  const req : QueryChatHistoryRequest = {
    page: 0,
    limit: 0
  }
  var response = queryChatSession(req).then((data)=>{
    console.log("queryChatSession.....")
    console.log(data)
  })
})
</script>

<template>
  <div class="session_container">
    <!-- 最近参与单聊的列表 和 参与的群聊的列表 -->
    <div class="session_list">
      session_list
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