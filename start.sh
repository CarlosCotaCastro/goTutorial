#!/bin/bash

# Go Tutorial Platform Startup Script
echo "🚀 Starting Go Tutorial Platform..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if required tools are installed
check_dependencies() {
    print_status "Checking dependencies..."
    
    if ! command -v go &> /dev/null; then
        print_error "Go is not installed. Please install Go 1.21+ from https://golang.org/dl/"
        exit 1
    fi
    
    if ! command -v node &> /dev/null; then
        print_error "Node.js is not installed. Please install Node.js 18+ from https://nodejs.org/"
        exit 1
    fi
    
    if ! command -v docker &> /dev/null; then
        print_warning "Docker is not installed. Code execution will use fallback method."
    fi
    
    print_success "Dependencies check completed"
}

# Build Docker image for code execution
build_docker_image() {
    if command -v docker &> /dev/null; then
        print_status "Building Docker image for code execution..."
        cd docker
        docker build -t go-executor:latest .
        cd ..
        print_success "Docker image built successfully"
    else
        print_warning "Skipping Docker build (Docker not available)"
    fi
}

# Start the backend server
start_backend() {
    print_status "Starting Go backend server..."
    cd backend
    
    # Install dependencies
    go mod tidy
    
    # Start the server in background
    go run . &
    BACKEND_PID=$!
    
    cd ..
    print_success "Backend server started (PID: $BACKEND_PID)"
}

# Start the frontend development server
start_frontend() {
    print_status "Starting React frontend..."
    cd frontend
    
    # Install dependencies if needed
    if [ ! -d "node_modules" ]; then
        print_status "Installing frontend dependencies..."
        npm install
    fi
    
    # Start the development server
    npm run dev &
    FRONTEND_PID=$!
    
    cd ..
    print_success "Frontend server started (PID: $FRONTEND_PID)"
}

# Wait for servers to be ready
wait_for_servers() {
    print_status "Waiting for servers to be ready..."
    
    # Wait for backend
    for i in {1..30}; do
        if curl -s http://localhost:8080/api/health > /dev/null; then
            print_success "Backend server is ready!"
            break
        fi
        sleep 1
    done
    
    # Wait for frontend
    for i in {1..30}; do
        if curl -s http://localhost:5173 > /dev/null; then
            print_success "Frontend server is ready!"
            break
        fi
        sleep 1
    done
}

# Main execution
main() {
    echo "=========================================="
    echo "🎯 Go Tutorial Platform Startup"
    echo "=========================================="
    
    check_dependencies
    build_docker_image
    start_backend
    start_frontend
    wait_for_servers
    
    echo ""
    echo "=========================================="
    print_success "🚀 Go Tutorial Platform is running!"
    echo "=========================================="
    echo ""
    echo "📚 Frontend: http://localhost:5173"
    echo "🔧 Backend API: http://localhost:8080/api"
    echo "📖 Health Check: http://localhost:8080/api/health"
    echo ""
    echo "Press Ctrl+C to stop all services"
    echo ""
    
    # Keep script running and handle cleanup
    trap 'echo ""; print_status "Shutting down services..."; kill $BACKEND_PID $FRONTEND_PID 2>/dev/null; exit 0' INT
    
    # Wait for user interrupt
    wait
}

# Run main function
main
