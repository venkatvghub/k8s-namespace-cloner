import { actions } from './deployments.reducers';
import apiCall from '../../utils/api-call-executor';

export const getAllNamespaceDeployments = (namespaceName) => async (dispatch) => {
  try {
    dispatch(actions.startFetching());
    const { data } = await apiCall('FETCH_ALL_NAMESPACE_DEPLOYMENTS', { namespaceName });
    if (Array.isArray(data?.deployments)) {
      dispatch(actions.getAllNamespaceDeployments(data.deployments));
    } else {
      dispatch(actions.getAllNamespaceDeployments([]));
    }
  } catch (error) {
    // console.log(error);
  }
};

export const updateDeploymentImage = (image, deployment) => async (dispatch) => {
  try {
    await apiCall('UPDATE_DEPLOYMENT_IMAGE', { image, deployment });
    dispatch(getAllNamespaceDeployments(deployment.namespace));
  } catch (error) {
    // console.log(error);
  }
};

export const increaseDeploymentReplicas = (deployment) => async (dispatch) => {
  try {
    await apiCall('INCREASE_DEPLOYMENT_REPLICAS', { deployment });
    dispatch(getAllNamespaceDeployments(deployment.namespace));
  } catch (error) {
    // console.log(error);
  }
};
