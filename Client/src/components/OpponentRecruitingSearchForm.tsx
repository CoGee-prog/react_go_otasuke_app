import { useState } from 'react'
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
import { LocalizationProvider, DatePicker } from '@mui/x-date-pickers-pro'
import CustomDatePicker from './CustomDatePicker'

interface SearchFormProps {
  onSearch: (params: string) => void
}

const daysOfWeek = ['日曜日', '月曜日', '火曜日', '水曜日', '木曜日', '金曜日', '土曜日']

const OpponentRecruitingSearchForm: React.FC<SearchFormProps> = ({ onSearch }) => {
  const [hasGround, setHasGround] = useState<string>('')
  const [prefectureId, setPrefectureId] = useState<string>('')
  const [isActive, setIsActive] = useState<boolean>(false)
  const [date, setDate] = useState<Date | null>(null)
  const [day, setDay] = useState<string>('')
  const [dateOrDay, setDateOrDay] = useState<string>('date')

  const handleHasGroundChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const value = event.target.value
    setHasGround(value === hasGround ? '' : value)
  }

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()

    const queryParams = new URLSearchParams({
      has_ground: hasGround,
      prefecture_id: prefectureId,
      is_active: isActive ? 'true' : '',
      date: dateOrDay === 'date' && date ? date.toISOString().split('T')[0] : '',
      day: dateOrDay === 'day' ? day : '',
    }).toString()

    onSearch(queryParams)
  }

  return (
    <Grid container justifyContent='center'>
      <Grid item xs={12} sm={10} md={8} lg={6}>
        <Box
          sx={{
            backgroundColor: '#f0f0f0',
            padding: 2,
            borderRadius: 2,
            maxWidth: 500,
            marginTop: 2,
            marginLeft: 'auto',
            marginRight: 'auto',
          }}
        >
          <form onSubmit={handleSubmit}>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <Typography variant='subtitle1'>グラウンドの有無</Typography>
                <RadioGroup row value={hasGround} onChange={handleHasGroundChange}>
                  <FormControlLabel value='true' control={<Radio />} label='有り' />
                  <FormControlLabel value='false' control={<Radio />} label='無し' />
                  <FormControlLabel value='' control={<Radio />} label='選択なし' />{' '}
                </RadioGroup>
              </Grid>
              <Grid item xs={12}>
                <TextField
                  select
                  label='都道府県'
                  value={prefectureId}
                  onChange={(e) => setPrefectureId(e.target.value)}
                  fullWidth
                >
                  {prefectures.map((prefecture) => (
                    <MenuItem key={prefecture.id} value={prefecture.id}>
                      {prefecture.name}
                    </MenuItem>
                  ))}
                </TextField>
              </Grid>
              <Grid item xs={12}>
                <FormControlLabel
                  control={
                    <Checkbox checked={isActive} onChange={(e) => setIsActive(e.target.checked)} />
                  }
                  label='募集中'
                />
              </Grid>
              <Grid item xs={12}>
                <RadioGroup row value={dateOrDay} onChange={(e) => setDateOrDay(e.target.value)}>
                  <FormControlLabel value='date' control={<Radio />} label='日付' />
                  <FormControlLabel value='day' control={<Radio />} label='曜日' />
                  <FormControlLabel value='none' control={<Radio />} label='選択しない' />
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
                  >
                    {daysOfWeek.map((day, index) => (
                      <MenuItem key={index} value={day}>
                        {day}
                      </MenuItem>
                    ))}
                  </TextField>
                ) : null}
              </Grid>
              <Grid item xs={12}>
                <Button type='submit' variant='contained' color='primary' fullWidth>
                  検索
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
