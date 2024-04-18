import TabContext from '@mui/lab/TabContext'
import TabList from '@mui/lab/TabList'
import TabPanel from '@mui/lab/TabPanel'
import Box from '@mui/material/Box'
import Tab from '@mui/material/Tab'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'

export default function Tabs() {
  const router = useRouter()
  // URLに基づいてタブの値を設定
  const [value, setValue] = useState('/')

  useEffect(() => {
    const path = router.pathname.split('/')[1]
    const basePath = `/${path}`
    setValue(basePath)
  }, [router.pathname])

  const handleChange = (event: React.SyntheticEvent, newValue: string) => {
    if (newValue !== value) {
      router.push(newValue)
    }
  }

  const handleClick = (event: React.MouseEvent<HTMLDivElement>) => {
    const target = event.target as HTMLElement
    const tabValue = target.closest('[role="tab"]')?.getAttribute('value') // タブの値を取得
    // 現在のパスと異なる場合に遷移
    if (tabValue && tabValue !== value) {
      router.push(tabValue)
    }
  }

  const handleTabClick = (path: string) => {
    // 現在のパスと異なる場合に遷移
    // if (value && value !== path) {
    router.push(path)
    // }
  }

  return (
    <Box sx={{ width: '100%', typography: 'body1' }}>
      <TabContext value={value}>
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
        <TabPanel value='/opponent_recruitings'>対戦相手募集</TabPanel>
        <TabPanel value='/schedules'>スケジュール管理</TabPanel>
      </TabContext>
    </Box>
  )
}
