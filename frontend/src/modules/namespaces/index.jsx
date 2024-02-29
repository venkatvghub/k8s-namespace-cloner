import { useEffect, useState, useMemo } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import DefaultLayout from '../../Layout/DefaultLayout';

import { cloneNamespace, getAllNamespaces } from './namespaces.actions';
import DataTable from '../../components/DataTable';
import { Dialog } from '@mui/material';
import CloneNamespaceDialog from './cloneNamespaceDialog';
import NamespaceMenuItems from './namespaceMenuItems';

const initialSortState = [
  { sortDirection: 'none', accessor: 'namespace' },
  { sortDirection: 'none', accessor: 'Pod' },
  { sortDirection: 'asc', accessor: 'app' },
];

const Namespaces = () => {
  const [sort, setSort] = useState(initialSortState);
  const [currentNamespace, setCurrentNamespace] = useState('');
  const [open, setOpen] = useState(false);
  const [page, setPage] = useState(1);
  const pageSize = 10;
  const namespaces = useSelector((state) => state.namespaces);

  const headers = [
    {
      Header: 'Namespace Name',
      accessor: 'namespace',
      sortDirection:
        sort[0] && sort[0].accessor === 'name' ? sort[0].sortDirection : 'none',
    },
    {
      Header: 'Pod',
      accessor: 'Pod',
      sortDirection: sort[1] && sort[1].accessor === 'eventName' ? sort[1].sortDirection : 'none',
    },
    {
      Header: 'App',
      accessor: 'app',
      sortDirection:
        sort[2] && sort[2].accessor === 'app' ? sort[2].sortDirection : 'none',
    },
    {
      Header: 'Cloned',
      accessor: 'cloned',
      sortDirection:
        sort[2] && sort[2].accessor === 'cloned' ? sort[2].sortDirection : 'none',
    },
    {
      Header: 'Action',
      disableSortBy: true,
      Cell: (row) => {
        return <NamespaceMenuItems row={row} setCloneNamespaceDialogContext={setCloneNamespaceDialogContext} />;
      }
    },
  ];
  const columns = useMemo(() => headers, [sort]);

  const setCloneNamespaceDialogContext = (namespace) => {
    setCurrentNamespace(namespace);
    setOpen(true);
  };

  const handleClose = () => {
    setCurrentNamespace('');
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
      getAllNamespaces(),
    );
  }, [dispatch, page, sort]);

  const onCloneClick = (_e, targetNamespace, namespaceName) => {
    dispatch(
      cloneNamespace(targetNamespace, namespaceName),
    );
    
  };

  return (
    <DefaultLayout>
      <div className="w-11/12 mx-auto">
        <div className="">
          <Dialog fullWidth={true} maxWidth="sm" open={open} fullScreen={false} onClose={handleClose}>
            <CloneNamespaceDialog namespaceName={currentNamespace} onCloneClick={onCloneClick} />
          </Dialog>
          <DataTable
            columns={columns}
            data={namespaces.namespaces || []}
            isLoading={namespaces.loading || false}
            TableHeading="Namespaces"
            TableDescription="List of namespaces"
            page={page}
            setPage={setPage}
            pageSize={pageSize}
            onHeaderClick={columnHeaderClick}
          />
        </div>
      </div>
    </DefaultLayout>
  );
};

export default Namespaces;
