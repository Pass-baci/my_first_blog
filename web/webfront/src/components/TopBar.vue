<template>
  <div>
<v-app-bar mobileBreakpoint="sm" app dark flat color="#336699">
      <v-app-bar-nav-icon dark class="hidden-md-and-up" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>

      <v-tabs dark class="hidden-sm-and-down ml-10">
        <v-tab @click="$router.push('/')">首页</v-tab>
        <v-tab
          v-for="item in cateList"
          :key="item.id"
          text
          @click="gotoCate(item.id)"
        >{{ item.name }}</v-tab>
      </v-tabs>

      <v-spacer></v-spacer>

      <v-responsive class="hidden-sm-and-down" color="white">
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
      <v-toolbar-title>
        <v-app-bar-nav-icon class="mx-5">
          <v-avatar size="40" color="grey" @click="toRescue">
            <img src="https://xx.mzqcgc.cn/logo.png" alt />
          </v-avatar>
        </v-app-bar-nav-icon>
      </v-toolbar-title>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" color="#336699" dark app temporary>
      <v-list>
        <v-list-item-title>
          <v-btn href="/" dark text>
            <v-icon small>mdi-home</v-icon>首页
          </v-btn>
        </v-list-item-title>

        <v-list-item
          v-model="group"
          active-class="deep-purple--text text--accent-4"
          v-for="item in cateList"
          :key="item.id"
        >
          <v-list-item-title>
            <v-btn dark text @click="gotoCate(item.id)">{{ item.name }}</v-btn>
          </v-list-item-title>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>

<script>
export default {
  data () {
    return {
      drawer: false,
      group: null,
      valid: true,
      registerformvalid: true,
      cateList: [],
      searchName: '',
      dialog: false,
      url: 'https://github.com/Pass-baci'
    }
  },
  watch: {
    group () {
      this.drawer = false
    }
  },
  created () {
    this.GetCateList()
  },
  methods: {
    // 获取分类
    async GetCateList () {
      const { data: res } = await this.$http.get('category')
      this.cateList = res.data
    },
    // 查找文章标题
    searchTitle (title) {
      if (title.length === 0) return this.$message.error('空空如也，请填入内容喔！')
      this.$router.push(`/search/${title}`)
    },
    gotoCate (cid) {
      this.$router.push(`/category/${cid}`).catch((err) => err)
    },
    // logo跳转
    toRescue () {
      var tempwindow = window.open('_blank')
      tempwindow.location = this.url
    }
  }
}
</script>

<style></style>
