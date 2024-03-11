import { Grid, Box } from '@kuma-ui/core';
import Sidebar from '@/components/Sidebar';
import Search from '@/components/tweet/Search';

export default function Window({ children }: { children: React.ReactNode }) {
  return (
    <Grid gridTemplateColumns="repeat(3, auto)" gap={6}>
      <Box height="80px" >
        <Sidebar />
      </Box>
      <Box height="80px" >
        {children}
      </Box>
      <Box height="80px" >
        <Search />
      </Box>
    </Grid>
  )
}
