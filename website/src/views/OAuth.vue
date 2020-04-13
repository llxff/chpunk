<template></template>

<script>
    import { OAUTH_REQUEST } from "@/store/actions/auth";
    export default {
        name: "OAuth",
        created() {
            this.$store.dispatch(OAUTH_REQUEST, this.$route.query).then(() => {
                this.$router.push("/");
            });
        }
    };
</script>
