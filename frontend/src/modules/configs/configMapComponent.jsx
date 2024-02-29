import { useEffect, useState, useMemo } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { getAllConfigMaps, updateConfigMap } from './configs.actions';
import DataTable from '../../components/DataTable';
import { ConfigMapsMenuItems } from './configMenuItems';
import UpdateConfigMapDialog from './updateConfigMapDialog';
import { Dialog } from '@mui/material';

const initialSortState = [
  { sortDirection: 'asc', accessor: 'key' },
  { sortDirection: 'none', accessor: 'value' },
];

const ConfigMapComponent = (props) => {
  const namespaceName = props.namespaceName;
  const [sort, setSort] = useState(initialSortState);
  const [page, setPage] = useState(1);
  const [currentConfig, setCurrentConfig] = useState({});
  const [open, setOpen] = useState(false);
  const pageSize = 10;
  const configs = useSelector((state) => state.configs);
  const configMaps = configs.configMaps.map(configMap => {
    return {
      key: Object.keys(configMap.data)[0],
      value: Object.values(configMap.data)[0],
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
        return <ConfigMapsMenuItems row={row} setUpdateConfigMapsDialogOpen={setUpdateConfigMapsDialogOpen} />;
      },
    },
  ];
  const columns = useMemo(() => headers, [sort]);

  const setUpdateConfigMapsDialogOpen = (configMap) => {
    setCurrentConfig(configMap);
    setOpen(true);
  };

  const handleClose = () => {
    setCurrentConfig({});
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
        getAllConfigMaps(namespaceName),
    );
  }, [dispatch, page, sort, namespaceName]);

  const onSubmit = (_e, key, value) => {
    const updateObject = {};
    updateObject[key] = value;
    dispatch(
      updateConfigMap(updateObject, namespaceName, configs.configMaps[0].name),
    );
  };

  return (
    <div className="w-11/12 mx-auto">
      <Dialog fullWidth={true} maxWidth="sm" open={open} fullScreen={false} onClose={handleClose}>
        <UpdateConfigMapDialog config={currentConfig} onSubmit={onSubmit} />
      </Dialog>
      <DataTable
        columns={columns}
        data={configMaps || []}
        isLoading={configs.loadingConfigMap || false}
        TableHeading="Config Maps"
        TableDescription="List of config maps"
        page={page}
        setPage={setPage}
        pageSize={pageSize}
        onHeaderClick={columnHeaderClick}
      />
    </div>
  );
};

export default ConfigMapComponent;
