import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import { AppLayout } from './components/AppLayout';
import { TooltipProvider } from './components/ui/tooltip';
import { HomePage } from './components/HomePage';
import { SettingsPage } from './components/SettingsPage';
import { ChatPage } from './components/ChatPage';

const router = createBrowserRouter([
  {
    path: "/",
    element: <AppLayout />,
    children: [{
      path: "/",
      element: <HomePage />,
    }, {
        path: "/settings",
        element: <SettingsPage />
      }, {
        path: "/chat",
        element: <ChatPage />
      }]
  },
]);

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <TooltipProvider>
      <RouterProvider router={router} />
    </TooltipProvider>
  </React.StrictMode>,
)
