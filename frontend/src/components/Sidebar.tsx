import React from 'react';
import { CheckCircle, Circle, Play } from 'lucide-react';
import type { Lesson, UserProgress } from '../types';

interface SidebarProps {
  lessons: Lesson[];
  currentLesson: Lesson | null;
  onSelectLesson: (lesson: Lesson) => void;
  userProgress: UserProgress[];
}

const Sidebar: React.FC<SidebarProps> = ({
  lessons,
  currentLesson,
  onSelectLesson,
  userProgress,
}) => {
  const getLessonStatus = (lessonId: number) => {
    const progress = userProgress?.find(p => p.lesson_id === lessonId);
    return progress?.completed ? 'completed' : 'locked';
  };

  const getDifficultyColor = (difficulty: string) => {
    switch (difficulty) {
      case 'beginner':
        return 'bg-green-100 text-green-800';
      case 'intermediate':
        return 'bg-yellow-100 text-yellow-800';
      case 'advanced':
        return 'bg-red-100 text-red-800';
      default:
        return 'bg-gray-100 text-gray-800';
    }
  };

  return (
    <div className="w-80 bg-white border-r border-gray-200 overflow-y-auto">
      <div className="p-6">
        <h2 className="text-lg font-semibold text-gray-900 mb-4">
          Learning Path
        </h2>
        
        <div className="space-y-2">
          {lessons.map((lesson) => {
            const status = getLessonStatus(lesson.id);
            const isCurrent = currentLesson?.id === lesson.id;
            
            return (
              <div
                key={lesson.id}
                className={`p-4 rounded-lg border cursor-pointer transition-all ${
                  isCurrent
                    ? 'border-go-blue bg-blue-50'
                    : status === 'completed'
                    ? 'border-green-200 bg-green-50 hover:bg-green-100'
                    : 'border-gray-200 bg-white hover:bg-gray-50'
                }`}
                onClick={() => onSelectLesson(lesson)}
              >
                <div className="flex items-start space-x-3">
                  <div className="flex-shrink-0 mt-1">
                    {status === 'completed' ? (
                      <CheckCircle className="w-5 h-5 text-green-600" />
                    ) : isCurrent ? (
                      <Play className="w-5 h-5 text-go-blue" />
                    ) : (
                      <Circle className="w-5 h-5 text-gray-400" />
                    )}
                  </div>
                  
                  <div className="flex-1 min-w-0">
                    <div className="flex items-center justify-between mb-1">
                      <h3 className={`text-sm font-medium ${
                        isCurrent ? 'text-go-blue' : 'text-gray-900'
                      }`}>
                        {lesson.title}
                      </h3>
                      <span className={`px-2 py-1 text-xs rounded-full ${getDifficultyColor(lesson.difficulty)}`}>
                        {lesson.difficulty}
                      </span>
                    </div>
                    
                    <p className="text-xs text-gray-600 line-clamp-2">
                      {lesson.description}
                    </p>
                    
                    <div className="flex items-center mt-2 text-xs text-gray-500">
                      <span>Lesson {lesson.order}</span>
                      {status === 'completed' && (
                        <span className="ml-2 text-green-600">✓ Completed</span>
                      )}
                    </div>
                  </div>
                </div>
              </div>
            );
          })}
        </div>
        
        {/* Progress Summary */}
        <div className="mt-6 p-4 bg-gray-50 rounded-lg">
          <h3 className="text-sm font-medium text-gray-900 mb-2">
            Your Progress
          </h3>
          <div className="space-y-2">
            <div className="flex justify-between text-sm">
              <span className="text-gray-600">Completed</span>
              <span className="font-medium">
                {userProgress?.filter(p => p.completed).length} / {lessons.length}
              </span>
            </div>
            <div className="w-full bg-gray-200 rounded-full h-2">
              <div
                className="bg-go-blue h-2 rounded-full transition-all duration-300"
                style={{
                  width: `${(userProgress?.filter(p => p.completed).length / lessons.length) * 100}%`,
                }}
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Sidebar;
