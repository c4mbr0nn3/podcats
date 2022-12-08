import axios from "axios";

const axiosOptions = {
  baseURL: `api/v1`,
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
};

const axiosClient = axios.create(axiosOptions);

export default axiosClient;
