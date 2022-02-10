declare namespace NodeJS {
  interface ProcessEnv {
    NEXT_PUBLIC_AUTH_API_URL: string
    NEXT_PUBLIC_SLIDE_API_URL: string
    NEXT_PUBLIC_SUMMARIZE_API_URL: string
    NEXT_PUBLIC_GITHUB_AUTH_CALLBACK_URL: string
    NEXT_PUBLIC_GITHUB_CLIENT_ID: string
  }
}