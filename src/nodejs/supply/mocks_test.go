// Code generated by MockGen. DO NOT EDIT.
// Source: supply.go

// Package supply_test is a generated GoMock package.
package supply_test

import (
	libbuildpack "github.com/cloudfoundry/libbuildpack"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockCommand is a mock of Command interface
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCommandMockRecorder
}

// MockCommandMockRecorder is the mock recorder for MockCommand
type MockCommandMockRecorder struct {
	mock *MockCommand
}

// NewMockCommand creates a new mock instance
func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &MockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommand) EXPECT() *MockCommandMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockCommand) Execute(arg0 string, arg1, arg2 io.Writer, arg3 string, arg4 ...string) error {
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Execute", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockCommandMockRecorder) Execute(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCommand)(nil).Execute), varargs...)
}

// MockManifest is a mock of Manifest interface
type MockManifest struct {
	ctrl     *gomock.Controller
	recorder *MockManifestMockRecorder
}

// MockManifestMockRecorder is the mock recorder for MockManifest
type MockManifestMockRecorder struct {
	mock *MockManifest
}

// NewMockManifest creates a new mock instance
func NewMockManifest(ctrl *gomock.Controller) *MockManifest {
	mock := &MockManifest{ctrl: ctrl}
	mock.recorder = &MockManifestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManifest) EXPECT() *MockManifestMockRecorder {
	return m.recorder
}

