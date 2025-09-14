# 🚀 The Ultimate Go Language Tutorial Platform

Welcome to the most comprehensive, interactive Go language learning platform! Learn Go programming through hands-on exercises with live code execution directly in your browser.

## ✨ Features

- **🎯 Interactive Code Editor** - Monaco editor with Go syntax highlighting
- **⚡ Live Code Execution** - Run Go code directly in the browser
- **📚 Progressive Tutorial System** - Step-by-step lessons with exercises
- **🔄 Real-time Feedback** - Immediate validation and hints
- **📊 Progress Tracking** - Visual learning progress and achievements
- **🎨 Modern UI/UX** - Beautiful, responsive design
- **🔒 Secure Execution** - Docker-based code execution sandbox

## 🛠 Tech Stack

### Backend
- **Go 1.21+** - Main server language
- **Gin** - Web framework
- **Docker** - Secure code execution sandbox
- **WebSocket** - Real-time communication

### Frontend
- **React 18** with TypeScript
- **Monaco Editor** - VS Code's editor with Go syntax highlighting
- **Tailwind CSS** - Modern styling
- **Vite** - Fast development and building
- **Axios** - HTTP client

## 🚀 Quick Start

### Prerequisites
- Go 1.21 or higher
- Node.js 18 or higher
- Docker (optional, for secure code execution)

### Option 1: Automated Startup (Recommended)
```bash
# Clone and navigate to the project
cd ~/goturorial

# Run the automated startup script
./start.sh
```

### Option 2: Manual Setup

#### Backend Setup
```bash
cd backend
go mod tidy
go run .
```

#### Frontend Setup
```bash
cd frontend
npm install
npm run dev
```

#### Docker Setup (Optional)
```bash
# Build the code execution container
cd docker
docker build -t go-executor:latest .
```

## 📚 Tutorial Structure

The tutorial is organized into progressive modules:

### 🎯 Learning Path
1. **Getting Started** - Hello World, Variables, Types
2. **Control Flow** - If/Else, Loops, Switch
3. **Functions** - Function basics, Multiple return values
4. **Data Structures** - Arrays, Slices, Maps, Structs
5. **Methods & Interfaces** - Object-oriented concepts
6. **Concurrency** - Goroutines, Channels (Coming Soon)
7. **Advanced Topics** - Error handling, Testing, Packages (Coming Soon)

### 📖 Lesson Format
Each lesson includes:
- **📝 Theory** - Clear explanations with examples
- **💻 Interactive Examples** - Run code snippets
- **🏋️ Exercises** - Hands-on practice
- **🎯 Challenges** - Advanced problems
- **📊 Progress Tracking** - Visual learning progress

## 🌐 Access Points

Once running, access the platform at:
- **Frontend**: http://localhost:5173
- **Backend API**: http://localhost:8080/api
- **Health Check**: http://localhost:8080/api/health

## 🎮 How to Use

1. **Select a Lesson** - Choose from the sidebar
2. **Read the Content** - Understand the concepts
3. **Write Code** - Use the Monaco editor with Go syntax highlighting
4. **Run Code** - Click "Run Code" or use Ctrl+Enter
5. **See Results** - View output in the right panel
6. **Check Solution** - Compare with the provided solution
7. **Track Progress** - Monitor your learning journey

## 🔧 Development

### Project Structure
```
goturorial/
├── backend/           # Go backend server
│   ├── main.go       # Main server file
│   ├── lessons.go    # Tutorial content
│   ├── executor.go   # Code execution service
│   └── go.mod        # Go dependencies
├── frontend/         # React frontend
│   ├── src/
│   │   ├── components/  # React components
│   │   ├── services/     # API services
│   │   ├── types/       # TypeScript types
│   │   └── App.tsx      # Main app component
│   └── package.json     # Node dependencies
├── docker/           # Docker configuration
│   ├── Dockerfile    # Code execution container
│   └── execute.go    # Execution helper
└── start.sh         # Automated startup script
```

### API Endpoints
- `GET /api/health` - Health check
- `GET /api/lessons` - Get all lessons
- `GET /api/lessons/:id` - Get specific lesson
- `POST /api/execute` - Execute Go code
- `GET /api/progress/:user_id` - Get user progress
- `POST /api/progress` - Update progress
- `GET /api/ws` - WebSocket connection

## 🎯 Current Lessons

1. **Hello, Go!** - Write your first Go program
2. **Variables and Types** - Understanding Go's type system
3. **Functions** - Creating and using functions
4. **Control Flow - If/Else** - Making decisions
5. **Loops** - Repeating code execution
6. **Arrays and Slices** - Working with collections
7. **Maps** - Storing key-value pairs
8. **Structs** - Creating custom data types
9. **Methods** - Adding behavior to structs
10. **Interfaces** - Defining behavior contracts

## 🔒 Security Features

- **Docker Sandbox** - Code execution in isolated containers
- **Timeout Protection** - Prevents infinite loops
- **Resource Limits** - Memory and CPU constraints
- **Input Validation** - Sanitized code execution

## 🚀 Future Enhancements

- [ ] User authentication and accounts
- [ ] More advanced Go topics (goroutines, channels)
- [ ] Code challenges and competitions
- [ ] Social features (leaderboards, sharing)
- [ ] Mobile app version
- [ ] Offline mode support

## 🤝 Contributing

This is a learning project! Contributions are welcome:
- Add new lessons and exercises
- Improve the UI/UX
- Enhance code execution security
- Add new features

## 📄 License

MIT License - Feel free to use and modify for your learning journey!

---

**Happy Learning! 🎉** Start your Go programming journey today with the most interactive tutorial platform available!