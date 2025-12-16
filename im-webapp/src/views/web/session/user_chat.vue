<script setup lang="ts">
// 声明从路由接收的 props
import {onActivated, onMounted, ref, watch} from "vue";
import SvgIcon from "@/components/SvgIcon.vue";
import {useUserInfoStore} from "@/stores";

const props = defineProps({
  id: { // 对应路由配置中的 :id
    type: [String, Number],
    required: true
  },
});


const store = useUserInfoStore()

onMounted(() => {
  console.log("onMounted...." + props.id)
  //TODO: 根据uid查询用户详细信息
})

// 2. 核心解决方案：监听 props.id 的变化
// 当路由参数变化，导致 props.id 变化时，此 watch 钩子会触发
watch(
    () => props.id, // 监听的响应式源
    (newId, oldId) => {
      console.log(`--- watch 触发: ID 从 ${oldId} 变为 ${newId} ---`);
    },
);

const inputText = ref(''); // 输入框内容

function sendMsg() {
  const webSocket : WebSocket = useUserInfoStore().ws
  webSocket.send(JSON.stringify({
    "type": "login",
    "data": {
      "to_uid" : props.id,
      "msg": inputText.value
    }
  }))
}

</script>

<template>
<div class="user-chat-container">
  <div class="chat-header">
    与user_chat {{props.id}} 对话中
  </div>

  <div class="chat-history">
    <el-input v-model="inputText" type="textarea" placeholder="请输入内容"  />
  </div>

  <div class="chat-input">
    <el-input  placeholder="请输入简介{{props.id}}"  v-model="inputText" />
    <svg-icon icon-name="icon-zhifeiji" @click="sendMsg"></svg-icon>
  </div>
</div>

</template>

<style scoped lang="scss">
.user-chat-container {
  width: 600px;
  height: 100%;
  display: flex;
  flex-direction: column;
  .chat-header {
    flex-grow: 1;
    flex-basis: 0;
  }
  .chat-history {
    flex-grow: 12;
    flex-basis: 0;
  }
  .chat-input {
    flex-grow: 1.5;
    flex-basis: 0;
    display: flex;
  }
}
</style>
