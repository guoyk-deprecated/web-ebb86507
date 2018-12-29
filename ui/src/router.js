import VueRouter from 'vue-router'

import Index from './components/Index'

export default new VueRouter({
    routes: [
        {
            path: '/',
            component: Index
        }
    ]
})