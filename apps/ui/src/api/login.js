import axios from 'axios'

export function loginByUser (dn, password) {
  var data = {
    dn: dn,
    password: password
  }
  return axios.post(
    'login',
    data
  )
}
