import axios from "axios";
import { useUserStore } from "../stores/user";
import request from './request';

export function getUserInfo() {
    return request({
      url: `/api1/api/home/getInfo`,
    });
  }
export const ResetPsd = formData => {
    const registerRequest = axios({
        method: "post",
      url: "http://127.0.0.1:8888/api/personal/resetPassword",
	  headers: {
        "Content-Type": "multipart/form-data",
      },
      data: formData,
    });
    
    return registerRequest;
};