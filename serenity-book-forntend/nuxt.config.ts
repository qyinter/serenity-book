// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  nitro: {
    compressPublicAssets: { gzip: true, brotli: true }
  },

  devtools: { enabled: process.env.NODE_ENV !== 'production' },

  css: ['~/assets/style/global.scss'],

  modules: [
    '@bg-dev/nuxt-naiveui',
    '@nuxtjs/eslint-module',
    'nuxt-simple-sitemap',
    'nuxt-simple-robots'
  ],

  features: {
    inlineStyles: true
  },

  runtimeConfig: {
    public: {
      baseURL: process.env.NUXT_API_URL,
      cdnURL: process.env.NUXT_CDN_URL
    }
  },

  app: {
    head: {
      charset: 'utf-8',
      viewport:
        'width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0, viewport-fit=cover',
      meta: [
        {
          name: 'format-detection',
          content: 'telephone=no'
        },
        {
          name: 'apple-mobile-web-app-capable',
          content: 'yes'
        }
      ],
      link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }]
    }
  },

  sitemap: {
    sitemaps: false,
    exclude: ['/404'],
    autoLastmod: true
  },

  robots: {
    blockNonSeoBots: true,
    disallowNonIndexableRoutes: true
  },

  routeRules: {
    '/sitemap.xml': {
      swr: 3600 * 24
    }
  }
})
