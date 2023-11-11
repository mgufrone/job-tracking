package application

import (
	"bytes"
	"context"
	json2 "encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/goravel/framework/contracts/auth"
	mocks2 "github.com/goravel/framework/contracts/auth/mocks"
	http2 "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/support/json"
	mock2 "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/mock"
	"io"
	"mgufrone.dev/job-tracking/app/http/requests/application"
	"mgufrone.dev/job-tracking/app/http/requests/job"
	"mgufrone.dev/job-tracking/app/models"
	"mgufrone.dev/job-tracking/app/service"
	"mgufrone.dev/job-tracking/app/service/mocks"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"mgufrone.dev/job-tracking/tests"
)

type CreateSuite struct {
	tests.HttpTestCase
	auth *mocks2.Auth
}

func TestCreateSuite(t *testing.T) {
	suite.Run(t, new(CreateSuite))
}

// SetupTest will run before each test in the suite.
func (s *CreateSuite) SetupTest() {
	s.HttpTestCase.SetupTest()
	if s.auth == nil {
		s.auth = mock2.Auth()
	}
}

// TearDownTest will run after each test in the suite.
func (s *CreateSuite) TearDownTest() {
}

func (s *CreateSuite) TestCreateFail0() {
	post, err := http.Post(s.ApiServer.URL+"/application", "application/json", nil)
	s.Assert().Equal(400, post.StatusCode)
	tee := io.TeeReader(post.Body, os.Stdout)
	io.ReadAll(tee)
	s.Assert().Nil(err)
}

func (s *CreateSuite) TestCreateFail1() {
	body := application.Create{}
	by, _ := json.Marshal(body)
	post, err := http.Post(s.ApiServer.URL+"/application", "application/json", bytes.NewReader(by))
	defer post.Body.Close()
	res, _ := io.ReadAll(post.Body)
	bodyResponse := string(res)
	s.Assert().Equal(400, post.StatusCode)
	s.Assert().Contains(bodyResponse, "job.link")
	s.Assert().Contains(bodyResponse, "job.company")
	s.Assert().Contains(bodyResponse, "job.role")
	s.Assert().Contains(bodyResponse, "resume_name")
	s.Assert().Nil(err)
}
func (s *CreateSuite) TestCreateFail2() {
	body := application.Create{
		ResumeName: "random1",
		Job: job.Create{
			Link: "https://google.dev/lowongan",
		},
	}
	by, _ := json.Marshal(body)
	post, err := http.Post(s.ApiServer.URL+"/application", "application/json", bytes.NewReader(by))
	res, _ := io.ReadAll(post.Body)
	bodyResponse := string(res)
	s.Assert().Equal(400, post.StatusCode)
	s.Assert().Contains(bodyResponse, "job.company")
	s.Assert().Contains(bodyResponse, "job.role")
	s.Assert().Nil(err)
}

func (s *CreateSuite) TestCreateFail3() {
	body := application.Create{
		ResumeName: "random1",
		Job: job.Create{
			Link: "https://google.dev/lowongan",
			Role: "Software Engineer",
		},
	}
	by, _ := json.Marshal(body)
	post, err := http.Post(s.ApiServer.URL+"/application", "application/json", bytes.NewReader(by))
	res, _ := io.ReadAll(post.Body)
	bodyResponse := string(res)
	s.Assert().Equal(400, post.StatusCode)
	s.Assert().Contains(bodyResponse, "job.company")
	s.Assert().Nil(err)
}

