import {createRouter, createWebHistory} from 'vue-router'

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
            path: '/',
            name: 'web',
            component: () => import('@/views/web/index.vue'),
            children: [
                {
                    path: '',
                    name: 'contact',
                    component: () => import('@/views/web/contact/index.vue'),
                    children: [
                        {
                            path: '',
                            name: 'user_list',
                            component: () => import('@/views/web/contact/user_list.vue')
                        },
                      {
                        path: '/welcome',
                        name: 'welcome',
                        component: () => import('@/views/web/contact/welcome.vue')
                      }
                    ]
                },
                {
                    path: 'info',
                    name: 'info',
                    component: () => import('@/views/web/info/index.vue')
                },
                {
                    path: 'session',
                    name: 'session',
                    component: () => import('@/views/web/session/index.vue')
                },
                {
                    path: 'notice',
                    name: 'notice',
                    component: () => import('@/views/web/notice/index.vue')
                }
            ]
        },
    ]
})

export default router
