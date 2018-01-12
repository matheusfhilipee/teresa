// Package storage provides access to the Cloud Storage JSON API.
//
// See https://developers.google.com/storage/docs/json_api/
//
// Usage example:
//
//   import "google.golang.org/api/storage/v1"
//   ...
//   storageService, err := storage.New(oauthHttpClient)
package storage

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "storage:v1"
const apiName = "storage"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/storage/v1/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View your data across Google Cloud Platform services
	CloudPlatformReadOnlyScope = "https://www.googleapis.com/auth/cloud-platform.read-only"

	// Manage your data and permissions in Google Cloud Storage
	DevstorageFullControlScope = "https://www.googleapis.com/auth/devstorage.full_control"

	// View your data in Google Cloud Storage
	DevstorageReadOnlyScope = "https://www.googleapis.com/auth/devstorage.read_only"

	// Manage your data in Google Cloud Storage
	DevstorageReadWriteScope = "https://www.googleapis.com/auth/devstorage.read_write"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.BucketAccessControls = NewBucketAccessControlsService(s)
	s.Buckets = NewBucketsService(s)
	s.Channels = NewChannelsService(s)
	s.DefaultObjectAccessControls = NewDefaultObjectAccessControlsService(s)
	s.Notifications = NewNotificationsService(s)
	s.ObjectAccessControls = NewObjectAccessControlsService(s)
	s.Objects = NewObjectsService(s)
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	BucketAccessControls *BucketAccessControlsService

	Buckets *BucketsService

	Channels *ChannelsService

	DefaultObjectAccessControls *DefaultObjectAccessControlsService

	Notifications *NotificationsService

	ObjectAccessControls *ObjectAccessControlsService

	Objects *ObjectsService

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewBucketAccessControlsService(s *Service) *BucketAccessControlsService {
	rs := &BucketAccessControlsService{s: s}
	return rs
}

type BucketAccessControlsService struct {
	s *Service
}

func NewBucketsService(s *Service) *BucketsService {
	rs := &BucketsService{s: s}
	return rs
}

type BucketsService struct {
	s *Service
}

func NewChannelsService(s *Service) *ChannelsService {
	rs := &ChannelsService{s: s}
	return rs
}

type ChannelsService struct {
	s *Service
}

func NewDefaultObjectAccessControlsService(s *Service) *DefaultObjectAccessControlsService {
	rs := &DefaultObjectAccessControlsService{s: s}
	return rs
}

type DefaultObjectAccessControlsService struct {
	s *Service
}

func NewNotificationsService(s *Service) *NotificationsService {
	rs := &NotificationsService{s: s}
	return rs
}

type NotificationsService struct {
	s *Service
}

func NewObjectAccessControlsService(s *Service) *ObjectAccessControlsService {
	rs := &ObjectAccessControlsService{s: s}
	return rs
}

type ObjectAccessControlsService struct {
	s *Service
}

func NewObjectsService(s *Service) *ObjectsService {
	rs := &ObjectsService{s: s}
	return rs
}

type ObjectsService struct {
	s *Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.ServiceAccount = NewProjectsServiceAccountService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	ServiceAccount *ProjectsServiceAccountService
}

func NewProjectsServiceAccountService(s *Service) *ProjectsServiceAccountService {
	rs := &ProjectsServiceAccountService{s: s}
	return rs
}

type ProjectsServiceAccountService struct {
	s *Service
}

