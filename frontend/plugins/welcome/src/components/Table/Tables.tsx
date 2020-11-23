import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntBooking } from '../../api/models/EntBooking';
import moment from 'moment';

const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [bookings, setBookings] = useState<EntBooking[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getBookings = async () => {
     const res = await api.listBooking({ limit: 10, offset: 0 });
     setLoading(false);
     setBookings(res);
   };
   getBookings();
 }, [loading]);
 
 const deleteBookings = async (id: number) => {
  const res = await api.deleteBooking({ id: id });
  setLoading(true);
};

 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">เลขที่</TableCell>
           <TableCell align="center">ชื่อแพทย์ผู้ทำการผ่าตัด</TableCell>
           <TableCell align="center">ชื่อคนไข้ที่เข้ารับการผ่าตัด</TableCell>
           <TableCell align="center">ห้องผ่าตัดที่ใช้</TableCell>
           <TableCell align="center">วันที่ใช้ห้องผ่าตัด</TableCell>
           <TableCell align="center">ลบการจองห้อง</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {bookings.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.doctorID?.doctorName}</TableCell>
             <TableCell align="center">{item.edges?.patientID?.patientName}</TableCell>
             <TableCell align="center">{item.edges?.operationroomID?.operationroomName}</TableCell>
             <TableCell align="center">{moment(item.date).format("DD/MM/YYYY")}</TableCell>
             <TableCell align="center">
             <Button
                 onClick={() => {
                   deleteBookings(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}