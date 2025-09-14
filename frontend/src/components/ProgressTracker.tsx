import React, { useState, useEffect } from 'react';
import { Trophy, Target, Clock, Award } from 'lucide-react';
import type { UserProgress } from '../types';

interface ProgressTrackerProps {
  userProgress: UserProgress[];
  totalLessons: number;
}

const ProgressTracker: React.FC<ProgressTrackerProps> = ({
  userProgress,
  totalLessons,
}) => {
  const [sessionTime, setSessionTime] = useState<number>(0);
  const [totalTimeSpent, setTotalTimeSpent] = useState<number>(0);

  const completedLessons = userProgress.filter(p => p.completed).length;
  const completionPercentage = totalLessons > 0 ? (completedLessons / totalLessons) * 100 : 0;

  // Track session time
  useEffect(() => {
    const startTime = Date.now();
    const interval = setInterval(() => {
      const elapsed = Math.floor((Date.now() - startTime) / 1000 / 60); // minutes
      setSessionTime(elapsed);
    }, 60000); // Update every minute

    // Save time when component unmounts
    return () => {
      clearInterval(interval);
      const finalElapsed = Math.floor((Date.now() - startTime) / 1000 / 60);
      if (finalElapsed > 0) {
        const currentTotal = parseInt(localStorage.getItem('totalTimeSpent') || '0', 10);
        localStorage.setItem('totalTimeSpent', (currentTotal + finalElapsed).toString());
      }
    };
  }, []);

  // Load total time spent from localStorage
  useEffect(() => {
    const savedTime = localStorage.getItem('totalTimeSpent');
    if (savedTime) {
      setTotalTimeSpent(parseInt(savedTime, 10));
    }
  }, []);

  // Save total time spent to localStorage
  useEffect(() => {
    if (totalTimeSpent > 0) {
      localStorage.setItem('totalTimeSpent', totalTimeSpent.toString());
    }
  }, [totalTimeSpent]);

  const getAchievementLevel = () => {
    if (completionPercentage >= 100) return { level: 'Master', color: 'text-purple-600', icon: Award };
    if (completionPercentage >= 75) return { level: 'Expert', color: 'text-blue-600', icon: Trophy };
    if (completionPercentage >= 50) return { level: 'Intermediate', color: 'text-green-600', icon: Target };
    if (completionPercentage >= 25) return { level: 'Beginner', color: 'text-yellow-600', icon: Clock };
    return { level: 'Getting Started', color: 'text-gray-600', icon: Clock };
  };

  const achievement = getAchievementLevel();
  const AchievementIcon = achievement.icon;

  const getStreakDays = () => {
    if (!userProgress || userProgress.length === 0) return 0;
    
    const completedProgress = userProgress
      .filter(p => p.completed && p.completed_at)
      .map(p => new Date(p.completed_at!).toDateString())
      .sort((a, b) => new Date(b).getTime() - new Date(a).getTime());

    if (completedProgress.length === 0) return 0;

    let streak = 0;
    const today = new Date().toDateString();
    const yesterday = new Date(Date.now() - 24 * 60 * 60 * 1000).toDateString();
    
    // Check if user completed something today or yesterday
    if (completedProgress.includes(today) || completedProgress.includes(yesterday)) {
      streak = 1;
      
      // Count consecutive days
      for (let i = 1; i < completedProgress.length; i++) {
        const currentDate = new Date(completedProgress[i]);
        const previousDate = new Date(completedProgress[i - 1]);
        const diffTime = previousDate.getTime() - currentDate.getTime();
        const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));
        
        if (diffDays === 1) {
          streak++;
        } else {
          break;
        }
      }
    }
    
    return streak;
  };

  const getTotalTimeSpent = () => {
    return totalTimeSpent + sessionTime;
  };

  return (
    <div className="w-80 bg-white border-l border-gray-200 overflow-y-auto">
      <div className="p-6">
        <h2 className="text-lg font-semibold text-gray-900 mb-6">
          Your Progress
        </h2>

        {/* Achievement Level */}
        <div className="bg-gradient-to-r from-blue-50 to-purple-50 rounded-lg p-4 mb-6">
          <div className="flex items-center space-x-3 mb-3">
            <div className="p-2 bg-white rounded-lg shadow-sm">
              <AchievementIcon className={`w-6 h-6 ${achievement.color}`} />
            </div>
            <div>
              <h3 className="font-semibold text-gray-900">Achievement Level</h3>
              <p className={`text-sm font-medium ${achievement.color}`}>
                {achievement.level}
              </p>
            </div>
          </div>
          
          <div className="space-y-2">
            <div className="flex justify-between text-sm">
              <span className="text-gray-600">Progress</span>
              <span className="font-medium">
                {completedLessons} / {totalLessons} lessons
              </span>
            </div>
            <div className="w-full bg-gray-200 rounded-full h-3">
              <div
                className="bg-gradient-to-r from-blue-500 to-purple-500 h-3 rounded-full transition-all duration-500"
                style={{ width: `${completionPercentage}%` }}
              />
            </div>
            <p className="text-xs text-gray-500 text-center">
              {completionPercentage.toFixed(1)}% Complete
            </p>
          </div>
        </div>

        {/* Stats */}
        <div className="space-y-4 mb-6">
          <div className="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div className="flex items-center space-x-3">
              <Trophy className="w-5 h-5 text-yellow-500" />
              <span className="text-sm font-medium text-gray-700">Lessons Completed</span>
            </div>
            <span className="text-lg font-bold text-gray-900">{completedLessons}</span>
          </div>

          <div className="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div className="flex items-center space-x-3">
              <Clock className="w-5 h-5 text-blue-500" />
              <span className="text-sm font-medium text-gray-700">Time Spent</span>
            </div>
            <span className="text-lg font-bold text-gray-900">{getTotalTimeSpent()}m</span>
          </div>

          <div className="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div className="flex items-center space-x-3">
              <Target className="w-5 h-5 text-green-500" />
              <span className="text-sm font-medium text-gray-700">Current Streak</span>
            </div>
            <span className="text-lg font-bold text-gray-900">{getStreakDays()} days</span>
          </div>
        </div>

        {/* Recent Activity */}
        <div className="mb-6">
          <h3 className="text-sm font-semibold text-gray-900 mb-3">
            Recent Activity
          </h3>
          <div className="space-y-2">
            {userProgress
              .filter(p => p.completed)
              .slice(-3)
              .map((progress, index) => (
                <div key={index} className="flex items-center space-x-3 p-2 bg-green-50 rounded-lg">
                  <div className="w-2 h-2 bg-green-500 rounded-full" />
                  <span className="text-sm text-gray-700">
                    Completed Lesson {progress.lesson_id}
                  </span>
                </div>
              ))}
            {userProgress.filter(p => p.completed).length === 0 && (
              <p className="text-sm text-gray-500 text-center py-4">
                No completed lessons yet
              </p>
            )}
          </div>
        </div>

        {/* Next Goals */}
        <div className="bg-yellow-50 border border-yellow-200 rounded-lg p-4">
          <h3 className="text-sm font-semibold text-yellow-800 mb-2">
            Next Goals
          </h3>
          <div className="space-y-2 text-sm text-yellow-700">
            <p>• Complete 5 more lessons to reach Intermediate level</p>
            <p>• Practice coding for 30 minutes daily</p>
            <p>• Master Go functions and methods</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ProgressTracker;
