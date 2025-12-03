<script setup lang="ts">

import {onMounted, ref} from "vue";
import {getFriendApplyList, type FriendApplyItem, type UserInfo} from "@/api/user_api";
import { ArrowDown } from '@element-plus/icons-vue';
import {ElMessage} from "element-plus";


const applyList = ref<FriendApplyItem[]>([])

onMounted(() => {
  //获取我的好友申请验证列表
  getFriendApplyList().then(res => {
    applyList.value = res.list
  })
})

const handleCommand = (command: string) => {
  console.log(command)
}

</script>
<template>
  <el-menu>
    <el-sub-menu index="1">
      <template #title>
        <span>我的好友申请</span>
      </template>
      <el-menu-item index="1-3" v-for="(item, index) in applyList" :key="index">
       <span>{{item.nickname}}</span>

        <div v-if="item.status === 0">
              <el-dropdown @command=handleCommand>
                <span class="el-dropdown-link" @click="handleCommand('agree')">
                  同意<el-icon> <arrow-down/></el-icon>
                </span>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item command="reject" >拒绝</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
        </div>
        <div v-else-if="item.status === 1">
          已同意
        </div>
        <div v-else-if="item.status === 2">
          已拒绝
        </div>

      </el-menu-item>
    </el-sub-menu>
  </el-menu>

</template>

<style scoped>
.el-dropdown-link {
  cursor: pointer;
  color: var(--el-color-primary);
  display: flex;
  align-items: center;
}
</style>