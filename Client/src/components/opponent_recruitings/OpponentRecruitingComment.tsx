import React, { useContext, useState } from 'react'
import {
  Card,
  CardContent,
  Typography,
  Chip,
  IconButton,
  Menu,
  MenuItem,
  TextField,
  Box,
} from '@mui/material'
import MoreVertIcon from '@mui/icons-material/MoreVert'
import { OpponentRecruitingComment } from 'src/types/opponentRecruiting'
import { AuthContext } from 'src/contexts/AuthContext'
import PrimaryButton from '../commons/PrimaryButton'

interface OpponentRecruitingCommentProps {
  comment: OpponentRecruitingComment
  opponentRecruitingTeamId: number
  isEditing: boolean
  onEdit: () => void
  onUpdate: (commentId: number, updatedContent: string) => void
  onDelete: (commentId: number) => void
}

const OpponentRecruitingComment: React.FC<OpponentRecruitingCommentProps> = ({
  comment,
  opponentRecruitingTeamId,
  isEditing,
  onEdit,
  onUpdate,
  onDelete,
}) => {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null)
  const [editedContent, setEditedContent] = useState(comment.content)
  const [error, setError] = useState(false)
  const { user } = useContext(AuthContext)

  const handleOpenMenu = (event: React.MouseEvent<HTMLButtonElement>) => {
    setAnchorEl(event.currentTarget)
  }

  const handleCloseMenu = () => {
    setAnchorEl(null)
  }

  const handleEdit = () => {
    onEdit()
    handleCloseMenu()
  }

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const value = event.target.value
    setEditedContent(value)
    setError(value.length > 1000)
  }

  const handleUpdate = () => {
    if (!error) {
      onUpdate(comment.id, editedContent)
    }
  }

  const handleDelete = () => {
    onDelete(comment.id)
    handleCloseMenu()
  }

  return (
    <Card variant='outlined' sx={{ position: 'relative' }}>
      <CardContent>
        <Box sx={{ display: 'flex', alignItems: 'center', mb: 0.5 }}>
          {comment.team_id === opponentRecruitingTeamId ? (
            <Chip
              label='募集チーム'
              size='small'
              sx={{ bgcolor: 'primary.main', color: 'white', mr: 1 }}
            />
          ) : (
            <Typography variant='body2' sx={{ mr: 1, fontWeight: 'bold' }}>
              チーム: {comment.team_name}
            </Typography>
          )}
        </Box>
        <Typography variant='body2' gutterBottom sx={{ fontWeight: 'bold' }}>
          投稿者: {comment.user_name}
        </Typography>
        {isEditing ? (
          <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
            <TextField
              fullWidth
              multiline
              rows={4}
              value={editedContent}
              onChange={handleChange}
              error={error}
              helperText={
                <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
                  {error ? 'コメントは1000文字以内で入力してください。' : <span>&nbsp;</span>}
                  <span>{`${editedContent.length}/1000`}</span>
                </Box>
              }
            />
            <PrimaryButton onClick={handleUpdate} disabled={error}>
              完了
            </PrimaryButton>
          </Box>
        ) : (
          <Typography variant='body1' sx={{ whiteSpace: 'pre-wrap', wordBreak: 'break-word' }}>
            {comment.content}
          </Typography>
        )}
      </CardContent>
      {user && user.current_team_id === comment.team_id && (
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
            {!isEditing && <MenuItem onClick={handleEdit}>編集</MenuItem>}
            <MenuItem onClick={handleDelete}>削除</MenuItem>
          </Menu>
        </>
      )}
    </Card>
  )
}

export default OpponentRecruitingComment
