# Contributing to LabData TCC

Thank you for your interest in contributing to the labdata-tcc project! I welcome contributions from the community and am excited to see what you'll bring to the project. This document outlines the guidelines for contributing to the project.

## Getting Started

### Fork the Repository

1. Fork the repository by clicking the "Fork" button on the top right of the repository page.
2. Clone your fork to your local machine:

   ```sh
   git clone https://github.com/<your-username>/labdata-tcc.git
   cd labdata-tcc
   ```

3. Add the upstream repository:

   ```sh
   git remote add upstream https://github.com/FabioCaffarello/labdata-tcc.git
   ```

### Setting Up the Development Environment

1. **Running Locally**

   - **Prerequisites**

     Ensure you have the required software installed:
     - [Python 3.10](https://www.python.org/downloads/)
       - [Poetry](https://pypi.org/project/poetry/1.8.3/)
         ```shell
         pip install poetry==1.8.3
         ```
     - [Golang 1.22.x](https://golang.google.cn/)
       - [Wire](https://pkg.go.dev/github.com/google/wire)
         ```shell
         go install github.com/google/wire/cmd/wire@latest 
         ```
       - [gomarkdoc](https://github.com/princjef/gomarkdoc)
         ```shell
         go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
         ```
     - [Node.js 22.x](https://nodejs.org/en/download/)
     - [Docker 26.1.1](https://www.docker.com/)

   - **Set Up the Environment**

     - Install NodeJS dependencies:
       ```sh
       npm install
       ```

     - Install Python dependencies:
       ```sh
       poetry install
       ```

2. **Running in a Devcontainer Locally**

   Ensure you have Docker and Visual Studio Code installed with the Remote - Containers extension. Open the repository in Visual Studio Code and click on "Reopen in Container" when prompted. The environment will be set up automatically.

3. **Running in GitHub Codespaces**

   You can also run the project in GitHub Codespaces directly. Navigate to the repository on GitHub, click the "Code" button, and select "Open with Codespaces" to create a new codespace. The environment will be set up automatically.

- **Set Up Husky**

    If you choose to run the project locally, you must set up Husky for managing Git hooks:

    ```sh
    npm prepare
    ```


### Commit Message Linting

I use commit message linting to maintain a consistent commit history. Please follow the commit message format:

```sh
git commit -m "type(scope): Description of your changes"
```

For example:

```sh
git commit -m "feat(api): Add new endpoint for data retrieval"
```

The available types are:
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `perf`: A code change that improves performance
- `test`: Adding missing or correcting existing tests
- `chore`: Changes to the build process or auxiliary tools and libraries such as documentation generation

## Making Changes

1. **Create a Branch**

   Create a new branch for your feature or bugfix:

   ```sh
   git checkout -b <branch-name>
   ```

2. **Make Your Changes**

   Implement your changes. Ensure your code follows the project's coding standards and includes appropriate tests.

3. **Commit Your Changes**

   Commit your changes with a clear and descriptive commit message following the commit message linting guidelines:

   ```sh
   git add .
   git commit -m "type(scope): Description of your changes"
   ```

4. **Push Your Changes**

   Push your changes to your forked repository:

   ```sh
   git push origin <branch-name>
   ```

5. **Submit a Pull Request**

   Go to the original repository on GitHub and open a pull request. Provide a detailed description of your changes and any additional context that may be helpful for the reviewers.

## Code Style

- Follow the coding style guidelines and best practices for Python, Go, and JavaScript.
- Use descriptive variable and function names.
- Write clear and concise comments where necessary.
- **Use Docstrings**:
  - For Python, follow the [Google Python Style Guide](https://google.github.io/styleguide/pyguide.html#38-comments-and-docstrings)
  - For Go, use [GoDoc](https://blog.golang.org/godoc-documenting-go-code) style for documentation.
  - For JavaScript, use [JSDoc](https://jsdoc.app/about-getting-started.html) style for documentation.
- **Update README Files**:
  - Ensure that you update the README files in relevant directories to reflect any changes you make to the code or documentation.
  - Provide clear instructions and examples where necessary to help others understand and use the code.

## Running Tests

Ensure that all tests pass before submitting your pull request. You can run the tests using the following commands:

- **Run tests for a specific project**:
  ```sh
  make check project=<project-name>
  ```

  If you are unsure about the project name, you can check it in the `project.json` file of the respective project.

- **Run all tests**:
  ```sh
  make check-all
  ```

## Reporting Issues

If you encounter any issues or have suggestions for improvements, please open an issue on GitHub. Provide as much detail as possible to help us understand and address the issue.

## License

By contributing to LabData TCC, you agree that your contributions will be licensed under the MIT License.

## Thank You!

Thank you for your contributions to labdata-tcc! Your efforts help make this project better for everyone. If you have any questions or need further assistance, feel free to contact us at [LinkedIn](https://www.linkedin.com/in/fabio-caffarello/) or open an issue on GitHub.



