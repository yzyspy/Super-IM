<template>
  <div class="login-container">
    <div class="content">
      <div class="banner">
      </div>
      <el-input type="text" v-model="user_name" placeholder="用户名"/>
      <el-input type="password" v-model="password" placeholder="密码"/>
      <el-button type="primary" @click="register">注册</el-button>
      <div class="bottom">
        <router-link to="/login">去登录</router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref} from "vue";
import {service} from "@/api";
import {useUserInfoStore} from "@/stores";
import {ElMessage} from "element-plus";

const user_name = ref('')
const password = ref('')

function register() {
  service.request({
    method: 'post',
    url: '/api/auth/register',
    data: {
      user_name: user_name.value,
      password: password.value
    }
  }).then((res: any) => {
    console.log(res)
    if (res.code == 0) {
      //保存token
      useUserInfoStore().setToken(res.data.token)
      ElMessage.info("注册成功")
    } else {
      ElMessage.info(res.msg)
    }
  });
}
</script>

<style scoped>

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
  padding-left: 20px;
}
</style>