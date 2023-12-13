import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Header from './../components/Header';
import SideBar from './../components/SideBar';

interface UserData {
  ID: number;
  Name: string;
  Email: string;
  Password: string;
  CreatedAt: string;
  UpdatedAt: string;
}

export const UserIndex = () => {
  const [userId, setUserId] = useState<number>(1);
  const [userData, setUserData] = useState<UserData | null>(null);

  useEffect(() => {
    axios
      .get(`http://localhost:8080/users/1`)
      .then((response) => {
        setUserData(response.data);
      })
      .catch((error) => {
        console.error('Error fetching data: ', error);
      });
  }, [userId]); // userIdが変更されたときにuseEffectを再実行

  const handleUserIdChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUserId(Number(event.target.value)); // 入力された値を数値に変換してuserIdを更新
  };

  if (!userData) return <div>Loading...</div>;

  return (
    <div>
      <Header title="User Information" />
      <div style={{ float: 'right' }} className="sidebar">
        {/* Add className to adjust SideBar */}
        <SideBar />
      </div>
      <h1>User Data</h1>
      <input type="number" value={userId} onChange={handleUserIdChange} placeholder="Enter User ID" />
      <p>ID: {userData.ID}</p>
      <p>Name: {userData.Name}</p>
      <p>Email: {userData.Email}</p>
      <p>Created At: {userData.CreatedAt}</p>
      <p>Updated At: {userData.UpdatedAt}</p>
    </div>
  );
};