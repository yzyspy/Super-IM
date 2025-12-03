<script setup lang="ts">
import {useRouter} from "vue-router";
import {searchUser, type SearchUserRequest} from "@/api/user_api";
import {ref, watch} from "vue";

const router = useRouter()


const props = defineProps({
  keyword: { // 对应路由配置中的 :id
    type: [String, Number],
    required: true
  },
});

watch(
    () => props.keyword, // 监听的响应式源
    (newId, oldId) => {
      console.log(`--- watch 触发: ID 从 ${oldId} 变为 ${newId} ---`);
      searchTxt.value = newId
    },
);

const searchTxt = ref(props.keyword)

</script>

<template>

  <router-link :to="{name:'search_user', params: { keyword: searchTxt}}">找人</router-link>
  <router-link :to="{name:'search_group', params: { keyword: searchTxt}}">找群</router-link>

 <router-view></router-view>

</template>

<style scoped>

</style>