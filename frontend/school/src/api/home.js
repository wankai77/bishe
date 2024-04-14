import request from './request';

export function getUserInfo() {
    return request({
      url: `/api1/api/home/getInfo`,
    });
  }
