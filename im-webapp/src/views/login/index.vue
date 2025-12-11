<template>
  <div class="container">
    <div class="fim_login">
      <div class="banner">

      </div>
      <div class="login_form">
        <el-form :model="form">
          <el-form-item>
            <el-input type="text" v-model="form.user_name" placeholder="用户名"/>
          </el-form-item>
          <el-form-item>
            <el-input type="password" v-model="form.password" placeholder="密码"/>
          </el-form-item>
          <el-form-item>
            <el-checkbox>记住密码</el-checkbox>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="login" style="width: 100%;">登录</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="bottom">
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

import {reactive, ref} from 'vue'
import {ElMessage} from "element-plus";
import {type AuthLoginRequest, doLogin} from "@/api/auth_api";
import {parseToken} from "@/utils/common";
import {queryUserInfo, type QueryUserInfoRequest} from "@/api/user_api";
import {initWebSocket} from "@/api/websocket_utils";

// 这个地方等语法要好好理解一下
const form = reactive<AuthLoginRequest>({
  user_name: '',
  password: ''
})
const router = useRouter()

const store = useUserInfoStore()

async function login() {
  let res = await doLogin(form);
  if (res.code!== 0) {
    ElMessage.info(res.msg)
    return
  } else {
    const userInfo = parseToken(res.data.token)
    store.setUserInfo(userInfo)
    //拉取用户信息（头像、配置等）
    let userInfoExtra = await queryUserInfo({})

    store.userInfo.abstract = userInfoExtra.data.abstract
    store.userInfo.avatar = userInfoExtra.data.avatar
    store.refreshUserInfo()

    //登录成功以后，创建websocket连接
    const ws : WebSocket = initWebSocket(res.data.token);
     store.setWebSocket(ws)

    const ws2 : WebSocket = useUserInfoStore().ws
    ws2.send(JSON.stringify({
      "type": "login",
      "data": {
        "user_id": store.userInfo.user_id
      }
    }))

    ElMessage.info("登录成功")
    router.replace("/")
  }
}

</script>


<style scoped lang="scss">
$color: #f0f0f0;
.container {
  display: flex;
  height: 100vh;
}

.fim_login {
  background: $color;
  width: 600px;
  height: 400px;
  margin: auto; /* 居中 */
  display: flex;
  flex-direction: column;
}

.login_form {
  margin-top: 30px;
  padding-left: 30px;
  padding-right: 30px;
}

.banner {
  background: url('../../assets/qq.jpg') no-repeat center center;
  height: 100px;
}
.bottom {
  display: flex;
}
</style>