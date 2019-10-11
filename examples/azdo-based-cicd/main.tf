
provider "azuredevops" {
  version = ">= 0.0.1"
}


// defines the project in AzDO
resource "azuredevops_project" "project" {
  project_name       = "terraform-provider-azuredevops"
  description        = ""
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
}

// defines a Git repository in the project in AzDO
resource "azuredevops_git_repo" "repository" {
  project_id = azuredevops_project.project.id
  name       = "my-cool-repository"
}

// defines an ARM service connection
resource "azuredevops_serviceendpoint" "arm" {
  project_id                 = azuredevops_project.project.id
  service_endpoint_name      = "ARM Service Connection"
  service_endpoint_type      = "arm"
  service_principal_username = var.service_principal_username
  service_principal_password = var.service_principal_password
  subscription_id            = var.subscription_id
  tenant_id                  = var.tenant_id
}

// defines variable groups that can be used for provisioning
resource "azuredevops_variable_group" "group1" {
    project_id = azuredevops_project.project.id
    name       = "My First Variable Group"
    variables  = {
      "VARIABLE_1" = "VALUE_1"
      "VARIABLE_2" = "VALUE_2"
    }
}

// defines a variable group that can be used for provisioning
resource "azuredevops_variable_group" "group2" {
    project_id = azuredevops_project.project.id
    name       = "My Second Variable Group"
    variables  = {
      "VARIABLE_3" = "VALUE_3"
    }
}

// A build that kicks off an infrastructure provisioning pipeline, using variable groups
resource "azuredevops_build_definition" "infra" {
  project_id            = azuredevops_project.project.id
  agent_pool_name       = "Hosted Ubuntu 1604"
  name                  = "Infrastructure Pipeilne"
  service_connection_id = azuredevops_serviceendpoint.arm.id
  variables_groups      = [
    azuredevops_variable_group.group1.id,
    azuredevops_variable_group.group2.id
  ]

  repository {
    repo_type   = azuredevops_git_repo.type
    repo_name   = azuredevops_git_repo.name
    branch_name = azuredevops_git_repo.default_branch
    yml_path    = "cicd/azure-pipelines-infra.yml"
  }
}



// note: see the types of branch policies here:
//   https://dev.azure.com/$ORG/$PROJECT/_apis/policy/types?api-version=5.0

// Look up the ID of the branch policy that ensures each PR is linked to a work item, then
// configure the policy
data "azuredevops_branch_policy_type" "work_item_policy" {
  name = "Work item linking"
}
resource "azuredevops_branch_policy" "work_item_linked_policy" {
  project_id    = azuredevops_project.project.id
  type          = azuredevops_branch_policy_type.work_item_policy.id
  repository_id = azuredevops_git_repo.repository.id
  branch        = azuredevops_git_repo.repository.default_branch
}



// Look up the ID of the branch policy that ensures each PR reviewed by 2 people, then
// configure the policy
data "azuredevops_branch_policy_type" "min_reviewer_count" {
  name = "Minimum number of reviewers"
}
resource "azuredevops_branch_policy" "min_reviewer_count_policy" {
  project_id    = azuredevops_project.project.id
  type          = azuredevops_branch_policy_type.min_reviewer_count.id
  repository_id = azuredevops_git_repo.repository.id
  branch        = azuredevops_git_repo.repository.default_branch
  settings = {
    "minimumApproverCount" = 2
  }
}



// Look up the ID of the branch policy that ensures each PR does not break a build, then
// configure the policy
data "azuredevops_branch_policy_type" "build" {
  name = "Minimum number of reviewers"
}
resource "azuredevops_branch_policy" "build_policy" {
  project_id    = azuredevops_project.project.id
  type          = azuredevops_branch_policy_type.build.id
  repository_id = azuredevops_git_repo.repository.id
  branch        = azuredevops_git_repo.repository.default_branch
  settings = {
    "buildDefinitionId" = azuredevops_build_definition.infra.id
  }
}

// group assignment....