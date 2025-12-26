<script setup lang="ts">
import SvgIcon from "@/components/SvgIcon.vue";
</script>

<template>
  <div class="container">
    <div class="title">欢迎使用FIM即时通讯</div>
    <div>
      <div>
        <svg-icon icon-name="icon-shipintonghua-tianchong"></svg-icon>
        这是一款基于vue3+go-zero的前后端微服务项目
        前端:vue3 typescript elementUI-pluse WebRTC
        后端:go-zero gorm elasticsearch websocket
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
   .title {
     font-size: 40px;
   }
}
</style>
