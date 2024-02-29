import { useState } from 'react';
import PropTypes from 'prop-types';
import { useTable, useSortBy } from 'react-table';
import { IoMdMore } from 'react-icons/io';
import { FaSort, FaSortDown, FaSortUp } from 'react-icons/fa';
import PreLoader from '../PreLoader';
import { Button } from '@mui/material';

const DataTable = ({
  columns,
  data,
  isLoading = false,
  TableHeading = 'Table Name',
  TableDescription = 'Enter Table description',
  page = 1,
  setPage,
  pageSize = 10,
  onHeaderClick,
}) => {
  const previousPage = () => {
    setPage((previousPage) => previousPage - 1);
  };
  const nextPage = () => {
    setPage((previousPage) => previousPage + 1);
  };
  const canPreviousPage = () => {
    return page !== 1;
  };
  const canNextPage = () => {
    return data.length === pageSize;
  };

  const { getTableProps, getTableBodyProps, headerGroups, rows, prepareRow } = useTable(
    {
      columns,
      data,
      manualSortBy: true,
    },
    useSortBy,
  );

  const [menuId, setMenuId] = useState('');

  if (isLoading) {
    return <PreLoader />;
  }

  const handleClickMore = (event, value) => {
    event.preventDefault();
    if (menuId !== value) {
      setMenuId(value);
    } else {
      setMenuId('');
    }
  };
  return (
    <>
      {headerGroups?.map((headerGroup) =>
        headerGroup.headers?.map((column) =>
          column?.Filter ? <div key={column.id}>{column.render('Filter')}</div> : null,
        ),
      )}
      <div className="mt-2 flex flex-col shadow-sm overflow-hidden border border-gray-200 sm:rounded-lg">
        <div className="flex justify-between px-4 py-5 sm:px-6 border-b">
          <div>
            <h3 className="text-2xl leading-6 font-medium text-gray-900">
              {TableHeading || 'Table'}
            </h3>
            <p className="mt-0 mb-0 max-w-2xl text-sm text-gray-500">
              {TableDescription || 'Data Table'}
            </p>
          </div>
        </div>

        <div className="-my-2 overflow-x-auto -mx-4 sm:-mx-6 lg:-mx-8">
          <div className="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8">
            <div className="">
              <table
                {...getTableProps()}
                className="min-w-full divide-y divide-gray-200 w-full text-sm text-left text-gray-500 dark:text-gray-400"
              >
                <thead className="text-xs text-gray-700 uppercase border-b bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                  {headerGroups?.map((headerGroup, idx) => (
                    <tr {...headerGroup.getHeaderGroupProps()} key={idx}>
                      {headerGroup.headers?.map((column, idx) => (
                        <th
                          scope="col"
                          key={idx}
                          className="group px-6 py-3 text-left text-base  text-gray-900 uppercase tracking-wider"
                          {...column.getHeaderProps(column.getSortByToggleProps(), {
                            style: {
                              minWidth: column.minWidth,
                              width: column.width,
                            },
                          })}
                          onClick={() => onHeaderClick(column)}
                        >
                          <div className="flex items-center justify-between">
                            {column.render('Header')}

                            {(column.disableSortBy !== true) && (
                              <span>
                                {column.sortDirection === 'asc' ? (
                                  <FaSortUp />
                                ) : column.sortDirection === 'desc' ? (
                                  <FaSortDown />
                                ) : (
                                  <FaSort />
                                )}
                              </span>
                            )}
                          </div>
                        </th>
                      ))}
                    </tr>
                  ))}
                </thead>
                <tbody
                  {...getTableBodyProps()}
                  className="bg-white text-gray-900 divide-y divide-gray-200"
                >
                  {rows?.map((row, i) => {
                    prepareRow(row);

                    return (
                      <tr {...row.getRowProps()} key={i}>
                        {row.cells?.map((cell, idx) => {
                          return (
                            <td
                              key={idx}
                              {...cell.getCellProps()}
                              className="text-xs px-6 py-4 whitespace-nowrap"
                            >
                              <div className="text-base font-semibold">
                                {cell.column.Header !== 'Action' && (cell.row.original[cell.column.id] || '-')}
                              </div>
                              <div className="flex">
                                {cell.column.Header === 'Action' && (
                                  <div className="flex">
                                    <button
                                      onClick={(event) =>
                                        handleClickMore(event, cell.row.id)
                                      }
                                    >
                                      <IoMdMore className="w-6 h-6" />
                                    </button>
                                    {cell.row.id === menuId && cell.column.Cell(cell.row.original)}
                                  </div>
                                )}
                              </div>
                            </td>
                          );
                        })}
                      </tr>
                    );
                  })}
                </tbody>
              </table>
            </div>
          </div>
        </div>
        <div className="p-3 flex bg-gray-50 items-center justify-between border-t border-gray-200">
          <div className="flex-1 flex justify-between">
            <Button onClick={() => previousPage()} disabled={!canPreviousPage()}>
              Previous
            </Button>
            <Button onClick={() => nextPage()} disabled={!canNextPage()}>
              Next
            </Button>
          </div>
        </div>
      </div>
    </>
  );
};

DataTable.propTypes = {
  columns: PropTypes.array.isRequired,
  data: PropTypes.array.isRequired,
  isLoading: PropTypes.bool,
  TableHeading: PropTypes.string,
  TableDescription: PropTypes.string,
  page: PropTypes.number,
  setPage: PropTypes.func,
  pageSize: PropTypes.number,
  onHeaderClick: PropTypes.func,
};

export default DataTable;
