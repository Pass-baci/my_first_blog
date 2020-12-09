<template>
    <div class="menu">
        <router-link to="/" class="amenu" style="color: black; margin-left: 7%">Paradise</router-link>
        <router-link to="/Article" class="amenu" style="margin-left: 10%">文章</router-link>
        <router-link to="/CategoryCom" class="amenu" style="margin-left: 10%">分类</router-link>
        <div class="menu_icon"><a-icon type="github" :style="{ fontSize: '25px' }" @click="toRescue" /></div>
    </div>
</template>

<script>
export default {
  data () {
    return {
      url: 'https://github.com/Pass-baci'
    }
  },
  methods: {
    toRescue () {
      var tempwindow = window.open('_blank')
      tempwindow.location = this.url
    }
  }
}
</script>

<style scoped>
.menu {
  width: 100%;
  height: 7%;
  display: flex;
  align-items: center;
  border-bottom-style: solid;
  box-shadow: #74777d 0px 0px 1px 1px;
}
.amenu {
  color: #5A5A5A;
  font-family: Microsoft YaHei;
  font-size: 27px;
  font-weight: 700;
}
a:hover {
    color: #5A5A5A;
}
.menu_icon {
    position: absolute;
    right: 7%;
    top: 1.5%;
    bottom: 1%;
}
.menu_icon:hover {
    cursor: pointer;
}
</style>
