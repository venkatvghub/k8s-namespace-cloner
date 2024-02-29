import PropTypes from 'prop-types';
const DefaultLayout = ({ children }) => {
  return (
    <div className="w-full h-screen">
      <div className="relative flex flex-col flex-1 min-h-[90vh] max-auto mx-auto px-2 pt-4 lg:px-4 xl:px-0">
        {children}
      </div>
    </div>
  );
};

DefaultLayout.propTypes = {
  children: PropTypes.node,
};

export default DefaultLayout;
