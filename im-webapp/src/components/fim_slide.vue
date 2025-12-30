<script setup lang="ts">

import SvgIcon from "@/components/SvgIcon.vue";
import {useRouter} from "vue-router";
import {useUserInfoStore} from "@/stores";


const router = useRouter()
const userStore = useUserInfoStore();

function change(menu : string) {
  //router.push({ name: menu})
  router.push("/" + menu);
}
</script>

<template>
  <div class="fim_slide">
    <div class="avatar">
      <el-avatar :src="userStore.userInfo.avatar"></el-avatar>
    </div>
    <div class="fim_menu">
      <div class="icon">
        <svg-icon iconName="icon-chat1" @click="change('session')"></svg-icon>
      </div>
      <div class="icon">
        <svg-icon iconName="icon-lianxiren" @click="change('contact')"></svg-icon>
      </div>
      <div class="icon">
        <svg-icon iconName="icon-tongzhi" @click="change('notice')"></svg-icon>
      </div>
    </div>
    <div class="caidan-container">
      <div class="caidan">
        <svg-icon iconName="icon-caidan" @click="change('info')"></svg-icon>
      </div>
    </div>
  </div>
</template>

<style scoped>

.fim_slide {
  background-color: #dededf;
  width: 60px;
  flex-shrink: 0; /* 禁止压缩 */

  display: flex;
  flex-direction: column;
  align-items: center;
  height: 100%;
}

.fim_menu {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  height: 30%;
}
.icon {
  cursor: pointer;
  margin-top: 30px;
}

.avatar {
  cursor: pointer;
  margin-top: 20px;
}
.caidan {
  cursor: pointer;
}
.caidan-container {
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  height: 280px;
}
.caidan {
  margin-bottom: 10px;
}
</style>