<script setup lang="ts">
import type {TextMsg, UserBaseType} from "@/api/chat_api";

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

</script>

<template>

</template>

<style scoped>

</style>