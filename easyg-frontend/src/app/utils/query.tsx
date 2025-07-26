import axios from 'axios';
import Cookies from 'js-cookie';
const query = axios.create({ baseURL: '/api' });


// 请求拦截
query.interceptors.request.use(
    config => {
        const token = Cookies.get("accessToken");
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
                const refreshToken = Cookies.get("refreshToken");
                const response = await axios.post('/api/api/user/refreshaccesstoken', { refreshToken: refreshToken });
                const accessToken = response.data.token;
                Cookies.set("accessToken", accessToken, { secure: true, sameSite: "Strict" });
                originalRequest.headers.Authorization = `Bearer ${accessToken}`;
                return query(originalRequest);
            } catch (refreshError) {
                Cookies.remove("accessToken");
                Cookies.remove("refreshToken");
                window.location.href = "/login";
                return Promise.reject(refreshError);
            }
        }
        return Promise.reject(error);
    }
);

export default query;