// AllDependencyVersions mocks base method
func (m *MockManifest) AllDependencyVersions(arg0 string) []string {
	ret := m.ctrl.Call(m, "AllDependencyVersions", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// AllDependencyVersions indicates an expected call of AllDependencyVersions
func (mr *MockManifestMockRecorder) AllDependencyVersions(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllDependencyVersions", reflect.TypeOf((*MockManifest)(nil).AllDependencyVersions), arg0)
}

// DefaultVersion mocks base method
func (m *MockManifest) DefaultVersion(arg0 string) (libbuildpack.Dependency, error) {
	ret := m.ctrl.Call(m, "DefaultVersion", arg0)
	ret0, _ := ret[0].(libbuildpack.Dependency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DefaultVersion indicates an expected call of DefaultVersion
func (mr *MockManifestMockRecorder) DefaultVersion(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultVersion", reflect.TypeOf((*MockManifest)(nil).DefaultVersion), arg0)
}

// InstallDependency mocks base method
func (m *MockManifest) InstallDependency(arg0 libbuildpack.Dependency, arg1 string) error {
	ret := m.ctrl.Call(m, "InstallDependency", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallDependency indicates an expected call of InstallDependency
func (mr *MockManifestMockRecorder) InstallDependency(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallDependency", reflect.TypeOf((*MockManifest)(nil).InstallDependency), arg0, arg1)
}

// InstallOnlyVersion mocks base method
func (m *MockManifest) InstallOnlyVersion(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "InstallOnlyVersion", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallOnlyVersion indicates an expected call of InstallOnlyVersion
func (mr *MockManifestMockRecorder) InstallOnlyVersion(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallOnlyVersion", reflect.TypeOf((*MockManifest)(nil).InstallOnlyVersion), arg0, arg1)
}

// MockNPM is a mock of NPM interface
type MockNPM struct {
	ctrl     *gomock.Controller
	recorder *MockNPMMockRecorder
}

// MockNPMMockRecorder is the mock recorder for MockNPM
type MockNPMMockRecorder struct {
	mock *MockNPM
}

// NewMockNPM creates a new mock instance
func NewMockNPM(ctrl *gomock.Controller) *MockNPM {
	mock := &MockNPM{ctrl: ctrl}
	mock.recorder = &MockNPMMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNPM) EXPECT() *MockNPMMockRecorder {
	return m.recorder
}

// Build mocks base method
func (m *MockNPM) Build(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "Build", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Build indicates an expected call of Build
func (mr *MockNPMMockRecorder) Build(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockNPM)(nil).Build), arg0, arg1)
}

// Rebuild mocks base method
func (m *MockNPM) Rebuild(arg0 string) error {
	ret := m.ctrl.Call(m, "Rebuild", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rebuild indicates an expected call of Rebuild
func (mr *MockNPMMockRecorder) Rebuild(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rebuild", reflect.TypeOf((*MockNPM)(nil).Rebuild), arg0)
}

// MockYarn is a mock of Yarn interface
type MockYarn struct {
	ctrl     *gomock.Controller
	recorder *MockYarnMockRecorder
}

// MockYarnMockRecorder is the mock recorder for MockYarn
type MockYarnMockRecorder struct {
	mock *MockYarn
}

// NewMockYarn creates a new mock instance
func NewMockYarn(ctrl *gomock.Controller) *MockYarn {
	mock := &MockYarn{ctrl: ctrl}
	mock.recorder = &MockYarnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockYarn) EXPECT() *MockYarnMockRecorder {
	return m.recorder
}

// Build mocks base method
func (m *MockYarn) Build(arg0, arg1, arg2 string) error {
	ret := m.ctrl.Call(m, "Build", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Build indicates an expected call of Build
func (mr *MockYarnMockRecorder) Build(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockYarn)(nil).Build), arg0, arg1, arg2)
}

// MockStager is a mock of Stager interface
type MockStager struct {
	ctrl     *gomock.Controller
	recorder *MockStagerMockRecorder
}

// MockStagerMockRecorder is the mock recorder for MockStager
type MockStagerMockRecorder struct {
	mock *MockStager
}

// NewMockStager creates a new mock instance
func NewMockStager(ctrl *gomock.Controller) *MockStager {
	mock := &MockStager{ctrl: ctrl}
	mock.recorder = &MockStagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStager) EXPECT() *MockStagerMockRecorder {
	return m.recorder
}

// BuildDir mocks base method
func (m *MockStager) BuildDir() string {
	ret := m.ctrl.Call(m, "BuildDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildDir indicates an expected call of BuildDir
func (mr *MockStagerMockRecorder) BuildDir() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildDir", reflect.TypeOf((*MockStager)(nil).BuildDir))
}

// CacheDir mocks base method
func (m *MockStager) CacheDir() string {
	ret := m.ctrl.Call(m, "CacheDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// CacheDir indicates an expected call of CacheDir
func (mr *MockStagerMockRecorder) CacheDir() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheDir", reflect.TypeOf((*MockStager)(nil).CacheDir))
}

// DepDir mocks base method
func (m *MockStager) DepDir() string {
	ret := m.ctrl.Call(m, "DepDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// DepDir indicates an expected call of DepDir
func (mr *MockStagerMockRecorder) DepDir() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepDir", reflect.TypeOf((*MockStager)(nil).DepDir))
}

// DepsIdx mocks base method
func (m *MockStager) DepsIdx() string {
	ret := m.ctrl.Call(m, "DepsIdx")
	ret0, _ := ret[0].(string)
	return ret0
}

// DepsIdx indicates an expected call of DepsIdx
func (mr *MockStagerMockRecorder) DepsIdx() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepsIdx", reflect.TypeOf((*MockStager)(nil).DepsIdx))
}

// LinkDirectoryInDepDir mocks base method
func (m *MockStager) LinkDirectoryInDepDir(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "LinkDirectoryInDepDir", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkDirectoryInDepDir indicates an expected call of LinkDirectoryInDepDir
func (mr *MockStagerMockRecorder) LinkDirectoryInDepDir(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkDirectoryInDepDir", reflect.TypeOf((*MockStager)(nil).LinkDirectoryInDepDir), arg0, arg1)
}

// WriteEnvFile mocks base method
func (m *MockStager) WriteEnvFile(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "WriteEnvFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteEnvFile indicates an expected call of WriteEnvFile
func (mr *MockStagerMockRecorder) WriteEnvFile(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteEnvFile", reflect.TypeOf((*MockStager)(nil).WriteEnvFile), arg0, arg1)
}

// WriteProfileD mocks base method
func (m *MockStager) WriteProfileD(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "WriteProfileD", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteProfileD indicates an expected call of WriteProfileD
func (mr *MockStagerMockRecorder) WriteProfileD(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteProfileD", reflect.TypeOf((*MockStager)(nil).WriteProfileD), arg0, arg1)
}

// SetStagingEnvironment mocks base method
func (m *MockStager) SetStagingEnvironment() error {
	ret := m.ctrl.Call(m, "SetStagingEnvironment")
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStagingEnvironment indicates an expected call of SetStagingEnvironment
func (mr *MockStagerMockRecorder) SetStagingEnvironment() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStagingEnvironment", reflect.TypeOf((*MockStager)(nil).SetStagingEnvironment))
}
