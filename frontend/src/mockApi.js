const mockApi = {
    generateCommitMessage: async (pattern) => {
      const exampleMessages = {
        conventional: 'feat(ui): add dark mode toggle',
        gitmoji: '✨ Add new feature for dark mode',
        angular: 'feat(core): add dark mode toggle',
        semantic: 'Added: New API endpoint for retrieving user data',
        custom: '[FEATURE] Implement user registration (#123)',
        scalable: 'ui/theme: add dark mode toggle',
      };
  
      // Simula um atraso para replicar um backend real
      await new Promise((resolve) => setTimeout(resolve, 500));
  
      return { message: exampleMessages[pattern] || 'Default message' };
    },
  };
  
  export default mockApi;
  