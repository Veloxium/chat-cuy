import axios from "axios";
import { useUserStore } from "@/store/userStore";

const baseURL = import.meta.env.VITE_API_URL

const axiosInstance = axios.create({
    baseURL,
    timeout: 10000,
    headers: {
        'Content-Type': 'application/json',
    },
});

axiosInstance.interceptors.request.use((config) => {
    const token = useUserStore.getState().user?.token;
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

axiosInstance.interceptors.response.use(
    (response) => response,
    (error) => {
        if (error.response?.status === 401) {
            useUserStore.getState().logout();
            window.location.href = "/";
        }
        return Promise.reject(error);
    }
);

export default axiosInstance;