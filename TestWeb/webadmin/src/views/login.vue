<template>
    <div class="container">
        <div class="loginBox">
            <a-form-model ref="loginFormRef" :rules="rules" :model="formdata" class="loginForm">
                <a-form-model-item  prop="username">
                    <a-input v-model="formdata.username" placeholder="请输入用户名">
                        <a-icon slot="prefix" type="user" style="color:rgba(0,0,0,.25)"/>
                    </a-input>
                </a-form-model-item>
                <a-form-model-item  prop="password">
                    <a-input v-model="formdata.password" v-on:keyup.enter="login" type="password" placeholder="请输入密码">
                        <a-icon slot="prefix" type="lock" style="color:rgba(0,0,0,.25)"/>
                    </a-input>
                </a-form-model-item>
                <a-form-model-item class="loginBtn">
                    <a-button type="primary" style="margin:10px" @click="login">登录</a-button>
                    <a-button style="margin:10px" @click="resetForm">取消</a-button>
                </a-form-model-item>
            </a-form-model>
        </div>
    </div>
</template>

<script>
export default {
  data () {
    return {
      formdata: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 4, max: 12, message: '用户名应在4-12个字符之间', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 20, message: '密码长度应为6-20个字符', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    resetForm () {
      this.$refs.loginFormRef.resetFields()
    },
    login () {
      this.$refs.loginFormRef.validate(async vaild => {
        if (!vaild) return this.$message.error('用户名或密码错误')
        const { data: res } = await this.$http.post('login', this.formdata)
        if (res.status !== 200) return this.$message.error(res.message)
        window.sessionStorage.setItem('token', res.token)
        this.$router.push('admin/Index')
      })
    }
  }
}
</script>

<style scoped>
.container {
    height: 100%;
    background-image: url('../assets/br.jpg');
    background-repeat:no-repeat;
    background-size:100% 100%
}
.loginBox {
    width: 400px;
    height: 300px;
    min-height: 200px;
    background-color: #fafafa;
    position: absolute;
    left:50%;
    top:50%;
    transform: translate(-50%, -50%);
    border: 1px dashed #e9e9e9;
    border-radius: 6px
}
.loginForm {
    width: 100%;
    position: absolute;
    bottom: 10%;
    padding: 0 20px;
    box-sizing: border-box;
}
.loginBtn {
    padding: 0 10px;
    float: right;
}
</style>
