import axios from "axios";
const apiClient = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BACKEND_URL,
});

export interface BaseResponse {
  status: "success" | "error";
  data: any;
  message: string;
  success: boolean;
}

export interface SuccessResponse<T> extends BaseResponse {
  data: T | null;
  error: false;
  success: true;
}

export interface ErrorResponse extends BaseResponse {
  data: null;
  error: true;
  success: false;
}

apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    const normalizedError = {
      message: "An error occurred",
      statusCode: error.response?.status || 500,
      data: {},
    };
    if (error.response) {
      normalizedError.message = error.response.data.message || "Request failed";
      normalizedError.data = error.response.data;
    } else if (error.request)
      normalizedError.message = "No response from server";
    else normalizedError.message = error.message || "Unknown error";
    return Promise.reject(normalizedError);
  },
);
export default apiClient;
