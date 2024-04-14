import request from './request';

export function getActInfo() {
    return request({
      url: `/api1/api/audit/getrelinfo`,
    });
  }
