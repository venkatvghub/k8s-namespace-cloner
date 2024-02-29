import { useNavigate } from 'react-router';

const NamespaceMenuItems = (props) => {
  const namespace = props.row;
  const navigate = useNavigate();

  return (
    <div className="w-40 absolute ml-6 border-1 border-black bg-slate-50 text-base p-2">
      {namespace.cloned === 'true' ? (
        <>
          <div key={'deployments'} className="hover:font-semibold text-blue-500 hover:text-blue-700">
            <button onClick={() => navigate(`/namespaces/${namespace.namespace}/deployments`)}>View deployments</button>
          </div>
          <div key={'configs'} className="hover:font-semibold text-blue-500 hover:text-blue-700">
            <button onClick={() => navigate(`/namespaces/${namespace.namespace}/configs`)}>View configs</button>
          </div>
        </>
      ) : (
        <>
          <div key={'clone'} className="hover:font-semibold text-blue-500 hover:text-blue-700">
            <button onClick={() => props.setCloneNamespaceDialogContext(namespace.namespace)}>Clone namespace</button>
          </div>
        </>
      )}
    </div>
  );
};

export default NamespaceMenuItems;
