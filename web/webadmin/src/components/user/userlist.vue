<template>
    <div>
        <span class="sp">欢迎来到用户管理页面</span>
        <a-card>
            <a-row>
                <a-col :span="6">
                    <a-input-search v-model="username" placeholder="输入用户进行查找" style="width: 100%" @search="getUserList" allowClear />
                </a-col>
                <a-col :span="4">
                    <a-button type="primary"  @click="addUservisible = true" style="margin-left:8%;">新增</a-button>
                </a-col>
            </a-row>
            <a-table rowKey="ID" :columns="columns" :pagination='paginationOption' :dataSource="userlist" bordered>
                <span slot="role" slot-scope="role">{{ role === 1 ? '管理员':'阅读者' }}</span>
                <template slot="action" slot-scope="data">
                    <div class="actionSlot">
                        <a-button type="primary" style="margin-right: 15px" @click="editUser(data.ID)">编辑</a-button>
                        <a-button type="danger" @click="deleteUser(data.ID)">删除</a-button>
                    </div>
                </template>
            </a-table>
        </a-card>
        <a-modal
          closable
          destroyOnclose
          title="新增用户"
          :visible="addUservisible"
          @ok="addUserOk"
          @cancel="addUserCancel"
        >
        <a-form-model :model="userInfo" :rules="userRules" ref="addUserRef">
          <a-form-model-item label="用户名" prop="username">
            <a-input v-model="userInfo.username"></a-input>
          </a-form-model-item>
          <a-form-model-item label="密码" prop="password">
            <a-input-password v-model="userInfo.password"></a-input-password>
          </a-form-model-item>
          <a-form-model-item label="确认密码" prop="repassword">
            <a-input-password v-model="userInfo.repassword"></a-input-password>
          </a-form-model-item>
          <a-form-model-item label="是否为管理员" prop="role">
            <a-select defaultValue="2" style="100px" @change="adminChange">
              <a-select-option key="1" value="1">是</a-select-option>
              <a-select-option key="2" value="2">否</a-select-option>
            </a-select>
          </a-form-model-item>
        </a-form-model>
        </a-modal>
        <a-modal
          closable
          destroyOnclose
          title="编辑用户"
          :visible="editUservisible"
          @ok="editUserOk"
          @cancel="editUserCancel"
        >
        <a-form-model :model="userInfo" :rules="editRules" ref="editUserRef">
          <a-form-model-item label="用户名" prop="username">
            <a-input v-model="userInfo.username"></a-input>
          </a-form-model-item>
          <a-form-model-item label="是否为管理员" prop="role">
            <a-select defaultValue="2" style="100px" @change="adminChange">
              <a-select-option key="1" value="1">是</a-select-option>
              <a-select-option key="2" value="2">否</a-select-option>
            </a-select>
          </a-form-model-item>
        </a-form-model>
        </a-modal>
    </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '5%',
    key: 'ID',
    align: 'center'
  },
  {
    title: '用户名',
    dataIndex: 'username',
    width: '20%',
    key: 'username',
    align: 'center'
  },
  {
    title: '权限',
    dataIndex: 'role',
    width: '10%',
    key: 'role',
    scopedSlots: { customRender: 'role' },
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
      userInfo: {
        id: 0,
        username: '',
        password: '',
        repassword: '',
        role: 2
      },
      userlist: [],
      userRules: {
        username: [
          {
            required: true,
            message: '请输入用户名',
            trigger: 'blur'
          },
          {
            min: 4,
            max: 12,
            message: '用户名长度不正确',
            trigger: 'blur'
          }
        ],
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.password === '') {
                callback(new Error('请输入密码'))
              }
              if ([...this.userInfo.password].length < 6 || [...this.userInfo.password].length > 20) {
                callback(new Error('密码格式错误'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        repassword: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.repassword === '') {
                callback(new Error('请输入密码'))
              }
              if (this.userInfo.repassword !== this.userInfo.password) {
                callback(new Error('密码不一致，请重新输入'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      editRules: {
        username: [
          {
            required: true,
            message: '请输入用户名',
            trigger: 'blur'
          },
          {
            min: 4,
            max: 12,
            message: '用户名长度不正确',
            trigger: 'blur'
          }
        ]
      },
      columns,
      username: '',
      addUservisible: false,
      editUservisible: false
    }
  },
  created () {
    this.getUserList()
  },
  methods: {
    async getUserList () {
      const { data: res } = await this.$http.get('users', {
        params: {
          username: this.username,
          pageSize: this.paginationOption.defaultPageSize,
          pagenum: this.paginationOption.defaultCurrent
        }
      })
      if (res.status !== 200) return this.$message.console.error(res.message)
      this.userlist = res.data
      this.paginationOption.total = res.total
    },
    deleteUser (id) {
      this.$confirm({
        title: '确定要删除该用户吗？',
        content: '一旦删除，无法恢复！',
        onOk: async () => {
          const res = await this.$http.delete(`/users/${id}`)
          if (res.status !== 200) return this.$message.console.error(res.message)
          this.$message.success('删除成功')
          this.getUserList()
        },
        onCancel: () => {
          this.$message.info('已取消！')
        }
      })
    },
    addUserOk () {
      this.$refs.addUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.post('users/add', {
          username: this.userInfo.username,
          password: this.userInfo.password,
          role: this.userInfo.role
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$message.success('添加用户成功！')
        this.addUservisible = false
        this.getUserList()
      })
    },
    addUserCancel () {
      this.$refs.addUserRef.resetFields()
      this.addUservisible = false
    },
    adminChange (value) {
      this.userInfo.role = Number(value)
    },
    async editUser (id) {
      this.editUservisible = true
      const { data: res } = await this.$http.get(`users/${id}`)
      this.userInfo.username = res.data.username
      this.userInfo.id = id
    },
    editUserOk () {
      this.$refs.editUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`users/${this.userInfo.id}`, {
          username: this.userInfo.username,
          role: this.userInfo.role
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$message.success('更新用户成功！')
        this.editUservisible = false
        this.getUserList()
      })
    },
    editUserCancel () {
      this.$refs.editUserRef.resetFields()
      this.editUservisible = false
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
