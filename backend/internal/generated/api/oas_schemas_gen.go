// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"

	ht "github.com/ogen-go/ogen/http"
)

// DownloadGetInternalServerError is response for DownloadGet operation.
type DownloadGetInternalServerError struct{}

func (*DownloadGetInternalServerError) downloadGetRes() {}

// DownloadGetNotFound is response for DownloadGet operation.
type DownloadGetNotFound struct{}

func (*DownloadGetNotFound) downloadGetRes() {}

type DownloadGetOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s DownloadGetOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

// DownloadGetOKHeaders wraps DownloadGetOK with response headers.
type DownloadGetOKHeaders struct {
	ContentLength OptInt64
	Response      DownloadGetOK
}

// GetContentLength returns the value of ContentLength.
func (s *DownloadGetOKHeaders) GetContentLength() OptInt64 {
	return s.ContentLength
}

// GetResponse returns the value of Response.
func (s *DownloadGetOKHeaders) GetResponse() DownloadGetOK {
	return s.Response
}

// SetContentLength sets the value of ContentLength.
func (s *DownloadGetOKHeaders) SetContentLength(val OptInt64) {
	s.ContentLength = val
}

// SetResponse sets the value of Response.
func (s *DownloadGetOKHeaders) SetResponse(val DownloadGetOK) {
	s.Response = val
}

func (*DownloadGetOKHeaders) downloadGetRes() {}

type HealthGetOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s HealthGetOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

// ListGetInternalServerError is response for ListGet operation.
type ListGetInternalServerError struct{}

func (*ListGetInternalServerError) listGetRes() {}

type ListGetOKApplicationJSON []ModInfo

func (*ListGetOKApplicationJSON) listGetRes() {}

// Ref: #/components/schemas/ModInfo
type ModInfo struct {
	// Название файла мода.
	Name OptString `json:"name"`
	// Признак, активен ли мод.
	Active OptBool `json:"active"`
}

// GetName returns the value of Name.
func (s *ModInfo) GetName() OptString {
	return s.Name
}

// GetActive returns the value of Active.
func (s *ModInfo) GetActive() OptBool {
	return s.Active
}

// SetName sets the value of Name.
func (s *ModInfo) SetName(val OptString) {
	s.Name = val
}

// SetActive sets the value of Active.
func (s *ModInfo) SetActive(val OptBool) {
	s.Active = val
}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt64 returns new OptInt64 with value set to v.
func NewOptInt64(v int64) OptInt64 {
	return OptInt64{
		Value: v,
		Set:   true,
	}
}

// OptInt64 is optional int64.
type OptInt64 struct {
	Value int64
	Set   bool
}

// IsSet returns true if OptInt64 was set.
func (o OptInt64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt64) Reset() {
	var v int64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt64) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt64) Get() (v int64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt64) Or(d int64) int64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/UploadConflictResponse
type UploadConflictResponse struct {
	Message OptString `json:"message"`
	// Файлы, успешно загруженные на сервер.
	Uploaded []string `json:"uploaded"`
	// Файлы, которые уже существуют в папке mods.
	Conflicts []string `json:"conflicts"`
	// Файлы с недопустимым расширением или содержимым.
	Invalid []string `json:"invalid"`
}

// GetMessage returns the value of Message.
func (s *UploadConflictResponse) GetMessage() OptString {
	return s.Message
}

// GetUploaded returns the value of Uploaded.
func (s *UploadConflictResponse) GetUploaded() []string {
	return s.Uploaded
}

// GetConflicts returns the value of Conflicts.
func (s *UploadConflictResponse) GetConflicts() []string {
	return s.Conflicts
}

// GetInvalid returns the value of Invalid.
func (s *UploadConflictResponse) GetInvalid() []string {
	return s.Invalid
}

// SetMessage sets the value of Message.
func (s *UploadConflictResponse) SetMessage(val OptString) {
	s.Message = val
}

// SetUploaded sets the value of Uploaded.
func (s *UploadConflictResponse) SetUploaded(val []string) {
	s.Uploaded = val
}

// SetConflicts sets the value of Conflicts.
func (s *UploadConflictResponse) SetConflicts(val []string) {
	s.Conflicts = val
}

// SetInvalid sets the value of Invalid.
func (s *UploadConflictResponse) SetInvalid(val []string) {
	s.Invalid = val
}

// UploadPostBadRequest is response for UploadPost operation.
type UploadPostBadRequest struct{}

func (*UploadPostBadRequest) uploadPostRes() {}

type UploadPostConflict UploadConflictResponse

func (*UploadPostConflict) uploadPostRes() {}

// UploadPostInternalServerError is response for UploadPost operation.
type UploadPostInternalServerError struct{}

func (*UploadPostInternalServerError) uploadPostRes() {}

type UploadPostMultiStatus UploadConflictResponse

func (*UploadPostMultiStatus) uploadPostRes() {}

type UploadPostReq struct {
	// Один или несколько `.jar` файлов.
	Files []ht.MultipartFile `json:"files"`
}

// GetFiles returns the value of Files.
func (s *UploadPostReq) GetFiles() []ht.MultipartFile {
	return s.Files
}

// SetFiles sets the value of Files.
func (s *UploadPostReq) SetFiles(val []ht.MultipartFile) {
	s.Files = val
}

// Ref: #/components/schemas/UploadSuccessResponse
type UploadSuccessResponse struct {
	Message OptString `json:"message"`
	// Файлы, успешно загруженные на сервер.
	Uploaded []string `json:"uploaded"`
}

// GetMessage returns the value of Message.
func (s *UploadSuccessResponse) GetMessage() OptString {
	return s.Message
}

// GetUploaded returns the value of Uploaded.
func (s *UploadSuccessResponse) GetUploaded() []string {
	return s.Uploaded
}

// SetMessage sets the value of Message.
func (s *UploadSuccessResponse) SetMessage(val OptString) {
	s.Message = val
}

// SetUploaded sets the value of Uploaded.
func (s *UploadSuccessResponse) SetUploaded(val []string) {
	s.Uploaded = val
}

func (*UploadSuccessResponse) uploadPostRes() {}
