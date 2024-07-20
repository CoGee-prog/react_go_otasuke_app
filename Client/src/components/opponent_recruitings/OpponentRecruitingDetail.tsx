import React, { useContext, useEffect, useState } from 'react'
import {
  Box,
  Card,
  CardContent,
  Typography,
  Divider,
  TextField,
  Grid,
  Chip,
  IconButton,
  ListItemIcon,
  ListItemText,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  Menu,
  MenuItem,
  Button,
} from '@mui/material'
import { OpponentRecruitingWithComments } from 'src/types/opponentRecruiting'
import PrimaryButton from '../commons/PrimaryButton'
import { formatTimeRange } from 'src/utils/formatDateTime'
import { AuthContext } from 'src/contexts/AuthContext'
import useApiWithFlashMessage from 'src/hooks/useApiWithFlashMessage'
import { GetOpponentRecruitingApiResponse } from 'src/types/apiResponses'
import OpponentRecruitingCommentForm from './OpponentRecruitingCommentForm'
import { TeamRole } from 'src/types/teamRole'
import Link from 'next/link'
import { useRouter } from 'next/router'
import DangerButton from '../commons/DangerButton'
import EditIcon from '@mui/icons-material/Edit'
import DeleteIcon from '@mui/icons-material/Delete'
import MoreVertIcon from '@mui/icons-material/MoreVert'
import OpponentRecruitingForm from './OpponentRecruitingForm'
import { OpponentRecruitingsFormData } from './OpponentRecruitingForm'
import { getPrefectureIdFromName } from 'src/utils/prefectures'
import { ArrowBack } from '@mui/icons-material'
import { useNavigateOpponentRecruitingsIndex } from 'src/hooks/useNavigateOpponentRecruitingsIndex'

interface OpponentRecruitingDetailProps {
  initialOpponentRecruitingWithComments: OpponentRecruitingWithComments
  id: string
}

// // OpponentRecruitingWithComments オブジェクトを OpponentRecruitingsFormData に変換
function mapToFormData(recruiting: OpponentRecruitingWithComments): OpponentRecruitingsFormData {
  const date = recruiting.start_time.split('T')[0]
  // 時間はフォームに必要な分単位(HH:mm)の形式で取り出す
  const startTime = recruiting.start_time.split('T')[1].split('+')[0].slice(0, 5)
  const endTime = recruiting.end_time.split('T')[1].split('+')[0].slice(0, 5)
  const prefectureId = getPrefectureIdFromName(recruiting.prefecture)
  return {
    title: recruiting.title,
    has_ground: recruiting.has_ground,
    ground_name: recruiting.ground_name || '',
    prefecture_id: prefectureId,
    date: date,
    start_time: startTime,
    end_time: endTime,
    detail: recruiting.detail,
  }
}

