const moment = require('moment');
module.exports = {
    title: '编程人生',
    description: 'coder at work',
    head: [
        ['link', { rel: 'icon', href: '/favicon.ico' }],
        ['link', { rel: 'stylesheet', href: 'https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.5.1/katex.min.css' }],
        ['link', { rel: 'stylesheet', href: 'https://cdn.jsdelivr.net/github-markdown-css/2.2.1/github-markdown.css' }],
        ['script', { type: "text/javascript" }, `var _hmt = _hmt || [];
        (function() {
          var hm = document.createElement("script");
          hm.src = "https://hm.baidu.com/hm.js?1ac83de30e327b8683cabc7b35aaf728";
          var s = document.getElementsByTagName("script")[0];
          s.parentNode.insertBefore(hm, s);
        })();
      `],
        ['meta', { name: 'keywords', content: 'mac,dev,php,go,golang,fe,java,kotlin,python,tech,coder,work' }]
    ],
    markdown: {
        lineNumbers: true, // 代码块显示行号
        extendMarkdown: md => {
            md.use(require('markdown-it-katex'))
        }
    },
    themeConfig: {
        nav: [
            { text: '通识', link: '/know/' },
            { text: '技艺', link: '/a/' },
            { text: 'blog', link: '/blog/' },
        ],
        // 假定是 GitHub. 同时也可以是一个完整的 GitLab URL
        repo: 'daydaygo/coder',
        // 自定义仓库链接文字。默认从 `themeConfig.repo` 中自动推断为
        // "GitHub"/"GitLab"/"Bitbucket" 其中之一，或是 "Source"。
        repoLabel: 'GitHub',
        // 假如文档不是放在仓库的根目录下：
        docsDir: 'docs',
        // 假如文档放在一个特定的分支下：
        docsBranch: 'master',
        editLinks: true,
        editLinkText: '在 github.com 上编辑此页',
        sidebar: {
            sidebar: "auto",
            "/know/": [
                {
                    title: 'know',
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        '/know/',
                    ]
                },
            ],
        },
        // https://www.vuepress.cn/zh/theme/default-theme-config.html#algolia-%E6%90%9C%E7%B4%A2
        // algolia: {
        //     apiKey: '<API_KEY>',
        //     indexName: '<INDEX_NAME>'
        // },
        sidebarDepth: 2,
        lastUpdated: '上次更新',
        serviceWorker: {
            updatePopup: {
                message: "发现新内容可用",
                buttonText: "刷新"
            }
        },
    },
    plugins: [
        ['@vuepress/last-updated',
            {
                transformer: (timestamp, lang) => {
                    // 不要忘了安装 moment
                    const moment = require('moment')
                    moment.locale("zh-cn")
                    return moment(timestamp).format('YYYY-MM-DD HH:mm:ss')
                },

                dateOptions: {
                    hours12: true
                }
            }],
        '@vuepress/back-to-top',
        '@vuepress/active-header-links',
        '@vuepress/medium-zoom',
        '@vuepress/nprogress'
    ]
}
