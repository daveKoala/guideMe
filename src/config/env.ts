// Single, type-safe config surface for per-environment values.
// Reads import.meta.env (typed in env.d.ts), applies defaults, coerces types.
// Import anywhere via: import { env } from '@/config/env'
export const env = Object.freeze({
  apiBaseUrl: import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:3000',
  useMockData: (import.meta.env.VITE_USE_MOCK_DATA ?? 'true') === 'true',
})

export type Env = typeof env
