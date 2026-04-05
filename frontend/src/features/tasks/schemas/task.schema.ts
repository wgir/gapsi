import { z } from 'zod';

export const taskSchema = z.object({
  id: z.string().optional(),
  title: z.string().min(1, 'Title is required'),
  description: z.string().min(1, 'Description is required'),
  status: z.enum(['TODO', 'DONE', 'CANCELLED']).default('TODO'),
  createdAt: z.string().optional(),
  updatedAt: z.string().optional(),
});

export type Task = z.infer<typeof taskSchema>;

export interface TaskResponse {
  data: Task[];
  next_id: string | null;
}
