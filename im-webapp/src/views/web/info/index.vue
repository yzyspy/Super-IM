<script setup lang="ts">
</script>

<template>
  <div class="fim_info_body">
    <div class="header">
      <div class="tab">
        <router-link :to="{name:'my_info'}">个人资料</router-link>
      </div>
      <div class="tab">
        <router-link :to="{name:'base_info'}">基本信息</router-link>
      </div>
      <div class="tab" >
        <router-link :to="{name:'safe_info'}">安全设置</router-link>
      </div>
      <div class="tab">
        <router-link :to="{name:'role_info'}">权限设置</router-link>
      </div>
    </div>
    <div class="body">
      <router-view/>
    </div>
  </div>
</template>

<style scoped lang="scss">
.header {
  display: flex;
}
.tab {
  // 上、右、下、左
  margin: 10px 10px 0 10px;
}
</style>