import { ref, computed } from 'vue'
import { defineStore } from 'pinia'


export const useUserInfoStore = defineStore('userInfo', {
  // 为了完整类型推理，推荐使用箭头函数
  state: () => ({
    token: '',
    userInfo: {
      name: ''
    }
  }),
  getters: {
    isLogin: (state) => {
      return state.token!== ''
    },
    getUserInfo: (state) => {
      return state.userInfo
    }
  },
  actions: {
    setToken(token: string) {
      this.token = token
    },
    setUserInfo(userInfo: any) {
      this.userInfo = userInfo
    }
  }
})

