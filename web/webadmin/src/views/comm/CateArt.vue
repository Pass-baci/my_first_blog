<template>
  <div>
  <Menus></Menus>
    <a-card style="margin-top: 10px">
        <a-table rowKey="ID" :columns="columns" :pagination="paginationOption" :dataSource="Artlist" style="margin-top: 10px">
          <span slot="CreatedAt" slot-scope="CreatedAt">{{ CreatedAt.substr(0,10) }}</span>
          <span class="titleho" slot="action" slot-scope="data" @click="ChangetoArt(data.ID)">{{ data.title }}</span>
          <span slot="Category2" slot-scope="Category1">{{ Category1 }}</span>
        </a-table>
    </a-card>
  </div>
</template>

<script>
import Menus from '../../components/menu/menu'

const columns = [
  {
    title: '发布时间',
    dataIndex: 'CreatedAt',
    width: '6%',
    key: 'CreatedAt',
    scopedSlots: { customRender: 'CreatedAt' },
    align: 'center'
  },
  {
    title: '标题',
    width: '10%',
    key: 'action',
    scopedSlots: { customRender: 'action' },
    align: 'center'
  },
  {
    title: '简述',
    dataIndex: 'desc',
    width: '20%',
    key: 'desc',
    align: 'center'
  },
  {
    title: '分类',
    dataIndex: 'Category.name',
    width: '10%',
    key: 'Category.name',
    slots: { title: 'Category1' },
    scopedSlots: { customRender: 'Category2' },
    align: 'center'
  }
]
export default {
  components: { Menus },
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
      Artlist: [],
      columns,
      id: 0,
      CateTitle: ''
    }
  },
  created () {
    this.getArtList()
  },
  methods: {
    async getArtList () {
      this.id = this.$route.query.CateId
      this.id = Number(this.id)
      const { data: res } = await this.$http.get(`article/List/${this.id}`, {
        query: {
          pagesize: this.paginationOption.defaultPageSize,
          pagenum: this.paginationOption.defaultCurrent
        }
      })
      if (res.status !== 200) return this.$message.console.error(res.message)
      this.Artlist = res.data
      this.paginationOption.total = res.total
    },
    async ChangetoArt (id) {
      const { data: res } = await this.$http.get(`article/info/${id}`)
      this.id = res.ID
      this.$router.push({
        path: 'ArtOpenCom',
        name: 'ArtOpenCom',
        query: {
          ArtId: id
        }
      })
    }
  }
}
</script>

<style scoped>
.titleho:hover {
  cursor: pointer;
}
</style>
