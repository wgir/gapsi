import React from 'react';
import { motion } from 'framer-motion';
import { Trash2, Calendar, Edit3, ClipboardList, CheckCircle, XCircle } from 'lucide-react';
import type { Task } from '../schemas/task.schema';
import { Button } from '../../../components/ui/Button';

interface TaskItemProps {
  task: Task;
  onStatusChange: (id: string, status: 'TODO' | 'DONE' | 'CANCELLED') => void;
  onDelete: (id: string) => void;
  onEdit: (task: Task) => void;
  isUpdating?: boolean;
  isDeleting?: boolean;
}

export const TaskItem: React.FC<TaskItemProps> = ({
  task,
  onStatusChange,
  onDelete,
  onEdit,
  isUpdating,
  isDeleting,
}) => {
  const statusConfig = {
    TODO: {
      color: 'bg-slate-100 text-slate-700 dark:bg-slate-800 dark:text-slate-300',
      icon: <ClipboardList className="h-5 w-5" />,
      label: 'To Do',
    },
    DONE: {
      color: 'bg-green-100 text-green-700 dark:bg-green-900/30 dark:text-green-400',
      icon: <CheckCircle className="h-5 w-5" />,
      label: 'Done',
    },
    CANCELLED: {
      color: 'bg-red-100 text-red-700 dark:bg-red-900/30 dark:text-red-400',
      icon: <XCircle className="h-5 w-5" />,
      label: 'Cancelled',
    },
  };

  const config = statusConfig[task.status as keyof typeof statusConfig] || statusConfig.TODO;

  return (
    <motion.div
      layout
      initial={{ opacity: 0, y: 10 }}
      animate={{ opacity: 1, y: 0 }}
      exit={{ opacity: 0, scale: 0.95 }}
      className="group relative overflow-hidden rounded-xl border border-slate-100 bg-white p-5 shadow-sm transition-all hover:border-primary-100 hover:shadow-md dark:border-slate-800 dark:bg-slate-900 dark:hover:border-primary-900"
    >
      <div className="flex items-start gap-4">
        <div className={ts('mt-1 flex-shrink-0 transition-colors', config.color)}>
            {config.icon}
        </div>

        <div className="flex-grow min-w-0">
          <div className="flex items-start justify-between gap-2">
            <h4
              className={ts(
                'text-base font-semibold transition-all duration-200',
                task.status === 'DONE'
                  ? 'text-slate-400 line-through dark:text-slate-500'
                  : 'text-slate-900 dark:text-white'
              )}
            >
              {task.title}
            </h4>
            <span className={ts('inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium', config.color)}>
              {config.label}
            </span>
          </div>
          <p
            className={ts(
              'mt-1 text-sm leading-relaxed text-slate-600 line-clamp-2 dark:text-slate-400',
              task.status === 'DONE' && 'text-slate-400 opacity-60 dark:text-slate-600'
            )}
          >
            {task.description}
          </p>

          <div className="mt-4 flex items-center gap-4 text-xs font-medium text-slate-400 dark:text-slate-500">
            <div className="flex items-center gap-1.5">
              <Calendar size={14} />
              <span>{task.createdAt ? new Date(task.createdAt).toLocaleDateString() : 'Today'}</span>
            </div>
            
            <div className="flex items-center gap-2 ml-auto">
                {task.status !== 'DONE' && (
                    <button 
                        onClick={() => onStatusChange(task.id!, 'DONE')}
                        disabled={isUpdating}
                        className="text-green-600 hover:text-green-700 dark:text-green-500 dark:hover:text-green-400 transition-colors font-semibold"
                    >
                        Mark Done
                    </button>
                )}
                {task.status !== 'CANCELLED' && (
                     <button 
                         onClick={() => onStatusChange(task.id!, 'CANCELLED')}
                         disabled={isUpdating}
                         className="text-red-500 hover:text-red-600 dark:text-red-400 dark:hover:text-red-300 transition-colors font-semibold"
                     >
                         Cancel
                     </button>
                )}
            </div>
          </div>
        </div>

        <div className="flex flex-col gap-2 scale-90 group-hover:scale-100 opacity-0 group-hover:opacity-100 transition-all duration-200">
          <Button
            variant="ghost"
            size="icon"
            onClick={() => onEdit(task)}
            className="h-8 w-8 hover:bg-primary-50 hover:text-primary-600 dark:hover:bg-primary-900/20"
          >
            <Edit3 size={16} />
          </Button>
          <Button
            variant="ghost"
            size="icon"
            isLoading={isDeleting}
            onClick={() => onDelete(task.id!)}
            className="h-8 w-8 hover:bg-red-50 hover:text-red-600 dark:hover:bg-red-900/20"
          >
            <Trash2 size={16} />
          </Button>
        </div>
      </div>
    </motion.div>
  );
};

function ts(...classes: (string | boolean | undefined)[]) {
  return classes.filter(Boolean).join(' ');
}