// Bucket: A bucket.
type Bucket struct {
	// Acl: Access controls on the bucket.
	Acl []*BucketAccessControl `json:"acl,omitempty"`

	// Billing: The bucket's billing configuration.
	Billing *BucketBilling `json:"billing,omitempty"`

	// Cors: The bucket's Cross-Origin Resource Sharing (CORS)
	// configuration.
	Cors []*BucketCors `json:"cors,omitempty"`

	// DefaultEventBasedHold: Defines the default value for Event-Based hold
	// on newly created objects in this bucket. Event-Based hold is a way to
	// retain objects indefinitely until an event occurs, signified by the
	// hold's release. After being released, such objects will be subject to
	// bucket-level retention (if any). One sample use case of this flag is
	// for banks to hold loan documents for at least 3 years after loan is
	// paid in full. Here bucket-level retention is 3 years and the event is
	// loan being paid in full. In this example these objects will be held
	// intact for any number of years until the event has occurred (hold is
	// released) and then 3 more years after that. Objects under Event-Based
	// hold cannot be deleted, overwritten or archived until the hold is
	// removed.
	DefaultEventBasedHold bool `json:"defaultEventBasedHold,omitempty"`

	// DefaultObjectAcl: Default access controls to apply to new objects
	// when no ACL is provided.
	DefaultObjectAcl []*ObjectAccessControl `json:"defaultObjectAcl,omitempty"`

	// Encryption: Encryption configuration used by default for newly
	// inserted objects, when no encryption config is specified.
	Encryption *BucketEncryption `json:"encryption,omitempty"`

	// Etag: HTTP 1.1 Entity tag for the bucket.
	Etag string `json:"etag,omitempty"`

	// Id: The ID of the bucket. For buckets, the id and name properties are
	// the same.
	Id string `json:"id,omitempty"`

	// Kind: The kind of item this is. For buckets, this is always
	// storage#bucket.
	Kind string `json:"kind,omitempty"`

	// Labels: User-provided labels, in key/value pairs.
	Labels map[string]string `json:"labels,omitempty"`

	// Lifecycle: The bucket's lifecycle configuration. See lifecycle
	// management for more information.
	Lifecycle *BucketLifecycle `json:"lifecycle,omitempty"`

	// Location: The location of the bucket. Object data for objects in the
	// bucket resides in physical storage within this region. Defaults to
	// US. See the developer's guide for the authoritative list.
	Location string `json:"location,omitempty"`

	// Logging: The bucket's logging configuration, which defines the
	// destination bucket and optional name prefix for the current bucket's
	// logs.
	Logging *BucketLogging `json:"logging,omitempty"`

	// Metageneration: The metadata generation of this bucket.
	Metageneration int64 `json:"metageneration,omitempty,string"`

	// Name: The name of the bucket.
	Name string `json:"name,omitempty"`

	// Owner: The owner of the bucket. This is always the project team's
	// owner group.
	Owner *BucketOwner `json:"owner,omitempty"`

	// ProjectNumber: The project number of the project the bucket belongs
	// to.
	ProjectNumber uint64 `json:"projectNumber,omitempty,string"`

	// RetentionPolicy: Defines the retention policy for a bucket. The
	// Retention policy enforces a minimum retention time for all objects
	// contained in the bucket, based on their creation time. Any attempt to
	// overwrite or delete objects younger than the retention period will
	// result in a PERMISSION_DENIED error. An unlocked retention policy can
	// be modified or removed from the bucket via the UpdateBucketMetadata
	// RPC. A locked retention policy cannot be removed or shortened in
	// duration for the lifetime of the bucket. Attempting to remove or
	// decrease period of a locked retention policy will result in a
	// PERMISSION_DENIED error.
	RetentionPolicy *BucketRetentionPolicy `json:"retentionPolicy,omitempty"`

	// SelfLink: The URI of this bucket.
	SelfLink string `json:"selfLink,omitempty"`

	// StorageClass: The bucket's default storage class, used whenever no
	// storageClass is specified for a newly-created object. This defines
	// how objects in the bucket are stored and determines the SLA and the
	// cost of storage. Values include MULTI_REGIONAL, REGIONAL, STANDARD,
	// NEARLINE, COLDLINE, and DURABLE_REDUCED_AVAILABILITY. If this value
	// is not specified when the bucket is created, it will default to
	// STANDARD. For more information, see storage classes.
	StorageClass string `json:"storageClass,omitempty"`

	// TimeCreated: The creation time of the bucket in RFC 3339 format.
	TimeCreated string `json:"timeCreated,omitempty"`

	// Updated: The modification time of the bucket in RFC 3339 format.
	Updated string `json:"updated,omitempty"`

	// Versioning: The bucket's versioning configuration.
	Versioning *BucketVersioning `json:"versioning,omitempty"`

	// Website: The bucket's website configuration, controlling how the
	// service behaves when accessing bucket contents as a web site. See the
	// Static Website Examples for more information.
	Website *BucketWebsite `json:"website,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Acl") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Acl") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Bucket) MarshalJSON() ([]byte, error) {
	type NoMethod Bucket
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BucketBilling: The bucket's billing configuration.
type BucketBilling struct {
	// RequesterPays: When set to true, Requester Pays is enabled for this
	// bucket.
	RequesterPays bool `json:"requesterPays,omitempty"`

	// ForceSendFields is a list of field names (e.g. "RequesterPays") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "RequesterPays") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BucketBilling) MarshalJSON() ([]byte, error) {
	type NoMethod BucketBilling
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type BucketCors struct {
	// MaxAgeSeconds: The value, in seconds, to return in the
	// Access-Control-Max-Age header used in preflight responses.
	MaxAgeSeconds int64 `json:"maxAgeSeconds,omitempty"`

	// Method: The list of HTTP methods on which to include CORS response
	// headers, (GET, OPTIONS, POST, etc) Note: "*" is permitted in the list
	// of methods, and means "any method".
	Method []string `json:"method,omitempty"`

	// Origin: The list of Origins eligible to receive CORS response
	// headers. Note: "*" is permitted in the list of origins, and means
	// "any Origin".
	Origin []string `json:"origin,omitempty"`

	// ResponseHeader: The list of HTTP headers other than the simple
	// response headers to give permission for the user-agent to share
	// across domains.
	ResponseHeader []string `json:"responseHeader,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MaxAgeSeconds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MaxAgeSeconds") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BucketCors) MarshalJSON() ([]byte, error) {
	type NoMethod BucketCors
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BucketEncryption: Encryption configuration used by default for newly
// inserted objects, when no encryption config is specified.
type BucketEncryption struct {
	// DefaultKmsKeyName: A Cloud KMS key that will be used to encrypt
	// objects inserted into this bucket, if no encryption method is
	// specified. Limited availability; usable only by enabled projects.
	DefaultKmsKeyName string `json:"defaultKmsKeyName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DefaultKmsKeyName")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DefaultKmsKeyName") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *BucketEncryption) MarshalJSON() ([]byte, error) {
	type NoMethod BucketEncryption
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BucketLifecycle: The bucket's lifecycle configuration. See lifecycle
// management for more information.
type BucketLifecycle struct {
	// Rule: A lifecycle management rule, which is made of an action to take
	// and the condition(s) under which the action will be taken.
	Rule []*BucketLifecycleRule `json:"rule,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Rule") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Rule") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BucketLifecycle) MarshalJSON() ([]byte, error) {
	type NoMethod BucketLifecycle
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type BucketLifecycleRule struct {
	// Action: The action to take.
	Action *BucketLifecycleRuleAction `json:"action,omitempty"`

	// Condition: The condition(s) under which the action will be taken.
	Condition *BucketLifecycleRuleCondition `json:"condition,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BucketLifecycleRule) MarshalJSON() ([]byte, error) {
	type NoMethod BucketLifecycleRule
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BucketLifecycleRuleAction: The action to take.
type BucketLifecycleRuleAction struct {
	// StorageClass: Target storage class. Required iff the type of the
	// action is SetStorageClass.
	StorageClass string `json:"storageClass,omitempty"`

	// Type: Type of the action. Currently, only Delete and SetStorageClass
	// are supported.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "StorageClass") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "StorageClass") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BucketLifecycleRuleAction) MarshalJSON() ([]byte, error) {
	type NoMethod BucketLifecycleRuleAction
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BucketLifecycleRuleCondition: The condition(s) under which the action
// will be taken.
type BucketLifecycleRuleCondition struct {
	// Age: Age of an object (in days). This condition is satisfied when an
	// object reaches the specified age.
	Age int64 `json:"age,omitempty"`

	// CreatedBefore: A date in RFC 3339 format with only the date part (for
	// instance, "2013-01-15"). This condition is satisfied when an object
	// is created before midnight of the specified date in UTC.
	CreatedBefore string `json:"createdBefore,omitempty"`

	// IsLive: Relevant only for versioned objects. If the value is true,
	// this condition matches live objects; if the value is false, it
	// matches archived objects.
	IsLive *bool `json:"isLive,omitempty"`

	// MatchesStorageClass: Objects having any of the storage classes
	// specified by this condition will be matched. Values include
	// MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, STANDARD, and
	// DURABLE_REDUCED_AVAILABILITY.
	MatchesStorageClass []string `json:"matchesStorageClass,omitempty"`

	// NumNewerVersions: Relevant only for versioned objects. If the value
	// is N, this condition is satisfied when there are at least N versions
	// (including the live version) newer than this version of the object.
	NumNewerVersions int64 `json:"numNewerVersions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Age") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Age") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BucketLifecycleRuleCondition) MarshalJSON() ([]byte, error) {
	type NoMethod BucketLifecycleRuleCondition
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BucketLogging: The bucket's logging configuration, which defines the
// destination bucket and optional name prefix for the current bucket's
// logs.
type BucketLogging struct {
	// LogBucket: The destination bucket where the current bucket's logs
	// should be placed.
	LogBucket string `json:"logBucket,omitempty"`

	// LogObjectPrefix: A prefix for log object names.
	LogObjectPrefix string `json:"logObjectPrefix,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LogBucket") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LogBucket") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BucketLogging) MarshalJSON() ([]byte, error) {
	type NoMethod BucketLogging
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BucketOwner: The owner of the bucket. This is always the project
// team's owner group.
type BucketOwner struct {
	// Entity: The entity, in the form project-owner-projectId.
	Entity string `json:"entity,omitempty"`

	// EntityId: The ID for the entity.
	EntityId string `json:"entityId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Entity") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Entity") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BucketOwner) MarshalJSON() ([]byte, error) {
	type NoMethod BucketOwner
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BucketRetentionPolicy: Defines the retention policy for a bucket. The
// Retention policy enforces a minimum retention time for all objects
// contained in the bucket, based on their creation time. Any attempt to
// overwrite or delete objects younger than the retention period will
// result in a PERMISSION_DENIED error. An unlocked retention policy can
// be modified or removed from the bucket via the UpdateBucketMetadata
// RPC. A locked retention policy cannot be removed or shortened in
// duration for the lifetime of the bucket. Attempting to remove or
// decrease period of a locked retention policy will result in a
// PERMISSION_DENIED error.
type BucketRetentionPolicy struct {
	// EffectiveTime: The time from which policy was enforced and effective.
	// RFC 3339 format.
	EffectiveTime string `json:"effectiveTime,omitempty"`

	// IsLocked: Once locked, an object retention policy cannot be modified.
	IsLocked bool `json:"isLocked,omitempty"`

	// RetentionPeriod: Specifies the duration that objects need to be
	// retained. Retention duration must be greater than zero and less than
	// 100 years. Note that enforcement of retention periods less than a day
	// is not guaranteed. Such periods should only be used for testing
	// purposes.
	RetentionPeriod int64 `json:"retentionPeriod,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "EffectiveTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EffectiveTime") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BucketRetentionPolicy) MarshalJSON() ([]byte, error) {
	type NoMethod BucketRetentionPolicy
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BucketVersioning: The bucket's versioning configuration.
type BucketVersioning struct {
	// Enabled: While set to true, versioning is fully enabled for this
	// bucket.
	Enabled bool `json:"enabled,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Enabled") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Enabled") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BucketVersioning) MarshalJSON() ([]byte, error) {
	type NoMethod BucketVersioning
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BucketWebsite: The bucket's website configuration, controlling how
// the service behaves when accessing bucket contents as a web site. See
// the Static Website Examples for more information.
type BucketWebsite struct {
	// MainPageSuffix: If the requested object path is missing, the service
	// will ensure the path has a trailing '/', append this suffix, and
	// attempt to retrieve the resulting object. This allows the creation of
	// index.html objects to represent directory pages.
	MainPageSuffix string `json:"mainPageSuffix,omitempty"`

	// NotFoundPage: If the requested object path is missing, and any
	// mainPageSuffix object is missing, if applicable, the service will
	// return the named object from this bucket as the content for a 404 Not
	// Found result.
	NotFoundPage string `json:"notFoundPage,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MainPageSuffix") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MainPageSuffix") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *BucketWebsite) MarshalJSON() ([]byte, error) {
	type NoMethod BucketWebsite
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BucketAccessControl: An access-control entry.
type BucketAccessControl struct {
	// Bucket: The name of the bucket.
	Bucket string `json:"bucket,omitempty"`

	// Domain: The domain associated with the entity, if any.
	Domain string `json:"domain,omitempty"`

	// Email: The email address associated with the entity, if any.
	Email string `json:"email,omitempty"`

	// Entity: The entity holding the permission, in one of the following
	// forms:
	// - user-userId
	// - user-email
	// - group-groupId
	// - group-email
	// - domain-domain
	// - project-team-projectId
	// - allUsers
	// - allAuthenticatedUsers Examples:
	// - The user liz@example.com would be user-liz@example.com.
	// - The group example@googlegroups.com would be
	// group-example@googlegroups.com.
	// - To refer to all members of the Google Apps for Business domain
	// example.com, the entity would be domain-example.com.
	Entity string `json:"entity,omitempty"`

	// EntityId: The ID for the entity, if any.
	EntityId string `json:"entityId,omitempty"`

	// Etag: HTTP 1.1 Entity tag for the access-control entry.
	Etag string `json:"etag,omitempty"`

	// Id: The ID of the access-control entry.
	Id string `json:"id,omitempty"`

	// Kind: The kind of item this is. For bucket access control entries,
	// this is always storage#bucketAccessControl.
	Kind string `json:"kind,omitempty"`

	// ProjectTeam: The project team associated with the entity, if any.
	ProjectTeam *BucketAccessControlProjectTeam `json:"projectTeam,omitempty"`

	// Role: The access permission for the entity.
	Role string `json:"role,omitempty"`

	// SelfLink: The link to this access-control entry.
	SelfLink string `json:"selfLink,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Bucket") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bucket") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BucketAccessControl) MarshalJSON() ([]byte, error) {
	type NoMethod BucketAccessControl
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BucketAccessControlProjectTeam: The project team associated with the
// entity, if any.
type BucketAccessControlProjectTeam struct {
	// ProjectNumber: The project number.
	ProjectNumber string `json:"projectNumber,omitempty"`

	// Team: The team.
	Team string `json:"team,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ProjectNumber") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ProjectNumber") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BucketAccessControlProjectTeam) MarshalJSON() ([]byte, error) {
	type NoMethod BucketAccessControlProjectTeam
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BucketAccessControls: An access-control list.
type BucketAccessControls struct {
	// Items: The list of items.
	Items []*BucketAccessControl `json:"items,omitempty"`

	// Kind: The kind of item this is. For lists of bucket access control
	// entries, this is always storage#bucketAccessControls.
	Kind string `json:"kind,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BucketAccessControls) MarshalJSON() ([]byte, error) {
	type NoMethod BucketAccessControls
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Buckets: A list of buckets.
type Buckets struct {
	// Items: The list of items.
	Items []*Bucket `json:"items,omitempty"`

	// Kind: The kind of item this is. For lists of buckets, this is always
	// storage#buckets.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The continuation token, used to page through large
	// result sets. Provide this value in a subsequent request to return the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Buckets) MarshalJSON() ([]byte, error) {
	type NoMethod Buckets
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Channel: An notification channel used to watch for resource changes.
type Channel struct {
	// Address: The address where notifications are delivered for this
	// channel.
	Address string `json:"address,omitempty"`

	// Expiration: Date and time of notification channel expiration,
	// expressed as a Unix timestamp, in milliseconds. Optional.
	Expiration int64 `json:"expiration,omitempty,string"`

	// Id: A UUID or similar unique string that identifies this channel.
	Id string `json:"id,omitempty"`

	// Kind: Identifies this as a notification channel used to watch for
	// changes to a resource. Value: the fixed string "api#channel".
	Kind string `json:"kind,omitempty"`

	// Params: Additional parameters controlling delivery channel behavior.
	// Optional.
	Params map[string]string `json:"params,omitempty"`

	// Payload: A Boolean value to indicate whether payload is wanted.
	// Optional.
	Payload bool `json:"payload,omitempty"`

	// ResourceId: An opaque ID that identifies the resource being watched
	// on this channel. Stable across different API versions.
	ResourceId string `json:"resourceId,omitempty"`

	// ResourceUri: A version-specific identifier for the watched resource.
	ResourceUri string `json:"resourceUri,omitempty"`

	// Token: An arbitrary string delivered to the target address with each
	// notification delivered over this channel. Optional.
	Token string `json:"token,omitempty"`

	// Type: The type of delivery mechanism used for this channel.
	Type string `json:"type,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Address") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Address") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Channel) MarshalJSON() ([]byte, error) {
	type NoMethod Channel
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ComposeRequest: A Compose request.
type ComposeRequest struct {
	// Destination: Properties of the resulting object.
	Destination *Object `json:"destination,omitempty"`

	// Kind: The kind of item this is.
	Kind string `json:"kind,omitempty"`

	// SourceObjects: The list of source objects that will be concatenated
	// into a single object.
	SourceObjects []*ComposeRequestSourceObjects `json:"sourceObjects,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Destination") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Destination") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ComposeRequest) MarshalJSON() ([]byte, error) {
	type NoMethod ComposeRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ComposeRequestSourceObjects struct {
	// Generation: The generation of this object to use as the source.
	Generation int64 `json:"generation,omitempty,string"`

	// Name: The source object's name. The source object's bucket is
	// implicitly the destination bucket.
	Name string `json:"name,omitempty"`

	// ObjectPreconditions: Conditions that must be met for this operation
	// to execute.
	ObjectPreconditions *ComposeRequestSourceObjectsObjectPreconditions `json:"objectPreconditions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Generation") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Generation") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ComposeRequestSourceObjects) MarshalJSON() ([]byte, error) {
	type NoMethod ComposeRequestSourceObjects
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ComposeRequestSourceObjectsObjectPreconditions: Conditions that must
// be met for this operation to execute.
type ComposeRequestSourceObjectsObjectPreconditions struct {
	// IfGenerationMatch: Only perform the composition if the generation of
	// the source object that would be used matches this value. If this
	// value and a generation are both specified, they must be the same
	// value or the call will fail.
	IfGenerationMatch int64 `json:"ifGenerationMatch,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "IfGenerationMatch")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IfGenerationMatch") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ComposeRequestSourceObjectsObjectPreconditions) MarshalJSON() ([]byte, error) {
	type NoMethod ComposeRequestSourceObjectsObjectPreconditions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Notification: A subscription to receive Google PubSub notifications.
type Notification struct {
	// CustomAttributes: An optional list of additional attributes to attach
	// to each Cloud PubSub message published for this notification
	// subscription.
	CustomAttributes map[string]string `json:"custom_attributes,omitempty"`

	// Etag: HTTP 1.1 Entity tag for this subscription notification.
	Etag string `json:"etag,omitempty"`

	// EventTypes: If present, only send notifications about listed event
	// types. If empty, sent notifications for all event types.
	EventTypes []string `json:"event_types,omitempty"`

	// Id: The ID of the notification.
	Id string `json:"id,omitempty"`

	// Kind: The kind of item this is. For notifications, this is always
	// storage#notification.
	Kind string `json:"kind,omitempty"`

	// ObjectNamePrefix: If present, only apply this notification
	// configuration to object names that begin with this prefix.
	ObjectNamePrefix string `json:"object_name_prefix,omitempty"`

	// PayloadFormat: The desired content of the Payload.
	PayloadFormat string `json:"payload_format,omitempty"`

	// SelfLink: The canonical URL of this notification.
	SelfLink string `json:"selfLink,omitempty"`

	// Topic: The Cloud PubSub topic to which this subscription publishes.
	// Formatted as:
	// '//pubsub.googleapis.com/projects/{project-identifier}/topics/{my-topi
	// c}'
	Topic string `json:"topic,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CustomAttributes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CustomAttributes") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Notification) MarshalJSON() ([]byte, error) {
	type NoMethod Notification
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Notifications: A list of notification subscriptions.
type Notifications struct {
	// Items: The list of items.
	Items []*Notification `json:"items,omitempty"`

	// Kind: The kind of item this is. For lists of notifications, this is
	// always storage#notifications.
	Kind string `json:"kind,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Notifications) MarshalJSON() ([]byte, error) {
	type NoMethod Notifications
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Object: An object.
type Object struct {
	// Acl: Access controls on the object.
	Acl []*ObjectAccessControl `json:"acl,omitempty"`

	// Bucket: The name of the bucket containing this object.
	Bucket string `json:"bucket,omitempty"`

	// CacheControl: Cache-Control directive for the object data. If
	// omitted, and the object is accessible to all anonymous users, the
	// default will be public, max-age=3600.
	CacheControl string `json:"cacheControl,omitempty"`

	// ComponentCount: Number of underlying components that make up this
	// object. Components are accumulated by compose operations.
	ComponentCount int64 `json:"componentCount,omitempty"`

	// ContentDisposition: Content-Disposition of the object data.
	ContentDisposition string `json:"contentDisposition,omitempty"`

	// ContentEncoding: Content-Encoding of the object data.
	ContentEncoding string `json:"contentEncoding,omitempty"`

	// ContentLanguage: Content-Language of the object data.
	ContentLanguage string `json:"contentLanguage,omitempty"`

	// ContentType: Content-Type of the object data. If an object is stored
	// without a Content-Type, it is served as application/octet-stream.
	ContentType string `json:"contentType,omitempty"`

	// Crc32c: CRC32c checksum, as described in RFC 4960, Appendix B;
	// encoded using base64 in big-endian byte order. For more information
	// about using the CRC32c checksum, see Hashes and ETags: Best
	// Practices.
	Crc32c string `json:"crc32c,omitempty"`

	// CustomerEncryption: Metadata of customer-supplied encryption key, if
	// the object is encrypted by such a key.
	CustomerEncryption *ObjectCustomerEncryption `json:"customerEncryption,omitempty"`

	// Etag: HTTP 1.1 Entity tag for the object.
	Etag string `json:"etag,omitempty"`

	// EventBasedHold: Defines the Event-Based hold for an object.
	// Event-Based hold is a way to retain objects indefinitely until an
	// event occurs, signified by the hold's release. After being released,
	// such objects will be subject to bucket-level retention (if any). One
	// sample use case of this flag is for banks to hold loan documents for
	// at least 3 years after loan is paid in full. Here bucket-level
	// retention is 3 years and the event is loan being paid in full. In
	// this example these objects will be held intact for any number of
	// years until the event has occurred (hold is released) and then 3 more
	// years after that.
	EventBasedHold bool `json:"eventBasedHold,omitempty"`

	// Generation: The content generation of this object. Used for object
	// versioning.
	Generation int64 `json:"generation,omitempty,string"`

	// Id: The ID of the object, including the bucket name, object name, and
	// generation number.
	Id string `json:"id,omitempty"`

	// Kind: The kind of item this is. For objects, this is always
	// storage#object.
	Kind string `json:"kind,omitempty"`

	// KmsKeyName: Cloud KMS Key used to encrypt this object, if the object
	// is encrypted by such a key. Limited availability; usable only by
	// enabled projects.
	KmsKeyName string `json:"kmsKeyName,omitempty"`

	// Md5Hash: MD5 hash of the data; encoded using base64. For more
	// information about using the MD5 hash, see Hashes and ETags: Best
	// Practices.
	Md5Hash string `json:"md5Hash,omitempty"`

	// MediaLink: Media download link.
	MediaLink string `json:"mediaLink,omitempty"`

	// Metadata: User-provided metadata, in key/value pairs.
	Metadata map[string]string `json:"metadata,omitempty"`

	// Metageneration: The version of the metadata for this object at this
	// generation. Used for preconditions and for detecting changes in
	// metadata. A metageneration number is only meaningful in the context
	// of a particular generation of a particular object.
	Metageneration int64 `json:"metageneration,omitempty,string"`

	// Name: The name of the object. Required if not specified by URL
	// parameter.
	Name string `json:"name,omitempty"`

	// Owner: The owner of the object. This will always be the uploader of
	// the object.
	Owner *ObjectOwner `json:"owner,omitempty"`

	// RetentionExpirationTime: Specifies the earliest time that the
	// object's retention period expires. This value is server-determined
	// and is in RFC 3339 format. Note 1: This field is not provided for
	// objects with an active Event-Based hold, since retention expiration
	// is unknown until the hold is removed. Note 2: This value can be
	// provided even when TemporaryHold is set (so that the user can reason
	// about policy without having to first unset the TemporaryHold).
	RetentionExpirationTime string `json:"retentionExpirationTime,omitempty"`

	// SelfLink: The link to this object.
	SelfLink string `json:"selfLink,omitempty"`

	// Size: Content-Length of the data in bytes.
	Size uint64 `json:"size,omitempty,string"`

	// StorageClass: Storage class of the object.
	StorageClass string `json:"storageClass,omitempty"`

	// TemporaryHold: Defines the temporary hold for an object. This flag is
	// used to enforce a temporary hold on an object. While it is set to
	// true, the object is protected against deletion and overwrites. A
	// common use case of this flag is regulatory investigations where
	// objects need to be retained while the investigation is ongoing.
	TemporaryHold bool `json:"temporaryHold,omitempty"`

	// TimeCreated: The creation time of the object in RFC 3339 format.
	TimeCreated string `json:"timeCreated,omitempty"`

	// TimeDeleted: The deletion time of the object in RFC 3339 format. Will
	// be returned if and only if this version of the object has been
	// deleted.
	TimeDeleted string `json:"timeDeleted,omitempty"`

	// TimeStorageClassUpdated: The time at which the object's storage class
	// was last changed. When the object is initially created, it will be
	// set to timeCreated.
	TimeStorageClassUpdated string `json:"timeStorageClassUpdated,omitempty"`

	// Updated: The modification time of the object metadata in RFC 3339
	// format.
	Updated string `json:"updated,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Acl") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Acl") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Object) MarshalJSON() ([]byte, error) {
	type NoMethod Object
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ObjectCustomerEncryption: Metadata of customer-supplied encryption
// key, if the object is encrypted by such a key.
type ObjectCustomerEncryption struct {
	// EncryptionAlgorithm: The encryption algorithm.
	EncryptionAlgorithm string `json:"encryptionAlgorithm,omitempty"`

	// KeySha256: SHA256 hash value of the encryption key.
	KeySha256 string `json:"keySha256,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EncryptionAlgorithm")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EncryptionAlgorithm") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ObjectCustomerEncryption) MarshalJSON() ([]byte, error) {
	type NoMethod ObjectCustomerEncryption
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ObjectOwner: The owner of the object. This will always be the
// uploader of the object.
type ObjectOwner struct {
	// Entity: The entity, in the form user-userId.
	Entity string `json:"entity,omitempty"`

	// EntityId: The ID for the entity.
	EntityId string `json:"entityId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Entity") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Entity") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ObjectOwner) MarshalJSON() ([]byte, error) {
	type NoMethod ObjectOwner
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ObjectAccessControl: An access-control entry.
type ObjectAccessControl struct {
	// Bucket: The name of the bucket.
	Bucket string `json:"bucket,omitempty"`

	// Domain: The domain associated with the entity, if any.
	Domain string `json:"domain,omitempty"`

	// Email: The email address associated with the entity, if any.
	Email string `json:"email,omitempty"`

	// Entity: The entity holding the permission, in one of the following
	// forms:
	// - user-userId
	// - user-email
	// - group-groupId
	// - group-email
	// - domain-domain
	// - project-team-projectId
	// - allUsers
	// - allAuthenticatedUsers Examples:
	// - The user liz@example.com would be user-liz@example.com.
	// - The group example@googlegroups.com would be
	// group-example@googlegroups.com.
	// - To refer to all members of the Google Apps for Business domain
	// example.com, the entity would be domain-example.com.
	Entity string `json:"entity,omitempty"`

	// EntityId: The ID for the entity, if any.
	EntityId string `json:"entityId,omitempty"`

	// Etag: HTTP 1.1 Entity tag for the access-control entry.
	Etag string `json:"etag,omitempty"`

	// Generation: The content generation of the object, if applied to an
	// object.
	Generation int64 `json:"generation,omitempty,string"`

	// Id: The ID of the access-control entry.
	Id string `json:"id,omitempty"`

	// Kind: The kind of item this is. For object access control entries,
	// this is always storage#objectAccessControl.
	Kind string `json:"kind,omitempty"`

	// Object: The name of the object, if applied to an object.
	Object string `json:"object,omitempty"`

	// ProjectTeam: The project team associated with the entity, if any.
	ProjectTeam *ObjectAccessControlProjectTeam `json:"projectTeam,omitempty"`

	// Role: The access permission for the entity.
	Role string `json:"role,omitempty"`

	// SelfLink: The link to this access-control entry.
	SelfLink string `json:"selfLink,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Bucket") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bucket") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ObjectAccessControl) MarshalJSON() ([]byte, error) {
	type NoMethod ObjectAccessControl
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ObjectAccessControlProjectTeam: The project team associated with the
// entity, if any.
type ObjectAccessControlProjectTeam struct {
	// ProjectNumber: The project number.
	ProjectNumber string `json:"projectNumber,omitempty"`

	// Team: The team.
	Team string `json:"team,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ProjectNumber") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ProjectNumber") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ObjectAccessControlProjectTeam) MarshalJSON() ([]byte, error) {
	type NoMethod ObjectAccessControlProjectTeam
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ObjectAccessControls: An access-control list.
type ObjectAccessControls struct {
	// Items: The list of items.
	Items []*ObjectAccessControl `json:"items,omitempty"`

	// Kind: The kind of item this is. For lists of object access control
	// entries, this is always storage#objectAccessControls.
	Kind string `json:"kind,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ObjectAccessControls) MarshalJSON() ([]byte, error) {
	type NoMethod ObjectAccessControls
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Objects: A list of objects.
type Objects struct {
	// Items: The list of items.
	Items []*Object `json:"items,omitempty"`

	// Kind: The kind of item this is. For lists of objects, this is always
	// storage#objects.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The continuation token, used to page through large
	// result sets. Provide this value in a subsequent request to return the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Prefixes: The list of prefixes of objects matching-but-not-listed up
	// to and including the requested delimiter.
	Prefixes []string `json:"prefixes,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Objects) MarshalJSON() ([]byte, error) {
	type NoMethod Objects
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Policy: A bucket/object IAM policy.
type Policy struct {
	// Bindings: An association between a role, which comes with a set of
	// permissions, and members who may assume that role.
	Bindings []*PolicyBindings `json:"bindings,omitempty"`

	// Etag: HTTP 1.1  Entity tag for the policy.
	Etag string `json:"etag,omitempty"`

	// Kind: The kind of item this is. For policies, this is always
	// storage#policy. This field is ignored on input.
	Kind string `json:"kind,omitempty"`

	// ResourceId: The ID of the resource to which this policy belongs. Will
	// be of the form projects/_/buckets/bucket for buckets, and
	// projects/_/buckets/bucket/objects/object for objects. A specific
	// generation may be specified by appending #generationNumber to the end
	// of the object name, e.g.
	// projects/_/buckets/my-bucket/objects/data.txt#17. The current
	// generation can be denoted with #0. This field is ignored on input.
	ResourceId string `json:"resourceId,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Bindings") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bindings") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Policy) MarshalJSON() ([]byte, error) {
	type NoMethod Policy
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PolicyBindings struct {
	Condition interface{} `json:"condition,omitempty"`

	// Members: A collection of identifiers for members who may assume the
	// provided role. Recognized identifiers are as follows:
	// - allUsers — A special identifier that represents anyone on the
	// internet; with or without a Google account.
	// - allAuthenticatedUsers — A special identifier that represents
	// anyone who is authenticated with a Google account or a service
	// account.
	// - user:emailid — An email address that represents a specific
	// account. For example, user:alice@gmail.com or user:joe@example.com.
	//
	// - serviceAccount:emailid — An email address that represents a
	// service account. For example,
	// serviceAccount:my-other-app@appspot.gserviceaccount.com .
	// - group:emailid — An email address that represents a Google group.
	// For example, group:admins@example.com.
	// - domain:domain — A Google Apps domain name that represents all the
	// users of that domain. For example, domain:google.com or
	// domain:example.com.
	// - projectOwner:projectid — Owners of the given project. For
	// example, projectOwner:my-example-project
	// - projectEditor:projectid — Editors of the given project. For
	// example, projectEditor:my-example-project
	// - projectViewer:projectid — Viewers of the given project. For
	// example, projectViewer:my-example-project
	Members []string `json:"members,omitempty"`

	// Role: The role to which members belong. Two types of roles are
	// supported: new IAM roles, which grant permissions that do not map
	// directly to those provided by ACLs, and legacy IAM roles, which do
	// map directly to ACL permissions. All roles are of the format
	// roles/storage.specificRole.
	// The new IAM roles are:
	// - roles/storage.admin — Full control of Google Cloud Storage
	// resources.
	// - roles/storage.objectViewer — Read-Only access to Google Cloud
	// Storage objects.
	// - roles/storage.objectCreator — Access to create objects in Google
	// Cloud Storage.
	// - roles/storage.objectAdmin — Full control of Google Cloud Storage
	// objects.   The legacy IAM roles are:
	// - roles/storage.legacyObjectReader — Read-only access to objects
	// without listing. Equivalent to an ACL entry on an object with the
	// READER role.
	// - roles/storage.legacyObjectOwner — Read/write access to existing
	// objects without listing. Equivalent to an ACL entry on an object with
	// the OWNER role.
	// - roles/storage.legacyBucketReader — Read access to buckets with
	// object listing. Equivalent to an ACL entry on a bucket with the
	// READER role.
	// - roles/storage.legacyBucketWriter — Read access to buckets with
	// object listing/creation/deletion. Equivalent to an ACL entry on a
	// bucket with the WRITER role.
	// - roles/storage.legacyBucketOwner — Read and write access to
	// existing buckets with object listing/creation/deletion. Equivalent to
	// an ACL entry on a bucket with the OWNER role.
	Role string `json:"role,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Condition") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Condition") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PolicyBindings) MarshalJSON() ([]byte, error) {
	type NoMethod PolicyBindings
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RewriteResponse: A rewrite response.
type RewriteResponse struct {
	// Done: true if the copy is finished; otherwise, false if the copy is
	// in progress. This property is always present in the response.
	Done bool `json:"done,omitempty"`

	// Kind: The kind of item this is.
	Kind string `json:"kind,omitempty"`

	// ObjectSize: The total size of the object being copied in bytes. This
	// property is always present in the response.
	ObjectSize int64 `json:"objectSize,omitempty,string"`

	// Resource: A resource containing the metadata for the copied-to
	// object. This property is present in the response only when copying
	// completes.
	Resource *Object `json:"resource,omitempty"`

	// RewriteToken: A token to use in subsequent requests to continue
	// copying data. This token is present in the response only when there
	// is more data to copy.
	RewriteToken string `json:"rewriteToken,omitempty"`

	// TotalBytesRewritten: The total bytes written so far, which can be
	// used to provide a waiting user with a progress indicator. This
	// property is always present in the response.
	TotalBytesRewritten int64 `json:"totalBytesRewritten,omitempty,string"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RewriteResponse) MarshalJSON() ([]byte, error) {
	type NoMethod RewriteResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ServiceAccount: A subscription to receive Google PubSub
// notifications.
type ServiceAccount struct {
	// EmailAddress: The ID of the notification.
	EmailAddress string `json:"email_address,omitempty"`

	// Kind: The kind of item this is. For notifications, this is always
	// storage#notification.
	Kind string `json:"kind,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "EmailAddress") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EmailAddress") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ServiceAccount) MarshalJSON() ([]byte, error) {
	type NoMethod ServiceAccount
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TestIamPermissionsResponse: A
// storage.(buckets|objects).testIamPermissions response.
type TestIamPermissionsResponse struct {
	// Kind: The kind of item this is.
	Kind string `json:"kind,omitempty"`

	// Permissions: The permissions held by the caller. Permissions are
	// always of the format storage.resource.capability, where resource is
	// one of buckets or objects. The supported permissions are as follows:
	//
	// - storage.buckets.delete — Delete bucket.
	// - storage.buckets.get — Read bucket metadata.
	// - storage.buckets.getIamPolicy — Read bucket IAM policy.
	// - storage.buckets.create — Create bucket.
	// - storage.buckets.list — List buckets.
	// - storage.buckets.setIamPolicy — Update bucket IAM policy.
	// - storage.buckets.update — Update bucket metadata.
	// - storage.objects.delete — Delete object.
	// - storage.objects.get — Read object data and metadata.
	// - storage.objects.getIamPolicy — Read object IAM policy.
	// - storage.objects.create — Create object.
	// - storage.objects.list — List objects.
	// - storage.objects.setIamPolicy — Update object IAM policy.
	// - storage.objects.update — Update object metadata.
	Permissions []string `json:"permissions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Kind") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TestIamPermissionsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod TestIamPermissionsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "storage.bucketAccessControls.delete":

type BucketAccessControlsDeleteCall struct {
	s          *Service
	bucket     string
	entity     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Permanently deletes the ACL entry for the specified entity on
// the specified bucket.
func (r *BucketAccessControlsService) Delete(bucket string, entity string) *BucketAccessControlsDeleteCall {
	c := &BucketAccessControlsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.entity = entity
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *BucketAccessControlsDeleteCall) UserProject(userProject string) *BucketAccessControlsDeleteCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketAccessControlsDeleteCall) Fields(s ...googleapi.Field) *BucketAccessControlsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BucketAccessControlsDeleteCall) Context(ctx context.Context) *BucketAccessControlsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BucketAccessControlsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BucketAccessControlsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/acl/{entity}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"entity": c.entity,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.bucketAccessControls.delete" call.
func (c *BucketAccessControlsDeleteCall) Do(opts ...googleapi.CallOption) error {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Permanently deletes the ACL entry for the specified entity on the specified bucket.",
	//   "httpMethod": "DELETE",
	//   "id": "storage.bucketAccessControls.delete",
	//   "parameterOrder": [
	//     "bucket",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/acl/{entity}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.bucketAccessControls.get":

type BucketAccessControlsGetCall struct {
	s            *Service
	bucket       string
	entity       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Returns the ACL entry for the specified entity on the specified
// bucket.
func (r *BucketAccessControlsService) Get(bucket string, entity string) *BucketAccessControlsGetCall {
	c := &BucketAccessControlsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.entity = entity
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *BucketAccessControlsGetCall) UserProject(userProject string) *BucketAccessControlsGetCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketAccessControlsGetCall) Fields(s ...googleapi.Field) *BucketAccessControlsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BucketAccessControlsGetCall) IfNoneMatch(entityTag string) *BucketAccessControlsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BucketAccessControlsGetCall) Context(ctx context.Context) *BucketAccessControlsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BucketAccessControlsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BucketAccessControlsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/acl/{entity}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"entity": c.entity,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.bucketAccessControls.get" call.
// Exactly one of *BucketAccessControl or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *BucketAccessControl.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BucketAccessControlsGetCall) Do(opts ...googleapi.CallOption) (*BucketAccessControl, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BucketAccessControl{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the ACL entry for the specified entity on the specified bucket.",
	//   "httpMethod": "GET",
	//   "id": "storage.bucketAccessControls.get",
	//   "parameterOrder": [
	//     "bucket",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/acl/{entity}",
	//   "response": {
	//     "$ref": "BucketAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.bucketAccessControls.insert":

type BucketAccessControlsInsertCall struct {
	s                   *Service
	bucket              string
	bucketaccesscontrol *BucketAccessControl
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Insert: Creates a new ACL entry on the specified bucket.
func (r *BucketAccessControlsService) Insert(bucket string, bucketaccesscontrol *BucketAccessControl) *BucketAccessControlsInsertCall {
	c := &BucketAccessControlsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.bucketaccesscontrol = bucketaccesscontrol
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *BucketAccessControlsInsertCall) UserProject(userProject string) *BucketAccessControlsInsertCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketAccessControlsInsertCall) Fields(s ...googleapi.Field) *BucketAccessControlsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BucketAccessControlsInsertCall) Context(ctx context.Context) *BucketAccessControlsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BucketAccessControlsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BucketAccessControlsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.bucketaccesscontrol)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/acl")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.bucketAccessControls.insert" call.
// Exactly one of *BucketAccessControl or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *BucketAccessControl.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BucketAccessControlsInsertCall) Do(opts ...googleapi.CallOption) (*BucketAccessControl, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BucketAccessControl{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new ACL entry on the specified bucket.",
	//   "httpMethod": "POST",
	//   "id": "storage.bucketAccessControls.insert",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/acl",
	//   "request": {
	//     "$ref": "BucketAccessControl"
	//   },
	//   "response": {
	//     "$ref": "BucketAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.bucketAccessControls.list":

type BucketAccessControlsListCall struct {
	s            *Service
	bucket       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves ACL entries on the specified bucket.
func (r *BucketAccessControlsService) List(bucket string) *BucketAccessControlsListCall {
	c := &BucketAccessControlsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *BucketAccessControlsListCall) UserProject(userProject string) *BucketAccessControlsListCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketAccessControlsListCall) Fields(s ...googleapi.Field) *BucketAccessControlsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BucketAccessControlsListCall) IfNoneMatch(entityTag string) *BucketAccessControlsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BucketAccessControlsListCall) Context(ctx context.Context) *BucketAccessControlsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BucketAccessControlsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BucketAccessControlsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/acl")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.bucketAccessControls.list" call.
