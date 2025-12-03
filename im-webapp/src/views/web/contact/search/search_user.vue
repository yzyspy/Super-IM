<script setup lang="ts">
import {useRouter} from "vue-router";
import {onMounted, ref} from "vue";
import {searchUser, type SearchUserRequest} from "@/api/user_api";

const router = useRouter();

const props = defineProps({
  keyword: {
    type: [String, Number],
    required: true
  },
});

onMounted(() => {
  const keyword = props.keyword
  console.log(keyword)

  const req: SearchUserRequest = {
    uid: keyword,
    nickname: keyword,
  }
  let ret = searchUser(req).then((res) => {
    console.log(res.list)
  })
})
const searchTxt = ref(props.keyword)

</script>



<template>
  <div class="contact-search">
    <el-input v-model="searchTxt" placeholder="请输入用户号或者用户昵称"></el-input>
  </div>
  <div class="search-result">
    <div class="search-user">

    </div>
    <div class="search-user">

    </div>
    <div class="search-user">

    </div>
    <div class="search-user">

    </div>
    <div class="search-user">

    </div>
    <div class="search-user">

    </div>
    <div class="search-user">

    </div>
  </div>

</template>

<style scoped>
.search-result {
  display: flex;
  flex-wrap: wrap;
}
.search-user {
  width: 100px;
  height: 100px;
  background-color: #669900;
  border: red solid 1px;
}
</style>