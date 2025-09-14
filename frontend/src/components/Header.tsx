import React from 'react';
import { Play, BookOpen, Code, Trophy } from 'lucide-react';

const Header: React.FC = () => {
  return (
    <header className="fixed top-0 left-0 right-0 bg-white shadow-sm border-b border-gray-200 z-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          {/* Logo and Title */}
          <div className="flex items-center space-x-3">
            <div className="flex items-center space-x-2">
              <div className="w-8 h-8 bg-go-blue rounded-lg flex items-center justify-center">
                <Code className="w-5 h-5 text-white" />
              </div>
              <h1 className="text-xl font-bold text-gray-900">
                Go Tutorial
              </h1>
            </div>
            <div className="hidden md:flex items-center space-x-1 text-sm text-gray-500">
              <BookOpen className="w-4 h-4" />
              <span>Interactive Learning Platform</span>
            </div>
          </div>

          {/* Navigation */}
          <nav className="hidden md:flex items-center space-x-8">
            <a href="#lessons" className="text-gray-700 hover:text-go-blue transition-colors">
              Lessons
            </a>
            <a href="#exercises" className="text-gray-700 hover:text-go-blue transition-colors">
              Exercises
            </a>
            <a href="#progress" className="text-gray-700 hover:text-go-blue transition-colors">
              Progress
            </a>
          </nav>

          {/* User Actions */}
          <div className="flex items-center space-x-4">
            <div className="flex items-center space-x-2 text-sm text-gray-600">
              <Trophy className="w-4 h-4" />
              <span>Level: Beginner</span>
            </div>
            
            <button className="bg-go-blue text-white px-4 py-2 rounded-lg hover:bg-blue-600 transition-colors flex items-center space-x-2">
              <Play className="w-4 h-4" />
              <span>Run Code</span>
            </button>
          </div>
        </div>
      </div>
    </header>
  );
};

export default Header;
