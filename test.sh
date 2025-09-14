#!/bin/bash

# Go Tutorial Platform Test Script
echo "🧪 Testing Go Tutorial Platform..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[TEST]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[PASS]${NC} $1"
}

print_error() {
    echo -e "${RED}[FAIL]${NC} $1"
}

# Test backend health endpoint
test_backend_health() {
    print_status "Testing backend health endpoint..."
    
    response=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/api/health)
    
    if [ "$response" = "200" ]; then
        print_success "Backend health check passed"
        return 0
    else
        print_error "Backend health check failed (HTTP $response)"
        return 1
    fi
}

# Test lessons endpoint
test_lessons_endpoint() {
    print_status "Testing lessons endpoint..."
    
    response=$(curl -s http://localhost:8080/api/lessons)
    
    if echo "$response" | grep -q "Hello, Go!"; then
        print_success "Lessons endpoint working"
        return 0
    else
        print_error "Lessons endpoint failed"
        return 1
    fi
}

# Test code execution
test_code_execution() {
    print_status "Testing code execution..."
    
    code='{"code":"package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, Test!\")\n}"}'
    
    response=$(curl -s -X POST -H "Content-Type: application/json" -d "$code" http://localhost:8080/api/execute)
    
    if echo "$response" | grep -q "Hello, Test!"; then
        print_success "Code execution working"
        return 0
    else
        print_error "Code execution failed"
        return 1
    fi
}

# Test frontend availability
test_frontend() {
    print_status "Testing frontend availability..."
    
    response=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:5173)
    
    if [ "$response" = "200" ]; then
        print_success "Frontend is accessible"
        return 0
    else
        print_error "Frontend not accessible (HTTP $response)"
        return 1
    fi
}

# Main test function
main() {
    echo "=========================================="
    echo "🧪 Go Tutorial Platform Test Suite"
    echo "=========================================="
    
    local tests_passed=0
    local total_tests=4
    
    test_backend_health && ((tests_passed++))
    test_lessons_endpoint && ((tests_passed++))
    test_code_execution && ((tests_passed++))
    test_frontend && ((tests_passed++))
    
    echo ""
    echo "=========================================="
    echo "📊 Test Results: $tests_passed/$total_tests passed"
    echo "=========================================="
    
    if [ $tests_passed -eq $total_tests ]; then
        print_success "🎉 All tests passed! Platform is working correctly."
        echo ""
        echo "🌐 Access your Go Tutorial Platform:"
        echo "   Frontend: http://localhost:5173"
        echo "   Backend:  http://localhost:8080/api"
        echo ""
    else
        print_error "❌ Some tests failed. Check the output above."
        exit 1
    fi
}

# Run tests
main
