package firestore

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/wgir/gapsi-todo/internal/domain"
	"google.golang.org/api/iterator"
)

type taskRepository struct {
	client     *firestore.Client
	collection string
}

const tasksCollection = "tasks"

// NewTaskRepository creates a new Firestore implementation of domain.TaskRepository.
func NewTaskRepository(client *firestore.Client) domain.TaskRepository {
	return &taskRepository{
		client:     client,
		collection: tasksCollection,
	}
}

func (r *taskRepository) Create(ctx context.Context, task *domain.Task) error {
	docRef := r.client.Collection(r.collection).NewDoc()
	task.ID = docRef.ID
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	_, err := docRef.Set(ctx, task)
	return err
}

func (r *taskRepository) GetAll(ctx context.Context, status domain.TaskStatus) ([]domain.Task, error) {
	var tasks []domain.Task
	query := r.client.Collection(r.collection).Query

	if status != "" {
		query = query.Where("status", "==", string(status))
	}

	iter := query.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var task domain.Task
		if err := doc.DataTo(&task); err != nil {
			return nil, err
		}
		task.ID = doc.Ref.ID
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *taskRepository) GetByID(ctx context.Context, id string) (*domain.Task, error) {
	doc, err := r.client.Collection(r.collection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var task domain.Task
	if err := doc.DataTo(&task); err != nil {
		return nil, err
	}
	task.ID = doc.Ref.ID

	return &task, nil
}

func (r *taskRepository) Update(ctx context.Context, task *domain.Task) error {
	task.UpdatedAt = time.Now()
	_, err := r.client.Collection(r.collection).Doc(task.ID).Set(ctx, task)
	return err
}

func (r *taskRepository) Delete(ctx context.Context, id string) error {
	_, err := r.client.Collection(r.collection).Doc(id).Delete(ctx)
	return err
}
