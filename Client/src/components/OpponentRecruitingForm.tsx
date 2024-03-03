import {
  Container,
  Typography,
  Grid,
  TextField,
  FormControlLabel,
  Checkbox,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Button,
  SelectChangeEvent,
} from '@mui/material'
import React, { useState } from 'react'
import useApiWithFlashMessage from 'src/hooks/useApiWithFlashMessage'
import { useNavigateOpponentRecruitingsCreate } from 'src/hooks/useNavigateOpponentRecrutingsCreate'
import { useNavigateOpponentRecruitingsIndex } from 'src/hooks/useNavigateOpponentRecrutingsIndex'
import { CreateOpponentRecruitingsApiRequest } from 'src/types/apiRequests'
import { prefectures } from 'src/utils/prefectures'
import PrimaryButton from './PrimaryButton'
import DangerButton from './DangerButton'

type Errors = {
  [key in keyof CreateOpponentRecruitingsApiRequest]?: string
}

function OpponentRecruitingForm() {
  const [formData, setFormData] = useState<CreateOpponentRecruitingsApiRequest>({
    title: '',
    has_ground: false,
    ground_name: '',
    prefecture_id: '',
    start_time: '',
    end_time: '',
    detail: '',
  })
  const [errors, setErrors] = useState<Errors>({})
  const { request } = useApiWithFlashMessage<CreateOpponentRecruitingsApiRequest>()

  const navigateOpponentRecruitingsIndex = useNavigateOpponentRecruitingsIndex()
  const navigateOpponentRecruitingsCreate = useNavigateOpponentRecruitingsCreate()

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, type, checked } = e.target
    setFormData({
      ...formData,
      [name]: type === 'checkbox' ? checked : value,
    })
  }

  const handleSelectChange = (e: SelectChangeEvent<string>) => {
    const name = e.target.name as keyof FormData
    const value = e.target.value
    setFormData({
      ...formData,
      [name]: value,
    })
  }

  const handleTextareaChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    const { name, value } = e.target
    setFormData({
      ...formData,
      [name]: value,
    })
  }

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const validationErrors = validateForm()
    if (Object.keys(validationErrors).length === 0) {
      try {
        const options: RequestInit = {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(formData),
        }
        await request('/opponent_recruitings', options)

        // フォームをクリアする
        setFormData({
          title: '',
          has_ground: false,
          ground_name: '',
          prefecture_id: '',
          start_time: '',
          end_time: '',
          detail: '',
        })
        // 対戦相手募集リストに移動
        navigateOpponentRecruitingsIndex()
      } catch (error) {
        console.error('対戦相手募集の作成に失敗しました。', error)
      }
    } else {
      setErrors(validationErrors)
      // 対戦相手募集作成に移動
      navigateOpponentRecruitingsCreate()
    }
  }

  const validateForm = () => {
    const errors: Errors = {}
    if (!formData.title) errors.title = 'タイトルは必須です。'
    if (!formData.prefecture_id) errors.prefecture_id = '都道府県の選択は必須です。'
    if (!formData.start_time) errors.start_time = '開始日時の選択は必須です。'
    if (!formData.end_time) errors.end_time = '終了日時の選択は必須です。'
    if (formData.has_ground && !formData.ground_name)
      errors.ground_name = 'グラウンド名の入力は必須です。'
    return errors
  }

  return (
    <Container maxWidth='sm'>
      <Typography variant='h4' component='h2' gutterBottom marginTop={2}>
        対戦相手募集作成
      </Typography>
      <form onSubmit={handleSubmit}>
        <Grid container spacing={3}>
          <Grid item xs={12}>
            <TextField
              label='タイトル'
              name='title'
              value={formData.title}
              onChange={handleInputChange}
              fullWidth
              error={Boolean(errors.title)}
              helperText={errors.title}
            />
          </Grid>
          <Grid item xs={12}>
            <FormControlLabel
              control={
                <Checkbox
                  name='has_ground'
                  checked={formData.has_ground}
                  onChange={handleInputChange}
                />
              }
              label='グラウンド有'
            />
          </Grid>
          {formData.has_ground && (
            <Grid item xs={12}>
              <TextField
                label='グラウンド名'
                name='ground_name'
                value={formData.ground_name}
                onChange={handleInputChange}
                fullWidth
                error={Boolean(errors.ground_name)}
                helperText={errors.ground_name}
              />
            </Grid>
          )}
          <Grid item xs={12}>
            <FormControl fullWidth error={Boolean(errors.prefecture_id)}>
              <InputLabel id='prefecture-label'>都道府県</InputLabel>
              <Select
                labelId='prefecture-label'
                name='prefecture_id'
                value={formData.prefecture_id}
                onChange={handleSelectChange}
                label='都道府県'
              >
                <MenuItem value=''>
                  <em>選択してください</em>
                </MenuItem>
                {prefectures.map((prefecture) => (
                  <MenuItem key={prefecture.id} value={prefecture.id}>
                    {prefecture.name}
                  </MenuItem>
                ))}
              </Select>
              {errors.prefecture_id && (
                <Typography variant='caption' color='error'>
                  {errors.prefecture_id}
                </Typography>
              )}
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <TextField
              label='開始日時'
              type='datetime-local'
              name='start_time'
              value={formData.start_time}
              onChange={handleInputChange}
              fullWidth
              InputLabelProps={{
                shrink: true,
              }}
              error={Boolean(errors.start_time)}
              helperText={errors.start_time}
            />
          </Grid>
          <Grid item xs={12}>
            <TextField
              label='終了日時'
              type='time'
              name='end_time'
              value={formData.end_time}
              onChange={handleInputChange}
              fullWidth
              InputLabelProps={{
                shrink: true,
              }}
              error={Boolean(errors.end_time)}
              helperText={errors.end_time}
            />
          </Grid>
          <Grid item xs={12}>
            <TextField
              label='詳細'
              name='detail'
              value={formData.detail}
              onChange={handleTextareaChange}
              fullWidth
              multiline
              rows={4}
            />
          </Grid>
          <Grid item xs={12}>
            <PrimaryButton>作成</PrimaryButton>
          </Grid>
        </Grid>
      </form>
    </Container>
  )
}

export default OpponentRecruitingForm