const OpponentRecruitingDetail: React.FC<OpponentRecruitingDetailProps> = ({
  initialOpponentRecruitingWithComments,
  id,
}) => {
  const router = useRouter()
  const [newComment, setNewComment] = useState('')
  const [error, setError] = useState(false)
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null)
  const { request, data } = useApiWithFlashMessage<GetOpponentRecruitingApiResponse>()
  const [opponentRecruitingWithComments, setOpponentRecruitingWithComments] = useState(
    initialOpponentRecruitingWithComments,
  )
  const [isEditing, setIsEditing] = useState(false)
  const [editingCommentId, setEditingCommentId] = useState<number | null>(null)
  const { user } = useContext(AuthContext)
  const [openDeleteDialog, setOpenDeleteDialog] = useState(false)
  const navigateOpponentRecruitingsIndex = useNavigateOpponentRecruitingsIndex()

  const handleOpenDeleteDialog = () => {
    setOpenDeleteDialog(true)
  }

  const handleCloseDeleteDialog = () => {
    setOpenDeleteDialog(false)
  }

  const handleDeleteConfirmed = async () => {
    handleCloseDeleteDialog()
    try {
      const options: RequestInit = {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
      }
      await request(`/opponent_recruitings/${id}`, options)
      // 対戦相手募集リストに移動
      navigateOpponentRecruitingsIndex()
    } catch (error) {
      console.error('対戦相手募集の削除に失敗しました:', error)
    }
  }

  const handleUpdateOpponentRecruitingStatus = async (isActive: boolean) => {
    try {
      const options: RequestInit = {
        method: 'PATCH',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ is_active: isActive }),
        credentials: 'include',
      }
      await request(`/opponent_recruitings/${id}/status`, options)
    } catch (error) {
      console.error('状態の更新に失敗しました', error)
    }
  }

  const handleOpenMenu = (event: React.MouseEvent<HTMLButtonElement>) => {
    setAnchorEl(event.currentTarget)
  }

  const handleCloseMenu = () => {
    setAnchorEl(null)
  }

  const toggleEdit = () => {
    setIsEditing(!isEditing)
    handleCloseMenu()
  }

  const handleUpdateSuccess = (updatedData: GetOpponentRecruitingApiResponse) => {
    setIsEditing(false)
    setOpponentRecruitingWithComments(updatedData.opponent_recruiting)
  }

  const handleUpdateComment = async (commentId: number, updatedComment: string) => {
    setEditingCommentId(null)
    try {
      const options: RequestInit = {
        method: 'PATCH',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ content: updatedComment }),
        credentials: 'include',
      }
      await request(`/opponent_recruitings/${id}/comments/${commentId}`, options)
    } catch (error) {
      console.error('コメント更新に失敗しました', error)
    }
  }

  const handleDeleteComment = async (commentId: number) => {
    setEditingCommentId(null)
    try {
      const options: RequestInit = {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
      }
      await request(`/opponent_recruitings/${id}/comments/${commentId}`, options)
    } catch (error) {
      console.error('コメント削除に失敗しました', error)
    }
  }

  useEffect(() => {
    if (data) {
      // APIからのレスポンスで状態を更新
      setOpponentRecruitingWithComments(data.opponent_recruiting)
    }
  }, [data])

  const handlePostComment = async () => {
    if (newComment.length > 1000) {
      setError(true)
      return
    }
    try {
      const options: RequestInit = {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ content: newComment }),
        credentials: 'include',
      }
      await request(`/opponent_recruitings/${id}/comments`, options)
      setNewComment('')
    } catch (error) {
      console.error('コメントに失敗しました', error)
    }
  }

  const handleCommentChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setNewComment(e.target.value)
    if (e.target.value.length > 1000) {
      setError(true)
    } else {
      setError(false)
    }
  }

  return (
    <Box sx={{ maxWidth: 600, mx: 'auto', my: 4, px: 2 }}>
      <Grid item>
        <IconButton
          onClick={() => router.back()}
          aria-label='戻る'
          sx={{
            color: 'black',
            border: '1px solid black',
            borderRadius: 0,
            marginRight: 1,
          }}
        >
          <ArrowBack />
          <Typography variant='button' sx={{ color: 'black', ml: 1 }}>
            戻る
          </Typography>
        </IconButton>
      </Grid>
      <Grid item xs>
        <Typography variant='h4' component='h2' gutterBottom marginTop={2}>
          対戦相手募集詳細
        </Typography>
      </Grid>
      <Grid
        container
        spacing={2}
        direction='column'
        alignItems='center'
        justifyContent='center'
        style={{ marginTop: '3px', marginBottom: '10px' }}
      >
        <Grid item xs={12} style={{ display: 'flex', justifyContent: 'center' }}>
          <Box sx={{ maxWidth: 500, width: '100%', textAlign: 'center' }}>
            {user && opponentRecruitingWithComments && 
            user.current_team_id === opponentRecruitingWithComments.team.id &&
            (user.current_team_role === TeamRole.ADMIN ||
              user.current_team_role === TeamRole.SUB_ADMIN) ? (
              opponentRecruitingWithComments.is_active ? (
                <DangerButton onClick={() => handleUpdateOpponentRecruitingStatus(false)}>
                  募集を終了する
                </DangerButton>
              ) : (
                <PrimaryButton onClick={() => handleUpdateOpponentRecruitingStatus(true)}>
                  募集を再開する
                </PrimaryButton>
              )
            ) : null}
          </Box>
        </Grid>
      </Grid>

      <Card
        sx={{
          mb: 2,
          backgroundColor: opponentRecruitingWithComments.is_active ? 'white' : '#d0d0d0',
          position: 'relative',
        }}
      >
        {/* 編集中であれば編集フォームを出す */}
        {isEditing ? (
          <OpponentRecruitingForm
            isEditing={true}
            initialData={mapToFormData(initialOpponentRecruitingWithComments)}
            id={id}
            onUpdateSuccess={handleUpdateSuccess}
          />
        ) : (
          <CardContent>
            <Box sx={{ display: 'flex', alignItems: 'center', mb: 0.5 }}>
              <Chip
                label={opponentRecruitingWithComments.is_active ? '募集中' : '募集終了'}
                size='small'
                color={opponentRecruitingWithComments.is_active ? 'primary' : 'default'}
              />
            </Box>
            <Typography variant='h5' gutterBottom>
              {opponentRecruitingWithComments.title}
            </Typography>
            <Typography sx={{ fontWeight: 'bold', mb: 1.5 }}>
              {formatTimeRange(
                opponentRecruitingWithComments.start_time,
                opponentRecruitingWithComments.end_time,
              )
                .text.split(' ')
                .map((part, index, array) => (
                  <span
                    key={index}
                    style={{
                      color:
                        index === 1 || index === array.length - 2
                          ? index === 1
                            ? formatTimeRange(
                                opponentRecruitingWithComments.start_time,
                                opponentRecruitingWithComments.end_time,
                              ).dayOfWeekColor
                            : formatTimeRange(
                                opponentRecruitingWithComments.start_time,
                                opponentRecruitingWithComments.end_time,
                              ).endDayOfWeekColor
                          : 'inherit',
                    }}
                  >
                    {part}{' '}
                  </span>
                ))}
            </Typography>
            {[
              { label: '都道府県', value: opponentRecruitingWithComments.prefecture },
              {
                label: 'グラウンド名',
                value: opponentRecruitingWithComments.ground_name,
                condition: opponentRecruitingWithComments.has_ground,
              },
              { label: 'チーム', value: opponentRecruitingWithComments.team.name },
              { label: 'レベル', value: opponentRecruitingWithComments.team.level },
              { label: '詳細', value: opponentRecruitingWithComments.detail },
            ]
              .filter((item) => item.condition !== false)
              .map((item, index, arr) => (
                <Box key={index} sx={{ my: 1 }}>
                  <Typography variant='body1' sx={{ fontWeight: 'bold' }} gutterBottom>
                    {item.label}
                  </Typography>
                  <Typography variant='body1' gutterBottom>
                    {item.value}
                  </Typography>
                  {index < arr.length - 1 && <Divider />}
                </Box>
              ))}
          </CardContent>
        )}
        {user &&
        user.current_team_id === opponentRecruitingWithComments.team.id &&
        (user.current_team_role === TeamRole.ADMIN ||
          user.current_team_role === TeamRole.SUB_ADMIN) ? (
          <>
            <IconButton
              aria-label='more'
              aria-controls='comment-menu'
              aria-haspopup='true'
              onClick={handleOpenMenu}
              sx={{ position: 'absolute', top: 8, right: 8 }}
            >
              <MoreVertIcon />
            </IconButton>
            <Menu
              id='comment-menu'
              anchorEl={anchorEl}
              keepMounted
              open={Boolean(anchorEl)}
              onClose={handleCloseMenu}
            >
              {isEditing ? (
                <MenuItem onClick={toggleEdit}>
                  <ListItemText>キャンセル</ListItemText>
                </MenuItem>
              ) : (
                opponentRecruitingWithComments.is_active && (
                  <MenuItem onClick={toggleEdit}>
                    <ListItemIcon>
                      <EditIcon fontSize='small' />
                    </ListItemIcon>
                    <ListItemText>編集</ListItemText>
                  </MenuItem>
                )
              )}
              {!isEditing && (
                <MenuItem onClick={handleOpenDeleteDialog}>
                  <ListItemIcon>
                    <DeleteIcon fontSize='small' />
                  </ListItemIcon>
                  <ListItemText>削除</ListItemText>
                </MenuItem>
              )}
            </Menu>
            <Dialog
              open={openDeleteDialog}
              onClose={handleCloseDeleteDialog}
              aria-labelledby='alert-dialog-title'
              aria-describedby='alert-dialog-description'
            >
              <DialogTitle id='alert-dialog-title'>{'削除確認'}</DialogTitle>
              <DialogContent>
                <DialogContentText id='alert-dialog-description'>
                  本当に削除してもよろしいですか？
                </DialogContentText>
              </DialogContent>
              <DialogActions>
                <Button onClick={handleCloseDeleteDialog} color='primary'>
                  キャンセル
                </Button>
                <Button onClick={handleDeleteConfirmed} color='primary' autoFocus>
                  削除
                </Button>
              </DialogActions>
            </Dialog>
          </>
        ) : null}
      </Card>
      <Typography variant='h6' component='h2' gutterBottom marginTop={4}>
        コメント
      </Typography>
      {opponentRecruitingWithComments.comments.map((comment, index) => (
        <Box key={index} sx={{ my: 2 }}>
          <OpponentRecruitingCommentForm
            comment={comment}
            opponentRecruitingTeamId={opponentRecruitingWithComments.team.id}
            isEditing={editingCommentId === comment.id}
            onEdit={() => setEditingCommentId(comment.id)}
            onUpdate={handleUpdateComment}
            onDelete={handleDeleteComment}
            isActiveOpponentRecruiting={opponentRecruitingWithComments.is_active}
          />
          {index < opponentRecruitingWithComments.comments.length - 1 && <Divider />}
        </Box>
      ))}
      {user ? (
        user.current_team_role === TeamRole.ADMIN ||
        user.current_team_role === TeamRole.SUB_ADMIN ? (
          opponentRecruitingWithComments.is_active ? (
            <Card variant='outlined'>
              <CardContent>
                <Box
                  sx={{
                    my: 2,
                    display: 'flex',
                    flexDirection: 'row',
                    justifyContent: 'space-between',
                    alignItems: 'center',
                  }}
                >
                  <TextField
                    label='チーム'
                    variant='filled'
                    value={user?.current_team_name}
                    InputProps={{ readOnly: true }}
                    sx={{ width: '48%' }}
                    InputLabelProps={{ shrink: true }}
                  />
                  <TextField
                    label='投稿者'
                    variant='filled'
                    value={user?.name}
                    InputProps={{ readOnly: true }}
                    sx={{ width: '48%' }}
                    InputLabelProps={{ shrink: true }}
                  />
                </Box>
                <TextField
                  label='コメント'
                  variant='outlined'
                  fullWidth
                  multiline
                  rows={4}
                  value={newComment}
                  onChange={handleCommentChange}
                  error={error}
                  helperText={
                    <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
                      {error ? 'コメントは1000文字以内で入力してください。' : <span>&nbsp;</span>}
                      <span>{`${newComment.length}/1000`}</span>
                    </Box>
                  }
                  sx={{ my: 1 }}
                />
                <PrimaryButton variant='contained' color='primary' onClick={handlePostComment}>
                  コメントを投稿
                </PrimaryButton>
              </CardContent>
            </Card>
          ) : (
            <Grid item xs={12} style={{ display: 'flex', justifyContent: 'center' }}>
              <Box sx={{ maxWidth: 500, width: '100%', textAlign: 'center' }}>
                <p>募集が終了しているのでコメントできません</p>
              </Box>
            </Grid>
          )
        ) : (
          <Grid item xs={12} style={{ display: 'flex', justifyContent: 'center' }}>
            <Box sx={{ maxWidth: 500, width: '100%', textAlign: 'center' }}>
              <p>チームの管理者か副管理者のみコメントを作成できます</p>
            </Box>
          </Grid>
        )
      ) : (
        <Link href={`/login?from=${encodeURIComponent(router.asPath)}`} passHref>
          <PrimaryButton color='inherit'>ログインしてコメントする</PrimaryButton>
        </Link>
      )}
    </Box>
  )
}

export default OpponentRecruitingDetail
