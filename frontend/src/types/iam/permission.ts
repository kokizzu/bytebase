import type { PickLiteral } from "../utils";

export type Permission =
  | "bb.auditLogs.export"
  | "bb.auditLogs.search"
  | "bb.branches.admin"
  | "bb.branches.create"
  | "bb.branches.delete"
  | "bb.branches.get"
  | "bb.branches.list"
  | "bb.branches.update"
  | "bb.changeHistories.get"
  | "bb.changeHistories.list"
  | "bb.vcsConnectors.create"
  | "bb.vcsConnectors.delete"
  | "bb.vcsConnectors.get"
  | "bb.vcsConnectors.list"
  | "bb.vcsConnectors.update"
  | "bb.changelists.create"
  | "bb.changelists.delete"
  | "bb.changelists.get"
  | "bb.changelists.list"
  | "bb.changelists.update"
  | "bb.databaseSecrets.delete"
  | "bb.databaseSecrets.list"
  | "bb.databaseSecrets.update"
  | "bb.databases.adviseIndex"
  | "bb.databases.check"
  | "bb.databases.get"
  | "bb.databases.getSchema"
  | "bb.databases.list"
  | "bb.databases.sync"
  | "bb.databases.update"
  | "bb.issueComments.list"
  | "bb.issueComments.create"
  | "bb.issueComments.update"
  | "bb.issues.create"
  | "bb.issues.get"
  | "bb.issues.list"
  | "bb.issues.update"
  | "bb.planCheckRuns.list"
  | "bb.planCheckRuns.run"
  | "bb.plans.create"
  | "bb.plans.get"
  | "bb.plans.list"
  | "bb.plans.update"
  | "bb.projects.get"
  | "bb.projects.getIamPolicy"
  | "bb.projects.setIamPolicy"
  | "bb.projects.update"
  | "bb.rollouts.create"
  | "bb.rollouts.get"
  | "bb.rollouts.list"
  | "bb.rollouts.preview"
  | "bb.slowQueries.list"
  | "bb.sql.select"
  | "bb.sql.info"
  | "bb.sql.explain"
  | "bb.sql.dml"
  | "bb.sql.ddl"
  | "bb.sql.export"
  | "bb.taskRuns.create"
  | "bb.taskRuns.list"
  | "bb.environments.create"
  | "bb.environments.delete"
  | "bb.environments.get"
  | "bb.environments.list"
  | "bb.environments.undelete"
  | "bb.environments.update"
  | "bb.vcsProviders.create"
  | "bb.vcsProviders.delete"
  | "bb.vcsProviders.get"
  | "bb.vcsProviders.list"
  | "bb.vcsProviders.listProjects"
  | "bb.vcsProviders.searchProjects"
  | "bb.vcsProviders.update"
  | "bb.identityProviders.create"
  | "bb.identityProviders.delete"
  | "bb.identityProviders.get"
  | "bb.identityProviders.undelete"
  | "bb.identityProviders.update"
  | "bb.sql.admin"
  | "bb.instances.create"
  | "bb.instances.delete"
  | "bb.instances.get"
  | "bb.instances.list"
  | "bb.instances.sync"
  | "bb.instances.undelete"
  | "bb.instances.update"
  | "bb.policies.create"
  | "bb.policies.delete"
  | "bb.policies.get"
  | "bb.policies.list"
  | "bb.policies.update"
  | "bb.projects.create"
  | "bb.projects.delete"
  | "bb.projects.list"
  | "bb.projects.undelete"
  | "bb.risks.create"
  | "bb.risks.delete"
  | "bb.risks.list"
  | "bb.risks.update"
  | "bb.roles.create"
  | "bb.roles.delete"
  | "bb.roles.list"
  | "bb.roles.get"
  | "bb.roles.update"
  | "bb.groups.create"
  | "bb.groups.update"
  | "bb.groups.delete"
  | "bb.groups.list"
  | "bb.groups.get"
  | "bb.users.create"
  | "bb.users.update"
  | "bb.users.delete"
  | "bb.users.undelete"
  | "bb.settings.get"
  | "bb.settings.list"
  | "bb.settings.set"
  | "bb.worksheets.manage"
  | "bb.worksheets.get";

export type QueryPermission = PickLiteral<
  Permission,
  | "bb.sql.select"
  | "bb.sql.info"
  | "bb.sql.explain"
  | "bb.sql.ddl"
  | "bb.sql.dml"
  | "bb.sql.admin"
>;

export const QueryPermissionQueryOnly: QueryPermission[] = ["bb.sql.select"];
export const QueryPermissionQueryAny: QueryPermission[] = [
  "bb.sql.select",
  "bb.sql.info",
  "bb.sql.explain",
  "bb.sql.ddl",
  "bb.sql.dml",
  "bb.sql.admin",
];
