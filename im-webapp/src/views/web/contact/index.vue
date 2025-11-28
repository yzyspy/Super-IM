<script setup lang="ts">
</script>

<template>
  <div class="contact_view">
    <div class="contact_slide">
        <div class="head">
            <span>搜索</span>
           <span>添加群聊</span>
        </div>
         <div class="content_menu">
           <el-menu
               default-active="2"
               class="el-menu-vertical-demo"
               @open="handleOpen"
               @close="handleClose"
           >
             <el-sub-menu index="1">
               <template #title>
                 <span>我创建的群聊</span>
               </template>
               <el-menu-item index="1-3">item three</el-menu-item>
               <el-menu-item index="1-3">item three</el-menu-item>
             </el-sub-menu>

             <el-sub-menu index="1">
               <template #title>
                 <span>我参加的群聊</span>
               </template>
               <el-menu-item index="1-3">item three</el-menu-item>
               <el-menu-item index="1-3">item three</el-menu-item>
             </el-sub-menu>

             <el-sub-menu index="1">
               <template #title>
                 <span>我的好友</span>
               </template>
               <el-menu-item-group title="家人分组">
                 <el-menu-item index="1-1">item one</el-menu-item>
                 <el-menu-item index="1-2">item two</el-menu-item>
               </el-menu-item-group>
               <el-sub-menu index="1-4">
                 <template #title>二级菜单</template>
                 <el-menu-item index="1-4-1">item one</el-menu-item>
               </el-sub-menu>
             </el-sub-menu>
           </el-menu>
         </div>
    </div>
    <div class="contact_main">
         <router-view></router-view>
    </div>
  </div>
</template>

<style scoped>

</style>
