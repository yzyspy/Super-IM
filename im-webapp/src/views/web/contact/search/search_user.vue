<script setup lang="ts">
import {useRouter} from "vue-router";
import {onMounted, ref} from "vue";
import {searchUser, type SearchUserRequest, type UserInfo} from "@/api/user_api";

const router = useRouter();

const props = defineProps({
  keyword: {
    type: [String],
    required: true
  },
});

onMounted(() => {
  const req: SearchUserRequest = {
    keyword : props.keyword
  }
  let ret = searchUser(req).then((res) => {
    searchUserList.value = res.list
  })
})
const searchTxt = ref(props.keyword)
const searchUserList = ref<UserInfo[]>([])
</script>



<template>
  <div class="contact-search">
    <el-input v-model="searchTxt" placeholder="请输入用户号或者用户昵称"></el-input>
  </div>
  <div class="search-result">
    <div class="search-user" v-for="item in searchUserList">
      <div class="user_item">
        <div>
          <el-avatar :src="item.avatar" width="30px" height="30px" />
        </div>
        <div class="nickname">
          {{item.nickname}}
        </div>
      </div>
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