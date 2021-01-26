<template>
    <div>
        <span class="sp">欢迎来到文章管理页面</span>
        <a-card>
            <a-row>
                <a-col :span="6">
                    <a-input-search v-model="arttitle" placeholder="输入文章标题进行查找" style="width: 100%" @search="getArtList" allowClear />
                </a-col>
                <a-col :span="3">
                    <a-button type="primary" @click="$router.push('/addart1')" style="margin-left:8%;">新增文章</a-button>
                </a-col>
            </a-row>
            <a-table rowKey="ID" :columns="columns" :pagination='paginationOption' :dataSource="artlist" bordered>
                <span slot="CreatedAt" slot-scope="CreatedAt">{{ CreatedAt.substr(0,10) }}</span>
                <template slot="action" slot-scope="data">
                    <div class="actionSlot">
                        <a-button type="primary" style="margin-right: 15px" @click="$router.push(`addart/${data.ID}`)">编辑</a-button>
                        <a-button type="danger" @click="deleteArt(data.ID)">删除</a-button>
                    </div>
                </template>
            </a-table>
        </a-card>
    </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '3%',
    key: 'ID',
    align: 'center'
  },
  {
    title: '发布时间',
    dataIndex: 'CreatedAt',
    width: '6%',
    key: 'CreatedAt',
    scopedSlots: { customRender: 'CreatedAt' },
    align: 'center'
  },
  {
    title: '分类名称',
    dataIndex: 'Category.name',
    width: '8%',
    key: 'Categoryname',
    scopedSlots: { customRender: 'Categoryname' },
    align: 'center'
  },
  {
    title: '文章标题',
    dataIndex: 'title',
    width: '15%',
    key: 'arttitle',
    align: 'center'
  },
  {
    title: '操作',
    width: '8%',
    key: 'action',
    scopedSlots: { customRender: 'action' },
    align: 'center'
  }
]
export default {
  data () {
    return {
      paginationOption: {
        pagSizeOptions: ['5', '5', '10'],
        defaultCurrent: 1,
        defaultPageSize: 10,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
        onchage: (page, pageSize) => {
          this.paginationOption.defaultCurrent = page
          this.paginationOption.defaultPageSize = pageSize
        },
        onshowSizeChange: (current, size) => {
          this.paginationOption.defaultCurrent = current
          this.paginationOption.defaultPageSize = size
        }
      },
      artInfo: {
        id: 0,
        arttitle: '',
        artcate: '',
        artcontent: ''
      },
      artlist: [],
      Catelist: [],
      artRules: {
        arttitle: [
          {
            required: true,
            message: '请输入标题',
            trigger: 'blur'
          },
          {
            max: 100,
            message: '标题长度过长',
            trigger: 'blur'
          }
        ],
        artcontent: [
          {
            required: true,
            message: '请输入标题',
            trigger: 'blur'
          }
        ]
      },
      columns,
      arttitle: '',
      addArtvisible: false,
      editArtvisible: false
    }
  },
  created () {
    this.getArtList()
    this.getCateList()
  },
  methods: {
    async getArtList () {
      const { data: res } = await this.$http.get('article', {
        params: {
          title: this.arttitle,
          pageSize: this.paginationOption.defaultPageSize,
          pagenum: this.paginationOption.defaultCurrent
        }
      })
      if (res.status !== 200) return this.$message.console.error(res.message)
      this.artlist = res.data
      this.paginationOption.total = res.total
    },
    deleteArt (id) {
      this.$confirm({
        title: '确定要删除该文章吗？',
        content: '一旦删除，无法恢复！',
        onOk: async () => {
          const res = await this.$http.delete(`/article/${id}`)
          if (res.status !== 200) return this.$message.console.error(res.message)
          this.$message.success('删除成功')
          this.getArtList()
        },
        onCancel: () => {
          this.$message.info('已取消！')
        }
      })
    },
    async getCateList () {
      const { data: res } = await this.$http.get('category')
      if (res.status !== 200) return this.$message.console.error(res.message)
      this.Catelist = res.data
      this.paginationOption.total = res.total
    }
  }
}
</script>

<style scoped>
.actionSlot {
    display: flex;
    justify-content: center;
}
.sp {
    font-size: 20px;
    margin-top: 1%;
    display: flex;
    justify-content: center;
}

</style>
