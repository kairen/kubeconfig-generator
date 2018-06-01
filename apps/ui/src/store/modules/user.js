import { loginByUser } from '@/api/login'
import { removeUser } from '@/utils/auth'

const user = {
  state: {
    // User Information
    dn: '',
    endpoint: '',
    ca: '',
    token: ''
  },
  mutations: {
    SET_DN: (state, dn) => {
      state.dn = dn
    },
    SET_ENDPOINT: (state, endpoint) => {
      state.endpoint = endpoint
    },
    SET_CA: (state, ca) => {
      state.ca = ca
    },
    SET_TOKEN: (state, token) => {
      state.token = token
    }
  },
  actions: {
    LoginByUser ({ commit }, userInfo) {
      const dn = userInfo.dn.trim()
      return new Promise((resolve, reject) => {
        loginByUser(dn, userInfo.password).then(response => {
          commit('SET_DN', response.data.username)
          commit('SET_ENDPOINT', response.data.endpoint)
          commit('SET_CA', response.data.ca)
          commit('SET_TOKEN', response.data.token)
          resolve()
        }).catch(error => {
          reject(error.response)
        })
      })
    },
    Logout ({ commit, state }) {
      // Need to call Logout API to destroy token
      return new Promise((resolve, reject) => {
        commit('SET_TOKEN', '')
        removeUser()
        resolve()
      })
    },
    FrontEndLogout ({ commit }) {
      return new Promise(resolve => {
        commit('SET_TOKEN', '')
        removeUser()
        resolve()
      })
    }
  }
}

export default user
