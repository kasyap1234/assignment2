Sure! Here are the instructions to start your Task Management Application with the repository URL.

---

# Task Management Application

## Getting Started
[Screen Recorder.webm](https://github.com/kasyap1234/assignment2/assets/76957823/05555ec5-1bc4-47aa-aede-0662c232cfae)



### Prerequisites

- [Node.js](https://nodejs.org/) (for frontend)
- [Golang](https://golang.org/) (for backend)

### Frontend Setup

1. Clone the repository:

```bash
git clone https://github.com/kasyap1234/assignment2.git
cd tasmanager-frontend
```

2. Install dependencies:

```bash
npm install
```

3. Start the development server:

```bash
npm run dev
```

The frontend will be running at `http://localhost:3000`.

### Backend Setup

1. Navigate to the backend directory:

```bash
cd tasmanager-backend
```

2. Install dependencies and build the server:

```bash
go mod tidy
go build
```

3. Start the backend server:

```bash
go run main.go
```

The backend will be running at `http://localhost:8000`.

## Running the Application

1. Ensure your backend server is running on `http://localhost:8000`.
2. Ensure your frontend server is running on `http://localhost:5173`.

With both servers running, you can access the Task Management Application in your browser by navigating to `http://localhost:5173`.

---

Feel free to adjust any specifics if necessary.
