// middleware.ts
import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

export async function middleware(request: NextRequest) {
  const requestHeaders = new Headers(request.headers);
  requestHeaders.set('x-current-path', request.nextUrl.pathname);

  // 检查是否是管理员路由
  if (request.nextUrl.pathname.startsWith('/admin')) {
    const token = request.cookies.get('accessToken')?.value;
    
    if (!token) {
      return NextResponse.redirect(new URL('/login', request.url));
    }

    try {
      // 验证token
      const response = await fetch(`${request.nextUrl.origin}/api/api/user/verify`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        }
      });

      if (!response.ok) {
        // accessToken 校验失败，尝试用 refreshToken 获取新 accessToken
        const refreshToken = request.cookies.get('refreshToken')?.value;
        if (refreshToken) {
          try {
            const refreshRes = await fetch(`${request.nextUrl.origin}/api/api/user/refreshaccesstoken`, {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json'
              },
              body: JSON.stringify({ refreshToken })
            });
            if (refreshRes.ok) {
              const data = await refreshRes.json();
              const newAccessToken = data.token;
              if (newAccessToken) {
                // 设置新的 accessToken 到 cookie
                const response = NextResponse.next({
                  request: { headers: requestHeaders }
                });
                response.cookies.set('accessToken', newAccessToken, { path: '/' });
                // 重新校验 accessToken
                const verifyRes = await fetch(`${request.nextUrl.origin}/api/api/user/verify`, {
                  method: 'POST',
                  headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${newAccessToken}`
                  }
                });
                if (verifyRes.ok) {
                  return response;
                }
              }
            }
          } catch (e) {}
        }
        return NextResponse.redirect(new URL('/login', request.url));
      }
    } catch (error) {
      return NextResponse.redirect(new URL('/login', request.url));
    }
  }

  return NextResponse.next({
    request: {
      headers: requestHeaders,
    },
  });
}

export const config = {
  matcher: ['/((?!api|_next/static|_next/image|favicon.ico).*)'],
};
