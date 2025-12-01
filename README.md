# Go Calculator + Multi-Stage Docker Project

## This project demonstrates a multi-container Docker setup with a Go backend and a static frontend, allowing users to:

1. Perform simple calculations (add two numbers)

2. Display all results dynamically on a browser

3. The project is fully containerized with separate backend and frontend services.

## 📂 Project Structure

``` bash
project/
│
├── BE/                     # Backend service
│   ├── Dockerfile          # Backend Dockerfile
│   └── main.go             # Go backend code
│
├── UI/                     # Frontend service
│   ├── Dockerfile          # Frontend Dockerfile
│   ├── index.html          # Main HTML page
│   └── static/
│       └── styles.css      # Static files
├── docker-compose.yaml     # Docker Compose configuration
└── README.md
```

## Features:

### Frontend

* Accessible in the browser (port 9000)

* Allows the user to:

  - Enter two numbers and click Calculate or press Enter

  - Display calculation results dynamically in a table

  - Press Restart to refresh the page

* Styled with soft grey and orange tones and white backgrounds for inputs and table

### Backend

* Written in Go

* Exposes APIs:

  - /add → store a calculation dinamically

* Multi-stage Docker build keeps image small and efficient

---

## Docker Architecture:

* Backend container

  - Runs Go backend

  - Listens on port 8000

* Frontend container

  - Serves static HTML/CSS/JS

  - Listens on port 9000

  - Communicates with backend via API

---

## Running the Project:

1. Clone the repository

``` bash
git clone <your-repository-url>
cd project
```

2. Build and run containers

``` bash
docker compose up --build
```

3. Using the App

- Open browser → ```http://localhost:9000```

- Enter two numbers and click Calculate or press Enter

- View the result in the table

- Press Restart to refresh the page

---

## Summary

This project uses Docker Compose to run two separate containers:

* Backend container – runs the Go server handling calculation requests.

* Frontend container – serves the static HTML, CSS, and JS interface in the browser.

#### The containers are isolated, yet communicate via exposed ports. The backend listens on port 8000, and the frontend serves the app on port 9000. Using Docker Compose makes it easy to build, run, and manage the application consistently across environments.


