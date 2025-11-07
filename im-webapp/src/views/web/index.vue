<script setup lang="ts">
import SvgIcon from "@/components/SvgIcon.vue";
import {useRouter} from "vue-router";

const router = useRouter()


function change(menu : string) {
  //router.push({ name: menu})
  router.push("/" + menu);
}

</script>

<template>
  <div class="container">
    <div class="fim_web">
      <div class="fim_slide">
        <div class="avatar">
          <svg-icon iconName="icon-qqqingliao" @click="change('info')"></svg-icon>
        </div>
        <div class="fim_menu">
          <div clas="icon">
            <svg-icon iconName="icon-chat1" @click="change('info')"></svg-icon>
          </div>
          <div clas="icon">
            <svg-icon iconName="icon-chat" @click="change('contact')"></svg-icon>
          </div>
          <div clas="icon">
            <svg-icon iconName="icon-chat" @click="change('notice')"></svg-icon>
          </div>
          <div clas="icon">
            <svg-icon iconName="icon-chat" @click="change('session')"></svg-icon>
          </div>
        </div>
      </div>
      <div class="fim_content">
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
  background-color: lightblue; /* 为了方便演示 */

  display: flex;
}

.fim_slide {
  background-color: #ffecec;
  width: 80px;

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
</style>
