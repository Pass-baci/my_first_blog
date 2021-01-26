<template>
    <div class="index">
        <h1>Paradise Baci</h1><br>
        <ul>
            <li><router-link to="/Article">文章</router-link></li>
            <li><router-link to="/CategoryCom">分类</router-link></li>
        </ul>
        <span class="timeDate">生命在流逝：{{ date }}</span>
    </div>
</template>

<script>
import dayjs from 'dayjs'
export default {
  data () {
    return {
      date: new Date()
    }
  },
  mounted () {
    const _this = this
    this.timer = setInterval(() => {
      _this.date = dayjs().format('YYYY-MM-DD HH:mm:ss')
    }, 10)
  },
  beforeDestroy () {
    if (this.timer) {
      clearInterval(this.timer)
    }
  }
}
</script>

<style scoped>
h1, span {
    font-family: Microsoft YaHei;
    font-weight: 700;
    text-align: center;
}
li {
    text-align: center;
    letter-spacing: 10px;
    margin-top: 10px;
    margin-right: 5px;
}
.index {
    width: 100%;
    height: 100%;
    display: flex;
    flex-grow: 1;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    margin: auto;
}
.timeDate {
    position: absolute;
    right: 1%;
    bottom: 2%;
}
a {
    color: #5A5A5A;
    font-family: Microsoft YaHei;
    font-size: 25px;
    font-weight: 700;
}
a:hover {
    color: #5A5A5A;
}

</style>
