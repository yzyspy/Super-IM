<script setup lang="ts">
import SvgIcon from "@/components/SvgIcon.vue";
</script>

<template>
  <div class="fim_web">
    <div class="fim_slide">
        <div class="avatar">
<!--          <img src="../../assets/logo.svg" alt="Avatar">-->
          <svg-icon iconName="icon-qqqingliao"></svg-icon>
        </div>
        <div class="fim_menu">
          <div clas="icon">
            <svg-icon iconName="icon-chat"></svg-icon>
          </div>
          <div clas="icon">
            <svg-icon iconName="icon-chat"></svg-icon>
          </div>
          <div clas="icon">
            <svg-icon iconName="icon-chat"></svg-icon>
          </div>
          <div clas="icon">
            <svg-icon iconName="icon-chat"></svg-icon>
          </div>
        </div>
    </div>
    <router-view></router-view>
  </div>
</template>

<style scoped>

</style>
