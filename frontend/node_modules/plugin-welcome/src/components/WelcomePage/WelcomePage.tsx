import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
//import Typography from '@material-ui/core/Typography';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'ระบบจองห้องผ่าตัด' };

 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`ยินดีต้อนรับเข้าสู่ ${profile.givenName || ':)'}`}
       subtitle="ผ่าแล้วเย็บคืนด้วย"
     ></Header>
     <Content>
       <ContentHeader title="ตารางการจองห้องผ่าตัด">
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
       <br></br>
       
       <Link component={RouterLink} to="/booking">
           <Button variant="contained" color="primary" >
             เพิ่มการจองห้อง
           </Button>
         </Link>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;