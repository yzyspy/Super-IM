<script setup lang="ts">
// 声明从路由接收的 props
import {onActivated, onMounted, ref, watch} from "vue";
import SvgIcon from "@/components/SvgIcon.vue";

const props = defineProps({
  id: { // 对应路由配置中的 :id
    type: [String, Number],
    required: true
  },
});
onMounted(() => {
  console.log("onMounted...." + props.id)
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

</script>

<template>
<div class="user-chat-container">
  <div class="chat-header">
    与user_chat {{props.id}} 对话中
  </div>

  <div class="chat-history">
    聊天记录
  </div>

  <div class="chat-input">
    <el-input  placeholder="请输入简介{{props.id}}"  v-model="inputText" />
    <svg-icon icon-name="icon-zhifeiji"></svg-icon>
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
    flex-grow: 10;
    flex-basis: 0;
  }
  .chat-input {
    flex-grow: 1;
    flex-basis: 0;
    display: flex;
  }
}
</style>
