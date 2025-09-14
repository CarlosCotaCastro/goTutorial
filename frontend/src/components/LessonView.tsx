import React, { useState } from 'react';
import { BookOpen, Lightbulb, RotateCcw, Eye, Code, ChevronDown, ChevronRight } from 'lucide-react';
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
  const [showExplanation, setShowExplanation] = useState(false);
  const [showVariants, setShowVariants] = useState(false);
  const [selectedVariant, setSelectedVariant] = useState(0);

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

        {/* Detailed Explanation */}
        {lesson.explanation && (
          <div className="mb-6">
            <button
              onClick={() => setShowExplanation(!showExplanation)}
              className="flex items-center space-x-2 w-full p-4 bg-blue-50 border border-blue-200 rounded-lg hover:bg-blue-100 transition-colors"
            >
              <BookOpen className="w-5 h-5 text-blue-600" />
              <span className="text-blue-900 font-semibold">Detailed Explanation</span>
              {showExplanation ? (
                <ChevronDown className="w-4 h-4 text-blue-600 ml-auto" />
              ) : (
                <ChevronRight className="w-4 h-4 text-blue-600 ml-auto" />
              )}
            </button>
            
            {showExplanation && (
              <div className="mt-4 p-4 bg-blue-50 border border-blue-200 rounded-lg">
                <div className="prose prose-sm max-w-none text-blue-900 whitespace-pre-wrap">
                  {lesson.explanation}
                </div>
              </div>
            )}
          </div>
        )}

        {/* Code Variants */}
        {lesson.variants && lesson.variants.length > 0 && (
          <div className="mb-6">
            <button
              onClick={() => setShowVariants(!showVariants)}
              className="flex items-center space-x-2 w-full p-4 bg-green-50 border border-green-200 rounded-lg hover:bg-green-100 transition-colors"
            >
              <Code className="w-5 h-5 text-green-600" />
              <span className="text-green-900 font-semibold">Code Variants & Examples</span>
              {showVariants ? (
                <ChevronDown className="w-4 h-4 text-green-600 ml-auto" />
              ) : (
                <ChevronRight className="w-4 h-4 text-green-600 ml-auto" />
              )}
            </button>
            
            {showVariants && (
              <div className="mt-4">
                {/* Variant Selector */}
                <div className="mb-4">
                  <div className="flex space-x-2">
                    {lesson.variants.map((_, index) => (
                      <button
                        key={index}
                        onClick={() => setSelectedVariant(index)}
                        className={`px-3 py-1 text-sm rounded ${
                          selectedVariant === index
                            ? 'bg-green-600 text-white'
                            : 'bg-green-100 text-green-800 hover:bg-green-200'
                        }`}
                      >
                        Variant {index + 1}
                      </button>
                    ))}
                  </div>
                </div>
                
                {/* Selected Variant Code */}
                <div className="bg-gray-900 text-gray-100 p-4 rounded-lg overflow-x-auto">
                  <pre className="text-sm">
                    <code>{lesson.variants[selectedVariant]}</code>
                  </pre>
                </div>
              </div>
            )}
          </div>
        )}

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
