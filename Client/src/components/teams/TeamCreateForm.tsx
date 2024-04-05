import React, { useContext, useEffect, useState } from 'react'
import {
  Container,
  Typography,
  Grid,
  TextField,
  Button,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  SelectChangeEvent,
  FormHelperText,
} from '@mui/material'
import { teamLevels } from 'src/utils/teamLevel'
import { CreateTeamsApiRequest } from 'src/types/apiRequests'
import useApiWithFlashMessage from 'src/hooks/useApiWithFlashMessage'
import { prefectures } from 'src/utils/prefectures'
import { useNavigateOpponentRecruitingsIndex } from 'src/hooks/useNavigateOpponentRecruitingsIndex'
import PrimaryButton from '../commons/PrimaryButton'
import { AuthContext } from 'src/contexts/AuthContext'
import { User } from 'src/types/user'
import { CreateTeamsApiResponse } from 'src/types/apiResponses'
import { saveDataWithExpiry } from 'src/utils/localStrageHelper'

type Errors = {
  [key in keyof CreateTeamsApiRequest]?: string
}

function TeamCreateForm() {
  const [formData, setFormData] = useState<CreateTeamsApiRequest>({
    name: '',
    prefecture_id: '',
    level_id: '',
    home_page_url: '',
    other: '',
  })
  const [errors, setErrors] = useState<Errors>({})
  const { request, data } = useApiWithFlashMessage<CreateTeamsApiResponse>()
  const { user, setUser } = useContext(AuthContext)
  const navigateOpponentRecruitingsIndex = useNavigateOpponentRecruitingsIndex()

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target
    setFormData({
      ...formData,
      [name]: value,
    })
  }

  const handleSelectChange = (e: SelectChangeEvent<number>) => {
    const { name, value } = e.target
    setFormData({
      ...formData,
      [name]: value === '' ? '' : Number(value),
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
          credentials: 'include',
        }
        await request('/teams', options)
      } catch (error) {
        console.error('チーム作成に失敗しました。', error)
      }
    } else {
      setErrors(validationErrors)
    }
  }

  useEffect(() => {
    if (data) {
      setFormData({
        name: '',
        prefecture_id: '',
        level_id: '',
        home_page_url: '',
        other: '',
      })
      const userData: User = {
        name: user?.name,
        current_team_id: data.current_team_id,
        current_team_name: data.current_team_name,
        current_team_role: data.current_team_role,
      }
      // ユーザー情報をローカルストレージにキャッシュ
      saveDataWithExpiry<User>('user', userData, 3600)
      setUser(userData)
      // 対戦相手募集リストに移動
      navigateOpponentRecruitingsIndex()
    }
  }, [data])

  const validateForm = () => {
    const errors: Errors = {}
    if (!formData.name) errors.name = 'チーム名は必須です。'
    if (formData.name.length > 32) errors.name = 'チーム名は32文字以内でなければなりません。'
    if (!formData.prefecture_id) errors.prefecture_id = '活動拠点は必須です。'
    if (!formData.level_id) errors.level_id = 'チームレベルは必須です。'
    if (formData.other.length > 500) errors.other = 'その他は500文字以内でなければなりません。'
    return errors
  }

  return (
    <Container maxWidth='sm'>
      <Typography variant='h4' component='h2' gutterBottom marginTop={2}>
        チーム作成
      </Typography>
      <form onSubmit={handleSubmit}>
        <Grid container spacing={3}>
          <Grid item xs={12}>
            <TextField
              label='チーム名'
              name='name'
              value={formData.name}
              onChange={handleInputChange}
              fullWidth
              error={Boolean(errors.name) || formData.name.length > 32}
              helperText={
                <div style={{ display: 'flex', justifyContent: 'space-between' }}>
                  <span
                    style={{
                      color: errors.name || formData.name.length > 32 ? 'error' : 'inherit',
                    }}
                  >
                    {errors.name ||
                      (formData.name.length > 32
                        ? 'チーム名は32文字以内でなければなりません。'
                        : '')}
                  </span>
                  <span>{formData.name.length}/32</span>
                </div>
              }
            />
          </Grid>
          <Grid item xs={12}>
            <FormControl fullWidth error={Boolean(errors.prefecture_id)}>
              <InputLabel id='prefecture-label'>活動拠点</InputLabel>
              <Select
                labelId='prefecture-label'
                name='prefecture_id'
                value={formData.prefecture_id}
                onChange={handleSelectChange}
                label='活動拠点'
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
              <FormHelperText>{errors.prefecture_id}</FormHelperText>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <FormControl fullWidth error={Boolean(errors.level_id)}>
              <InputLabel id='level-label'>チームレベル</InputLabel>
              <Select
                labelId='level-label'
                name='level_id'
                value={formData.level_id}
                onChange={handleSelectChange}
                label='チームレベル'
              >
                <MenuItem value=''>
                  <em>選択してください</em>
                </MenuItem>
                {teamLevels.map((teamLevel) => (
                  <MenuItem key={teamLevel.id} value={teamLevel.id}>
                    {teamLevel.name}
                  </MenuItem>
                ))}
              </Select>
              <FormHelperText>{errors.level_id}</FormHelperText>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <TextField
              label='ホームページリンク'
              name='home_page_url'
              value={formData.home_page_url}
              onChange={handleInputChange}
              fullWidth
            />
          </Grid>
          <Grid item xs={12}>
            <TextField
              label='その他'
              name='other'
              value={formData.other}
              onChange={handleInputChange}
              fullWidth
              multiline
              rows={4}
              error={Boolean(errors.other) || formData.other.length > 500}
              helperText={
                <div style={{ display: 'flex', justifyContent: 'space-between' }}>
                  <span
                    style={{
                      color: errors.other || formData.other.length > 500 ? 'error' : 'inherit',
                    }}
                  >
                    {errors.other ||
                      (formData.other.length > 500
                        ? 'その他は500文字以内でなければなりません。'
                        : '')}
                  </span>
                  <span>{formData.other.length}/500</span>
                </div>
              }
            />
          </Grid>
          <Grid item xs={12} style={{ marginBottom: '20px' }}>
            <PrimaryButton>作成</PrimaryButton>
          </Grid>
        </Grid>
      </form>
    </Container>
  )
}

export default TeamCreateForm
