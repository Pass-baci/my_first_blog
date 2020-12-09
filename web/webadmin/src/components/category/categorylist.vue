<template>
    <div>
        <span class="sp">欢迎来到分类管理页面</span>
        <a-card>
            <a-row>
                <a-col :span="6">
                    <a-input-search v-model="catename" placeholder="输入分类名称进行查找" style="width: 100%" @search="getCateList" allowClear />
                </a-col>
                <a-col :span="4">
                    <a-button type="primary"  @click="addCatevisible = true" style="margin-left:8%;">新增</a-button>
                </a-col>
            </a-row>
            <a-table rowKey="ID" :columns="columns" :pagination='paginationOption' :dataSource="catelist" bordered>
                <template slot="action" slot-scope="data">
                    <div class="actionSlot">
                        <a-button type="primary" style="margin-right: 15px" @click="editCate(data.id)">编辑</a-button>
                        <a-button type="danger" @click="deleteCate(data.id)">删除</a-button>
                    </div>
                </template>
            </a-table>
        </a-card>
        <a-modal
          closable
          destroyOnclose
          title="新增分类"
          :visible="addCatevisible"
          @ok="addCateOk"
          @cancel="addCateCancel"
        >
        <a-form-model :model="cateInfo" :rules="cateRules" ref="addCateRef">
          <a-form-model-item label="分类名" prop="catename">
            <a-input v-model="cateInfo.catename"  v-on:keyup.enter="addCateOk"></a-input>
          </a-form-model-item>
        </a-form-model>
        </a-modal>
        <a-modal
          closable
          destroyOnclose
          title="编辑分类"
          :visible="editCatevisible"
          @ok="editCateOk"
          @cancel="editCateCancel"
        >
        <a-form-model :model="cateInfo" :rules="cateRules" ref="addCateRef">
          <a-form-model-item label="分类名称" prop="catename">
            <a-input v-model="cateInfo.catename"></a-input>
          </a-form-model-item>
        </a-form-model>
        </a-modal>
    </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    width: '5%',
    key: 'ID',
    align: 'center'
  },
  {
    title: '分类名称',
    dataIndex: 'name',
    width: '20%',
    key: 'catename',
    align: 'center'
  },
  {
    title: '操作',
    width: '10%',
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
      cateInfo: {
        id: 0,
        catename: ''
      },
      catelist: [],
      cateRules: {
        catename: [
          {
            required: true,
            message: '请输入分类名称',
            trigger: 'blur'
          }
        ]
      },
      columns,
      catename: '',
      addCatevisible: false,
      editCatevisible: false
    }
  },
  created () {
    this.getCateList()
  },
  methods: {
    async getCateList () {
      const { data: res } = await this.$http.get('category', {
        params: {
          pageSize: this.paginationOption.defaultPageSize,
          pageNum: this.paginationOption.defaultCurrent,
          catename: this.catename
        }
      })
      if (res.status !== 200) return this.$message.console.error(res.message)
      this.catelist = res.data
      this.paginationOption.total = res.total
    },
    deleteCate (id) {
      this.$confirm({
        title: '确定要删除该分类吗？',
        content: '一旦删除，无法恢复！',
        onOk: async () => {
          const res = await this.$http.delete(`/category/${id}`)
          if (res.status !== 200) return this.$message.console.error(res.message)
          this.$message.success('删除成功')
          this.getCateList()
        },
        onCancel: () => {
          this.$message.info('已取消！')
        }
      })
    },
    addCateOk () {
      this.$refs.addCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.post('category/add', {
          name: this.cateInfo.catename
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$message.success('添加分类成功！')
        this.addCatevisible = false
        this.getCateList()
      })
    },
    addCateCancel () {
      this.$refs.addCateRef.resetFields()
      this.addCatevisible = false
    },
    async editCate (id) {
      this.editCatevisible = true
      const { data: res } = await this.$http.get(`category/${id}`)
      this.cateInfo.catename = res.data.name
      this.cateInfo.id = id
    },
    editCateOk () {
      this.$refs.addCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`category/${this.cateInfo.id}`, {
          name: this.cateInfo.catename
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$message.success('更新分类成功！')
        this.editCatevisible = false
        this.getCateList()
      })
    },
    editCateCancel () {
      this.$refs.addCateRef.resetFields()
      this.editCatevisible = false
      this.$message.info('已取消！')
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
    margin-top: 1%;
    font-size: 20px;
    display: flex;
    justify-content: center;
}

</style>
