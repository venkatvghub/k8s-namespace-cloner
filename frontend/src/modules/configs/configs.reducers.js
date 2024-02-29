import { createSlice } from '@reduxjs/toolkit';

const initialState = { loadingConfigMap: false, loadingSecrets: false, configMaps: [], secrets: [] };

export const configsSlice = createSlice({
  name: 'configs',
  initialState,
  reducers: {
    getAllConfigMaps: (state, action) => {
      return { ...state, loadingConfigMap: false, configMaps: action.payload };
    },
    startFetchingConfigMaps: (state) => {
      return { ...state, loadingConfigMap: true };
    },

    getAllSecrets: (state, action) => {
      return { ...state, loadingSecrets: false, secrets: action.payload };
    },
    startFetchingSecrets: (state) => {
      return { ...state, loadingSecrets: true };
    }
  },
});

export const actions = configsSlice.actions;
export default configsSlice.reducer;