// Exactly one of *BucketAccessControls or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *BucketAccessControls.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BucketAccessControlsListCall) Do(opts ...googleapi.CallOption) (*BucketAccessControls, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BucketAccessControls{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves ACL entries on the specified bucket.",
	//   "httpMethod": "GET",
	//   "id": "storage.bucketAccessControls.list",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/acl",
	//   "response": {
	//     "$ref": "BucketAccessControls"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.bucketAccessControls.patch":

type BucketAccessControlsPatchCall struct {
	s                   *Service
	bucket              string
	entity              string
	bucketaccesscontrol *BucketAccessControl
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Patch: Updates an ACL entry on the specified bucket. This method
// supports patch semantics.
func (r *BucketAccessControlsService) Patch(bucket string, entity string, bucketaccesscontrol *BucketAccessControl) *BucketAccessControlsPatchCall {
	c := &BucketAccessControlsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.entity = entity
	c.bucketaccesscontrol = bucketaccesscontrol
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *BucketAccessControlsPatchCall) UserProject(userProject string) *BucketAccessControlsPatchCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketAccessControlsPatchCall) Fields(s ...googleapi.Field) *BucketAccessControlsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BucketAccessControlsPatchCall) Context(ctx context.Context) *BucketAccessControlsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BucketAccessControlsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BucketAccessControlsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.bucketaccesscontrol)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/acl/{entity}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"entity": c.entity,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.bucketAccessControls.patch" call.
// Exactly one of *BucketAccessControl or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *BucketAccessControl.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BucketAccessControlsPatchCall) Do(opts ...googleapi.CallOption) (*BucketAccessControl, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BucketAccessControl{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an ACL entry on the specified bucket. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "storage.bucketAccessControls.patch",
	//   "parameterOrder": [
	//     "bucket",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/acl/{entity}",
	//   "request": {
	//     "$ref": "BucketAccessControl"
	//   },
	//   "response": {
	//     "$ref": "BucketAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.bucketAccessControls.update":

type BucketAccessControlsUpdateCall struct {
	s                   *Service
	bucket              string
	entity              string
	bucketaccesscontrol *BucketAccessControl
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Update: Updates an ACL entry on the specified bucket.
func (r *BucketAccessControlsService) Update(bucket string, entity string, bucketaccesscontrol *BucketAccessControl) *BucketAccessControlsUpdateCall {
	c := &BucketAccessControlsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.entity = entity
	c.bucketaccesscontrol = bucketaccesscontrol
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *BucketAccessControlsUpdateCall) UserProject(userProject string) *BucketAccessControlsUpdateCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketAccessControlsUpdateCall) Fields(s ...googleapi.Field) *BucketAccessControlsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BucketAccessControlsUpdateCall) Context(ctx context.Context) *BucketAccessControlsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BucketAccessControlsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BucketAccessControlsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.bucketaccesscontrol)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/acl/{entity}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"entity": c.entity,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.bucketAccessControls.update" call.
// Exactly one of *BucketAccessControl or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *BucketAccessControl.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BucketAccessControlsUpdateCall) Do(opts ...googleapi.CallOption) (*BucketAccessControl, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BucketAccessControl{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an ACL entry on the specified bucket.",
	//   "httpMethod": "PUT",
	//   "id": "storage.bucketAccessControls.update",
	//   "parameterOrder": [
	//     "bucket",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/acl/{entity}",
	//   "request": {
	//     "$ref": "BucketAccessControl"
	//   },
	//   "response": {
	//     "$ref": "BucketAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.buckets.delete":

type BucketsDeleteCall struct {
	s          *Service
	bucket     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Permanently deletes an empty bucket.
func (r *BucketsService) Delete(bucket string) *BucketsDeleteCall {
	c := &BucketsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	return c
}

// IfMetagenerationMatch sets the optional parameter
// "ifMetagenerationMatch": If set, only deletes the bucket if its
// metageneration matches this value.
func (c *BucketsDeleteCall) IfMetagenerationMatch(ifMetagenerationMatch int64) *BucketsDeleteCall {
	c.urlParams_.Set("ifMetagenerationMatch", fmt.Sprint(ifMetagenerationMatch))
	return c
}

// IfMetagenerationNotMatch sets the optional parameter
// "ifMetagenerationNotMatch": If set, only deletes the bucket if its
// metageneration does not match this value.
func (c *BucketsDeleteCall) IfMetagenerationNotMatch(ifMetagenerationNotMatch int64) *BucketsDeleteCall {
	c.urlParams_.Set("ifMetagenerationNotMatch", fmt.Sprint(ifMetagenerationNotMatch))
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *BucketsDeleteCall) UserProject(userProject string) *BucketsDeleteCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketsDeleteCall) Fields(s ...googleapi.Field) *BucketsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BucketsDeleteCall) Context(ctx context.Context) *BucketsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BucketsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BucketsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.buckets.delete" call.
func (c *BucketsDeleteCall) Do(opts ...googleapi.CallOption) error {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Permanently deletes an empty bucket.",
	//   "httpMethod": "DELETE",
	//   "id": "storage.buckets.delete",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ifMetagenerationMatch": {
	//       "description": "If set, only deletes the bucket if its metageneration matches this value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationNotMatch": {
	//       "description": "If set, only deletes the bucket if its metageneration does not match this value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.buckets.get":

type BucketsGetCall struct {
	s            *Service
	bucket       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Returns metadata for the specified bucket.
func (r *BucketsService) Get(bucket string) *BucketsGetCall {
	c := &BucketsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	return c
}

// IfMetagenerationMatch sets the optional parameter
// "ifMetagenerationMatch": Makes the return of the bucket metadata
// conditional on whether the bucket's current metageneration matches
// the given value.
func (c *BucketsGetCall) IfMetagenerationMatch(ifMetagenerationMatch int64) *BucketsGetCall {
	c.urlParams_.Set("ifMetagenerationMatch", fmt.Sprint(ifMetagenerationMatch))
	return c
}

// IfMetagenerationNotMatch sets the optional parameter
// "ifMetagenerationNotMatch": Makes the return of the bucket metadata
// conditional on whether the bucket's current metageneration does not
// match the given value.
func (c *BucketsGetCall) IfMetagenerationNotMatch(ifMetagenerationNotMatch int64) *BucketsGetCall {
	c.urlParams_.Set("ifMetagenerationNotMatch", fmt.Sprint(ifMetagenerationNotMatch))
	return c
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to noAcl.
//
// Possible values:
//   "full" - Include all properties.
//   "noAcl" - Omit owner, acl and defaultObjectAcl properties.
func (c *BucketsGetCall) Projection(projection string) *BucketsGetCall {
	c.urlParams_.Set("projection", projection)
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *BucketsGetCall) UserProject(userProject string) *BucketsGetCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketsGetCall) Fields(s ...googleapi.Field) *BucketsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BucketsGetCall) IfNoneMatch(entityTag string) *BucketsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BucketsGetCall) Context(ctx context.Context) *BucketsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BucketsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BucketsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.buckets.get" call.
// Exactly one of *Bucket or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Bucket.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *BucketsGetCall) Do(opts ...googleapi.CallOption) (*Bucket, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Bucket{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns metadata for the specified bucket.",
	//   "httpMethod": "GET",
	//   "id": "storage.buckets.get",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ifMetagenerationMatch": {
	//       "description": "Makes the return of the bucket metadata conditional on whether the bucket's current metageneration matches the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationNotMatch": {
	//       "description": "Makes the return of the bucket metadata conditional on whether the bucket's current metageneration does not match the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to noAcl.",
	//       "enum": [
	//         "full",
	//         "noAcl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit owner, acl and defaultObjectAcl properties."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}",
	//   "response": {
	//     "$ref": "Bucket"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.buckets.getIamPolicy":

type BucketsGetIamPolicyCall struct {
	s            *Service
	bucket       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetIamPolicy: Returns an IAM policy for the specified bucket.
func (r *BucketsService) GetIamPolicy(bucket string) *BucketsGetIamPolicyCall {
	c := &BucketsGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *BucketsGetIamPolicyCall) UserProject(userProject string) *BucketsGetIamPolicyCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketsGetIamPolicyCall) Fields(s ...googleapi.Field) *BucketsGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BucketsGetIamPolicyCall) IfNoneMatch(entityTag string) *BucketsGetIamPolicyCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BucketsGetIamPolicyCall) Context(ctx context.Context) *BucketsGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BucketsGetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BucketsGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/iam")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.buckets.getIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *BucketsGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns an IAM policy for the specified bucket.",
	//   "httpMethod": "GET",
	//   "id": "storage.buckets.getIamPolicy",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/iam",
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.buckets.insert":

type BucketsInsertCall struct {
	s          *Service
	bucket     *Bucket
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Insert: Creates a new bucket.
func (r *BucketsService) Insert(projectid string, bucket *Bucket) *BucketsInsertCall {
	c := &BucketsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.urlParams_.Set("project", projectid)
	c.bucket = bucket
	return c
}

// PredefinedAcl sets the optional parameter "predefinedAcl": Apply a
// predefined set of access controls to this bucket.
//
// Possible values:
//   "authenticatedRead" - Project team owners get OWNER access, and
// allAuthenticatedUsers get READER access.
//   "private" - Project team owners get OWNER access.
//   "projectPrivate" - Project team members get access according to
// their roles.
//   "publicRead" - Project team owners get OWNER access, and allUsers
// get READER access.
//   "publicReadWrite" - Project team owners get OWNER access, and
// allUsers get WRITER access.
func (c *BucketsInsertCall) PredefinedAcl(predefinedAcl string) *BucketsInsertCall {
	c.urlParams_.Set("predefinedAcl", predefinedAcl)
	return c
}

// PredefinedDefaultObjectAcl sets the optional parameter
// "predefinedDefaultObjectAcl": Apply a predefined set of default
// object access controls to this bucket.
//
// Possible values:
//   "authenticatedRead" - Object owner gets OWNER access, and
// allAuthenticatedUsers get READER access.
//   "bucketOwnerFullControl" - Object owner gets OWNER access, and
// project team owners get OWNER access.
//   "bucketOwnerRead" - Object owner gets OWNER access, and project
// team owners get READER access.
//   "private" - Object owner gets OWNER access.
//   "projectPrivate" - Object owner gets OWNER access, and project team
// members get access according to their roles.
//   "publicRead" - Object owner gets OWNER access, and allUsers get
// READER access.
func (c *BucketsInsertCall) PredefinedDefaultObjectAcl(predefinedDefaultObjectAcl string) *BucketsInsertCall {
	c.urlParams_.Set("predefinedDefaultObjectAcl", predefinedDefaultObjectAcl)
	return c
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to noAcl, unless the bucket resource
// specifies acl or defaultObjectAcl properties, when it defaults to
// full.
//
// Possible values:
//   "full" - Include all properties.
//   "noAcl" - Omit owner, acl and defaultObjectAcl properties.
func (c *BucketsInsertCall) Projection(projection string) *BucketsInsertCall {
	c.urlParams_.Set("projection", projection)
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request.
func (c *BucketsInsertCall) UserProject(userProject string) *BucketsInsertCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketsInsertCall) Fields(s ...googleapi.Field) *BucketsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BucketsInsertCall) Context(ctx context.Context) *BucketsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BucketsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BucketsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.bucket)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.buckets.insert" call.
// Exactly one of *Bucket or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Bucket.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *BucketsInsertCall) Do(opts ...googleapi.CallOption) (*Bucket, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Bucket{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new bucket.",
	//   "httpMethod": "POST",
	//   "id": "storage.buckets.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "predefinedAcl": {
	//       "description": "Apply a predefined set of access controls to this bucket.",
	//       "enum": [
	//         "authenticatedRead",
	//         "private",
	//         "projectPrivate",
	//         "publicRead",
	//         "publicReadWrite"
	//       ],
	//       "enumDescriptions": [
	//         "Project team owners get OWNER access, and allAuthenticatedUsers get READER access.",
	//         "Project team owners get OWNER access.",
	//         "Project team members get access according to their roles.",
	//         "Project team owners get OWNER access, and allUsers get READER access.",
	//         "Project team owners get OWNER access, and allUsers get WRITER access."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "predefinedDefaultObjectAcl": {
	//       "description": "Apply a predefined set of default object access controls to this bucket.",
	//       "enum": [
	//         "authenticatedRead",
	//         "bucketOwnerFullControl",
	//         "bucketOwnerRead",
	//         "private",
	//         "projectPrivate",
	//         "publicRead"
	//       ],
	//       "enumDescriptions": [
	//         "Object owner gets OWNER access, and allAuthenticatedUsers get READER access.",
	//         "Object owner gets OWNER access, and project team owners get OWNER access.",
	//         "Object owner gets OWNER access, and project team owners get READER access.",
	//         "Object owner gets OWNER access.",
	//         "Object owner gets OWNER access, and project team members get access according to their roles.",
	//         "Object owner gets OWNER access, and allUsers get READER access."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "A valid API project identifier.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to noAcl, unless the bucket resource specifies acl or defaultObjectAcl properties, when it defaults to full.",
	//       "enum": [
	//         "full",
	//         "noAcl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit owner, acl and defaultObjectAcl properties."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b",
	//   "request": {
	//     "$ref": "Bucket"
	//   },
	//   "response": {
	//     "$ref": "Bucket"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.buckets.list":

type BucketsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of buckets for a given project.
func (r *BucketsService) List(projectid string) *BucketsListCall {
	c := &BucketsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.urlParams_.Set("project", projectid)
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of buckets to return in a single response. The service will use this
// parameter or 1,000 items, whichever is smaller.
func (c *BucketsListCall) MaxResults(maxResults int64) *BucketsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A
// previously-returned page token representing part of the larger set of
// results to view.
func (c *BucketsListCall) PageToken(pageToken string) *BucketsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Prefix sets the optional parameter "prefix": Filter results to
// buckets whose names begin with this prefix.
func (c *BucketsListCall) Prefix(prefix string) *BucketsListCall {
	c.urlParams_.Set("prefix", prefix)
	return c
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to noAcl.
//
// Possible values:
//   "full" - Include all properties.
//   "noAcl" - Omit owner, acl and defaultObjectAcl properties.
func (c *BucketsListCall) Projection(projection string) *BucketsListCall {
	c.urlParams_.Set("projection", projection)
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request.
func (c *BucketsListCall) UserProject(userProject string) *BucketsListCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketsListCall) Fields(s ...googleapi.Field) *BucketsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BucketsListCall) IfNoneMatch(entityTag string) *BucketsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BucketsListCall) Context(ctx context.Context) *BucketsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BucketsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BucketsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.buckets.list" call.
// Exactly one of *Buckets or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Buckets.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *BucketsListCall) Do(opts ...googleapi.CallOption) (*Buckets, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Buckets{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of buckets for a given project.",
	//   "httpMethod": "GET",
	//   "id": "storage.buckets.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "default": "1000",
	//       "description": "Maximum number of buckets to return in a single response. The service will use this parameter or 1,000 items, whichever is smaller.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A previously-returned page token representing part of the larger set of results to view.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "prefix": {
	//       "description": "Filter results to buckets whose names begin with this prefix.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "A valid API project identifier.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to noAcl.",
	//       "enum": [
	//         "full",
	//         "noAcl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit owner, acl and defaultObjectAcl properties."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b",
	//   "response": {
	//     "$ref": "Buckets"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BucketsListCall) Pages(ctx context.Context, f func(*Buckets) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "storage.buckets.lockRetentionPolicy":

type BucketsLockRetentionPolicyCall struct {
	s          *Service
	bucket     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// LockRetentionPolicy: Locks retention policy on a bucket.
func (r *BucketsService) LockRetentionPolicy(bucket string, ifMetagenerationMatch int64) *BucketsLockRetentionPolicyCall {
	c := &BucketsLockRetentionPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.urlParams_.Set("ifMetagenerationMatch", fmt.Sprint(ifMetagenerationMatch))
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *BucketsLockRetentionPolicyCall) UserProject(userProject string) *BucketsLockRetentionPolicyCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketsLockRetentionPolicyCall) Fields(s ...googleapi.Field) *BucketsLockRetentionPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BucketsLockRetentionPolicyCall) Context(ctx context.Context) *BucketsLockRetentionPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BucketsLockRetentionPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BucketsLockRetentionPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/lockRetentionPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.buckets.lockRetentionPolicy" call.
// Exactly one of *Bucket or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Bucket.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *BucketsLockRetentionPolicyCall) Do(opts ...googleapi.CallOption) (*Bucket, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Bucket{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Locks retention policy on a bucket.",
	//   "httpMethod": "POST",
	//   "id": "storage.buckets.lockRetentionPolicy",
	//   "parameterOrder": [
	//     "bucket",
	//     "ifMetagenerationMatch"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ifMetagenerationMatch": {
	//       "description": "Makes the operation conditional on whether bucket's current metageneration matches the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/lockRetentionPolicy",
	//   "response": {
	//     "$ref": "Bucket"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.buckets.patch":

