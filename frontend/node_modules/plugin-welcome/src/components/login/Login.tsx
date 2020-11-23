import React, { useState,useEffect} from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import { EntUser } from '../../api/models/EntUser';

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'right',
   },
   margin: {
     margin: theme.spacing(3),
   },
   withoutLabel: {
     marginTop: theme.spacing(3),
   },
   textField: {
     width: '30ch',
   },
 }),
);

export default function Create() {
 const classes = useStyles();
 const profile = {thisName: 'โปรดเข้าสู่ระบบ เพื่อใช้งานระบบการจองห้องผ่าตัด' };
 const api = new DefaultApi();
 const [users, setUsers] = useState<EntUser[]>([]);
 const [userid, setUser] = useState(Number);
 const [loading, setLoading] = useState(true);
 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);
 
 useEffect(() => {
   const getUser = async () => {
    const res = await api.listUser({limit:10,offset :0});
    setLoading(false);
    setUsers(res);
   };
   getUser();

}, [loading]);

 const handleUserchange = (event: React.ChangeEvent<{value: unknown}>) => {
  setUser(event.target.value as number);
};

 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`${profile.thisName}`}
       subtitle="เลือกให้ตรงกับชื่อของตัวเอง"
     ></Header>
     <Content>
       <ContentHeader title="เลือกชื่อแพทย์ผู้ใช้งาน เพื่อเข้าสู่ระบบ">
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 เข้าสู่ระบบสำเร็จ!
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 This is a warning alert — check it out!
               </Alert>
             )}
           </div>
         ) : null}
       </ContentHeader>
       <center>
       <div className={classes.root}>
         <form noValidate autoComplete="off">
         <FormControl
             fullWidth
             className={classes.root}
             variant="outlined"
           >
             <InputLabel id="user-label"></InputLabel>
             <Select
              labelId="user-label"
              id="user"
              value={userid}
              onChange={handleUserchange}
              style={{width:600}}
        >
          {users.map((item: EntUser)=> (
            <MenuItem value={item.id}>{item.doctorName}</MenuItem>))}
            </Select>
           </FormControl>
 
           <div className={classes.margin}>
             <Button
               //style={{ marginLeft: 150 }}
               component={RouterLink}
               to="/welcomepage"
               variant="contained"
             >
               เข้าสู่ระบบ
             </Button>
           </div>
         </form>
       </div>
     </Content>
     </center>
   </Page>
 );
}