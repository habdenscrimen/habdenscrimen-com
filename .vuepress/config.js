module.exports = {
  title: 'habdenscrimen',
  description: 'Блог Вячеслава Ефремова',
  head: [
    ['link', { rel: 'icon', href: '/favicon.png' }],
    ['link', { rel: 'shortcut icon', href: '/favicon.ico' }],
    ['meta', { name: 'theme-color', content: '#ffffff' }],
    ['meta', { name: 'viewport', content: 'width=device-width, initial-scale=1.0' }],
  ],

  // theme: require.resolve('../../vuepress-theme-blog'),
  theme: '@vuepress/theme-blog',
  themeConfig: {
    dateFormat: 'MMM D YYYY',
    pwa: true,
    nav: [],
    smoothScroll: true,
    footer: {
      contact: [],
      copyright: [],
    },
    telegramLink: 'https://t.me/habdenscrimen_com',
  },
  plugins: [
    [
      '@vuepress/blog',
      {
        directories: [
          {
            // Unique ID of current classification
            id: 'posts',
            // Target directory
            dirname: '_posts',
            // Path of the `entry page` (or `list page`)
            path: '/',
            pagination: {
              lengthPerPage: 7,
            },
          },
        ],
        sitemap: {
          hostname: 'https://habdenscrimen.com',
        },
        // newsletter: {
        //   endpoint:
        //     'https://billyyyyy3320.us4.list-manage.com/subscribe/post?u=4905113ee00d8210c2004e038&amp;id=bd18d40138',
        // },
      },
    ],
  ],
}
