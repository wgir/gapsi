import React from 'react';
import { useForm, type SubmitHandler } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import type { Task } from '../schemas/task.schema';
import { Input } from '../../../components/ui/Input';
import { Button } from '../../../components/ui/Button';
import { z } from 'zod';

const formSchema = z.object({
  title: z.string().min(1, 'Title is required'),
  description: z.string().min(1, 'Description is required'),
  status: z.enum(['TODO', 'DONE', 'CANCELLED']).default('TODO'),
});

type TaskFormData = z.infer<typeof formSchema>;

interface TaskFormProps {
  onSubmit: (data: TaskFormData) => Promise<void>;
  initialData?: Partial<Task>;
  isLoading?: boolean;
}

export const TaskForm: React.FC<TaskFormProps> = ({ onSubmit, initialData, isLoading }) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<TaskFormData>({
    resolver: zodResolver(formSchema) as any,
    defaultValues: {
      title: initialData?.title || '',
      description: initialData?.description || '',
      status: (initialData?.status as any) || 'TODO',
    },
  });

  const handleFormSubmit: SubmitHandler<TaskFormData> = async (data) => {
    try {
      await onSubmit(data);
    } catch (error) {
      console.error('Failed to submit form:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit(handleFormSubmit)} className="space-y-4">
      <Input
        id="title"
        label="Task Title"
        placeholder="Enter task title..."
        {...register('title')}
        error={errors.title?.message}
      />
      
      <div className="space-y-1.5">
        <label
          htmlFor="description"
          className="block text-sm font-medium text-slate-700 dark:text-slate-300"
        >
          Description
        </label>
        <textarea
          id="description"
          rows={3}
          className="flex w-full rounded-lg border border-slate-200 bg-white px-3 py-2 text-sm ring-offset-white placeholder:text-slate-500 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-400 focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-800 dark:bg-slate-950 dark:ring-offset-slate-950 dark:placeholder:text-slate-400 dark:focus-visible:ring-primary-800"
          placeholder="Detailed task description..."
          {...register('description')}
        />
        {errors.description?.message && (
          <p className="text-xs font-medium text-red-500 animate-in fade-in slide-in-from-top-1">
            {errors.description.message}
          </p>
        )}
      </div>

      {initialData?.id && (
        <div className="space-y-1.5">
          <label
            htmlFor="status"
            className="block text-sm font-medium text-slate-700 dark:text-slate-300"
          >
            Status
          </label>
          <select
            id="status"
            className="flex h-10 w-full rounded-lg border border-slate-200 bg-white px-3 py-2 text-sm ring-offset-white focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-400 focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-800 dark:bg-slate-950 dark:ring-offset-slate-950 dark:focus-visible:ring-primary-800"
            {...register('status')}
          >
            <option value="TODO">To Do</option>
            <option value="DONE">Done</option>
            <option value="CANCELLED">Cancelled</option>
          </select>
          {errors.status?.message && (
            <p className="text-xs font-medium text-red-500 animate-in fade-in slide-in-from-top-1">
              {errors.status.message}
            </p>
          )}
        </div>
      )}

      <div className="flex justify-end pt-2">
        <Button type="submit" isLoading={isLoading} className="w-full sm:w-auto">
          {initialData?.id ? 'Update Task' : 'Create Task'}
        </Button>
      </div>
    </form>
  );
};
