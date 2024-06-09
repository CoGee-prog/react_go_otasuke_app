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
  SelectChangeEvent,
} from '@mui/material'
import React, { useContext, useEffect, useState } from 'react'
import useApiWithFlashMessage from 'src/hooks/useApiWithFlashMessage'
import { useNavigateOpponentRecruitingsCreate } from 'src/hooks/useNavigateOpponentRecruitingsCreate'
import { useNavigateOpponentRecruitingsIndex } from 'src/hooks/useNavigateOpponentRecruitingsIndex'
import { CreateOpponentRecruitingsApiRequest } from 'src/types/apiRequests'
import { prefectures } from 'src/utils/prefectures'
import PrimaryButton from '../commons/PrimaryButton'
import { useRouter } from 'next/router'
import { TeamRole } from 'src/types/teamRole'
import { AuthContext } from 'src/contexts/AuthContext'
import CustomDatePicker from '../commons/CustomDatePicker'
import { useNavigateOpponentRecruitingDetail } from 'src/hooks/useNavigateOpponentRecruitingDetail'
import { OpponentRecruitingWithComments } from 'src/types/opponentRecruiting'
import { GetOpponentRecruitingApiResponse } from 'src/types/apiResponses'

type Errors = {
  [key in keyof OpponentRecruitingsFormData]?: string
}

export interface OpponentRecruitingsFormData {
  title: string
  has_ground: boolean
  ground_name: string
  prefecture_id: number
  date: string
  start_time: string
  end_time: string
  detail: string
}

interface OpponentRecruitingFormProps {
  isEditing: boolean
  initialData?: OpponentRecruitingsFormData
  id?: string
  onUpdateSuccess?: (updatedData: GetOpponentRecruitingApiResponse) => void
}

