<script setup lang="ts">
import {useRouter} from "vue-router";
import {onMounted, ref} from "vue";
import {applyFriend, type ApplyFriendRequest, searchUser, type SearchUserRequest, type UserInfo} from "@/api/user_api";
import {ElMessage} from "element-plus";

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
   searchUser(req).then((res) => {
      searchUserList.value = res.list
    })
})
const searchTxt = ref(props.keyword)
const searchUserList = ref<UserInfo[]>([])

const addFriend = (uid: number) => {
  const req : ApplyFriendRequest = {
    uid : uid
  }
  applyFriend(req).then((res) => {
    if(res.code === 0 && res.data) {
      ElMessage.success('申请成功')
    } else {
      ElMessage.error(res.msg)
    }
  })
}

function doSearchUser() {
  const req: SearchUserRequest = {
    keyword : searchTxt.value
  }
  console.log(req)
  searchUser(req).then((res) => {
    searchUserList.value = res.list
  })
}

</script>

<template>
  <div class="contact-search">
    <el-input v-model="searchTxt" placeholder="请输入用户号或者用户昵称"></el-input>
    <el-button @click="doSearchUser">搜索</el-button>
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
      <el-button @click="addFriend(item.uid)">加好友</el-button>
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
}
.contact-search {
  display: flex;
}
.user_item {
  display: flex;
}
</style>