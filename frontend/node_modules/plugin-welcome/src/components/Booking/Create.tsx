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
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
//import Autocomplete from '@material-ui/lab/Autocomplete';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
//import Typography from '@material-ui/core/Typography';
import Select from '@material-ui/core/Select';

import { EntUser } from '../../api/models/EntUser';
import { EntPatient } from '../../api/models/EntPatient';
import { EntOperationroom } from '../../api/models/EntOperationroom';
//import { EntBooking } from '../../api/models/EntBooking';
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
 
//const initialUserState = {
// name: 'นายแพทย์เชี่ยวชาญ ฉับฉับ',
// email: 'doctor01@gmail.com',
//};
//const username={thisname:'นายแพทย์เชี่ยวชาญ ฉับฉับ'}
 
export default function Create() {
 const classes = useStyles();
 const profile = {thisName: 'ระบบจองห้องผ่าตัด' };
 const api = new DefaultApi();
 
 const [users, setUsers] = useState<EntUser[]>([]);
 const [patients, setPatients] = useState<EntPatient[]>([]);
 const [operationrooms, setOperationrooms] = useState<EntOperationroom[]>([]);
 //const [bookings, setBookings] = useState<EntBooking[]>([]);
 const [patientid, setPatient] = useState(Number);
 const [userid, setUser] = useState(Number);
 const [operationroomid, setOperationroom] = useState(Number);
 const [datetime, setdatetime] = useState(String);
 const [loading, setLoading] = useState(true);
 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);
 
 useEffect(() => {
   const getPatient = async () => {
     const res = await api.listPatient({limit: 10,offset :0});
     setLoading(false);
     setPatients(res);
     console.log(res);
   };
   getPatient();

   const getOperationroom = async () => {
     const res = await api.listOperationroom({limit:10,offset :0});
     setLoading(false);
     setOperationrooms(res);
   };
   getOperationroom();

   const getUser = async () => {
    const res = await api.listUser({limit:10,offset :0});
    setLoading(false);
    setUsers(res);
   };
   getUser();
//}, [loading]);

 //const getBooking = async () => {
 //  const res = await api.listBooking({limit: 10,offset: 0});
 //  setBookings(res);
 //};
 //getBooking();
}, [loading]);

 const handledatetimeChange = (event: any) => {
   setdatetime(event.target.value as string);
 };

 const handlePatientchange = (event: React.ChangeEvent<{value: unknown}>) => {
   setPatient(event.target.value as number);
 };

 const handleUserchange = (event: React.ChangeEvent<{value: unknown}>) => {
  setUser(event.target.value as number);
};
  
 const handleOperationroomchange = (event: React.ChangeEvent<{value: unknown}>) => {
    setOperationroom(event.target.value as number);
 };

 const createBooking = async ()=>{
   const booking ={
     user: userid,
     patient: patientid,
     operationroom: operationroomid,
     added: datetime +"T00:00:00Z"
   };
   console.log(booking);
   const res: any = await api.createBooking({booking : booking});
   setStatus(true);
   if(res.id != ''){
     setAlert(true);
   }else{
     setAlert(false);
   }
   //const timer = setTimeout(()=>{
  //   setStatus(false);
  // },1000);
 };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`${profile.thisName}`}
       subtitle="ผ่าแล้วเย็บคืนด้วย"
     ></Header>

     <Content>
       <ContentHeader title="กรอกรายการจองห้อง">
       
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 จองห้องผ่าตัดสำเร็จ! กดกลับ เพื่อดูผลการจอง!
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
              <font align="left">ชื่อ-นามสกุลแพทย์ผู้ทำการผ่าตัด</font>
         <FormControl
             fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <InputLabel id="user-label"></InputLabel>
             <center>
             <Select
              labelId="user-label"
              id="user"
              value={userid}
              onChange={handleUserchange}
              style={{width:400}}
        >
          {users.map((item: EntUser)=> (
            <MenuItem value={item.id}>{item.doctorName}</MenuItem>))}
            </Select>
            </center>
           </FormControl>
           ชื่อ-นามสกุล ผู้เข้ารับการผ่าตัด
         <FormControl
             fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <InputLabel id="patient-label"></InputLabel>
             <center>
             <Select
              labelId="patient-label"
              id="patient"
              value={patientid}
              onChange={handlePatientchange}
              style={{width:400}}
        >
          {patients.map((item: EntPatient)=> (
            <MenuItem value={item.id}>{item.patientName}</MenuItem>))}
            </Select>
            </center>
           </FormControl>
           ห้องผ่าตัดที่ต้องการจอง
         <FormControl
             fullWidth
             className={classes.margin}
             variant="outlined"
           >
            <InputLabel id="operationroom-label"></InputLabel>
            <center>
            <Select
                  labelId="operationroom-label"
                  id="operationroom"
                  value={operationroomid}
                  onChange={handleOperationroomchange}
                  style={{ width: 400 }}
                >
                {operationrooms.map((item: EntOperationroom) => (
                  <MenuItem value={item.id}>{item.operationroomName}</MenuItem>
                ))}
                </Select>
                </center>
           </FormControl>
           <FormControl className={classes.margin} >
                <TextField
                 id="date_time"
                  label="วันที่ใช้ห้องผ่าตัด"
                  type="date"
                  value={datetime}
                  onChange={handledatetimeChange}
                  className={classes.textField}
                  InputLabelProps={{
                   shrink: true,
                 }}
        
                />
                </FormControl>
 
           <div className={classes.margin}>
             <Button
               onClick={() => {
                 createBooking();
               }}
               variant="contained"
               color="primary"
             >
               ยืนยันการจองห้อง
             </Button>
             <Button
               style={{ marginLeft: 20 }}
               component={RouterLink}
               to="/"
               variant="contained"
             >
               ย้อนกลับ
             </Button>
           </div>
         </form>
       </div>
     </Content>
     </center>
   </Page>
 );
}