func (s *CreateSuite) TestCreateFail00Unauthenticated() {
	body := application.Create{
		ResumeName: "random1",
		Job: job.Create{
			Link:    "https://google.dev/lowongan",
			Role:    "Software Engineer",
			Company: "Random Company",
		},
	}
	by, _ := json.Marshal(body)
	s.auth.On("Parse", mock.Anything, mock.Anything).Return(&auth.Payload{}, nil).Once()
	s.auth.On("User", mock.Anything, mock.Anything).Return(func(context2 http2.Context, usr interface{}) error {
		return errors.New("invalid user")
	}).Once()
	req, _ := http.NewRequestWithContext(context.TODO(), "POST", s.ApiServer.URL+"/application", bytes.NewReader(by))
	req.Header.Set("Content-Type", "application/json")
	post, err := http.DefaultClient.Do(req)
	//res, _ := io.ReadAll(post.Body)
	//bodyResponse := string(res)
	s.Assert().Nil(err)
	s.Assert().Equal(401, post.StatusCode)
}
func (s *CreateSuite) TestCreateFail00Unauthenticated01() {
	body := application.Create{
		ResumeName: "random1",
		Job: job.Create{
			Link:    "https://google.dev/lowongan",
			Role:    "Software Engineer",
			Company: "Random Company",
		},
	}
	by, _ := json.Marshal(body)
	s.auth.On("Parse", mock.Anything, mock.Anything).Return(nil, errors.New("authentication failed")).Once()
	req, _ := http.NewRequestWithContext(context.TODO(), "POST", s.ApiServer.URL+"/application", bytes.NewReader(by))
	req.Header.Set("Content-Type", "application/json")
	post, err := http.DefaultClient.Do(req)
	//res, _ := io.ReadAll(post.Body)
	//bodyResponse := string(res)
	s.Assert().Nil(err)
	s.Assert().Equal(401, post.StatusCode)
}
func (s *CreateSuite) TestCreateFail01Exists() {
	jobApplicationService := mocks.NewJobApplication(s.T())
	jobApplicationService.
		On("CreateUserApplication", mock.Anything, mock.Anything, mock.Anything).
		Return(service.ErrExists).Once()
	app := mock2.App()
	app.On("Make", service.JobApplicationService).Return(jobApplicationService, nil).Once()
	body := application.Create{
		ResumeName: "random1",
		Job: job.Create{
			Link:    "https://google.dev/lowongan",
			Role:    "Software Engineer",
			Company: "Random Company",
		},
	}
	by, _ := json.Marshal(body)
	var usr models.User
	usr.ID = uuid.New()
	usr.ExternalID = uuid.NewString()
	usr.Name = "randomname"
	token := "randomtoken"
	s.auth.On("Parse", mock.Anything, mock.Anything).Return(&auth.Payload{}, nil).Once()
	s.auth.On("User", mock.Anything, mock.Anything).Return(func(context2 http2.Context, usr interface{}) error {
		usr.(*models.User).ID = uuid.New()
		return nil
	}).Once()
	req, _ := http.NewRequestWithContext(context.TODO(), "POST", s.ApiServer.URL+"/application", bytes.NewReader(by))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	//post, err := http.Post(s.ApiServer.URL+"/application", "application/json", bytes.NewReader(by))
	post, err := http.DefaultClient.Do(req)
	s.Assert().Equal(303, post.StatusCode)
	s.Assert().Nil(err)
}
func (s *CreateSuite) TestCreateFail01Exists01() {
	jobApplicationService := mocks.NewJobApplication(s.T())
	jobApplicationService.
		On("CreateUserApplication", mock.Anything, mock.Anything, mock.Anything).
		Return(errors.New("db connection failed")).Once()
	app := mock2.App()
	app.On("Make", service.JobApplicationService).Return(jobApplicationService, nil).Once()
	body := application.Create{
		ResumeName: "random1",
		Job: job.Create{
			Link:    "https://google.dev/lowongan",
			Role:    "Software Engineer",
			Company: "Random Company",
		},
	}
	by, _ := json.Marshal(body)
	var usr models.User
	usr.ID = uuid.New()
	usr.ExternalID = uuid.NewString()
	usr.Name = "randomname"
	token := "randomtoken"
	s.auth.On("Parse", mock.Anything, mock.Anything).Return(&auth.Payload{}, nil).Once()
	s.auth.On("User", mock.Anything, mock.Anything).Return(func(context2 http2.Context, usr interface{}) error {
		usr.(*models.User).ID = uuid.New()
		return nil
	}).Once()
	req, _ := http.NewRequestWithContext(context.TODO(), "POST", s.ApiServer.URL+"/application", bytes.NewReader(by))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	//post, err := http.Post(s.ApiServer.URL+"/application", "application/json", bytes.NewReader(by))
	post, err := http.DefaultClient.Do(req)
	s.Assert().Equal(http.StatusInternalServerError, post.StatusCode)
	s.Assert().Nil(err)
}

func (s *CreateSuite) TestCreateSuccessful() {
	jobApplicationService := mocks.NewJobApplication(s.T())
	newId := uuid.New()
	jobApplicationService.
		On("CreateUserApplication", mock.Anything, mock.Anything, mock.Anything).
		Return(func(ctx context.Context, usr *models.User, appEntity *models.JobApplication) error {
			appEntity.ID = newId
			return nil
		}).Once()
	app := mock2.App()
	app.On("Make", service.JobApplicationService).Return(jobApplicationService, nil).Once()
	body := application.Create{
		ResumeName: "random1",
		Job: job.Create{
			Link:    "https://google.dev/lowongan",
			Role:    "Software Engineer",
			Company: "Random Company",
		},
	}
	by, _ := json.Marshal(body)
	var usr models.User
	usr.ID = uuid.New()
	usr.ExternalID = uuid.NewString()
	usr.Name = "randomname"
	token := "randomtoken"
	s.auth.On("Parse", mock.Anything, mock.Anything).Return(&auth.Payload{}, nil).Once()
	s.auth.On("User", mock.Anything, mock.Anything).Return(func(context2 http2.Context, usr interface{}) error {
		usr.(*models.User).ID = uuid.New()
		return nil
	}).Once()
	req, _ := http.NewRequestWithContext(context.TODO(), "POST", s.ApiServer.URL+"/application", bytes.NewReader(by))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	//post, err := http.Post(s.ApiServer.URL+"/application", "application/json", bytes.NewReader(by))
	post, err := http.DefaultClient.Do(req)
	result := map[string]string{}
	json2.NewDecoder(post.Body).Decode(&result)
	s.Assert().Equal(http.StatusCreated, post.StatusCode)
	s.Assert().Equal(newId.String(), result["application"])
	s.Assert().Nil(err)
}
