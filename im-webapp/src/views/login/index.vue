<template>
  <div class="login-container">
    <div class="content">
      <div class="banner">
      </div>
      <input type="text" v-model="user_name" placeholder="用户名">
      <input type="password" v-model="password" placeholder="密码">
      <button @click="login">登录</button>
    </div>
  </div>
</template>


<script setup lang="ts">

// 在ES6模块中：
// 当导出的是默认导出（export default）时，导入时可以不加花括号，且可以任意命名。
// 当导出的是命名导出（export）时，导入时必须加花括号，且名称必须与导出时一致（或者使用as重命名）

import {service} from '@/api/index'; //如果是导出的时候是 export default service 不需要加大括号， 如果不是default 就需要加大括号
import {useUserInfoStore} from '@/stores/index';

import { ref } from 'vue'

const user_name = ref('')
const password = ref('')

function login() {
  console.log('login');
  service.request({
    method: 'post',
    url: '/api/auth/login',
    data: {
      user_name: user_name,
      password: password
    }
  }).then((res: any) => {
    console.log(res);
   // useUserInfoStore().setToken(res.data.)
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
</style>