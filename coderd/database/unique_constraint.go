// Code generated by gen/enum. DO NOT EDIT.
package database

// UniqueConstraint represents a named unique constraint on a table.
type UniqueConstraint string

// UniqueConstraint enums.
const (
	UniqueLicensesIDKey                            UniqueConstraint = "licenses_id_key"                                // ALTER TABLE ONLY licenses ADD CONSTRAINT licenses_id_key UNIQUE (id);
	UniqueLicensesJWTKey                           UniqueConstraint = "licenses_jwt_key"                               // ALTER TABLE ONLY licenses ADD CONSTRAINT licenses_jwt_key UNIQUE (jwt);
	UniqueParameterSchemasJobIDNameKey             UniqueConstraint = "parameter_schemas_job_id_name_key"              // ALTER TABLE ONLY parameter_schemas ADD CONSTRAINT parameter_schemas_job_id_name_key UNIQUE (job_id, name);
	UniqueParameterValuesScopeIDNameKey            UniqueConstraint = "parameter_values_scope_id_name_key"             // ALTER TABLE ONLY parameter_values ADD CONSTRAINT parameter_values_scope_id_name_key UNIQUE (scope_id, name);
	UniqueProvisionerDaemonsNameKey                UniqueConstraint = "provisioner_daemons_name_key"                   // ALTER TABLE ONLY provisioner_daemons ADD CONSTRAINT provisioner_daemons_name_key UNIQUE (name);
	UniqueSiteConfigsKeyKey                        UniqueConstraint = "site_configs_key_key"                           // ALTER TABLE ONLY site_configs ADD CONSTRAINT site_configs_key_key UNIQUE (key);
	UniqueTemplateVersionsTemplateIDNameKey        UniqueConstraint = "template_versions_template_id_name_key"         // ALTER TABLE ONLY template_versions ADD CONSTRAINT template_versions_template_id_name_key UNIQUE (template_id, name);
	UniqueWorkspaceAppsAgentIDNameKey              UniqueConstraint = "workspace_apps_agent_id_name_key"               // ALTER TABLE ONLY workspace_apps ADD CONSTRAINT workspace_apps_agent_id_name_key UNIQUE (agent_id, name);
	UniqueWorkspaceBuildsJobIDKey                  UniqueConstraint = "workspace_builds_job_id_key"                    // ALTER TABLE ONLY workspace_builds ADD CONSTRAINT workspace_builds_job_id_key UNIQUE (job_id);
	UniqueWorkspaceBuildsWorkspaceIDBuildNumberKey UniqueConstraint = "workspace_builds_workspace_id_build_number_key" // ALTER TABLE ONLY workspace_builds ADD CONSTRAINT workspace_builds_workspace_id_build_number_key UNIQUE (workspace_id, build_number);
	UniqueIndexOrganizationName                    UniqueConstraint = "idx_organization_name"                          // CREATE UNIQUE INDEX idx_organization_name ON organizations USING btree (name);
	UniqueIndexOrganizationNameLower               UniqueConstraint = "idx_organization_name_lower"                    // CREATE UNIQUE INDEX idx_organization_name_lower ON organizations USING btree (lower(name));
	UniqueIndexUsersEmail                          UniqueConstraint = "idx_users_email"                                // CREATE UNIQUE INDEX idx_users_email ON users USING btree (email) WHERE (deleted = false);
	UniqueIndexUsersUsername                       UniqueConstraint = "idx_users_username"                             // CREATE UNIQUE INDEX idx_users_username ON users USING btree (username) WHERE (deleted = false);
	UniqueTemplatesOrganizationIDNameIndex         UniqueConstraint = "templates_organization_id_name_idx"             // CREATE UNIQUE INDEX templates_organization_id_name_idx ON templates USING btree (organization_id, lower((name)::text)) WHERE (deleted = false);
	UniqueUsersEmailLowerIndex                     UniqueConstraint = "users_email_lower_idx"                          // CREATE UNIQUE INDEX users_email_lower_idx ON users USING btree (lower(email)) WHERE (deleted = false);
	UniqueUsersUsernameLowerIndex                  UniqueConstraint = "users_username_lower_idx"                       // CREATE UNIQUE INDEX users_username_lower_idx ON users USING btree (lower(username)) WHERE (deleted = false);
	UniqueWorkspacesOwnerIDLowerIndex              UniqueConstraint = "workspaces_owner_id_lower_idx"                  // CREATE UNIQUE INDEX workspaces_owner_id_lower_idx ON workspaces USING btree (owner_id, lower((name)::text)) WHERE (deleted = false);
)
