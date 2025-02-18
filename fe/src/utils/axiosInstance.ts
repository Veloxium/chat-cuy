import axios from "axios";

// const baseURL = import.meta.env.VITE_API_URL;

const API_URL = "http://localhost:8080";

const axiosInstance = axios.create({
   baseURL: API_URL,
   // timeout: 10000,
   // headers: {
   //    "Content-Type": "application/json",
   // },
});

export default axiosInstance;
