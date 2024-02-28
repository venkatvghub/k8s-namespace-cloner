const DeploymentMenuItems = (props) => {
  const deployment = props.row;

  const onScaleUpClick = () => {
    if(window.confirm("Are you sure you want to increase the replicas to 1?")) {
      props.increaseReplicas(deployment);
    }
  }

  return (
    <div className="w-40 absolute ml-6 border-1 border-black bg-slate-50 text-base p-2">
      <div key={'clone'} className="hover:font-semibold text-blue-500 hover:text-blue-700">
        <button onClick={() => props.setUpdateDeploymentDialogOpen(deployment)}>Update deployment</button>
      </div>
      {deployment.Replicas === 0 ? (
        <div key={'replicas'} className="hover:font-semibold text-blue-500 hover:text-blue-700">
          <button onClick={onScaleUpClick}>Scale up</button>
        </div>
      ) : (
        <></>
      )}
    </div>
  );
};

export default DeploymentMenuItems;
