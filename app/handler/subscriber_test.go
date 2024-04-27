package handler

import (
	"energy-response-assignment/app/service"
	"energy-response-assignment/config"
	"energy-response-assignment/entity"
	mockService "energy-response-assignment/mocks/app/service"
	"energy-response-assignment/util/validator"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNewSubscriber(t *testing.T) {
	type args struct {
		cfg               *config.Config
		subscriberService service.Subscriber
	}
	tests := []struct {
		name string
		args args
		want Subscriber
	}{
		{
			name: "success",
			args: args{
				cfg:               &config.Config{},
				subscriberService: mockService.NewSubscriber(t),
			},
			want: &subscriberImpl{
				Config:            &config.Config{},
				SubscriberService: mockService.NewSubscriber(t),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSubscriber(tt.args.cfg, tt.args.subscriberService)
			assert.Equalf(t, tt.want, got, "NewSubscriber() = %v, want %v", got, tt.want)
		})
	}
}

func Test_subscriberImpl_Subscribe(t *testing.T) {
	type fields struct {
		Config            *config.Config
		SubscriberService service.Subscriber
	}
	type args struct {
		e       *echo.Echo
		reqJson string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		rec        *httptest.ResponseRecorder
		resJson    string
		mockFn     func(args args, fields fields, returner *entity.Subscriber)
		statusCode int
		wantErr    assert.ErrorAssertionFunc
		subscriber *entity.Subscriber
	}{
		{
			name: "Subscribe Test",
			fields: fields{
				Config:            &config.Config{},
				SubscriberService: mockService.NewSubscriber(t),
			},
			args: args{
				e:       echo.New(),
				reqJson: `{"email":"test@example.com"}`,
			},
			rec: httptest.NewRecorder(),
			subscriber: &entity.Subscriber{
				ID:        uuid.New(),
				Email:     "test@example.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Active:    true,
			},
			resJson: `{"email":"test@example.com","active":true,"message":"Subscribe success"}`,
			mockFn: func(args args, fields fields, returner *entity.Subscriber) {
				args.e.Validator = validator.GetEchoValidator()
				fields.SubscriberService.(*mockService.Subscriber).EXPECT().Subscribe(mock.Anything, returner.Email).Return(returner, nil)
			},
			statusCode: http.StatusOK,
			wantErr:    assert.NoError,
		},
		{
			name: "Invalid Request",
			fields: fields{
				Config:            &config.Config{},
				SubscriberService: mockService.NewSubscriber(t),
			},
			args: args{
				e:       echo.New(),
				reqJson: `{"wrongKey":"test@example.com"}`,
			},
			rec:        httptest.NewRecorder(),
			subscriber: nil,
			resJson:    `{"email":"test@example.com","active":true,"message":"Subscribe success"}`,
			mockFn: func(args args, fields fields, returner *entity.Subscriber) {
				args.e.Validator = validator.GetEchoValidator()
			},
			statusCode: http.StatusBadRequest,
			wantErr:    assert.Error,
		},
		{
			name: "Validate Request",
			fields: fields{
				Config:            &config.Config{},
				SubscriberService: mockService.NewSubscriber(t),
			},
			args: args{
				e:       echo.New(),
				reqJson: `{"email":"testexample.com"}`,
			},
			rec:        httptest.NewRecorder(),
			subscriber: nil,
			resJson:    `{"email":"test@example.com","active":true,"message":"Subscribe success"}`,
			mockFn: func(args args, fields fields, returner *entity.Subscriber) {
				args.e.Validator = validator.GetEchoValidator()
			},
			statusCode: http.StatusBadRequest,
			wantErr:    assert.Error,
		},
		{
			name: "Internal Error",
			fields: fields{
				Config:            &config.Config{},
				SubscriberService: mockService.NewSubscriber(t),
			},
			args: args{
				e:       echo.New(),
				reqJson: `{"email":"test@example.com"}`,
			},
			rec: httptest.NewRecorder(),
			subscriber: &entity.Subscriber{
				ID:        uuid.New(),
				Email:     "test@example.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Active:    true,
			},
			resJson: `{"email":"test@example.com","active":true,"message":"Subscribe success"}`,
			mockFn: func(args args, fields fields, returner *entity.Subscriber) {
				args.e.Validator = validator.GetEchoValidator()
				fields.SubscriberService.(*mockService.Subscriber).EXPECT().Subscribe(mock.Anything, returner.Email).Return(nil, fmt.Errorf("error"))
			},
			statusCode: http.StatusInternalServerError,
			wantErr:    assert.Error,
		},
		{
			name: "Invalid Request format",
			fields: fields{
				Config:            &config.Config{},
				SubscriberService: mockService.NewSubscriber(t),
			},
			args: args{
				e:       echo.New(),
				reqJson: `{"wrongKey"`,
			},
			rec:        httptest.NewRecorder(),
			subscriber: nil,
			resJson:    `{"email":"test@example.com","active":true,"message":"Subscribe success"}`,
			mockFn: func(args args, fields fields, returner *entity.Subscriber) {
				args.e.Validator = validator.GetEchoValidator()
			},
			statusCode: http.StatusBadRequest,
			wantErr:    assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &subscriberImpl{
				Config:            tt.fields.Config,
				SubscriberService: tt.fields.SubscriberService,
			}
			tt.mockFn(tt.args, tt.fields, tt.subscriber)
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.args.reqJson))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			echoContext := tt.args.e.NewContext(req, tt.rec)
			err := s.Subscribe(echoContext)
			if tt.wantErr(t, err, fmt.Sprintf("Subscribe(%v)", echoContext)) {
				if err == nil {
					assert.Equal(t, tt.statusCode, tt.rec.Code)
					assert.JSONEq(t, tt.resJson, tt.rec.Body.String())
				} else {
					assert.IsType(t, &echo.HTTPError{}, err)
					assert.Equal(t, tt.statusCode, err.(*echo.HTTPError).Code)
					assert.NotEmpty(t, tt.resJson)
				}
			}
		})
	}
}

