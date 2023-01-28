package store

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/common"
	api "github.com/bytebase/bytebase/backend/legacyapi"
)

// TaskMessage is the message for tasks.
type TaskMessage struct {
	ID int

	// Standard fields
	CreatorID int
	CreatedTs int64
	UpdaterID int
	UpdatedTs int64

	// Related fields
	PipelineID int
	StageID    int
	InstanceID int
	// Could be empty for creating database task when the task isn't yet completed successfully.
	DatabaseID          *int
	TaskRunRawList      []*TaskRunMessage
	TaskCheckRunRawList []*taskCheckRunRaw

	// Domain specific fields
	Name              string
	Status            api.TaskStatus
	Type              api.TaskType
	Payload           string
	EarliestAllowedTs int64
	BlockedBy         []string
}

func (task *TaskMessage) toTask() *api.Task {
	return &api.Task{
		ID: task.ID,

		// Standard fields
		CreatorID: task.CreatorID,
		CreatedTs: task.CreatedTs,
		UpdaterID: task.UpdaterID,
		UpdatedTs: task.UpdatedTs,

		// Related fields
		PipelineID: task.PipelineID,
		StageID:    task.StageID,
		InstanceID: task.InstanceID,
		// Could be empty for creating database task when the task isn't yet completed successfully.
		DatabaseID: task.DatabaseID,

		// Domain specific fields
		Name:              task.Name,
		Status:            task.Status,
		Type:              task.Type,
		Payload:           task.Payload,
		EarliestAllowedTs: task.EarliestAllowedTs,
		BlockedBy:         task.BlockedBy,
	}
}

// GetTaskByID gets a task by ID.
func (s *Store) GetTaskByID(ctx context.Context, id int) (*api.Task, error) {
	task, err := s.GetTaskV2ByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get Task with ID %d", id)
	}
	composedTask, err := s.composeTask(ctx, task)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to compose task %+v", task)
	}
	return composedTask, nil
}

// FindTask finds a list of Task instances.
func (s *Store) FindTask(ctx context.Context, find *api.TaskFind) ([]*api.Task, error) {
	tasks, err := s.ListTasks(ctx, find)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find Task list with TaskFind[%+v]", find)
	}
	var composedTasks []*api.Task
	for _, task := range tasks {
		composedTask, err := s.composeTask(ctx, task)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to compose task %+v", task)
		}
		composedTasks = append(composedTasks, composedTask)
	}
	return composedTasks, nil
}

// PatchTask patches an instance of Task.
func (s *Store) PatchTask(ctx context.Context, patch *api.TaskPatch) (*api.Task, error) {
	taskRaw, err := s.UpdateTaskV2(ctx, patch)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to patch Task with TaskPatch[%+v]", patch)
	}
	task, err := s.composeTask(ctx, taskRaw)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to compose Task with taskRaw[%+v]", taskRaw)
	}
	return task, nil
}

// PatchTaskStatus patches a task status.
func (s *Store) PatchTaskStatus(ctx context.Context, patch *api.TaskStatusPatch) (*api.Task, error) {
	task, err := s.UpdateTaskStatusV2(ctx, patch)
	if err != nil {
		return nil, FormatError(err)
	}

	composedTask, err := s.composeTask(ctx, task)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to compose task %+v", task)
	}
	return composedTask, nil
}

// BatchPatchTaskStatus patches status for a list of tasks.
func (s *Store) BatchPatchTaskStatus(ctx context.Context, taskIDs []int, status api.TaskStatus, updaterID int) error {
	var ids []string
	for _, id := range taskIDs {
		ids = append(ids, fmt.Sprintf("%d", id))
	}
	query := fmt.Sprintf(`
		UPDATE task
		SET status = $1, updater_id = $2
		WHERE id IN (%s);
	`, strings.Join(ids, ","))
	if _, err := s.db.db.ExecContext(ctx, query, status, updaterID); err != nil {
		return FormatError(err)
	}
	return nil
}

