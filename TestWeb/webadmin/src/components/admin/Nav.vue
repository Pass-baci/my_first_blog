<template>
  <a-layout-sider breakpoint="lg" v-model="collapsed" theme="light">
    <div class="log">
      <span>{{ collapsed ? 'Baby' : 'Hello Baby' }}</span>
    </div>
    <a-menu theme="light" mode="inline" @click="gotoPage">
      <a-menu-item key="index">
        <a-icon type="appstore"/><span>仪表盘</span>
      </a-menu-item>
      <a-sub-menu>
        <span slot="title"><a-icon type="file" /><span>文章管理</span></span>
        <a-menu-item key="addart1">
        <a-icon type="form" /><span>新增文章</span>
        </a-menu-item>
        <a-menu-item key="artlist">
          <a-icon type="unordered-list" /><span>文章列表</span>
        </a-menu-item>
      </a-sub-menu>
      <a-menu-item key="catelist">
        <a-icon type="ordered-list" /><span>分类列表</span>
      </a-menu-item>
      <a-menu-item key="userlist">
        <a-icon type="user"/><span>用户列表</span>
      </a-menu-item>
    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {
  data () {
    return {
      collapsed: false
    }
  },
  methods: {
    gotoPage (item) {
      this.$router.push('/admin/' + item.key).catch((err) => err)
    }
  }
}
</script>

<style scoped>
.log{
    height: 45px;
    width: 77%;
    background-color: #eee;
    margin: 15px;
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>
