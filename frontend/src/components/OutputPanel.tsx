import React from 'react';
import { Terminal, CheckCircle, XCircle, Loader } from 'lucide-react';

interface OutputPanelProps {
  output: string;
  isExecuting: boolean;
}

const OutputPanel: React.FC<OutputPanelProps> = ({ output, isExecuting }) => {
  const getStatusIcon = () => {
    if (isExecuting) {
      return <Loader className="w-4 h-4 text-blue-500 animate-spin" />;
    }
    if (output.includes('Error:')) {
      return <XCircle className="w-4 h-4 text-red-500" />;
    }
    if (output.trim()) {
      return <CheckCircle className="w-4 h-4 text-green-500" />;
    }
    return <Terminal className="w-4 h-4 text-gray-400" />;
  };

  const getStatusText = () => {
    if (isExecuting) {
      return 'Executing...';
    }
    if (output.includes('Error:')) {
      return 'Error';
    }
    if (output.trim()) {
      return 'Success';
    }
    return 'Ready';
  };

  const getStatusColor = () => {
    if (isExecuting) {
      return 'text-blue-600';
    }
    if (output.includes('Error:')) {
      return 'text-red-600';
    }
    if (output.trim()) {
      return 'text-green-600';
    }
    return 'text-gray-600';
  };

  return (
    <div className="w-96 bg-white border-l border-gray-200 flex flex-col">
      {/* Output Header */}
      <div className="flex items-center justify-between p-4 border-b border-gray-200 bg-gray-50">
        <div className="flex items-center space-x-2">
          <Terminal className="w-5 h-5 text-gray-600" />
          <h3 className="text-lg font-semibold text-gray-900">
            Output
          </h3>
        </div>
        
        <div className="flex items-center space-x-2">
          {getStatusIcon()}
          <span className={`text-sm font-medium ${getStatusColor()}`}>
            {getStatusText()}
          </span>
        </div>
      </div>

      {/* Output Content */}
      <div className="flex-1 p-4 overflow-y-auto">
        {isExecuting ? (
          <div className="flex items-center justify-center h-full">
            <div className="text-center">
              <Loader className="w-8 h-8 text-blue-500 animate-spin mx-auto mb-2" />
              <p className="text-gray-600">Running your Go code...</p>
            </div>
          </div>
        ) : output ? (
          <div className="space-y-2">
            <pre className="text-sm font-mono text-gray-800 whitespace-pre-wrap break-words">
              {output}
            </pre>
          </div>
        ) : (
          <div className="flex items-center justify-center h-full text-gray-500">
            <div className="text-center">
              <Terminal className="w-12 h-12 mx-auto mb-4 text-gray-300" />
              <p className="text-lg font-medium mb-2">No output yet</p>
              <p className="text-sm">
                Run your Go code to see the results here
              </p>
            </div>
          </div>
        )}
      </div>

      {/* Output Footer */}
      <div className="p-3 border-t border-gray-200 bg-gray-50 text-xs text-gray-500">
        <div className="flex items-center justify-between">
          <span>Go Runtime Output</span>
          <div className="flex items-center space-x-2">
            <span>Lines: {output.split('\n').length}</span>
            <span>•</span>
            <span>Characters: {output.length}</span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default OutputPanel;
