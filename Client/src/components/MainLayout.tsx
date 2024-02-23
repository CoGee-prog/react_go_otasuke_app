import { Box } from "@mui/material";
import Header from "./Header";
import Tabs from "./Tabs";
import { ReactNode } from "react";

interface MainLayoutProps {
  children: ReactNode
}

const MainLayout: React.FC<MainLayoutProps> = ({ children }) => {
  return (
    <Box>
      <Header />
      <Tabs />
      <main>{children}</main>
    </Box>
  )
}

export default MainLayout
