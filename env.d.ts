/// <reference types="vite/client" />
/// <reference types="vite-plugin-pwa/client" />

interface ImportMetaEnv {
  readonly VITE_API_BASE_URL: string
  readonly VITE_USE_MOCK_DATA: string // env vars are always strings; coerced in src/config/env.ts
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
