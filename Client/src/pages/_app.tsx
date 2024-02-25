import MainLayout from 'src/components/MainLayout'
import '../styles/globals.css'
import type { AppProps } from 'next/app'
import Head from 'next/head'
import { AuthProvider } from 'src/providers/AuthProvider'
import { FlashMessageProvider } from 'src/providers/FlashMessageProvider'
import { useRouter } from 'next/router'

export default function App({ Component, pageProps }: AppProps) {
  const router = useRouter()

  // 404ページのパスを判断するロジック（例として、カスタム404ページが無い場合を想定）
  const isNotFoundPage = router.pathname === '/404'
  return (
    <>
      <Head>
        <script src='https://www.gstatic.com/firebasejs/ui/6.0.2/firebase-ui-auth__ja.js'></script>
        <link
          type='text/css'
          rel='stylesheet'
          href='https://www.gstatic.com/firebasejs/ui/6.0.2/firebase-ui-auth.css'
        />
      </Head>
      {isNotFoundPage ? (
        <Component {...pageProps} />
      ) : (
        <FlashMessageProvider>
          <AuthProvider>
            <MainLayout>
              <Component {...pageProps} />
            </MainLayout>
          </AuthProvider>
        </FlashMessageProvider>
      )}
    </>
  )
}
