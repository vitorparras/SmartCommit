import { defineStore } from 'pinia';

export const useConfigStore = defineStore('config', {
  state: () => ({
    commitPattern: 'conventional',
  }),
  actions: {
    setPattern(pattern) {
      this.commitPattern = pattern;
    },
  },
});
