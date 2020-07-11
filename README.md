# github-list-repos
Example of using [go-github](https://github.com/google/go-github) to output a list of repositories for an enterprise github instance

Lists all of the repositories that the authorized user has access to.<br><br>
The following environment variables must be set:

**GITHUB_BASE_URL** - base URL to the target enterprise github instance, 
example: https://github.my-company.com/

**GITHUB_AUTH_TOKEN** - authentication token to the target github instance
