<template>
    <div>
        <a-card>
            <h3>{{ artInfo.id ? '更新' : '新增' }}文章页面</h3>
            <a-form-model :model="artInfo" ref="artInfoRef" :rules="artInfoRules" :hideRequiredMark="true">
                <a-form-model-item label="文章标题" prop="title">
                    <a-input style="width:300px" v-model="artInfo.title"></a-input>
                </a-form-model-item>
                <a-form-model-item label="文章分类" prop="cid">
                    <a-select style="width:200px" v-model="artInfo.cid" placeholder="请选择分类" @change="cateChange">
                        <a-select-option v-for="item in Catelist" :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
                    </a-select>
                </a-form-model-item>
                <a-form-model-item label="文章描述" prop="desc">
                    <a-input type="textarea" v-model="artInfo.desc"></a-input>
                </a-form-model-item>
                <a-form-model-item label="文章缩略图" prop="img">
                  <a-upload
                    listType="picture"
                    :defaultFileList="fileList"
                    name="file"
                    :action="upUrl"
                    :headers="headers"
                    @change="upChange"
                  >
                    <a-button>
                      <a-icon type="upload" />点击上传
                    </a-button>
                    <br/>
                    <br/>
                    <template v-show="id">
                      <img :src="artInfo.img" style="width:120px;height:100px" />
                    </template>
                  </a-upload>
                </a-form-model-item>
                <a-form-model-item label="文章内容" prop="Content">
                    <Editor style="margin: 0; padding: 0" v-model="artInfo.content"></Editor>
                </a-form-model-item>
                <a-form-model-item>
                    <a-button type="danger" style="margin-right: 15px" @click="artOk(artInfo.id)">{{ artInfo.id ? '更新' : '提交' }}</a-button>
                    <a-button type="primary" @click="addcancle">取消</a-button>
                </a-form-model-item>
            </a-form-model>
        </a-card>
    </div>
</template>

<script>
import { Url } from '../../plugin/http'
import Editor from '../editor/index'
export default {
  components: { Editor },
  props: ['id'],
  data () {
    return {
      artInfo: {
        id: 0,
        title: '',
        cid: undefined,
        desc: '',
        content: '',
        img: ''
      },
      Catelist: [],
      upUrl: Url + 'upload',
      headers: '',
      fileList: [],
      artInfoRules: {
        title: [{
          required: true,
          message: '请输入标题',
          trigger: 'blur'
        }],
        cid: [{
          required: true,
          message: '请选择分类',
          trigger: 'blur'
        }],
        desc: [{
          required: true,
          message: '请输入文章描述',
          trigger: 'blur'
        },
        {
          max: 120,
          message: '输入字符不可超过120个',
          trigger: 'change'
        }],
        content: [{
          required: true,
          message: '请输入文章内容',
          trigger: 'blur'
        }]
      }
    }
  },
  created () {
    this.getCateList()
    this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
    if (this.id) {
      this.getArtInfo(this.id)
    }
  },
  methods: {
    async getArtInfo (id) {
      const { data: res } = await this.$http.get(`article/info/${id}`)
      if (res.status !== 200) return this.$message.console.error(res.message)
      this.artInfo = res.data
      this.artInfo.id = res.data.ID
    },
    async getCateList () {
      const { data: res } = await this.$http.get('category')
      if (res.status !== 200) return this.$message.console.error(res.message)
      this.Catelist = res.data
    },
    cateChange (value) {
      this.artInfo.cid = value
    },
    upChange (info) {
      if (info.file.status !== 'uploading') {
      }
      if (info.file.status === 'done') {
        this.$message.success('图片上传成功')
        const imgUrl = info.file.response.url
        this.artInfo.img = imgUrl
      } else if (info.file.status === 'error') {
        this.$message.error('图片上传失败')
      }
    },
    artOk (id) {
      this.$refs.artInfoRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数验证未通过，请按要求录入文章')
        if (id === 0) {
          const { data: res } = await this.$http.post('article/add', this.artInfo)
          if (res.status !== 200) return this.$message.console.error(res.message)
          this.$router.push('/admin/artlist')
          this.$message.success('添加文章成功')
        } else {
          const { data: res } = await this.$http.put(`article/${id}`, this.artInfo)
          if (res.status !== 200) return this.$message.console.error(res.message)
          this.$router.push('/admin/artlist')
          this.$message.success('更新文章成功')
        }
      })
    },
    addcancle () {
      this.$refs.artInfoRef.resetFields()
      this.$router.push('/admin/artlist')
    }
  }
}
</script>
