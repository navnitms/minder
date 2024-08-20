// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
)

type Querier interface {
	BulkGetProfilesByID(ctx context.Context, profileIds []uuid.UUID) ([]BulkGetProfilesByIDRow, error)
	CountProfilesByEntityType(ctx context.Context) ([]CountProfilesByEntityTypeRow, error)
	CountProfilesByName(ctx context.Context, name string) (int64, error)
	CountRepositories(ctx context.Context) (int64, error)
	CountUsers(ctx context.Context) (int64, error)
	// CreateEntity adds an entry to the entity_instances table so it can be tracked by Minder.
	CreateEntity(ctx context.Context, arg CreateEntityParams) (EntityInstance, error)
	// CreateEntityWithID adds an entry to the entities table with a specific ID so it can be tracked by Minder.
	CreateEntityWithID(ctx context.Context, arg CreateEntityWithIDParams) (EntityInstance, error)
	// CreateInvitation creates a new invitation. The code is a secret that is sent
	// to the invitee, and the email is the address to which the invitation will be
	// sent. The role is the role that the invitee will have when they accept the
	// invitation. The project is the project to which the invitee will be invited.
	// The sponsor is the user who is inviting the invitee.
	CreateInvitation(ctx context.Context, arg CreateInvitationParams) (UserInvite, error)
	// CreateOrEnsureEntityByID adds an entry to the entity_instances table if it does not exist, or returns the existing entry.
	CreateOrEnsureEntityByID(ctx context.Context, arg CreateOrEnsureEntityByIDParams) (EntityInstance, error)
	CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error)
	CreateProfileForEntity(ctx context.Context, arg CreateProfileForEntityParams) (EntityProfile, error)
	CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error)
	CreateProjectWithID(ctx context.Context, arg CreateProjectWithIDParams) (Project, error)
	CreateProvider(ctx context.Context, arg CreateProviderParams) (Provider, error)
	CreatePullRequest(ctx context.Context, arg CreatePullRequestParams) (PullRequest, error)
	CreateRepository(ctx context.Context, arg CreateRepositoryParams) (Repository, error)
	CreateRuleType(ctx context.Context, arg CreateRuleTypeParams) (RuleType, error)
	CreateSelector(ctx context.Context, arg CreateSelectorParams) (ProfileSelector, error)
	CreateSessionState(ctx context.Context, arg CreateSessionStateParams) (SessionStore, error)
	// Subscriptions --
	CreateSubscription(ctx context.Context, arg CreateSubscriptionParams) (Subscription, error)
	CreateUser(ctx context.Context, identitySubject string) (User, error)
	DeleteArtifact(ctx context.Context, id uuid.UUID) error
	// DeleteEntity removes an entity from the entity_instances table for a project.
	DeleteEntity(ctx context.Context, arg DeleteEntityParams) error
	// DeleteEntityByName removes an entity from the entity_instances table for a project.
	DeleteEntityByName(ctx context.Context, arg DeleteEntityByNameParams) error
	DeleteEvaluationHistoryByIDs(ctx context.Context, evaluationids []uuid.UUID) (int64, error)
	DeleteExpiredSessionStates(ctx context.Context) (int64, error)
	DeleteInstallationIDByAppID(ctx context.Context, appInstallationID int64) error
	// DeleteInvitation deletes an invitation by its code. This is intended to be
	// called by a user who has issued an invitation and then accepted it, declined
	// it or the sponsor has decided to revoke it.
	DeleteInvitation(ctx context.Context, code string) (UserInvite, error)
	DeleteNonUpdatedRules(ctx context.Context, arg DeleteNonUpdatedRulesParams) error
	DeleteProfile(ctx context.Context, arg DeleteProfileParams) error
	DeleteProfileForEntity(ctx context.Context, arg DeleteProfileForEntityParams) error
	DeleteProject(ctx context.Context, id uuid.UUID) ([]DeleteProjectRow, error)
	DeleteProvider(ctx context.Context, arg DeleteProviderParams) error
	DeletePullRequest(ctx context.Context, arg DeletePullRequestParams) error
	DeleteRepository(ctx context.Context, id uuid.UUID) error
	DeleteRuleType(ctx context.Context, id uuid.UUID) error
	DeleteSelector(ctx context.Context, id uuid.UUID) error
	DeleteSelectorsByProfileID(ctx context.Context, profileID uuid.UUID) error
	DeleteSessionStateByProjectID(ctx context.Context, arg DeleteSessionStateByProjectIDParams) error
	DeleteUser(ctx context.Context, id int32) error
	EnqueueFlush(ctx context.Context, arg EnqueueFlushParams) (FlushCache, error)
	// FindProviders allows us to take a trait and filter
	// providers by it. It also optionally takes a name, in case we want to
	// filter by name as well.
	FindProviders(ctx context.Context, arg FindProvidersParams) ([]Provider, error)
	FlushCache(ctx context.Context, arg FlushCacheParams) (FlushCache, error)
	GetAccessTokenByEnrollmentNonce(ctx context.Context, arg GetAccessTokenByEnrollmentNonceParams) (ProviderAccessToken, error)
	GetAccessTokenByProjectID(ctx context.Context, arg GetAccessTokenByProjectIDParams) (ProviderAccessToken, error)
	GetAccessTokenByProvider(ctx context.Context, provider string) ([]ProviderAccessToken, error)
	GetAccessTokenSinceDate(ctx context.Context, arg GetAccessTokenSinceDateParams) (ProviderAccessToken, error)
	GetArtifactByID(ctx context.Context, arg GetArtifactByIDParams) (Artifact, error)
	GetArtifactByName(ctx context.Context, arg GetArtifactByNameParams) (Artifact, error)
	GetBundle(ctx context.Context, arg GetBundleParams) (Bundle, error)
	GetChildrenProjects(ctx context.Context, id uuid.UUID) ([]GetChildrenProjectsRow, error)
	// GetEntitiesByType retrieves all entities of a given type for a project or hierarchy of projects.
	// this is how one would get all repositories, artifacts, etc.
	GetEntitiesByType(ctx context.Context, arg GetEntitiesByTypeParams) ([]EntityInstance, error)
	// GetEntityByID retrieves an entity by its ID for a project or hierarchy of projects.
	GetEntityByID(ctx context.Context, arg GetEntityByIDParams) (EntityInstance, error)
	// GetEntityByName retrieves an entity by its name for a project or hierarchy of projects.
	GetEntityByName(ctx context.Context, arg GetEntityByNameParams) (EntityInstance, error)
	GetEvaluationHistory(ctx context.Context, arg GetEvaluationHistoryParams) (GetEvaluationHistoryRow, error)
	// GetFeatureInProject verifies if a feature is available for a specific project.
	// It returns the settings for the feature if it is available.
	GetFeatureInProject(ctx context.Context, arg GetFeatureInProjectParams) (json.RawMessage, error)
	// GetImmediateChildrenProjects is a query that returns all the immediate children of a project.
	GetImmediateChildrenProjects(ctx context.Context, parentID uuid.UUID) ([]Project, error)
	GetInstallationIDByAppID(ctx context.Context, appInstallationID int64) (ProviderGithubAppInstallation, error)
	GetInstallationIDByEnrollmentNonce(ctx context.Context, arg GetInstallationIDByEnrollmentNonceParams) (ProviderGithubAppInstallation, error)
	GetInstallationIDByProviderID(ctx context.Context, providerID uuid.NullUUID) (ProviderGithubAppInstallation, error)
	// GetInvitationByCode retrieves an invitation by its code. This is intended to
	// be called by a user who has received an invitation email and is following the
	// link to accept the invitation or when querying for additional info about the
	// invitation.
	GetInvitationByCode(ctx context.Context, code string) (GetInvitationByCodeRow, error)
	// GetInvitationsByEmail retrieves all invitations for a given email address.
	// This is intended to be called by a logged in user with their own email address,
	// to allow them to accept invitations even if email delivery was not working.
	// Note that this requires that the destination email address matches the email
	// address of the logged in user in the external identity service / auth token.
	// This clarification is related solely for user's ListInvitations calls and does
	// not affect to resolving invitations intended for other mail addresses.
	GetInvitationsByEmail(ctx context.Context, email string) ([]GetInvitationsByEmailRow, error)
	// GetInvitationsByEmailAndProject retrieves all invitations by email and project.
	GetInvitationsByEmailAndProject(ctx context.Context, arg GetInvitationsByEmailAndProjectParams) ([]GetInvitationsByEmailAndProjectRow, error)
	// Copyright 2024 Stacklok, Inc
	//
	// Licensed under the Apache License, Version 2.0 (the "License");
	// you may not use this file except in compliance with the License.
	// You may obtain a copy of the License at
	//
	//      http://www.apache.org/licenses/LICENSE-2.0
	//
	// Unless required by applicable law or agreed to in writing, software
	// distributed under the License is distributed on an "AS IS" BASIS,
	// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	// See the License for the specific language governing permissions and
	// limitations under the License.
	GetLatestEvalStateForRuleEntity(ctx context.Context, arg GetLatestEvalStateForRuleEntityParams) (EvaluationStatus, error)
	GetParentProjects(ctx context.Context, id uuid.UUID) ([]uuid.UUID, error)
	GetParentProjectsUntil(ctx context.Context, arg GetParentProjectsUntilParams) ([]uuid.UUID, error)
	GetProfileByID(ctx context.Context, arg GetProfileByIDParams) (Profile, error)
	GetProfileByIDAndLock(ctx context.Context, arg GetProfileByIDAndLockParams) (Profile, error)
	GetProfileByNameAndLock(ctx context.Context, arg GetProfileByNameAndLockParams) (Profile, error)
	GetProfileByProjectAndID(ctx context.Context, arg GetProfileByProjectAndIDParams) ([]GetProfileByProjectAndIDRow, error)
	GetProfileByProjectAndName(ctx context.Context, arg GetProfileByProjectAndNameParams) ([]GetProfileByProjectAndNameRow, error)
	GetProfileStatusByIdAndProject(ctx context.Context, arg GetProfileStatusByIdAndProjectParams) (GetProfileStatusByIdAndProjectRow, error)
	GetProfileStatusByNameAndProject(ctx context.Context, arg GetProfileStatusByNameAndProjectParams) (GetProfileStatusByNameAndProjectRow, error)
	GetProfileStatusByProject(ctx context.Context, projectID uuid.UUID) ([]GetProfileStatusByProjectRow, error)
	GetProjectByID(ctx context.Context, id uuid.UUID) (Project, error)
	GetProjectByName(ctx context.Context, name string) (Project, error)
	GetProjectIDBySessionState(ctx context.Context, sessionState string) (GetProjectIDBySessionStateRow, error)
	GetProviderByID(ctx context.Context, id uuid.UUID) (Provider, error)
	GetProviderByIDAndProject(ctx context.Context, arg GetProviderByIDAndProjectParams) (Provider, error)
	// GetProviderByName allows us to get a provider by its name. This takes
	// into account the project hierarchy, so it will only return the provider
	// if it exists in the project or any of its ancestors. It'll return the first
	// provider that matches the name.
	GetProviderByName(ctx context.Context, arg GetProviderByNameParams) (Provider, error)
	// get a list of repos with webhooks belonging to a provider
	// is used for webhook cleanup during provider deletion
	GetProviderWebhooks(ctx context.Context, providerID uuid.UUID) ([]GetProviderWebhooksRow, error)
	GetPullRequest(ctx context.Context, arg GetPullRequestParams) (PullRequest, error)
	GetPullRequestByID(ctx context.Context, id uuid.UUID) (PullRequest, error)
	GetRepoPathFromArtifactID(ctx context.Context, id uuid.UUID) (GetRepoPathFromArtifactIDRow, error)
	GetRepoPathFromPullRequestID(ctx context.Context, id uuid.UUID) (GetRepoPathFromPullRequestIDRow, error)
	// avoid using this, where possible use GetRepositoryByIDAndProject instead
	GetRepositoryByID(ctx context.Context, id uuid.UUID) (Repository, error)
	GetRepositoryByIDAndProject(ctx context.Context, arg GetRepositoryByIDAndProjectParams) (Repository, error)
	GetRepositoryByRepoID(ctx context.Context, repoID int64) (Repository, error)
	GetRepositoryByRepoName(ctx context.Context, arg GetRepositoryByRepoNameParams) (Repository, error)
	GetRuleInstancesEntityInProjects(ctx context.Context, arg GetRuleInstancesEntityInProjectsParams) ([]RuleInstance, error)
	GetRuleInstancesForProfile(ctx context.Context, profileID uuid.UUID) ([]RuleInstance, error)
	GetRuleTypeByID(ctx context.Context, id uuid.UUID) (RuleType, error)
	GetRuleTypeByName(ctx context.Context, arg GetRuleTypeByNameParams) (RuleType, error)
	// intended as a temporary transition query
	// this will be removed once rule_instances is used consistently in the engine
	GetRuleTypeIDByRuleNameEntityProfile(ctx context.Context, arg GetRuleTypeIDByRuleNameEntityProfileParams) (uuid.UUID, error)
	// intended as a temporary transition query
	// this will be removed once the evaluation history tables replace the old state tables
	GetRuleTypeNameByID(ctx context.Context, id uuid.UUID) (string, error)
	GetRuleTypesByEntityInHierarchy(ctx context.Context, arg GetRuleTypesByEntityInHierarchyParams) ([]RuleType, error)
	GetSelectorByID(ctx context.Context, id uuid.UUID) (ProfileSelector, error)
	GetSelectorsByProfileID(ctx context.Context, profileID uuid.UUID) ([]ProfileSelector, error)
	GetSubscriptionByProjectBundle(ctx context.Context, arg GetSubscriptionByProjectBundleParams) (Subscription, error)
	GetUnclaimedInstallationsByUser(ctx context.Context, ghID sql.NullString) ([]ProviderGithubAppInstallation, error)
	GetUserByID(ctx context.Context, id int32) (User, error)
	GetUserBySubject(ctx context.Context, identitySubject string) (User, error)
	GlobalListProviders(ctx context.Context) ([]Provider, error)
	GlobalListProvidersByClass(ctx context.Context, class ProviderClass) ([]Provider, error)
	InsertAlertEvent(ctx context.Context, arg InsertAlertEventParams) error
	InsertEvaluationRuleEntity(ctx context.Context, arg InsertEvaluationRuleEntityParams) (uuid.UUID, error)
	InsertEvaluationStatus(ctx context.Context, arg InsertEvaluationStatusParams) (uuid.UUID, error)
	InsertRemediationEvent(ctx context.Context, arg InsertRemediationEventParams) error
	ListArtifactsByRepoID(ctx context.Context, repositoryID uuid.NullUUID) ([]Artifact, error)
	ListEvaluationHistory(ctx context.Context, arg ListEvaluationHistoryParams) ([]ListEvaluationHistoryRow, error)
	ListEvaluationHistoryStaleRecords(ctx context.Context, arg ListEvaluationHistoryStaleRecordsParams) ([]ListEvaluationHistoryStaleRecordsRow, error)
	ListFlushCache(ctx context.Context) ([]FlushCache, error)
	// ListInvitationsForProject collects the information visible to project
	// administrators after an invitation has been issued.  In particular, it
	// *does not* report the invitation code, which is a secret intended for
	// the invitee.
	ListInvitationsForProject(ctx context.Context, project uuid.UUID) ([]ListInvitationsForProjectRow, error)
	// ListNonOrgProjects is a query that lists all non-organization projects.
	// projects have a boolean field is_organization that is set to true if the project is an organization.
	// this flag is no longer used and will be removed in the future.
	ListNonOrgProjects(ctx context.Context) ([]Project, error)
	// ListOrgProjects is a query that lists all organization projects.
	// projects have a boolean field is_organization that is set to true if the project is an organization.
	// this flag is no longer used and will be removed in the future.
	ListOldOrgProjects(ctx context.Context) ([]Project, error)
	// ListOldestRuleEvaluationsByRepositoryId has casts in select statement as sqlc generates incorrect types.
	// cast after MIN is required due to a known bug in sqlc: https://github.com/sqlc-dev/sqlc/issues/1965
	ListOldestRuleEvaluationsByRepositoryId(ctx context.Context, repositoryIds []uuid.UUID) ([]ListOldestRuleEvaluationsByRepositoryIdRow, error)
	ListProfilesByProjectIDAndLabel(ctx context.Context, arg ListProfilesByProjectIDAndLabelParams) ([]ListProfilesByProjectIDAndLabelRow, error)
	ListProfilesInstantiatingRuleType(ctx context.Context, ruleTypeID uuid.UUID) ([]string, error)
	// ListProvidersByProjectID allows us to list all providers
	// for a given array of projects.
	ListProvidersByProjectID(ctx context.Context, projects []uuid.UUID) ([]Provider, error)
	// ListProvidersByProjectIDPaginated allows us to lits all providers for a given project
	// with pagination taken into account. In this case, the cursor is the creation date.
	ListProvidersByProjectIDPaginated(ctx context.Context, arg ListProvidersByProjectIDPaginatedParams) ([]Provider, error)
	ListRegisteredRepositoriesByProjectIDAndProvider(ctx context.Context, arg ListRegisteredRepositoriesByProjectIDAndProviderParams) ([]Repository, error)
	ListRepositoriesAfterID(ctx context.Context, arg ListRepositoriesAfterIDParams) ([]Repository, error)
	ListRepositoriesByProjectID(ctx context.Context, arg ListRepositoriesByProjectIDParams) ([]Repository, error)
	ListRuleEvaluationsByProfileId(ctx context.Context, arg ListRuleEvaluationsByProfileIdParams) ([]ListRuleEvaluationsByProfileIdRow, error)
	ListRuleTypesByProject(ctx context.Context, projectID uuid.UUID) ([]RuleType, error)
	// When doing a key/algorithm rotation, identify the secrets which need to be
	// rotated. The criteria for rotation are:
	// 1) The encrypted_access_token is NULL (this should be removed when we make
	//    this column non-nullable).
	// 2) The access token does not use the configured default algorithm.
	// 3) The access token does not use the default key version.
	// This query accepts the default key version/algorithm as arguments since
	// that information is not known to the database.
	ListTokensToMigrate(ctx context.Context, arg ListTokensToMigrateParams) ([]ProviderAccessToken, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	// LockIfThresholdNotExceeded is used to lock an entity for execution. It will
	// attempt to insert or update the entity_execution_lock table only if the
	// last_lock_time is older than the threshold. If the lock is successful, it
	// will return the lock record. If the lock is unsuccessful, it will return
	// NULL.
	LockIfThresholdNotExceeded(ctx context.Context, arg LockIfThresholdNotExceededParams) (EntityExecutionLock, error)
	// OrphanProject is a query that sets the parent_id of a project to NULL.
	OrphanProject(ctx context.Context, arg OrphanProjectParams) (Project, error)
	// ReleaseLock is used to release a lock on an entity. It will delete the
	// entity_execution_lock record if the lock is held by the given locked_by
	// value.
	ReleaseLock(ctx context.Context, arg ReleaseLockParams) error
	RepositoryExistsAfterID(ctx context.Context, id uuid.UUID) (bool, error)
	SetCurrentVersion(ctx context.Context, arg SetCurrentVersionParams) error
	TemporaryPopulateArtifacts(ctx context.Context) error
	// TemporaryPopulateEvaluationHistory sets the entity_instance_id column for
	// all existing evaluation_rule_entities records to the id of the entity
	// instance that the rule entity is associated with. We derive this from the entity_type
	// and the corresponding entity id (repository_id, pull_request_id, or artifact_id).
	// Note that there are cases where repository_id and pull_request_id will both be set,
	// so we need to rely on the entity_type to determine which one to use. The same
	// applies to repository_id and artifact_id.
	TemporaryPopulateEvaluationHistory(ctx context.Context) error
	TemporaryPopulatePullRequests(ctx context.Context) error
	TemporaryPopulateRepositories(ctx context.Context) error
	UpdateEncryptedSecret(ctx context.Context, arg UpdateEncryptedSecretParams) error
	// UpdateInvitationRole updates an invitation by its code. This is intended to be
	// called by a user who has issued an invitation and then decided to change the
	// role of the invitee.
	UpdateInvitationRole(ctx context.Context, arg UpdateInvitationRoleParams) (UserInvite, error)
	UpdateLease(ctx context.Context, arg UpdateLeaseParams) error
	UpdateProfile(ctx context.Context, arg UpdateProfileParams) (Profile, error)
	UpdateProjectMeta(ctx context.Context, arg UpdateProjectMetaParams) (Project, error)
	UpdateProvider(ctx context.Context, arg UpdateProviderParams) error
	UpdateReminderLastSentById(ctx context.Context, id uuid.UUID) error
	UpdateRuleType(ctx context.Context, arg UpdateRuleTypeParams) (RuleType, error)
	UpdateSelector(ctx context.Context, arg UpdateSelectorParams) (ProfileSelector, error)
	UpsertAccessToken(ctx context.Context, arg UpsertAccessTokenParams) (ProviderAccessToken, error)
	UpsertArtifact(ctx context.Context, arg UpsertArtifactParams) (Artifact, error)
	// Copyright 2024 Stacklok, Inc
	//
	// Licensed under the Apache License, Version 2.0 (the "License");
	// you may not use this file except in compliance with the License.
	// You may obtain a copy of the License at
	//
	//      http://www.apache.org/licenses/LICENSE-2.0
	//
	// Unless required by applicable law or agreed to in writing, software
	// distributed under the License is distributed on an "AS IS" BASIS,
	// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	// See the License for the specific language governing permissions and
	// limitations under the License.
	// Bundles --
	UpsertBundle(ctx context.Context, arg UpsertBundleParams) error
	UpsertInstallationID(ctx context.Context, arg UpsertInstallationIDParams) (ProviderGithubAppInstallation, error)
	UpsertLatestEvaluationStatus(ctx context.Context, arg UpsertLatestEvaluationStatusParams) error
	UpsertProfileForEntity(ctx context.Context, arg UpsertProfileForEntityParams) (EntityProfile, error)
	UpsertPullRequest(ctx context.Context, arg UpsertPullRequestParams) (PullRequest, error)
	// Copyright 2024 Stacklok, Inc
	//
	// Licensed under the Apache License, Version 2.0 (the "License");
	// you may not use this file except in compliance with the License.
	// You may obtain a copy of the License at
	//
	//      http://www.apache.org/licenses/LICENSE-2.0
	//
	// Unless required by applicable law or agreed to in writing, software
	// distributed under the License is distributed on an "AS IS" BASIS,
	// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	// See the License for the specific language governing permissions and
	// limitations under the License.
	UpsertRuleInstance(ctx context.Context, arg UpsertRuleInstanceParams) (uuid.UUID, error)
}

var _ Querier = (*Queries)(nil)
