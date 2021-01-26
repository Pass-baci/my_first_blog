<template>
    <div>
        <div class="text">
            <p>本页面暂无内容，内容正在更新，点击下方按钮进入文章列表</p>
        </div>
        <div class="btu">
            <a-button @click="$router.push('/admin/artlist')">文章列表</a-button>
        </div>
    </div>
</template>

<style scoped>
.text {
    font-size: 24px;
    margin-top: 10%;
    display: flex;
    justify-content: center;
    align-items: center;
}
.btu {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top:10px;
}
</style>
