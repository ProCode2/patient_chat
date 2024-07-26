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
import { LoginPage } from './components/LoginPage';
import { SignupPage } from './components/SignupPage';

import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'

// Create a client
const queryClient = new QueryClient()

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
    }, {
      path: "/login",
      element: <LoginPage />
    }, {
      path: "/signup",
      element: <SignupPage />
    }]
  },
]);

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <TooltipProvider>
        <RouterProvider router={router} />
      </TooltipProvider>
    </QueryClientProvider>
  </React.StrictMode >,
)
