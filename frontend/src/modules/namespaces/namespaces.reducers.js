import { createSlice } from '@reduxjs/toolkit';

const initialState = { loading: false, namespaces: [], clonedNamespace: {} };

export const namespacesSlice = createSlice({
  name: 'namespaces',
  initialState,
  reducers: {
    getAllNamespaces: (state, action) => {
      return { ...state, loading: false, namespaces: action.payload };
    },
    startFetching: (state) => {
      return { ...state, loading: true };
    },
    cloneNamespace: (state, action) => {
      return { ...state, clonedNamespace: action.payload }
    }
  },
});

export const actions = namespacesSlice.actions;
export default namespacesSlice.reducer;
