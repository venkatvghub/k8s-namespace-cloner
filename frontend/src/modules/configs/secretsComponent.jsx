import { useEffect, useState, useMemo } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { getAllSecrets, updateSecret } from './configs.actions';
import DataTable from '../../components/DataTable';
import { SecretsMenuItems } from './configMenuItems';
import UpdateSecretDialog from './updateSecretDialog';
import { Dialog } from '@mui/material';

const initialSortState = [
  { sortDirection: 'none', accessor: 'namespace' },
  { sortDirection: 'none', accessor: 'Pod' },
  { sortDirection: 'asc', accessor: 'app' },
];

const SecretsComponent = (props) => {
  const namespaceName = props.namespaceName;
  const [sort, setSort] = useState(initialSortState);
  const [page, setPage] = useState(1);
  const [currentSecret, setCurrentSecret] = useState({});
  const [open, setOpen] = useState(false);
  const pageSize = 10;
  const configs = useSelector((state) => state.configs);
  const secrets = configs.secrets.map(secret => {
    return {
      key: Object.keys(secret.data)[0],
      value: Object.values(secret.data)[0],
    }
  });

  const headers = [
    {
      Header: 'Key',
      accessor: 'key',
      sortDirection:
        sort[0] && sort[0].accessor === 'key' ? sort[0].sortDirection : 'none',
    },
    {
      Header: 'Value',
      accessor: 'value',
      sortDirection: sort[1] && sort[1].accessor === 'value' ? sort[1].sortDirection : 'none',
    },
    {
      Header: 'Action',
      disableSortBy: true,
      Cell: (row) => {
        return <SecretsMenuItems row={row} setUpdateSecretsDialogOpen={setUpdateSecretsDialogOpen} />;
      },
    },
  ];
  const columns = useMemo(() => headers, [sort]);

  const setUpdateSecretsDialogOpen = (secret) => {
    setCurrentSecret(secret);
    setOpen(true);
  };

  const handleClose = () => {
    setCurrentSecret({});
    setOpen(false);
  };

  const columnHeaderClick = async (column) => {
    switch (column.sortDirection) {
      case 'none': {
        const noneCase = sort.map((item) =>
          item.accessor === column.id
            ? { ...item, sortDirection: 'asc' }
            : { ...item, sortDirection: 'none' },
        );
        setSort(noneCase);
        break;
      }
      case 'asc': {
        const asc = sort.map((item) =>
          item.accessor === column.id
            ? { ...item, sortDirection: 'desc' }
            : { ...item, sortDirection: 'none' },
        );
        setSort(asc);
        break;
      }
      case 'desc': {
        const desc = sort.map((item) =>
          item.accessor === column.id
            ? { ...item, sortDirection: 'asc' }
            : { ...item, sortDirection: 'none' },
        );
        setSort(desc);
        break;
      }
    }
  };

  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(
        getAllSecrets(namespaceName),
    );
  }, [dispatch, page, sort]);

  const onSubmit = (_e, key, value) => {
    const updateObject = {};
    updateObject[key] = value;
    dispatch(
      updateSecret(updateObject, namespaceName, configs.secrets[0].name),
    );
  };

  return (
    <div className="w-11/12 mx-auto">
      <Dialog fullWidth={true} maxWidth="sm" open={open} fullScreen={false} onClose={handleClose}>
        <UpdateSecretDialog secret={currentSecret} onSubmit={onSubmit} />
      </Dialog>
      <DataTable
          columns={columns}
          data={secrets || []}
          isLoading={configs.loadingSecrets || false}
          TableHeading="Secrets"
          TableDescription="List of secrets"
          page={page}
          setPage={setPage}
          pageSize={pageSize}
          onHeaderClick={columnHeaderClick}
      />
    </div>
  );
};

export default SecretsComponent;