//
// private functions
//

func (s *Store) composeTask(ctx context.Context, raw *TaskMessage) (*api.Task, error) {
	task := raw.toTask()

	creator, err := s.GetPrincipalByID(ctx, task.CreatorID)
	if err != nil {
		return nil, err
	}
	task.Creator = creator

	updater, err := s.GetPrincipalByID(ctx, task.UpdaterID)
	if err != nil {
		return nil, err
	}
	task.Updater = updater

	taskRunRawList, err := s.listTaskRun(ctx, &TaskRunFind{
		TaskID: &task.ID,
	})
	if err != nil {
		return nil, err
	}
	taskCheckRunFind := &api.TaskCheckRunFind{
		TaskID: &task.ID,
	}
	taskCheckRunRawList, err := s.listTaskCheckRun(ctx, taskCheckRunFind)
	if err != nil {
		return nil, err
	}
	for _, taskRunRaw := range taskRunRawList {
		taskRun := taskRunRaw.toTaskRun()
		creator, err := s.GetPrincipalByID(ctx, taskRun.CreatorID)
		if err != nil {
			return nil, err
		}
		taskRun.Creator = creator

		updater, err := s.GetPrincipalByID(ctx, taskRun.UpdaterID)
		if err != nil {
			return nil, err
		}
		taskRun.Updater = updater
		task.TaskRunList = append(task.TaskRunList, taskRun)
	}
	for _, taskCheckRunRaw := range taskCheckRunRawList {
		taskCheckRun := taskCheckRunRaw.toTaskCheckRun()
		creator, err := s.GetPrincipalByID(ctx, taskCheckRun.CreatorID)
		if err != nil {
			return nil, err
		}
		taskCheckRun.Creator = creator

		updater, err := s.GetPrincipalByID(ctx, taskCheckRun.UpdaterID)
		if err != nil {
			return nil, err
		}
		taskCheckRun.Updater = updater
		task.TaskCheckRunList = append(task.TaskCheckRunList, taskCheckRun)
	}

	blockedBy := []string{}
	dags, err := s.ListTaskDags(ctx, &TaskDAGFind{ToTaskID: &raw.ID})
	if err != nil {
		return nil, err
	}
	for _, dag := range dags {
		blockedBy = append(blockedBy, strconv.Itoa(dag.FromTaskID))
	}
	task.BlockedBy = blockedBy

	instance, err := s.GetInstanceByID(ctx, task.InstanceID)
	if err != nil {
		return nil, err
	}
	if instance == nil {
		return nil, errors.Errorf("instance not found with ID %v", task.InstanceID)
	}
	task.Instance = instance

	if task.DatabaseID != nil {
		database, err := s.GetDatabase(ctx, &api.DatabaseFind{ID: task.DatabaseID})
		if err != nil {
			return nil, err
		}
		if database == nil {
			return nil, errors.Errorf("database not found with ID %v", task.DatabaseID)
		}
		task.Database = database
	}

	return task, nil
}

// GetTaskV2ByID gets a task by ID.
func (s *Store) GetTaskV2ByID(ctx context.Context, id int) (*TaskMessage, error) {
	tasks, err := s.ListTasks(ctx, &api.TaskFind{ID: &id})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get Task with ID %d", id)
	}
	if len(tasks) == 0 {
		return nil, nil
	} else if len(tasks) > 1 {
		return nil, errors.Errorf("found %v tasks with id %v", len(tasks), id)
	}
	return tasks[0], nil
}

