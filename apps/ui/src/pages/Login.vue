<template>
  <v-app>
    <v-container fluid fill-height>
      <v-layout align-center justify-center>
        <v-flex xs12 sm8 md4>
          <v-card>
            <v-toolbar dark color="primary">
              <v-toolbar-title>Kubeconfig Generator</v-toolbar-title>
            </v-toolbar>
            <v-card-text>
              <v-form ref="loginform" v-model="valid" lazy-validation>
                <v-text-field
                  prepend-icon="person"
                  name="login"
                  label="Login"
                  type="text"
                  v-model="username"
                  :rules="loginFormRules.username"
                  required>
                </v-text-field>
                <v-text-field
                  id="password"
                  prepend-icon="lock"
                  name="password"
                  label="Password"
                  type="password"
                  v-model="password"
                  :rules="loginFormRules.passowrd"
                  required></v-text-field>
              </v-form>
            </v-card-text>
            <v-card-actions>
              <v-btn color="primary" v-on:click="login" block>Login</v-btn>
            </v-card-actions>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
    <v-snackbar
      v-model="visiable"
      :timeout="5000"
      :top="true"
      :right="true"
      :multi-line="true">
      Login Failed
      <v-btn flat color="pink" @click.native="visiable = false">Close</v-btn>
    </v-snackbar>
  </v-app>
</template>

<script>
// import axios from 'axios'

export default {
  name: 'Login',
  data () {
    return {
      valid: true,
      visiable: false,
      username: '',
      password: '',
      response: '',
      loginFormRules: {
        username: [v => !!v || 'Username is required'],
        passowrd: [v => !!v || 'Password is required']
      }
    }
  },
  methods: {
    login () {
      if (this.$refs.loginform.validate()) {
        var data = {
          dn: this.username,
          password: this.password
        }
        this.$store.dispatch('LoginByUser', data).then(() => {
          this.$router.push({ path: '/home' })
        }).catch(() => {
          console.log('Failed')
          this.visiable = true
        })
      }
    }
  }
}
</script>
