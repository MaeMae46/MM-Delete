// index.tsx หรือ main.tsx
import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import Stock1 from './pages/Stock1';
import Stock2 from './pages/Stock2';
import Stock3 from './pages/Stock3';
import Stock4 from './pages/Stock4';
import Order from './pages/order';

const router = createBrowserRouter([
    { path: "/", element: <Stock1 /> },
    { path: "/Stock2", element: <Stock2 /> },
    { path: "/Stock3", element: <Stock3 /> },
    { path: "/Stock4", element: <Stock4 /> },
    { path: "/order", element: <Order /> }
  ]);
  
  
  ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>

                   <RouterProvider router={router} />

    </React.StrictMode>
  );