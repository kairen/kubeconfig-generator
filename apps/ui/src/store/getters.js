const getters = {
  dn: state => state.user.dn,
  endpoint: state => state.user.endpoint,
  ca: state => state.user.ca,
  token: state => state.user.token
}

export default getters
