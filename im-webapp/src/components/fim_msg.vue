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
   <el-avatar :src="props.data.user.avatar" size="small"></el-avatar>
  <div class="info">
    <el-text>{{ props.data.user.nickName }}</el-text>
    <div class="msg">
      <div class="text_msg" v-if="props.data.msg.type === 1">
        {{props.data.msg.textMsg.content}}
      </div>
    </div>
  </div>
</div>
</template>

<style scoped>

</style>