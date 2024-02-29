export const ConfigMapsMenuItems = (props) => {
  const configMap = props.row;

  return (
    <div className="w-40 absolute ml-6 border-1 border-black bg-slate-50 text-base p-2">
      <div key={'clone'} className="hover:font-semibold text-blue-500 hover:text-blue-700">
        <button onClick={() => props.setUpdateConfigMapsDialogOpen(configMap)}>Update config map</button>
      </div>
    </div>
  );
};

export const SecretsMenuItems = (props) => {
  const secret = props.row;

  return (
    <div className="w-40 absolute ml-6 border-1 border-black bg-slate-50 text-base p-2">
      <div key={'clone'} className="hover:font-semibold text-blue-500 hover:text-blue-700">
        <button onClick={() => props.setUpdateSecretsDialogOpen(secret)}>Update secret</button>
      </div>
    </div>
  );
};
