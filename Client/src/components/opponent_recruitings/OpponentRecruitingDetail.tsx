import React, { useContext, useState } from 'react'
import {
  Box,
  Card,
  CardContent,
  Typography,
  Chip,
  Divider,
  TextField,
} from '@mui/material'
import { format } from 'date-fns'
import { OpponentRecruitingWithComments } from 'src/types/opponentRecruiting'
import PrimaryButton from '../commons/PrimaryButton'
import { formatTimeRange } from 'src/utils/formatDateTime'
import { AuthContext } from 'src/contexts/AuthContext'

interface OpponentRecruitingDetailProps {
  opponentRecruitingWithComments: OpponentRecruitingWithComments
}

const OpponentRecruitingDetail: React.FC<OpponentRecruitingDetailProps> = ({
  opponentRecruitingWithComments,
}) => {
  const [newComment, setNewComment] = useState('')
  const [error, setError] = useState(false)
  const handlePostComment = async () => {
    if (newComment.length > 1000) {
      setError(true)
      return
    }
    // APIリクエストを送る処理
    // ...
    // コメントを投稿後、ページを再読み込み
    window.location.reload()
  }

	const { user } = useContext(AuthContext)

  const handleCommentChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setNewComment(e.target.value)
    if (e.target.value.length > 1000) {
      setError(true)
    } else {
      setError(false)
    }
  }

  return (
    <Box sx={{ maxWidth: 600, mx: 'auto', my: 4 }}>
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
            { label: 'チーム', value: opponentRecruitingWithComments.team.name },
            { label: 'レベル', value: opponentRecruitingWithComments.team.level },
            { label: '都道府県', value: opponentRecruitingWithComments.prefecture },
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
          <Card variant='outlined'>
            <CardContent>
              <Typography variant='body1' gutterBottom>
                チーム: {comment.team_name}
              </Typography>
              <Typography variant='body1' gutterBottom>
                投稿者: {comment.user_name}
              </Typography>
              <Typography variant='body2'>
                {comment.content}
                {comment.edited && (
                  <Chip label='編集済み' size='small' sx={{ ml: 1, bgcolor: 'grey.300' }} />
                )}
                {comment.deleted && (
                  <Chip label='削除済み' size='small' sx={{ ml: 1, bgcolor: 'grey.300' }} />
                )}
              </Typography>
            </CardContent>
          </Card>
          {index < opponentRecruitingWithComments.comments.length - 1 && <Divider />}
        </Box>
      ))}
      <Box sx={{ my: 2 }}>
        <Typography variant='body1' gutterBottom>
          チーム: {user?.current_team_name}
        </Typography>
        <Typography variant='body1' gutterBottom>
          投稿者: {user?.name}
        </Typography>
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
      </Box>
    </Box>
  )
}

export default OpponentRecruitingDetail
