<script setup lang="ts">
import {ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();

const props = defineProps({
  keyword: {
    type: [String, Number],
    required: true
  },
});

const searchTxt = ref(props.keyword)
</script>

<template>
搜索到到群列表


  <el-input v-model="searchTxt" placeholder="请输入用户号或者用户昵称"></el-input>
</template>

<style scoped>

</style>