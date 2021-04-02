<template >
  <div>
    <div class="d-flex justify-center pa-3 ma-1 text-h4 font-weight-bold">{{artInfo.title}}</div>
    <div class="d-flex justify-center align-center">
      <div class="d-flex mx-10 justify-center">
        <v-icon class="mr-1" color="indigo" small>{{'mdi-calendar-month'}}</v-icon>
        <span>{{artInfo.CreatedAt | dateformat('YYYY-MM-DD HH:MM')}}</span>
      </div>
    </div>
    <v-divider class="pa-3 ma-3"></v-divider>
    <v-alert class="ma-4 font-italic caption d-flex justify-center" elevation="1" color="grey lighten-1" dark border="left" outlined>{{artInfo.desc}}</v-alert>
    <div class="content_box">
      <div class="content ma-5 pa-3 text-justify" v-html="artInfo.content"></div>
    </div>
    <v-divider class="ma-5"></v-divider>
  </div>
</template>
<script>
import Prism from '../assets/prism'
export default {
  props: ['id'],
  data () {
    return {
      artInfo: {}
    }
  },
  created () {
    this.getArtInfo()
  },
  methods: {
    // 查询文章
    async getArtInfo () {
      const { data: res } = await this.$http.get(`article/info/${this.id}`)
      this.artInfo = res.data
      this.timer = setTimeout(() => {
        Prism.highlightAll()
      }, 0)
    }
  }
}
</script>
<style scoped>
.content_box {
  max-width: 100%;
}
.content >>> img {
  width: auto;
  max-width: 100%;
}
.content >>> pre,
code {
  margin: 10px;
  padding: 14px;
  overflow: auto;
  font-size: 85%;
  line-height: 1.45;
}
</style>