type BucketsPatchCall struct {
	s          *Service
	bucket     string
	bucket2    *Bucket
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Updates a bucket. Changes to the bucket will be readable
// immediately after writing, but configuration changes may take time to
// propagate. This method supports patch semantics.
func (r *BucketsService) Patch(bucket string, bucket2 *Bucket) *BucketsPatchCall {
	c := &BucketsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.bucket2 = bucket2
	return c
}

// IfMetagenerationMatch sets the optional parameter
// "ifMetagenerationMatch": Makes the return of the bucket metadata
// conditional on whether the bucket's current metageneration matches
// the given value.
func (c *BucketsPatchCall) IfMetagenerationMatch(ifMetagenerationMatch int64) *BucketsPatchCall {
	c.urlParams_.Set("ifMetagenerationMatch", fmt.Sprint(ifMetagenerationMatch))
	return c
}

// IfMetagenerationNotMatch sets the optional parameter
// "ifMetagenerationNotMatch": Makes the return of the bucket metadata
// conditional on whether the bucket's current metageneration does not
// match the given value.
func (c *BucketsPatchCall) IfMetagenerationNotMatch(ifMetagenerationNotMatch int64) *BucketsPatchCall {
	c.urlParams_.Set("ifMetagenerationNotMatch", fmt.Sprint(ifMetagenerationNotMatch))
	return c
}

// PredefinedAcl sets the optional parameter "predefinedAcl": Apply a
// predefined set of access controls to this bucket.
//
// Possible values:
//   "authenticatedRead" - Project team owners get OWNER access, and
// allAuthenticatedUsers get READER access.
//   "private" - Project team owners get OWNER access.
//   "projectPrivate" - Project team members get access according to
// their roles.
//   "publicRead" - Project team owners get OWNER access, and allUsers
// get READER access.
//   "publicReadWrite" - Project team owners get OWNER access, and
// allUsers get WRITER access.
func (c *BucketsPatchCall) PredefinedAcl(predefinedAcl string) *BucketsPatchCall {
	c.urlParams_.Set("predefinedAcl", predefinedAcl)
	return c
}

// PredefinedDefaultObjectAcl sets the optional parameter
// "predefinedDefaultObjectAcl": Apply a predefined set of default
// object access controls to this bucket.
//
// Possible values:
//   "authenticatedRead" - Object owner gets OWNER access, and
// allAuthenticatedUsers get READER access.
//   "bucketOwnerFullControl" - Object owner gets OWNER access, and
// project team owners get OWNER access.
//   "bucketOwnerRead" - Object owner gets OWNER access, and project
// team owners get READER access.
//   "private" - Object owner gets OWNER access.
//   "projectPrivate" - Object owner gets OWNER access, and project team
// members get access according to their roles.
//   "publicRead" - Object owner gets OWNER access, and allUsers get
// READER access.
func (c *BucketsPatchCall) PredefinedDefaultObjectAcl(predefinedDefaultObjectAcl string) *BucketsPatchCall {
	c.urlParams_.Set("predefinedDefaultObjectAcl", predefinedDefaultObjectAcl)
	return c
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to full.
//
// Possible values:
//   "full" - Include all properties.
//   "noAcl" - Omit owner, acl and defaultObjectAcl properties.
func (c *BucketsPatchCall) Projection(projection string) *BucketsPatchCall {
	c.urlParams_.Set("projection", projection)
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *BucketsPatchCall) UserProject(userProject string) *BucketsPatchCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketsPatchCall) Fields(s ...googleapi.Field) *BucketsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BucketsPatchCall) Context(ctx context.Context) *BucketsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BucketsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BucketsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.bucket2)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.buckets.patch" call.
// Exactly one of *Bucket or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Bucket.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *BucketsPatchCall) Do(opts ...googleapi.CallOption) (*Bucket, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Bucket{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a bucket. Changes to the bucket will be readable immediately after writing, but configuration changes may take time to propagate. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "storage.buckets.patch",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ifMetagenerationMatch": {
	//       "description": "Makes the return of the bucket metadata conditional on whether the bucket's current metageneration matches the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationNotMatch": {
	//       "description": "Makes the return of the bucket metadata conditional on whether the bucket's current metageneration does not match the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "predefinedAcl": {
	//       "description": "Apply a predefined set of access controls to this bucket.",
	//       "enum": [
	//         "authenticatedRead",
	//         "private",
	//         "projectPrivate",
	//         "publicRead",
	//         "publicReadWrite"
	//       ],
	//       "enumDescriptions": [
	//         "Project team owners get OWNER access, and allAuthenticatedUsers get READER access.",
	//         "Project team owners get OWNER access.",
	//         "Project team members get access according to their roles.",
	//         "Project team owners get OWNER access, and allUsers get READER access.",
	//         "Project team owners get OWNER access, and allUsers get WRITER access."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "predefinedDefaultObjectAcl": {
	//       "description": "Apply a predefined set of default object access controls to this bucket.",
	//       "enum": [
	//         "authenticatedRead",
	//         "bucketOwnerFullControl",
	//         "bucketOwnerRead",
	//         "private",
	//         "projectPrivate",
	//         "publicRead"
	//       ],
	//       "enumDescriptions": [
	//         "Object owner gets OWNER access, and allAuthenticatedUsers get READER access.",
	//         "Object owner gets OWNER access, and project team owners get OWNER access.",
	//         "Object owner gets OWNER access, and project team owners get READER access.",
	//         "Object owner gets OWNER access.",
	//         "Object owner gets OWNER access, and project team members get access according to their roles.",
	//         "Object owner gets OWNER access, and allUsers get READER access."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to full.",
	//       "enum": [
	//         "full",
	//         "noAcl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit owner, acl and defaultObjectAcl properties."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}",
	//   "request": {
	//     "$ref": "Bucket"
	//   },
	//   "response": {
	//     "$ref": "Bucket"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.buckets.setIamPolicy":

type BucketsSetIamPolicyCall struct {
	s          *Service
	bucket     string
	policy     *Policy
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// SetIamPolicy: Updates an IAM policy for the specified bucket.
func (r *BucketsService) SetIamPolicy(bucket string, policy *Policy) *BucketsSetIamPolicyCall {
	c := &BucketsSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.policy = policy
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *BucketsSetIamPolicyCall) UserProject(userProject string) *BucketsSetIamPolicyCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketsSetIamPolicyCall) Fields(s ...googleapi.Field) *BucketsSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BucketsSetIamPolicyCall) Context(ctx context.Context) *BucketsSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BucketsSetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BucketsSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.policy)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/iam")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.buckets.setIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *BucketsSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an IAM policy for the specified bucket.",
	//   "httpMethod": "PUT",
	//   "id": "storage.buckets.setIamPolicy",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/iam",
	//   "request": {
	//     "$ref": "Policy"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.buckets.testIamPermissions":

type BucketsTestIamPermissionsCall struct {
	s            *Service
	bucket       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// TestIamPermissions: Tests a set of permissions on the given bucket to
// see which, if any, are held by the caller.
func (r *BucketsService) TestIamPermissions(bucket string, permissions []string) *BucketsTestIamPermissionsCall {
	c := &BucketsTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.urlParams_.SetMulti("permissions", append([]string{}, permissions...))
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *BucketsTestIamPermissionsCall) UserProject(userProject string) *BucketsTestIamPermissionsCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketsTestIamPermissionsCall) Fields(s ...googleapi.Field) *BucketsTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BucketsTestIamPermissionsCall) IfNoneMatch(entityTag string) *BucketsTestIamPermissionsCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BucketsTestIamPermissionsCall) Context(ctx context.Context) *BucketsTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BucketsTestIamPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BucketsTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/iam/testPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.buckets.testIamPermissions" call.
// Exactly one of *TestIamPermissionsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *TestIamPermissionsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BucketsTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*TestIamPermissionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestIamPermissionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Tests a set of permissions on the given bucket to see which, if any, are held by the caller.",
	//   "httpMethod": "GET",
	//   "id": "storage.buckets.testIamPermissions",
	//   "parameterOrder": [
	//     "bucket",
	//     "permissions"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "permissions": {
	//       "description": "Permissions to test.",
	//       "location": "query",
	//       "repeated": true,
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/iam/testPermissions",
	//   "response": {
	//     "$ref": "TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.buckets.update":

type BucketsUpdateCall struct {
	s          *Service
	bucket     string
	bucket2    *Bucket
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Updates a bucket. Changes to the bucket will be readable
// immediately after writing, but configuration changes may take time to
// propagate.
func (r *BucketsService) Update(bucket string, bucket2 *Bucket) *BucketsUpdateCall {
	c := &BucketsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.bucket2 = bucket2
	return c
}

// IfMetagenerationMatch sets the optional parameter
// "ifMetagenerationMatch": Makes the return of the bucket metadata
// conditional on whether the bucket's current metageneration matches
// the given value.
func (c *BucketsUpdateCall) IfMetagenerationMatch(ifMetagenerationMatch int64) *BucketsUpdateCall {
	c.urlParams_.Set("ifMetagenerationMatch", fmt.Sprint(ifMetagenerationMatch))
	return c
}

// IfMetagenerationNotMatch sets the optional parameter
// "ifMetagenerationNotMatch": Makes the return of the bucket metadata
// conditional on whether the bucket's current metageneration does not
// match the given value.
func (c *BucketsUpdateCall) IfMetagenerationNotMatch(ifMetagenerationNotMatch int64) *BucketsUpdateCall {
	c.urlParams_.Set("ifMetagenerationNotMatch", fmt.Sprint(ifMetagenerationNotMatch))
	return c
}

// PredefinedAcl sets the optional parameter "predefinedAcl": Apply a
// predefined set of access controls to this bucket.
//
// Possible values:
//   "authenticatedRead" - Project team owners get OWNER access, and
// allAuthenticatedUsers get READER access.
//   "private" - Project team owners get OWNER access.
//   "projectPrivate" - Project team members get access according to
// their roles.
//   "publicRead" - Project team owners get OWNER access, and allUsers
// get READER access.
//   "publicReadWrite" - Project team owners get OWNER access, and
// allUsers get WRITER access.
func (c *BucketsUpdateCall) PredefinedAcl(predefinedAcl string) *BucketsUpdateCall {
	c.urlParams_.Set("predefinedAcl", predefinedAcl)
	return c
}

// PredefinedDefaultObjectAcl sets the optional parameter
// "predefinedDefaultObjectAcl": Apply a predefined set of default
// object access controls to this bucket.
//
// Possible values:
//   "authenticatedRead" - Object owner gets OWNER access, and
// allAuthenticatedUsers get READER access.
//   "bucketOwnerFullControl" - Object owner gets OWNER access, and
// project team owners get OWNER access.
//   "bucketOwnerRead" - Object owner gets OWNER access, and project
// team owners get READER access.
//   "private" - Object owner gets OWNER access.
//   "projectPrivate" - Object owner gets OWNER access, and project team
// members get access according to their roles.
//   "publicRead" - Object owner gets OWNER access, and allUsers get
// READER access.
func (c *BucketsUpdateCall) PredefinedDefaultObjectAcl(predefinedDefaultObjectAcl string) *BucketsUpdateCall {
	c.urlParams_.Set("predefinedDefaultObjectAcl", predefinedDefaultObjectAcl)
	return c
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to full.
//
// Possible values:
//   "full" - Include all properties.
//   "noAcl" - Omit owner, acl and defaultObjectAcl properties.
func (c *BucketsUpdateCall) Projection(projection string) *BucketsUpdateCall {
	c.urlParams_.Set("projection", projection)
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *BucketsUpdateCall) UserProject(userProject string) *BucketsUpdateCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketsUpdateCall) Fields(s ...googleapi.Field) *BucketsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BucketsUpdateCall) Context(ctx context.Context) *BucketsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BucketsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BucketsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.bucket2)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.buckets.update" call.
// Exactly one of *Bucket or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Bucket.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *BucketsUpdateCall) Do(opts ...googleapi.CallOption) (*Bucket, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Bucket{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a bucket. Changes to the bucket will be readable immediately after writing, but configuration changes may take time to propagate.",
	//   "httpMethod": "PUT",
	//   "id": "storage.buckets.update",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ifMetagenerationMatch": {
	//       "description": "Makes the return of the bucket metadata conditional on whether the bucket's current metageneration matches the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationNotMatch": {
	//       "description": "Makes the return of the bucket metadata conditional on whether the bucket's current metageneration does not match the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "predefinedAcl": {
	//       "description": "Apply a predefined set of access controls to this bucket.",
	//       "enum": [
	//         "authenticatedRead",
	//         "private",
	//         "projectPrivate",
	//         "publicRead",
	//         "publicReadWrite"
	//       ],
	//       "enumDescriptions": [
	//         "Project team owners get OWNER access, and allAuthenticatedUsers get READER access.",
	//         "Project team owners get OWNER access.",
	//         "Project team members get access according to their roles.",
	//         "Project team owners get OWNER access, and allUsers get READER access.",
	//         "Project team owners get OWNER access, and allUsers get WRITER access."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "predefinedDefaultObjectAcl": {
	//       "description": "Apply a predefined set of default object access controls to this bucket.",
	//       "enum": [
	//         "authenticatedRead",
	//         "bucketOwnerFullControl",
	//         "bucketOwnerRead",
	//         "private",
	//         "projectPrivate",
	//         "publicRead"
	//       ],
	//       "enumDescriptions": [
	//         "Object owner gets OWNER access, and allAuthenticatedUsers get READER access.",
	//         "Object owner gets OWNER access, and project team owners get OWNER access.",
	//         "Object owner gets OWNER access, and project team owners get READER access.",
	//         "Object owner gets OWNER access.",
	//         "Object owner gets OWNER access, and project team members get access according to their roles.",
	//         "Object owner gets OWNER access, and allUsers get READER access."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to full.",
	//       "enum": [
	//         "full",
	//         "noAcl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit owner, acl and defaultObjectAcl properties."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}",
	//   "request": {
	//     "$ref": "Bucket"
	//   },
	//   "response": {
	//     "$ref": "Bucket"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.channels.stop":

type ChannelsStopCall struct {
	s          *Service
	channel    *Channel
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Stop: Stop watching resources through this channel
func (r *ChannelsService) Stop(channel *Channel) *ChannelsStopCall {
	c := &ChannelsStopCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.channel = channel
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChannelsStopCall) Fields(s ...googleapi.Field) *ChannelsStopCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ChannelsStopCall) Context(ctx context.Context) *ChannelsStopCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ChannelsStopCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ChannelsStopCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.channel)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "channels/stop")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.channels.stop" call.
func (c *ChannelsStopCall) Do(opts ...googleapi.CallOption) error {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Stop watching resources through this channel",
	//   "httpMethod": "POST",
	//   "id": "storage.channels.stop",
	//   "path": "channels/stop",
	//   "request": {
	//     "$ref": "Channel",
	//     "parameterName": "resource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.defaultObjectAccessControls.delete":

type DefaultObjectAccessControlsDeleteCall struct {
	s          *Service
	bucket     string
	entity     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Permanently deletes the default object ACL entry for the
// specified entity on the specified bucket.
func (r *DefaultObjectAccessControlsService) Delete(bucket string, entity string) *DefaultObjectAccessControlsDeleteCall {
	c := &DefaultObjectAccessControlsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.entity = entity
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *DefaultObjectAccessControlsDeleteCall) UserProject(userProject string) *DefaultObjectAccessControlsDeleteCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DefaultObjectAccessControlsDeleteCall) Fields(s ...googleapi.Field) *DefaultObjectAccessControlsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DefaultObjectAccessControlsDeleteCall) Context(ctx context.Context) *DefaultObjectAccessControlsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DefaultObjectAccessControlsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DefaultObjectAccessControlsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/defaultObjectAcl/{entity}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"entity": c.entity,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.defaultObjectAccessControls.delete" call.
func (c *DefaultObjectAccessControlsDeleteCall) Do(opts ...googleapi.CallOption) error {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Permanently deletes the default object ACL entry for the specified entity on the specified bucket.",
	//   "httpMethod": "DELETE",
	//   "id": "storage.defaultObjectAccessControls.delete",
	//   "parameterOrder": [
	//     "bucket",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/defaultObjectAcl/{entity}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.defaultObjectAccessControls.get":

type DefaultObjectAccessControlsGetCall struct {
	s            *Service
	bucket       string
	entity       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Returns the default object ACL entry for the specified entity on
// the specified bucket.
func (r *DefaultObjectAccessControlsService) Get(bucket string, entity string) *DefaultObjectAccessControlsGetCall {
	c := &DefaultObjectAccessControlsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.entity = entity
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *DefaultObjectAccessControlsGetCall) UserProject(userProject string) *DefaultObjectAccessControlsGetCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DefaultObjectAccessControlsGetCall) Fields(s ...googleapi.Field) *DefaultObjectAccessControlsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *DefaultObjectAccessControlsGetCall) IfNoneMatch(entityTag string) *DefaultObjectAccessControlsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DefaultObjectAccessControlsGetCall) Context(ctx context.Context) *DefaultObjectAccessControlsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DefaultObjectAccessControlsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DefaultObjectAccessControlsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/defaultObjectAcl/{entity}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"entity": c.entity,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.defaultObjectAccessControls.get" call.
// Exactly one of *ObjectAccessControl or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ObjectAccessControl.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *DefaultObjectAccessControlsGetCall) Do(opts ...googleapi.CallOption) (*ObjectAccessControl, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ObjectAccessControl{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the default object ACL entry for the specified entity on the specified bucket.",
	//   "httpMethod": "GET",
	//   "id": "storage.defaultObjectAccessControls.get",
	//   "parameterOrder": [
	//     "bucket",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/defaultObjectAcl/{entity}",
	//   "response": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.defaultObjectAccessControls.insert":

type DefaultObjectAccessControlsInsertCall struct {
	s                   *Service
	bucket              string
	objectaccesscontrol *ObjectAccessControl
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Insert: Creates a new default object ACL entry on the specified
// bucket.
func (r *DefaultObjectAccessControlsService) Insert(bucket string, objectaccesscontrol *ObjectAccessControl) *DefaultObjectAccessControlsInsertCall {
	c := &DefaultObjectAccessControlsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.objectaccesscontrol = objectaccesscontrol
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *DefaultObjectAccessControlsInsertCall) UserProject(userProject string) *DefaultObjectAccessControlsInsertCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DefaultObjectAccessControlsInsertCall) Fields(s ...googleapi.Field) *DefaultObjectAccessControlsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DefaultObjectAccessControlsInsertCall) Context(ctx context.Context) *DefaultObjectAccessControlsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DefaultObjectAccessControlsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DefaultObjectAccessControlsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.objectaccesscontrol)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/defaultObjectAcl")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.defaultObjectAccessControls.insert" call.
// Exactly one of *ObjectAccessControl or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ObjectAccessControl.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *DefaultObjectAccessControlsInsertCall) Do(opts ...googleapi.CallOption) (*ObjectAccessControl, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ObjectAccessControl{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new default object ACL entry on the specified bucket.",
	//   "httpMethod": "POST",
	//   "id": "storage.defaultObjectAccessControls.insert",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/defaultObjectAcl",
	//   "request": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "response": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.defaultObjectAccessControls.list":

type DefaultObjectAccessControlsListCall struct {
	s            *Service
	bucket       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves default object ACL entries on the specified bucket.
func (r *DefaultObjectAccessControlsService) List(bucket string) *DefaultObjectAccessControlsListCall {
	c := &DefaultObjectAccessControlsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	return c
}

// IfMetagenerationMatch sets the optional parameter
// "ifMetagenerationMatch": If present, only return default ACL listing
// if the bucket's current metageneration matches this value.
func (c *DefaultObjectAccessControlsListCall) IfMetagenerationMatch(ifMetagenerationMatch int64) *DefaultObjectAccessControlsListCall {
	c.urlParams_.Set("ifMetagenerationMatch", fmt.Sprint(ifMetagenerationMatch))
	return c
}

// IfMetagenerationNotMatch sets the optional parameter
// "ifMetagenerationNotMatch": If present, only return default ACL
// listing if the bucket's current metageneration does not match the
// given value.
func (c *DefaultObjectAccessControlsListCall) IfMetagenerationNotMatch(ifMetagenerationNotMatch int64) *DefaultObjectAccessControlsListCall {
	c.urlParams_.Set("ifMetagenerationNotMatch", fmt.Sprint(ifMetagenerationNotMatch))
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *DefaultObjectAccessControlsListCall) UserProject(userProject string) *DefaultObjectAccessControlsListCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DefaultObjectAccessControlsListCall) Fields(s ...googleapi.Field) *DefaultObjectAccessControlsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *DefaultObjectAccessControlsListCall) IfNoneMatch(entityTag string) *DefaultObjectAccessControlsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DefaultObjectAccessControlsListCall) Context(ctx context.Context) *DefaultObjectAccessControlsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DefaultObjectAccessControlsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DefaultObjectAccessControlsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/defaultObjectAcl")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.defaultObjectAccessControls.list" call.
// Exactly one of *ObjectAccessControls or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ObjectAccessControls.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *DefaultObjectAccessControlsListCall) Do(opts ...googleapi.CallOption) (*ObjectAccessControls, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ObjectAccessControls{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves default object ACL entries on the specified bucket.",
	//   "httpMethod": "GET",
	//   "id": "storage.defaultObjectAccessControls.list",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ifMetagenerationMatch": {
	//       "description": "If present, only return default ACL listing if the bucket's current metageneration matches this value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationNotMatch": {
	//       "description": "If present, only return default ACL listing if the bucket's current metageneration does not match the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/defaultObjectAcl",
	//   "response": {
	//     "$ref": "ObjectAccessControls"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.defaultObjectAccessControls.patch":

type DefaultObjectAccessControlsPatchCall struct {
	s                   *Service
	bucket              string
	entity              string
	objectaccesscontrol *ObjectAccessControl
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Patch: Updates a default object ACL entry on the specified bucket.
// This method supports patch semantics.
func (r *DefaultObjectAccessControlsService) Patch(bucket string, entity string, objectaccesscontrol *ObjectAccessControl) *DefaultObjectAccessControlsPatchCall {
	c := &DefaultObjectAccessControlsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.entity = entity
	c.objectaccesscontrol = objectaccesscontrol
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *DefaultObjectAccessControlsPatchCall) UserProject(userProject string) *DefaultObjectAccessControlsPatchCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DefaultObjectAccessControlsPatchCall) Fields(s ...googleapi.Field) *DefaultObjectAccessControlsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DefaultObjectAccessControlsPatchCall) Context(ctx context.Context) *DefaultObjectAccessControlsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DefaultObjectAccessControlsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DefaultObjectAccessControlsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.objectaccesscontrol)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/defaultObjectAcl/{entity}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"entity": c.entity,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.defaultObjectAccessControls.patch" call.
// Exactly one of *ObjectAccessControl or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ObjectAccessControl.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *DefaultObjectAccessControlsPatchCall) Do(opts ...googleapi.CallOption) (*ObjectAccessControl, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ObjectAccessControl{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a default object ACL entry on the specified bucket. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "storage.defaultObjectAccessControls.patch",
	//   "parameterOrder": [
	//     "bucket",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/defaultObjectAcl/{entity}",
	//   "request": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "response": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.defaultObjectAccessControls.update":

