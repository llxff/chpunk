<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
import axios from "axios"
import { AUTH_LOGOUT } from "@/store/actions/auth";

export default {
  name: "app",
  created() {
    axios.defaults.headers.common['Authorization'] = this.$store.state.auth.token;

    axios.interceptors.response.use(undefined, function (err) {
      return new Promise(function (resolve, reject) {
        if (err.status === 401 && err.config && !err.config.__isRetryRequest) {
          this.$store.dispatch(AUTH_LOGOUT)
        }
        throw err;
      });
    });
  }
}
</script>

<style lang="scss">
  div#app {
    padding: 0 15px;
    margin-top: 5px;
  }
</style>
