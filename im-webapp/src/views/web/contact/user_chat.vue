<script setup lang="ts">
// 声明从路由接收的 props
import {onActivated, onMounted} from "vue";

const props = defineProps({
  id: { // 对应路由配置中的 :id
    type: [String, Number],
    required: true
  },
});
onMounted(() => {
  console.log("onMounted...." + props.id)
})

onActivated(() => {
  console.log("onActivated...." + props.id)
})

</script>

<template>
user_chat
</template>

<style scoped>

</style>
