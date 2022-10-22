import * as React from 'react';
import Box from '@mui/material/Box';
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';
import { Link as RouterLink } from "react-router-dom";
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { green } from '@mui/material/colors';

const theme = createTheme({
  palette: {
    primary: {
      main: green[400],
    },
    secondary: {
      main: '#e8f5e9',
    },
  },
});

export default function VerticalTabs() {
  const [value, setValue] = React.useState(0);
  
  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setValue(newValue);
  };
 
  return (
    <ThemeProvider theme={theme}>
    <Box
      sx={{ flexGrow: 1, bgcolor: 'background.paper',  height: 224 ,marginTop:5}}
    >
      <Tabs
        orientation="vertical"
        value={value}
        variant="scrollable"
        onChange={handleChange}
        // aria-label="Vertical tabs example"
        sx={{ borderRight: 1, borderColor: 'divider' }}
      >
        <Tab label="เพิ่มผู้ป่วย"   component={RouterLink} to="/" />
        <Tab label="ผู้ป่วยของแพทย์"   component={RouterLink} to="/patient" />     
     </Tabs>
   </Box>
   </ThemeProvider>
  );
}


