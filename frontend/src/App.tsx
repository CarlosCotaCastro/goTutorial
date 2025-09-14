import { useState, useEffect } from 'react';
import { BrowserRouter as Router } from 'react-router-dom';
import Header from './components/Header';
import Sidebar from './components/Sidebar';
import LessonView from './components/LessonView';
import CodeEditor from './components/CodeEditor';
import OutputPanel from './components/OutputPanel';
import ProgressTracker from './components/ProgressTracker';
import type { Lesson, CodeExecutionResponse, UserProgress } from './types';
import { apiService } from './services/api';

function App() {
  const [lessons, setLessons] = useState<Lesson[]>([]);
  const [currentLesson, setCurrentLesson] = useState<Lesson | null>(null);
  const [code, setCode] = useState<string>('');
  const [output, setOutput] = useState<string>('');
  const [isExecuting, setIsExecuting] = useState<boolean>(false);
  const [userProgress, setUserProgress] = useState<UserProgress[]>([]);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    loadLessons();
    loadUserProgress();
  }, []);

  const loadLessons = async () => {
    try {
      const data = await apiService.getLessons();
      setLessons(data);
      if (data.length > 0) {
        setCurrentLesson(data[0]);
        setCode(data[0].exercise);
      }
    } catch (error) {
      console.error('Failed to load lessons:', error);
      // Set some default lessons if API fails
      const defaultLessons: Lesson[] = [
        {
          id: 1,
          title: "Hello, Go!",
          description: "Write your first Go program",
          content: "Welcome to Go! Learn the basics of Go programming.",
          exercise: "Write a program that prints 'Hello, World!'",
          solution: "package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, World!\")\n}",
          difficulty: "beginner",
          order: 1,
        }
      ];
      setLessons(defaultLessons);
      setCurrentLesson(defaultLessons[0]);
      setCode(defaultLessons[0].exercise);
    } finally {
      setLoading(false);
    }
  };

  const loadUserProgress = async () => {
    try {
      const userId = 'demo-user';
      const progress = await apiService.getUserProgress(userId);
      setUserProgress(progress);
    } catch (error) {
      console.error('Failed to load user progress:', error);
      // Try to load from localStorage as fallback
      const savedProgress = localStorage.getItem('userProgress');
      if (savedProgress) {
        try {
          const parsedProgress = JSON.parse(savedProgress);
          setUserProgress(parsedProgress);
        } catch (parseError) {
          console.error('Failed to parse saved progress:', parseError);
          setUserProgress([]);
        }
      } else {
        setUserProgress([]);
      }
    }
  };

  const executeCode = async () => {
    if (!code.trim()) return;

    setIsExecuting(true);
    setOutput('');

    try {
      const response: CodeExecutionResponse = await apiService.executeCode(code);
      setOutput(response.output);
      
      if (response.error) {
        setOutput(`Error: ${response.error}`);
      } else {
        // Check if the code execution was successful and mark lesson as completed
        checkAndMarkLessonCompleted();
      }
    } catch (error) {
      setOutput(`Error: ${error}`);
    } finally {
      setIsExecuting(false);
    }
  };

  const checkAndMarkLessonCompleted = async () => {
    if (!currentLesson) return;

    // Check if lesson is already completed
    const existingProgress = userProgress.find(p => p.lesson_id === currentLesson.id);
    if (existingProgress?.completed) return;

    // Enhanced completion check
    const isCompleted = checkLessonCompletion(currentLesson, code, output);

    if (isCompleted) {
      const newProgress: UserProgress = {
        user_id: 'demo-user',
        lesson_id: currentLesson.id,
        completed: true,
        completed_at: new Date().toISOString(),
      };

      try {
        await apiService.updateProgress(newProgress);
        
        // Update local state
        setUserProgress(prev => {
          const filtered = prev.filter(p => p.lesson_id !== currentLesson.id);
          return [...filtered, newProgress];
        });

        // Save to localStorage for persistence
        const savedProgress = JSON.parse(localStorage.getItem('userProgress') || '[]');
        const updatedProgress = savedProgress.filter((p: UserProgress) => p.lesson_id !== currentLesson.id);
        updatedProgress.push(newProgress);
        localStorage.setItem('userProgress', JSON.stringify(updatedProgress));

        console.log(`🎉 Lesson "${currentLesson.title}" completed!`);
      } catch (error) {
        console.error('Failed to update progress:', error);
      }
    }
  };

  const checkLessonCompletion = (lesson: Lesson, code: string, output: string): boolean => {
    // Basic checks
    if (!code.trim() || output.includes('Error:')) return false;

    // Lesson-specific completion criteria
    switch (lesson.id) {
      case 1: // Hello, Go!
        return code.includes('fmt.Println') && output.includes('Hello');
      case 2: // Variables
        return code.includes('var ') || code.includes(':=');
      case 3: // Functions
        return code.includes('func ');
      case 4: // Loops
        return code.includes('for ');
      case 5: // Conditionals
        return code.includes('if ');
      default:
        // Generic completion - code runs without errors and has meaningful content
        return code.trim().length > 20 && !output.includes('Error:');
    }
  };

  const selectLesson = (lesson: Lesson) => {
    setCurrentLesson(lesson);
    setCode(lesson.exercise);
    setOutput('');
  };

  const checkSolution = () => {
    if (currentLesson) {
      setCode(currentLesson.solution);
    }
  };

  const resetExercise = () => {
    if (currentLesson) {
      setCode(currentLesson.exercise);
      setOutput('');
    }
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600 mx-auto mb-4"></div>
          <p className="text-gray-600">Loading Go Tutorial...</p>
        </div>
      </div>
    );
  }

  return (
    <Router>
      <div className="min-h-screen bg-gray-50">
        <Header />
        
        <div className="flex h-screen pt-16">
          <Sidebar 
            lessons={lessons}
            currentLesson={currentLesson}
            onSelectLesson={selectLesson}
            userProgress={userProgress}
          />
          
          <div className="flex-1 flex flex-col">
            <div className="flex-1 flex">
              {/* Main Content Area */}
              <div className="flex-1 flex flex-col">
                {currentLesson && (
                  <LessonView 
                    lesson={currentLesson}
                    onCheckSolution={checkSolution}
                    onResetExercise={resetExercise}
                  />
                )}
                
                <div className="flex-1 flex">
                  <CodeEditor 
                    code={code}
                    onChange={setCode}
                    onExecute={executeCode}
                    isExecuting={isExecuting}
                  />
                  
                  <OutputPanel 
                    output={output}
                    isExecuting={isExecuting}
                  />
                </div>
              </div>
              
              {/* Progress Panel */}
              <ProgressTracker 
                userProgress={userProgress}
                totalLessons={lessons.length}
              />
            </div>
          </div>
        </div>
      </div>
    </Router>
  );
}

export default App;