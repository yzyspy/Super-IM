<script setup lang="ts">
import type {Msg, MsgType, TextMsg, UserBaseType} from "@/api/chat_api";
import {watch} from "vue";
import {useUserInfoStore} from "@/stores";


interface Props {
  data: Msg
}

const props = defineProps<Props>()

</script>

<template>
<div class="fim_msg">
  <div class="info">
<!--    <el-text>{{ props.data.user.nickName }}</el-text>-->
    <el-avatar :src="props.data.user.avatar" size="small"></el-avatar>
    <div class="msg">
      <div class="text_msg">
        {{props.data.msg.textMsg.content}}
      </div>
    </div>
  </div>
</div>
</template>

<style scoped>
.info {
  display: flex;
  margin-top: 20px;
}
.text_msg {
  background-color: #D5E2FC;
  margin-left: 10px;
}
</style>