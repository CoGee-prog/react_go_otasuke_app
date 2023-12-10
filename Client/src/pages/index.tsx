import Head from 'next/head'
import Header from '../../components/Header'
import Tabs from '../../components/Tabs'

export default function Home() {
  return (
    <div>
      <Head>
        <script src='https://www.gstatic.com/firebasejs/ui/6.0.2/firebase-ui-auth__ja.js'></script>
        <link
          type='text/css'
          rel='stylesheet'
          href='https://www.gstatic.com/firebasejs/ui/6.0.2/firebase-ui-auth.css'
        />
      </Head>
      <Header />
      <Tabs />
    </div>
  )
}
