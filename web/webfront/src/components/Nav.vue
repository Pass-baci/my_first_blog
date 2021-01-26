<template>
  <v-card class="mx-auto" max-width="300">
    <v-img src="../assets/ground.jpg">
      <v-card-title>
        <v-col align="center">
          <v-avatar size="130" color="grey">
            <img src="../assets/logo.png" alt />
          </v-avatar>
        </v-col>
      </v-card-title>
      <v-divider></v-divider>
    </v-img>
        <v-card-title>About Me:</v-card-title>
    <v-card-text><span>凡是过去，皆为序章</span></v-card-text>
    <v-card-text><span>程序员的学习之旅</span></v-card-text>

    <v-divider color="indigo"></v-divider>

    <v-list nav dense>
      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="blue darken-2">{{'mdi-qqchat'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text"><span>778281938</span></v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="green darken-2">{{'mdi-wechat'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text"><span>嘘嘘嘘！！！</span></v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="indigo">{{'mdi-email'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text"><span>778281938@qq.com</span></v-list-item-content>
      </v-list-item>
    </v-list>
  </v-card>
</template>
<style lang="">
</style>
