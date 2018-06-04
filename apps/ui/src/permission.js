import router from './router'
import store from './store'

const whiteList = ['/login', '/home']

router.beforeEach((to, from, next) => {
  if (to.path === '/home') {
    if (!store.getters.username) {
      next('/login')
    } else {
      next()
    }
  } else {
    if (whiteList.indexOf(to.path) !== -1) {
      next()
    } else {
      next('/login')
    }
  }
})
