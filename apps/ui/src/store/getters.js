const getters = {
  username: state => state.user.username,
  endpoint: state => state.user.endpoint,
  ca: state => state.user.ca,
  token: state => state.user.token
}

export default getters
