import React, { useState, useEffect } from 'react';
import { Sun, Moon, LayoutDashboard, Menu, X, Plus } from 'lucide-react';
import { Button } from '../components/ui/Button';

interface MainLayoutProps {
  children: React.ReactNode;
  onNewTask?: () => void;
}

export const MainLayout: React.FC<MainLayoutProps> = ({ children, onNewTask }) => {
  const [isDarkMode, setIsDarkMode] = useState(false);
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);

  useEffect(() => {
    if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
      document.documentElement.classList.add('dark');
      setIsDarkMode(true);
    } else {
      document.documentElement.classList.remove('dark');
      setIsDarkMode(false);
    }
  }, []);

  const toggleDarkMode = () => {
    if (isDarkMode) {
      document.documentElement.classList.remove('dark');
      localStorage.theme = 'light';
      setIsDarkMode(false);
    } else {
      document.documentElement.classList.add('dark');
      localStorage.theme = 'dark';
      setIsDarkMode(true);
    }
  };

  return (
    <div className="flex h-screen bg-slate-50 dark:bg-slate-950 transition-colors duration-300">
      {/* Sidebar Desktop */}
      <aside className="hidden w-64 flex-col border-r border-slate-200 bg-white dark:border-slate-800 dark:bg-slate-900 md:flex">
        <div className="flex h-16 items-center px-6 border-b border-slate-100 dark:border-slate-800">
          <div className="flex items-center gap-2 font-bold text-primary-600 dark:text-primary-400">
            <div className="h-8 w-8 rounded-lg bg-primary-600 dark:bg-primary-500" />
            <span className="text-xl tracking-tight text-slate-900 dark:text-white">Gapsi Todo</span>
          </div>
        </div>
        <nav className="flex-grow space-y-1 p-4">
          <a href="#" className="flex items-center gap-3 rounded-lg bg-primary-50 px-3 py-2 text-sm font-medium text-primary-700 dark:bg-primary-900/20 dark:text-primary-400 transition-colors">
            <LayoutDashboard size={20} />
            Dashboard
          </a>
        </nav>
      </aside>

      {/* Main Content */}
      <main className="flex flex-1 flex-col overflow-hidden">
        {/* Header */}
        <header className="flex h-16 items-center justify-between border-b border-slate-200 bg-white px-6 dark:border-slate-800 dark:bg-slate-900">
          <div className="flex items-center gap-4">
            <button className="md:hidden text-slate-500" onClick={() => setIsSidebarOpen(true)}>
              <Menu size={24} />
            </button>
            <h1 className="text-lg font-semibold text-slate-900 dark:text-white">Dashboard</h1>
          </div>
          <div className="flex items-center gap-3">
            <Button variant="secondary" size="md" onClick={onNewTask} className="hidden sm:flex">
              <Plus size={18} className="mr-1" />
              New Task
            </Button>
            <Button variant="ghost" size="icon" onClick={toggleDarkMode} className="text-slate-500">
              {isDarkMode ? <Sun size={20} /> : <Moon size={20} />}
            </Button>
          </div>
        </header>

        {/* Page Content */}
        <div className="flex-1 overflow-y-auto p-4 sm:p-6 lg:p-8">
          <div className="mx-auto max-w-5xl">
            {children}
          </div>
        </div>
      </main>

      {/* Sidebar Mobile Overlay */}
      {isSidebarOpen && (
        <div className="fixed inset-0 z-50 flex items-center justify-start md:hidden">
          <div className="absolute inset-0 bg-slate-900/40 backdrop-blur-sm" onClick={() => setIsSidebarOpen(false)} />
          <aside className="relative flex h-full w-64 flex-col bg-white dark:bg-slate-900">
            <div className="flex h-16 items-center justify-between border-b border-slate-100 px-6 dark:border-slate-800">
              <span className="text-xl font-bold text-slate-900 dark:text-white">Todo App</span>
              <button onClick={() => setIsSidebarOpen(false)}>
                <X size={24} className="text-slate-500" />
              </button>
            </div>
            <nav className="flex-grow p-4 space-y-2">
              <a href="#" className="flex items-center gap-3 rounded-lg bg-primary-50 px-3 py-2 text-sm font-medium text-primary-700 dark:bg-primary-900/20 dark:text-primary-400">
                <LayoutDashboard size={20} />
                Dashboard
              </a>
            </nav>
          </aside>
        </div>
      )}
    </div>
  );
};
