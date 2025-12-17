<script setup lang="ts">
// 声明从路由接收的 props
import {onActivated, onMounted, ref, watch} from "vue";
import SvgIcon from "@/components/SvgIcon.vue";
import {useRouter} from "vue-router";

const router = useRouter();

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
onMounted(() => {
  console.log("onMounted...." + props.uid)
})

// 2. 核心解决方案：监听 props.id 的变化
// 当路由参数变化，导致 props.id 变化时，此 watch 钩子会触发
watch(
    () => props.uid, // 监听的响应式源
    (newId, oldId) => {
      console.log(`--- watch 触发: ID 从 ${oldId} 变为 ${newId} ---`);
    },
);

function handleSendMsg() {
  console.log("handleSendMsg...." + props.uid)
  router.push({  name: 'user_chat' , params: { id: props.uid} })
}


</script>

<template>
  <div class="user-chat-container">
    <div class="chat-header">
      user_chat {{props.nick_name}} 的详情页面
    </div>
    <el-button @click="handleSendMsg">发送消息</el-button>
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
