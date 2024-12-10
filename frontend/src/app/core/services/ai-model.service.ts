import { Injectable } from '@angular/core';
import { AIModel } from '../models/commit.model';

@Injectable({
  providedIn: 'root'
})
export class AIModelService {
  constructor() {}

  async getAvailableModels(): Promise<AIModel[]> {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 800));

    // Mock response
    return [
      { id: 'gpt3', name: 'GPT-3', isOnline: true },
      { id: 'gpt4', name: 'GPT-4', isOnline: true },
      { id: 'codex', name: 'Codex', isOnline: true },
      { id: 'localai', name: 'Local AI', isOnline: false },
      { id: 'customai', name: 'Custom AI', isOnline: false },
    ];
  }
}

