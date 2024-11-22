# Web Development
- https://en.wikipedia.org/wiki/HTML
- https://en.wikipedia.org/wiki/CSS
- https://en.wikipedia.org/wiki/JavaScript

Tool
- - JSFiddle - Code Playground https://jsfiddle.net


## Best Practice
- https://frontendmasters.com/guides/front-end-handbook/2018/what-is-a-FD.html
- https://developer.mozilla.org/en-US/docs/Web
- https://developer.mozilla.org/en-US/docs/Learn
- https://microsoft.github.io/Web-Dev-For-Beginners/
- https://github.com/microsoft/Web-Dev-For-Beginners
- https://github.com/Microsoft/frontend-bootcamp
- https://github.com/Chalarangelo/30-seconds-of-code
- https://www.udemy.com/course/foundations-of-front-end-development/
- 前端九部 - 入门者手册2019 https://www.yuque.com/fe9/basic
- Patterns.dev - Improve how you architect webapps https://www.patterns.dev

### 原理
Web browser
- https://en.wikipedia.org/wiki/Web_browser
- How Web Works https://github.com/vasanthk/how-web-works
- 图解浏览器的基本工作原理 https://zhuanlan.zhihu.com/p/47407398
- Chrome 浏览器运行原理你了解多少？https://mp.weixin.qq.com/s/wjrcO2Ej7BEThWVsCnXEtA
- Inside look at modern web browser (part 1) https://developer.chrome.com/blog/inside-browser-part1/
- OWASP Cheat Sheet Series https://cheatsheetseries.owasp.org/

Life of a URLRequest
- Life of a URLRequest https://chromium.googlesource.com/chromium/src/+/refs/heads/main/net/docs/life-of-a-url-request.md
- What happens behind the scenes when we type www.google.com in a browser?  https://github.com/vasanthk/how-web-works
- What Happens When You Type a Url in Your Browser and Press Enter https://blog.nicolasmesa.co/posts/2018/08/what-happens-when-you-type-a-url-in-your-browser-and-press-enter/
- What happens when you type a URL into your browser? https://aws.amazon.com/blogs/mobile/what-happens-when-you-type-a-url-into-your-browser/
- What happens when you type a URL into your browser? https://blog.bytebytego.com/p/what-happens-when-you-type-a-url


## JavaScript Package Manager
### npm
- https://www.npmjs.com
- https://www.npmjs.com/package/package
- https://docs.npmjs.com
- https://github.com/npm/npm
- https://github.com/npm/cli
- https://docs.npmjs.com/cli/v10
- npm介绍、package.json详解、npm install原理、常用指令、npx工具、发布自己的包 https://www.cnblogs.com/yaopengfei/p/16144782.html

Docs
- https://docs.npmjs.com/cli
- https://docs.npmjs.com/cli/v10/commands

npmmirror 镜像站
- https://npmmirror.com


### npx
- https://docs.npmjs.com/cli/v10/commands/npx
- https://www.freecodecamp.org/news/npm-vs-npx-whats-the-difference/
- npx与npm的差异解析，以及包管理器yarn与Node版本管理工具nvm的使用方法详解 https://blog.csdn.net/shanghai597/article/details/133930521

1.package.json 中的 scripts 配置逻辑

```
"scripts": {
  "dev": "vue-cli-service serve",
  "build:prod": "vue-cli-service build",
  "lint": eslint --ext .js,.vue src
}


上面配置的 npm script 可分别对应 npx 执行命令
npm run dev
npm run build:prod
npm run lint

npx vue-cli-service serve
npx vue-cli-service build
npx eslint --ext .js,.vue src
```

2.环境变量设置

vue-cli-service 会自动设置 NODE_ENV 环境变量，以确保在开发和生产模式下使用正确的环境配置。Vue CLI 会在不同的构建模式中根据命令自动定义 NODE_ENV 的值：
- serve 命令：当你运行 vue-cli-service serve（开发服务器）时，Vue CLI 会将 NODE_ENV 自动设置为 development。
- build 命令：当你运行 vue-cli-service build（生产构建）时，Vue CLI 会将 NODE_ENV 自动设置为 production。
- --mode 选项：如果使用 --mode 选项指定模式，例如 --mode staging，Vue CLI 会根据 .env.staging 文件加载环境变量，并且 NODE_ENV 通常仍会被设置为 production，以确保使用生产模式的构建优化。



### yarn
- https://github.com/yarnpkg/yarn
- https://classic.yarnpkg.com

### pnpm
- https://github.com/pnpm/pnpm


## Toolkit
### webpack
- https://github.com/webpack/webpack
- https://webpack.js.org
- https://webpack.js.org/concepts/
- https://webpack.js.org/guides/
- https://webpack.js.org/guides/getting-started/
- https://webpack.js.org/api/
- https://webpack.js.org/configuration/

Rspack
- https://github.com/web-infra-dev/rspack


## Library
### jQuery
- https://api.jquery.com
- https://github.com/jquery
- You Don't Need jQuery https://github.com/camsong/You-Dont-Need-jQuery


## Third Party
Service
- https://www.bootcdn.cn

Font & Icon
- https://fontawesome.com.cn
- https://www.iconfont.cn
- https://github.com/FortAwesome/Font-Awesome
