import { useInfiniteQuery, useMutation, useQueryClient, type InfiniteData } from '@tanstack/react-query';
import { tasksApi } from '../api/tasks.api';
import type { Task } from '../schemas/task.schema';

export const useTasks = () => {
  const queryClient = useQueryClient();

  const {
    data,
    error,
    fetchNextPage,
    hasNextPage,
    isFetching,
    isFetchingNextPage,
    status,
  } = useInfiniteQuery<Task[], Error, InfiniteData<Task[]>, string[], string | undefined>({
    queryKey: ['tasks'],
    queryFn: ({ pageParam }) => tasksApi.getTasks(pageParam, 2),
    initialPageParam: undefined,
    getNextPageParam: (lastPage) => {
      // If there's no lastPage or it's empty, we're done
      if (!lastPage || lastPage.length === 0) return undefined;
      
      // If we got fewer items than requested (assume 10 or current limit), we're done
      // The user previously changed this limit to '2' in the UI for testing
      if (lastPage.length < 2) return undefined;

      return lastPage[lastPage.length - 1].id;
    },
  });

  const createTaskMutation = useMutation({
    mutationFn: tasksApi.createTask,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] });
    },
  });

  const completeTaskMutation = useMutation({
    mutationFn: (id: string) => tasksApi.completeTask(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] });
    },
  });

  const deleteTaskMutation = useMutation({
    mutationFn: (id: string) => tasksApi.deleteTask(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] });
    },
  });

  const updateTaskMutation = useMutation({
    mutationFn: ({ id, task }: { id: string; task: Partial<Task> }) => tasksApi.updateTask(id, task),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] });
    },
  });

  return {
    tasks: data?.pages.flat() ?? [],
    status,
    error,
    fetchNextPage,
    hasNextPage,
    isFetching,
    isFetchingNextPage,
    createTask: createTaskMutation.mutateAsync,
    completeTask: completeTaskMutation.mutateAsync,
    deleteTask: deleteTaskMutation.mutateAsync,
    updateTask: updateTaskMutation.mutateAsync,
    isCreating: createTaskMutation.isPending,
    isCompleting: completeTaskMutation.isPending,
    isDeleting: deleteTaskMutation.isPending,
    isUpdating: updateTaskMutation.isPending,
  };
};
