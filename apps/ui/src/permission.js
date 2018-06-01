import router from './router'

const whiteList = ['/login', '/home']

router.beforeEach((to, from, next) => {
  if (whiteList.indexOf(to.path) !== -1) {
    next()
  } else {
    next('/login')
  }
})
