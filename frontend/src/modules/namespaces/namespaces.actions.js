import { actions } from './namespaces.reducers';
import apiCall from '../../utils/api-call-executor';

export const getAllNamespaces = () => async (dispatch) => {
  try {
    dispatch(actions.startFetching());
    const { data } = await apiCall('FETCH_ALL_NAMESPACES');
    if (Array.isArray(data)) {
      dispatch(actions.getAllNamespaces(data));
    } else {
      dispatch(actions.getAllNamespaces([]));
    }
  } catch (error) {
    // console.log(error);
  }
};

export const cloneNamespace = (targetNamespace, namespaceName) => async (dispatch) => {
  try {
    const { data } = await apiCall('CLONE_NAMESPACE', { targetNamespace, namespaceName });
    dispatch(actions.cloneNamespace(data));
  } catch (error) {
    // console.log(error);
  }
};
