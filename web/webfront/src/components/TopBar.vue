<template>
  <div>
    <v-app-bar app color="#336699">
        <v-avatar class="mx-2" size="40" color="#CCFFFF"><img src="../assets/logo.png" alt /></v-avatar>
        <v-container class="py-0 fill-height">
        <v-btn text color="white" @click="$router.push('/')">首页</v-btn>
        <v-btn
          v-for="item in cateList"
          :key="item.id"
          text
          color="white"
          @click="gotoCate(item.id)"
        >{{item.name}}</v-btn>
    </v-container>
        <v-spacer></v-spacer>

      <v-responsive max-width="50%" class="mr-5">
        <v-text-field
          dense
          flat
          hide-details
          solo-inverted
          rounded
          placeholder="请输入文章标题查找"
          dark
          append-icon="mdi-text-search"
          v-model="searchName"
          @change="searchTitle(searchName)"
        ></v-text-field>
      </v-responsive>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  data () {
    return {
      cateList: [],
      searchName: ''
    }
  },
  created () {
    this.GetCateList()
  },
  methods: {
    // 获取文章分类
    async GetCateList () {
      const { data: res } = await this.$http.get('/category')
      this.cateList = res.data
    },
    searchTitle (title) {
      if (title === '') {
        this.$router.push('/')
      } else {
        this.$router.push(`/Search/${title}`)
      }
    },
    gotoCate (cid) {
      this.$router.push(`/category/${cid}`).catch((err) => err)
    }
  }
}
</script>

<style>

</style>
