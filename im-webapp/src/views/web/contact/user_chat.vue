<script setup lang="ts">
// 声明从路由接收的 props
import {onActivated, onMounted, watch} from "vue";

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
    // { immediate: true } // 如果想在组件初始化时就执行，可以加上这个
);

</script>

<template>
user_chat {{props.id}}
</template>

<style scoped>

</style>
