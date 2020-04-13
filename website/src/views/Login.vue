<template>
  <div class="container">
    <div class="row">
      <div class="col-sm"></div>
      <div class="col-sm">
        <button class="btn btn-danger btn-block" v-on:click="login">Sign in with <b>Google</b></button>
      </div>
      <div class="col-sm"></div>
    </div>
  </div>
</template>

<script>
  import { AUTH_REQUEST } from "@/store/actions/auth";
  export default {
    name: "Login",
    methods: {
      login: function() {
        this.$store.dispatch(AUTH_REQUEST).then(() => {
          this.$router.push("/");
        });
      }
    }
  };
</script>
