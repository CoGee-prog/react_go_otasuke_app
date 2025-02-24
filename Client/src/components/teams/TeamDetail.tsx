import React, { useContext, useEffect, useState } from 'react'
import {
  Box,
  Card,
  CardContent,
  Typography,
  Divider,
  Grid,
  IconButton,
  ListItemIcon,
  ListItemText,
  Menu,
  MenuItem,
} from '@mui/material'
import { AuthContext } from 'src/contexts/AuthContext'
import { UpdateTeamApiResponse } from 'src/types/apiResponses'
import { TeamRole } from 'src/types/teamRole'
import { useRouter } from 'next/router'
import EditIcon from '@mui/icons-material/Edit'
import MoreVertIcon from '@mui/icons-material/MoreVert'
import { getPrefectureNameFromId } from 'src/utils/prefectures'
import { ArrowBack } from '@mui/icons-material'
import { Team } from 'src/types/team'
import TeamForm, { TeamFormData } from './TeamForm'
import { getLevelNameFromId } from 'src/utils/teamLevel'
import Link from 'next/link'

interface TeamProps {
  initialTeam: Team
}

// Team オブジェクトを TeamFormData に変換
function mapToFormData(team: Team): TeamFormData {
  return {
    name: team.name,
    prefecture_id: team.prefecture_id,
    level_id: team.level_id,
    home_page_url: team.home_page_url,
    other: team.other,
  }
}

const TeamDetail: React.FC<TeamProps> = ({ initialTeam }) => {
  const router = useRouter()
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null)
  const [team, setTeam] = useState(initialTeam)
  const [isEditing, setIsEditing] = useState(false)
  const { user } = useContext(AuthContext)

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

  const handleUpdateSuccess = (updatedData: UpdateTeamApiResponse) => {
    setIsEditing(false)
    setTeam(updatedData.team)
  }

  return (
    <Box sx={{ maxWidth: 600, mx: 'auto', my: 4, px: 2 }}>
      {!isEditing && (
        <>
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
              チーム詳細
            </Typography>
          </Grid>
        </>
      )}

      <Card
        sx={{
          mb: 2,
          position: 'relative',
        }}
      >
        {/* 編集中であれば編集フォームを出す */}
        {isEditing ? (
          <TeamForm
            isEditing={true}
            initialData={mapToFormData(team)}
            id={team.id}
            onUpdateSuccess={handleUpdateSuccess}
          />
        ) : (
          <>
            <CardContent>
              {[
                { label: 'チーム名', value: team.name },
                { label: '活動拠点', value: getPrefectureNameFromId(team.prefecture_id) },
                {
                  label: 'チームレベル',
                  value: getLevelNameFromId(team.level_id),
                },
                { label: 'ホームページリンク', value: team.home_page_url },
                { label: 'その他', value: team.other },
              ].map((item, index, arr) => (
                <Box key={index} sx={{ my: 1 }}>
                  <Typography variant='body1' sx={{ fontWeight: 'bold' }} gutterBottom>
                    {item.label}
                  </Typography>
                  {item.label === 'ホームページリンク' ? (
                    <Typography
                      variant='body1'
                      gutterBottom
                      color='primary'
                      sx={{ textDecoration: 'underline', cursor: 'pointer' }}
                      component='a'
                      href={item.value}
                      target='_blank'
                      rel='noopener noreferrer'
                    >
                      {item.value}
                    </Typography>
                  ) : (
                    <Typography variant='body1' gutterBottom>
                      {item.value}
                    </Typography>
                  )}
                  {index < arr.length - 1 && <Divider />}
                </Box>
              ))}
            </CardContent>
          </>
        )}
        {user &&
        user.current_team_id === team.id &&
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
                <MenuItem onClick={toggleEdit}>
                  <ListItemIcon>
                    <EditIcon fontSize='small' />
                  </ListItemIcon>
                  <ListItemText>編集</ListItemText>
                </MenuItem>
              )}
            </Menu>
          </>
        ) : null}
      </Card>
    </Box>
  )
}

export default TeamDetail
