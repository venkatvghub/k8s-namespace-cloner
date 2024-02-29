import { createSlice } from '@reduxjs/toolkit';

const initialState = { loading: false, deployments: [] };

export const deploymentsSlice = createSlice({
  name: 'namespaceDeployments',
  initialState,
  reducers: {
    getAllNamespaceDeployments: (state, action) => {
      return { ...state, loading: false, deployments: action.payload };
    },
    startFetching: (state) => {
      return { ...state, loading: true };
    },
  },
});

export const actions = deploymentsSlice.actions;
export default deploymentsSlice.reducer;
