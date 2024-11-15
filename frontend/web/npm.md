# npm Q&A


## Node Version Manager
- https://github.com/nvm-sh/nvm

nvm 常用命令
- nvm install <version>     # 安装特定版本的 Node.js
- nvm use <version>         # 切换到指定的 Node.js 版本
- nvm ls                    # 列出所有已安装的 Node.js 版本
- nvm uninstall <version>   # 卸载指定版本的 Node.js


## package.json
package.json 是 Node.js 项目中的一个配置文件，用于管理项目的元数据、依赖项、脚本命令和其他信息。它是由 npm（Node Package Manager）生成的，通常位于项目的根目录。


## npm 常用命令
1.npm install
- `npm install` 安装所有依赖项
- `npm install <package>`: 安装一个新的依赖项，并添加到 dependencies。
- `npm install <package> --save-dev`: 安装一个开发依赖项，并添加到 devDependencies。
- `npm run <script>`: 运行 scripts 中定义的脚本，例如 npm run dev。

2.npm run
- npm run start：通常启动应用程序（开发或生产环境）。
- npm run dev：运行开发服务器，通常用于开发模式。
- npm run serve：常用于启动开发服务器，Vue CLI 中较常见。
- npm run build：构建生产环境的优化代码。
- npm run test：运行项目测试，适合用在 CI/CD 中。
- npm run lint：运行代码检查，通常使用 linter 工具。

3.npm run build
- 构建生产环境的静态文件，它会优化、压缩代码并生成静态资源，以便部署到服务器上。
- 执行 npm run build 后，Vite 会在项目目录下生成一个 dist 文件夹，其中包含已优化和压缩的生产环境文件：
  - JavaScript 和 CSS 文件会被分割、压缩和哈希化，以便于缓存。
  - HTML 文件 和 静态资源 也会经过优化。
  - 动态导入 会根据路由或模块分离，生成相应的文件块，以减小首次加载体积。
- 可以通过以下命令在本地预览构建后的文件，npm run preview 启动一个本地服务器，模拟生产环境，通常会在 http://localhost:4173 打开项目页面，帮助检查构建后的效果。
- 可以将 dist 文件夹上传到 Web 服务器（如 Nginx、Apache）或静态网站托管平台（如 Vercel、Netlify、GitHub Pages）中，提供给用户访问。


## npm Q&A
### 一、npm run dev 和 npm run serve 的区别
```
npm run dev 和 npm run serve 都是用于启动开发服务器的命令，但具体使用哪个取决于所用的工具或框架，下面是它们的区别：

npm run dev:
- 常用于现代构建工具，例如 Vite 和 Nuxt 3 等。
- 通常表示启动开发模式，开发服务器会自动监视文件更改并重新加载页面。
- 一般速度更快，适用于现代的轻量级构建工具。
- 示例：Vite 项目会默认使用 npm run dev 启动开发服务器。

npm run serve:
- 在 Vue CLI 3 和 Vue CLI 4 项目中常用。
- 也是启动开发服务器并监视文件变更，适合基于 Webpack 的项目。
- 示例：Vue CLI 使用 Webpack 构建项目，因此使用 npm run serve 来启动开发模式。

选择哪个命令？
- 如果您使用 Vite 或 Nuxt 3，通常使用 npm run dev。
- 如果您使用 Vue CLI（Vue 2 或 3 项目用 Webpack），通常使用 npm run serve。

总结：虽然两者的效果相似，但具体用法取决于项目的构建工具。
```
