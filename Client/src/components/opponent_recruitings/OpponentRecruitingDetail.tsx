import React, { useContext, useEffect, useState } from 'react'
import {
  Box,
  Card,
  CardContent,
  Typography,
  Divider,
  TextField,
  Grid,
} from '@mui/material'
import { OpponentRecruitingWithComments } from 'src/types/opponentRecruiting'
import PrimaryButton from '../commons/PrimaryButton'
import { formatTimeRange } from 'src/utils/formatDateTime'
import { AuthContext } from 'src/contexts/AuthContext'
import useApiWithFlashMessage from 'src/hooks/useApiWithFlashMessage'
import { GetOpponentRecruitingApiResponse } from 'src/types/apiResponses'
import OpponentRecruitingComment from './OpponentRecruitingComment'
import { TeamRole } from 'src/types/teamRole'
import Link from 'next/link'
import { useRouter } from 'next/router'

interface OpponentRecruitingDetailProps {
  initialOpponentRecruitingWithComments: OpponentRecruitingWithComments
  id: string
}

const OpponentRecruitingDetail: React.FC<OpponentRecruitingDetailProps> = ({
  initialOpponentRecruitingWithComments,
  id,
}) => {
  const [newComment, setNewComment] = useState('')
  const [error, setError] = useState(false)
  const { request, data } = useApiWithFlashMessage<GetOpponentRecruitingApiResponse>()
  const [opponentRecruitingWithComments, setOpponentRecruitingWithComments] = useState(
    initialOpponentRecruitingWithComments,
  )
  const [editingCommentId, setEditingCommentId] = useState<number | null>(null)
  const { user } = useContext(AuthContext)
	const router = useRouter()
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
      <Typography variant='h4' component='h2' gutterBottom marginTop={2}>
        対戦相手募集詳細
      </Typography>
      <Card sx={{ mb: 2 }}>
        <CardContent>
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
            { label: 'チーム', value: opponentRecruitingWithComments.team.name },
            { label: 'レベル', value: opponentRecruitingWithComments.team.level },
            { label: '詳細', value: opponentRecruitingWithComments.detail },
          ].map((item, index, arr) => (
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
      </Card>
      <Typography variant='h6' component='h2' gutterBottom marginTop={4}>
        コメント
      </Typography>
      {opponentRecruitingWithComments.comments.map((comment, index) => (
        <Box key={index} sx={{ my: 2 }}>
          <OpponentRecruitingComment
            comment={comment}
            opponentRecruitingTeamId={opponentRecruitingWithComments.team.id}
            isEditing={editingCommentId === comment.id}
            onEdit={() => setEditingCommentId(comment.id)}
            onUpdate={handleUpdateComment}
            onDelete={handleDeleteComment}
          />
          {index < opponentRecruitingWithComments.comments.length - 1 && <Divider />}
        </Box>
      ))}
      {user ? (
        user.current_team_role === TeamRole.ADMIN ||
        user.current_team_role === TeamRole.SUB_ADMIN ? (
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
