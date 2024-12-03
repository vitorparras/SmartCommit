<template>
    <div class="selector">
      <label for="commitPattern" class="block font-medium">Selecione o Padrão de Commit:</label>
      <select
        v-model="commitPattern"
        id="commitPattern"
        class="border p-2 rounded"
      >
        <option value="conventional">Conventional Commits</option>
        <option value="gitmoji">Gitmoji</option>
        <option value="angular">AngularJS</option>
        <option value="semantic">Semantic Commit Messages</option>
        <option value="custom">Custom Commit Messages</option>
        <option value="scalable">Scalable Commit Messages</option>
      </select>
  
      <p class="mt-4 text-sm text-gray-600">
        <strong>Exemplo:</strong> {{ example }}
      </p>
    </div>
  </template>
  
  <script>
  import { computed } from 'vue'; // Importando computed
  import { useConfigStore } from '@/stores/configStore'; // Store
  
  export default {
    setup() {
      const configStore = useConfigStore(); // Acessando a store
  
      const examples = {
        conventional: 'feat(ui): add button for dark mode',
        gitmoji: '✨ Add new feature for dark mode',
        angular: 'fix(auth): resolve login timeout issue',
        semantic: 'Added: New API endpoint for retrieving user data',
        custom: '[FEATURE] Implement user registration (#123)',
        scalable: 'auth/login: fix token expiration issue',
      };
  
      return {
        commitPattern: computed({
          get: () => configStore.commitPattern, // Vinculando v-model ao estado global
          set: (value) => (configStore.commitPattern = value),
        }),
        example: computed(() => examples[configStore.commitPattern]), // Mostrando exemplo baseado no padrão
      };
    },
  };
  </script>
  
  <style scoped>
  .selector {
    margin-bottom: 1rem;
  }
  </style>
  