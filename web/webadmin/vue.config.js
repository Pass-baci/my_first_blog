const CompressionWebpackPlugin = require('compression-webpack-plugin')
const productionGzipExtensions = ['js', 'css']

module.exports = {
    publicPath: '/admin/',
    outputDir: 'dist',
    // 警告 webpack 的性能提示
    configureWebpack: {
      performance: {
        hints: 'warning',
        // 入口起点的最大体积 整数类型（以字节为单位）
        maxEntrypointSize: 50000000,
        // 生成文件的最大体积 整数类型（以字节为单位 300k）
        maxAssetSize: 30000000,
        // 只给出 js 文件的性能提示
        assetFilter: function(assetFilename) {
          return assetFilename.endsWith('.js')
        }
      }
    },
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