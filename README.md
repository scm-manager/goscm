# goscm

Goscm is a custom webhook for SCM-Manager support in Argo CD. It is based on the [Library webhooks project](https://github.com/go-playground/webhooks/).

## Local development

If you want to use your local Goscm version in Argo CD development, you need to apply this comment within the Argo CD repository:

```
go mod edit -replace=github.com/scm-manager/goscm@<INSERT_VERSION>=<PATH>
```

First, replace *<INSERT_VERSION>* with the version listed in the *go.mod* file in the upper command (e.g. v0.0.6). 
Then replace *<RELATIVE_PATH>* with its path (e.g. */home/someSubfolder/goscm*).

**Keep in mind to exclude the changed go.mod file from your Argo CD commit!** Otherwise, it is going to cause a build fail on other systems.

## Notes

OpenAPI Spec: https://ecosystem.cloudogu.com/scm/openapi

## Tests

In order for tests to run, a valid API key must be supplied via the environment variable `SCM_BEARER_TOKEN`.

Tests are run utilizing: https://stagex.cloudogu.com/scm

## Disclaimer
This go client for SCM-Manager is not complete and therefore does not support all features from SCM-Manager. 
It simply provides some basic functions which can be used by third party applications. 

The client is in a beta stage. It is possible that the API will change.