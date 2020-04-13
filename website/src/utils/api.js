import axios from "axios"

const apiCall = ({ url, method, data }) =>
    new Promise((resolve, reject) => {
        axios({url: "http://localhost" + url, method: method, data: data })
        .then(resolve)
        .catch(reject)
    });

export default apiCall;
