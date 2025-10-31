import { ref, computed } from 'vue'
import { defineStore } from 'pinia'


// 在ES6模块中：
// 当导出的是默认导出（export default）时，导入时可以不加花括号，且可以任意命名。
// 当导出的是命名导出（export）时，导入时必须加花括号，且名称必须与导出时一致（或者使用as重命名）

export const useUserInfoStore = defineStore('userInfo', {
  // 为了完整类型推理，推荐使用箭头函数
  state: () => ({
    userInfo: {
      name: '',
      token: '',
    }
  }),
  getters: {
    isLogin: (state) => {
      return state.userInfo.token!== ''
    },
    getUserInfo: (state) => {
      return state.userInfo
    }
  },
  actions: {
    setToken(token: string) {
      this.userInfo.token = token
    },
    setUserInfo(userInfo: any) {
      this.userInfo = userInfo
    }
  }
})

