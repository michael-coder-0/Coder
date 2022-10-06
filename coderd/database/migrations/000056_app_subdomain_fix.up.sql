-- There was a mistake in the last migration which set "subdomain" to be the
-- opposite of the deprecated value "relative_path", however the "relative_path"
-- value may not have been correct as it was not consumed anywhere prior to this
-- point.
--
-- Force all workspace apps to use path based routing until rebuild. This should
-- not impact any existing workspaces as the only supported routing method has
-- been path based routing prior to this point.
--
-- On rebuild the value from the Terraform template will be used instead
-- (defaulting to false if unspecified).
UPDATE "workspace_apps" SET "subdomain" = false;