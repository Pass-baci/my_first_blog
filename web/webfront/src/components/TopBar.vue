<template>
  <div>
    <v-app-bar app color="#336699">
        <v-avatar class="mx-0" size="40" color="#CCFFFF" @click="toRescue"><img src="../assets/logo.png" alt /></v-avatar>
        <v-container class="py-0 fill-height">
          <v-btn text color="white" size="90%" @click="$router.push('/')">首页</v-btn>
          <v-btn
            v-for="item in cateList.slice(0, 1)"
            :key="item.id"
            text
            color="white"
            @click="gotoCate(item.id)"
            size="90%"
          >{{item.name}}</v-btn>
          <v-menu bottom origin="center center" transition="scale-transition" >
            <template v-slot:activator="{ on, attrs }">
              <v-btn
                text
                color="white"
                dark
                v-bind="attrs"
                v-on="on"
              >
                所有标签
              </v-btn>
            </template>

            <v-list color="#336699">
              <v-list-item
                v-for="item in cateList"
                :key="item.id"
                text
                color="white"
              >
          <v-btn color="white" text @click="gotoCate(item.id)">{{ item.name }}</v-btn>
        </v-list-item>
      </v-list>
    </v-menu>
        </v-container>
        <v-spacer></v-spacer>

      <v-responsive max-width="20%" class="mr-2">
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
      searchName: '',
      url: 'https://github.com/Pass-baci'
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
    },
    toRescue () {
      var tempwindow = window.open('_blank')
      tempwindow.location = this.url
    }
  }
}
</script>

<style>

</style>
