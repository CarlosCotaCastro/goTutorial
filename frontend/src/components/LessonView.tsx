import React from 'react';
import { BookOpen, Lightbulb, RotateCcw, Eye } from 'lucide-react';
import type { Lesson } from '../types';

interface LessonViewProps {
  lesson: Lesson;
  onCheckSolution: () => void;
  onResetExercise: () => void;
}

const LessonView: React.FC<LessonViewProps> = ({
  lesson,
  onCheckSolution,
  onResetExercise,
}) => {
  return (
    <div className="bg-white border-b border-gray-200 p-6">
      <div className="max-w-4xl mx-auto">
        {/* Lesson Header */}
        <div className="mb-6">
          <div className="flex items-center space-x-3 mb-2">
            <BookOpen className="w-6 h-6 text-go-blue" />
            <h1 className="text-2xl font-bold text-gray-900">
              {lesson.title}
            </h1>
            <span className={`px-3 py-1 text-sm rounded-full ${
              lesson.difficulty === 'beginner'
                ? 'bg-green-100 text-green-800'
                : lesson.difficulty === 'intermediate'
                ? 'bg-yellow-100 text-yellow-800'
                : 'bg-red-100 text-red-800'
            }`}>
              {lesson.difficulty}
            </span>
          </div>
          <p className="text-gray-600 text-lg">
            {lesson.description}
          </p>
        </div>

        {/* Lesson Content */}
        <div className="prose prose-lg max-w-none mb-6">
          <div className="bg-gray-50 p-4 rounded-lg">
            <h3 className="text-lg font-semibold text-gray-900 mb-3 flex items-center">
              <Lightbulb className="w-5 h-5 text-yellow-500 mr-2" />
              What You'll Learn
            </h3>
            <div className="text-gray-700 whitespace-pre-wrap">
              {lesson.content}
            </div>
          </div>
        </div>

        {/* Exercise Section */}
        <div className="bg-blue-50 border border-blue-200 rounded-lg p-4 mb-6">
          <h3 className="text-lg font-semibold text-blue-900 mb-3">
            Exercise
          </h3>
          <div className="text-blue-800 whitespace-pre-wrap">
            {lesson.exercise}
          </div>
        </div>

        {/* Action Buttons */}
        <div className="flex items-center space-x-4">
          <button
            onClick={onResetExercise}
            className="flex items-center space-x-2 px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors"
          >
            <RotateCcw className="w-4 h-4" />
            <span>Reset Exercise</span>
          </button>
          
          <button
            onClick={onCheckSolution}
            className="flex items-center space-x-2 px-4 py-2 bg-go-blue text-white rounded-lg hover:bg-blue-600 transition-colors"
          >
            <Eye className="w-4 h-4" />
            <span>Show Solution</span>
          </button>
        </div>
      </div>
    </div>
  );
};

export default LessonView;
