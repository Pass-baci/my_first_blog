<template>
   <a-button type="danger" @click="loginout">
          注销
    </a-button>
</template>

<script>
export default {
  methods: {
    loginout () {
      window.sessionStorage.clear('token')
      this.$router.push('login')
    }
  }
}
</script>
