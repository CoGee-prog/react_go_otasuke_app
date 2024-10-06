// pages/404.js
import Header from 'src/components/layouts/Header'
import styles from '../styles/custom404.module.css'

export default function Custom404() {
  return (
    <>
      <Header />
      <div className={styles.container}>
        <h1>404 - ページが見つかりません</h1>
      </div>
    </>
  )
}
