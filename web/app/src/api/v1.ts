import axios from "axios";
export default axios.create({
    baseURL: '/api/v1/',
    timeout: 5000,
    validateStatus: function (status) {
        return status >= 200 && status < 500;
    }
});
