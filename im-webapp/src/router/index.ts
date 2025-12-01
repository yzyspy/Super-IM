import {createRouter, createWebHistory} from 'vue-router'
import {useUserInfoStore} from '@/stores/index';

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/login',
            name: 'login',
            component: () => import('@/views/login/index.vue')
        },
        {
            path: '/register',
            name: 'register',
            component: () => import('@/views/register/index.vue')
        },
        {
            path: '',
            name: 'web',
            component: () => import('@/views/web/index.vue'),
            children: [
                {
                    path: '/contact',
                    name: 'contact',
                    component: () => import('@/views/web/contact/index.vue'),
                    children: [
                        {
                            path: '',
                            name: 'user_list',
                            component: () => import('@/views/web/contact/user_list.vue')
                        },
                        {
                          path: 'welcome',//加反斜杠是绝对路径，匹配的是http://localhost:5173/welcome，去掉反斜杠是相对路径，匹配的是http://localhost:5173/contact/welcome
                          name: 'welcome',
                          component: () => import('@/views/web/contact/welcome.vue')
                        },
                        {
                            path: 'user_chat/:id',//对话页面,path需要配置参数
                            name: 'user_chat',
                            props: true,//需要设置为true
                            component: () => import('@/views/web/contact/user_chat.vue'),
                        }
                    ]
                },
                {
                    path: '/info',
                    name: 'info',
                    component: () => import('@/views/web/info/index.vue'),
                    children: [
                        {
                            path: '',
                            name: 'my_info',
                            component: () => import('@/views/web/info/my_info.vue')
                        },
                        {
                            path: 'base_info',
                            name: 'base_info',
                            component: () => import('@/views/web/info/base_info.vue')
                        },
                        {
                            path: 'role_info',
                            name: 'role_info',
                            component: () => import('@/views/web/info/role_info.vue')
                        },
                        {
                            path: 'safe_info',
                            name: 'safe_info',
                            component: () => import('@/views/web/info/safe_info.vue')
                        },
                    ]
                },
                {
                    path: '/session',
                    name: 'session',
                    component: () => import('@/views/web/session/index.vue')
                },
                {
                    path: '/notice',
                    name: 'notice',
                    component: () => import('@/views/web/notice/index.vue')
                }
            ]
        },
    ]
})

// 路由守卫, 判断要跳转的路由是否需要登录，如果需要登录但是没有登录，则跳转到登录页面
router.beforeEach((to, from, next) => {
    let login = useUserInfoStore().isLogin
    console.info('router.beforeEach', to, from, login)
    if (to.name === 'login' || to.name ==='register') {
        // 登录页面和注册页面不需要登录
        next()
    } else {
        let login = useUserInfoStore().isLogin
        if (login) {
            next()
        } else {
            next({name: 'login'})
        }
    }
})

export default router
