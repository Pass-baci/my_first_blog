<template>
    <div>
        <Menus></Menus>
        <div class="cate_sy">
            <p class="catename_sy" v-for="item in Categorylist" :key="item.id" :value="item.id" @click="toCateArtinfo(item.id)">{{ item.name }}</p>
        </div>
    </div>
</template>

<script>
import Menus from '../../components/menu/menu'
export default {
  components: { Menus },
  data () {
    return {
      Categorylist: [],
      CateId: 0
    }
  },
  created () {
    this.getCateList()
  },
  methods: {
    async getCateList () {
      const { data: res } = await this.$http.get('category')
      if (res.status !== 200) return this.$message.console.error(res.message)
      this.Categorylist = res.data
    },
    async toCateArtinfo (id) {
      const { data: res } = await this.$http.get(`article/List/${id}`)
      this.id = res.ID
      this.$router.push({
        path: 'CateArt',
        name: 'CateArt',
        query: {
          CateId: id
        }
      })
    }
  }
}
</script>

<style scoped>
.cate_sy {
    margin-top: 10%;
}
.catename_sy {
    font-family: "PingFang SC", "Lantinghei SC", "Microsoft Yahei", "Hiragino Sans GB", "Microsoft Sans Serif", "WenQuanYi Micro Hei", Helvetica, sans-serif;
    font-size: 25px;
    font-weight:bold;
    margin-left: 20%;
    margin-right: 20%;
    border-bottom: 2px dashed;
    line-height: 2;
    display: flex;
    justify-content: center;
}
.catename_sy:hover {
    cursor: pointer;
}
</style>
