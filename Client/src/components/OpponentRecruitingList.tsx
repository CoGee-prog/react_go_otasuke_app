// pages/index.tsx

import {
  Container,
  Grid,
  Card,
  CardContent,
  Typography,
  CardActionArea,
  Pagination,
  Box, // Boxコンポーネントをインポート
} from '@mui/material'
import Link from 'next/link'
import { useState, useEffect } from 'react'
import { Page } from 'src/types/page'
import { getOpponentRecruitingsApiResponse } from 'src/types/apiResponses'
import { OpponentRecruiting } from 'src/types/opponentRecruiting'
import fetchAPI from 'src/utils/fetchApi'
import { useNavigateHome } from 'src/hooks/useNavigateHome'
import formatDateTime from 'src/utils/formatDateTime'

interface OpponentRecruitingListProps {
  initialRecruitings: OpponentRecruiting[]
  initialPage: Page
}

export const OpponentRecruitingList: React.FC<OpponentRecruitingListProps> = ({
  initialRecruitings,
  initialPage,
}) => {
  const [opponentRecruitings, setOpponentRecruitings] =
    useState<OpponentRecruiting[]>(initialRecruitings)
  const [page, setPage] = useState<number>(initialPage.number)
  const [totalPages, setTotalPages] = useState<number>(initialPage.total_pages)
  const navigateHome = useNavigateHome()

  useEffect(() => {
    // ページが変更されるたびに、新しいデータを取得して状態を更新する
    handleChangePage(null, page)
  }, [])

  const handleChangePage = async (event: React.ChangeEvent<unknown> | null, value: number) => {
    const options: RequestInit = {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    }
    fetchAPI<getOpponentRecruitingsApiResponse>(`/opponent_recruitings?page=${value}`, options)
      .then((responseData) => {
        setOpponentRecruitings(responseData.result.opponent_recruitings)
        setPage(responseData.result.page.number)
        setTotalPages(responseData.result.page.total_pages)
      })
      .catch((error) => {
        console.error(error)
        navigateHome()
      })
  }

  return (
    <Container maxWidth='lg'>
      <Box display='flex' justifyContent='center' marginTop={2}>
        <Pagination count={totalPages} page={page} onChange={handleChangePage} />
      </Box>
      <Grid container spacing={2} direction='column' alignItems='center' justifyContent='center'>
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
            {' '}
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
                }}
              >
                <CardActionArea>
                  <CardContent>
                    <Typography variant='h5' component='div'>
                      {recruitment.team.name}
                    </Typography>
                    <Typography sx={{ mb: 1.5 }} color='text.secondary'>
                      {formatDateTime(recruitment.start_time)} ~ {formatDateTime(recruitment.end_time)}
                    </Typography>
                    <Typography variant='body2'>
                      {recruitment.prefecture} - {recruitment.detail}
                    </Typography>
                  </CardContent>
                </CardActionArea>
              </Card>
            </Link>
          </Grid>
        ))}
      </Grid>
      <Box display='flex' justifyContent='center' marginTop={2}>
        <Pagination count={totalPages} page={page} onChange={handleChangePage} />
      </Box>
    </Container>
  )
}
