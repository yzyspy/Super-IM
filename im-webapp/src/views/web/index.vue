<script setup lang="ts">
import SvgIcon from "@/components/SvgIcon.vue";
import Fim_slide from "@/components/fim_slide.vue";
import {reactive} from "vue";

interface sessionType {
  id: number
  title: string
  avatar: string
  date: string
  msgPreview: string
  isTop: boolean
  type: "user" | "group"
}

const data = reactive({
  list: [
    {
      "userId": 4,
      "avatar": ""
    }
  ]
})

</script>

<template>
  <div class="container">
    <div class="fim_web">
      <!-- 左侧菜单栏 -->
      <fim_slide></fim_slide>
      <div class="session_main">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  display: flex;
  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */
  height: 100vh; /* 关键：父容器需要有一个明确的高度，这里使用视口高度 */
}

.fim_web {
  width: 800px; /* 你的div宽度 */
  height: 500px; /* 你的div高度 */
  background-color: #f3f3f3; /* 为了方便演示 */

  display: flex;
}
</style>
