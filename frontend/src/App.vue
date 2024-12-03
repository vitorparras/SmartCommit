<template>
  <div class="app-container">
    <h1 class="text-2xl font-bold mb-4">Gerador de Mensagens de Commit</h1>
    <SelectorComponent />
    <MessageDisplayComponent />
  </div>
</template>

<script>
import SelectorComponent from './components/Config/SelectorComponent.vue';
import MessageDisplayComponent from './components/Results/MessageDisplayComponent.vue';

export default {
  components: {
    SelectorComponent,
    MessageDisplayComponent,
  },
};
</script>

<style>
.app-container {
  padding: 2rem;
  font-family: Arial, sans-serif;
}
</style>
