import request from './request';

import axios from 'axios'; // 导入 axios 库

// 发送获取 table.json 文件数据的 GET 请求
export const fetchData = query => {
    const getRequest = axios({
        url: './table.json',
        method: 'get',
        params: query
    });
    
    return getRequest;
};

// 发送注册用户的 POST 请求
// export const registerUser = formData => {
//     return request({
//       url: `/api1/api/user/register`,
//       method: "post",
//       headers: {
//         "Content-Type": "multipart/form-data",
//     },
//       data: formData
//     });
//   }
export const registerUser = formData => {
    const registerRequest = axios({
        method: "post",
      url: "http://127.0.0.1:8888/api/user/register",
	  headers: {
        "Content-Type": "multipart/form-data",
      },
      data: formData,
    });
    
    return registerRequest;
};




export const loginUser = formData => {
    const loginRequest = axios({
        method: "post",
      url: "http://127.0.0.1:8888/api/user/login",
	  headers: {
        "Content-Type": "multipart/form-data",
      },
      data: formData,
    });
    
    return loginRequest;
};
