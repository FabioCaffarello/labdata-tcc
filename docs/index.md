# Getting Started

## Introduction

Welcome to the LabData TCC project! This project is part of my final work for the master's program in Data Engineering at FIA Business School. The aim of this project is to ... <!-- TODO: briefly describe the main goal of the project -->. This repository is a monorepo, containing multiple packages and services that work together to achieve the project's goals.

## Prerequisites
Depending on how you choose to run the project, the requirements may vary. Please ensure you have met the following requirements based on your preferred method:

1. **Running Locally**

   - You have installed:
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
     - [docker 26.1.1](https://www.docker.com/)

2. **Running in a Devcontainer Locally**

   - You have installed:
     - Docker
     - Visual Studio Code with the Remote - Containers extension

3. **Running in GitHub Codespaces**

   - You have a GitHub account with access to GitHub Codespaces.

## Installation

1. **Clone the Repository**

   First, clone the repository to your local machine using git:

   ```sh
   git clone https://github.com/FabioCaffarello/labdata-tcc.git
   cd labdata-tcc
   ```

## Running the Project

1. **Running Locally**

   - **Set Up the Environment**

     Ensure you have all the required dependencies:

     - NodeJS:
        ```sh
        npm install
        ```

     - Python:
        ```sh
        poetry install
        ```

2. **Running in a Devcontainer Locally**

   Ensure you have Docker and Visual Studio Code installed with the Remote - Containers extension. Open the repository in Visual Studio Code and click on "Reopen in Container" when prompted. The environment will be set up automatically.

3. **Running in GitHub Codespaces**

   You can also run the project in GitHub Codespaces directly. Navigate to the repository on GitHub, click the "Code" button, and select "Open with Codespaces" to create a new codespace. The environment will be set up automatically.

<!-- ## Usage

To start using the project, run the following command:

```sh
python main.py
```

Or if you are using Docker:

```sh
docker-compose up
```

You can access the application at `http://localhost:8000`.

#### Features

- **Feature 1**: Brief description of feature 1.
- **Feature 2**: Brief description of feature 2.
- **Feature 3**: Brief description of feature 3. -->

## Contributing
Contributions are welcome! Please read the [CONTRIBUTING.md](https://github.com/FabioCaffarello/labdata-tcc/blob/main/docs/CONTRIBUTING.md) file for details on the process for submitting pull requests.

#### License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/FabioCaffarello/labdata-tcc/blob/main/LICENSE) file for details.

#### Contact

If you have any questions or need further assistance, feel free to contact me at [linkedin](https://www.linkedin.com/in/fabio-caffarello/) or open an issue on GitHub.
