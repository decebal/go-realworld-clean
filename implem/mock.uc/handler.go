// Code generated by MockGen. DO NOT EDIT.
// Source: ./uc/HANDLER.go

// Package uc is a generated GoMock package.
package uc

import (
	domain "github.com/err0r500/go-realworld-clean/domain"
	uc "github.com/err0r500/go-realworld-clean/uc"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHandler is a mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// ProfileGet mocks base method
func (m *MockHandler) ProfileGet(userName string) (*domain.Profile, error) {
	ret := m.ctrl.Call(m, "ProfileGet", userName)
	ret0, _ := ret[0].(*domain.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProfileGet indicates an expected call of ProfileGet
func (mr *MockHandlerMockRecorder) ProfileGet(userName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProfileGet", reflect.TypeOf((*MockHandler)(nil).ProfileGet), userName)
}

// ProfileUpdateFollow mocks base method
func (m *MockHandler) ProfileUpdateFollow(loggedInUsername, username string, follow bool) (*domain.User, error) {
	ret := m.ctrl.Call(m, "ProfileUpdateFollow", loggedInUsername, username, follow)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProfileUpdateFollow indicates an expected call of ProfileUpdateFollow
func (mr *MockHandlerMockRecorder) ProfileUpdateFollow(loggedInUsername, username, follow interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProfileUpdateFollow", reflect.TypeOf((*MockHandler)(nil).ProfileUpdateFollow), loggedInUsername, username, follow)
}

// UserCreate mocks base method
func (m *MockHandler) UserCreate(username, email, password string) (*domain.User, string, error) {
	ret := m.ctrl.Call(m, "UserCreate", username, email, password)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UserCreate indicates an expected call of UserCreate
func (mr *MockHandlerMockRecorder) UserCreate(username, email, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserCreate", reflect.TypeOf((*MockHandler)(nil).UserCreate), username, email, password)
}

// UserLogin mocks base method
func (m *MockHandler) UserLogin(email, password string) (*domain.User, string, error) {
	ret := m.ctrl.Call(m, "UserLogin", email, password)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UserLogin indicates an expected call of UserLogin
func (mr *MockHandlerMockRecorder) UserLogin(email, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLogin", reflect.TypeOf((*MockHandler)(nil).UserLogin), email, password)
}

// UserGet mocks base method
func (m *MockHandler) UserGet(userName string) (*domain.User, string, error) {
	ret := m.ctrl.Call(m, "UserGet", userName)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UserGet indicates an expected call of UserGet
func (mr *MockHandlerMockRecorder) UserGet(userName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserGet", reflect.TypeOf((*MockHandler)(nil).UserGet), userName)
}

// UserEdit mocks base method
func (m *MockHandler) UserEdit(userName string, newUser map[uc.UpdatableProperty]*string) (*domain.User, string, error) {
	ret := m.ctrl.Call(m, "UserEdit", userName, newUser)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UserEdit indicates an expected call of UserEdit
func (mr *MockHandlerMockRecorder) UserEdit(userName, newUser interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserEdit", reflect.TypeOf((*MockHandler)(nil).UserEdit), userName, newUser)
}

// ArticlesFeed mocks base method
func (m *MockHandler) ArticlesFeed(username string, limit, offset int) (domain.ArticleCollection, int, error) {
	ret := m.ctrl.Call(m, "ArticlesFeed", username, limit, offset)
	ret0, _ := ret[0].(domain.ArticleCollection)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ArticlesFeed indicates an expected call of ArticlesFeed
func (mr *MockHandlerMockRecorder) ArticlesFeed(username, limit, offset interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArticlesFeed", reflect.TypeOf((*MockHandler)(nil).ArticlesFeed), username, limit, offset)
}

// GetArticles mocks base method
func (m *MockHandler) GetArticles(limit, offset int, filters uc.Filters) (domain.ArticleCollection, int, error) {
	ret := m.ctrl.Call(m, "GetArticles", limit, offset, filters)
	ret0, _ := ret[0].(domain.ArticleCollection)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetArticles indicates an expected call of GetArticles
func (mr *MockHandlerMockRecorder) GetArticles(limit, offset, filters interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticles", reflect.TypeOf((*MockHandler)(nil).GetArticles), limit, offset, filters)
}

// ArticleGet mocks base method
func (m *MockHandler) ArticleGet(slug string) (*domain.Article, error) {
	ret := m.ctrl.Call(m, "ArticleGet", slug)
	ret0, _ := ret[0].(*domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArticleGet indicates an expected call of ArticleGet
func (mr *MockHandlerMockRecorder) ArticleGet(slug interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArticleGet", reflect.TypeOf((*MockHandler)(nil).ArticleGet), slug)
}

// ArticlePost mocks base method
func (m *MockHandler) ArticlePost(username string, article domain.Article) (*domain.Article, error) {
	ret := m.ctrl.Call(m, "ArticlePost", username, article)
	ret0, _ := ret[0].(*domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArticlePost indicates an expected call of ArticlePost
func (mr *MockHandlerMockRecorder) ArticlePost(username, article interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArticlePost", reflect.TypeOf((*MockHandler)(nil).ArticlePost), username, article)
}

// ArticlePut mocks base method
func (m *MockHandler) ArticlePut(username, slug string, article domain.Article) (*domain.Article, error) {
	ret := m.ctrl.Call(m, "ArticlePut", username, slug, article)
	ret0, _ := ret[0].(*domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArticlePut indicates an expected call of ArticlePut
func (mr *MockHandlerMockRecorder) ArticlePut(username, slug, article interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArticlePut", reflect.TypeOf((*MockHandler)(nil).ArticlePut), username, slug, article)
}

// ArticleDelete mocks base method
func (m *MockHandler) ArticleDelete(username, slug string) error {
	ret := m.ctrl.Call(m, "ArticleDelete", username, slug)
	ret0, _ := ret[0].(error)
	return ret0
}

// ArticleDelete indicates an expected call of ArticleDelete
func (mr *MockHandlerMockRecorder) ArticleDelete(username, slug interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArticleDelete", reflect.TypeOf((*MockHandler)(nil).ArticleDelete), username, slug)
}

// CommentsGet mocks base method
func (m *MockHandler) CommentsGet(slug string) ([]domain.Comment, error) {
	ret := m.ctrl.Call(m, "CommentsGet", slug)
	ret0, _ := ret[0].([]domain.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommentsGet indicates an expected call of CommentsGet
func (mr *MockHandlerMockRecorder) CommentsGet(slug interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommentsGet", reflect.TypeOf((*MockHandler)(nil).CommentsGet), slug)
}

// CommentsPost mocks base method
func (m *MockHandler) CommentsPost(username, slug, comment string) (*domain.Comment, error) {
	ret := m.ctrl.Call(m, "CommentsPost", username, slug, comment)
	ret0, _ := ret[0].(*domain.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommentsPost indicates an expected call of CommentsPost
func (mr *MockHandlerMockRecorder) CommentsPost(username, slug, comment interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommentsPost", reflect.TypeOf((*MockHandler)(nil).CommentsPost), username, slug, comment)
}

// CommentsDelete mocks base method
func (m *MockHandler) CommentsDelete(username, slug string, id int) error {
	ret := m.ctrl.Call(m, "CommentsDelete", username, slug, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommentsDelete indicates an expected call of CommentsDelete
func (mr *MockHandlerMockRecorder) CommentsDelete(username, slug, id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommentsDelete", reflect.TypeOf((*MockHandler)(nil).CommentsDelete), username, slug, id)
}

// FavoritesUpdate mocks base method
func (m *MockHandler) FavoritesUpdate(username, slug string, favortie bool) (*domain.Article, error) {
	ret := m.ctrl.Call(m, "FavoritesUpdate", username, slug, favortie)
	ret0, _ := ret[0].(*domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FavoritesUpdate indicates an expected call of FavoritesUpdate
func (mr *MockHandlerMockRecorder) FavoritesUpdate(username, slug, favortie interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FavoritesUpdate", reflect.TypeOf((*MockHandler)(nil).FavoritesUpdate), username, slug, favortie)
}

// Tags mocks base method
func (m *MockHandler) Tags() ([]string, error) {
	ret := m.ctrl.Call(m, "Tags")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tags indicates an expected call of Tags
func (mr *MockHandlerMockRecorder) Tags() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tags", reflect.TypeOf((*MockHandler)(nil).Tags))
}

// MockProfileLogic is a mock of ProfileLogic interface
type MockProfileLogic struct {
	ctrl     *gomock.Controller
	recorder *MockProfileLogicMockRecorder
}

// MockProfileLogicMockRecorder is the mock recorder for MockProfileLogic
type MockProfileLogicMockRecorder struct {
	mock *MockProfileLogic
}

// NewMockProfileLogic creates a new mock instance
func NewMockProfileLogic(ctrl *gomock.Controller) *MockProfileLogic {
	mock := &MockProfileLogic{ctrl: ctrl}
	mock.recorder = &MockProfileLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProfileLogic) EXPECT() *MockProfileLogicMockRecorder {
	return m.recorder
}

// ProfileGet mocks base method
func (m *MockProfileLogic) ProfileGet(userName string) (*domain.Profile, error) {
	ret := m.ctrl.Call(m, "ProfileGet", userName)
	ret0, _ := ret[0].(*domain.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProfileGet indicates an expected call of ProfileGet
func (mr *MockProfileLogicMockRecorder) ProfileGet(userName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProfileGet", reflect.TypeOf((*MockProfileLogic)(nil).ProfileGet), userName)
}

// ProfileUpdateFollow mocks base method
func (m *MockProfileLogic) ProfileUpdateFollow(loggedInUsername, username string, follow bool) (*domain.User, error) {
	ret := m.ctrl.Call(m, "ProfileUpdateFollow", loggedInUsername, username, follow)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProfileUpdateFollow indicates an expected call of ProfileUpdateFollow
func (mr *MockProfileLogicMockRecorder) ProfileUpdateFollow(loggedInUsername, username, follow interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProfileUpdateFollow", reflect.TypeOf((*MockProfileLogic)(nil).ProfileUpdateFollow), loggedInUsername, username, follow)
}

// MockUserLogic is a mock of UserLogic interface
type MockUserLogic struct {
	ctrl     *gomock.Controller
	recorder *MockUserLogicMockRecorder
}

// MockUserLogicMockRecorder is the mock recorder for MockUserLogic
type MockUserLogicMockRecorder struct {
	mock *MockUserLogic
}

// NewMockUserLogic creates a new mock instance
func NewMockUserLogic(ctrl *gomock.Controller) *MockUserLogic {
	mock := &MockUserLogic{ctrl: ctrl}
	mock.recorder = &MockUserLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserLogic) EXPECT() *MockUserLogicMockRecorder {
	return m.recorder
}

// UserCreate mocks base method
func (m *MockUserLogic) UserCreate(username, email, password string) (*domain.User, string, error) {
	ret := m.ctrl.Call(m, "UserCreate", username, email, password)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UserCreate indicates an expected call of UserCreate
func (mr *MockUserLogicMockRecorder) UserCreate(username, email, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserCreate", reflect.TypeOf((*MockUserLogic)(nil).UserCreate), username, email, password)
}

// UserLogin mocks base method
func (m *MockUserLogic) UserLogin(email, password string) (*domain.User, string, error) {
	ret := m.ctrl.Call(m, "UserLogin", email, password)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UserLogin indicates an expected call of UserLogin
func (mr *MockUserLogicMockRecorder) UserLogin(email, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLogin", reflect.TypeOf((*MockUserLogic)(nil).UserLogin), email, password)
}

// UserGet mocks base method
func (m *MockUserLogic) UserGet(userName string) (*domain.User, string, error) {
	ret := m.ctrl.Call(m, "UserGet", userName)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UserGet indicates an expected call of UserGet
func (mr *MockUserLogicMockRecorder) UserGet(userName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserGet", reflect.TypeOf((*MockUserLogic)(nil).UserGet), userName)
}

// UserEdit mocks base method
func (m *MockUserLogic) UserEdit(userName string, newUser map[uc.UpdatableProperty]*string) (*domain.User, string, error) {
	ret := m.ctrl.Call(m, "UserEdit", userName, newUser)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UserEdit indicates an expected call of UserEdit
func (mr *MockUserLogicMockRecorder) UserEdit(userName, newUser interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserEdit", reflect.TypeOf((*MockUserLogic)(nil).UserEdit), userName, newUser)
}

// MockArticlesLogic is a mock of ArticlesLogic interface
type MockArticlesLogic struct {
	ctrl     *gomock.Controller
	recorder *MockArticlesLogicMockRecorder
}

// MockArticlesLogicMockRecorder is the mock recorder for MockArticlesLogic
type MockArticlesLogicMockRecorder struct {
	mock *MockArticlesLogic
}

// NewMockArticlesLogic creates a new mock instance
func NewMockArticlesLogic(ctrl *gomock.Controller) *MockArticlesLogic {
	mock := &MockArticlesLogic{ctrl: ctrl}
	mock.recorder = &MockArticlesLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockArticlesLogic) EXPECT() *MockArticlesLogicMockRecorder {
	return m.recorder
}

// ArticlesFeed mocks base method
func (m *MockArticlesLogic) ArticlesFeed(username string, limit, offset int) (domain.ArticleCollection, int, error) {
	ret := m.ctrl.Call(m, "ArticlesFeed", username, limit, offset)
	ret0, _ := ret[0].(domain.ArticleCollection)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ArticlesFeed indicates an expected call of ArticlesFeed
func (mr *MockArticlesLogicMockRecorder) ArticlesFeed(username, limit, offset interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArticlesFeed", reflect.TypeOf((*MockArticlesLogic)(nil).ArticlesFeed), username, limit, offset)
}

// GetArticles mocks base method
func (m *MockArticlesLogic) GetArticles(limit, offset int, filters uc.Filters) (domain.ArticleCollection, int, error) {
	ret := m.ctrl.Call(m, "GetArticles", limit, offset, filters)
	ret0, _ := ret[0].(domain.ArticleCollection)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetArticles indicates an expected call of GetArticles
func (mr *MockArticlesLogicMockRecorder) GetArticles(limit, offset, filters interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticles", reflect.TypeOf((*MockArticlesLogic)(nil).GetArticles), limit, offset, filters)
}

// MockArticleLogic is a mock of ArticleLogic interface
type MockArticleLogic struct {
	ctrl     *gomock.Controller
	recorder *MockArticleLogicMockRecorder
}

// MockArticleLogicMockRecorder is the mock recorder for MockArticleLogic
type MockArticleLogicMockRecorder struct {
	mock *MockArticleLogic
}

// NewMockArticleLogic creates a new mock instance
func NewMockArticleLogic(ctrl *gomock.Controller) *MockArticleLogic {
	mock := &MockArticleLogic{ctrl: ctrl}
	mock.recorder = &MockArticleLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockArticleLogic) EXPECT() *MockArticleLogicMockRecorder {
	return m.recorder
}

// ArticleGet mocks base method
func (m *MockArticleLogic) ArticleGet(slug string) (*domain.Article, error) {
	ret := m.ctrl.Call(m, "ArticleGet", slug)
	ret0, _ := ret[0].(*domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArticleGet indicates an expected call of ArticleGet
func (mr *MockArticleLogicMockRecorder) ArticleGet(slug interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArticleGet", reflect.TypeOf((*MockArticleLogic)(nil).ArticleGet), slug)
}

// ArticlePost mocks base method
func (m *MockArticleLogic) ArticlePost(username string, article domain.Article) (*domain.Article, error) {
	ret := m.ctrl.Call(m, "ArticlePost", username, article)
	ret0, _ := ret[0].(*domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArticlePost indicates an expected call of ArticlePost
func (mr *MockArticleLogicMockRecorder) ArticlePost(username, article interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArticlePost", reflect.TypeOf((*MockArticleLogic)(nil).ArticlePost), username, article)
}

// ArticlePut mocks base method
func (m *MockArticleLogic) ArticlePut(username, slug string, article domain.Article) (*domain.Article, error) {
	ret := m.ctrl.Call(m, "ArticlePut", username, slug, article)
	ret0, _ := ret[0].(*domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArticlePut indicates an expected call of ArticlePut
func (mr *MockArticleLogicMockRecorder) ArticlePut(username, slug, article interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArticlePut", reflect.TypeOf((*MockArticleLogic)(nil).ArticlePut), username, slug, article)
}

// ArticleDelete mocks base method
func (m *MockArticleLogic) ArticleDelete(username, slug string) error {
	ret := m.ctrl.Call(m, "ArticleDelete", username, slug)
	ret0, _ := ret[0].(error)
	return ret0
}

// ArticleDelete indicates an expected call of ArticleDelete
func (mr *MockArticleLogicMockRecorder) ArticleDelete(username, slug interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArticleDelete", reflect.TypeOf((*MockArticleLogic)(nil).ArticleDelete), username, slug)
}

// MockCommentsLogic is a mock of CommentsLogic interface
type MockCommentsLogic struct {
	ctrl     *gomock.Controller
	recorder *MockCommentsLogicMockRecorder
}

// MockCommentsLogicMockRecorder is the mock recorder for MockCommentsLogic
type MockCommentsLogicMockRecorder struct {
	mock *MockCommentsLogic
}

// NewMockCommentsLogic creates a new mock instance
func NewMockCommentsLogic(ctrl *gomock.Controller) *MockCommentsLogic {
	mock := &MockCommentsLogic{ctrl: ctrl}
	mock.recorder = &MockCommentsLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommentsLogic) EXPECT() *MockCommentsLogicMockRecorder {
	return m.recorder
}

// CommentsGet mocks base method
func (m *MockCommentsLogic) CommentsGet(slug string) ([]domain.Comment, error) {
	ret := m.ctrl.Call(m, "CommentsGet", slug)
	ret0, _ := ret[0].([]domain.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommentsGet indicates an expected call of CommentsGet
func (mr *MockCommentsLogicMockRecorder) CommentsGet(slug interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommentsGet", reflect.TypeOf((*MockCommentsLogic)(nil).CommentsGet), slug)
}

// CommentsPost mocks base method
func (m *MockCommentsLogic) CommentsPost(username, slug, comment string) (*domain.Comment, error) {
	ret := m.ctrl.Call(m, "CommentsPost", username, slug, comment)
	ret0, _ := ret[0].(*domain.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommentsPost indicates an expected call of CommentsPost
func (mr *MockCommentsLogicMockRecorder) CommentsPost(username, slug, comment interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommentsPost", reflect.TypeOf((*MockCommentsLogic)(nil).CommentsPost), username, slug, comment)
}

// CommentsDelete mocks base method
func (m *MockCommentsLogic) CommentsDelete(username, slug string, id int) error {
	ret := m.ctrl.Call(m, "CommentsDelete", username, slug, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommentsDelete indicates an expected call of CommentsDelete
func (mr *MockCommentsLogicMockRecorder) CommentsDelete(username, slug, id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommentsDelete", reflect.TypeOf((*MockCommentsLogic)(nil).CommentsDelete), username, slug, id)
}

// MockFavoritesLogic is a mock of FavoritesLogic interface
type MockFavoritesLogic struct {
	ctrl     *gomock.Controller
	recorder *MockFavoritesLogicMockRecorder
}

// MockFavoritesLogicMockRecorder is the mock recorder for MockFavoritesLogic
type MockFavoritesLogicMockRecorder struct {
	mock *MockFavoritesLogic
}

// NewMockFavoritesLogic creates a new mock instance
func NewMockFavoritesLogic(ctrl *gomock.Controller) *MockFavoritesLogic {
	mock := &MockFavoritesLogic{ctrl: ctrl}
	mock.recorder = &MockFavoritesLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFavoritesLogic) EXPECT() *MockFavoritesLogicMockRecorder {
	return m.recorder
}

// FavoritesUpdate mocks base method
func (m *MockFavoritesLogic) FavoritesUpdate(username, slug string, favortie bool) (*domain.Article, error) {
	ret := m.ctrl.Call(m, "FavoritesUpdate", username, slug, favortie)
	ret0, _ := ret[0].(*domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FavoritesUpdate indicates an expected call of FavoritesUpdate
func (mr *MockFavoritesLogicMockRecorder) FavoritesUpdate(username, slug, favortie interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FavoritesUpdate", reflect.TypeOf((*MockFavoritesLogic)(nil).FavoritesUpdate), username, slug, favortie)
}

// MockTagsLogic is a mock of TagsLogic interface
type MockTagsLogic struct {
	ctrl     *gomock.Controller
	recorder *MockTagsLogicMockRecorder
}

// MockTagsLogicMockRecorder is the mock recorder for MockTagsLogic
type MockTagsLogicMockRecorder struct {
	mock *MockTagsLogic
}

// NewMockTagsLogic creates a new mock instance
func NewMockTagsLogic(ctrl *gomock.Controller) *MockTagsLogic {
	mock := &MockTagsLogic{ctrl: ctrl}
	mock.recorder = &MockTagsLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTagsLogic) EXPECT() *MockTagsLogicMockRecorder {
	return m.recorder
}

// Tags mocks base method
func (m *MockTagsLogic) Tags() ([]string, error) {
	ret := m.ctrl.Call(m, "Tags")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tags indicates an expected call of Tags
func (mr *MockTagsLogicMockRecorder) Tags() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tags", reflect.TypeOf((*MockTagsLogic)(nil).Tags))
}
