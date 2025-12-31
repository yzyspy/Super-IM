<script setup lang="ts">
import type {TextMsg, UserBaseType} from "@/api/chat_api";
import {watch} from "vue";
import {useUserInfoStore} from "@/stores";

interface MsgType {
  "id": number
  "user": UserBaseType
  "isMe": boolean
  "createdAt": string
  "msg": {
    "type":number
    "textMsg":TextMsg
  }
}
interface Props {
  data: MsgType
}

const props = defineProps<Props>()
const store = useUserInfoStore()
watch(
    () => store.latestMsg,
    (newId, oldId) => {
      console.log(`--- watch 触发: latestMsg 从 ${oldId} 变为 ${newId} ---`);
    },
);
</script>

<template>
<div>
  this is fim_msg {{store.latestMsg.msg.content}}
</div>
</template>

<style scoped>

</style>