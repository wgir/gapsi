import axiosInstance from '../../../api/axiosInstance';
import type { Task } from '../schemas/task.schema';

export const tasksApi = {
  getTasks: async (last_id?: string, limit: number = 10): Promise<Task[]> => {
    const params = new URLSearchParams();
    params.append('limit', limit.toString());
    if (last_id) {
      params.append('last_id', last_id);
    }
    const response = await axiosInstance.get<Task[]>(`/tasks?${params.toString()}`);
    
    // Safety check: if response data is not an array (e.g. empty body), return empty array
    if (!Array.isArray(response.data)) {
      return [];
    }
    
    return response.data;
  },

  createTask: async (task: Omit<Task, 'id' | 'createdAt' | 'updatedAt' | 'status'>): Promise<Task> => {
    const { data } = await axiosInstance.post<Task>('/tasks', task);
    return data;
  },

  completeTask: async (id: string): Promise<Task> => {
    const { data } = await axiosInstance.patch<Task>(`/tasks/${id}/complete`);
    return data;
  },

  deleteTask: async (id: string): Promise<void> => {
    await axiosInstance.delete(`/tasks/${id}`);
  },

  getTaskById: async (id: string): Promise<Task> => {
    const { data } = await axiosInstance.get<Task>(`/tasks/${id}`);
    return data;
  },

  updateTask: async (id: string, task: Partial<Task>): Promise<Task> => {
    const { data } = await axiosInstance.put<Task>(`/tasks/${id}`, task);
    return data;
  },
};
