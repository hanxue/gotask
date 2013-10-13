package task

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestTaskParser_Load(t *testing.T) {
	l := taskParser{"../examples"}
	funcs, err := l.Parse()

	assert.Equal(t, nil, err)
	assert.Equal(t, "github.com/jingweno/gotask/examples", funcs.ImportPath)
	assert.Equal(t, 1, len(funcs.Funcs))
}

func TestTaskParser_filterTaskFiles(t *testing.T) {
	files := []string{"file.go", "file_task.go", "task.go"}
	taskFiles := filterTaskFiles(files)

	assert.Equal(t, 1, len(taskFiles))
	assert.Equal(t, "file_task.go", taskFiles[0])
}

func TestTaskParser_parseTaskNames(t *testing.T) {
	names, _ := parseTaskNames("../examples/example_task.go")

	assert.Equal(t, 1, len(names))
	assert.Equal(t, "TaskHelloWorld", names[0])
}