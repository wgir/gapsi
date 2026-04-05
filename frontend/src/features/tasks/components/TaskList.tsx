import React, { useRef, useEffect } from 'react';
import { AnimatePresence } from 'framer-motion';
import { useTasks } from '../hooks/useTasks';
import { TaskItem } from './TaskItem';
import type { Task } from '../schemas/task.schema';
import { Button } from '../../../components/ui/Button';

interface TaskListProps {
  onEdit: (task: Task) => void;
}

export const TaskList: React.FC<TaskListProps> = ({ onEdit }) => {
  const {
    tasks,
    status,
    error,
    fetchNextPage,
    hasNextPage,
    isFetchingNextPage,
    updateTask,
    deleteTask,
  } = useTasks();

  const observerRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    if (!hasNextPage || isFetchingNextPage) return;

    const observer = new IntersectionObserver(
      (entries) => {
        if (entries[0].isIntersecting) {
          fetchNextPage();
        }
      },
      { threshold: 1.0 }
    );

    if (observerRef.current) {
      observer.observe(observerRef.current);
    }

    return () => observer.disconnect();
  }, [hasNextPage, isFetchingNextPage, fetchNextPage]);

  if (status === 'pending') {
    return (
      <div className="space-y-4">
        {[1, 2, 3].map((i) => (
          <div key={i} className="h-28 w-full animate-pulse rounded-xl bg-slate-100 dark:bg-slate-800" />
        ))}
      </div>
    );
  }

  if (status === 'error') {
    return (
      <div className="flex flex-col items-center justify-center rounded-xl border border-red-100 bg-red-50 p-10 text-center dark:border-red-900/20 dark:bg-red-900/10">
        <p className="text-sm font-medium text-red-600 dark:text-red-400">
          Failed to load tasks: {(error as Error).message}
        </p>
        <Button
          variant="ghost"
          size="sm"
          className="mt-4 text-red-600 dark:text-red-400"
          onClick={() => window.location.reload()}
        >
          Retry
        </Button>
      </div>
    );
  }

  if (tasks.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center rounded-2xl border-2 border-dashed border-slate-100 p-16 text-center dark:border-slate-800">
        <div className="flex h-16 w-16 items-center justify-center rounded-full bg-slate-50 dark:bg-slate-800 text-slate-300">
            <svg className="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4" />
            </svg>
        </div>
        <h3 className="mt-4 text-lg font-semibold text-slate-900 dark:text-white">All set!</h3>
        <p className="mt-2 text-sm text-slate-500 dark:text-slate-400">
          No tasks found matching your filters.
        </p>
      </div>
    );
  }

  const handleStatusChange = async (id: string, newStatus: 'TODO' | 'DONE' | 'CANCELLED') => {
    try {
      await updateTask({ id, task: { status: newStatus } });
    } catch (err) {
      console.error('Failed to update status:', err);
    }
  };

  return (
    <div className="space-y-4">
      <AnimatePresence mode="popLayout" initial={false}>
        {tasks.map((task) => (
          <TaskItem
            key={task.id}
            task={task}
            onStatusChange={handleStatusChange}
            onDelete={() => deleteTask(task.id!)}
            onEdit={onEdit}
          />
        ))}
      </AnimatePresence>

      <div ref={observerRef} className="h-4 w-full">
        {isFetchingNextPage && (
          <div className="flex items-center justify-center py-4">
            <div className="h-6 w-6 animate-spin rounded-full border-2 border-primary-500 border-t-transparent" />
          </div>
        )}
      </div>
    </div>
  );
};
