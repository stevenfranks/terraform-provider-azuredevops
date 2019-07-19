// Package foo implements the Azure ARM Foo service API version 5.1-preview.
//
//
package foo

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "context"
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "github.com/Azure/go-autorest/tracing"
    "github.com/satori/go.uuid"
    "net/http"
)

const (
// DefaultBaseURI is the default URI used for the service Foo
DefaultBaseURI = "https://dev.azure.com")

// BaseClient is the base client for Foo.
type BaseClient struct {
    autorest.Client
    BaseURI string
}

// New creates an instance of the BaseClient client.
func New()BaseClient {
    return NewWithBaseURI(DefaultBaseURI, )
}

// NewWithBaseURI creates an instance of the BaseClient client.
func NewWithBaseURI(baseURI string, ) BaseClient {
    return BaseClient{
        Client: autorest.NewClientWithUserAgent(UserAgent()),
        BaseURI: baseURI,
    }
}

    // Create queues a project to be created. Use the [GetOperation](../../operations/operations/get) to periodically check
    // for create project status.
        // Parameters:
            // organization - the name of the Azure DevOps organization.
            // body - the project to create.
    func (client BaseClient) Create(ctx context.Context, organization string, body TeamProject) (result OperationReference, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.Create")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
            req, err := client.CreatePreparer(ctx, organization, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "foo.BaseClient", "Create", nil , "Failure preparing request")
        return
        }

                resp, err := client.CreateSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "Create", resp, "Failure sending request")
                return
                }

                result, err = client.CreateResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "Create", resp, "Failure responding to request")
                }

        return
        }

        // CreatePreparer prepares the Create request.
        func (client BaseClient) CreatePreparer(ctx context.Context, organization string, body TeamProject) (*http.Request, error) {
                pathParameters := map[string]interface{} {
                "organization": autorest.Encode("path",organization),
                }

                            const APIVersion = "5.1-preview"
            queryParameters := map[string]interface{} {
            "api-version": APIVersion,
            }

            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPathParameters("/{organization}/_apis/projects",pathParameters),
        autorest.WithJSON(body),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // CreateSender sends the Create request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) CreateSender(req *http.Request) (*http.Response, error) {
            sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                return autorest.SendWithSender(client, req, sd...)
                }

    // CreateResponder handles the response to the Create request. The method always
    // closes the http.Response Body.
    func (client BaseClient) CreateResponder(resp *http.Response) (result OperationReference, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // Delete queues a project to be deleted. Use the [GetOperation](../../operations/operations/get) to periodically check
    // for delete project status.
        // Parameters:
            // organization - the name of the Azure DevOps organization.
            // projectID - the project id of the project to delete.
    func (client BaseClient) Delete(ctx context.Context, organization string, projectID uuid.UUID) (result OperationReference, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.Delete")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
            req, err := client.DeletePreparer(ctx, organization, projectID)
        if err != nil {
        err = autorest.NewErrorWithError(err, "foo.BaseClient", "Delete", nil , "Failure preparing request")
        return
        }

                resp, err := client.DeleteSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "Delete", resp, "Failure sending request")
                return
                }

                result, err = client.DeleteResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "Delete", resp, "Failure responding to request")
                }

        return
        }

        // DeletePreparer prepares the Delete request.
        func (client BaseClient) DeletePreparer(ctx context.Context, organization string, projectID uuid.UUID) (*http.Request, error) {
                pathParameters := map[string]interface{} {
                "organization": autorest.Encode("path",organization),
                "projectId": autorest.Encode("path",projectID),
                }

                            const APIVersion = "5.1-preview"
            queryParameters := map[string]interface{} {
            "api-version": APIVersion,
            }

            preparer := autorest.CreatePreparer(
        autorest.AsDelete(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPathParameters("/{organization}/_apis/projects/{projectId}",pathParameters),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // DeleteSender sends the Delete request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) DeleteSender(req *http.Request) (*http.Response, error) {
            sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                return autorest.SendWithSender(client, req, sd...)
                }

    // DeleteResponder handles the response to the Delete request. The method always
    // closes the http.Response Body.
    func (client BaseClient) DeleteResponder(resp *http.Response) (result OperationReference, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // Get get project with the specified id or name, optionally including capabilities.
        // Parameters:
            // organization - the name of the Azure DevOps organization.
            // includeCapabilities - include capabilities (such as source control) in the team project result (default:
            // false).
            // includeHistory - search within renamed projects (that had such name in the past).
    func (client BaseClient) Get(ctx context.Context, organization string, projectID string, includeCapabilities *bool, includeHistory *bool) (result TeamProject, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.Get")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
            req, err := client.GetPreparer(ctx, organization, projectID, includeCapabilities, includeHistory)
        if err != nil {
        err = autorest.NewErrorWithError(err, "foo.BaseClient", "Get", nil , "Failure preparing request")
        return
        }

                resp, err := client.GetSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "Get", resp, "Failure sending request")
                return
                }

                result, err = client.GetResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "Get", resp, "Failure responding to request")
                }

        return
        }

        // GetPreparer prepares the Get request.
        func (client BaseClient) GetPreparer(ctx context.Context, organization string, projectID string, includeCapabilities *bool, includeHistory *bool) (*http.Request, error) {
                pathParameters := map[string]interface{} {
                "organization": autorest.Encode("path",organization),
                "projectId": autorest.Encode("path",projectID),
                }

                            const APIVersion = "5.1-preview"
            queryParameters := map[string]interface{} {
            "api-version": APIVersion,
            }
                if includeCapabilities != nil {
                queryParameters["includeCapabilities"] = autorest.Encode("query",*includeCapabilities)
                }
                if includeHistory != nil {
                queryParameters["includeHistory"] = autorest.Encode("query",*includeHistory)
                }

            preparer := autorest.CreatePreparer(
        autorest.AsGet(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPathParameters("/{organization}/_apis/projects/{projectId}",pathParameters),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // GetSender sends the Get request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) GetSender(req *http.Request) (*http.Response, error) {
            sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                return autorest.SendWithSender(client, req, sd...)
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client BaseClient) GetResponder(resp *http.Response) (result TeamProject, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // GetProjectProperties get a collection of team project properties.
        // Parameters:
            // organization - the name of the Azure DevOps organization.
            // projectID - the team project ID.
            // keys - a comma-delimited string of team project property names. Wildcard characters ("?" and "*") are
            // supported. If no key is specified, all properties will be returned.
    func (client BaseClient) GetProjectProperties(ctx context.Context, organization string, projectID uuid.UUID, keys string) (result ListProjectProperty, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.GetProjectProperties")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
            req, err := client.GetProjectPropertiesPreparer(ctx, organization, projectID, keys)
        if err != nil {
        err = autorest.NewErrorWithError(err, "foo.BaseClient", "GetProjectProperties", nil , "Failure preparing request")
        return
        }

                resp, err := client.GetProjectPropertiesSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "GetProjectProperties", resp, "Failure sending request")
                return
                }

                result, err = client.GetProjectPropertiesResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "GetProjectProperties", resp, "Failure responding to request")
                }

        return
        }

        // GetProjectPropertiesPreparer prepares the GetProjectProperties request.
        func (client BaseClient) GetProjectPropertiesPreparer(ctx context.Context, organization string, projectID uuid.UUID, keys string) (*http.Request, error) {
                pathParameters := map[string]interface{} {
                "organization": autorest.Encode("path",organization),
                "projectId": autorest.Encode("path",projectID),
                }

                            const APIVersion = "5.1-preview"
            queryParameters := map[string]interface{} {
            "api-version": APIVersion,
            }
                if len(keys) > 0 {
                queryParameters["keys"] = autorest.Encode("query",keys)
                }

            preparer := autorest.CreatePreparer(
        autorest.AsGet(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPathParameters("/{organization}/_apis/projects/{projectId}/properties",pathParameters),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // GetProjectPropertiesSender sends the GetProjectProperties request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) GetProjectPropertiesSender(req *http.Request) (*http.Response, error) {
            sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                return autorest.SendWithSender(client, req, sd...)
                }

    // GetProjectPropertiesResponder handles the response to the GetProjectProperties request. The method always
    // closes the http.Response Body.
    func (client BaseClient) GetProjectPropertiesResponder(resp *http.Response) (result ListProjectProperty, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result.Value),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // List get all projects in the organization that the authenticated user has access to.
        // Parameters:
            // organization - the name of the Azure DevOps organization.
            // stateFilter - filter on team projects in a specific team project state (default: WellFormed).
    func (client BaseClient) List(ctx context.Context, organization string, stateFilter ProjectState, top *int32, skip *int32, continuationToken string, getDefaultTeamImageURL *bool) (result ListTeamProjectReference, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.List")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
            req, err := client.ListPreparer(ctx, organization, stateFilter, top, skip, continuationToken, getDefaultTeamImageURL)
        if err != nil {
        err = autorest.NewErrorWithError(err, "foo.BaseClient", "List", nil , "Failure preparing request")
        return
        }

                resp, err := client.ListSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "List", resp, "Failure sending request")
                return
                }

                result, err = client.ListResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "List", resp, "Failure responding to request")
                }

        return
        }

        // ListPreparer prepares the List request.
        func (client BaseClient) ListPreparer(ctx context.Context, organization string, stateFilter ProjectState, top *int32, skip *int32, continuationToken string, getDefaultTeamImageURL *bool) (*http.Request, error) {
                pathParameters := map[string]interface{} {
                "organization": autorest.Encode("path",organization),
                }

                            const APIVersion = "5.1-preview"
            queryParameters := map[string]interface{} {
            "api-version": APIVersion,
            }
                if len(string(stateFilter)) > 0 {
                queryParameters["stateFilter"] = autorest.Encode("query",stateFilter)
                }
                if top != nil {
                queryParameters["$top"] = autorest.Encode("query",*top)
                }
                if skip != nil {
                queryParameters["$skip"] = autorest.Encode("query",*skip)
                }
                if len(continuationToken) > 0 {
                queryParameters["continuationToken"] = autorest.Encode("query",continuationToken)
                }
                if getDefaultTeamImageURL != nil {
                queryParameters["getDefaultTeamImageUrl"] = autorest.Encode("query",*getDefaultTeamImageURL)
                }

            preparer := autorest.CreatePreparer(
        autorest.AsGet(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPathParameters("/{organization}/_apis/projects",pathParameters),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // ListSender sends the List request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) ListSender(req *http.Request) (*http.Response, error) {
            sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                return autorest.SendWithSender(client, req, sd...)
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client BaseClient) ListResponder(resp *http.Response) (result ListTeamProjectReference, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result.Value),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // RemoveProjectAvatar removes the avatar for the project.
        // Parameters:
            // organization - the name of the Azure DevOps organization.
            // projectID - the ID or name of the project.
    func (client BaseClient) RemoveProjectAvatar(ctx context.Context, organization string, projectID string) (result autorest.Response, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.RemoveProjectAvatar")
            defer func() {
                sc := -1
                if result.Response != nil {
                    sc = result.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
            req, err := client.RemoveProjectAvatarPreparer(ctx, organization, projectID)
        if err != nil {
        err = autorest.NewErrorWithError(err, "foo.BaseClient", "RemoveProjectAvatar", nil , "Failure preparing request")
        return
        }

                resp, err := client.RemoveProjectAvatarSender(req)
                if err != nil {
                result.Response = resp
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "RemoveProjectAvatar", resp, "Failure sending request")
                return
                }

                result, err = client.RemoveProjectAvatarResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "RemoveProjectAvatar", resp, "Failure responding to request")
                }

        return
        }

        // RemoveProjectAvatarPreparer prepares the RemoveProjectAvatar request.
        func (client BaseClient) RemoveProjectAvatarPreparer(ctx context.Context, organization string, projectID string) (*http.Request, error) {
                pathParameters := map[string]interface{} {
                "organization": autorest.Encode("path",organization),
                "projectId": autorest.Encode("path",projectID),
                }

                            const APIVersion = "5.1-preview"
            queryParameters := map[string]interface{} {
            "api-version": APIVersion,
            }

            preparer := autorest.CreatePreparer(
        autorest.AsDelete(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPathParameters("/{organization}/_apis/projects/{projectId}/avatar",pathParameters),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // RemoveProjectAvatarSender sends the RemoveProjectAvatar request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) RemoveProjectAvatarSender(req *http.Request) (*http.Response, error) {
            sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                return autorest.SendWithSender(client, req, sd...)
                }

    // RemoveProjectAvatarResponder handles the response to the RemoveProjectAvatar request. The method always
    // closes the http.Response Body.
    func (client BaseClient) RemoveProjectAvatarResponder(resp *http.Response) (result autorest.Response, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByClosing())
        result.Response = resp
            return
        }

    // SetProjectAvatar sets the avatar for the project.
        // Parameters:
            // organization - the name of the Azure DevOps organization.
            // body - the avatar blob data object to upload.
            // projectID - the ID or name of the project.
    func (client BaseClient) SetProjectAvatar(ctx context.Context, organization string, body ProjectAvatar, projectID string) (result autorest.Response, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.SetProjectAvatar")
            defer func() {
                sc := -1
                if result.Response != nil {
                    sc = result.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
            req, err := client.SetProjectAvatarPreparer(ctx, organization, body, projectID)
        if err != nil {
        err = autorest.NewErrorWithError(err, "foo.BaseClient", "SetProjectAvatar", nil , "Failure preparing request")
        return
        }

                resp, err := client.SetProjectAvatarSender(req)
                if err != nil {
                result.Response = resp
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "SetProjectAvatar", resp, "Failure sending request")
                return
                }

                result, err = client.SetProjectAvatarResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "SetProjectAvatar", resp, "Failure responding to request")
                }

        return
        }

        // SetProjectAvatarPreparer prepares the SetProjectAvatar request.
        func (client BaseClient) SetProjectAvatarPreparer(ctx context.Context, organization string, body ProjectAvatar, projectID string) (*http.Request, error) {
                pathParameters := map[string]interface{} {
                "organization": autorest.Encode("path",organization),
                "projectId": autorest.Encode("path",projectID),
                }

                            const APIVersion = "5.1-preview"
            queryParameters := map[string]interface{} {
            "api-version": APIVersion,
            }

            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPut(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPathParameters("/{organization}/_apis/projects/{projectId}/avatar",pathParameters),
        autorest.WithJSON(body),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // SetProjectAvatarSender sends the SetProjectAvatar request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) SetProjectAvatarSender(req *http.Request) (*http.Response, error) {
            sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                return autorest.SendWithSender(client, req, sd...)
                }

    // SetProjectAvatarResponder handles the response to the SetProjectAvatar request. The method always
    // closes the http.Response Body.
    func (client BaseClient) SetProjectAvatarResponder(resp *http.Response) (result autorest.Response, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByClosing())
        result.Response = resp
            return
        }

    // SetProjectProperties create, update, and delete team project properties.
        // Parameters:
            // organization - the name of the Azure DevOps organization.
            // projectID - the team project ID.
            // body - a JSON Patch document that represents an array of property operations. See RFC 6902 for more details
            // on JSON Patch. The accepted operation verbs are Add and Remove, where Add is used for both creating and
            // updating properties. The path consists of a forward slash and a property name.
    func (client BaseClient) SetProjectProperties(ctx context.Context, organization string, projectID uuid.UUID, body interface{}) (result autorest.Response, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.SetProjectProperties")
            defer func() {
                sc := -1
                if result.Response != nil {
                    sc = result.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
            req, err := client.SetProjectPropertiesPreparer(ctx, organization, projectID, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "foo.BaseClient", "SetProjectProperties", nil , "Failure preparing request")
        return
        }

                resp, err := client.SetProjectPropertiesSender(req)
                if err != nil {
                result.Response = resp
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "SetProjectProperties", resp, "Failure sending request")
                return
                }

                result, err = client.SetProjectPropertiesResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "SetProjectProperties", resp, "Failure responding to request")
                }

        return
        }

        // SetProjectPropertiesPreparer prepares the SetProjectProperties request.
        func (client BaseClient) SetProjectPropertiesPreparer(ctx context.Context, organization string, projectID uuid.UUID, body interface{}) (*http.Request, error) {
                pathParameters := map[string]interface{} {
                "organization": autorest.Encode("path",organization),
                "projectId": autorest.Encode("path",projectID),
                }

                            const APIVersion = "5.1-preview"
            queryParameters := map[string]interface{} {
            "api-version": APIVersion,
            }

            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json-patch+json; charset=utf-8"),
        autorest.AsPatch(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPathParameters("/{organization}/_apis/projects/{projectId}/properties",pathParameters),
        autorest.WithJSON(body),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // SetProjectPropertiesSender sends the SetProjectProperties request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) SetProjectPropertiesSender(req *http.Request) (*http.Response, error) {
            sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                return autorest.SendWithSender(client, req, sd...)
                }

    // SetProjectPropertiesResponder handles the response to the SetProjectProperties request. The method always
    // closes the http.Response Body.
    func (client BaseClient) SetProjectPropertiesResponder(resp *http.Response) (result autorest.Response, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByClosing())
        result.Response = resp
            return
        }

    // Update update an existing project's name, abbreviation, description, or restore a project.
        // Parameters:
            // organization - the name of the Azure DevOps organization.
            // body - the updates for the project. The state must be set to wellFormed to restore the project.
            // projectID - the project id of the project to update.
    func (client BaseClient) Update(ctx context.Context, organization string, body TeamProject, projectID uuid.UUID) (result OperationReference, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.Update")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
            req, err := client.UpdatePreparer(ctx, organization, body, projectID)
        if err != nil {
        err = autorest.NewErrorWithError(err, "foo.BaseClient", "Update", nil , "Failure preparing request")
        return
        }

                resp, err := client.UpdateSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "Update", resp, "Failure sending request")
                return
                }

                result, err = client.UpdateResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "foo.BaseClient", "Update", resp, "Failure responding to request")
                }

        return
        }

        // UpdatePreparer prepares the Update request.
        func (client BaseClient) UpdatePreparer(ctx context.Context, organization string, body TeamProject, projectID uuid.UUID) (*http.Request, error) {
                pathParameters := map[string]interface{} {
                "organization": autorest.Encode("path",organization),
                "projectId": autorest.Encode("path",projectID),
                }

                            const APIVersion = "5.1-preview"
            queryParameters := map[string]interface{} {
            "api-version": APIVersion,
            }

            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPatch(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPathParameters("/{organization}/_apis/projects/{projectId}",pathParameters),
        autorest.WithJSON(body),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // UpdateSender sends the Update request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) UpdateSender(req *http.Request) (*http.Response, error) {
            sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                return autorest.SendWithSender(client, req, sd...)
                }

    // UpdateResponder handles the response to the Update request. The method always
    // closes the http.Response Body.
    func (client BaseClient) UpdateResponder(resp *http.Response) (result OperationReference, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

