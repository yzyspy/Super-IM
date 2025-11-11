<script setup lang="ts">
import { useUserInfoStore } from "@/stores/index";
const userStore = useUserInfoStore();
userStore.loadUserInfo();

</script>

<template>
  <router-view></router-view>
</template>

<style scoped>

</style>
