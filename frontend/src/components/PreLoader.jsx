import PropTypes from 'prop-types';
import { HashLoader } from 'react-spinners';

const PreLoader = ({ color = '#9CA3AF', loadingText = 'Fetching Data' }) => {
  return (
    <div className="flex w-full h-full justify-center text-center m-auto">
      <div className="flex my-6">
        <HashLoader size={18} color={color} />
        <span className="-mt-0.5 ml-2 font-medium tracking-wide uppercase" style={{ color }}>
          {loadingText}
        </span>
      </div>
    </div>
  );
};

PreLoader.propTypes = {
  color: PropTypes.string,
  loadingText: PropTypes.string,
};

export default PreLoader;
