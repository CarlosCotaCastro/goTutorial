export interface Lesson {
  id: number;
  title: string;
  description: string;
  content: string;
  exercise: string;
  solution: string;
  difficulty: 'beginner' | 'intermediate' | 'advanced';
  order: number;
}

export interface CodeExecutionRequest {
  code: string;
}

export interface CodeExecutionResponse {
  output: string;
  error?: string;
}

export interface UserProgress {
  user_id: string;
  lesson_id: number;
  completed: boolean;
  completed_at?: string;
}

export interface ApiResponse<T> {
  data: T;
  error?: string;
}
