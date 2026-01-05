<script setup lang="ts">
// 声明从路由接收的 props
import {onActivated, onMounted, ref, watch} from "vue";
import SvgIcon from "@/components/SvgIcon.vue";
import {useUserInfoStore} from "@/stores";
import Fim_msg from "@/components/fim_msg.vue";
import type {Msg, MsgType, UserBaseType} from "@/api/chat_api";


const props = defineProps({
  uid: { // 对应路由配置中的 :uid
    type: [String],
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

watch(
    () => store.latestMsg,
    (newId, oldId) => {
      console.log(`--- watch 触发: latestMsg 从 ${oldId} 变为 ${newId} ---`);
      msgList.value.push(store.latestMsg)
    },
    { deep: true } // 如果不加这个，修改 userInfo.age 是监听不到的
);


const inputText = ref(''); // 输入框内容

function sendMsg() {
  const w : WebSocket = store.ws
  w.send(JSON.stringify({
    "rev_user_id": props.uid,
    "msg": {
      "msg_type" : 1,
      "content": inputText.value
    }
  }))
  console.log("sendMsg...." + inputText.value)

  var userInfo = store.userInfo;
  const msg : MsgType = {
    rev_user: {
      id:props.uid,
      nickName: props.nick_name,
      avatar: props.avatar
    },
    send_user: {
      id:userInfo.user_id,
      nickName: userInfo.user_name,
      avatar: userInfo.avatar
    },
    msg: {
      msg_type: 0,
      content: inputText.value
    },
    create_at: "",
    is_me: true
  }
  msgList.value.push(msg)
  inputText.value = ''
}
 const msgList = ref<MsgType[]>( [])

const convertMsg = (msg : MsgType) => {
  const m : Msg = {
    id: 0,
    user: {
      id: msg.is_me ? msg.rev_user.id : msg.send_user.id,
      nickName: msg.is_me ? msg.rev_user.nickName : msg.send_user.nickName,
      avatar:  msg.is_me ? msg.rev_user.avatar : msg.send_user.avatar,
    },
    isMe: msg.is_me,
    createdAt: '',
    msg: {
      type: 0,
      textMsg: {
        content: msg.msg.content
      }
    }
  }
  return m
}
</script>

<template>
<div class="user-chat-container">
  <div class="chat-header">
    <el-avatar :src="props.avatar"/>{{props.nick_name}}
  </div>
  <div class="user_chat_inner_view">
    <div class="user_chat_inner_head">
      <el-scrollbar height="100%">
         <div class="msg">
             <template v-for="msg in msgList">
               <fim_msg :data="convertMsg(msg)"></fim_msg>
             </template>
         </div>
      </el-scrollbar>
    </div>

    <div class="user_chat_inner_menu">
      <svg-icon icon-name="icon-shangchuantupian" @click="sendMsg"></svg-icon>
      <svg-icon icon-name="icon-shipintonghua-tianchong" @click="sendMsg"></svg-icon>
      <svg-icon icon-name="icon-wenjianjia" @click="sendMsg"></svg-icon>
      <svg-icon icon-name="icon-biaoqing" @click="sendMsg"></svg-icon>
      <svg-icon icon-name="icon-yuyintonghuaiconx" @click="sendMsg"></svg-icon>
      <svg-icon icon-name="icon-shipintonghua" @click="sendMsg"></svg-icon>
    </div>

    <div class="user_chat_inner_box">
      <el-input  placeholder="发送消息"  v-model="inputText" @keyup.enter="sendMsg"/>
      <svg-icon icon-name="icon-zhifeiji" @click="sendMsg"></svg-icon>
    </div>
  </div>
</div>

</template>

<style scoped lang="scss">
.user-chat-container {
  width: 720px;
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
  .user_chat_inner_view {
     height: 100%;
    .user_chat_inner_box {
      flex-grow: 2;
      flex-basis: 0;
      display: flex;
    }
    .user_chat_inner_head {
       height: calc(100% - 75px);
    }
    .user_chat_inner_menu {
       border-top: 1px solid #e2e2e2;
       cursor: pointer;
    }
    .user_chat_inner_box {

    }
  }
}
</style>
