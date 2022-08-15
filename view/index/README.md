# 农晓信知识学院管理后台项目

该项目是基于 [Create React App](https://github.com/facebook/create-react-app) 进行二次开发的。

## 指令列表

在当前项目中，您可以运行：

### `yarn start`

在开发模式下运行应用程序。\
打开 [http://localhost:3000](http://localhost:3000) 在浏览器中查看。

如果您进行编辑，页面将重新加载。\
您还将在控制台中看到所有的 Lint 错误。

### `yarn test`

以交互模式启动测试运行器。\
有关详细信息，请参阅有关 [运行测试](https://facebook.github.io/create-react-app/docs/running-tests) 的部分。

### `yarn build`

将用于生产的应用程序构建到 `build` 文件夹。\
它在生产模式下正确捆绑 React 并优化构建以获得最佳性能。构建文件将被压缩，文件名包括哈希。

有关详细信息，请参阅有关 [部署](https://facebook.github.io/create-react-app/docs/deployment) 的部分。

### `yarn lint`

将自动检查项目中的错误并尝试修复。\
在构建生产项目前需要运行该脚本可以修复潜在的错误。

### `yarn analyze`

** 只能在 `yarn build` 指令之后运行 **
用于分析项目中的依赖构成及文件大小，有利于项目优化。

## 了解更多

您可以在 [创建 React 应用程序文档](https://facebook.github.io/create-react-app/docs/getting-started) 中了解更多信息。

要学习 React，请查看 [React 文档](https://reactjs.org/)。

## 项目结构

```txt
- src
  - assets 资源文件，用于放置资源类文件，比如图片、csv这类非脚本文件
  - components 全局组件，一些全局同用类的组件可以放置在这个目录下
  - helper 工具类脚本
  - pages 项目页面主目录
    - account 账户管理相关页面
      - admin 管理员页面
      - agency 机构管理页面
      - department 组织管理页面
      - role 角色管理页面
      - user 用户管理页面
    - biz 业务板块相关页面
      - course 课程管理相关页面
        - resource 资源相关页面
    - dashboard 数据看板页面
    - errors 错误类页面
    - login 登录页面
    - operation 运营板块相关页面
      - tag 知识分类相关页面
        - components 页面相关组件
  - services 接口类文件
  - state 全局状态
  - styles 全局样式变量
  - config.ts 静态配置项
  - constant.ts 静态变量
  - global.less 全局样式
  - index.tsx 入口文件
  - react-app-env.d.ts TypeScript 定义，全局定义或者覆盖 types/nxkb 定义
  - router.tsx 路由注册，用于页面注册，目前页面统一采用懒加载模式，以减少网站的初次加载时间
  - setupTests.ts 测试脚本入口，暂时未使用
- config-overrides.js 用于覆盖底层的 webpack 配置项
```

### 主要依赖

```txt
"@ant-design/pro-card": "^1.20.4"       Antd Pro 卡片组件
"@ant-design/pro-list": "^1.21.62"      Antd Pro 列表组件
"@ant-design/pro-table": "^2.74.1"      Antd Pro 高级表格组件
"@meemo/lamp": "^0.6.1"                 一些简单的通用脚本库
"ahooks": "^3.3.12"                     一些常用 Hooks
"antd": "^4.20.3"                       Ant Design 组件库
"axios": "^0.27.2"                      请求库
"dayjs": "^1.11.2"                      轻量级日期库
"lodash": "^4.17.21"                    一些通用脚本库
"nprogress": "^0.2.0"                   加载进度条
"react": "^18.1.0"                      React
"react-dom": "^18.1.0"                  React Dom
"react-router-dom": "^6.3.0"            React 路由库
"recoil": "^0.7.3-alpha.2"              React 轻量级状态库
"web-vitals": "^2.1.4"                  页面加载时长监控，暂时未使用
```

### 组件规范

组件默认未页面组件，即放置在在对应页面下 `components` 文件中，当有多个页面都在使用相同的组件时才考虑提升到全局组件。

项目主要使用 [Antd](https://ant.design/components) 和 [Antd Pro Componetns](https://procomponents.ant.design/components) 这两个 UI 库。目前管理后台的设计基本遵循 [Ant Design](https://gw.alipayobjects.com/os/bmw-prod/22208f9d-f8c5-4d7c-b87a-fec290e96527.sketch) 设计规范，因此很多页面级组件可以直接在以上两个 UI 库中找到对应的组件。

### 样式规范

项目采用 `less`，less 用法与 `sass` 极为接近。
在页面中优先采用 module style 的形式进行引入，即 `style.mudule.less`，这样可以保证样式只应用于当前页面（组件）。

### 全局引用

为了减少文件的引入层级，目前支持 `@` 引用变量，该变量指向根目录的 `src` 文件夹，比如引用用户 api，可以使用 `@/servers/user`。