func Test_subscriberImpl_UnSubscribe(t *testing.T) {
	type fields struct {
		Config            *config.Config
		SubscriberService service.Subscriber
	}
	type args struct {
		e       *echo.Echo
		reqJson string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		rec        *httptest.ResponseRecorder
		resJson    string
		mockFn     func(args args, fields fields, returner *entity.Subscriber)
		statusCode int
		wantErr    assert.ErrorAssertionFunc
		subscriber *entity.Subscriber
	}{
		{
			name: "UnSubscribe Test",
			fields: fields{
				Config:            &config.Config{},
				SubscriberService: mockService.NewSubscriber(t),
			},
			args: args{
				e:       echo.New(),
				reqJson: `{"email":"test@example.com"}`,
			},
			rec: httptest.NewRecorder(),
			subscriber: &entity.Subscriber{
				ID:        uuid.New(),
				Email:     "test@example.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Active:    false,
			},
			resJson: `{"email":"test@example.com","active":false,"message":"UnSubscribe success"}`,
			mockFn: func(args args, fields fields, returner *entity.Subscriber) {
				args.e.Validator = validator.GetEchoValidator()
				fields.SubscriberService.(*mockService.Subscriber).EXPECT().Unsubscribe(mock.Anything, returner.Email).Return(returner, nil)
			},
			statusCode: http.StatusOK,
			wantErr:    assert.NoError,
		},
		{
			name: "Invalid Request",
			fields: fields{
				Config:            &config.Config{},
				SubscriberService: mockService.NewSubscriber(t),
			},
			args: args{
				e:       echo.New(),
				reqJson: `{"wrongKey":"test@example.com"}`,
			},
			rec:        httptest.NewRecorder(),
			subscriber: nil,
			resJson:    `{"email":"test@example.com","active":true,"message":"Subscribe success"}`,
			mockFn: func(args args, fields fields, returner *entity.Subscriber) {
				args.e.Validator = validator.GetEchoValidator()
			},
			statusCode: http.StatusBadRequest,
			wantErr:    assert.Error,
		},
		{
			name: "Validate Request",
			fields: fields{
				Config:            &config.Config{},
				SubscriberService: mockService.NewSubscriber(t),
			},
			args: args{
				e:       echo.New(),
				reqJson: `{"email":"testexample.com"}`,
			},
			rec:        httptest.NewRecorder(),
			subscriber: nil,
			resJson:    `{"email":"test@example.com","active":true,"message":"Subscribe success"}`,
			mockFn: func(args args, fields fields, returner *entity.Subscriber) {
				args.e.Validator = validator.GetEchoValidator()
			},
			statusCode: http.StatusBadRequest,
			wantErr:    assert.Error,
		},
		{
			name: "Internal Error",
			fields: fields{
				Config:            &config.Config{},
				SubscriberService: mockService.NewSubscriber(t),
			},
			args: args{
				e:       echo.New(),
				reqJson: `{"email":"test@example.com"}`,
			},
			rec: httptest.NewRecorder(),
			subscriber: &entity.Subscriber{
				ID:        uuid.New(),
				Email:     "test@example.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Active:    true,
			},
			resJson: `{"email":"test@example.com","active":true,"message":"Subscribe success"}`,
			mockFn: func(args args, fields fields, returner *entity.Subscriber) {
				args.e.Validator = validator.GetEchoValidator()
				fields.SubscriberService.(*mockService.Subscriber).EXPECT().Unsubscribe(mock.Anything, returner.Email).Return(nil, fmt.Errorf("error"))
			},
			statusCode: http.StatusInternalServerError,
			wantErr:    assert.Error,
		},
		{
			name: "Invalid Request Format",
			fields: fields{
				Config:            &config.Config{},
				SubscriberService: mockService.NewSubscriber(t),
			},
			args: args{
				e:       echo.New(),
				reqJson: `{"wrongKey":"te`,
			},
			rec:        httptest.NewRecorder(),
			subscriber: nil,
			resJson:    `{"email":"test@example.com","active":true,"message":"Subscribe success"}`,
			mockFn: func(args args, fields fields, returner *entity.Subscriber) {
				args.e.Validator = validator.GetEchoValidator()
			},
			statusCode: http.StatusBadRequest,
			wantErr:    assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &subscriberImpl{
				Config:            tt.fields.Config,
				SubscriberService: tt.fields.SubscriberService,
			}
			tt.mockFn(tt.args, tt.fields, tt.subscriber)
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.args.reqJson))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			echoContext := tt.args.e.NewContext(req, tt.rec)
			err := s.UnSubscribe(echoContext)
			if tt.wantErr(t, err, fmt.Sprintf("UnSubscribe(%v)", echoContext)) {
				if err == nil {
					assert.Equal(t, tt.statusCode, tt.rec.Code)
					assert.JSONEq(t, tt.resJson, tt.rec.Body.String())
				} else {
					assert.IsType(t, &echo.HTTPError{}, err)
					assert.Equal(t, tt.statusCode, err.(*echo.HTTPError).Code)
					assert.NotEmpty(t, tt.resJson)
				}
			}
		})
	}
}
