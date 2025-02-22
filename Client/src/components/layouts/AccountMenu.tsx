import * as React from 'react'
import Box from '@mui/material/Box'
import Avatar from '@mui/material/Avatar'
import Menu from '@mui/material/Menu'
import MenuItem from '@mui/material/MenuItem'
import ListIcon from '@mui/icons-material/List'
import ListItemIcon from '@mui/material/ListItemIcon'
import Divider from '@mui/material/Divider'
import GroupIcon from '@mui/icons-material/Group'
import IconButton from '@mui/material/IconButton'
import Typography from '@mui/material/Typography'
import Tooltip from '@mui/material/Tooltip'
import Settings from '@mui/icons-material/Settings'
import Logout from '@mui/icons-material/Logout'
import ChangeCircleIcon from '@mui/icons-material/ChangeCircle'
import GroupAddIcon from '@mui/icons-material/GroupAdd'
import { useContext } from 'react'
import { AuthContext } from 'src/contexts/AuthContext'
import Link from 'next/link'
import { TeamRole } from 'src/types/teamRole'
import { InfoOutlined } from '@mui/icons-material'

export default function AccountMenu() {
  const [accountAnchorEl, setAccountAnchorEl] = React.useState<null | HTMLElement>(null)
  const [teamAnchorEl, setTeamAnchorEl] = React.useState<null | HTMLElement>(null)
  const accountMenuOpen = Boolean(accountAnchorEl)
  const teamMenuOpen = Boolean(teamAnchorEl)

  const handleTeamClick = (event: React.MouseEvent<HTMLElement>) => {
    if (isTeamAdminUser) {
      setTeamAnchorEl(event.currentTarget)
    }
  }
  const handleTeamMenuClose = () => {
    setTeamAnchorEl(null)
  }
  const handleAccountClick = (event: React.MouseEvent<HTMLElement>) => {
    setAccountAnchorEl(event.currentTarget)
  }
  const handleAccountMenuClose = () => {
    setAccountAnchorEl(null)
  }
  const handleLogout = () => {
    logout()
    setAccountAnchorEl(null)
  }
  const { user, logout } = useContext(AuthContext)
  console.log(user && user.current_team_role)
  const isTeamAdminUser =
    user &&
    (user.current_team_role === TeamRole.ADMIN || user.current_team_role === TeamRole.SUB_ADMIN)

  return (
    <React.Fragment>
      <Box sx={{ display: 'flex', alignItems: 'center', textAlign: 'center' }}>
        <Tooltip title={isTeamAdminUser ? 'Team menus' : ''}>
          <Typography
            onClick={handleTeamClick}
            sx={{
              minWidth: 100,
              color: 'white',
              cursor: isTeamAdminUser ? 'pointer' : 'default',
              '&:hover': isTeamAdminUser
                ? {
                    color: 'lightgray',
                    textDecoration: 'underline',
                  }
                : undefined,
            }}
          >
            {user?.current_team_name}
          </Typography>
        </Tooltip>
        <Tooltip title='Account settings'>
          <IconButton
            onClick={handleAccountClick}
            size='small'
            sx={{ ml: 2 }}
            aria-controls={accountMenuOpen ? 'account-menu' : undefined}
            aria-haspopup='true'
            aria-expanded={accountMenuOpen ? 'true' : undefined}
          >
            <Avatar sx={{ width: 32, height: 32 }}>{user?.name?.charAt(0)}</Avatar>
          </IconButton>
        </Tooltip>
      </Box>

      {/* Team Menu */}
      <Menu
        anchorEl={teamAnchorEl}
        id='team-menu'
        open={teamMenuOpen}
        onClose={handleTeamMenuClose}
        onClick={handleTeamMenuClose}
        PaperProps={{
          elevation: 0,
          sx: {
            overflow: 'visible',
            filter: 'drop-shadow(0px 2px 8px rgba(0,0,0,0.32))',
            mt: 1.5,
            '& .MuiAvatar-root': {
              width: 32,
              height: 32,
              ml: -0.5,
              mr: 1,
            },
            '&::before': {
              content: '""',
              display: 'block',
              position: 'absolute',
              top: 0,
              right: 14,
              width: 10,
              height: 10,
              bgcolor: 'background.paper',
              transform: 'translateY(-50%) rotate(45deg)',
              zIndex: 0,
            },
          },
        }}
        transformOrigin={{ horizontal: 'right', vertical: 'top' }}
        anchorOrigin={{ horizontal: 'right', vertical: 'bottom' }}
      >
        <Link href={`/teams/${user?.current_team_id}`} passHref>
          <MenuItem onClick={handleTeamMenuClose}>
            <ListItemIcon>
              <InfoOutlined fontSize='small' />
            </ListItemIcon>
            チーム詳細
          </MenuItem>
        </Link>
        <Link href='/opponent_recruitings/my_team' passHref>
          <MenuItem onClick={handleTeamMenuClose}>
            <ListItemIcon>
              <ListIcon fontSize='small' />
            </ListItemIcon>
            自チーム募集一覧
          </MenuItem>
        </Link>
        <MenuItem onClick={handleTeamMenuClose}>
          <ListItemIcon>
            <GroupIcon fontSize='small' />
          </ListItemIcon>
          メンバー招待
        </MenuItem>
      </Menu>

      {/* Account Menu */}
      <Menu
        anchorEl={accountAnchorEl}
        id='account-menu'
        open={accountMenuOpen}
        onClose={handleAccountMenuClose}
        onClick={handleAccountMenuClose}
        PaperProps={{
          elevation: 0,
          sx: {
            overflow: 'visible',
            filter: 'drop-shadow(0px 2px 8px rgba(0,0,0,0.32))',
            mt: 1.5,
            '& .MuiAvatar-root': {
              width: 32,
              height: 32,
              ml: -0.5,
              mr: 1,
            },
            '&::before': {
              content: '""',
              display: 'block',
              position: 'absolute',
              top: 0,
              right: 14,
              width: 10,
              height: 10,
              bgcolor: 'background.paper',
              transform: 'translateY(-50%) rotate(45deg)',
              zIndex: 0,
            },
          },
        }}
        transformOrigin={{ horizontal: 'right', vertical: 'top' }}
        anchorOrigin={{ horizontal: 'right', vertical: 'bottom' }}
      >
        <MenuItem onClick={handleAccountMenuClose}>
          <Avatar /> プロフィール
        </MenuItem>
        <Divider />
        <MenuItem onClick={handleAccountMenuClose}>
          <ListItemIcon>
            <ChangeCircleIcon fontSize='small' />
          </ListItemIcon>
          チーム切り替え
        </MenuItem>
        <Link href='/teams/create' passHref>
          <MenuItem onClick={handleAccountMenuClose}>
            <ListItemIcon>
              <GroupAddIcon fontSize='small' />
            </ListItemIcon>
            チーム作成
          </MenuItem>
        </Link>
        <MenuItem onClick={handleAccountMenuClose}>
          <ListItemIcon>
            <Settings fontSize='small' />
          </ListItemIcon>
          設定
        </MenuItem>
        <MenuItem onClick={handleLogout}>
          <ListItemIcon>
            <Logout fontSize='small' />
          </ListItemIcon>
          ログアウト
        </MenuItem>
      </Menu>
    </React.Fragment>
  )
}