function OpponentRecruitingForm({
  isEditing = false,
  initialData,
  id,
  onUpdateSuccess,
}: OpponentRecruitingFormProps) {
  const router = useRouter()
  const { user } = useContext(AuthContext)
  const navigateOpponentRecruitingsIndex = useNavigateOpponentRecruitingsIndex()
  const [isAccessAllowed, setIsAccessAllowed] = useState(false)

  const [formData, setFormData] = useState<OpponentRecruitingsFormData>(
    initialData || {
      title: '',
      has_ground: false,
      ground_name: '',
      prefecture_id: 0,
      date: '',
      start_time: '',
      end_time: '',
      detail: '',
    },
  )
  const [errors, setErrors] = useState<Errors>({})
  const { request, data } = useApiWithFlashMessage<GetOpponentRecruitingApiResponse>()
  const navigateOpponentRecruitingsCreate = useNavigateOpponentRecruitingsCreate()
  const navigateOpponentRecruitingDetail = useNavigateOpponentRecruitingDetail(id!)

  useEffect(() => {
    // ユーザーの役割が管理者または副管理者であれば 、アクセス可能とする
    if (
      user &&
      (user.current_team_role == TeamRole.ADMIN || user.current_team_role == TeamRole.SUB_ADMIN)
    ) {
      setIsAccessAllowed(true)
    } else {
      // 適切な権限がない場合、対戦相手募集リストにリダイレクト
      navigateOpponentRecruitingsIndex()
    }
  }, [router])

  useEffect(() => {
    if (isEditing && data && onUpdateSuccess) {
      // 編集完了を親コンポーネントに通知
      onUpdateSuccess(data!)
      // 対戦相手募集詳細に移動
      navigateOpponentRecruitingDetail()
    }
  }, [data])

  if (!isAccessAllowed) {
    // 認証されていない場合は、何も表示しない
    return null
  }

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, type, checked } = e.target
    setFormData({
      ...formData,
      [name]: type === 'checkbox' ? checked : value,
    })
  }

  const handleSelectChange = (e: SelectChangeEvent<string>) => {
    const name = e.target.name as keyof OpponentRecruitingsFormData
    const value = parseInt(e.target.value)
    setFormData({
      ...formData,
      [name]: value,
    })
  }

  const handleDateChange = (newValue: Date | null) => {
    if (newValue && !isNaN(newValue.getTime())) {
      setFormData({
        ...formData,
        date: newValue.toISOString().split('T')[0],
      })
    } else {
      setFormData({
        ...formData,
        date: '',
      })
    }
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
        // 開始日時と終了日時をフォーマット
        // 日付と時間を結合してフォーマット
        const formattedStartTime = `${formData.date}T${formData.start_time}:00+09:00`
        const formattedEndTime = `${formData.date}T${formData.end_time}:00+09:00`
        // dateプロパティを除く
        const { date, ...formDataExcludeDate } = formData
        const requestData = {
          ...formDataExcludeDate,
          start_time: formattedStartTime,
          end_time: formattedEndTime,
        }
        const options: RequestInit = {
          method: isEditing ? 'PATCH' : 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(requestData),
          credentials: 'include',
        }
        await request(isEditing ? `/opponent_recruitings/${id}` : '/opponent_recruitings', options)

        if (!isEditing) {
          // 対戦相手募集リストに移動
          navigateOpponentRecruitingsIndex()
        } else {
        }

        // フォームをクリアする
        setFormData({
          title: '',
          has_ground: false,
          ground_name: '',
          prefecture_id: 0,
          date: '',
          start_time: '',
          end_time: '',
          detail: '',
        })
      } catch (error) {
        if (isEditing) {
          console.error('対戦相手募集の変更に失敗しました。', error)
        } else {
          console.error('対戦相手募集の作成に失敗しました。', error)
        }
      }
    } else {
      setErrors(validationErrors)
      if (isEditing) {
        // 対戦相手募集詳細に移動
        navigateOpponentRecruitingDetail()
      } else {
        // 対戦相手募集作成に移動
        navigateOpponentRecruitingsCreate()
      }
    }
  }

  const validateForm = () => {
    const errors: Errors = {}
    if (!formData.title) errors.title = 'タイトルは必須です。'
    if (formData.title.length > 50) errors.title = 'タイトルは50文字以内でなければなりません。'
    if (!formData.prefecture_id) errors.prefecture_id = '都道府県の選択は必須です。'
    if (!formData.date) errors.date = '日付は必須です。'
    if (!formData.start_time) errors.start_time = '開始時間は必須です。'
    if (!formData.end_time) errors.end_time = '終了時間は必須です。'
    if (formData.start_time && formData.end_time && formData.end_time < formData.start_time) {
      errors.end_time = '終了時間は開始時間より後でなければなりません。'
    }
    if (formData.has_ground && !formData.ground_name)
      errors.ground_name = 'グラウンド名の入力は必須です。'
    if (!formData.ground_name && formData.ground_name.length > 50)
      errors.ground_name = 'グラウンド名は50文字以内でなければなりません。'
    if (formData.detail.length > 1000) errors.detail = '詳細は1000文字以内でなければなりません。'
    return errors
  }

  return (
    <Container maxWidth='sm'>
      <Typography variant='h4' component='h2' gutterBottom marginTop={2}>
        {isEditing ? '対戦相手募集編集' : '対戦相手募集作成'}
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
              error={Boolean(errors.title) || formData.title.length > 50}
              helperText={
                <div style={{ display: 'flex', justifyContent: 'space-between' }}>
                  <span
                    style={{
                      color: errors.title || formData.title.length > 50 ? 'error' : 'inherit',
                    }}
                  >
                    {errors.title ||
                      (formData.title.length > 50
                        ? 'タイトルは50文字以内でなければなりません。'
                        : '')}
                  </span>
                  <span>{formData.title.length}/50</span>
                </div>
              }
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
                error={Boolean(errors.ground_name) || formData.ground_name.length > 50}
                helperText={
                  <div style={{ display: 'flex', justifyContent: 'space-between' }}>
                    <span
                      style={{
                        color:
                          errors.ground_name || formData.ground_name.length > 50
                            ? 'error'
                            : 'inherit',
                      }}
                    >
                      {errors.ground_name ||
                        (formData.ground_name.length > 50
                          ? 'グラウンド名は50文字以内でなければなりません。'
                          : '')}
                    </span>
                    <span>{formData.ground_name.length}/50</span>
                  </div>
                }
              />
            </Grid>
          )}
          <Grid item xs={12}>
            <FormControl fullWidth error={Boolean(errors.prefecture_id)}>
              <InputLabel id='prefecture-label'>都道府県</InputLabel>
              <Select
                labelId='prefecture-label'
                name='prefecture_id'
                value={formData.prefecture_id.toString()}
                onChange={handleSelectChange}
                label='都道府県'
              >
                <MenuItem value='0'>
                  <em>選択してください</em>
                </MenuItem>
                {prefectures.map((prefecture) => (
                  <MenuItem key={prefecture.id} value={prefecture.id}>
                    {prefecture.name}
                  </MenuItem>
                ))}
              </Select>
              {errors.prefecture_id && (
                <Typography
                  variant='caption'
                  color='error'
                  sx={{ display: 'block', marginTop: '4px', marginLeft: '14px' }}
                >
                  {errors.prefecture_id}
                </Typography>
              )}
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <CustomDatePicker
              value={formData.date ? new Date(formData.date) : null}
              onChange={handleDateChange}
              error={Boolean(errors.date)}
              helperText={errors.date}
            />
          </Grid>
          <Grid
            item
            xs={12}
            container
            alignItems='flex-start'
            justifyContent='space-between'
            spacing={2}
          >
            <Grid item xs={5}>
              <TextField
                label='開始時間'
                type='time'
                name='start_time'
                value={formData.start_time}
                onChange={handleInputChange}
                fullWidth
                InputLabelProps={{
                  shrink: true,
                }}
                inputProps={{
                  // 5分刻み
                  step: 300,
                }}
                error={Boolean(errors.start_time)}
                helperText={errors.start_time || ' '} // 空白を追加して高さを保持
              />
            </Grid>
            <Grid item xs={1} container justifyContent='center'>
              <Typography variant='h6' sx={{ my: 2 }}>
                ~
              </Typography>{' '}
            </Grid>
            <Grid item xs={5}>
              <TextField
                label='終了時間'
                type='time'
                name='end_time'
                value={formData.end_time}
                onChange={handleInputChange}
                fullWidth
                InputLabelProps={{
                  shrink: true,
                }}
                inputProps={{
                  // 5分刻み
                  step: 300,
                }}
                error={Boolean(errors.end_time)}
                helperText={errors.end_time || ' '} // 空白を追加して高さを保持
              />
            </Grid>
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
              error={Boolean(errors.detail) || formData.detail.length > 1000}
              helperText={
                <div style={{ display: 'flex', justifyContent: 'space-between' }}>
                  <span
                    style={{
                      color: errors.detail || formData.detail.length > 1000 ? 'error' : 'inherit',
                    }}
                  >
                    {errors.detail ||
                      (formData.detail.length > 1000
                        ? '詳細は1000文字以内でなければなりません。'
                        : '')}
                  </span>
                  <span>{formData.detail.length}/1000</span>
                </div>
              }
            />
          </Grid>
          <Grid item xs={12} style={{ marginBottom: '20px' }}>
            <PrimaryButton>{isEditing ? '更新' : '作成'}</PrimaryButton>
          </Grid>
        </Grid>
      </form>
    </Container>
  )
}

export default OpponentRecruitingForm
