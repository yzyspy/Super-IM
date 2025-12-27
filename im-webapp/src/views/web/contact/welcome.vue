<script setup lang="ts">
import SvgIcon from "@/components/SvgIcon.vue";
</script>

<template>
  <div class="container">
    <div class="title">欢迎使用FIM即时通讯</div>
    <div>
      <div>
        这是一款基于vue3+go-zero的前后端微服务项目
        <div>前端:vue3 typescript elementUI-pluse WebRTC</div>
        <div>后端:go-zero gorm elasticsearch websocket</div>
      </div>
      基本功能:
      <ul>
        <li>友单聊功能,支持文字、表情、图片、文件</li>
        <li>好友搜索,好友申请</li>
        <li>用户群聊功能,群文件、群禁言、管理员、群搜索 </li>
        <li>用户视频通话功能 </li>
      </ul>
    </div>
  </div>
</template>

<style scoped lang="scss">
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
   .title {
     font-size: 40px;
     color: #0086b3;
   }
}
</style>
