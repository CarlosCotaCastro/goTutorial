import axios from 'axios';
import type { Lesson, CodeExecutionRequest, CodeExecutionResponse, UserProgress } from '../types';

import { config } from '../config';

const API_BASE_URL = config.API_URL;

const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const apiService = {
  // Health check
  async healthCheck(): Promise<boolean> {
    try {
      const response = await api.get('/health');
      return response.status === 200;
    } catch (error) {
      console.error('Health check failed:', error);
      return false;
    }
  },

  // Get all lessons
  async getLessons(): Promise<Lesson[]> {
    try {
      const response = await api.get('/lessons');
      return response.data;
    } catch (error) {
      console.error('Failed to fetch lessons:', error);
      throw error;
    }
  },

  // Get a specific lesson
  async getLesson(id: number): Promise<Lesson> {
    try {
      const response = await api.get(`/lessons/${id}`);
      return response.data;
    } catch (error) {
      console.error(`Failed to fetch lesson ${id}:`, error);
      throw error;
    }
  },

  // Execute Go code
  async executeCode(code: string): Promise<CodeExecutionResponse> {
    try {
      const request: CodeExecutionRequest = { code };
      const response = await api.post('/execute', request);
      return response.data;
    } catch (error) {
      console.error('Code execution failed:', error);
      throw error;
    }
  },

  // Get user progress
  async getUserProgress(userId: string): Promise<UserProgress[]> {
    try {
      const response = await api.get(`/progress/${userId}`);
      return response.data;
    } catch (error) {
      console.error('Failed to fetch user progress:', error);
      throw error;
    }
  },

  // Update user progress
  async updateProgress(progress: UserProgress): Promise<void> {
    try {
      await api.post('/progress', progress);
    } catch (error) {
      console.error('Failed to update progress:', error);
      throw error;
    }
  },

  // WebSocket connection for real-time features
  createWebSocketConnection(): WebSocket {
    return new WebSocket(config.WS_URL);
  },
};
