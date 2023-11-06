# CommuniTEA Application

Bringing your local community together over a cuppa.

## Table of Contents
- [Table of Contents](#table-of-contents)
- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Installation](#installation)

# Getting Started
To get started with this project, follow these steps:
1. clone the repository:
```sh
https://github.com/CommuniTEAM/CommuniTEA
```

2. Install the project dependencies:
```sh
npm install
```

3. Start the development server:
```sh
npm run dev
```

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Node.js: [Download](https://nodejs.org/)
- Docker: [Download](https://www.docker.com/get-started)

## Installation

1. Install Node.js packages:

   ```sh
   npm install
   ```

2. Build the Docker image:

   ```sh
   docker build -t communitea .
   ```


## Usage

1. Start the Docker container:

   ```sh
   docker run -p 3000:3000 communitea npm run dev
   ```

2. Access the application in your web browser at `http://localhost:3000`.



