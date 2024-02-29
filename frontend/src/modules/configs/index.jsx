import { useParams } from 'react-router';
import DefaultLayout from '../../Layout/DefaultLayout';
import ConfigMapComponent from './configMapComponent';
import SecretsComponent from './secretsComponent';

const Configs = () => {
  const { namespaceName } = useParams();
  return (
    <DefaultLayout>
      <ConfigMapComponent namespaceName={namespaceName} />
      <SecretsComponent namespaceName={namespaceName} />
    </DefaultLayout>
  );
};

export default Configs;
