import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import { VitePWA } from 'vite-plugin-pwa'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),
    VitePWA({
      registerType: 'autoUpdate',
      injectRegister: 'auto',
      // SW off in `vite dev`; test offline against a `vite build` + `vite preview`.
      devOptions: { enabled: false },
      manifest: {
        name: 'Guide Me',
        short_name: 'GuideMe',
        description:
          'Plan trips and manage travel documents for the people you care about — works offline.',
        theme_color: '#1769aa',
        background_color: '#ffffff',
        display: 'standalone',
        start_url: '/',
        scope: '/',
        icons: [
          { src: 'pwa-192x192.png', sizes: '192x192', type: 'image/png' },
          { src: 'pwa-512x512.png', sizes: '512x512', type: 'image/png' },
          {
            src: 'maskable-icon-512x512.png',
            sizes: '512x512',
            type: 'image/png',
            purpose: 'maskable',
          },
        ],
      },
      workbox: {
        // Precache app shell + icons, but NOT the heavy demo test-doc PNGs (small first load).
        globPatterns: ['**/*.{js,css,html,ico,svg,woff,woff2,png}'],
        globIgnores: ['**/test-docs/**'],
        // Router uses createWebHistory → deep links must fall back to index.html offline.
        navigateFallback: 'index.html',
        runtimeCaching: [
          {
            // Demo document images: fetched on demand, cached for offline (CacheFirst).
            urlPattern: ({ url }) => url.pathname.startsWith('/test-docs/'),
            handler: 'CacheFirst',
            options: {
              cacheName: 'test-docs',
              expiration: { maxEntries: 30, maxAgeSeconds: 60 * 60 * 24 * 30 },
              cacheableResponse: { statuses: [0, 200] },
            },
          },
        ],
      },
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
})
