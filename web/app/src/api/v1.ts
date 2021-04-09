import axios from "axios";
export const ApiBase = '/api/v1/'
export default axios.create({
    baseURL: ApiBase,
    validateStatus: function (status) {
        return status >= 200 && status < 500;
    }
});
