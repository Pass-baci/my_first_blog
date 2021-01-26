<template>
  <div>
  <Menus></Menus>
    <a-card style="margin-top: 10px">
      <a-row>
        <a-col :span="6" style="width: 100%">
          <a-input-search v-model="arttitle" placeholder="搜索标题" style="width: 20%; float:right" @search="getArtList" allowClear />
        </a-col>
      </a-row>
        <a-table rowKey="ID" :columns="columns" :pagination="paginationOption" :dataSource="Artlist" style="margin-top: 10px" :customRow="topathclick" :rowClassName="rowClass">
          <span slot="UpdatedAt" slot-scope="UpdatedAt">{{ UpdatedAt.substr(0,10) }}</span>
          <span slot="Category2" slot-scope="Category1">{{ Category1 ? Category1 : '暂无分类' }}</span>
        </a-table>
    </a-card>
  </div>
</template>

<script>
import Menus from '../../components/menu/menu'

const columns = [
  {
    title: '更新时间',
    dataIndex: 'UpdatedAt',
    width: '6%',
    key: 'UpdatedAt',
    scopedSlots: { customRender: 'UpdatedAt' },
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
    align: 'center',
    className: ''
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
      arttitle: '',
      spinningOk: true
    }
  },
  created () {
    this.spinningOk = true
    this.getArtList()
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
      this.Artlist = res.data
      this.paginationOption.total = res.total
      this.spinningOk = false
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
    },
    topathclick (record, index) {
      return {
        on: {
          click: () => {
            this.id = record.ID
            this.$router.push({
              path: 'ArtOpenCom',
              name: 'ArtOpenCom',
              query: {
                ArtId: this.id
              }
            })
          }
        }
      }
    },
    rowClass (record, index) {
      const className = 'rowclass1'
      return className
    }
  }
}
</script>

<style>
.rowclass1 {
  cursor: pointer;
}
</style>
