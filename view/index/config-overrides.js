const path = require('path')
const { override, addWebpackAlias, fixBabelImports, addWebpackPlugin } = require('customize-cra')
const AntdDayjsWebpackPlugin = require('antd-dayjs-webpack-plugin')
const addLessLoader = require('customize-cra-less-loader')

module.exports = {
  webpack: override(
    // 增加全局 Alias
    addWebpackAlias({
      '@': path.resolve(__dirname, 'src'),
    }),
    // 按需加载 lodash
    fixBabelImports('lodash', {
      libraryDirectory: '',
      camel2DashComponentName: false,
    }),
    fixBabelImports('import', {
      libraryName: 'antd',
      libraryDirectory: 'es',
      style: true,
    }),
    addLessLoader({
      lessLoaderOptions: {
        lessOptions: {
          javascriptEnabled: true,
          modifyVars: {
            '@menu-dark-bg': '#FFFFFF',
            '@collapse-header-padding': '0 0',
          },
        },
      },
    }),
    // 使用 dayjs 替换 antd 中默认的 Moment 日期库
    addWebpackPlugin(new AntdDayjsWebpackPlugin())
  ),
}