type DefaultObjectAccessControlsUpdateCall struct {
	s                   *Service
	bucket              string
	entity              string
	objectaccesscontrol *ObjectAccessControl
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Update: Updates a default object ACL entry on the specified bucket.
func (r *DefaultObjectAccessControlsService) Update(bucket string, entity string, objectaccesscontrol *ObjectAccessControl) *DefaultObjectAccessControlsUpdateCall {
	c := &DefaultObjectAccessControlsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.entity = entity
	c.objectaccesscontrol = objectaccesscontrol
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *DefaultObjectAccessControlsUpdateCall) UserProject(userProject string) *DefaultObjectAccessControlsUpdateCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DefaultObjectAccessControlsUpdateCall) Fields(s ...googleapi.Field) *DefaultObjectAccessControlsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DefaultObjectAccessControlsUpdateCall) Context(ctx context.Context) *DefaultObjectAccessControlsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DefaultObjectAccessControlsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DefaultObjectAccessControlsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.objectaccesscontrol)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/defaultObjectAcl/{entity}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"entity": c.entity,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.defaultObjectAccessControls.update" call.
// Exactly one of *ObjectAccessControl or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ObjectAccessControl.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *DefaultObjectAccessControlsUpdateCall) Do(opts ...googleapi.CallOption) (*ObjectAccessControl, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ObjectAccessControl{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a default object ACL entry on the specified bucket.",
	//   "httpMethod": "PUT",
	//   "id": "storage.defaultObjectAccessControls.update",
	//   "parameterOrder": [
	//     "bucket",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/defaultObjectAcl/{entity}",
	//   "request": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "response": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.notifications.delete":

type NotificationsDeleteCall struct {
	s            *Service
	bucket       string
	notification string
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Delete: Permanently deletes a notification subscription.
func (r *NotificationsService) Delete(bucket string, notification string) *NotificationsDeleteCall {
	c := &NotificationsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.notification = notification
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *NotificationsDeleteCall) UserProject(userProject string) *NotificationsDeleteCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *NotificationsDeleteCall) Fields(s ...googleapi.Field) *NotificationsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *NotificationsDeleteCall) Context(ctx context.Context) *NotificationsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *NotificationsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *NotificationsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/notificationConfigs/{notification}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket":       c.bucket,
		"notification": c.notification,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.notifications.delete" call.
func (c *NotificationsDeleteCall) Do(opts ...googleapi.CallOption) error {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Permanently deletes a notification subscription.",
	//   "httpMethod": "DELETE",
	//   "id": "storage.notifications.delete",
	//   "parameterOrder": [
	//     "bucket",
	//     "notification"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "The parent bucket of the notification.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "notification": {
	//       "description": "ID of the notification to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/notificationConfigs/{notification}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.notifications.get":

type NotificationsGetCall struct {
	s            *Service
	bucket       string
	notification string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: View a notification configuration.
func (r *NotificationsService) Get(bucket string, notification string) *NotificationsGetCall {
	c := &NotificationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.notification = notification
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *NotificationsGetCall) UserProject(userProject string) *NotificationsGetCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *NotificationsGetCall) Fields(s ...googleapi.Field) *NotificationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *NotificationsGetCall) IfNoneMatch(entityTag string) *NotificationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *NotificationsGetCall) Context(ctx context.Context) *NotificationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *NotificationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *NotificationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/notificationConfigs/{notification}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket":       c.bucket,
		"notification": c.notification,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.notifications.get" call.
// Exactly one of *Notification or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Notification.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *NotificationsGetCall) Do(opts ...googleapi.CallOption) (*Notification, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Notification{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "View a notification configuration.",
	//   "httpMethod": "GET",
	//   "id": "storage.notifications.get",
	//   "parameterOrder": [
	//     "bucket",
	//     "notification"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "The parent bucket of the notification.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "notification": {
	//       "description": "Notification ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/notificationConfigs/{notification}",
	//   "response": {
	//     "$ref": "Notification"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.notifications.insert":

type NotificationsInsertCall struct {
	s            *Service
	bucket       string
	notification *Notification
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Insert: Creates a notification subscription for a given bucket.
func (r *NotificationsService) Insert(bucket string, notification *Notification) *NotificationsInsertCall {
	c := &NotificationsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.notification = notification
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *NotificationsInsertCall) UserProject(userProject string) *NotificationsInsertCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *NotificationsInsertCall) Fields(s ...googleapi.Field) *NotificationsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *NotificationsInsertCall) Context(ctx context.Context) *NotificationsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *NotificationsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *NotificationsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.notification)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/notificationConfigs")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.notifications.insert" call.
// Exactly one of *Notification or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Notification.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *NotificationsInsertCall) Do(opts ...googleapi.CallOption) (*Notification, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Notification{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a notification subscription for a given bucket.",
	//   "httpMethod": "POST",
	//   "id": "storage.notifications.insert",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "The parent bucket of the notification.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/notificationConfigs",
	//   "request": {
	//     "$ref": "Notification"
	//   },
	//   "response": {
	//     "$ref": "Notification"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.notifications.list":

type NotificationsListCall struct {
	s            *Service
	bucket       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of notification subscriptions for a given
// bucket.
func (r *NotificationsService) List(bucket string) *NotificationsListCall {
	c := &NotificationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *NotificationsListCall) UserProject(userProject string) *NotificationsListCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *NotificationsListCall) Fields(s ...googleapi.Field) *NotificationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *NotificationsListCall) IfNoneMatch(entityTag string) *NotificationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *NotificationsListCall) Context(ctx context.Context) *NotificationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *NotificationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *NotificationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/notificationConfigs")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.notifications.list" call.
// Exactly one of *Notifications or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Notifications.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *NotificationsListCall) Do(opts ...googleapi.CallOption) (*Notifications, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Notifications{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of notification subscriptions for a given bucket.",
	//   "httpMethod": "GET",
	//   "id": "storage.notifications.list",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a Google Cloud Storage bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/notificationConfigs",
	//   "response": {
	//     "$ref": "Notifications"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.objectAccessControls.delete":

type ObjectAccessControlsDeleteCall struct {
	s          *Service
	bucket     string
	object     string
	entity     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Permanently deletes the ACL entry for the specified entity on
// the specified object.
func (r *ObjectAccessControlsService) Delete(bucket string, object string, entity string) *ObjectAccessControlsDeleteCall {
	c := &ObjectAccessControlsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.object = object
	c.entity = entity
	return c
}

// Generation sets the optional parameter "generation": If present,
// selects a specific revision of this object (as opposed to the latest
// version, the default).
func (c *ObjectAccessControlsDeleteCall) Generation(generation int64) *ObjectAccessControlsDeleteCall {
	c.urlParams_.Set("generation", fmt.Sprint(generation))
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectAccessControlsDeleteCall) UserProject(userProject string) *ObjectAccessControlsDeleteCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectAccessControlsDeleteCall) Fields(s ...googleapi.Field) *ObjectAccessControlsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectAccessControlsDeleteCall) Context(ctx context.Context) *ObjectAccessControlsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectAccessControlsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectAccessControlsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/o/{object}/acl/{entity}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
		"entity": c.entity,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objectAccessControls.delete" call.
func (c *ObjectAccessControlsDeleteCall) Do(opts ...googleapi.CallOption) error {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Permanently deletes the ACL entry for the specified entity on the specified object.",
	//   "httpMethod": "DELETE",
	//   "id": "storage.objectAccessControls.delete",
	//   "parameterOrder": [
	//     "bucket",
	//     "object",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "generation": {
	//       "description": "If present, selects a specific revision of this object (as opposed to the latest version, the default).",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}/acl/{entity}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.objectAccessControls.get":

type ObjectAccessControlsGetCall struct {
	s            *Service
	bucket       string
	object       string
	entity       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Returns the ACL entry for the specified entity on the specified
// object.
func (r *ObjectAccessControlsService) Get(bucket string, object string, entity string) *ObjectAccessControlsGetCall {
	c := &ObjectAccessControlsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.object = object
	c.entity = entity
	return c
}

// Generation sets the optional parameter "generation": If present,
// selects a specific revision of this object (as opposed to the latest
// version, the default).
func (c *ObjectAccessControlsGetCall) Generation(generation int64) *ObjectAccessControlsGetCall {
	c.urlParams_.Set("generation", fmt.Sprint(generation))
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectAccessControlsGetCall) UserProject(userProject string) *ObjectAccessControlsGetCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectAccessControlsGetCall) Fields(s ...googleapi.Field) *ObjectAccessControlsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ObjectAccessControlsGetCall) IfNoneMatch(entityTag string) *ObjectAccessControlsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectAccessControlsGetCall) Context(ctx context.Context) *ObjectAccessControlsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectAccessControlsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectAccessControlsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/o/{object}/acl/{entity}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
		"entity": c.entity,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objectAccessControls.get" call.
// Exactly one of *ObjectAccessControl or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ObjectAccessControl.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ObjectAccessControlsGetCall) Do(opts ...googleapi.CallOption) (*ObjectAccessControl, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ObjectAccessControl{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the ACL entry for the specified entity on the specified object.",
	//   "httpMethod": "GET",
	//   "id": "storage.objectAccessControls.get",
	//   "parameterOrder": [
	//     "bucket",
	//     "object",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "generation": {
	//       "description": "If present, selects a specific revision of this object (as opposed to the latest version, the default).",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}/acl/{entity}",
	//   "response": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.objectAccessControls.insert":

type ObjectAccessControlsInsertCall struct {
	s                   *Service
	bucket              string
	object              string
	objectaccesscontrol *ObjectAccessControl
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Insert: Creates a new ACL entry on the specified object.
func (r *ObjectAccessControlsService) Insert(bucket string, object string, objectaccesscontrol *ObjectAccessControl) *ObjectAccessControlsInsertCall {
	c := &ObjectAccessControlsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.object = object
	c.objectaccesscontrol = objectaccesscontrol
	return c
}

// Generation sets the optional parameter "generation": If present,
// selects a specific revision of this object (as opposed to the latest
// version, the default).
func (c *ObjectAccessControlsInsertCall) Generation(generation int64) *ObjectAccessControlsInsertCall {
	c.urlParams_.Set("generation", fmt.Sprint(generation))
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectAccessControlsInsertCall) UserProject(userProject string) *ObjectAccessControlsInsertCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectAccessControlsInsertCall) Fields(s ...googleapi.Field) *ObjectAccessControlsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectAccessControlsInsertCall) Context(ctx context.Context) *ObjectAccessControlsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectAccessControlsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectAccessControlsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.objectaccesscontrol)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/o/{object}/acl")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objectAccessControls.insert" call.
// Exactly one of *ObjectAccessControl or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ObjectAccessControl.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ObjectAccessControlsInsertCall) Do(opts ...googleapi.CallOption) (*ObjectAccessControl, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ObjectAccessControl{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new ACL entry on the specified object.",
	//   "httpMethod": "POST",
	//   "id": "storage.objectAccessControls.insert",
	//   "parameterOrder": [
	//     "bucket",
	//     "object"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "generation": {
	//       "description": "If present, selects a specific revision of this object (as opposed to the latest version, the default).",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}/acl",
	//   "request": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "response": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.objectAccessControls.list":

type ObjectAccessControlsListCall struct {
	s            *Service
	bucket       string
	object       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves ACL entries on the specified object.
func (r *ObjectAccessControlsService) List(bucket string, object string) *ObjectAccessControlsListCall {
	c := &ObjectAccessControlsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.object = object
	return c
}

// Generation sets the optional parameter "generation": If present,
// selects a specific revision of this object (as opposed to the latest
// version, the default).
func (c *ObjectAccessControlsListCall) Generation(generation int64) *ObjectAccessControlsListCall {
	c.urlParams_.Set("generation", fmt.Sprint(generation))
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectAccessControlsListCall) UserProject(userProject string) *ObjectAccessControlsListCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectAccessControlsListCall) Fields(s ...googleapi.Field) *ObjectAccessControlsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ObjectAccessControlsListCall) IfNoneMatch(entityTag string) *ObjectAccessControlsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectAccessControlsListCall) Context(ctx context.Context) *ObjectAccessControlsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectAccessControlsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectAccessControlsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/o/{object}/acl")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objectAccessControls.list" call.
// Exactly one of *ObjectAccessControls or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ObjectAccessControls.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ObjectAccessControlsListCall) Do(opts ...googleapi.CallOption) (*ObjectAccessControls, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ObjectAccessControls{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves ACL entries on the specified object.",
	//   "httpMethod": "GET",
	//   "id": "storage.objectAccessControls.list",
	//   "parameterOrder": [
	//     "bucket",
	//     "object"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "generation": {
	//       "description": "If present, selects a specific revision of this object (as opposed to the latest version, the default).",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}/acl",
	//   "response": {
	//     "$ref": "ObjectAccessControls"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.objectAccessControls.patch":

type ObjectAccessControlsPatchCall struct {
	s                   *Service
	bucket              string
	object              string
	entity              string
	objectaccesscontrol *ObjectAccessControl
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Patch: Updates an ACL entry on the specified object. This method
// supports patch semantics.
func (r *ObjectAccessControlsService) Patch(bucket string, object string, entity string, objectaccesscontrol *ObjectAccessControl) *ObjectAccessControlsPatchCall {
	c := &ObjectAccessControlsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.object = object
	c.entity = entity
	c.objectaccesscontrol = objectaccesscontrol
	return c
}

// Generation sets the optional parameter "generation": If present,
// selects a specific revision of this object (as opposed to the latest
// version, the default).
func (c *ObjectAccessControlsPatchCall) Generation(generation int64) *ObjectAccessControlsPatchCall {
	c.urlParams_.Set("generation", fmt.Sprint(generation))
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectAccessControlsPatchCall) UserProject(userProject string) *ObjectAccessControlsPatchCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectAccessControlsPatchCall) Fields(s ...googleapi.Field) *ObjectAccessControlsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectAccessControlsPatchCall) Context(ctx context.Context) *ObjectAccessControlsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectAccessControlsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectAccessControlsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.objectaccesscontrol)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/o/{object}/acl/{entity}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
		"entity": c.entity,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objectAccessControls.patch" call.
// Exactly one of *ObjectAccessControl or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ObjectAccessControl.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ObjectAccessControlsPatchCall) Do(opts ...googleapi.CallOption) (*ObjectAccessControl, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ObjectAccessControl{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an ACL entry on the specified object. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "storage.objectAccessControls.patch",
	//   "parameterOrder": [
	//     "bucket",
	//     "object",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "generation": {
	//       "description": "If present, selects a specific revision of this object (as opposed to the latest version, the default).",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}/acl/{entity}",
	//   "request": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "response": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.objectAccessControls.update":

type ObjectAccessControlsUpdateCall struct {
	s                   *Service
	bucket              string
	object              string
	entity              string
	objectaccesscontrol *ObjectAccessControl
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Update: Updates an ACL entry on the specified object.
func (r *ObjectAccessControlsService) Update(bucket string, object string, entity string, objectaccesscontrol *ObjectAccessControl) *ObjectAccessControlsUpdateCall {
	c := &ObjectAccessControlsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.object = object
	c.entity = entity
	c.objectaccesscontrol = objectaccesscontrol
	return c
}

// Generation sets the optional parameter "generation": If present,
// selects a specific revision of this object (as opposed to the latest
// version, the default).
func (c *ObjectAccessControlsUpdateCall) Generation(generation int64) *ObjectAccessControlsUpdateCall {
	c.urlParams_.Set("generation", fmt.Sprint(generation))
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectAccessControlsUpdateCall) UserProject(userProject string) *ObjectAccessControlsUpdateCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectAccessControlsUpdateCall) Fields(s ...googleapi.Field) *ObjectAccessControlsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectAccessControlsUpdateCall) Context(ctx context.Context) *ObjectAccessControlsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectAccessControlsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectAccessControlsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.objectaccesscontrol)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/o/{object}/acl/{entity}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
		"entity": c.entity,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objectAccessControls.update" call.
// Exactly one of *ObjectAccessControl or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ObjectAccessControl.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ObjectAccessControlsUpdateCall) Do(opts ...googleapi.CallOption) (*ObjectAccessControl, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ObjectAccessControl{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an ACL entry on the specified object.",
	//   "httpMethod": "PUT",
	//   "id": "storage.objectAccessControls.update",
	//   "parameterOrder": [
	//     "bucket",
	//     "object",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "generation": {
	//       "description": "If present, selects a specific revision of this object (as opposed to the latest version, the default).",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}/acl/{entity}",
	//   "request": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "response": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.objects.compose":

type ObjectsComposeCall struct {
	s                 *Service
	destinationBucket string
	destinationObject string
	composerequest    *ComposeRequest
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Compose: Concatenates a list of existing objects into a new object in
// the same bucket.
func (r *ObjectsService) Compose(destinationBucket string, destinationObject string, composerequest *ComposeRequest) *ObjectsComposeCall {
	c := &ObjectsComposeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.destinationBucket = destinationBucket
	c.destinationObject = destinationObject
	c.composerequest = composerequest
	return c
}

// DestinationPredefinedAcl sets the optional parameter
// "destinationPredefinedAcl": Apply a predefined set of access controls
// to the destination object.
//
// Possible values:
//   "authenticatedRead" - Object owner gets OWNER access, and
// allAuthenticatedUsers get READER access.
//   "bucketOwnerFullControl" - Object owner gets OWNER access, and
// project team owners get OWNER access.
//   "bucketOwnerRead" - Object owner gets OWNER access, and project
// team owners get READER access.
//   "private" - Object owner gets OWNER access.
//   "projectPrivate" - Object owner gets OWNER access, and project team
// members get access according to their roles.
//   "publicRead" - Object owner gets OWNER access, and allUsers get
// READER access.
func (c *ObjectsComposeCall) DestinationPredefinedAcl(destinationPredefinedAcl string) *ObjectsComposeCall {
	c.urlParams_.Set("destinationPredefinedAcl", destinationPredefinedAcl)
	return c
}

// IfGenerationMatch sets the optional parameter "ifGenerationMatch":
// Makes the operation conditional on whether the object's current
// generation matches the given value. Setting to 0 makes the operation
// succeed only if there are no live versions of the object.
func (c *ObjectsComposeCall) IfGenerationMatch(ifGenerationMatch int64) *ObjectsComposeCall {
	c.urlParams_.Set("ifGenerationMatch", fmt.Sprint(ifGenerationMatch))
	return c
}

// IfMetagenerationMatch sets the optional parameter
// "ifMetagenerationMatch": Makes the operation conditional on whether
// the object's current metageneration matches the given value.
func (c *ObjectsComposeCall) IfMetagenerationMatch(ifMetagenerationMatch int64) *ObjectsComposeCall {
	c.urlParams_.Set("ifMetagenerationMatch", fmt.Sprint(ifMetagenerationMatch))
	return c
}

// KmsKeyName sets the optional parameter "kmsKeyName": Resource name of
// the Cloud KMS key, of the form
// projects/my-project/locations/global/keyRings/my-kr/cryptoKeys/my-key,
//  that will be used to encrypt the object. Overrides the object
// metadata's kms_key_name value, if any.
func (c *ObjectsComposeCall) KmsKeyName(kmsKeyName string) *ObjectsComposeCall {
	c.urlParams_.Set("kmsKeyName", kmsKeyName)
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectsComposeCall) UserProject(userProject string) *ObjectsComposeCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsComposeCall) Fields(s ...googleapi.Field) *ObjectsComposeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectsComposeCall) Context(ctx context.Context) *ObjectsComposeCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectsComposeCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectsComposeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.composerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{destinationBucket}/o/{destinationObject}/compose")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"destinationBucket": c.destinationBucket,
		"destinationObject": c.destinationObject,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objects.compose" call.
// Exactly one of *Object or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Object.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ObjectsComposeCall) Do(opts ...googleapi.CallOption) (*Object, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Object{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Concatenates a list of existing objects into a new object in the same bucket.",
	//   "httpMethod": "POST",
	//   "id": "storage.objects.compose",
	//   "parameterOrder": [
	//     "destinationBucket",
	//     "destinationObject"
	//   ],
	//   "parameters": {
	//     "destinationBucket": {
	//       "description": "Name of the bucket in which to store the new object.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "destinationObject": {
	//       "description": "Name of the new object. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "destinationPredefinedAcl": {
	//       "description": "Apply a predefined set of access controls to the destination object.",
	//       "enum": [
	//         "authenticatedRead",
	//         "bucketOwnerFullControl",
	//         "bucketOwnerRead",
	//         "private",
	//         "projectPrivate",
	//         "publicRead"
	//       ],
	//       "enumDescriptions": [
	//         "Object owner gets OWNER access, and allAuthenticatedUsers get READER access.",
	//         "Object owner gets OWNER access, and project team owners get OWNER access.",
	//         "Object owner gets OWNER access, and project team owners get READER access.",
	//         "Object owner gets OWNER access.",
	//         "Object owner gets OWNER access, and project team members get access according to their roles.",
	//         "Object owner gets OWNER access, and allUsers get READER access."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifGenerationMatch": {
	//       "description": "Makes the operation conditional on whether the object's current generation matches the given value. Setting to 0 makes the operation succeed only if there are no live versions of the object.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationMatch": {
	//       "description": "Makes the operation conditional on whether the object's current metageneration matches the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "kmsKeyName": {
	//       "description": "Resource name of the Cloud KMS key, of the form projects/my-project/locations/global/keyRings/my-kr/cryptoKeys/my-key, that will be used to encrypt the object. Overrides the object metadata's kms_key_name value, if any.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{destinationBucket}/o/{destinationObject}/compose",
	//   "request": {
	//     "$ref": "ComposeRequest"
	//   },
	//   "response": {
	//     "$ref": "Object"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.objects.copy":

