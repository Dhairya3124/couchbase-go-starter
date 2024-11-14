<!---
This is a sample README.md file for a Couchbase Starter Kit. It includes a template for the README file that you can use in your repository. You can copy the contents of this file and replace the placeholders with the appropriate information for your starter kit.
-->

# Couchbase [Golang] Starter Kit

![Test Suite](https://github.com/codercatdev/couchbase-GO-starter/actions/workflows/test-connection.yml/badge.svg)
![Couchbase Capella](https://img.shields.io/badge/Couchbase_Capella-Enabled-red)
[![License: MIT](https://cdn.prod.website-files.com/5e0f1144930a8bc8aace526c/65dd9eb5aaca434fac4f1c34_License-MIT-blue.svg)](/LICENSE)
![Static Badge](https://img.shields.io/badge/Code_of_Conduct-Contributor_Covenant-violet.svg)

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/couchbase-starter-kit/URL)
[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/couchbase-starter-kit/URL)

## Configuration

| Variable Name                      | Description                                                 |      Default value       |
|:-----------------------------------|:------------------------------------------------------------|:------------------------:|
| GO_COUCHBASE_CONNECTION_STRING     | A couchbase connection string                               |            -             |
| GO_COUCHBASE_USERNAME              | Username for authentication with Couchbase                  |            -             |
| GO_COUCHBASE_PASSWORD              | Password for authentication with Couchbase                  |            -             |
| COUCHBASE_USE_CAPELLA              | Use to change the connection profile                        |          false           |
| COUCHBASE_DEFAULT_BUCKET           | The name of the Couchbase Bucket, parent of the scope       |         default          |
| COUCHBASE_DEFAULT_SCOPE            | The name of the Couchbase scope, parent of the collection   |         _default         |
| COUCHBASE_DEFAULT_COLLECTION       | The name of the Couchbase collection to store the Documents |         _default         |
| COUCHBASE_OTLP_ENABLED             | Enable traces and metrics OTLP export                       |          false           |
| COUCHBASE_OTLP_ENDPOINT            | The OTLP server endpoint to send metrics and traces         |            -             |

### Getting started

Follow these steps to get the starter kit up and running on your local machine.

**Prerequisites**

Make sure you have the following installed:

* Golang(version 1.16 or higher)

You will also need an account on Couchbase Capella to create a cluster and obtain the connection details.

**Clone the repository**

```bash
git clone https://github.com/Dhairya3124/couchbase-go-starter.git
cd URL
```
