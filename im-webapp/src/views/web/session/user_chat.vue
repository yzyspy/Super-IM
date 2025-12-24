<script setup lang="ts">
// 声明从路由接收的 props
import {onActivated, onMounted, ref, watch} from "vue";
import SvgIcon from "@/components/SvgIcon.vue";
import {useUserInfoStore} from "@/stores";


const props = defineProps({
  uid: { // 对应路由配置中的 :uid
    type: [Number],
    required: true
  },
  nick_name: { // 对应路由配置中的 :nick_name
    type: [String],
    required: true
  },
  avatar: { // 对应路由配置中的 :avatar
    type: [String],
    required: true
  },
});


const store = useUserInfoStore()

onMounted(() => {
  console.log("onMounted...." + props.uid)
  //TODO: 根据uid查询用户详细信息
})

// 2. 核心解决方案：监听 props.id 的变化
// 当路由参数变化，导致 props.id 变化时，此 watch 钩子会触发
watch(
    () => props.uid, // 监听的响应式源
    (newId, oldId) => {
      console.log(`--- watch 触发: ID 从 ${oldId} 变为 ${newId} ---`);
    },
);

const inputText = ref(''); // 输入框内容

function sendMsg() {
  const w : WebSocket = useUserInfoStore().ws
  w.send(JSON.stringify({
    "type": "login",
    "data": {
      "to_uid" : props.uid,
      "msg": inputText.value
    }
  }))
}

</script>

<template>
<div class="user-chat-container">
  <div class="chat-header">
    <el-avatar :src="props.avatar"/>{{props.nick_name}}
  </div>
  <div class="user_chat_inner_view">
    <div class="user_chat_inner_head">
      <el-scrollbar height="350px">
         <div class="msg">
            聊天记录
         </div>
      </el-scrollbar>
    </div>

    <div class="user_chat_inner_menu">
      <svg-icon icon-name="icon-zhifeiji" @click="sendMsg"></svg-icon>
      <svg-icon icon-name="icon-zhifeiji" @click="sendMsg"></svg-icon>
      <svg-icon icon-name="icon-zhifeiji" @click="sendMsg"></svg-icon>
      <svg-icon icon-name="icon-zhifeiji" @click="sendMsg"></svg-icon>
      <svg-icon icon-name="icon-zhifeiji" @click="sendMsg"></svg-icon>
    </div>

    <div class="user_chat_inner_box">
      <el-input  placeholder="发送消息"  v-model="inputText" />
      <svg-icon icon-name="icon-zhifeiji" @click="sendMsg"></svg-icon>
    </div>
  </div>
</div>

</template>

<style scoped lang="scss">
.msg {
  font-size: 200px;
}
.user-chat-container {
  width: 600px;
  height: 100%;
  display: flex;
  flex-direction: column;
  .chat-header {
    flex-grow: 1;
    flex-basis: 0;
    margin-top: 10px;
  }
  .chat-history {
    flex-grow: 12;
    flex-basis: 0;
  }
  .user_chat_inner_box {
    flex-grow: 1.5;
    flex-basis: 0;
    display: flex;
  }
}
</style>
