import apiMapping from './api-mapping';
import axios from 'axios';

const apiCall = (feature, params) => {
  let api = apiMapping.getRequestObject(feature, params);
  return axios(api);
};

export default apiCall;