type ObjectsCopyCall struct {
	s                 *Service
	sourceBucket      string
	sourceObject      string
	destinationBucket string
	destinationObject string
	object            *Object
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Copy: Copies a source object to a destination object. Optionally
// overrides metadata.
func (r *ObjectsService) Copy(sourceBucket string, sourceObject string, destinationBucket string, destinationObject string, object *Object) *ObjectsCopyCall {
	c := &ObjectsCopyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.sourceBucket = sourceBucket
	c.sourceObject = sourceObject
	c.destinationBucket = destinationBucket
	c.destinationObject = destinationObject
	c.object = object
	return c
}

// DestinationPredefinedAcl sets the optional parameter
// "destinationPredefinedAcl": Apply a predefined set of access controls
// to the destination object.
//
// Possible values:
//   "authenticatedRead" - Object owner gets OWNER access, and
// allAuthenticatedUsers get READER access.
//   "bucketOwnerFullControl" - Object owner gets OWNER access, and
// project team owners get OWNER access.
//   "bucketOwnerRead" - Object owner gets OWNER access, and project
// team owners get READER access.
//   "private" - Object owner gets OWNER access.
//   "projectPrivate" - Object owner gets OWNER access, and project team
// members get access according to their roles.
//   "publicRead" - Object owner gets OWNER access, and allUsers get
// READER access.
func (c *ObjectsCopyCall) DestinationPredefinedAcl(destinationPredefinedAcl string) *ObjectsCopyCall {
	c.urlParams_.Set("destinationPredefinedAcl", destinationPredefinedAcl)
	return c
}

// IfGenerationMatch sets the optional parameter "ifGenerationMatch":
// Makes the operation conditional on whether the destination object's
// current generation matches the given value. Setting to 0 makes the
// operation succeed only if there are no live versions of the object.
func (c *ObjectsCopyCall) IfGenerationMatch(ifGenerationMatch int64) *ObjectsCopyCall {
	c.urlParams_.Set("ifGenerationMatch", fmt.Sprint(ifGenerationMatch))
	return c
}

// IfGenerationNotMatch sets the optional parameter
// "ifGenerationNotMatch": Makes the operation conditional on whether
// the destination object's current generation does not match the given
// value. If no live object exists, the precondition fails. Setting to 0
// makes the operation succeed only if there is a live version of the
// object.
func (c *ObjectsCopyCall) IfGenerationNotMatch(ifGenerationNotMatch int64) *ObjectsCopyCall {
	c.urlParams_.Set("ifGenerationNotMatch", fmt.Sprint(ifGenerationNotMatch))
	return c
}

// IfMetagenerationMatch sets the optional parameter
// "ifMetagenerationMatch": Makes the operation conditional on whether
// the destination object's current metageneration matches the given
// value.
func (c *ObjectsCopyCall) IfMetagenerationMatch(ifMetagenerationMatch int64) *ObjectsCopyCall {
	c.urlParams_.Set("ifMetagenerationMatch", fmt.Sprint(ifMetagenerationMatch))
	return c
}

// IfMetagenerationNotMatch sets the optional parameter
// "ifMetagenerationNotMatch": Makes the operation conditional on
// whether the destination object's current metageneration does not
// match the given value.
func (c *ObjectsCopyCall) IfMetagenerationNotMatch(ifMetagenerationNotMatch int64) *ObjectsCopyCall {
	c.urlParams_.Set("ifMetagenerationNotMatch", fmt.Sprint(ifMetagenerationNotMatch))
	return c
}

// IfSourceGenerationMatch sets the optional parameter
// "ifSourceGenerationMatch": Makes the operation conditional on whether
// the source object's current generation matches the given value.
func (c *ObjectsCopyCall) IfSourceGenerationMatch(ifSourceGenerationMatch int64) *ObjectsCopyCall {
	c.urlParams_.Set("ifSourceGenerationMatch", fmt.Sprint(ifSourceGenerationMatch))
	return c
}

// IfSourceGenerationNotMatch sets the optional parameter
// "ifSourceGenerationNotMatch": Makes the operation conditional on
// whether the source object's current generation does not match the
// given value.
func (c *ObjectsCopyCall) IfSourceGenerationNotMatch(ifSourceGenerationNotMatch int64) *ObjectsCopyCall {
	c.urlParams_.Set("ifSourceGenerationNotMatch", fmt.Sprint(ifSourceGenerationNotMatch))
	return c
}

// IfSourceMetagenerationMatch sets the optional parameter
// "ifSourceMetagenerationMatch": Makes the operation conditional on
// whether the source object's current metageneration matches the given
// value.
func (c *ObjectsCopyCall) IfSourceMetagenerationMatch(ifSourceMetagenerationMatch int64) *ObjectsCopyCall {
	c.urlParams_.Set("ifSourceMetagenerationMatch", fmt.Sprint(ifSourceMetagenerationMatch))
	return c
}

// IfSourceMetagenerationNotMatch sets the optional parameter
// "ifSourceMetagenerationNotMatch": Makes the operation conditional on
// whether the source object's current metageneration does not match the
// given value.
func (c *ObjectsCopyCall) IfSourceMetagenerationNotMatch(ifSourceMetagenerationNotMatch int64) *ObjectsCopyCall {
	c.urlParams_.Set("ifSourceMetagenerationNotMatch", fmt.Sprint(ifSourceMetagenerationNotMatch))
	return c
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to noAcl, unless the object resource
// specifies the acl property, when it defaults to full.
//
// Possible values:
//   "full" - Include all properties.
//   "noAcl" - Omit the owner, acl property.
func (c *ObjectsCopyCall) Projection(projection string) *ObjectsCopyCall {
	c.urlParams_.Set("projection", projection)
	return c
}

// SourceGeneration sets the optional parameter "sourceGeneration": If
// present, selects a specific revision of the source object (as opposed
// to the latest version, the default).
func (c *ObjectsCopyCall) SourceGeneration(sourceGeneration int64) *ObjectsCopyCall {
	c.urlParams_.Set("sourceGeneration", fmt.Sprint(sourceGeneration))
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectsCopyCall) UserProject(userProject string) *ObjectsCopyCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsCopyCall) Fields(s ...googleapi.Field) *ObjectsCopyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectsCopyCall) Context(ctx context.Context) *ObjectsCopyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectsCopyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectsCopyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.object)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{sourceBucket}/o/{sourceObject}/copyTo/b/{destinationBucket}/o/{destinationObject}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"sourceBucket":      c.sourceBucket,
		"sourceObject":      c.sourceObject,
		"destinationBucket": c.destinationBucket,
		"destinationObject": c.destinationObject,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objects.copy" call.
// Exactly one of *Object or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Object.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ObjectsCopyCall) Do(opts ...googleapi.CallOption) (*Object, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Object{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Copies a source object to a destination object. Optionally overrides metadata.",
	//   "httpMethod": "POST",
	//   "id": "storage.objects.copy",
	//   "parameterOrder": [
	//     "sourceBucket",
	//     "sourceObject",
	//     "destinationBucket",
	//     "destinationObject"
	//   ],
	//   "parameters": {
	//     "destinationBucket": {
	//       "description": "Name of the bucket in which to store the new object. Overrides the provided object metadata's bucket value, if any.For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "destinationObject": {
	//       "description": "Name of the new object. Required when the object metadata is not otherwise provided. Overrides the object metadata's name value, if any.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "destinationPredefinedAcl": {
	//       "description": "Apply a predefined set of access controls to the destination object.",
	//       "enum": [
	//         "authenticatedRead",
	//         "bucketOwnerFullControl",
	//         "bucketOwnerRead",
	//         "private",
	//         "projectPrivate",
	//         "publicRead"
	//       ],
	//       "enumDescriptions": [
	//         "Object owner gets OWNER access, and allAuthenticatedUsers get READER access.",
	//         "Object owner gets OWNER access, and project team owners get OWNER access.",
	//         "Object owner gets OWNER access, and project team owners get READER access.",
	//         "Object owner gets OWNER access.",
	//         "Object owner gets OWNER access, and project team members get access according to their roles.",
	//         "Object owner gets OWNER access, and allUsers get READER access."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifGenerationMatch": {
	//       "description": "Makes the operation conditional on whether the destination object's current generation matches the given value. Setting to 0 makes the operation succeed only if there are no live versions of the object.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifGenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the destination object's current generation does not match the given value. If no live object exists, the precondition fails. Setting to 0 makes the operation succeed only if there is a live version of the object.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationMatch": {
	//       "description": "Makes the operation conditional on whether the destination object's current metageneration matches the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the destination object's current metageneration does not match the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifSourceGenerationMatch": {
	//       "description": "Makes the operation conditional on whether the source object's current generation matches the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifSourceGenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the source object's current generation does not match the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifSourceMetagenerationMatch": {
	//       "description": "Makes the operation conditional on whether the source object's current metageneration matches the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifSourceMetagenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the source object's current metageneration does not match the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to noAcl, unless the object resource specifies the acl property, when it defaults to full.",
	//       "enum": [
	//         "full",
	//         "noAcl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit the owner, acl property."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sourceBucket": {
	//       "description": "Name of the bucket in which to find the source object.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sourceGeneration": {
	//       "description": "If present, selects a specific revision of the source object (as opposed to the latest version, the default).",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sourceObject": {
	//       "description": "Name of the source object. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{sourceBucket}/o/{sourceObject}/copyTo/b/{destinationBucket}/o/{destinationObject}",
	//   "request": {
	//     "$ref": "Object"
	//   },
	//   "response": {
	//     "$ref": "Object"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.objects.delete":

type ObjectsDeleteCall struct {
	s          *Service
	bucket     string
	object     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes an object and its metadata. Deletions are permanent
// if versioning is not enabled for the bucket, or if the generation
// parameter is used.
func (r *ObjectsService) Delete(bucket string, object string) *ObjectsDeleteCall {
	c := &ObjectsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.object = object
	return c
}

// Generation sets the optional parameter "generation": If present,
// permanently deletes a specific revision of this object (as opposed to
// the latest version, the default).
func (c *ObjectsDeleteCall) Generation(generation int64) *ObjectsDeleteCall {
	c.urlParams_.Set("generation", fmt.Sprint(generation))
	return c
}

// IfGenerationMatch sets the optional parameter "ifGenerationMatch":
// Makes the operation conditional on whether the object's current
// generation matches the given value. Setting to 0 makes the operation
// succeed only if there are no live versions of the object.
func (c *ObjectsDeleteCall) IfGenerationMatch(ifGenerationMatch int64) *ObjectsDeleteCall {
	c.urlParams_.Set("ifGenerationMatch", fmt.Sprint(ifGenerationMatch))
	return c
}

// IfGenerationNotMatch sets the optional parameter
// "ifGenerationNotMatch": Makes the operation conditional on whether
// the object's current generation does not match the given value. If no
// live object exists, the precondition fails. Setting to 0 makes the
// operation succeed only if there is a live version of the object.
func (c *ObjectsDeleteCall) IfGenerationNotMatch(ifGenerationNotMatch int64) *ObjectsDeleteCall {
	c.urlParams_.Set("ifGenerationNotMatch", fmt.Sprint(ifGenerationNotMatch))
	return c
}

// IfMetagenerationMatch sets the optional parameter
// "ifMetagenerationMatch": Makes the operation conditional on whether
// the object's current metageneration matches the given value.
func (c *ObjectsDeleteCall) IfMetagenerationMatch(ifMetagenerationMatch int64) *ObjectsDeleteCall {
	c.urlParams_.Set("ifMetagenerationMatch", fmt.Sprint(ifMetagenerationMatch))
	return c
}

// IfMetagenerationNotMatch sets the optional parameter
// "ifMetagenerationNotMatch": Makes the operation conditional on
// whether the object's current metageneration does not match the given
// value.
func (c *ObjectsDeleteCall) IfMetagenerationNotMatch(ifMetagenerationNotMatch int64) *ObjectsDeleteCall {
	c.urlParams_.Set("ifMetagenerationNotMatch", fmt.Sprint(ifMetagenerationNotMatch))
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectsDeleteCall) UserProject(userProject string) *ObjectsDeleteCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsDeleteCall) Fields(s ...googleapi.Field) *ObjectsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectsDeleteCall) Context(ctx context.Context) *ObjectsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/o/{object}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objects.delete" call.
func (c *ObjectsDeleteCall) Do(opts ...googleapi.CallOption) error {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes an object and its metadata. Deletions are permanent if versioning is not enabled for the bucket, or if the generation parameter is used.",
	//   "httpMethod": "DELETE",
	//   "id": "storage.objects.delete",
	//   "parameterOrder": [
	//     "bucket",
	//     "object"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of the bucket in which the object resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "generation": {
	//       "description": "If present, permanently deletes a specific revision of this object (as opposed to the latest version, the default).",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifGenerationMatch": {
	//       "description": "Makes the operation conditional on whether the object's current generation matches the given value. Setting to 0 makes the operation succeed only if there are no live versions of the object.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifGenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the object's current generation does not match the given value. If no live object exists, the precondition fails. Setting to 0 makes the operation succeed only if there is a live version of the object.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationMatch": {
	//       "description": "Makes the operation conditional on whether the object's current metageneration matches the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the object's current metageneration does not match the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.objects.get":

type ObjectsGetCall struct {
	s            *Service
	bucket       string
	object       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves an object or its metadata.
func (r *ObjectsService) Get(bucket string, object string) *ObjectsGetCall {
	c := &ObjectsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.object = object
	return c
}

// Generation sets the optional parameter "generation": If present,
// selects a specific revision of this object (as opposed to the latest
// version, the default).
func (c *ObjectsGetCall) Generation(generation int64) *ObjectsGetCall {
	c.urlParams_.Set("generation", fmt.Sprint(generation))
	return c
}

// IfGenerationMatch sets the optional parameter "ifGenerationMatch":
// Makes the operation conditional on whether the object's current
// generation matches the given value. Setting to 0 makes the operation
// succeed only if there are no live versions of the object.
func (c *ObjectsGetCall) IfGenerationMatch(ifGenerationMatch int64) *ObjectsGetCall {
	c.urlParams_.Set("ifGenerationMatch", fmt.Sprint(ifGenerationMatch))
	return c
}

// IfGenerationNotMatch sets the optional parameter
// "ifGenerationNotMatch": Makes the operation conditional on whether
// the object's current generation does not match the given value. If no
// live object exists, the precondition fails. Setting to 0 makes the
// operation succeed only if there is a live version of the object.
func (c *ObjectsGetCall) IfGenerationNotMatch(ifGenerationNotMatch int64) *ObjectsGetCall {
	c.urlParams_.Set("ifGenerationNotMatch", fmt.Sprint(ifGenerationNotMatch))
	return c
}

// IfMetagenerationMatch sets the optional parameter
// "ifMetagenerationMatch": Makes the operation conditional on whether
// the object's current metageneration matches the given value.
func (c *ObjectsGetCall) IfMetagenerationMatch(ifMetagenerationMatch int64) *ObjectsGetCall {
	c.urlParams_.Set("ifMetagenerationMatch", fmt.Sprint(ifMetagenerationMatch))
	return c
}

// IfMetagenerationNotMatch sets the optional parameter
// "ifMetagenerationNotMatch": Makes the operation conditional on
// whether the object's current metageneration does not match the given
// value.
func (c *ObjectsGetCall) IfMetagenerationNotMatch(ifMetagenerationNotMatch int64) *ObjectsGetCall {
	c.urlParams_.Set("ifMetagenerationNotMatch", fmt.Sprint(ifMetagenerationNotMatch))
	return c
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to noAcl.
//
// Possible values:
//   "full" - Include all properties.
//   "noAcl" - Omit the owner, acl property.
func (c *ObjectsGetCall) Projection(projection string) *ObjectsGetCall {
	c.urlParams_.Set("projection", projection)
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectsGetCall) UserProject(userProject string) *ObjectsGetCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsGetCall) Fields(s ...googleapi.Field) *ObjectsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ObjectsGetCall) IfNoneMatch(entityTag string) *ObjectsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do and Download
// methods. Any pending HTTP request will be aborted if the provided
// context is canceled.
func (c *ObjectsGetCall) Context(ctx context.Context) *ObjectsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/o/{object}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Download fetches the API endpoint's "media" value, instead of the normal
// API response value. If the returned error is nil, the Response is guaranteed to
// have a 2xx status code. Callers must close the Response.Body as usual.
func (c *ObjectsGetCall) Download(opts ...googleapi.CallOption) (*http.Response, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("media")
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckMediaResponse(res); err != nil {
		res.Body.Close()
		return nil, err
	}
	return res, nil
}

// Do executes the "storage.objects.get" call.
// Exactly one of *Object or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Object.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ObjectsGetCall) Do(opts ...googleapi.CallOption) (*Object, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Object{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves an object or its metadata.",
	//   "httpMethod": "GET",
	//   "id": "storage.objects.get",
	//   "parameterOrder": [
	//     "bucket",
	//     "object"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of the bucket in which the object resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "generation": {
	//       "description": "If present, selects a specific revision of this object (as opposed to the latest version, the default).",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifGenerationMatch": {
	//       "description": "Makes the operation conditional on whether the object's current generation matches the given value. Setting to 0 makes the operation succeed only if there are no live versions of the object.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifGenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the object's current generation does not match the given value. If no live object exists, the precondition fails. Setting to 0 makes the operation succeed only if there is a live version of the object.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationMatch": {
	//       "description": "Makes the operation conditional on whether the object's current metageneration matches the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the object's current metageneration does not match the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to noAcl.",
	//       "enum": [
	//         "full",
	//         "noAcl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit the owner, acl property."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}",
	//   "response": {
	//     "$ref": "Object"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ],
	//   "supportsMediaDownload": true,
	//   "useMediaDownloadService": true
	// }

}

// method id "storage.objects.getIamPolicy":

type ObjectsGetIamPolicyCall struct {
	s            *Service
	bucket       string
	object       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetIamPolicy: Returns an IAM policy for the specified object.
func (r *ObjectsService) GetIamPolicy(bucket string, object string) *ObjectsGetIamPolicyCall {
	c := &ObjectsGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.object = object
	return c
}

// Generation sets the optional parameter "generation": If present,
// selects a specific revision of this object (as opposed to the latest
// version, the default).
func (c *ObjectsGetIamPolicyCall) Generation(generation int64) *ObjectsGetIamPolicyCall {
	c.urlParams_.Set("generation", fmt.Sprint(generation))
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectsGetIamPolicyCall) UserProject(userProject string) *ObjectsGetIamPolicyCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsGetIamPolicyCall) Fields(s ...googleapi.Field) *ObjectsGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ObjectsGetIamPolicyCall) IfNoneMatch(entityTag string) *ObjectsGetIamPolicyCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectsGetIamPolicyCall) Context(ctx context.Context) *ObjectsGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectsGetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectsGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/o/{object}/iam")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objects.getIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ObjectsGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns an IAM policy for the specified object.",
	//   "httpMethod": "GET",
	//   "id": "storage.objects.getIamPolicy",
	//   "parameterOrder": [
	//     "bucket",
	//     "object"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of the bucket in which the object resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "generation": {
	//       "description": "If present, selects a specific revision of this object (as opposed to the latest version, the default).",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}/iam",
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.objects.insert":

type ObjectsInsertCall struct {
	s          *Service
	bucket     string
	object     *Object
	urlParams_ gensupport.URLParams
	mediaInfo_ *gensupport.MediaInfo
	ctx_       context.Context
	header_    http.Header
}

// Insert: Stores a new object and metadata.
func (r *ObjectsService) Insert(bucket string, object *Object) *ObjectsInsertCall {
	c := &ObjectsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.object = object
	return c
}

// ContentEncoding sets the optional parameter "contentEncoding": If
// set, sets the contentEncoding property of the final object to this
// value. Setting this parameter is equivalent to setting the
// contentEncoding metadata property. This can be useful when uploading
// an object with uploadType=media to indicate the encoding of the
// content being uploaded.
func (c *ObjectsInsertCall) ContentEncoding(contentEncoding string) *ObjectsInsertCall {
	c.urlParams_.Set("contentEncoding", contentEncoding)
	return c
}

// IfGenerationMatch sets the optional parameter "ifGenerationMatch":
// Makes the operation conditional on whether the object's current
// generation matches the given value. Setting to 0 makes the operation
// succeed only if there are no live versions of the object.
func (c *ObjectsInsertCall) IfGenerationMatch(ifGenerationMatch int64) *ObjectsInsertCall {
	c.urlParams_.Set("ifGenerationMatch", fmt.Sprint(ifGenerationMatch))
	return c
}

// IfGenerationNotMatch sets the optional parameter
// "ifGenerationNotMatch": Makes the operation conditional on whether
// the object's current generation does not match the given value. If no
// live object exists, the precondition fails. Setting to 0 makes the
// operation succeed only if there is a live version of the object.
func (c *ObjectsInsertCall) IfGenerationNotMatch(ifGenerationNotMatch int64) *ObjectsInsertCall {
	c.urlParams_.Set("ifGenerationNotMatch", fmt.Sprint(ifGenerationNotMatch))
	return c
}

// IfMetagenerationMatch sets the optional parameter
// "ifMetagenerationMatch": Makes the operation conditional on whether
// the object's current metageneration matches the given value.
func (c *ObjectsInsertCall) IfMetagenerationMatch(ifMetagenerationMatch int64) *ObjectsInsertCall {
	c.urlParams_.Set("ifMetagenerationMatch", fmt.Sprint(ifMetagenerationMatch))
	return c
}

// IfMetagenerationNotMatch sets the optional parameter
// "ifMetagenerationNotMatch": Makes the operation conditional on
// whether the object's current metageneration does not match the given
// value.
func (c *ObjectsInsertCall) IfMetagenerationNotMatch(ifMetagenerationNotMatch int64) *ObjectsInsertCall {
	c.urlParams_.Set("ifMetagenerationNotMatch", fmt.Sprint(ifMetagenerationNotMatch))
	return c
}

// KmsKeyName sets the optional parameter "kmsKeyName": Resource name of
// the Cloud KMS key, of the form
// projects/my-project/locations/global/keyRings/my-kr/cryptoKeys/my-key,
//  that will be used to encrypt the object. Overrides the object
// metadata's kms_key_name value, if any. Limited availability; usable
// only by enabled projects.
func (c *ObjectsInsertCall) KmsKeyName(kmsKeyName string) *ObjectsInsertCall {
	c.urlParams_.Set("kmsKeyName", kmsKeyName)
	return c
}

