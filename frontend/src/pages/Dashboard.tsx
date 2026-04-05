import React, { useState } from 'react';
import { MainLayout } from '../layouts/MainLayout';
import { TaskList } from '../features/tasks/components/TaskList';
import { TaskForm } from '../features/tasks/components/TaskForm';
import { Modal } from '../components/ui/Modal';
import { useTasks } from '../features/tasks/hooks/useTasks';
import type { Task } from '../features/tasks/schemas/task.schema';

export const Dashboard: React.FC = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingTask, setEditingTask] = useState<Task | null>(null);
  const { createTask, updateTask, isCreating, isUpdating } = useTasks();

  const handleOpenModal = (task: Task | null = null) => {
    setEditingTask(task);
    setIsModalOpen(true);
  };

  const handleCloseModal = () => {
    setEditingTask(null);
    setIsModalOpen(false);
  };

  const handleFormSubmit = async (data: Omit<Task, 'id' | 'createdAt' | 'updatedAt'>) => {
    try {
      if (editingTask?.id) {
        await updateTask({ id: editingTask.id, task: data });
      } else {
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        const { status, ...createData } = data as any;
        await createTask(createData);
      }
      handleCloseModal();
    } catch (error) {
      console.error('Submission failed:', error);
    }
  };

  return (
    <MainLayout onNewTask={() => handleOpenModal()}>
      <div className="space-y-8">
        <section className="flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
          <div>
            <h2 className="text-2xl font-bold tracking-tight text-slate-900 dark:text-white">Tasks Overview</h2>
            <p className="text-sm text-slate-500 dark:text-slate-400">Manage your daily tasks and productivity.</p>
          </div>
        </section>

        <section className="pb-10">
          <TaskList onEdit={handleOpenModal} />
        </section>
      </div>

      <Modal
        isOpen={isModalOpen}
        onClose={handleCloseModal}
        title={editingTask ? 'Edit Task' : 'Create New Task'}
        size="md"
      >
        <TaskForm
          onSubmit={handleFormSubmit}
          initialData={editingTask || {}}
          isLoading={isCreating || isUpdating}
        />
      </Modal>
    </MainLayout>
  );
};
