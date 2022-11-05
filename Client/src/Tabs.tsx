import * as React from "react";
import Box from "@mui/material/Box";
import Tab from "@mui/material/Tab";
import TabContext from "@mui/lab/TabContext";
import TabList from "@mui/lab/TabList";
import TabPanel from "@mui/lab/TabPanel";

export default function Tabs() {
  const [value, setValue] = React.useState("1");

  const handleChange = (event: React.SyntheticEvent, newValue: string) => {
    setValue(newValue);
  };

  return (
    <Box sx={{ width: "100%", typography: "body1" }}>
      <TabContext value={value}>
        <Box
          sx={{
            borderBottom: 1,
            borderColor: "divider",
            justifyContent: "center",
          }}
        >
          <TabList
            onChange={handleChange}
            aria-label="lab API tabs example"
            centered
          >
            <Tab label="対戦相手募集" value="1" />
            <Tab label="スケジュール管理" value="2" />
          </TabList>
        </Box>
        <TabPanel value="1">対戦相手募集</TabPanel>
        <TabPanel value="2">スケジュール管理</TabPanel>
      </TabContext>
    </Box>
  );
}
