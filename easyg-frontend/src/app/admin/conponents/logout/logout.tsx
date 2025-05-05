import Cookies from 'js-cookie';
import { NextResponse } from 'next/server';

export const logout = async () => {
    const token = Cookies.get('accessToken');
    try {
        // 调用登出API
        const response = await fetch('/api/api/user/logout', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
          }
        });
  
        if (response.ok) {
          return NextResponse.redirect(new URL('/login', window.location.origin));
        }
      } catch (error) {
        console.error('登出失败:', error);
        return;
      }

    // 清除cookies中的token
    Cookies.remove('accessToken');
    Cookies.remove('refreshToken');
    
    // 重定向到登录页面
    window.location.href = '/login';
}

