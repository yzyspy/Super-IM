<script setup lang="ts">
import type {TextMsg, UserBaseType} from "@/api/chat_api";

interface Props {
  "id": number
  "user": UserBaseType
  "isMe": boolean
  "createdAt": string
  "msg": {
    "type":number
    "textMsg":TextMsg
  }
}

const props = defineProps<Props>()

</script>

<template>

</template>

<style scoped>

</style>