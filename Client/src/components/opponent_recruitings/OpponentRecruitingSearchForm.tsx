import { useEffect, useState } from 'react'
import {
  TextField,
  Button,
  FormControlLabel,
  Checkbox,
  RadioGroup,
  Radio,
  Grid,
  MenuItem,
  Box,
  Typography,
} from '@mui/material'
import { prefectures } from 'src/utils/prefectures'
import CustomDatePicker from '../commons/CustomDatePicker'
import { useRouter } from 'next/router'

interface SearchFormProps {
  onSearch: (params: string) => void
}

const daysOfWeek = [
  { jp: '日曜日', en: 'Sunday' },
  { jp: '月曜日', en: 'Monday' },
  { jp: '火曜日', en: 'Tuesday' },
  { jp: '水曜日', en: 'Wednesday' },
  { jp: '木曜日', en: 'Thursday' },
  { jp: '金曜日', en: 'Friday' },
  { jp: '土曜日', en: 'Saturday' },
]

const OpponentRecruitingSearchForm: React.FC<SearchFormProps> = ({ onSearch }) => {
  const router = useRouter()
  const [hasGround, setHasGround] = useState<string>('')
  const [prefectureId, setPrefectureId] = useState<string>('')
  const [isActive, setIsActive] = useState<boolean>(false)
  const [date, setDate] = useState<Date | null>(null)
  const [day, setDay] = useState<string>('')
  const [dateOrDay, setDateOrDay] = useState<string>('')

  useEffect(() => {
    // クエリパラメータから初期値を設定
    const query = router.query
    if (router.isReady) {
      setHasGround((query.has_ground as string) || '')
      setPrefectureId((query.prefecture_id as string) || '')
      setIsActive(query.is_active === 'true')
      setDateOrDay(query.date ? 'date' : query.day ? 'day' : '')
      if (query.date) setDate(new Date(query.date as string))
      const queryDay = query.day as string
      if (queryDay) {
        // 日本語の曜日名に変換する
        const dayObject = daysOfWeek.find((day) => day.en === queryDay)
        setDay(dayObject ? dayObject.jp : '')
      }
    }
  }, [router.isReady, router.query])

  const handleHasGroundChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const value = event.target.value
    setHasGround(value === hasGround ? '' : value)
  }

  const handleReset = () => {
    setHasGround('')
    setPrefectureId('')
    setIsActive(false)
    setDate(null)
    setDay('')
    setDateOrDay('')
  }

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()

    const params = new URLSearchParams()

    if (hasGround) params.append('has_ground', hasGround)
    if (prefectureId) params.append('prefecture_id', prefectureId)
    if (isActive) params.append('is_active', 'true')
    if (dateOrDay === 'date' && date) params.append('date', date.toISOString().split('T')[0])
    if (dateOrDay === 'day' && day) {
      const dayInEnglish = daysOfWeek.find((d) => d.jp === day)?.en
      if (dayInEnglish) {
        params.append('day', dayInEnglish)
      }
    }

    const queryParams = params.toString()

    onSearch(queryParams)
  }

  return (
    <Grid container justifyContent='center'>
      <Grid item xs={12} sm={10} md={8} lg={6}>
        <Box
          sx={{
            backgroundColor: '#f5f5f5',
            padding: 2,
            borderRadius: 2,
            maxWidth: 500,
            marginTop: 2,
            marginLeft: 'auto',
            marginRight: 'auto',
            boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
          }}
        >
          <form onSubmit={handleSubmit}>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <TextField
                  select
                  label='エリア'
                  value={prefectureId}
                  onChange={(e) => setPrefectureId(e.target.value)}
                  fullWidth
                  sx={{ '& .MuiInputBase-root': { color: '#333' } }}
                >
                  <MenuItem value=''>
                    <em>選択してください</em>
                  </MenuItem>
                  {prefectures.map((prefecture) => (
                    <MenuItem key={prefecture.id} value={prefecture.id}>
                      {prefecture.name}
                    </MenuItem>
                  ))}
                </TextField>
              </Grid>
              <Grid item xs={12}>
                <Typography variant='subtitle1' sx={{ fontWeight: 'bold', color: '#333' }}>
                  日程
                </Typography>
                <RadioGroup row value={dateOrDay} onChange={(e) => setDateOrDay(e.target.value)}>
                  <FormControlLabel value='' control={<Radio />} label='選択しない' />
                  <FormControlLabel value='date' control={<Radio />} label='日付' />
                  <FormControlLabel value='day' control={<Radio />} label='曜日' />
                </RadioGroup>
              </Grid>
              <Grid item xs={12}>
                {dateOrDay === 'date' ? (
                  <CustomDatePicker value={date} onChange={setDate} />
                ) : dateOrDay === 'day' ? (
                  <TextField
                    select
                    label='曜日'
                    value={day}
                    onChange={(e) => setDay(e.target.value)}
                    fullWidth
                    sx={{ '& .MuiInputBase-root': { color: '#333' } }}
                  >
                    {daysOfWeek.map((day, index) => (
                      <MenuItem key={index} value={day.jp}>
                        {day.jp}
                      </MenuItem>
                    ))}
                  </TextField>
                ) : null}
              </Grid>
              <Grid item xs={12}>
                <Typography variant='subtitle1' sx={{ fontWeight: 'bold', color: '#333' }}>
                  グラウンドの有無
                </Typography>
                <RadioGroup row value={hasGround} onChange={handleHasGroundChange}>
                  <FormControlLabel value='' control={<Radio />} label='選択しない' />
                  <FormControlLabel value='true' control={<Radio />} label='有' />
                  <FormControlLabel value='false' control={<Radio />} label='無' />
                </RadioGroup>
              </Grid>
              <Grid item xs={12}>
                <FormControlLabel
                  control={
                    <Checkbox checked={isActive} onChange={(e) => setIsActive(e.target.checked)} />
                  }
                  label='募集中のみ'
                />
              </Grid>
              <Grid item xs={12}>
                <Button
                  type='submit'
                  variant='contained'
                  sx={{
                    backgroundColor: '#4CAF50',
                    '&:hover': { backgroundColor: '#388E3C' },
                    color: '#fff',
                    fontWeight: 'bold',
                    fontSize: '16px',
                  }}
                  fullWidth
                >
                  検索
                </Button>
                <Button
                  onClick={handleReset}
                  sx={{
                    marginTop: 2,
                    color: 'gray',
                    fontWeight: 'bold',
                    fontSize: '16px',
                    border: '1px solid #ccc',
                    backgroundColor: 'transparent', 
                    '&:hover': {
                      backgroundColor: '#f0f0f0',
                      borderColor: '#aaa',
                    },
                  }}
                  fullWidth
                >
                  リセット
                </Button>
              </Grid>
            </Grid>
          </form>
        </Box>
      </Grid>
    </Grid>
  )
}

export default OpponentRecruitingSearchForm
