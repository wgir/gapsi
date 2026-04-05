import axios from 'axios';

const axiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080',
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add interceptors for error handling or auth if needed
axiosInstance.interceptors.response.use(
  (response) => response,
  (error) => {
    // Global error handler
    const message = error.response?.data?.message || 'An unexpected error occurred';
    console.error(`[API Error]: ${message}`);
    return Promise.reject(error);
  }
);

export default axiosInstance;