// Name sets the optional parameter "name": Name of the object. Required
// when the object metadata is not otherwise provided. Overrides the
// object metadata's name value, if any. For information about how to
// URL encode object names to be path safe, see Encoding URI Path Parts.
func (c *ObjectsInsertCall) Name(name string) *ObjectsInsertCall {
	c.urlParams_.Set("name", name)
	return c
}

// PredefinedAcl sets the optional parameter "predefinedAcl": Apply a
// predefined set of access controls to this object.
//
// Possible values:
//   "authenticatedRead" - Object owner gets OWNER access, and
// allAuthenticatedUsers get READER access.
//   "bucketOwnerFullControl" - Object owner gets OWNER access, and
// project team owners get OWNER access.
//   "bucketOwnerRead" - Object owner gets OWNER access, and project
// team owners get READER access.
//   "private" - Object owner gets OWNER access.
//   "projectPrivate" - Object owner gets OWNER access, and project team
// members get access according to their roles.
//   "publicRead" - Object owner gets OWNER access, and allUsers get
// READER access.
func (c *ObjectsInsertCall) PredefinedAcl(predefinedAcl string) *ObjectsInsertCall {
	c.urlParams_.Set("predefinedAcl", predefinedAcl)
	return c
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to noAcl, unless the object resource
// specifies the acl property, when it defaults to full.
//
// Possible values:
//   "full" - Include all properties.
//   "noAcl" - Omit the owner, acl property.
func (c *ObjectsInsertCall) Projection(projection string) *ObjectsInsertCall {
	c.urlParams_.Set("projection", projection)
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectsInsertCall) UserProject(userProject string) *ObjectsInsertCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Media specifies the media to upload in one or more chunks. The chunk
// size may be controlled by supplying a MediaOption generated by
// googleapi.ChunkSize. The chunk size defaults to
// googleapi.DefaultUploadChunkSize.The Content-Type header used in the
// upload request will be determined by sniffing the contents of r,
// unless a MediaOption generated by googleapi.ContentType is
// supplied.
// At most one of Media and ResumableMedia may be set.
func (c *ObjectsInsertCall) Media(r io.Reader, options ...googleapi.MediaOption) *ObjectsInsertCall {
	if ct := c.object.ContentType; ct != "" {
		options = append([]googleapi.MediaOption{googleapi.ContentType(ct)}, options...)
	}
	c.mediaInfo_ = gensupport.NewInfoFromMedia(r, options)
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be
// canceled with ctx.
//
// Deprecated: use Media instead.
//
// At most one of Media and ResumableMedia may be set. mediaType
// identifies the MIME media type of the upload, such as "image/png". If
// mediaType is "", it will be auto-detected. The provided ctx will
// supersede any context previously provided to the Context method.
func (c *ObjectsInsertCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *ObjectsInsertCall {
	c.ctx_ = ctx
	c.mediaInfo_ = gensupport.NewInfoFromResumableMedia(r, size, mediaType)
	return c
}

// ProgressUpdater provides a callback function that will be called
// after every chunk. It should be a low-latency function in order to
// not slow down the upload operation. This should only be called when
// using ResumableMedia (as opposed to Media).
func (c *ObjectsInsertCall) ProgressUpdater(pu googleapi.ProgressUpdater) *ObjectsInsertCall {
	c.mediaInfo_.SetProgressUpdater(pu)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsInsertCall) Fields(s ...googleapi.Field) *ObjectsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
// This context will supersede any context previously provided to the
// ResumableMedia method.
func (c *ObjectsInsertCall) Context(ctx context.Context) *ObjectsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.object)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/o")
	if c.mediaInfo_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
		c.urlParams_.Set("uploadType", c.mediaInfo_.UploadType())
	}
	if body == nil {
		body = new(bytes.Buffer)
		reqHeaders.Set("Content-Type", "application/json")
	}
	body, cleanup := c.mediaInfo_.UploadRequest(reqHeaders, body)
	defer cleanup()
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objects.insert" call.
// Exactly one of *Object or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Object.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ObjectsInsertCall) Do(opts ...googleapi.CallOption) (*Object, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	rx := c.mediaInfo_.ResumableUpload(res.Header.Get("Location"))
	if rx != nil {
		rx.Client = c.s.client
		rx.UserAgent = c.s.userAgent()
		ctx := c.ctx_
		if ctx == nil {
			ctx = context.TODO()
		}
		res, err = rx.Upload(ctx)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		if err := googleapi.CheckResponse(res); err != nil {
			return nil, err
		}
	}
	ret := &Object{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Stores a new object and metadata.",
	//   "httpMethod": "POST",
	//   "id": "storage.objects.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "*/*"
	//     ],
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/storage/v1/b/{bucket}/o"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/storage/v1/b/{bucket}/o"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of the bucket in which to store the new object. Overrides the provided object metadata's bucket value, if any.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "contentEncoding": {
	//       "description": "If set, sets the contentEncoding property of the final object to this value. Setting this parameter is equivalent to setting the contentEncoding metadata property. This can be useful when uploading an object with uploadType=media to indicate the encoding of the content being uploaded.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifGenerationMatch": {
	//       "description": "Makes the operation conditional on whether the object's current generation matches the given value. Setting to 0 makes the operation succeed only if there are no live versions of the object.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifGenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the object's current generation does not match the given value. If no live object exists, the precondition fails. Setting to 0 makes the operation succeed only if there is a live version of the object.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationMatch": {
	//       "description": "Makes the operation conditional on whether the object's current metageneration matches the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the object's current metageneration does not match the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "kmsKeyName": {
	//       "description": "Resource name of the Cloud KMS key, of the form projects/my-project/locations/global/keyRings/my-kr/cryptoKeys/my-key, that will be used to encrypt the object. Overrides the object metadata's kms_key_name value, if any. Limited availability; usable only by enabled projects.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "Name of the object. Required when the object metadata is not otherwise provided. Overrides the object metadata's name value, if any. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "predefinedAcl": {
	//       "description": "Apply a predefined set of access controls to this object.",
	//       "enum": [
	//         "authenticatedRead",
	//         "bucketOwnerFullControl",
	//         "bucketOwnerRead",
	//         "private",
	//         "projectPrivate",
	//         "publicRead"
	//       ],
	//       "enumDescriptions": [
	//         "Object owner gets OWNER access, and allAuthenticatedUsers get READER access.",
	//         "Object owner gets OWNER access, and project team owners get OWNER access.",
	//         "Object owner gets OWNER access, and project team owners get READER access.",
	//         "Object owner gets OWNER access.",
	//         "Object owner gets OWNER access, and project team members get access according to their roles.",
	//         "Object owner gets OWNER access, and allUsers get READER access."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to noAcl, unless the object resource specifies the acl property, when it defaults to full.",
	//       "enum": [
	//         "full",
	//         "noAcl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit the owner, acl property."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o",
	//   "request": {
	//     "$ref": "Object"
	//   },
	//   "response": {
	//     "$ref": "Object"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "storage.objects.list":

type ObjectsListCall struct {
	s            *Service
	bucket       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of objects matching the criteria.
func (r *ObjectsService) List(bucket string) *ObjectsListCall {
	c := &ObjectsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	return c
}

// Delimiter sets the optional parameter "delimiter": Returns results in
// a directory-like mode. items will contain only objects whose names,
// aside from the prefix, do not contain delimiter. Objects whose names,
// aside from the prefix, contain delimiter will have their name,
// truncated after the delimiter, returned in prefixes. Duplicate
// prefixes are omitted.
func (c *ObjectsListCall) Delimiter(delimiter string) *ObjectsListCall {
	c.urlParams_.Set("delimiter", delimiter)
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of items plus prefixes to return in a single page of responses. As
// duplicate prefixes are omitted, fewer total results may be returned
// than requested. The service will use this parameter or 1,000 items,
// whichever is smaller.
func (c *ObjectsListCall) MaxResults(maxResults int64) *ObjectsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A
// previously-returned page token representing part of the larger set of
// results to view.
func (c *ObjectsListCall) PageToken(pageToken string) *ObjectsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Prefix sets the optional parameter "prefix": Filter results to
// objects whose names begin with this prefix.
func (c *ObjectsListCall) Prefix(prefix string) *ObjectsListCall {
	c.urlParams_.Set("prefix", prefix)
	return c
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to noAcl.
//
// Possible values:
//   "full" - Include all properties.
//   "noAcl" - Omit the owner, acl property.
func (c *ObjectsListCall) Projection(projection string) *ObjectsListCall {
	c.urlParams_.Set("projection", projection)
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectsListCall) UserProject(userProject string) *ObjectsListCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Versions sets the optional parameter "versions": If true, lists all
// versions of an object as distinct results. The default is false. For
// more information, see Object Versioning.
func (c *ObjectsListCall) Versions(versions bool) *ObjectsListCall {
	c.urlParams_.Set("versions", fmt.Sprint(versions))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsListCall) Fields(s ...googleapi.Field) *ObjectsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ObjectsListCall) IfNoneMatch(entityTag string) *ObjectsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectsListCall) Context(ctx context.Context) *ObjectsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/o")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objects.list" call.
// Exactly one of *Objects or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Objects.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ObjectsListCall) Do(opts ...googleapi.CallOption) (*Objects, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Objects{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of objects matching the criteria.",
	//   "httpMethod": "GET",
	//   "id": "storage.objects.list",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of the bucket in which to look for objects.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "delimiter": {
	//       "description": "Returns results in a directory-like mode. items will contain only objects whose names, aside from the prefix, do not contain delimiter. Objects whose names, aside from the prefix, contain delimiter will have their name, truncated after the delimiter, returned in prefixes. Duplicate prefixes are omitted.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "1000",
	//       "description": "Maximum number of items plus prefixes to return in a single page of responses. As duplicate prefixes are omitted, fewer total results may be returned than requested. The service will use this parameter or 1,000 items, whichever is smaller.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A previously-returned page token representing part of the larger set of results to view.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "prefix": {
	//       "description": "Filter results to objects whose names begin with this prefix.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to noAcl.",
	//       "enum": [
	//         "full",
	//         "noAcl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit the owner, acl property."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "versions": {
	//       "description": "If true, lists all versions of an object as distinct results. The default is false. For more information, see Object Versioning.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "b/{bucket}/o",
	//   "response": {
	//     "$ref": "Objects"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ],
	//   "supportsSubscription": true
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ObjectsListCall) Pages(ctx context.Context, f func(*Objects) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "storage.objects.patch":

type ObjectsPatchCall struct {
	s          *Service
	bucket     string
	object     string
	object2    *Object
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Patches an object's metadata.
func (r *ObjectsService) Patch(bucket string, object string, object2 *Object) *ObjectsPatchCall {
	c := &ObjectsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.object = object
	c.object2 = object2
	return c
}

// Generation sets the optional parameter "generation": If present,
// selects a specific revision of this object (as opposed to the latest
// version, the default).
func (c *ObjectsPatchCall) Generation(generation int64) *ObjectsPatchCall {
	c.urlParams_.Set("generation", fmt.Sprint(generation))
	return c
}

// IfGenerationMatch sets the optional parameter "ifGenerationMatch":
// Makes the operation conditional on whether the object's current
// generation matches the given value. Setting to 0 makes the operation
// succeed only if there are no live versions of the object.
func (c *ObjectsPatchCall) IfGenerationMatch(ifGenerationMatch int64) *ObjectsPatchCall {
	c.urlParams_.Set("ifGenerationMatch", fmt.Sprint(ifGenerationMatch))
	return c
}

// IfGenerationNotMatch sets the optional parameter
// "ifGenerationNotMatch": Makes the operation conditional on whether
// the object's current generation does not match the given value. If no
// live object exists, the precondition fails. Setting to 0 makes the
// operation succeed only if there is a live version of the object.
func (c *ObjectsPatchCall) IfGenerationNotMatch(ifGenerationNotMatch int64) *ObjectsPatchCall {
	c.urlParams_.Set("ifGenerationNotMatch", fmt.Sprint(ifGenerationNotMatch))
	return c
}

// IfMetagenerationMatch sets the optional parameter
// "ifMetagenerationMatch": Makes the operation conditional on whether
// the object's current metageneration matches the given value.
func (c *ObjectsPatchCall) IfMetagenerationMatch(ifMetagenerationMatch int64) *ObjectsPatchCall {
	c.urlParams_.Set("ifMetagenerationMatch", fmt.Sprint(ifMetagenerationMatch))
	return c
}

// IfMetagenerationNotMatch sets the optional parameter
// "ifMetagenerationNotMatch": Makes the operation conditional on
// whether the object's current metageneration does not match the given
// value.
func (c *ObjectsPatchCall) IfMetagenerationNotMatch(ifMetagenerationNotMatch int64) *ObjectsPatchCall {
	c.urlParams_.Set("ifMetagenerationNotMatch", fmt.Sprint(ifMetagenerationNotMatch))
	return c
}

// PredefinedAcl sets the optional parameter "predefinedAcl": Apply a
// predefined set of access controls to this object.
//
// Possible values:
//   "authenticatedRead" - Object owner gets OWNER access, and
// allAuthenticatedUsers get READER access.
//   "bucketOwnerFullControl" - Object owner gets OWNER access, and
// project team owners get OWNER access.
//   "bucketOwnerRead" - Object owner gets OWNER access, and project
// team owners get READER access.
//   "private" - Object owner gets OWNER access.
//   "projectPrivate" - Object owner gets OWNER access, and project team
// members get access according to their roles.
//   "publicRead" - Object owner gets OWNER access, and allUsers get
// READER access.
func (c *ObjectsPatchCall) PredefinedAcl(predefinedAcl string) *ObjectsPatchCall {
	c.urlParams_.Set("predefinedAcl", predefinedAcl)
	return c
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to full.
//
// Possible values:
//   "full" - Include all properties.
//   "noAcl" - Omit the owner, acl property.
func (c *ObjectsPatchCall) Projection(projection string) *ObjectsPatchCall {
	c.urlParams_.Set("projection", projection)
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request, for Requester Pays buckets.
func (c *ObjectsPatchCall) UserProject(userProject string) *ObjectsPatchCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsPatchCall) Fields(s ...googleapi.Field) *ObjectsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectsPatchCall) Context(ctx context.Context) *ObjectsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.object2)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/o/{object}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objects.patch" call.
// Exactly one of *Object or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Object.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ObjectsPatchCall) Do(opts ...googleapi.CallOption) (*Object, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Object{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Patches an object's metadata.",
	//   "httpMethod": "PATCH",
	//   "id": "storage.objects.patch",
	//   "parameterOrder": [
	//     "bucket",
	//     "object"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of the bucket in which the object resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "generation": {
	//       "description": "If present, selects a specific revision of this object (as opposed to the latest version, the default).",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifGenerationMatch": {
	//       "description": "Makes the operation conditional on whether the object's current generation matches the given value. Setting to 0 makes the operation succeed only if there are no live versions of the object.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifGenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the object's current generation does not match the given value. If no live object exists, the precondition fails. Setting to 0 makes the operation succeed only if there is a live version of the object.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationMatch": {
	//       "description": "Makes the operation conditional on whether the object's current metageneration matches the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the object's current metageneration does not match the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "predefinedAcl": {
	//       "description": "Apply a predefined set of access controls to this object.",
	//       "enum": [
	//         "authenticatedRead",
	//         "bucketOwnerFullControl",
	//         "bucketOwnerRead",
	//         "private",
	//         "projectPrivate",
	//         "publicRead"
	//       ],
	//       "enumDescriptions": [
	//         "Object owner gets OWNER access, and allAuthenticatedUsers get READER access.",
	//         "Object owner gets OWNER access, and project team owners get OWNER access.",
	//         "Object owner gets OWNER access, and project team owners get READER access.",
	//         "Object owner gets OWNER access.",
	//         "Object owner gets OWNER access, and project team members get access according to their roles.",
	//         "Object owner gets OWNER access, and allUsers get READER access."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to full.",
	//       "enum": [
	//         "full",
	//         "noAcl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit the owner, acl property."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request, for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}",
	//   "request": {
	//     "$ref": "Object"
	//   },
	//   "response": {
	//     "$ref": "Object"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.objects.rewrite":

type ObjectsRewriteCall struct {
	s                 *Service
	sourceBucket      string
	sourceObject      string
	destinationBucket string
	destinationObject string
	object            *Object
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Rewrite: Rewrites a source object to a destination object. Optionally
// overrides metadata.
func (r *ObjectsService) Rewrite(sourceBucket string, sourceObject string, destinationBucket string, destinationObject string, object *Object) *ObjectsRewriteCall {
	c := &ObjectsRewriteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.sourceBucket = sourceBucket
	c.sourceObject = sourceObject
	c.destinationBucket = destinationBucket
	c.destinationObject = destinationObject
	c.object = object
	return c
}

// DestinationKmsKeyName sets the optional parameter
// "destinationKmsKeyName": Resource name of the Cloud KMS key, of the
// form
// projects/my-project/locations/global/keyRings/my-kr/cryptoKeys/my-key,
//  that will be used to encrypt the object. Overrides the object
// metadata's kms_key_name value, if any.
func (c *ObjectsRewriteCall) DestinationKmsKeyName(destinationKmsKeyName string) *ObjectsRewriteCall {
	c.urlParams_.Set("destinationKmsKeyName", destinationKmsKeyName)
	return c
}

// DestinationPredefinedAcl sets the optional parameter
// "destinationPredefinedAcl": Apply a predefined set of access controls
// to the destination object.
//
// Possible values:
//   "authenticatedRead" - Object owner gets OWNER access, and
// allAuthenticatedUsers get READER access.
//   "bucketOwnerFullControl" - Object owner gets OWNER access, and
// project team owners get OWNER access.
//   "bucketOwnerRead" - Object owner gets OWNER access, and project
// team owners get READER access.
//   "private" - Object owner gets OWNER access.
//   "projectPrivate" - Object owner gets OWNER access, and project team
// members get access according to their roles.
//   "publicRead" - Object owner gets OWNER access, and allUsers get
// READER access.
func (c *ObjectsRewriteCall) DestinationPredefinedAcl(destinationPredefinedAcl string) *ObjectsRewriteCall {
	c.urlParams_.Set("destinationPredefinedAcl", destinationPredefinedAcl)
	return c
}

// IfGenerationMatch sets the optional parameter "ifGenerationMatch":
// Makes the operation conditional on whether the object's current
// generation matches the given value. Setting to 0 makes the operation
// succeed only if there are no live versions of the object.
func (c *ObjectsRewriteCall) IfGenerationMatch(ifGenerationMatch int64) *ObjectsRewriteCall {
	c.urlParams_.Set("ifGenerationMatch", fmt.Sprint(ifGenerationMatch))
	return c
}

// IfGenerationNotMatch sets the optional parameter
// "ifGenerationNotMatch": Makes the operation conditional on whether
// the object's current generation does not match the given value. If no
// live object exists, the precondition fails. Setting to 0 makes the
// operation succeed only if there is a live version of the object.
func (c *ObjectsRewriteCall) IfGenerationNotMatch(ifGenerationNotMatch int64) *ObjectsRewriteCall {
	c.urlParams_.Set("ifGenerationNotMatch", fmt.Sprint(ifGenerationNotMatch))
	return c
}

// IfMetagenerationMatch sets the optional parameter
// "ifMetagenerationMatch": Makes the operation conditional on whether
// the destination object's current metageneration matches the given
// value.
func (c *ObjectsRewriteCall) IfMetagenerationMatch(ifMetagenerationMatch int64) *ObjectsRewriteCall {
	c.urlParams_.Set("ifMetagenerationMatch", fmt.Sprint(ifMetagenerationMatch))
	return c
}

// IfMetagenerationNotMatch sets the optional parameter
// "ifMetagenerationNotMatch": Makes the operation conditional on
// whether the destination object's current metageneration does not
// match the given value.
func (c *ObjectsRewriteCall) IfMetagenerationNotMatch(ifMetagenerationNotMatch int64) *ObjectsRewriteCall {
	c.urlParams_.Set("ifMetagenerationNotMatch", fmt.Sprint(ifMetagenerationNotMatch))
	return c
}

// IfSourceGenerationMatch sets the optional parameter
// "ifSourceGenerationMatch": Makes the operation conditional on whether
// the source object's current generation matches the given value.
func (c *ObjectsRewriteCall) IfSourceGenerationMatch(ifSourceGenerationMatch int64) *ObjectsRewriteCall {
	c.urlParams_.Set("ifSourceGenerationMatch", fmt.Sprint(ifSourceGenerationMatch))
	return c
}

// IfSourceGenerationNotMatch sets the optional parameter
// "ifSourceGenerationNotMatch": Makes the operation conditional on
// whether the source object's current generation does not match the
// given value.
func (c *ObjectsRewriteCall) IfSourceGenerationNotMatch(ifSourceGenerationNotMatch int64) *ObjectsRewriteCall {
	c.urlParams_.Set("ifSourceGenerationNotMatch", fmt.Sprint(ifSourceGenerationNotMatch))
	return c
}

// IfSourceMetagenerationMatch sets the optional parameter
// "ifSourceMetagenerationMatch": Makes the operation conditional on
// whether the source object's current metageneration matches the given
// value.
func (c *ObjectsRewriteCall) IfSourceMetagenerationMatch(ifSourceMetagenerationMatch int64) *ObjectsRewriteCall {
	c.urlParams_.Set("ifSourceMetagenerationMatch", fmt.Sprint(ifSourceMetagenerationMatch))
	return c
}

// IfSourceMetagenerationNotMatch sets the optional parameter
// "ifSourceMetagenerationNotMatch": Makes the operation conditional on
// whether the source object's current metageneration does not match the
// given value.
func (c *ObjectsRewriteCall) IfSourceMetagenerationNotMatch(ifSourceMetagenerationNotMatch int64) *ObjectsRewriteCall {
	c.urlParams_.Set("ifSourceMetagenerationNotMatch", fmt.Sprint(ifSourceMetagenerationNotMatch))
	return c
}

// MaxBytesRewrittenPerCall sets the optional parameter
// "maxBytesRewrittenPerCall": The maximum number of bytes that will be
// rewritten per rewrite request. Most callers shouldn't need to specify
// this parameter - it is primarily in place to support testing. If
// specified the value must be an integral multiple of 1 MiB (1048576).
// Also, this only applies to requests where the source and destination
// span locations and/or storage classes. Finally, this value must not
// change across rewrite calls else you'll get an error that the
// rewriteToken is invalid.
func (c *ObjectsRewriteCall) MaxBytesRewrittenPerCall(maxBytesRewrittenPerCall int64) *ObjectsRewriteCall {
	c.urlParams_.Set("maxBytesRewrittenPerCall", fmt.Sprint(maxBytesRewrittenPerCall))
	return c
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to noAcl, unless the object resource
// specifies the acl property, when it defaults to full.
//
// Possible values:
//   "full" - Include all properties.
//   "noAcl" - Omit the owner, acl property.
func (c *ObjectsRewriteCall) Projection(projection string) *ObjectsRewriteCall {
	c.urlParams_.Set("projection", projection)
	return c
}

// RewriteToken sets the optional parameter "rewriteToken": Include this
// field (from the previous rewrite response) on each rewrite request
// after the first one, until the rewrite response 'done' flag is true.
// Calls that provide a rewriteToken can omit all other request fields,
// but if included those fields must match the values provided in the
// first rewrite request.
func (c *ObjectsRewriteCall) RewriteToken(rewriteToken string) *ObjectsRewriteCall {
	c.urlParams_.Set("rewriteToken", rewriteToken)
	return c
}

// SourceGeneration sets the optional parameter "sourceGeneration": If
// present, selects a specific revision of the source object (as opposed
// to the latest version, the default).
func (c *ObjectsRewriteCall) SourceGeneration(sourceGeneration int64) *ObjectsRewriteCall {
	c.urlParams_.Set("sourceGeneration", fmt.Sprint(sourceGeneration))
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectsRewriteCall) UserProject(userProject string) *ObjectsRewriteCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsRewriteCall) Fields(s ...googleapi.Field) *ObjectsRewriteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectsRewriteCall) Context(ctx context.Context) *ObjectsRewriteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectsRewriteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectsRewriteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.object)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{sourceBucket}/o/{sourceObject}/rewriteTo/b/{destinationBucket}/o/{destinationObject}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"sourceBucket":      c.sourceBucket,
		"sourceObject":      c.sourceObject,
		"destinationBucket": c.destinationBucket,
		"destinationObject": c.destinationObject,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objects.rewrite" call.
// Exactly one of *RewriteResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *RewriteResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ObjectsRewriteCall) Do(opts ...googleapi.CallOption) (*RewriteResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &RewriteResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Rewrites a source object to a destination object. Optionally overrides metadata.",
	//   "httpMethod": "POST",
	//   "id": "storage.objects.rewrite",
	//   "parameterOrder": [
	//     "sourceBucket",
	//     "sourceObject",
	//     "destinationBucket",
	//     "destinationObject"
	//   ],
	//   "parameters": {
	//     "destinationBucket": {
	//       "description": "Name of the bucket in which to store the new object. Overrides the provided object metadata's bucket value, if any.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "destinationKmsKeyName": {
	//       "description": "Resource name of the Cloud KMS key, of the form projects/my-project/locations/global/keyRings/my-kr/cryptoKeys/my-key, that will be used to encrypt the object. Overrides the object metadata's kms_key_name value, if any.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "destinationObject": {
	//       "description": "Name of the new object. Required when the object metadata is not otherwise provided. Overrides the object metadata's name value, if any. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "destinationPredefinedAcl": {
	//       "description": "Apply a predefined set of access controls to the destination object.",
	//       "enum": [
	//         "authenticatedRead",
	//         "bucketOwnerFullControl",
	//         "bucketOwnerRead",
	//         "private",
	//         "projectPrivate",
	//         "publicRead"
	//       ],
	//       "enumDescriptions": [
	//         "Object owner gets OWNER access, and allAuthenticatedUsers get READER access.",
	//         "Object owner gets OWNER access, and project team owners get OWNER access.",
	//         "Object owner gets OWNER access, and project team owners get READER access.",
	//         "Object owner gets OWNER access.",
	//         "Object owner gets OWNER access, and project team members get access according to their roles.",
	//         "Object owner gets OWNER access, and allUsers get READER access."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifGenerationMatch": {
	//       "description": "Makes the operation conditional on whether the object's current generation matches the given value. Setting to 0 makes the operation succeed only if there are no live versions of the object.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifGenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the object's current generation does not match the given value. If no live object exists, the precondition fails. Setting to 0 makes the operation succeed only if there is a live version of the object.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationMatch": {
	//       "description": "Makes the operation conditional on whether the destination object's current metageneration matches the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the destination object's current metageneration does not match the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifSourceGenerationMatch": {
	//       "description": "Makes the operation conditional on whether the source object's current generation matches the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifSourceGenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the source object's current generation does not match the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifSourceMetagenerationMatch": {
	//       "description": "Makes the operation conditional on whether the source object's current metageneration matches the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifSourceMetagenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the source object's current metageneration does not match the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxBytesRewrittenPerCall": {
	//       "description": "The maximum number of bytes that will be rewritten per rewrite request. Most callers shouldn't need to specify this parameter - it is primarily in place to support testing. If specified the value must be an integral multiple of 1 MiB (1048576). Also, this only applies to requests where the source and destination span locations and/or storage classes. Finally, this value must not change across rewrite calls else you'll get an error that the rewriteToken is invalid.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to noAcl, unless the object resource specifies the acl property, when it defaults to full.",
	//       "enum": [
	//         "full",
	//         "noAcl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit the owner, acl property."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "rewriteToken": {
	//       "description": "Include this field (from the previous rewrite response) on each rewrite request after the first one, until the rewrite response 'done' flag is true. Calls that provide a rewriteToken can omit all other request fields, but if included those fields must match the values provided in the first rewrite request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sourceBucket": {
	//       "description": "Name of the bucket in which to find the source object.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sourceGeneration": {
	//       "description": "If present, selects a specific revision of the source object (as opposed to the latest version, the default).",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sourceObject": {
	//       "description": "Name of the source object. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{sourceBucket}/o/{sourceObject}/rewriteTo/b/{destinationBucket}/o/{destinationObject}",
	//   "request": {
	//     "$ref": "Object"
	//   },
	//   "response": {
	//     "$ref": "RewriteResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.objects.setIamPolicy":

type ObjectsSetIamPolicyCall struct {
	s          *Service
	bucket     string
	object     string
	policy     *Policy
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// SetIamPolicy: Updates an IAM policy for the specified object.
func (r *ObjectsService) SetIamPolicy(bucket string, object string, policy *Policy) *ObjectsSetIamPolicyCall {
	c := &ObjectsSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.object = object
	c.policy = policy
	return c
}

// Generation sets the optional parameter "generation": If present,
// selects a specific revision of this object (as opposed to the latest
// version, the default).
func (c *ObjectsSetIamPolicyCall) Generation(generation int64) *ObjectsSetIamPolicyCall {
	c.urlParams_.Set("generation", fmt.Sprint(generation))
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectsSetIamPolicyCall) UserProject(userProject string) *ObjectsSetIamPolicyCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsSetIamPolicyCall) Fields(s ...googleapi.Field) *ObjectsSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectsSetIamPolicyCall) Context(ctx context.Context) *ObjectsSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectsSetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectsSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.policy)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/o/{object}/iam")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objects.setIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ObjectsSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an IAM policy for the specified object.",
	//   "httpMethod": "PUT",
	//   "id": "storage.objects.setIamPolicy",
	//   "parameterOrder": [
	//     "bucket",
	//     "object"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of the bucket in which the object resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "generation": {
	//       "description": "If present, selects a specific revision of this object (as opposed to the latest version, the default).",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}/iam",
	//   "request": {
	//     "$ref": "Policy"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.objects.testIamPermissions":

