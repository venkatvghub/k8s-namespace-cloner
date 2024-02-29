import { configureStore } from '@reduxjs/toolkit';
import configsReducers from '../modules/configs/configs.reducers';
import deploymentsReducers from '../modules/deployments/deployments.reducers';
import namespacesReducers from '../modules/namespaces/namespaces.reducers';

const createStore = (preloadedState) =>
  configureStore({
    reducer: {
      namespaces: namespacesReducers,
      deployments: deploymentsReducers,
      configs: configsReducers,
    },
    preloadedState,
});

export default createStore;
