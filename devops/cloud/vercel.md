# Vercel
- https://vercel.com

Docs
- https://vercel.com/docs
- https://vercel.com/docs/getting-started-with-vercel
- https://vercel.com/new


## Vercel Project Settings
### Sub Domain
Vercel 配置
- projects - Domains: 添加域名
- project - Settings - Domains: 编辑自定义生产二级域名

Cloudflare 配置
- DNS - Records: 添加解析记录
  - Type: CNAME
  - Name: custom subdomain prefix
  - Content: cname.vercel-dns.com
  - Proxy status: 只能是 DNS only
