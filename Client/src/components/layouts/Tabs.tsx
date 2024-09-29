import TabContext from '@mui/lab/TabContext'
import TabList from '@mui/lab/TabList'
import Box from '@mui/material/Box'
import Tab from '@mui/material/Tab'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'

export default function Tabs() {
  const router = useRouter()
  // URLに基づいてタブの値を設定
  const [basePath, setBasePath] = useState<string | null>(null)

  useEffect(() => {
    const pathSegments = router.pathname.split('/')
    const firstSegment = pathSegments[1]
    const secondSegment = pathSegments[2]

    // /opponent_recruitingsか/opponent_recruitings/[id]のパスのみハイライトする
    if (firstSegment === 'opponent_recruitings' && (!secondSegment || secondSegment === '[id]')) {
      setBasePath('/opponent_recruitings')
    } else if (firstSegment === '/schedules') {
      setBasePath('/schedules')
    } else {
      // その他のパスではハイライトしない
      setBasePath(null)
    }
  }, [router.pathname])

  const handleTabClick = (nextPath: string) => {
    // 現在のパスと異なる場合に遷移
    if (router.pathname !== nextPath) {
      router.push(nextPath)
    }
  }

  return (
    <Box sx={{ width: '100%', typography: 'body1' }}>
      <TabContext value={basePath || ''}>
        <Box sx={{ borderBottom: 1, borderColor: 'divider', justifyContent: 'center' }}>
          <TabList onChange={() => {}} aria-label='Tabs' centered>
            <Tab
              label='対戦相手募集'
              value='/opponent_recruitings'
              onClick={() => handleTabClick('/opponent_recruitings')}
            />
            <Tab
              label='スケジュール管理'
              value='/schedules'
              onClick={() => handleTabClick('/schedules')}
            />
          </TabList>
        </Box>
      </TabContext>
    </Box>
  )
}
