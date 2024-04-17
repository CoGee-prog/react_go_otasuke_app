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
    setValue(newValue)
    // 新しいURLにナビゲート
    router.push(newValue)
  }

  return (
    <Box sx={{ width: '100%', typography: 'body1' }}>
      <TabContext value={value}>
        <Box sx={{ borderBottom: 1, borderColor: 'divider', justifyContent: 'center' }}>
          <TabList onChange={handleChange} aria-label='lab API tabs example' centered>
            <Tab label='対戦相手募集' value='/opponent_recruitings' />
            <Tab label='スケジュール管理' value='/schedules' />
          </TabList>
        </Box>
        <TabPanel value='/schedules'>スケジュール管理</TabPanel>
      </TabContext>
    </Box>
  )
}
