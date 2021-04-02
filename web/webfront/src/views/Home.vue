<template>
<div class="main">
 <v-app app>
    <TopBar></TopBar>

    <v-main class="grey lighten-3">
      <v-container>
        <v-row>
          <v-col cols="12" md="7">
            <v-sheet min-height="80vh" rounded="lg">
              <router-view :key="$route.path"></router-view>
            </v-sheet>
          </v-col>
          <v-col cols="12" md="5">
            <Nav></Nav>
          </v-col>
        </v-row>
      </v-container>
    </v-main>

    <Footer></Footer>
  </v-app>
</div>
</template>

<script>
import TopBar from '../components/TopBar'
import Footer from '../components/Footer'
import Nav from '../components/Nav'
export default {
  components: { TopBar, Footer, Nav }
}
</script>
