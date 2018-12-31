import VueRouter from 'vue-router'

import Index from './components/Index'
import Admin from './components/Admin'

export default new VueRouter({
    routes: [
        {
            path: '/',
            component: Index
        },
        {
            path: '/admin',
            component: Admin
        }
    ]
})