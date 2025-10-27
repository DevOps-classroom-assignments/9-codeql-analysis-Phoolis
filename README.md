# Go Web Server – CodeQL - CI

This Go web server (main.go) exposes two HTTP endpoints:
1. **/readfile**
- Accepts a query parameter `file` (e.g., `/readfile?file=hello.txt`).
- Reads the contents of the specified file and returns it in the HTTP response.

2. **/exec**
- Accepts a query parameter cmd (e.g. `/exec?cmd=ls`).
- Executes the provided shell command.

> [!WARNING]
>  
> This repository contains code for DevOps CodeQL exercises. Executing this code directly on your local system may be dangerous and could pose security risks.
>   
> **It is strongly recommended to run all code and workflows only inside a Docker container or other isolated environment.**  
>
> Do not run this code outside of controlled, sandboxed environments.

## Steps

### Step 1. Test web server locally

Build the container and run it locally.

```bash
docker build -t go-web-server .
docker run -p 8080:8080 -it --rm go-web-server
```
After container is running, test the `/readfile` endpoint using the following curl command:

```bash
curl "http://localhost:8080/readfile?file=hello.txt"
```

You should see the content of the `safe-files/hello.txt` in the respone.

> [!WARNING]
> What if the user adds a path in the filename? An attacker might try using the following `file` attributes:
>
> ```
> curl "http://localhost:8080/readfile?file=../main.go"
> curl "http://localhost:8080/readfile?file=../../etc/shadow"
> ```

### Step 2. Implement CodeQL - CI workflow

> **Note:**  
> Temporarily set the repository visibility to public to allow CodeQL analysis to run in the workflow.

Modify the existing `codeql.yml` file.

- Implement a workflow tha use CodeQL to detect vulnerabilities.
- Implement a CI workflow to ensure the container builds and runs successfully

You do not need to know the Go programming language to complete this assignment. However, you should review any CodeQL warnings and make sure you understand what they mean.

## About the exercise
This exercise has been created by Juha Hinkula and is licensed under the Creative Commons BY-NC-SA license.

AI tools such as ChatGPT and GitHub Copilot have been used in the implementation of the task description, source code, data files and tests.
