import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ApiKeyService {
  private readonly STORAGE_KEY = 'ai-model-api-keys';

  getApiKey(modelId: string): string | null {
    const keys = this.getAllApiKeys();
    return keys[modelId] || null;
  }

  saveApiKey(modelId: string, apiKey: string): void {
    const keys = this.getAllApiKeys();
    keys[modelId] = apiKey;
    localStorage.setItem(this.STORAGE_KEY, JSON.stringify(keys));
  }

  private getAllApiKeys(): Record<string, string> {
    const stored = localStorage.getItem(this.STORAGE_KEY);
    return stored ? JSON.parse(stored) : {};
  }
}

