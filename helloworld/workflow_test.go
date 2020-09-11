package helloworld

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
  // Create a WorkflowTestSuite
	testSuite := &testsuite.WorkflowTestSuite{}
  // Establish a new environment
	env := testSuite.NewTestWorkflowEnvironment()
	// Mock activity implementation
	env.OnActivity(HelloWorldActivity, mock.Anything, "World").Return("Hello World!", nil)
  // Execute the Workflow in the test environment
  env.ExecuteWorkflow(HelloWorldWorkflow, "World")
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	var result string
	require.NoError(t, env.GetWorkflowResult(&result))
	require.Equal(t, "Hello World!", result)
}