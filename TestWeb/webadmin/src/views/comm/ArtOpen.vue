<template>
  <div>
    <Menus></Menus>
    <div class="titlestyle"><h2>{{ ArtInfo.title }}</h2></div>
    <span class="timestyle">发布时间：{{ ArtInfo.time.substr(0,10) }}</span>
    <div class="contentstyle" v-html="ArtInfo.content"></div>
  </div>
</template>
<script>
import Menus from '../../components/menu/menu'

export default {
  components: { Menus },
  data () {
    return {
      ArtInfo: {
        title: '',
        time: '',
        content: ''
      },
      id: 0
    }
  },
  created () {
    this.getArtinfo()
  },
  adminChange (value) {
    this.id = Number(value)
  },
  methods: {
    async getArtinfo () {
      this.id = this.$route.query.ArtId
      this.id = Number(this.id)
      const { data: res } = await this.$http.get(`article/info/${this.id}`)
      this.ArtInfo.title = res.data.title
      this.ArtInfo.time = res.data.CreatedAt
      this.ArtInfo.content = res.data.content
    }
  }
}
</script>

<style scoped>
.titlemove:hover {
  cursor: pointer;
}
.titlestyle {
    font-family: Avenir,-apple-system,BlinkMacSystemFont,Segoe UI,PingFang SC,Hiragino Sans GB,Microsoft YaHei,Helvetica Neue,Helvetica,Arial,sans-serif,Apple Color Emoji,Segoe UI Emoji,Segoe UI Symbol;
    font-size: 20px;
    color: #0d1a26;
    font-variant: tabular-nums;
    line-height: 38px;
    margin-top: 5%;
    display: flex;
    justify-content: center;
}
.contentstyle {
    font-family: "PingFang SC", "Lantinghei SC", "Microsoft Yahei", "Hiragino Sans GB", "Microsoft Sans Serif", "WenQuanYi Micro Hei", Helvetica, sans-serif;
    font-weight: 400;
    font-size: 17px;
    line-height: 30px;
    color: #353535;
    letter-spacing: 5px;
    text-indent: 46px;
    margin: 5% 5% 5% 5%;
}
.timestyle {
  position: absolute;
  right: 10%;
}

</style>
