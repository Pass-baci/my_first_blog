<template>
  <a-layout class="contain">
     <Nav></Nav>
    <a-layout>
      <a-layout-header class="headBtn">
        <Header></Header>
      </a-layout-header>
      <a-layout-content>
        <router-view :key="$route.path"></router-view>
      </a-layout-content>
      <a-layout-footer>
          <Footer></Footer>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
import Nav from '../components/admin/Nav'
import Footer from '../components/admin/Footer'
import Header from '../components/admin/Header'
export default {
  components: { Nav, Footer, Header }
}
</script>

<style scoped>
.contain{
    height: 100%;
}
.headBtn{
    background-color: #fff;
    justify-content: flex-end;
    display: flex;
    align-items: center;
    border-bottom-style: solid;
    box-shadow: #74777d 0px 0px 1px 1px;
}
</style>
