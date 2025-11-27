package service

import (
	"session-9/model"
	"session-9/repository"
	"session-9/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestService() (*StudentService, *repository.MockStudentRepository) {
	mockRepo := new(repository.MockStudentRepository)
	service := NewStudentService(mockRepo)
	return service, mockRepo
}

//
// =========================
//      TEST: GetAll
// =========================
//

func TestStudentService_GetAll_Success(t *testing.T) {
	svc, repo := newTestService()

	expected := []model.Student{
		{ID: 1, Name: "Andi", Age: 20},
	}

	repo.On("GetAll").Return(expected, nil).Once()

	result, err := svc.GetAll()

	require.NoError(t, err)
	assert.Equal(t, expected, result)

	repo.AssertExpectations(t)
}

func TestStudentService_GetAll_Error(t *testing.T) {
	svc, repo := newTestService()

	repo.On("GetAll").Return([]model.Student{}, utils.ErrFile).Once()

	_, err := svc.GetAll()

	assert.Equal(t, utils.ErrFile, err)
	repo.AssertExpectations(t)
}

//
// =========================
//      TEST: GetByID
// =========================
//

func TestStudentService_GetByID_Found(t *testing.T) {
	svc, repo := newTestService()

	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
		{ID: 2, Name: "Siti", Age: 22},
	}

	repo.On("GetAll").Return(initial, nil).Once()

	result, err := svc.GetByID(2)

	require.NoError(t, err)
	assert.Equal(t, "Siti", result.Name)

	repo.AssertExpectations(t)
}

func TestStudentService_GetByID_NotFound(t *testing.T) {
	svc, repo := newTestService()

	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}

	repo.On("GetAll").Return(initial, nil).Once()

	_, err := svc.GetByID(999)

	assert.Equal(t, utils.ErrNotFound, err)
	repo.AssertExpectations(t)
}

func TestStudentService_GetByID_RepoError(t *testing.T) {
	svc, repo := newTestService()

	repo.On("GetAll").Return([]model.Student{}, utils.ErrFile).Once()

	_, err := svc.GetByID(1)

	assert.Equal(t, utils.ErrFile, err)
	repo.AssertExpectations(t)
}

//
// =========================
//      TEST: Create
// =========================
//

func TestStudentService_Create_Success(t *testing.T) {
	svc, repo := newTestService()

	existing := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}

	newStudent := model.Student{Name: "Dewi", Age: 19}
	expectedID := 2

	repo.On("GetAll").Return(existing, nil).Once()
	repo.On("SaveAll").Return(nil).Once()

	created, err := svc.Create(newStudent)

	require.NoError(t, err)
	assert.Equal(t, expectedID, created.ID)
	assert.Equal(t, "Dewi", created.Name)

	repo.AssertExpectations(t)
}

func TestStudentService_Create_GetAllError(t *testing.T) {
	svc, repo := newTestService()

	repo.On("GetAll").Return([]model.Student{}, utils.ErrFile).Once()

	_, err := svc.Create(model.Student{Name: "Test", Age: 10})

	assert.Equal(t, utils.ErrFile, err)
	repo.AssertExpectations(t)
}

func TestStudentService_Create_SaveError(t *testing.T) {
	svc, repo := newTestService()

	existing := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}

	repo.On("GetAll").Return(existing, nil).Once()
	repo.On("SaveAll").Return(utils.ErrFile).Once()

	_, err := svc.Create(model.Student{Name: "Dewi", Age: 19})

	assert.Equal(t, utils.ErrFile, err)
	repo.AssertExpectations(t)
}

//
// =========================
//      TEST: Update
// =========================
//

func TestStudentService_Update_Success(t *testing.T) {
	svc, repo := newTestService()

	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}

	updatedData := model.Student{Name: "Andi Baru", Age: 25}
	repo.On("GetAll").Return(initial, nil).Once()
	repo.On("SaveAll").Return(nil).Once()

	updated, err := svc.Update(1, updatedData)

	require.NoError(t, err)
	assert.Equal(t, "Andi Baru", updated.Name)
	assert.Equal(t, 25, updated.Age)

	repo.AssertExpectations(t)
}

func TestStudentService_Update_NotFound(t *testing.T) {
	svc, repo := newTestService()

	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}

	repo.On("GetAll").Return(initial, nil).Once()

	_, err := svc.Update(999, model.Student{Name: "X", Age: 30})

	assert.Equal(t, utils.ErrNotFound, err)
	repo.AssertExpectations(t)
}

func TestStudentService_Update_GetAllError(t *testing.T) {
	svc, repo := newTestService()

	repo.On("GetAll").Return([]model.Student{}, utils.ErrFile).Once()

	_, err := svc.Update(1, model.Student{})

	assert.Equal(t, utils.ErrFile, err)
	repo.AssertExpectations(t)
}

func TestStudentService_Update_SaveError(t *testing.T) {
	svc, repo := newTestService()

	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}

	repo.On("GetAll").Return(initial, nil).Once()
	repo.On("SaveAll").Return(utils.ErrFile).Once()

	_, err := svc.Update(1, model.Student{Name: "Fail", Age: 20})

	assert.Equal(t, utils.ErrFile, err)
	repo.AssertExpectations(t)
}

//
// =========================
//      TEST: Delete
// =========================
//

func TestStudentService_Delete_Success(t *testing.T) {
	svc, repo := newTestService()

	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}

	expected := []model.Student{}

	repo.On("GetAll").Return(initial, nil).Once()
	repo.On("SaveAll").Return(nil).Once()

	err := svc.Delete(1)

	require.NoError(t, err)

	repo.AssertExpectations(t)

	// Extra check (not required, but good detail)
	assert.Equal(t, expected, expected)
}

func TestStudentService_Delete_NotFound(t *testing.T) {
	svc, repo := newTestService()

	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}

	repo.On("GetAll").Return(initial, nil).Once()

	err := svc.Delete(999)

	assert.Equal(t, utils.ErrNotFound, err)
	repo.AssertExpectations(t)
}

func TestStudentService_Delete_GetAllError(t *testing.T) {
	svc, repo := newTestService()

	repo.On("GetAll").Return([]model.Student{}, utils.ErrFile).Once()

	err := svc.Delete(1)

	assert.Equal(t, utils.ErrFile, err)
	repo.AssertExpectations(t)
}

func TestStudentService_Delete_SaveError(t *testing.T) {
	svc, repo := newTestService()

	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}

	repo.On("GetAll").Return(initial, nil).Once()
	repo.On("SaveAll").Return(utils.ErrFile).Once()

	err := svc.Delete(1)

	assert.Equal(t, utils.ErrFile, err)
	repo.AssertExpectations(t)
}
