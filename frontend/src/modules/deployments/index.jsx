import { useEffect, useState, useMemo } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import DefaultLayout from '../../Layout/DefaultLayout';

import { getAllNamespaceDeployments, increaseDeploymentReplicas, updateDeploymentImage } from './deployments.actions';
import { useParams } from 'react-router-dom';
import DataTable from '../../components/DataTable';
import UpdateDeploymentDialog from './updateDeploymentDialog';
import { Dialog } from '@mui/material';
import DeploymentMenuItems from './deploymentMenuItems';

const initialSortState = [
  { sortDirection: 'none', accessor: 'Name' },
  { sortDirection: 'asc', accessor: 'Namespace' },
  { sortDirection: 'none', accessor: 'Pod' },
  { sortDirection: 'none', accessor: 'Service' },
];

const Deployments = () => {
  const { namespaceName } = useParams();
  const [sort, setSort] = useState(initialSortState);
  const [currentDeployment, setCurrentDeployment] = useState({});
  const [open, setOpen] = useState(false);
  const [page, setPage] = useState(1);
  const pageSize = 10;
  const deployments = useSelector((state) => state.deployments);

  const headers = [
    {
      Header: 'Deployment',
      accessor: 'name',
      sortDirection:
        sort[0] && sort[0].accessor === 'name' ? sort[0].sortDirection : 'none',
    },
    {
      Header: 'Namespace',
      accessor: 'namespace',
      sortDirection: sort[1] && sort[1].accessor === 'namespace' ? sort[1].sortDirection : 'none',
    },
    {
      Header: 'Pod',
      accessor: 'pod',
      sortDirection:
        sort[2] && sort[2].accessor === 'pod' ? sort[2].sortDirection : 'none',
    },
    {
      Header: 'App',
      accessor: 'app',
      sortDirection:
        sort[2] && sort[2].accessor === 'app' ? sort[2].sortDirection : 'none',
    },
    {
      Header: 'Action',
      disableSortBy: true,
      Cell: (row) => {
        return <DeploymentMenuItems
          row={row}
          setUpdateDeploymentDialogOpen={setUpdateDeploymentDialogOpen}
          increaseReplicas={increaseReplicas}
        />;
      },
    },
  ];
  const columns = useMemo(() => headers, [sort]);

  const setUpdateDeploymentDialogOpen = (deployment) => {
    setCurrentDeployment(deployment);
    setOpen(true);
  };

  const handleClose = () => {
    setCurrentDeployment({});
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
      getAllNamespaceDeployments(namespaceName),
    );
  }, [dispatch, page, sort, namespaceName]);

  const onSubmit = (_e, image, deployment) => {
    dispatch(
      updateDeploymentImage(image, deployment),
    );
  };

  const increaseReplicas = (deployment) => {
    dispatch(
      increaseDeploymentReplicas(deployment)
    );
  };

  return (
    <DefaultLayout>
      <div className="w-11/12 mx-auto">
        <div>
          <Dialog fullWidth={true} maxWidth="sm" open={open} fullScreen={false} onClose={handleClose}>
            <UpdateDeploymentDialog deployment={currentDeployment} onSubmit={onSubmit} />
          </Dialog>
          <DataTable
            columns={columns}
            data={deployments.deployments || []}
            isLoading={deployments.loading || false}
            TableHeading="Deployments"
            TableDescription="List of deployments"
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

export default Deployments;
