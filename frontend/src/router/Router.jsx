import React, { Suspense } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Namespaces from '../modules/namespaces';
import routePaths from './routePaths';

const RouterMod = () => {
  const Configs = React.lazy(() => import('../modules/configs'));
  const Deployments = React.lazy(() => import('../modules/deployments'));
  return (
    <BrowserRouter>
      <Suspense fallback={<div></div>}>
        <Routes>
          <Route exact path="/" name="namespaces" element={<Namespaces />} />
          <Route exact path={routePaths.namespaces} name="namespaces" element={<Namespaces />} />
          <Route
            exact
            path={routePaths.namespaceDeployments}
            name="Namespace Deployments"
            element={<Deployments />}
          />
          <Route
            exact
            path={routePaths.namespaceConfigs}
            name="Configs"
            element={<Configs />}
          />
        </Routes>
      </Suspense>
    </BrowserRouter>
  );
};

export default RouterMod;
