const CompressionWebpackPlugin = require('compression-webpack-plugin')
const productionGzipExtensions = ['js', 'css']

module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  assetsDir: 'static',
  productionSourceMap: false,
  configureWebpack: {
    plugins: [
        new CompressionWebpackPlugin({
            // asset: '[path].gz[query]',   // 提示compression-webpack-plugin@3.0.0的话asset改为filename
            algorithm: 'gzip',
            test: new RegExp('\\.(' + productionGzipExtensions.join('|') + ')$'),
            threshold: 10240,
            minRatio: 0.8
        })
    ]
  },
  chainWebpack: config => {
    config.plugin('html').tap(args => {
      args[0].title = '欢迎来到Paradise!'
      return args
    })
  }
}