// CreateTasksV2 creates a new task.
func (s *Store) CreateTasksV2(ctx context.Context, creates ...*api.TaskCreate) ([]*TaskMessage, error) {
	var query strings.Builder
	var values []interface{}
	var queryValues []string

	if _, err := query.WriteString(
		`INSERT INTO task (
			creator_id,
			updater_id,
			pipeline_id,
			stage_id,
			instance_id,
			database_id,
			name,
			status,
			type,
			payload,
			earliest_allowed_ts
		)
		VALUES
    `); err != nil {
		return nil, err
	}
	for i, create := range creates {
		if create.Payload == "" {
			create.Payload = "{}"
		}
		values = append(values,
			create.CreatorID,
			create.CreatorID,
			create.PipelineID,
			create.StageID,
			create.InstanceID,
			create.DatabaseID,
			create.Name,
			create.Status,
			create.Type,
			create.Payload,
			create.EarliestAllowedTs,
		)
		const count = 11
		queryValues = append(queryValues, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*count+1, i*count+2, i*count+3, i*count+4, i*count+5, i*count+6, i*count+7, i*count+8, i*count+9, i*count+10, i*count+11))
	}
	if _, err := query.WriteString(strings.Join(queryValues, ",")); err != nil {
		return nil, err
	}
	if _, err := query.WriteString(` RETURNING id, creator_id, created_ts, updater_id, updated_ts, pipeline_id, stage_id, instance_id, database_id, name, status, type, payload, earliest_allowed_ts`); err != nil {
		return nil, err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	var tasks []*TaskMessage
	rows, err := tx.QueryContext(ctx, query.String(), values...)
	if err != nil {
		return nil, FormatError(err)
	}
	defer rows.Close()
	for rows.Next() {
		task := &TaskMessage{}
		var databaseID sql.NullInt32
		if err := rows.Scan(
			&task.ID,
			&task.CreatorID,
			&task.CreatedTs,
			&task.UpdaterID,
			&task.UpdatedTs,
			&task.PipelineID,
			&task.StageID,
			&task.InstanceID,
			&databaseID,
			&task.Name,
			&task.Status,
			&task.Type,
			&task.Payload,
			&task.EarliestAllowedTs,
		); err != nil {
			return nil, FormatError(err)
		}
		if databaseID.Valid {
			val := int(databaseID.Int32)
			task.DatabaseID = &val
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, FormatError(err)
	}

	if err := tx.Commit(); err != nil {
		return nil, FormatError(err)
	}
	return tasks, nil
}

// ListTasks retrieves a list of tasks based on find.
func (s *Store) ListTasks(ctx context.Context, find *api.TaskFind) ([]*TaskMessage, error) {
	where, args := []string{"TRUE"}, []interface{}{}
	if v := find.ID; v != nil {
		where, args = append(where, fmt.Sprintf("id = $%d", len(args)+1)), append(args, *v)
	}
	if v := find.PipelineID; v != nil {
		where, args = append(where, fmt.Sprintf("pipeline_id = $%d", len(args)+1)), append(args, *v)
	}
	if v := find.StageID; v != nil {
		where, args = append(where, fmt.Sprintf("stage_id = $%d", len(args)+1)), append(args, *v)
	}
	if v := find.DatabaseID; v != nil {
		where, args = append(where, fmt.Sprintf("database_id = $%d", len(args)+1)), append(args, *v)
	}
	if v := find.StatusList; v != nil {
		list := []string{}
		for _, status := range *v {
			list = append(list, fmt.Sprintf("$%d", len(args)+1))
			args = append(args, status)
		}
		where = append(where, fmt.Sprintf("status in (%s)", strings.Join(list, ",")))
	}
	if v := find.TypeList; v != nil {
		var list []string
		for _, taskType := range *v {
			list = append(list, fmt.Sprintf("$%d", len(args)+1))
			args = append(args, taskType)
		}
		where = append(where, fmt.Sprintf("type in (%s)", strings.Join(list, ",")))
	}
	if v := find.Payload; v != "" {
		where = append(where, v)
	}

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
		SELECT
			id,
			creator_id,
			created_ts,
			updater_id,
			updated_ts,
			pipeline_id,
			stage_id,
			instance_id,
			database_id,
			name,
			status,
			type,
			payload,
			earliest_allowed_ts
		FROM task
		WHERE `+strings.Join(where, " AND ")+` ORDER BY id ASC`,
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*TaskMessage
	for rows.Next() {
		task := &TaskMessage{}
		if err := rows.Scan(
			&task.ID,
			&task.CreatorID,
			&task.CreatedTs,
			&task.UpdaterID,
			&task.UpdatedTs,
			&task.PipelineID,
			&task.StageID,
			&task.InstanceID,
			&task.DatabaseID,
			&task.Name,
			&task.Status,
			&task.Type,
			&task.Payload,
			&task.EarliestAllowedTs,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, FormatError(err)
	}
	if err := tx.Commit(); err != nil {
		return nil, FormatError(err)
	}
	return tasks, nil
}

// UpdateTaskV2 updates an existing task.
// Returns ENOTFOUND if task does not exist.
func (s *Store) UpdateTaskV2(ctx context.Context, patch *api.TaskPatch) (*TaskMessage, error) {
	set, args := []string{"updater_id = $1"}, []interface{}{patch.UpdaterID}
	if v := patch.DatabaseID; v != nil {
		set, args = append(set, fmt.Sprintf("database_id = $%d", len(args)+1)), append(args, *v)
	}
	if (patch.Statement != nil || patch.SchemaVersion != nil) && patch.Payload != nil {
		return nil, errors.Errorf("cannot set both statement/schemaVersion and payload for TaskPatch")
	}
	var payloadSet []string
	if v := patch.Statement; v != nil {
		payloadSet, args = append(payloadSet, fmt.Sprintf(`jsonb_build_object('statement', to_jsonb($%d::TEXT))`, len(args)+1)), append(args, *v)
	}
	if v := patch.SchemaVersion; v != nil {
		payloadSet, args = append(payloadSet, fmt.Sprintf(`jsonb_build_object('schemaVersion', to_jsonb($%d::TEXT))`, len(args)+1)), append(args, *v)
	}
	if len(payloadSet) != 0 {
		set = append(set, fmt.Sprintf(`payload = payload || %s`, strings.Join(payloadSet, "||")))
	}
	if v := patch.Payload; v != nil {
		payload := "{}"
		if *v != "" {
			payload = *v
		}
		set, args = append(set, fmt.Sprintf("payload = $%d", len(args)+1)), append(args, payload)
	}
	if v := patch.EarliestAllowedTs; v != nil {
		set, args = append(set, fmt.Sprintf("earliest_allowed_ts = $%d", len(args)+1)), append(args, *v)
	}
	args = append(args, patch.ID)

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	task := &TaskMessage{}
	// Execute update query with RETURNING.
	if err := tx.QueryRowContext(ctx, fmt.Sprintf(`
		UPDATE task
		SET `+strings.Join(set, ", ")+`
		WHERE id = $%d
		RETURNING id, creator_id, created_ts, updater_id, updated_ts, pipeline_id, stage_id, instance_id, database_id, name, status, type, payload, earliest_allowed_ts
	`, len(args)),
		args...,
	).Scan(
		&task.ID,
		&task.CreatorID,
		&task.CreatedTs,
		&task.UpdaterID,
		&task.UpdatedTs,
		&task.PipelineID,
		&task.StageID,
		&task.InstanceID,
		&task.DatabaseID,
		&task.Name,
		&task.Status,
		&task.Type,
		&task.Payload,
		&task.EarliestAllowedTs,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, &common.Error{Code: common.NotFound, Err: errors.Errorf("task not found with ID %d", patch.ID)}
		}
		return nil, FormatError(err)
	}

	if err := tx.Commit(); err != nil {
		return nil, FormatError(err)
	}

	return task, nil
}

// UpdateTaskStatusV2 updates the status of a task.
func (s *Store) UpdateTaskStatusV2(ctx context.Context, patch *api.TaskStatusPatch) (*TaskMessage, error) {
	task, err := s.GetTaskV2ByID(ctx, patch.ID)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, &common.Error{Code: common.NotFound, Err: errors.Errorf("task ID not found: %d", patch.ID)}
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	taskRunFind := &TaskRunFind{
		TaskID: &task.ID,
		StatusList: &[]api.TaskRunStatus{
			api.TaskRunRunning,
		},
	}
	taskRun, err := s.getTaskRunTx(ctx, tx, taskRunFind)
	if err != nil {
		return nil, err
	}
	if taskRun == nil {
		if patch.Status == api.TaskRunning {
			if err := s.createTaskRunImpl(ctx, tx, &TaskRunMessage{
				TaskID: task.ID,
				Name:   fmt.Sprintf("%s %d", task.Name, time.Now().Unix()),
				Type:   task.Type,
			}, patch.UpdaterID); err != nil {
				return nil, err
			}
		}
	} else {
		if patch.Status == api.TaskRunning {
			return nil, errors.Errorf("task is already running: %v", task.Name)
		}
		taskRunStatusPatch := &TaskRunStatusPatch{
			ID:        &taskRun.ID,
			UpdaterID: patch.UpdaterID,
			Code:      patch.Code,
			Result:    patch.Result,
			Comment:   patch.Comment,
		}
		switch patch.Status {
		case api.TaskDone:
			taskRunStatusPatch.Status = api.TaskRunDone
		case api.TaskFailed:
			taskRunStatusPatch.Status = api.TaskRunFailed
		case api.TaskPending:
		case api.TaskPendingApproval:
		case api.TaskCanceled:
			taskRunStatusPatch.Status = api.TaskRunCanceled
		}
		if _, err := s.patchTaskRunStatusImpl(ctx, tx, taskRunStatusPatch); err != nil {
			return nil, err
		}
	}

	// Updates the task
	// Build UPDATE clause.
	set, args := []string{"updater_id = $1"}, []interface{}{patch.UpdaterID}
	set, args = append(set, "status = $2"), append(args, patch.Status)
	var payloadSet []string
	if v := patch.Skipped; v != nil {
		payloadSet, args = append(payloadSet, fmt.Sprintf(`jsonb_build_object('skipped', to_jsonb($%d::BOOLEAN))`, len(args)+1)), append(args, *v)
	}
	if v := patch.SkippedReason; v != nil {
		payloadSet, args = append(payloadSet, fmt.Sprintf(`jsonb_build_object('skippedReason', to_jsonb($%d::TEXT))`, len(args)+1)), append(args, *v)
	}
	if len(payloadSet) != 0 {
		set = append(set, fmt.Sprintf(`payload = payload || %s`, strings.Join(payloadSet, "||")))
	}

	updatedTask := &TaskMessage{}
	// Execute update query with RETURNING.
	if err := tx.QueryRowContext(ctx, `
		UPDATE task
		SET `+strings.Join(set, ", ")+`
		WHERE id = `+fmt.Sprintf("%d", patch.ID)+`
		RETURNING id, creator_id, created_ts, updater_id, updated_ts, pipeline_id, stage_id, instance_id, database_id, name, status, type, payload, earliest_allowed_ts
	`,
		args...,
	).Scan(
		&updatedTask.ID,
		&updatedTask.CreatorID,
		&updatedTask.CreatedTs,
		&updatedTask.UpdaterID,
		&updatedTask.UpdatedTs,
		&updatedTask.PipelineID,
		&updatedTask.StageID,
		&updatedTask.InstanceID,
		&updatedTask.DatabaseID,
		&updatedTask.Name,
		&updatedTask.Status,
		&updatedTask.Type,
		&updatedTask.Payload,
		&updatedTask.EarliestAllowedTs,
	); err != nil {
		return nil, FormatError(err)
	}

	if err := tx.Commit(); err != nil {
		return nil, FormatError(err)
	}

	return updatedTask, nil
}
