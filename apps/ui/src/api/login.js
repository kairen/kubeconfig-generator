import axios from 'axios'

export function loginByUser (username, password) {
  var data = {
    dn: username,
    password: password
  }
  return axios.post(
    'login',
    data
  )
}
