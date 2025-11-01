<template>
  <div class="login-container">
    <div class="content">
      <div class="banner">
      </div>
      <el-input type="text" v-model="user_name" placeholder="用户名"/>
      <el-input type="password" v-model="password" placeholder="密码"/>
      <el-button type="primary" @click="login">登录</el-button>
      <div class="bottom">
        <router-link to="">忘记密码</router-link>
        <router-link to="/register">注册账号</router-link>
      </div>
    </div>
  </div>
</template>


<script setup lang="ts">

// 在ES6模块中：
// 当导出的是默认导出（export default）时，导入时可以不加花括号，且可以任意命名。
// 当导出的是命名导出（export）时，导入时必须加花括号，且名称必须与导出时一致（或者使用as重命名）

import {service} from '@/api/index'; //如果是导出的时候是 export default service 不需要加大括号， 如果不是default 就需要加大括号
import {useUserInfoStore} from '@/stores/index';

import {useRouter} from "vue-router";

import {ref} from 'vue'
import {ElMessage} from "element-plus";

const user_name = ref('')
const password = ref('')
const router = useRouter()

function login() {
  service.request({
    method: 'post',
    url: '/api/auth/login',
    data: {
      user_name: user_name.value,
      password: password.value
    }
  }).then((res: any) => {
    if (res.code == 0) {
      //保存token
      useUserInfoStore().setToken(res.data.token)
      ElMessage.info("登录成功")
      router.replace("/")
    } else {
      ElMessage.info(res.msg)
    }
  });
}
</script>


<style scoped>
.login-container {
  display: flex;
  align-items: center;
  min-height: 100vh; /* 视口高度 */
}

.content {
  background: #f0f0f0;
  width: 600px;
  height: 450px;
  margin: auto; /* 居中 */
  display: flex;
  flex-direction: column;
}

.banner {
  background: url('../../assets/qq.jpg') no-repeat center center;
  height: 100px;
}
.bottom {
  display: flex;
}
</style>