import {
  Container,
  Grid,
  Card,
  CardContent,
  Typography,
  CardActionArea,
  Pagination,
  Box,
  Chip,
  Divider,
} from '@mui/material'
import Link from 'next/link'
import { useState, useEffect, useContext } from 'react'
import { Page } from 'src/types/page'
import { GetOpponentRecruitingsApiResponse } from 'src/types/apiResponses'
import { OpponentRecruiting } from 'src/types/opponentRecruiting'
import fetchAPI from 'src/utils/fetchApi'
import { useNavigateHome } from 'src/hooks/useNavigateHome'
import { formatTimeRange } from 'src/utils/formatDateTime'
import PrimaryButton from '../commons/PrimaryButton'
import { TeamRole } from 'src/types/teamRole'
import { AuthContext } from 'src/contexts/AuthContext'
import OpponentRecruitingSearchForm from './OpponentRecruitingSearchForm'

interface OpponentRecruitingListProps {
  initialRecruitings: OpponentRecruiting[]
  initialPage: Page
}

const OpponentRecruitingList: React.FC<OpponentRecruitingListProps> = ({
  initialRecruitings,
  initialPage,
}) => {
  const [opponentRecruitings, setOpponentRecruitings] =
    useState<OpponentRecruiting[]>(initialRecruitings)
  const [page, setPage] = useState<number>(initialPage.number)
  const [totalPages, setTotalPages] = useState<number>(initialPage.total_pages)
  const { user } = useContext(AuthContext)
  const [queryParams, setQueryParams] = useState<string>('')
  const navigateHome = useNavigateHome()

  useEffect(() => {
    // ページが変更されるたびに、新しいデータを取得して状態を更新する
    handleChangePage(null, page)
  }, [page, queryParams])

  const handleChangePage = async (event: React.ChangeEvent<unknown> | null, value: number) => {
		// 同じページの場合は何もしない
		if (page == value) {
			return;
		}
    const options: RequestInit = {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    }
    fetchAPI<GetOpponentRecruitingsApiResponse>(
      `/opponent_recruitings?page=${value}&${queryParams}`,
      options,
    )
      .then((responseData) => {
        setOpponentRecruitings(responseData.result.opponent_recruitings)
        setPage(responseData.result.page.number)
        setTotalPages(responseData.result.page.total_pages)
      })
      .catch((error) => {
        console.error(error)
        // ホームページに移動
        navigateHome()
      })
  }

  return (
    <Container maxWidth='lg' sx={{ marginBottom: 2 }}>
      <Grid
        container
        spacing={2}
        direction='column'
        alignItems='center'
        justifyContent='center'
        style={{ marginTop: '3px' }}
      >
        <Grid item xs={12} style={{ display: 'flex', justifyContent: 'center' }}>
          <Box sx={{ maxWidth: 500, width: '100%', textAlign: 'center' }}>
            {user &&
            (user.current_team_role === TeamRole.ADMIN ||
              user.current_team_role === TeamRole.SUB_ADMIN) ? (
              <Link href='/opponent_recruitings/create' passHref>
                <PrimaryButton>対戦相手を募集する</PrimaryButton>
              </Link>
            ) : (
              <p>チームの管理者か副管理者のみ対戦相手募集を作成できます</p>
            )}
          </Box>
        </Grid>
      </Grid>
      <Grid container justifyContent='center'>
        <Grid item xs={12} style={{ display: 'flex', justifyContent: 'center' }}>
          <Box sx={{ maxWidth: 'md', width: '100%' }}>
            <OpponentRecruitingSearchForm onSearch={(params) => setQueryParams(params)} />
          </Box>
        </Grid>
      </Grid>
      {totalPages > 0 ? (
        <Box display='flex' justifyContent='center' marginTop={2}>
          <Pagination
            count={totalPages}
            page={page}
            boundaryCount={1}
            siblingCount={2}
            onChange={handleChangePage}
          />
        </Box>
      ) : (
        <Typography variant='h6' component='p' sx={{ marginTop: 2, textAlign: 'center' }}>
          対戦相手募集がありません
        </Typography>
      )}
      <Grid
        container
        spacing={2}
        direction='column'
        alignItems='center'
        justifyContent='center'
        sx={{ marginTop: 0.5, marginBottom: 1 }}
      >
        {opponentRecruitings.map((recruitment) => (
          <Grid
            item
            xs={12}
            sm={6}
            md={4}
            lg={3}
            key={recruitment.id}
            style={{ maxWidth: 500, width: '100%' }}
          >
            <Link href={`/opponent_recruitings/${recruitment.id}`} passHref>
              <Card
                sx={{
                  maxWidth: 500,
                  margin: 'auto',
                  transition: '0.3s',
                  boxShadow: '0 8px 40px -12px rgba(0,0,0,0.3)',
                  '&:hover': {
                    boxShadow: '0 16px 70px -12.125px rgba(0,0,0,0.3)',
                  },
                  borderRadius: 2,
                  backgroundColor: recruitment.is_active ? 'white' : '#d0d0d0',
                  position: 'relative',
                }}
              >
                <CardActionArea>
                  <CardContent>
                    <Box sx={{ maxWidth: '85%', wordWrap: 'break-word' }}>
                      <Typography
                        variant='h6'
                        component='div'
                        gutterBottom
                        sx={{ fontWeight: 'bold' }}
                      >
                        {recruitment.title}
                      </Typography>
                    </Box>
                    <Typography sx={{ fontWeight: 'bold', mb: 1.5 }}>
                      {formatTimeRange(recruitment.start_time, recruitment.end_time)
                        .text.split(' ')
                        .map((part, index, array) => (
                          <span
                            key={index}
                            style={{
                              color:
                                index === 1 || index === array.length - 2
                                  ? index === 1
                                    ? formatTimeRange(recruitment.start_time, recruitment.end_time)
                                        .dayOfWeekColor
                                    : formatTimeRange(recruitment.start_time, recruitment.end_time)
                                        .endDayOfWeekColor
                                  : 'inherit',
                            }}
                          >
                            {part}{' '}
                          </span>
                        ))}
                    </Typography>
                    <Box sx={{ display: 'flex', alignItems: 'center', mb: 1 }}>
                      <Typography
                        variant='body2'
                        component='div'
                        sx={{ fontWeight: 'bold', mr: 1 }}
                      >
                        エリア:
                      </Typography>
                      <Typography variant='body2' component='div'>
                        {recruitment.prefecture}
                      </Typography>
                    </Box>
                    <Box sx={{ display: 'flex', alignItems: 'center', mb: 1 }}>
                      <Typography
                        variant='body2'
                        component='div'
                        sx={{ fontWeight: 'bold', mr: 1 }}
                      >
                        グラウンド有無:
                      </Typography>
                      <Chip
                        label={recruitment.has_ground ? 'あり' : 'なし'}
                        size='small'
                        color={recruitment.has_ground ? 'success' : 'error'}
                        sx={{ mr: 1 }}
                      />
                    </Box>
                    {recruitment.has_ground && (
                      <Box sx={{ display: 'flex', alignItems: 'center', mb: 1 }}>
                        <Typography
                          variant='body2'
                          component='div'
                          sx={{ fontWeight: 'bold', mr: 1 }}
                        >
                          グラウンド名:
                        </Typography>
                        <Typography variant='body2' component='div'>
                          {recruitment.ground_name}
                        </Typography>
                      </Box>
                    )}
                    <Divider sx={{ mb: 1 }} />
                    <Typography variant='body2' component='div'>
                      チーム: {recruitment.team.name}
                    </Typography>
                    <Typography variant='body2' component='div'>
                      レベル: {recruitment.team.level}
                    </Typography>
                  </CardContent>
                  <Chip
                    label={recruitment.is_active ? '募集中' : '募集終了'}
                    size='small'
                    color={recruitment.is_active ? 'primary' : 'default'}
                    sx={{ position: 'absolute', top: 8, right: 8 }}
                  />
                </CardActionArea>
              </Card>
            </Link>
          </Grid>
        ))}
      </Grid>
      {totalPages > 0 ? (
        <Box display='flex' justifyContent='center' marginTop={2}>
          <Pagination
            count={totalPages}
            page={page}
            boundaryCount={1}
            siblingCount={2}
            onChange={handleChangePage}
          />
        </Box>
      ) : null}
    </Container>
  )
}

export default OpponentRecruitingList
