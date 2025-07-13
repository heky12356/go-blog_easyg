# go-blog_easyg 前端

这是基于 Next.js 和 TypeScript 构建的 go-blog_easyg 项目的前端部分。

## 目录
- [项目结构](#项目结构)
- [安装依赖](#安装依赖)
- [运行应用](#运行应用)
- [目录布局](#目录布局)

## 项目结构

前端项目主要包含以下模块：
- **App**：主应用程序组件，包括用户认证、管理面板、文章发布等。
- **Components**：可复用的 UI 组件，如页脚、导航栏、Markdown 渲染器等。
- **Utils**：工具函数，用于查询操作和字符串处理。
- **Public**：静态资源文件夹，例如图片或字体。
- **配置文件**：`next.config.ts`, `tsconfig.json` 等。

## 安装依赖

在开始之前，请确保已经安装了 Node.js 和 npm。然后执行以下命令来安装项目所需依赖：
```bash
npm install
```

## 运行应用

安装完成后，可以通过以下命令启动开发服务器：
```bash
npm run dev
```

打开浏览器并访问 [http://localhost:3000](http://localhost:3000) 查看应用。

## 目录布局

以下是项目的简化目录结构：
```
src/
├── app/
│   ├── (user)/
│   │   ├── login/
│   │   └── register/
│   ├── admin/
│   │   ├── components/
│   │   │   └── logout/
│   │   ├── create/
│   │   ├── delete/
│   │   ├── layout.tsx
│   │   └── page.tsx
│   ├── components/
│   │   ├── footer.tsx
│   │   ├── mark.tsx
│   │   ├── navbar.tsx
│   │   └── titlesuffix.tsx
│   ├── post/
│   │   ├── [uid]/
│   │   └── page.tsx
│   ├── test/
│   │   └── page.tsx
│   ├── utils/
│   │   ├── getsuffix.tsx
│   │   └── query.tsx
│   ├── globals.css
│   ├── layout.tsx
│   ├── loading.tsx
│   ├── page.module.css
│   └── page.tsx
├── middleware.ts
```

有关项目设置的更多详细信息，请参阅配置文件如 `next.config.ts` 和 `tsconfig.json`。