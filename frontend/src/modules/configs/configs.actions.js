import { actions } from './configs.reducers';
import apiCall from '../../utils/api-call-executor';

export const getAllConfigMaps = (namespaceName) => async (dispatch) => {
  try {
    dispatch(actions.startFetchingConfigMaps());
    const { data } = await apiCall('FETCH_ALL_NAMESPACE_CONFIG_MAPS', { namespaceName });
    if (Array.isArray(data)) {
      dispatch(actions.getAllConfigMaps(data));
    } else {
      dispatch(actions.getAllConfigMaps([]));
    }
  } catch (error) {
    console.log(error);
  }
};

export const getAllSecrets = (namespaceName) => async (dispatch) => {
  try {
    dispatch(actions.startFetchingSecrets());
    const { data } = await apiCall('FETCH_ALL_NAMESPACE_SECRETS', { namespaceName });
    if (Array.isArray(data)) {
      dispatch(actions.getAllSecrets(data));
    } else {
      dispatch(actions.getAllSecrets([]));
    }
  } catch (error) {
    console.log(error);
  }
};

export const updateConfigMap = (updateObject, namespaceName, configMap) => async (dispatch) => {
  try {
    await apiCall('UPDATE_CONFIG_MAP', { updateObject, namespaceName, configMap });
    dispatch(getAllConfigMaps(namespaceName));
  } catch (error) {
    console.log(error);
  }
};

export const updateSecret = (updateObject, namespaceName, secret) => async (dispatch) => {
  try {
    await apiCall('UPDATE_SECRET', { updateObject, namespaceName, secret });
    dispatch(getAllSecrets(namespaceName));
  } catch (error) {
    console.log(error);
  }
};
