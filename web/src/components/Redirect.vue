<template></template>

<script>
import axios from 'axios'
export default {
    created() {
        let k = this.$route.params.key;
        this.redirect(k)
    },
    methods: {
        redirect(key) {
            axios({
                method: "get",
                url: "http://127.0.0.1:8080/" + key,
            }).then(response => {
                window.location.href = response.data.msg;
                // console.log()
            }).catch(err => {
                console.log(err)
            })
        }
    }
}
</script>

<style>
</style>