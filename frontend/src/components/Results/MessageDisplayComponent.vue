<template>
    <div class="message-display">
      <h2 class="font-bold mb-2">Mensagem Gerada:</h2>
      <p v-if="isLoading" class="text-gray-500">Gerando mensagem...</p>
      <p v-else-if="generatedMessage" class="border p-4 rounded bg-gray-50">
        {{ generatedMessage }}
      </p>
      <p v-else class="text-gray-500">Selecione um padrão e gere uma mensagem.</p>
    </div>
  </template>
  
  <script>
  import { ref, watch } from 'vue';
  import { useConfigStore } from '@/stores/configStore';
  import mockApi from '@/mockApi';
  
  export default {
    setup() {
      const configStore = useConfigStore();
      const generatedMessage = ref('');
      const isLoading = ref(false);
  
      const fetchMessage = async () => {
        isLoading.value = true;
        const response = await mockApi.generateCommitMessage(configStore.commitPattern);
        generatedMessage.value = response.message;
        isLoading.value = false;
      };
  
      // Monitora mudanças no padrão e atualiza a mensagem
      watch(
        () => configStore.commitPattern,
        () => {
          fetchMessage();
        },
        { immediate: true }
      );
  
      return {
        generatedMessage,
        isLoading,
      };
    },
  };
  </script>
  
  <style scoped>
  .message-display {
    margin-top: 2rem;
  }
  </style>
  