export const __prod__ = process.env.NODE_ENV === 'production'

export const AUTH_API_URL = process.env.NEXT_PUBLIC_AUTH_API_URL
export const SLIDE_API_URL = process.env.NEXT_PUBLIC_SLIDE_API_URL
export const SUMMARIZE_API_URL = process.env.NEXT_PUBLIC_SUMMARIZE_API_URL
export const COMPUTE_API_URL = process.env.NEXT_PUBLIC_COMPUTE_API_URL
export const LINK_API_URL = process.env.NEXT_PUBLIC_LINK_API_URL
export const GITHUB_AUTH_CALLBACK_URL =
  process.env.NEXT_PUBLIC_GITHUB_AUTH_CALLBACK_URL
export const GITHUB_CLIENT_ID = process.env.NEXT_PUBLIC_GITHUB_CLIENT_ID