type ObjectsTestIamPermissionsCall struct {
	s            *Service
	bucket       string
	object       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// TestIamPermissions: Tests a set of permissions on the given object to
// see which, if any, are held by the caller.
func (r *ObjectsService) TestIamPermissions(bucket string, object string, permissions []string) *ObjectsTestIamPermissionsCall {
	c := &ObjectsTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.object = object
	c.urlParams_.SetMulti("permissions", append([]string{}, permissions...))
	return c
}

// Generation sets the optional parameter "generation": If present,
// selects a specific revision of this object (as opposed to the latest
// version, the default).
func (c *ObjectsTestIamPermissionsCall) Generation(generation int64) *ObjectsTestIamPermissionsCall {
	c.urlParams_.Set("generation", fmt.Sprint(generation))
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectsTestIamPermissionsCall) UserProject(userProject string) *ObjectsTestIamPermissionsCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsTestIamPermissionsCall) Fields(s ...googleapi.Field) *ObjectsTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ObjectsTestIamPermissionsCall) IfNoneMatch(entityTag string) *ObjectsTestIamPermissionsCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectsTestIamPermissionsCall) Context(ctx context.Context) *ObjectsTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectsTestIamPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectsTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/o/{object}/iam/testPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objects.testIamPermissions" call.
// Exactly one of *TestIamPermissionsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *TestIamPermissionsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ObjectsTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*TestIamPermissionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestIamPermissionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Tests a set of permissions on the given object to see which, if any, are held by the caller.",
	//   "httpMethod": "GET",
	//   "id": "storage.objects.testIamPermissions",
	//   "parameterOrder": [
	//     "bucket",
	//     "object",
	//     "permissions"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of the bucket in which the object resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "generation": {
	//       "description": "If present, selects a specific revision of this object (as opposed to the latest version, the default).",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "permissions": {
	//       "description": "Permissions to test.",
	//       "location": "query",
	//       "repeated": true,
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}/iam/testPermissions",
	//   "response": {
	//     "$ref": "TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.objects.update":

type ObjectsUpdateCall struct {
	s          *Service
	bucket     string
	object     string
	object2    *Object
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Updates an object's metadata.
func (r *ObjectsService) Update(bucket string, object string, object2 *Object) *ObjectsUpdateCall {
	c := &ObjectsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.object = object
	c.object2 = object2
	return c
}

// Generation sets the optional parameter "generation": If present,
// selects a specific revision of this object (as opposed to the latest
// version, the default).
func (c *ObjectsUpdateCall) Generation(generation int64) *ObjectsUpdateCall {
	c.urlParams_.Set("generation", fmt.Sprint(generation))
	return c
}

// IfGenerationMatch sets the optional parameter "ifGenerationMatch":
// Makes the operation conditional on whether the object's current
// generation matches the given value. Setting to 0 makes the operation
// succeed only if there are no live versions of the object.
func (c *ObjectsUpdateCall) IfGenerationMatch(ifGenerationMatch int64) *ObjectsUpdateCall {
	c.urlParams_.Set("ifGenerationMatch", fmt.Sprint(ifGenerationMatch))
	return c
}

// IfGenerationNotMatch sets the optional parameter
// "ifGenerationNotMatch": Makes the operation conditional on whether
// the object's current generation does not match the given value. If no
// live object exists, the precondition fails. Setting to 0 makes the
// operation succeed only if there is a live version of the object.
func (c *ObjectsUpdateCall) IfGenerationNotMatch(ifGenerationNotMatch int64) *ObjectsUpdateCall {
	c.urlParams_.Set("ifGenerationNotMatch", fmt.Sprint(ifGenerationNotMatch))
	return c
}

// IfMetagenerationMatch sets the optional parameter
// "ifMetagenerationMatch": Makes the operation conditional on whether
// the object's current metageneration matches the given value.
func (c *ObjectsUpdateCall) IfMetagenerationMatch(ifMetagenerationMatch int64) *ObjectsUpdateCall {
	c.urlParams_.Set("ifMetagenerationMatch", fmt.Sprint(ifMetagenerationMatch))
	return c
}

// IfMetagenerationNotMatch sets the optional parameter
// "ifMetagenerationNotMatch": Makes the operation conditional on
// whether the object's current metageneration does not match the given
// value.
func (c *ObjectsUpdateCall) IfMetagenerationNotMatch(ifMetagenerationNotMatch int64) *ObjectsUpdateCall {
	c.urlParams_.Set("ifMetagenerationNotMatch", fmt.Sprint(ifMetagenerationNotMatch))
	return c
}

// PredefinedAcl sets the optional parameter "predefinedAcl": Apply a
// predefined set of access controls to this object.
//
// Possible values:
//   "authenticatedRead" - Object owner gets OWNER access, and
// allAuthenticatedUsers get READER access.
//   "bucketOwnerFullControl" - Object owner gets OWNER access, and
// project team owners get OWNER access.
//   "bucketOwnerRead" - Object owner gets OWNER access, and project
// team owners get READER access.
//   "private" - Object owner gets OWNER access.
//   "projectPrivate" - Object owner gets OWNER access, and project team
// members get access according to their roles.
//   "publicRead" - Object owner gets OWNER access, and allUsers get
// READER access.
func (c *ObjectsUpdateCall) PredefinedAcl(predefinedAcl string) *ObjectsUpdateCall {
	c.urlParams_.Set("predefinedAcl", predefinedAcl)
	return c
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to full.
//
// Possible values:
//   "full" - Include all properties.
//   "noAcl" - Omit the owner, acl property.
func (c *ObjectsUpdateCall) Projection(projection string) *ObjectsUpdateCall {
	c.urlParams_.Set("projection", projection)
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectsUpdateCall) UserProject(userProject string) *ObjectsUpdateCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsUpdateCall) Fields(s ...googleapi.Field) *ObjectsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectsUpdateCall) Context(ctx context.Context) *ObjectsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.object2)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/o/{object}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objects.update" call.
// Exactly one of *Object or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Object.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ObjectsUpdateCall) Do(opts ...googleapi.CallOption) (*Object, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Object{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an object's metadata.",
	//   "httpMethod": "PUT",
	//   "id": "storage.objects.update",
	//   "parameterOrder": [
	//     "bucket",
	//     "object"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of the bucket in which the object resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "generation": {
	//       "description": "If present, selects a specific revision of this object (as opposed to the latest version, the default).",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifGenerationMatch": {
	//       "description": "Makes the operation conditional on whether the object's current generation matches the given value. Setting to 0 makes the operation succeed only if there are no live versions of the object.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifGenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the object's current generation does not match the given value. If no live object exists, the precondition fails. Setting to 0 makes the operation succeed only if there is a live version of the object.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationMatch": {
	//       "description": "Makes the operation conditional on whether the object's current metageneration matches the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ifMetagenerationNotMatch": {
	//       "description": "Makes the operation conditional on whether the object's current metageneration does not match the given value.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object. For information about how to URL encode object names to be path safe, see Encoding URI Path Parts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "predefinedAcl": {
	//       "description": "Apply a predefined set of access controls to this object.",
	//       "enum": [
	//         "authenticatedRead",
	//         "bucketOwnerFullControl",
	//         "bucketOwnerRead",
	//         "private",
	//         "projectPrivate",
	//         "publicRead"
	//       ],
	//       "enumDescriptions": [
	//         "Object owner gets OWNER access, and allAuthenticatedUsers get READER access.",
	//         "Object owner gets OWNER access, and project team owners get OWNER access.",
	//         "Object owner gets OWNER access, and project team owners get READER access.",
	//         "Object owner gets OWNER access.",
	//         "Object owner gets OWNER access, and project team members get access according to their roles.",
	//         "Object owner gets OWNER access, and allUsers get READER access."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to full.",
	//       "enum": [
	//         "full",
	//         "noAcl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit the owner, acl property."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}",
	//   "request": {
	//     "$ref": "Object"
	//   },
	//   "response": {
	//     "$ref": "Object"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.objects.watchAll":

type ObjectsWatchAllCall struct {
	s          *Service
	bucket     string
	channel    *Channel
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// WatchAll: Watch for changes on all objects in a bucket.
func (r *ObjectsService) WatchAll(bucket string, channel *Channel) *ObjectsWatchAllCall {
	c := &ObjectsWatchAllCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.channel = channel
	return c
}

// Delimiter sets the optional parameter "delimiter": Returns results in
// a directory-like mode. items will contain only objects whose names,
// aside from the prefix, do not contain delimiter. Objects whose names,
// aside from the prefix, contain delimiter will have their name,
// truncated after the delimiter, returned in prefixes. Duplicate
// prefixes are omitted.
func (c *ObjectsWatchAllCall) Delimiter(delimiter string) *ObjectsWatchAllCall {
	c.urlParams_.Set("delimiter", delimiter)
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of items plus prefixes to return in a single page of responses. As
// duplicate prefixes are omitted, fewer total results may be returned
// than requested. The service will use this parameter or 1,000 items,
// whichever is smaller.
func (c *ObjectsWatchAllCall) MaxResults(maxResults int64) *ObjectsWatchAllCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A
// previously-returned page token representing part of the larger set of
// results to view.
func (c *ObjectsWatchAllCall) PageToken(pageToken string) *ObjectsWatchAllCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Prefix sets the optional parameter "prefix": Filter results to
// objects whose names begin with this prefix.
func (c *ObjectsWatchAllCall) Prefix(prefix string) *ObjectsWatchAllCall {
	c.urlParams_.Set("prefix", prefix)
	return c
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to noAcl.
//
// Possible values:
//   "full" - Include all properties.
//   "noAcl" - Omit the owner, acl property.
func (c *ObjectsWatchAllCall) Projection(projection string) *ObjectsWatchAllCall {
	c.urlParams_.Set("projection", projection)
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request. Required for Requester Pays buckets.
func (c *ObjectsWatchAllCall) UserProject(userProject string) *ObjectsWatchAllCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Versions sets the optional parameter "versions": If true, lists all
// versions of an object as distinct results. The default is false. For
// more information, see Object Versioning.
func (c *ObjectsWatchAllCall) Versions(versions bool) *ObjectsWatchAllCall {
	c.urlParams_.Set("versions", fmt.Sprint(versions))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsWatchAllCall) Fields(s ...googleapi.Field) *ObjectsWatchAllCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ObjectsWatchAllCall) Context(ctx context.Context) *ObjectsWatchAllCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ObjectsWatchAllCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ObjectsWatchAllCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.channel)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "b/{bucket}/o/watch")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.objects.watchAll" call.
// Exactly one of *Channel or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Channel.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ObjectsWatchAllCall) Do(opts ...googleapi.CallOption) (*Channel, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Channel{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Watch for changes on all objects in a bucket.",
	//   "httpMethod": "POST",
	//   "id": "storage.objects.watchAll",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of the bucket in which to look for objects.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "delimiter": {
	//       "description": "Returns results in a directory-like mode. items will contain only objects whose names, aside from the prefix, do not contain delimiter. Objects whose names, aside from the prefix, contain delimiter will have their name, truncated after the delimiter, returned in prefixes. Duplicate prefixes are omitted.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "1000",
	//       "description": "Maximum number of items plus prefixes to return in a single page of responses. As duplicate prefixes are omitted, fewer total results may be returned than requested. The service will use this parameter or 1,000 items, whichever is smaller.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A previously-returned page token representing part of the larger set of results to view.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "prefix": {
	//       "description": "Filter results to objects whose names begin with this prefix.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to noAcl.",
	//       "enum": [
	//         "full",
	//         "noAcl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit the owner, acl property."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request. Required for Requester Pays buckets.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "versions": {
	//       "description": "If true, lists all versions of an object as distinct results. The default is false. For more information, see Object Versioning.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "b/{bucket}/o/watch",
	//   "request": {
	//     "$ref": "Channel",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "Channel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "storage.projects.serviceAccount.get":

type ProjectsServiceAccountGetCall struct {
	s            *Service
	projectId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Get the email address of this project's Google Cloud Storage
// service account.
func (r *ProjectsServiceAccountService) Get(projectId string) *ProjectsServiceAccountGetCall {
	c := &ProjectsServiceAccountGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	return c
}

// UserProject sets the optional parameter "userProject": The project to
// be billed for this request.
func (c *ProjectsServiceAccountGetCall) UserProject(userProject string) *ProjectsServiceAccountGetCall {
	c.urlParams_.Set("userProject", userProject)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsServiceAccountGetCall) Fields(s ...googleapi.Field) *ProjectsServiceAccountGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsServiceAccountGetCall) IfNoneMatch(entityTag string) *ProjectsServiceAccountGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsServiceAccountGetCall) Context(ctx context.Context) *ProjectsServiceAccountGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsServiceAccountGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsServiceAccountGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/serviceAccount")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "storage.projects.serviceAccount.get" call.
// Exactly one of *ServiceAccount or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ServiceAccount.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsServiceAccountGetCall) Do(opts ...googleapi.CallOption) (*ServiceAccount, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ServiceAccount{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the email address of this project's Google Cloud Storage service account.",
	//   "httpMethod": "GET",
	//   "id": "storage.projects.serviceAccount.get",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Project ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userProject": {
	//       "description": "The project to be billed for this request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/serviceAccount",
	//   "response": {
	//     "$ref": "ServiceAccount"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}
