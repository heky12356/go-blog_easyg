import axios from 'axios';
import Cookies from 'js-cookie';
import { Navigate } from 'react-router-dom';
const query = axios.create({ baseURL: '/api' });

var accessToken = Cookies.get("accessToken");
var refreshToken = Cookies.get("refreshToken");

// 请求拦截
query.interceptors.response.use(
    config => {
        const token = accessToken;
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    error => {
        return Promise.reject(error);
    }
);

// 响应拦截
query.interceptors.response.use(
    response => response,
    async (error) => {
        const originalRequest = error.config;
        if (error.response.status === 401 && !originalRequest._retry) {
            originalRequest._retry = true;
            try {
                const response = await axios.post('/api/api/user/refreshaccesstoken', { refreshToken: refreshToken });
                accessToken = response.data.token;
                query.defaults.headers.common['Authorization'] = `Bearer ${accessToken}`;
                Cookies.set("accessToken", response.data.token, { secure: true, sameSite: "Strict" });
                return query(originalRequest);
            } catch (refreshError) {
                return <Navigate to="/login" />
            }
        }
        return Promise.reject(error);
    }
);

export default query;