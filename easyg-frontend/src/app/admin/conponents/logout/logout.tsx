'use client';
import Cookies from 'js-cookie';
export const logout = async (router: any) => {
  // 清除cookies中的token
  Cookies.remove('accessToken');
  Cookies.remove('refreshToken');

  // 重定向到登录页面
  router.push('/login');

  // 异步调用 logout API，不阻塞跳转
  fetch('/api/api/user/logout', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${Cookies.get('accessToken')}`
    }
  }).catch(err => {
    console.warn('后台登出请求失败:', err);
  });
}

